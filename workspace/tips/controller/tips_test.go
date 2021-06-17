package controller

import (
	"bytes"
	"testing"
)

type mock struct {
	title string
}
type mockHelpKey struct {
	key string
}

func TestGetTipForTopic(t *testing.T) {
	sn := ScannerImpl{}
	assertEquals := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("delete", func(t *testing.T) {
		m := &mock{}
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, m)
		got := buffer.String()
		want := "Tip for delete is Delete remote branch : git push origin --delete <remote_branchname> \n"
		assertEquals(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, sn)
		got := buffer.String()
		want := "Default tip: \n \"Saving current state of tracked files without commiting\" \n \"git stash\" \n"
		assertEquals(t, got, want)
	})

	t.Run("git-tip --all", func(t *testing.T) {
		m := &mockHelpKey{}
		buffer := bytes.Buffer{}
		GetTipForTopic(&buffer, m)
		got := buffer.String()
		want := S
		assertEquals(t, got, want)

	})

}
func (m mock) scanTitleFromConsole() string {
	m.title = "delete"
	return m.title
}

func (m mockHelpKey) scanTitleFromConsole() string {
	m.key = "git-tip --all"
	return m.key
}
