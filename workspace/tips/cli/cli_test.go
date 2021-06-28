package cli

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//todo Table driven test
func init() {
	os.Setenv("GO_ENV", "test")
}
func TestInValidInput(t *testing.T) {
	t.Run("Checking title should be greater than len 3", func(t *testing.T) {
		got := isValidInput("jhg")
		want := false
		assert.Equal(t, got, want)
	})
	t.Run("checking for valid key", func(t *testing.T) {
		got := isValidInput("push")
		want := true
		assert.Equal(t, got, want)
	})
}
func TestReadInput(t *testing.T) {
	t.Run("Integration Testing : Reading data from console(userInput) ", func(t *testing.T) {
		//input_buffer := strings.NewReader("push")
		var buffer bytes.Buffer
		buffer.WriteString("push")
		//got := Stdin.readInput(input_buffer)
		got := Stdin.readInput(&buffer)
		want := "push"
		assert.Equal(t, got, want)
	})
}

type getTopic_invalid struct {
	mock_title string
}

func (getTopic_invalid *getTopic_invalid) readInput(i io.Reader) string {
	getTopic_invalid.mock_title = "git"
	return getTopic_invalid.mock_title
}

type getTopic_valid struct {
	title string
}

func (getTopic_valid *getTopic_valid) readInput(i io.Reader) string {
	getTopic_valid.title = "git commit"
	return getTopic_valid.title
}
func TestGetTopic(t *testing.T) {
	t.Run("Check GetTopic with reader interface", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		var key getTopic_valid
		got, _ := GetTopic(&key, &output_buffer)
		want := "git commit"
		assert.Equal(t, got, want)
	})

	t.Run("Check GetTopic invalid topic", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		var key getTopic_invalid
		_, err := GetTopic(&key, &output_buffer)
		assert.Error(t, err)
	})
}

type endToEnd struct {
	title string
}

func (endToEnd *endToEnd) readInput(i io.Reader) string {
	endToEnd.title = "push"
	return endToEnd.title
}

type endToEndWithValidation struct {
	title string
}

func (endToEndWithValidation *endToEndWithValidation) readInput(i io.Reader) string {
	endToEndWithValidation.title = "nor"
	return endToEndWithValidation.title
}
func TestRun(t *testing.T) {
	t.Run("End to End Test (e2e) for tips tool", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		input_buffer := &endToEnd{}
		Run(input_buffer, &output_buffer)
		got := output_buffer.String()
		want := "push"
		assert.Contains(t, got, want)
	})
	t.Run("End to End Test (e2e) for tips tool with validation", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		input_buffer := &endToEndWithValidation{}

		Run(input_buffer, &output_buffer)
		got := output_buffer.String()
		want := "\"key length should be greater than 3\""
		assert.Contains(t, got, want)
	})
}
