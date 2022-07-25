package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sshaunn/train-go/ms1/controller"
	"github.com/sshaunn/train-go/ms1/middleware"
)

func Routers() {
	router := gin.New()
	router.Use(middleware.Cors())
	router.POST("api/v1/upload", controller.Upload)
	router.Run("localhost:8080")
	// return router
}
