package model

import (
	"testing"
)

func TestGetTip(t *testing.T) {
	assertEquals := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Get Tip for valid Topic", func(t *testing.T) {
		got := GetTip("git status")
		want := "git status -s"
		assertEquals(t, got, want)
	})
	t.Run("Get Tip for invalid Topic", func(t *testing.T) {
		got := GetTip("dummy")
		want := "Tips Not Available for Topic"
		assertEquals(t, got, want)
	})

}
