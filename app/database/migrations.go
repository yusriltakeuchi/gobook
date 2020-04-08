package database

import (
	"fmt"

	"github.com/yusriltakeuchi/gobook/app/models"
)

func Migrate() {

	///Initializing database
	Init()

	//Migrations database
	db.AutoMigrate(&models.Books{})
	fmt.Println(" -> Successfully migrate Books Model")

	db.AutoMigrate(&models.User{})
	fmt.Println(" -> Successfully migrate User Model")
}
