package main

import (
	"os"

	"github.com/LJTian/maker-flow/templates/apps/go-cli/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
