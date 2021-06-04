package cli

//function(Anonymous func can accept inputs and return outputs) type which returning string

// :ToDo: Copy Paste Issue
type userInput func() string

//returning Title
func GetTopic(userInput userInput) string {
	//:ToDo: Get User Input - Mock GetInput in Test
	title := userInput() //title
	return title

}
