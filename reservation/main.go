package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"reservation/config"
	"reservation/handler"
	"reservation/middleware"
	"reservation/repository/grpc"
	repositoryDB "reservation/repository/mongodb"
	repositorySendGrid "reservation/repository/sendgrid"
	grpcClient "reservation/service/grpcClient/implement"
	invitationService "reservation/service/invitation/implement"
	reservationService "reservation/service/reservation/implement"
	"reservation/utils"
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
	grpcCinemaRepo := grpc.New(&grpc.Config{Network: "tcp", Port: appConfig.GRPCCinemaHost})
	grpcService := grpcClient.New(grpcAuthRepo, grpcCinemaRepo)

	middlewareService := middleware.New(grpcService)

	reservationDB := repositoryDB.New(dbConn, appConfig)
	reservationSRV := reservationService.New(reservationDB, uuid, grpcService, appConfig)

	sendGridDB := repositorySendGrid.New(appConfig.SendgridAPIKey)
	invitationSRV := invitationService.New(sendGridDB)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	_ = handler.New(reservationSRV, invitationSRV, middlewareService).RegisterRoutes(router)
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
