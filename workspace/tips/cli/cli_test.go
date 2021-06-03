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
	t.Run("Get Status Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInput)
		want := "git status"
		assertEquals(t, got, want)
	})
	t.Run("Get Empty String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputEmpty)
		want := ""
		assertEquals(t, got, want)
	})
	t.Run("Get Delete Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputDelete)
		want := "git delete remote branch"
		assertEquals(t, got, want)
	})
	t.Run("Get Help Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputHelp)
		want := "git help -g"
		assertEquals(t, got, want)
	})
}

func mockTestUserInput() string {
	return "git status"
}
func mockTestUserInputEmpty() string {
	return ""
}
func mockTestUserInputDelete() string {
	return "git delete remote branch"
}

func mockTestUserInputHelp() string {
	return "git help -g"
}
