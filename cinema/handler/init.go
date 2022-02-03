package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	cinemaController "cinema/handler/cinema"
	"cinema/middleware"
	service "cinema/service/cinema"
)

type Handlers struct {
	service    service.CinemaService
	middleware middleware.Service
	cinemaCtrl *cinemaController.Handlers
}

func New(service service.CinemaService, middleware middleware.Service) *Handlers {
	return &Handlers{service: service, middleware: middleware, cinemaCtrl: cinemaController.New(service)}
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

	api := router.Group("cinemas")
	{
		api.GET("", app.cinemaCtrl.All)
		api.GET(":id", app.cinemaCtrl.Read)
		api.POST("photo/:id", app.cinemaCtrl.Upload)

		// enhance := api.Group("", app.middleware.Enhance())
		{
			api.POST("", app.cinemaCtrl.Create)
			api.PATCH(":id", app.cinemaCtrl.Update)
			api.DELETE(":id", app.cinemaCtrl.Delete)
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
