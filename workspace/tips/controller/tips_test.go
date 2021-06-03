package controller

import (
	"bytes"
	"testing"
)

func TestGetTipForTopic(t *testing.T) {
	sn := ScannerImpl{}
	assertEquals := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("git delete", func(t *testing.T) {
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, sn)

		got := buffer.String()
		want := "Tip for \"\" is \"Tips Not Available for Topic\" \n"
		assertEquals(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, sn)

		got := buffer.String()
		want := "Tip for \"\" is \"Tips Not Available for Topic\" \n"
		assertEquals(t, got, want)
	})

	// t.Run("git command", func(t *testing.T) {
	// 	buffer := bytes.Buffer{}
	// 	GetTipForTopic(&buffer)
	// 	got := buffer.String()
	// 	want := "Tip for \"git status\" is \"git status -s\" \n"
	//assertEquals(t, got, want)
	// })

}
