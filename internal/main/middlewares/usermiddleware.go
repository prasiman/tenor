package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"main/internal/main/models"
	"main/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		// Validate jwt token
		bearerToken := strings.Split(header, "Bearer ")

		if len(bearerToken) != 2 {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "No Bearer token found"})
			return
		}

		id, err := jwt.ParseToken(bearerToken[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err})
			return
		}

		// Check if user exists in db
		user, err := models.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}
