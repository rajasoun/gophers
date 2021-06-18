package controller

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
	title string
}

func (m mock) scanTitleFromConsole() string {
	m.title = "delete"
	return m.title
}

func TestGetTipForTopic(t *testing.T) {
	sn := ScannerImpl{}
	t.Run("delete", func(t *testing.T) {
		m := &mock{}
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, m)
		got := buffer.String()
		want := "Tip for delete is Delete remote branch"
		assert.Contains(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, sn)
		got := buffer.String()
		want := "Default tip: \n \"Saving current state of tracked"
		assert.Contains(t, got, want)
	})
}
