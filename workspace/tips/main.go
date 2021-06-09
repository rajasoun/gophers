package main

import (
	"fmt"
	"os"

	"github/gophers/tips/controller"
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
