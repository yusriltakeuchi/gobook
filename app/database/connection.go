package database

import (
	"fmt"

	"github.com/yusriltakeuchi/gobook/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	var connectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.LoadEnv("DB_USERNAME"), config.LoadEnv("DB_PASSWORD"), config.LoadEnv("DB_NAME"))

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
