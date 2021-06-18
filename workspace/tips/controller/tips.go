package controller

import (
	"bufio"
	"fmt"
	"github/gophers/tips/cli"
	"github/gophers/tips/model"
	"io"
	"os"
	"strings"
)

type scanner interface {
	scanTitleFromConsole() string
}

type ScannerImpl struct {
	message string
}

const help = "git-tip --all"

//returning Tips in console according to user-input
func GetTipForTopic(writer io.Writer, scan scanner) {
	topic := cli.GetTopic(scan.scanTitleFromConsole)
	switch topic {
	case help:
		data, _ := model.LoadTipsFromJson()
		for index := range data {
			title := data[index].Title
			tip := data[index].Tip
			fmt.Fprintf(writer, " %q \n %q \n\n", title, tip)
		}
	case "":
		topic := "Saving current state of tracked files without commiting"
		tip := "git stash"
		fmt.Fprintf(writer, "Default tip: \n %q \n %q \n", topic, tip)
	default:
		//to do retrun actual title
		tip := model.GetTip(topic)
		fmt.Fprintf(writer, "Tip for %s is %s \n", topic, tip)
	}
}

//Implemention of interface methods with ScannerImpl struct type/class
func (scan ScannerImpl) scanTitleFromConsole() string {
	scanMessage := ScannerImpl{message: "->>> Enter key to get a tip or (git-tip --all)"}
	fmt.Println(scanMessage.message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}
