package main

import (
	"fmt"
	"github/gophers/tips/controller"
	"os"
)

func init() {
	fmt.Println("-->>Git Commands<<--")
	fmt.Println("<keyword> Gives the git tips consisting of the keyword")
}

func main() {
	scan := controller.ScannerImpl{}
	controller.GetTipForTopic(os.Stdout, scan)
	fmt.Println(" \n ")
	controller.MoreCommnads(os.Stdout, controller.Input)
}

//to do error and pointer with Tdd coverage
// to do reader interfaces
