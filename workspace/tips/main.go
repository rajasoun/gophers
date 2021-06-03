package main

import (
	"os"

	"github.com/gophers/tips/controller"
)

func main() {

	controller.GetTipForTopic(os.Stdout)
}
