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

type Scanner interface {
	scanTitleFromConsole() string
}

type ScannerImpl struct {
}

func GetTipForTopic(writer io.Writer, c Scanner) {
	topic := cli.GetTopic(c.scanTitleFromConsole)
	tip := model.GetTip(topic)
	fmt.Fprintf(writer, "Tip for %q is %q \n", topic, tip)
}

func (c ScannerImpl) scanTitleFromConsole() string {

	allTopics := model.LoadTipsFromJson()

	for index, _ := range allTopics {
		fmt.Println(allTopics[index].Title)
	}

	fmt.Println("Enter a title from above to get a tip:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	input = strings.TrimSuffix(input, "\n")

	return input
}
