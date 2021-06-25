package cli

import (
	"bufio"
	"errors"
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
	var validError = errors.New("key length should be greater than 3")
	user_input := readInput(reader)
	if isValidInput(user_input) {
		return user_input, nil
	}
	return "", validError
}

func Run(reader io.Reader, writer io.Writer) {
	fmt.Printf(" %q \n", "Enter Topic: ")
	//Get topic from User
	topic, err := GetTopic(reader, writer)
	//print error
	if err != nil {
		fmt.Fprintf(writer, "%q", err.Error())
	} else {
		//Print Tip for the Topic
		controller.GetTipForTopic(topic, writer)
	}
}
