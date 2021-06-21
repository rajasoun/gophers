package controller

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
	title string
}
type helpCommand struct {
	command string
}

func (m mock) ScanTitleFromConsole() string {
	m.title = "delete"
	return m.title
}
func (hc helpCommand) ScanTitleFromConsole() string {
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
		want := "Extract file from another branch."
		assert.Contains(t, got, want)
	})
}

func TestGetMoreCommands(t *testing.T) {
	t.Run("Check more other commands According to input", func(t *testing.T) {
		buffer := bytes.Buffer{}
		GetMoreCommands(&buffer, "stash")
		got := buffer.String()
		want := "Show list of all saved stashes : git stash list"
		assert.Contains(t, got, want)
	})
}

func TestGetTipForTopiccWithReader(t *testing.T) {
	t.Run("Delete", func(t *testing.T) {
		buffer := bytes.Buffer{}
		reader := strings.NewReader("delete")
		GetTipForTopicc(&buffer, reader)
		got := buffer.String()
		want := "Tip for delete is Delete remote branch"
		assert.Contains(t, got, want)
	})
	t.Run("Empty string", func(t *testing.T) {
		buffer := bytes.Buffer{}
		input := strings.NewReader("")
		GetTipForTopicc(&buffer, input)
		got := buffer.String()
		want := "Default tip: \n \"Saving current state of tracked"
		assert.Contains(t, got, want)
	})
	t.Run("git-tip --all", func(t *testing.T) {
		input := strings.NewReader("git-tip --all")
		buffer := bytes.Buffer{}
		GetTipForTopicc(&buffer, input)
		got := buffer.String()
		want := "Extract file from another branch."
		assert.Contains(t, got, want)
	})
	t.Run("Check title length should be greter than 3 ", func(t *testing.T) {
		input := strings.NewReader("pok")
		buffer := bytes.Buffer{}
		GetTipForTopicc(&buffer, input)
		got := buffer.String()
		want := "\"word should be greater than 3\""
		assert.Equal(t, got, want)
	})

}
