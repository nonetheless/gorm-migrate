package main

import (
	"github.com/nonetheless/gorm-migrate/cmd"
	"os"
)

func main() {
	cmd := cmd.NewRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}