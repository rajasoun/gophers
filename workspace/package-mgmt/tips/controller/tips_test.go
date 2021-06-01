package controller

import (
	"bytes"
	"testing"
)

func TestGetTipForTopic(t *testing.T) {
	buffer := bytes.Buffer{}
	GetTipForTopic(&buffer)

	got := buffer.String()
	want := "Tip for \"git status\" is \"git status -s\" \n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
