package model

import (
	"encoding/json"
)

type Tips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

//return tip
func GetTip(title string) string {
	// pass:1 - Make test pass
	// hardcoded
	//:ToDo: Reafactor To Load Tips from JSON
	data := LoadTipsFromJson()
	for index := range data {
		//if strings.Compare(title, data[index].Title) == 0 {
		if title == data[index].Title {
			return data[index].Tip
		}
	}
	return "Tips Not Available for Topic"
}

var data = `[{"title":"Initialize git repo","tip":"git init"},{"title":"git clone","tip":"git clone <repo-dir>"},{"title":"git config","tip":"git config --global user.email<email.id>"},{"title":"git status","tip":"git status -s"},{"title":"add code to github","tip":"git add ."},{"title":"git commit","tip":"git commit -m <commit message>"},{"title":"git push remote branch","tip":"git push -u origin <branch name>"},{"title":"pull code from remote","tip":"git pull --rebase"},{"title":"git checkout","tip":"git checkout <name of repo branch>"},{"title":"git merge","tip":" git merge <query>"},{"title":"git reset","tip":"git reset --hard"},{"title": "git help","tip": "git help -g"},{"title": "git delete remote branch","tip": "git push origin --delete <remote_branchname>"},{"title": "Saving current state of tracked files without commiting","tip": "git stash"},{"title": "Stash changes before rebasing","tip": "git rebase --autostash"},{"title": "Show both staged and unstaged changes","tip": "git diff HEAD"}]`

// func ReadJsonFile() ([]byte, error) {

// 	jsonData, err := ioutil.ReadFile("tipsData.json")
// 	if err != nil {
// 		return []byte{}, errors.New("file issue")
// 	}
// 	return jsonData, nil
// }

func LoadTipsFromJson() []Tips {

	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result
}
