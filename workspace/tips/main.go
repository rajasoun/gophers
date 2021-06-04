package main

import (
	"fmt"
	"os"

	"github/gophers/tips/controller"
)

func init() {
	fmt.Println("-->>Git Commands<<--")
}
func main() {
	scan := controller.ScannerImpl{}
	controller.GetTipForTopic(os.Stdout, scan)
}
