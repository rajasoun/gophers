package cli

type InputGetter func() string

func GetTopic(inputGetter InputGetter) string {
	// pass:1 - Make test pass
	// hardcoded
	//:ToDo: Get User Input - Mock GetInput in Test

	title := inputGetter() //title
	//fmt.Print(title)
	return title

}
