package cli

import (
	"bufio"
	"errors"
	"fmt"
	"github/gophers/tips/controller"
	"io"
	"os"
	"strings"
)

type reader interface{ readInput(io.Reader) string }

type ReaderInput struct {
	input string
}

var Stdin = ReaderInput{}

func (reader_input ReaderInput) readInput(reader io.Reader) string {
	inputReader := bufio.NewReader(reader)
	// ReadString will block until the delimiter is entered
	reader_input.input, _ = inputReader.ReadString('\n')
	// remove the delimeter from the string
	reader_input.input = strings.TrimSuffix(reader_input.input, "\n")
	return reader_input.input
}
func isValidInput(userInput string) bool {
	if len(userInput) <= 3 && userInput != "" {
		return false
	}
	return true
}

//returning Title
func getTopic(reader reader, writer io.Writer) (string, error) {
	var validError = errors.New("key length should be greater than 3")
	user_input := reader.readInput(os.Stdin)
	if isValidInput(user_input) {
		return user_input, nil
	}
	return "", validError
}

func Run(reader reader, writer io.Writer) {
	fmt.Printf(" %q \n", "Enter Topic: ")
	//Get topic from User
	topic, err := getTopic(reader, writer)
	//print error
	if err != nil {
		fmt.Fprintf(writer, " %q", err.Error())
	} else {
		//Print Tip for the Topic
		var controller_impl controller.Controller_Impl
		controller.GetTipForTopic(topic, writer, controller_impl)
	}
}
