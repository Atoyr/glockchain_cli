package main

import (
	"os"

	"github.com/atoyr/glockchain_cli/cli"
)

func main() {
	c := cli.NewCLI()
	c.App.Run(os.Args)
}
