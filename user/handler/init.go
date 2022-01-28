package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"user/middleware"
	"user/service/auth"
	"user/service/user"

	authController "user/handler/auth"
	userController "user/handler/user"
)

type Handlers struct {
	userService user.Service
	authService auth.Service
	middleware  middleware.Service
	userCtrl    *userController.Controller
	authCtrl    *authController.Controller
}

func New(user user.Service, auth auth.Service, middleware middleware.Service) *Handlers {
	return &Handlers{
		userService: user,
		authService: auth,
		middleware:  middleware,
		userCtrl:    userController.New(user),
		authCtrl:    authController.New(auth),
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

	user := router.Group("users")
	{
		user.POST("", app.userCtrl.Create)
		user.POST("login", app.authCtrl.Login)
		user.POST("photo/:id", app.userCtrl.Upload)

		simple := user.Group("", app.middleware.Simple())
		{
			simple.POST("logout", app.authCtrl.Logout)
			simple.GET("me", app.userCtrl.GetMe)
			simple.PATCH("me", app.userCtrl.UpdateMe)
			simple.DELETE("me", app.userCtrl.DeleteMe)
		}

		enhance := user.Group("", app.middleware.Enhance())
		{
			enhance.POST("logoutALL", app.authCtrl.LogoutAll)
			enhance.GET("", app.userCtrl.All)
			enhance.GET(":id", app.userCtrl.Read)
			enhance.PATCH(":id", app.userCtrl.Update)
			enhance.DELETE(":id", app.userCtrl.Delete)
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
