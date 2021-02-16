package main

import (
	controller "hw/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {

	authorized := router.Group("/")
	authorized.Use(controller.AuthRequired)
	{
		authorized.GET("/user", controller.GetUserAction)
		authorized.GET("/user/:id", controller.GetUserByIdAction)
		authorized.POST("/user", controller.InsertUserAction)
		authorized.PUT("/user/:id", controller.InserActionById)
		authorized.DELETE("/user/:id", controller.DeleteUser)
	}
	router.POST("/login", controller.LoginUser)
}
