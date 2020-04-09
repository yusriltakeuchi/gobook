package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yusriltakeuchi/gobook/app/database"
	"github.com/yusriltakeuchi/gobook/app/middlewares"
	"github.com/yusriltakeuchi/gobook/app/models"
	"github.com/yusriltakeuchi/gobook/app/response"
	"github.com/yusriltakeuchi/gobook/validator"
)

//Function to login and get access token
func Login(c *gin.Context) {
	//Data mapping & validatio
	var req models.LoginValidate
	err := validator.Validate(c, &req)
	if err != nil {
		response.BadValidator(c, err.Error())
		return
	}

	//model to map from query result
	var user models.User
	DB := database.GetDB()

	//Binding user request json
	c.BindJSON(&req)

	//Find user
	err = DB.Where("username = ?", req.Username).Find(&user).Error
	if err != nil {
		response.NotFound(c, "User")
		return
	}

	//Validate password
	valid := middlewares.CheckPasswordHash(req.Password, user.Password)
	if valid == false {
		response.BadCredentials(c)
		return
	}

	//Generate Access token
	token, exp, err := middlewares.GenerateToken(user.Username, user.UniqueKey)
	if err != nil {
		response.InternalError(c)
		return
	}
	response.Send(c, 200, "Successfully login", gin.H{
		"token":        token,
		"expired_time": exp})
}

//Function to register new users
func Register(c *gin.Context) {
	//Data mapping & validations
	var req models.RegisterValidate
	err := validator.Validate(c, &req)
	if err != nil {
		response.BadValidator(c, err.Error())
		return
	}

	//model to map from query result
	DB := database.GetDB()

	//Binding user request json
	c.BindJSON(&req)

	//Hashing password
	password, _ := middlewares.HashPassword(req.Password)

	//Generate random string for token
	uniqueKey := middlewares.RandomStr(10)

	//Data to save
	user := models.User{
		Username:  req.Username,
		Password:  password,
		Email:     req.Email,
		UniqueKey: uniqueKey,
		Profile: models.Profile{
			Name:  req.Name,
			Phone: req.Phone,
			City:  req.City}}

	//find existing user
	err = DB.Where("username = ?", req.Username).Find(&user).Error
	if err == nil {
		response.Exists(c, "User")
		return
	}

	//If not exists then create user
	DB.Create(&user)
	response.Send(c, 201, "Successfully registerd", &user)
	return
}

//Function to logout and delete old authorizations
func Logout(c *gin.Context) {
	//Get username from tokens
	username := middlewares.GetUsername()
	DB := database.GetDB()

	var user models.User

	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		response.NotFound(c, "User")
		return
	}

	//Generate new uniqueKey
	user.UniqueKey = middlewares.RandomStr(10)
	DB.Save(&user)
	response.Send(c, 200, "Successfully logout", nil)
}
