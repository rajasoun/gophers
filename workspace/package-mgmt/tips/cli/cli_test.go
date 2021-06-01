package cli

import (
	"testing"
)

func TestGetTopicFromConsole(t *testing.T) {
	assertEquals := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Get Topic String From Console", func(t *testing.T) {
		got := GetTopic()
		want := "git status"
		assertEquals(t, got, want)
	})
}
