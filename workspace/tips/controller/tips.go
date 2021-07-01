package controller

import (
	"fmt"
	"io"

	"github/gophers/tips/model"
)

// pass userinput to model and also write tip in console
func GetTipForTopic(topic string, writer io.Writer) {
	tip := model.GetTip(topic)
	fmt.Fprintf(writer, "  \n %q \n\n", tip)
}
