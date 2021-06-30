package controller

import (
	"fmt"
	"io"

	"github/gophers/tips/model"
)

func GetTipForTopic(topic string, writer io.Writer) {
	var model_Impl = model.File_reader_Impl{}
	tip := model.GetTip(topic, model_Impl)
	fmt.Fprintf(writer, "  \n %q \n\n", tip)
}
