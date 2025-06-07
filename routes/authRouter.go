package routes

import (
	"Movie-Management-System/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("users/signup", controllers.Signup())

}
