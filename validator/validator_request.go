package validator

import (
	"github.com/gin-gonic/gin"
)

//Validation user request
func Validate(c *gin.Context, req interface{}) error {
	return c.ShouldBindJSON(req)
}
