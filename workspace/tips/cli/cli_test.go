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
		want := "git help"
		assertEquals(t, got, want)
	})
	t.Run("Get Init Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputInit)
		want := "Initialize git repo"
		assertEquals(t, got, want)
	})
	t.Run("Get Clone Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputClone)
		want := "git clone"
		assertEquals(t, got, want)
	})
	t.Run("Get Config Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputConfig)
		want := "git config"
		assertEquals(t, got, want)
	})
	t.Run("Get Add Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputAdd)
		want := "add code to github"
		assertEquals(t, got, want)
	})
	t.Run("Get Commit Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputCommit)
		want := "git commit"
		assertEquals(t, got, want)
	})
	t.Run("Get Push Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputPush)
		want := "git push remote branch"
		assertEquals(t, got, want)
	})
	t.Run("Get Pull Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputPull)
		want := "pull code from remote"
		assertEquals(t, got, want)
	})
	t.Run("Get Checkout Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputCheckout)
		want := "git checkout"
		assertEquals(t, got, want)
	})
	t.Run("Get Merge Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputMerge)
		want := "git merge"
		assertEquals(t, got, want)
	})
	t.Run("Get Reset Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputReset)
		want := "git reset --hard"
		assertEquals(t, got, want)
	})
	t.Run("Get Stash Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputStash)
		want := "Saving current state of tracked files without commiting"
		assertEquals(t, got, want)
	})
	t.Run("Get Rebase Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputRebase)
		want := "Stash changes before rebasing"
		assertEquals(t, got, want)
	})
	t.Run("Get Diff Topic String From Console", func(t *testing.T) {

		got := GetTopic(mockTestUserInputDiff)
		want := "Show both staged and unstaged changes"
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

func mockTestUserInputInit() string {
	return "Initialize git repo"
}
func mockTestUserInputClone() string {
	return "git clone"
}
func mockTestUserInputConfig() string {
	return "git config"
}
func mockTestUserInputAdd() string {
	return "add code to github"
}
func mockTestUserInputCommit() string {
	return "git commit"
}
func mockTestUserInputPush() string {
	return "git push remote branch"
}
func mockTestUserInputPull() string {
	return "pull code from remote"
}
func mockTestUserInputCheckout() string {
	return "git checkout"
}
func mockTestUserInputMerge() string {
	return "git merge"
}
func mockTestUserInputHelp() string {
	return "git help"
}
func mockTestUserInputReset() string {
	return "git reset --hard"
}
func mockTestUserInputStash() string {
	return "Saving current state of tracked files without commiting"
}
func mockTestUserInputRebase() string {
	return "Stash changes before rebasing"
}

func mockTestUserInputDiff() string {
	return "Show both staged and unstaged changes"
}
