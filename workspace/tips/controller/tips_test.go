package controller

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
	title string
}
type helpCommand struct {
	command string
}

func (m mock) scanTitleFromConsole() string {
	m.title = "delete"
	return m.title
}
func (hc helpCommand) scanTitleFromConsole() string {
	hc.command = "git-tip --all"
	return hc.command
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
	t.Run("git-tip --all", func(t *testing.T) {
		hc := &helpCommand{}
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, hc)
		got := buffer.String()
		//want := 166
		want := "Everyday Git in twenty commands or so"
		assert.Contains(t, got, want)
	})
}
