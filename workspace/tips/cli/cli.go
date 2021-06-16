package cli

//function(Anonymous func can accept inputs and return outputs) type which returning string
type userInput func() string

//returning Title
func GetTopic(userInput userInput) string {
	title := userInput()
	return title
}
