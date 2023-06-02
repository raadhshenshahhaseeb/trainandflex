package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hyperversalblocks/trainandflex/pkg/jwt"
)

func (s *server) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		userinfo, err := jwt.ValidateToken(token, s.config.JwtSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "token expired")
			return
		}

		ctx.Set(jwt.KeyUserID, userinfo.UID)
		ctx.Next()
	}
}

func (s *server) ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO handle all errors?
	if c.Errors == nil || len(c.Errors) == 0 || c.Writer.Status() != http.StatusInternalServerError {
		return
	}

	for _, err := range c.Errors {
		s.logger.Error(err)
	}

	c.JSON(http.StatusInternalServerError, "internal server error")
}
