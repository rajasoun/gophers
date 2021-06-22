package cli

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

//function(Anonymous func can accept inputs and return outputs type which returning string
type userInput func() string

//returning Title
func GetTopic(userInput userInput) string {
	title := userInput()
	return title
}

func GetTopicc(reader io.Reader) (string, error) {

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
