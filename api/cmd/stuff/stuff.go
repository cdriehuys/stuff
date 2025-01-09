package main

import (
	"os"

	"github.com/cdriehuys/stuff/api/internal/cli"
)

func main() {
	cmd := cli.NewRootCmd(os.Stderr)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
