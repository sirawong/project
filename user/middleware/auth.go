package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (middleware Service) Simple() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		encodedToken := strings.ReplaceAll(header, "Bearer ", "")
		if encodedToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": "Please authenticate."})
			return
		}

		id, _, err := middleware.auth.Verify(c, encodedToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("token", encodedToken)
		c.Set("userid", id)
	}
}

func (middleware Service) Enhance() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		encodedToken := strings.ReplaceAll(header, "Bearer ", "")
		if encodedToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": "Please authenticate."})
			return
		}

		id, user, err := middleware.auth.Verify(c, encodedToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.Role != "superadmin" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("token", encodedToken)
		c.Set("userid", id)
		c.Set("role", user.Role)

	}
}
