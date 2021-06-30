// +build integration

package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func TestRun(t *testing.T) {
	t.Run("End to End Test (e2e) for tips tool", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		input_buffer := strings.NewReader("push")
		Run(input_buffer, &output_buffer)
		got := output_buffer.String()
		want := "push"
		assert.Contains(t, got, want)
	})
	t.Run("End to End Test (e2e) for tips tool with validation", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		input_buffer := strings.NewReader("nor")
		Run(input_buffer, &output_buffer)
		got := output_buffer.String()
		want := "\"key length should be greater than 3\""
		assert.Contains(t, got, want)
	})
}
