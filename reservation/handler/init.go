package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"reservation/middleware"

	reservController "reservation/handler/reservation"
	reservationService "reservation/service/reservation"
)

type Handlers struct {
	reservService reservationService.Service
	middleware    middleware.Service
	reservCtrl    *reservController.Handlers
}

func New(reservationService reservationService.Service, middleware middleware.Service) *Handlers {
	return &Handlers{
		reservService: reservationService,
		middleware:    middleware,
		reservCtrl:    reservController.New(reservationService),
	}
}

func (app *Handlers) RegisterRoutes(router *gin.Engine) *Handlers {
	// docs.SwaggerInfo.Title = "Authorization Service API"
	// docs.SwaggerInfo.Description = "API Spec Authorization Service."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = app.swaggerHost
	// docs.SwaggerInfo.BasePath = app.basePath
	// docPath := ginSwagger.URL(fmt.Sprintf("//%s/swagger/doc.json", app.swaggerHost))

	router.Use(app.corsMiddleware())

	router.GET("/system/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	api := router.Group("reservations")
	{
		api.GET(":id", app.reservCtrl.Read)
		api.GET("checkin/:id", app.reservCtrl.Checkin)

		// simple := api.Group("", app.middleware.Simple())
		{
			api.POST("", app.reservCtrl.Create)
			api.GET("", app.reservCtrl.All)
		}
		// enhance := api.Group("", app.middleware.Enhance())
		{
			api.PATCH(":id", app.reservCtrl.Update)
			api.DELETE(":id", app.reservCtrl.Delete)
		}
	}

	return app

}

func (app *Handlers) corsMiddleware() gin.HandlerFunc {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowCredentials = true
	corsConf.AddAllowHeaders("Authorization", "User-Agent")
	return cors.New(corsConf)
}
