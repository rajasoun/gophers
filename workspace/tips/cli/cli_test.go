// +build !integration

package cli

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGetTopic(t *testing.T) {
	t.Run("Check GetTopic with reader interface", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		var input_buffer bytes.Buffer
		input_buffer.WriteString("git commit")
		got, _ := getTopic(&input_buffer, &output_buffer)
		want := "git commit"
		assert.Equal(t, got, want)
	})

	t.Run("Check GetTopic invalid topic", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		var input_buffer bytes.Buffer
		input_buffer.WriteString("git")
		_, err := getTopic(&input_buffer, &output_buffer)
		assert.Error(t, err)
	})
}

func TestReadInput(t *testing.T) {
	t.Run("Reading data from console(userInput) ", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.WriteString("push")
		got := readInput(&buffer)
		want := "push"
		assert.Equal(t, got, want)
	})
}
