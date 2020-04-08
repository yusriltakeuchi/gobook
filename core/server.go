package core

import (
	"fmt"
	"os/exec"

	"github.com/yusriltakeuchi/gobook/app/database"
	"github.com/yusriltakeuchi/gobook/config"
	"github.com/yusriltakeuchi/gobook/router"
)

func Start() {

	//Setup router
	router.SetupRouter()

	defer database.CloseDB()
}

func InstallLibrary() {
	for _, value := range config.GetPackages() {
		cmd := exec.Command("go", "get", value)
		err := cmd.Run()
		var message string

		if err != nil {
			message = fmt.Sprintf(" -> Error Installing %s.\n%s", value, err.Error())
			fmt.Println(message)
			continue
		}
		message = fmt.Sprintf(" -> Successfully Install %s.", value)
		fmt.Println(message)
	}
	fmt.Println(" -> Required library successfully installed.")
}
