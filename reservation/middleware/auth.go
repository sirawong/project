package middleware

import (
	"fmt"
	"net/http"
	pbAuth "reservation/service/grpcClient/protobuf/auth"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
