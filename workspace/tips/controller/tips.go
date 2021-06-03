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

func GetTipForTopic(writer io.Writer) {
	topic := cli.GetTopic(scanTitleFromConsole)
	tip := model.GetTip(topic)
	fmt.Fprintf(writer, "Tip for %q is %q \n", topic, tip)
}

func scanTitleFromConsole() string {
	topics := model.LoadTipsFromJson()
	for index, _ := range topics {
		fmt.Println(topics[index].Title)
	}
	fmt.Print("Enter any title from above to get a tip: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	return input
}
