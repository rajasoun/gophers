// +build !integration

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

type getTopic_invalid struct{ mock_title string }
type getTopic_valid struct{ title string }

func (getTopic_invalid *getTopic_invalid) readInput(i io.Reader) string {
	getTopic_invalid.mock_title = "git"
	return getTopic_invalid.mock_title
}

func (getTopic_valid *getTopic_valid) readInput(i io.Reader) string {
	getTopic_valid.title = "git commit"
	return getTopic_valid.title
}
func TestGetTopic(t *testing.T) {
	t.Run("Check GetTopic with reader interface", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		var key getTopic_valid
		got, _ := getTopic(&key, &output_buffer)
		want := "git commit"
		assert.Equal(t, got, want)
	})

	t.Run("Check GetTopic invalid topic", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		var key getTopic_invalid
		_, err := getTopic(&key, &output_buffer)
		assert.Error(t, err)
	})
}
