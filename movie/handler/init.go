package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	movieController "movie/handler/movie"
	"movie/middleware"
	movieService "movie/service/movie"
)

type Handlers struct {
	service    movieService.Service
	middleware middleware.Service
	movieCtrl  *movieController.Handlers
}

func New(service movieService.Service, middleware middleware.Service) *Handlers {
	return &Handlers{
		service:    service,
		middleware: middleware,
		movieCtrl:  movieController.New(service),
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

	api := router.Group("movies")
	{
		api.GET("", app.movieCtrl.All)
		api.GET(":id", app.movieCtrl.Read)
		
		enhance := api.Group("", app.middleware.Enhance())
		{
			enhance.POST("", app.movieCtrl.Create)
			enhance.GET("photo/:id", app.movieCtrl.Upload)
			enhance.PUT(":id", app.movieCtrl.Update)
			enhance.DELETE(":id", app.movieCtrl.Delete)
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
