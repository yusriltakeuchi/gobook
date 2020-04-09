package database

import (
	"fmt"
	"github.com/yusriltakeuchi/gobook/app/models"
)

//Insert your model here
func GetModel() [][]interface{} {
	return [][]interface{}{
		{&models.Books{}, "Books"},
		{&models.User{}, "User"}}
}

func Migrate() {
	///Initializing database
	Init()

	//Migrations database
	for _, value := range GetModel() {
		exists := db.HasTable(value[0])
		if exists == false {
			db.AutoMigrate(value[0])
			fmt.Println(fmt.Sprintf(" -> Successfully migrate %s Model", value[1]))
		} else {
			fmt.Println(fmt.Sprintf(" -> Migration %s already defined", value[1]))
		}
	}
}
