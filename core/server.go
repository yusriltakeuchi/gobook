package core

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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

func GenerateCode(cmd []string) {
	switch cmd[1] {
	case "controller":
		MakeCode(cmd[2], "controllers")
	case "model":
		MakeCode(cmd[2], "models")
	}
}

func MakeCode(name string, types string) {
	basePath, _ := os.Getwd()
	destPath := fmt.Sprintf("%s/app/%s/%s.go", basePath, types, name)

	_code := GetTemplate(types)
	templateCode := strings.ReplaceAll(_code, "{{.projectPath}}", strings.Split(basePath, "src/")[1])
	templateCode = strings.ReplaceAll(templateCode, "{{.name}}", name)

	status := writeFile(destPath, templateCode)
	if status == true {
		fmt.Println(fmt.Sprintf(" -> Successfuly create %s %s.", types, name))
	} else {
		fmt.Println(fmt.Sprintf(" -> Failed to create %s", types))
	}

	if types == "models" {
		fmt.Println(fmt.Sprintf(" -> Creating Migrations %s.", name))
		AddingMigrations(basePath, name)
	}
}

func AddingMigrations(basePath string, name string) {
	sourcePath := fmt.Sprintf("%s/app/database/migrations.go", basePath)

	_code := ReadFile(sourcePath)

	_modelInsert := fmt.Sprintf("},\n		{&models.%s{}, '%s'}}", name, name)
	_modelInsert = strings.ReplaceAll(_modelInsert, "'", "\"")
	templateCode := strings.ReplaceAll(_code, "}}", _modelInsert)

	status := writeFile(sourcePath, templateCode)
	if status == false {
		fmt.Println(fmt.Sprintf(" -> Failed to create migrations %s", name))
	}
}

func AddingLibrary(name string) {
	basePath, _ := os.Getwd()
	sourcePath := fmt.Sprintf("%s/config/packages.go", basePath)

	_code := ReadFile(sourcePath)

	_packagesInsert := fmt.Sprintf("\"}\n		\"%s\"}", name)
	templateCode := strings.ReplaceAll(_code, "\"}", _packagesInsert)

	status := writeFile(sourcePath, templateCode)
	if status == false {
		fmt.Println(fmt.Sprintf(" -> Failed to import library %s", name))
	}
}
