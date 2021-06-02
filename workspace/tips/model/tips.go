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
	data := loadTipsFromJson()
	for index, _ := range data {
		//if strings.Compare(title, data[index].Title) == 0 {
		if title == data[index].Title {
			return data[index].Tip
		}
	}
	return "Tips Not Available for Topic"
}

func loadTipsFromJson() []Tips {
	data := `[{"title":"git status","tip":"git status -s"},{"title": "Everyday Git in twenty commands or so","tip": "git help everyday"},{"title": "Show helpful guides that come with Git","tip": "git help -g"},{"title": "Delete remote branch","tip": "git push origin --delete <remote_branchname>"},{"title": "Saving current state of tracked files without commiting","tip": "git stash"}]`
	var result []Tips
	json.Unmarshal([]byte(data), &result)
	return result
}
