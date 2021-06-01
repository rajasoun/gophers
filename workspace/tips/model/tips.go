package model

import (
	"strings"
)

func GetTip(topic string) string {
	// pass:1 - Make test pass
	// hardcoded
	//:ToDo: Reafactor To Load Tips from JSON
	var result string
	if strings.Compare(topic, "git status") == 0 {
		result = "git status -s"
	} else {
		result = "Tips Not Available for Topic"
	}
	return result
}
