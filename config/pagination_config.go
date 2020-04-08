package config

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Define controller constants
const PaginationLimit int = 10
const PaginationOffset int = 10

//Function to get pagination number
func Page(c *gin.Context) int {
	page := c.Query("p")
	var p int
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 0
	}
	return p
}
