package main

import (
	"log"

	"github.com/debdutdeb/semis/cmd"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
