package controller

import (
	"fmt"
	"github/gophers/tips/model"
	"io"
)

var Controller_impl = Controller_Impl{}

func GetTipForTopic(topic string, writer io.Writer, controller Controller) {
	tip := controller.getTip(topic)
	fmt.Fprintf(writer, "  \n %q \n\n", tip)
}

func (con_impl Controller_Impl) getTip(topic string) string {
	tip := model.GetTip(topic, model.Model_Impl)
	return tip
}

type Controller interface{ getTip(string) string }
type Controller_Impl struct{}
