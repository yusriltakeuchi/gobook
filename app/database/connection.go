package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
