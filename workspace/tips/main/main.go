package main

import (
	"fmt"
	"github/gophers/tips/controller"
	"os"
)

func init() {
	fmt.Println("-->>Git Commands<<--")
	//fmt.Println("options")
	//fmt.Println("git-tip --all")
	//fmt.Println("<keyword> Gives the git tips consisting of the keyword")
}

func main() {
	scan := controller.ScannerImpl{}
	controller.GetTipForTopic(os.Stdout, scan)

}

//todo add 1 json file with absolute data
//to do error and pointer with Tdd coverage
// to do reader interfaces
//
