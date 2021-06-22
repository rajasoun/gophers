package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopic(t *testing.T) {
	t.Run("Check GetTopic with reader interface", func(t *testing.T) {
		key := strings.NewReader("git commit")
		output_buffer := bytes.Buffer{}
		got, _ := GetTopic(key, &output_buffer)
		want := "git commit"
		assert.Equal(t, got, want)
	})
	t.Run("Check GetTopic invalid topic", func(t *testing.T) {
		key := strings.NewReader("git")
		output_buffer := bytes.Buffer{}
		_, err := GetTopic(key, &output_buffer)
		assert.Error(t, err)
	})
}

func TestRun(t *testing.T) {
	t.Run("End to End Test (e2e) for tips tool", func(t *testing.T) {
		input_buffer := strings.NewReader("push")
		output_buffer := bytes.Buffer{}
		Run(input_buffer, &output_buffer)
		got := output_buffer.String()
		want := "push"
		assert.Contains(t, got, want)
	})
}
