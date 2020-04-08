package core

import (
	"fmt"

	"github.com/yusriltakeuchi/gobook/app/database"
)

func ShowHelp() {
	fmt.Println("----=====[ Gobook ]=====----")
	fmt.Println(" -> go run main.go start - Start server with port 8080")
	fmt.Println(" -> go run main.go migrate - Migrate database")
	fmt.Println(" -> go run main.go install - Installing all required library")
}

func Command(cmd []string) {

	if len(cmd) == 0 {
		ShowHelp()
		return
	}

	switch cmd[0] {
	case "start":
		Start()
	case "migrate":
		database.Migrate()
	case "install":
		InstallLibrary()
	case "help":
		ShowHelp()
	}
}
