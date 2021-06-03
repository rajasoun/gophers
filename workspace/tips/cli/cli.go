package cli

// :ToDo: Copy Paste Issue
type InputGetter func() string

func GetTopic(inputGetter InputGetter) string {
	// pass:1 - Make test pass
	// hardcoded
	//:ToDo: Get User Input - Mock GetInput in Test

	title := inputGetter() //title
	return title

}
