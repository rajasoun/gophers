package controller

import (
	"fmt"
	"io"

	"github/gophers/tips/model"

	"github.com/sirupsen/logrus"
)

// pass userinput to model and also write tip in console
func GetTipForTopic(topic string, writer io.Writer) {
	tip := model.GetTip(topic)
	if logrus.GetLevel() == logrus.DebugLevel {
		logrus.WithField("tip", tip).Debug("successfully display tip of the valid topic ")
	} else {
		fmt.Fprintf(writer, "  \n %q \n\n", tip)
	}
}
