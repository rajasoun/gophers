package main

import (
	"os"

	"github/gophers/tips/controller"
)

func main() {

	controller.GetTipForTopic(os.Stdout)
}
