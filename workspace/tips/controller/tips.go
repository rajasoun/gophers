package controller

import (
	"fmt"
	"github/gophers/tips/model"
	"io"
)

const ALL = "git-tip --all"
const EMPTY = ""

//with Reader interface
func GetTipForTopic(topic string, writer io.Writer) {
	var tip = model.GetTip(topic)
	fmt.Fprintf(writer, " %q \n %q \n\n", topic, tip)
}
