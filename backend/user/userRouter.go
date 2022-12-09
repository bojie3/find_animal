package user

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", signup())
	incomingRoutes.POST("users/signin", signin())
}
