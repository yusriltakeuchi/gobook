package main

import (
	"os"

	"gobook/core"
)

func main() {

	args := os.Args[1:]
	core.Command(args)
}
