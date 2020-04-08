package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"github.com/yusriltakeuchi/gobook/config"
	"github.com/yusriltakeuchi/gobook/db"
	"github.com/yusriltakeuchi/gobook/models"
	"github.com/yusriltakeuchi/gobook/response"
	"github.com/yusriltakeuchi/gobook/validator"
)

//Function to create book
func CreateBook(c *gin.Context) {
	//Data mapping & validations
	var req models.BooksValidate
	err := validator.Validate(c, &req)
	if err != nil {
		response.BadValidator(c, err.Error())
		return
	}

	var book models.Books
	DB := db.GetDB()

	//Binding user request json
	c.BindJSON(&req)
	//Copy from request to book model
	deepcopier.Copy(req).To(&book)

	//Find existing data
	err = DB.Where("title = ?", book.Title).Find(&book).Error
	//If data found
	if err == nil {
		response.Exists(c, "Book")
		return
	}

	//If book not exists then create
	DB.Create(&book)
	response.Send(c, 201, "Successfully create book", book)
	return
}

//Function to delete book with specific id
func DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")

	var book models.Books
	DB := db.GetDB()

	//Find book with id
	err := DB.Where("id = ?", id).Find(&book).Error
	//If data not found
	if err != nil {
		response.NotFound(c, "Book")
		return
	}

	//If found then delete
	DB.Delete(&book)
	response.Send(c, 200, "Successfully delete book", nil)
	return
}

//Function to update book
func UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")

	//Data mapping & validations
	var req models.BooksValidate
	err := validator.Validate(c, &req)
	if err != nil {
		response.BadValidator(c, err.Error())
		return
	}

	var book models.Books
	DB := db.GetDB()

	//Binding user request json
	c.BindJSON(&req)

	//Find book with specific id
	err = DB.Where("id = ?", id).First(&book).Error
	//If book not found
	if err != nil {
		response.NotFound(c, "Book")
		return
	}

	//Copy from request to book model
	deepcopier.Copy(req).To(&book)
	DB.Save(&book)
	response.Send(c, 200, "Successfully update book", book)
}

//Function to get all books
func GetBooks(c *gin.Context) {
	var books []models.Books
	DB := db.GetDB()

	//Paginations the data
	p := config.Page(c)
	//Find data with limit offset
	err := DB.Limit(config.PaginationLimit).Offset(p * config.PaginationLimit).Find(&books).Error

	//If error then show 404
	if err != nil {
		response.NotFound(c, "Books")
		return
	} else {
		response.Send(c, 200, "Successfully get books", books)
		return
	}
}
