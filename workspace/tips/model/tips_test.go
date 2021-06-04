package model

import (
	"reflect"
	"testing"
)

func TestGetTip(t *testing.T) {
	//tips := Tips{}
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
	t.Run("Get Tip for valid Topic", func(t *testing.T) {
		got := GetTip("git delete remote branch")
		want := "git push origin --delete <remote_branchname>"
		assertEquals(t, got, want)
	})

}
func TestLoadTipsFromJson(t *testing.T) {
	_, got := LoadTipsFromJson()
	expected, _ := MockReadJsonFile()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %q want %q", got, expected)
	}
}

func TestReadJsonFileNegative(t *testing.T) {
	_, err := MockReadJsonFile()
	want := "file issue"

	if reflect.DeepEqual(err, want) {
		t.Errorf("err %q want %q", err, want)
	}
}

func MockReadJsonFile() (string, error) {
	myJson :=
		`[{"title":"Initialize git repo","tip":"git init"},{"title":"git clone","tip":"git clone <repo-dir>"},{"title":"git config","tip":"git config --global user.email<email.id>"},{"title":"git status","tip":"git status -s"},{"title":"add code to github","tip":"git add ."},{"title":"git commit","tip":"git commit -m <commit message>"},{"title":"git push remote branch","tip":"git push -u origin <branch name>"},{"title":"pull code from remote","tip":"git pull --rebase"},{"title":"git checkout","tip":"git checkout <name of repo branch>"},{"title":"git merge","tip":" git merge <query>"},{"title":"git reset","tip":"git reset --hard"},{"title": "git help","tip": "git help -g"},{"title": "git delete remote branch","tip": "git push origin --delete <remote_branchname>"},{"title": "Saving current state of tracked files without commiting","tip": "git stash"},{"title": "Stash changes before rebasing","tip": "git rebase --autostash"},{"title": "Show both staged and unstaged changes","tip": "git diff HEAD"}]`

	return string(myJson), nil
}
