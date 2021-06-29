package main

import (
	"github/gophers/tips/cli"
	"os"
)

func main() {
	cli.Run(cli.Stdin, os.Stdout)

}
