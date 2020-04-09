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
		//Check if library already installed
		cmd := exec.Command("go", "list", value)
		err := cmd.Run()

		//If package not installed then install
		if err != nil {
			cmd = exec.Command("go", "get", value)
			err = cmd.Run()

			if err != nil {
				fmt.Println(fmt.Sprintf(" -> Error Installing %s.\n%s", value, err.Error()))
				continue
			}
			fmt.Println(fmt.Sprintf(" -> Successfully Install %s.", value))
		} else {
			fmt.Println(fmt.Sprintf(" -> Library %s already installed.", value))
		}
	}
	fmt.Println(" -> Required library successfully installed.")
}
