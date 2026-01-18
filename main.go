package main

import (
	"flag"

	"climb/src/cli"
)

func main() {
	dryRun := flag.Bool("dry-run", false, "When enabled no files are modified, created or deleted")

	flag.Parse()

	var args = flag.Args()

	cli.Cmd(args, *dryRun)
}
