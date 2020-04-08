package response

import (
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context, item string) {
	c.AbortWithStatusJSON(404, gin.H{
		"status":  404,
		"message": item + " not found"})
}

func InternalError(c *gin.Context) {
	c.AbortWithStatusJSON(500, gin.H{
		"status":  500,
		"message": "Internal server error"})
}

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  401,
		"message": "Your request is not authorized"})
}

func InvalidToken(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  401,
		"message": "Invalid authorization token"})
}

func Exists(c *gin.Context, item string) {
	c.AbortWithStatusJSON(409, gin.H{
		"status":  409,
		"message": item + " already exists"})
}

func BadCredentials(c *gin.Context) {
	c.AbortWithStatusJSON(401, gin.H{
		"status":  401,
		"message": "Username or password is wrong"})
}

func BadValidator(c *gin.Context, err interface{}) {
	c.AbortWithStatusJSON(422, gin.H{
		"status":  422,
		"message": err})
}

func Send(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data})
}
