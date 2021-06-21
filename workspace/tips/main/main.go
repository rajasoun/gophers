package main

import (
	"fmt"
	"github/gophers/tips/controller"
	"os"
)

func init() {
	fmt.Println("-->>Git Commands<<--")
}

func main() {
	//code done with Scanner interface implement function (without io.Reader interface)
	//scan := controller.ScannerImpl{}
	//controller.GetTipForTopic(os.Stdout, scan)

	// Code done with io.Reader interface
	controller.GetTipForTopicc(os.Stdout, os.Stdin)
	fmt.Println(" \n ")
	controller.GetMoreCommands(os.Stdout, controller.Input)
}

//to do error and pointer with Tdd coverage
