package user

import (
	"github.com/bojie/animal/backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/register", register())
	incomingRoutes.POST("users/login", login())
	incomingRoutes.GET("/validate", middleware.Authenticate, validate())
	incomingRoutes.POST("/auth", refresh())
}
