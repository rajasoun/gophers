package cli

import (
	"testing"
)

func TestGetTopicFromConsole(t *testing.T) {
	assertEquals := func(t testing.TB, mockTestUserInput func() string, want string) {
		t.Helper()
		got := GetTopic(mockTestUserInput)
		expected := want
		if got != expected {
			t.Errorf("got %q want %q", got, want)
		}
	}
	mockTests := []struct {
		name string
		mock func() string
		want string
	}{
		{name: "Get Status Topic String From Console", mock: func() string { return "git status" }, want: "git status"},
		{name: "Get Empty String From Console", mock: func() string { return "" }, want: ""},
		{name: "Get Help Topic String From Console", mock: func() string { return "git help" }, want: "git help"},
		{name: "Get Delete Topic String From Console", mock: func() string { return "git delete remote branch" }, want: "git delete remote branch"},
		{name: "Get Init Topic String From Console", mock: func() string { return "Initialize git repo" }, want: "Initialize git repo"},
		{name: "Get Clone Topic String From Console", mock: func() string { return "git clone" }, want: "git clone"},
		{name: "Get Add Topic String From Console", mock: func() string { return "add code to github" }, want: "add code to github"},
		{name: "Get Commit Topic String From Console", mock: func() string { return "git commit" }, want: "git commit"},
		{name: "Get Push Topic String From Console", mock: func() string { return "git push remote branch" }, want: "git push remote branch"},
		{name: "Get Pull Topic String From Console", mock: func() string { return "pull code from remote" }, want: "pull code from remote"},
		{name: "Get Checkout Topic String From Console", mock: func() string { return "git checkout" }, want: "git checkout"},
		{name: "Get Merge Topic String From Console", mock: func() string { return "git merge" }, want: "git merge"},
		{name: "Get Reset Topic String From Console", mock: func() string { return "git reset --hard" }, want: "git reset --hard"},
		{name: "Get Stash Topic String From Console", mock: func() string { return "Saving current state of tracked files without commiting" }, want: "Saving current state of tracked files without commiting"},
		{name: "Get Rebase Topic String From Console", mock: func() string { return "Stash changes before rebasing" }, want: "Stash changes before rebasing"},
		{name: "Get Diff Topic String From Console", mock: func() string { return "Show both staged and unstaged changes" }, want: "Show both staged and unstaged changes"},
	}
	for _, tt := range mockTests {
		t.Run(tt.name, func(t *testing.T) {
			assertEquals(t, tt.mock, tt.want)
		})
	}

}
