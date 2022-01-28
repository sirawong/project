package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cinema/config"
	"cinema/handler"
	"cinema/middleware"
	"cinema/repository/grpc"
	repository "cinema/repository/mongodb"
	service "cinema/service/cinema/implement"
	grpcServer "cinema/service/grpc"
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

	grpcAuthRepo := grpc.New(&grpc.Config{Network: "tcp", Port: appConfig.GRPCAuthHost})
	grpcService := grpcClient.New(grpcAuthRepo)

	middlewareService := middleware.New(grpcService)

	cinemaDB := repository.New(dbConn, appConfig)
	cinemaSRV := service.New(cinemaDB, uuid, appConfig)

	go grpcServer.NewServer(cinemaSRV, appConfig)

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
