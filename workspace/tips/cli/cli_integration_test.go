package cli

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}
func TestReadInput(t *testing.T) {
	t.Run("Integration Testing : Reading data from console(userInput) ", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.WriteString("push")
		got := Stdin.readInput(&buffer)
		want := "push"
		assert.Equal(t, got, want)
	})
}

type endToEnd struct{ title string }
type endToEndWithValidation struct{ title string }

func (endToEnd *endToEnd) readInput(i io.Reader) string {
	endToEnd.title = "push"
	return endToEnd.title
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
