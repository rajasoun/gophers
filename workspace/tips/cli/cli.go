package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github/gophers/tips/controller"
)

// scan inpur from user
func readInput(reader io.Reader) string {
	inputReader := bufio.NewReader(reader)
	// ReadString will block until the delimiter is entered
	input, _ := inputReader.ReadString('\n')
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}

//check input validation
func isValidInput(userInput string) bool {
	if len(userInput) <= 3 && userInput != "" {
		return false
	}
	return true
}

//returning Title
func getTopic(reader io.Reader) (string, error) {
	var validError = errors.New("key length should be greater than 3")
	user_input := readInput(reader)
	if isValidInput(user_input) {
		return user_input, nil
	}
	return "", validError
}

//read user input and pass to controller
func Run(reader io.Reader, writer io.Writer) {
	//fmt.Printf(" %q \n", "Enter Topic: ")
	topic, err := getTopic(reader)
	if err != nil {
		fmt.Fprintf(writer, " %q", err.Error())
	} else {
		controller.GetTipForTopic(topic, writer)
	}
}
