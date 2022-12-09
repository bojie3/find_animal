package main

import (
	"github.com/bojie/animal/backend/user"
	"github.com/bojie/animal/backend/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	user.UserRoutes(router)

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"error":"no such api"}) })

	// Start listening and serving requests
	router.Run(":8080")

	defer db.DB.Close()
}
