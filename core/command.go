package core

import (
	"fmt"

	"gobook/app/database"
)

func ShowHelp() {
	fmt.Println("----=====[ Gobook ]=====----")
	fmt.Println(" -> gobook start - Start server with port 8080")
	fmt.Println(" -> gobook migrate - Migrate database")
	fmt.Println(" -> gobook install <packagename> - Installing required library")
	fmt.Println(" -> gobook make controller <name> - Create new controllers")
	fmt.Println(" -> gobook make model <name> - Create new models")
}

func Command(cmd []string) {

	if len(cmd) == 0 {
		Start()
		return
	}

	switch cmd[0] {
	case "start":
		Start()
	case "migrate":
		database.Migrate()
	case "install":
		if len(cmd) == 2 {
			fmt.Println(fmt.Sprintf(" -> Inserting library %s.", cmd[1]))
			AddingLibrary(cmd[1])
		}
		InstallLibrary()
	case "make":
		GenerateCode(cmd)
	case "help":
		ShowHelp()
	}
}
