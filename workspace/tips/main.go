package main

import (
	"os"

	"github/gophers/tips/controller"
)

func main() {
	sn := controller.ScannerImpl{}
	controller.GetTipForTopic(os.Stdout, sn)
}
