package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yusriltakeuchi/gobook/models"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	dbUsername, dbPassword, dbName := "root", "62569621144", "restgorm"
	var connectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbName)

	db, _ = gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err)
	}

	//Migrations database
	db.AutoMigrate(&models.Books{})
	db.AutoMigrate(&models.User{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
