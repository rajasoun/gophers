package controller

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"

	"github/gophers/tips/model"
)

// pass userinput to model and also write tip in console
func GetTipForTopic(topic string, writer io.Writer) {
	tip := model.GetTip(topic)
	logrus.WithField("tip", tip).Debug("successfully display tip of the valid topic ")
	fmt.Fprintf(writer, "  \n %q \n\n", tip)
}
