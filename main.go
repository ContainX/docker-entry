package main

import (
	"os"

	"github.com/ContainX/docker-entry/command"
)

func main() {

	if len(os.Args) < 2 {
		return
	}

	command.ExecuteCommand(os.Args[1:])
}
