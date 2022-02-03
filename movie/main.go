package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"movie/config"
	"movie/handler"
	"movie/logs"
	"movie/middleware"
	"movie/repository/grpc"
	gStorageDB "movie/repository/gstorage"
	repository "movie/repository/mongodb"
	grpcClient "movie/service/grpcClient/implement"
	service "movie/service/movie/implement"
	"movie/utils"
)

func main() {
	appConfig := config.Get()

	ctx := context.Background()
	dbConn := initDatabase(ctx, appConfig)
	defer func() {
		dbConn.Disconnect(ctx)
	}()

	uuid, err := utils.NewUUID()
	if err != nil {
		panic(err)
	}

	storageClient := initStorage(ctx)
	defer func() {
		storageClient.Close()
	}()

	storage := gStorageDB.New(storageClient, appConfig)
	grpcRepo := grpc.New(&grpc.Config{Network: "tcp", Port: appConfig.GRPCAuthHost})
	grpcService := grpcClient.New(grpcRepo)
	middlewareService := middleware.New(grpcService)

	movieDB := repository.New(dbConn, appConfig)
	movieSRV := service.New(movieDB, uuid, appConfig, storage)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	_ = handler.New(movieSRV, middlewareService).RegisterRoutes(router)
	router.Run(appConfig.AppPort)
}

func initDatabase(ctx context.Context, appConfig *config.Config) *mongo.Client {
	dsn := fmt.Sprintf("%v", appConfig.MongoDBEndpoint)

	db, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	return db
}

func initStorage(ctx context.Context) *storage.Client {
	// storageClient, err := storage.NewClient(ctx, setUpStorage())
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	return storageClient
}

// func setUpStorage() gOption.ClientOption {
// 	if _, err := os.Stat("config/key.json"); errors.Is(err, os.ErrNotExist) {
// 		return gOption.WithoutAuthentication()
// 	}
// 	return gOption.WithCredentialsFile("config/key.json")
// }
