package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"showtime/logs"
	pbAuth "showtime/service/grpcClient/protobuf/auth"
)

func (middleware Service) Simple() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.ReplaceAll(header, "Bearer ", "")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": "Please authenticate."})
			return
		}

		data := &pbAuth.TokenRequest{
			Id:    fmt.Sprint(time.Now().Unix()),
			Token: token,
		}

		grpcResponse, err := middleware.GRPCSrv.VerifyToken(data)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !grpcResponse.Verify {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userid", grpcResponse.Id)

	}
}

func (middleware Service) Enhance() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.ReplaceAll(header, "Bearer ", "")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": "Please authenticate."})
			return
		}

		data := &pbAuth.TokenRequest{
			Id:    fmt.Sprint(time.Now().Unix()),
			Token: token,
		}

		grpcResponse, err := middleware.GRPCSrv.VerifyToken(data)
		if err != nil {
			logs.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !grpcResponse.Verify {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if grpcResponse.Role != "superadmin" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userid", grpcResponse.Id)

	}
}
