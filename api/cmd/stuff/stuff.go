package main

import (
	"os"

	"github.com/cdriehuys/stuff/api/internal/cli"
	"github.com/cdriehuys/stuff/api/migrations"
)

func main() {
	cmd := cli.NewRootCmd(os.Stderr, migrations.Files)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
