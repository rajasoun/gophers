package main

import (
	"github/gophers/tips/cli"
	"os"
)

//read user input and write the tip in console
func main() {
	cli.Run(os.Stdin, os.Stdout)
}
