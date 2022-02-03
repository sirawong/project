package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinema/config"
	"cinema/handler"
	"cinema/logs"
	"cinema/middleware"
	"cinema/repository/grpc"
	gStorageDB "cinema/repository/gstorage"
	repository "cinema/repository/mongodb"
	service "cinema/service/cinema/implement"
	grpcClient "cinema/service/grpcClient/implement"
	"cinema/utils"
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
	grpcAuthRepo := grpc.New(&grpc.Config{Network: "tcp", Port: appConfig.GRPCAuthHost})
	grpcService := grpcClient.New(grpcAuthRepo)

	middlewareService := middleware.New(grpcService)

	cinemaDB := repository.New(dbConn, appConfig)
	cinemaSRV := service.New(cinemaDB, uuid, appConfig, storage)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	_ = handler.New(cinemaSRV, middlewareService).RegisterRoutes(router)
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
