package router

import (
	"gobook/app/controllers"
	"gobook/app/middlewares"

	"github.com/gin-gonic/gin"
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

		//Get Profile
		v1.GET("/user/profile", controllers.GetProfile)
	}

	//Running server in port 8080
	r.Run(":8080")
}
