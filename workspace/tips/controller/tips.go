package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github/gophers/tips/cli"
	"github/gophers/tips/model"
)

type scanner interface {
	scanTitleFromConsole() string
}

type ScannerImpl struct {
	message string
}

//returning Tips in console according to user-input
func GetTipForTopic(writer io.Writer, scan scanner) {
	topic := cli.GetTopic(scan.scanTitleFromConsole)
	tip := model.GetTip(topic)
	fmt.Fprintf(writer, "Tip for %q is %q \n", topic, tip)
}

//Implemention of interface methods with ScannerImpl struct type/class
func (scan ScannerImpl) scanTitleFromConsole() string {
	scanMessage := ScannerImpl{message: "->>> Enter Any Title  To Get a Tip:"}
	//allTopics, _ := model.LoadTipsFromJson()

	// for index, _ := range allTopics {
	// 	//fmt.Println(allTopics[index].Title)
	// }

	fmt.Println(scanMessage.message)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal("An error occured while reading input.Please try again")
	// }
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
