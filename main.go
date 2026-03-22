package main

import (
	"github.com/Compogo/sceleton/interface/cli"
)

func main() {
	cmd, err := cli.NewRootCommand()
	if err != nil {
		panic(err)
	}

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
