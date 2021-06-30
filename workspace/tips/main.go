package main

import (
	"github/gophers/tips/cli"
	"os"
)

func main() {
	cli.Run(os.Stdin, os.Stdout)
}
