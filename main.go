package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"github.com/yusriltakeuchi/gobook/db"
	"github.com/yusriltakeuchi/gobook/router"
)

func main() {
	//Initializing database
	db.Init()

	//Setup router
	router.SetupRouter()

	defer db.CloseDB()
}
