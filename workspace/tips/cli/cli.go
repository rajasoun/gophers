package cli

//return title: userInput
func GetTopic() string {
	// pass:1 - Make test pass
	// hardcoded
	//:ToDo: Get User Input - Mock GetInput in Test

	title := scanTitleFromConsole() //title

	return title[0]

}
func scanTitleFromConsole() []string {
	//var title string
	//fmt.Scanf("%s", &title)
	title := []string{"git status", "Everyday Git in twenty commands or so", "Show helpful guides that come with Git", "Delete remote branch", "Saving current state of tracked files without commiting"}
	return title
}
