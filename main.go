package main

import (
	"os"

	"github.com/yusriltakeuchi/gobook/core"
)

func main() {

	args := os.Args[1:]
	core.Command(args)
}
