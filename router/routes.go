package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yusriltakeuchi/gobook/app/controllers"
	"github.com/yusriltakeuchi/gobook/app/middlewares"
)

func SetupRouter() {
	r := gin.Default()
	v1 := r.Group("/v1")

	//Authentications
	v1.POST("/auth/login", controllers.Login)
	v1.POST("/auth/register", controllers.Register)

	v1.Use(middlewares.AuthHandler())
	{
		//Books Group
		v1.GET("/books", controllers.GetBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.PUT("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)

		//Authentication for logout
		v1.POST("/auth/logout", controllers.Logout)
	}

	//Running server in port 8080
	r.Run(":8080")
}
