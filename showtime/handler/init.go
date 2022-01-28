package showtime

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	showtimeController "showtime/handler/showtime"
	"showtime/middleware"
	service "showtime/service/showtime"
)

type Handlers struct {
	service      service.Service
	middleware   middleware.Service
	showtimeCtrl *showtimeController.Handlers
}

func New(service service.Service, middleware middleware.Service) *Handlers {
	return &Handlers{
		service:      service,
		middleware:   middleware,
		showtimeCtrl: showtimeController.New(service),
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

	api := router.Group("showtimes")
	{
		api.GET("", app.showtimeCtrl.All)
		api.GET(":id", app.showtimeCtrl.Read)

		enhance := api.Group("", app.middleware.Enhance())
		{
			enhance.POST("", app.showtimeCtrl.Create)
			enhance.PATCH(":id", app.showtimeCtrl.Update)
			enhance.DELETE(":id", app.showtimeCtrl.Delete)
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
