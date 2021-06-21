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

const help = "git-tip --all"

var data, _ = model.LoadTipsFromJson()

var Input string

//with Reader interface
func GetTipForTopicc(writer io.Writer, scan io.Reader) {
	fmt.Println("->>> Enter key to get a tip or (git-tip --all)")
	var Title, err = cli.GetTopicc(scan)
	if err != nil {
		fmt.Fprintf(writer, "%q", err)
	} else {
		Input = Title
		switch Title {
		case help:
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
			tip := model.GetTip(Title)
			fmt.Fprintf(writer, "Tip for %s is %s \n", Title, tip)
		}
	}

}

type scanner interface {
	ScanTitleFromConsole() string
}

type ScannerImpl struct {
	message string
}

//:ToDo: To Underdtand teh Logic
//returning Tips in console according to user-input
func GetTipForTopic(writer io.Writer, scan scanner) {
	topic := cli.GetTopic(scan.ScanTitleFromConsole)
	switch topic {
	case help:
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
func (scan ScannerImpl) ScanTitleFromConsole() string {
	scanMessage := ScannerImpl{message: "->>> Enter key to get a tip or (git-tip --all)"}
	fmt.Println(scanMessage.message)
	reader := bufio.NewReader(os.Stdin)
	Input, _ = reader.ReadString('\n')
	Input = strings.TrimSuffix(Input, "\n")
	return Input
}

func GetMoreCommands(writer io.Writer, input string) {
	if input != "" && input != "dummy" {
		fmt.Printf("More %q commands :\n", input)
		commands := model.GetAllCommands(data, input)
		for index := 1; index < len(commands); index++ {
			fmt.Fprintf(writer, " %q \n\n", commands[index])

		}
	}
}
