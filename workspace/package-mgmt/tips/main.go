package main

import (
	"fmt"

	"github.com/gophers/tips/cli"
	"github.com/gophers/tips/model"
)

func GetTipForTopic() {
	var topic = cli.GetTopic()
	var tip = model.GetTip(topic)
	fmt.Printf("Tip for %q is %q \n", topic, tip)
}

func main() {
	GetTipForTopic()
}
