package controllers

import (
	"gobook/app/database"
	"gobook/app/middlewares"
	"gobook/app/models"
	"gobook/app/response"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	//Get username from tokens
	username := middlewares.GetUsername()
	DB := database.GetDB()

	var user models.User
	err := DB.Preload("Profile").Where("username = ?", username).Find(&user).Error

	if err != nil {
		response.NotFound(c, "User")
		return
	}

	response.Send(c, 200, "Successfully get profile", user)
}
