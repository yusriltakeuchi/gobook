package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yusriltakeuchi/gobook/app/database"
	"github.com/yusriltakeuchi/gobook/app/middlewares"
	"github.com/yusriltakeuchi/gobook/app/models"
	"github.com/yusriltakeuchi/gobook/app/response"
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
