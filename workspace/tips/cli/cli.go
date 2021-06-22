package cli

import (
	"bufio"
	"fmt"
	"github/gophers/tips/controller"
	"io"
	"strings"
)

func isValidInput(userInput string) bool {
	if len(userInput) <= 3 && userInput != "" {
		return false
	}
	return true
}

func readInput(reader io.Reader) string {
	inputReader := bufio.NewReader(reader)
	// ReadString will block until the delimiter is entered
	input, _ := inputReader.ReadString('\n')
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}

//returning Title
func GetTopic(reader io.Reader, writer io.Writer) (string, error) {
	user_input := readInput(reader)
	if isValidInput(user_input) {
		return user_input, nil
	}
	return nil, err
}

func Run(reader io.Reader, writer io.Writer) {
	fmt.Printf(" %q \n", "Enter Topic: ")
	//Get topic from User
	topic, _ := GetTopic(reader, writer)
	//Print Tip for the Topic
	controller.GetTipForTopic(topic, writer)
}
