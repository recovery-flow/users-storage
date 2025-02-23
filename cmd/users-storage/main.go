package main

import (
	"os"

	"github.com/recovery-flow/users-storage/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
