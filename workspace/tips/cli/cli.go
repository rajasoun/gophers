package cli

import (
	"bufio"
	"errors"
	"fmt"
	"github/gophers/tips/controller"
	"io"
	"strings"
)

//returning Title
func GetTopic(reader io.Reader, writer io.Writer) (string, error) {
	fmt.Fprintf(writer, " %q \n", "Enter Topic: ")
	var error = errors.New("word should be greater than 3")
	inputReader := bufio.NewReader(reader)
	// ReadString will block until the delimiter is entered
	input, _ := inputReader.ReadString('\n')
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	if len(input) <= 3 && input != "" {
		return input, error
	}
	return input, nil
}

func Run(reader io.Reader, writer io.Writer) {
	//Get topic from User
	topic, _ := GetTopic(reader, writer)
	//Print Tip for the Topic
	controller.GetTipForTopic(topic, writer)
}
