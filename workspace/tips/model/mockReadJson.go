package model

func MockTips() string {
	jsonMockData :=
		`[
		{
			"title":"Everyday Git in twenty commands or so",
			"tip":"git help everyday"
		},
		{
			"title":"Show helpful guides that come with Git",
			"tip":"git help -g"
		},
		{
			"title":"Saving current state of unstaged changes to tracked files",
			"tip":"git stash -k",
			"alternatives":[
			   "git stash --keep-index",
			   "git stash push --keep-index"
			]
		 },
	]`
	return string(jsonMockData)
}
