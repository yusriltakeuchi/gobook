package middlewares

import (
	"fmt"
	"strings"

	"github.com/yusriltakeuchi/gobook/app/database"
	"github.com/yusriltakeuchi/gobook/app/models"
	"github.com/yusriltakeuchi/gobook/app/response"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var username string

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer: "
		if !strings.Contains(token, b) {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			response.InvalidToken(c)
			c.Abort()
			return
		}
		// Validate token
		valid, err := ValidateToken(t[1], SigningKey)
		if err != nil {
			response.InvalidToken(c)
			c.Abort()
			return
		}

		//Validate user is found in database
		DB := database.GetDB()
		var user models.User

		//Get username
		username = fmt.Sprintf("%v", valid.Claims.(jwt.MapClaims)["username"])

		//Find user
		err = DB.Where("username = ?", username).Find(&user).Error
		if err != nil {
			response.Unauthorized(c)
			return
		}

		//Validate uniqueKey
		//Get uniqueKey from token
		uniqueKey := valid.Claims.(jwt.MapClaims)["unique_key"]
		if user.UniqueKey != uniqueKey {
			response.Unauthorized(c)
			return
		}

		// set userId Variable
		c.Set("username", username)
		c.Next()
	}
}

func GetUsername() string {
	return username
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
