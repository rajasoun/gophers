package controller

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTipForTopiccWithReader(t *testing.T) {
	//:ToDo: Table Driven Tests
	t.Run("Delete", func(t *testing.T) {
		output_buffer := bytes.Buffer{}
		GetTipForTopic("delete", &output_buffer)
		got := output_buffer.String()
		want := "Tip for delete is Delete remote branch"
		assert.Contains(t, got, want)
	})
	t.Run("Empty string", func(t *testing.T) {
		buffer := bytes.Buffer{}
		GetTipForTopic("", &buffer)
		got := buffer.String()
		want := "Default tip: \n \"Saving current state of tracked"
		assert.Contains(t, got, want)
	})
}
