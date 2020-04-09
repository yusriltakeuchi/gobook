package core

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	// Open file for reading.
	text, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(text)
}

func writeFile(path string, code string) bool {
	//if not found then create file
	var _, err = os.Stat(path)

	if err != nil {
		var file, err = os.Create(path)
		if err != nil {
			return false
		}
		defer file.Close()
	}

	var file, erro = os.OpenFile(path, os.O_RDWR, 0644)
	if erro != nil {
		fmt.Println(erro.Error())
		return false
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(code)

	// Save file changes.
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
