package main

import (
	"flag"

	"climb/src/cli"
	"climb/src/utils"
)

func main() {
	flag.Parse()

	var args = flag.Args()

	utils.ValidateArgs(args)
	cli.Cmd(args)
}
