package controller

import (
	"fmt"
	"io"

	"github/gophers/tips/model"
)

func GetTipForTopic(topic string, writer io.Writer) {
	tip := model.GetTip(topic, model.Reader{})
	fmt.Fprintf(writer, "  \n %q \n\n", tip)
}
