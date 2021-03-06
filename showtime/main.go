package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"showtime/config"
	handler "showtime/handler"
	"showtime/middleware"
	"showtime/repository/grpc"
	repository "showtime/repository/mongodb"
	grpcClient "showtime/service/grpcClient/implement"
	service "showtime/service/showtime/implement"
	"showtime/utils"
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

	grpcRepo := grpc.New(&grpc.Config{Network: "tcp", Port: appConfig.GRPCAuthHost})
	grpcService := grpcClient.New(grpcRepo)
	middlewareService := middleware.New(grpcService)

	movieDB := repository.New(dbConn, appConfig)
	movieSRV := service.New(movieDB, uuid)

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
