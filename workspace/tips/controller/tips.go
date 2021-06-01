package controller

import (
	"fmt"
	"io"

	"github.com/gophers/tips/cli"
	"github.com/gophers/tips/model"
)

func GetTipForTopic(writer io.Writer) {
	var topic = cli.GetTopic()
	var tip = model.GetTip(topic)
	fmt.Fprintf(writer, "Tip for %q is %q \n", topic, tip)
}
