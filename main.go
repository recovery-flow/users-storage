package main

import (
	"os"

	"github.com/cifra-city/users-storage/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
