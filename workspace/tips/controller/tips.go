package controller

import (
	"fmt"
	"io"

	"github/gophers/tips/model"
)

// pass userinput to model and also write tip in console
func GetTipForTopic(topic string, writer io.Writer) {
	tip := model.GetTip(topic, model.Reader{})
	fmt.Fprintf(writer, "  \n %q \n\n", tip)
}
