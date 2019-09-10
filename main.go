package main

import (
	"os"

	"github.com/atoyr/glockchain_cli/cli"
)

func main() {
	cli := cli.NewCli()
	cli.App.Run(os.Args)
}
