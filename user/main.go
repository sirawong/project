package main

import (
	"context"
	"fmt"

	storage "cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"user/config"
	"user/handler"
	"user/logs"
	"user/middleware"
	gStorageDB "user/repository/gstorage"
	repository "user/repository/mongodb"
	authService "user/service/auth/implement"
	grpcServer "user/service/grpc"
	userService "user/service/user/implement"
	"user/utils"
)

func main() {
	appConfig := config.Get()
	ctx := context.Background()
	dbConn := initDatabase(ctx, appConfig)
	defer func() {
		dbConn.Disconnect(ctx)
	}()

	storageClient := initStorage(ctx)
	defer func() {
		storageClient.Close()
	}()

	uuid, err := utils.NewUUID()
	if err != nil {
		panic(err)
	}

	jwt := utils.NewJWT(appConfig)

	storage := gStorageDB.New(storageClient, appConfig)
	userDB := repository.New(dbConn, appConfig)
	authSrv := authService.New(userDB, jwt)
	userSrv := userService.New(userDB, authSrv, uuid, appConfig, storage)
	midSrv := middleware.New(authSrv)

	go grpcServer.NewServer(authSrv, appConfig)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	_ = handler.New(userSrv, authSrv, midSrv).RegisterRoutes(router)
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
