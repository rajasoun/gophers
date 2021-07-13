package cmd

import (
	"errors"
	"github/gophers/tips/controller"
	"io"

	"github.com/spf13/cobra"
)

var (
	rootCmd         = NewRootCmd()
	topic, subtopic string
)

const (
	validLen int    = 2
	validArg string = "git"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tips",
		Short:   "tips for command line interface function",
		Long:    "tips provides help for docker and git cli commands ",
		Aliases: []string{"flags", "arguments"},
		Version: "0.1v",
		Example: `-> tips --topic stash 
->"Saving current state of unstaged changes to tracked files : git stash -k" `,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			topic, err := getTopic(args)
			if err != nil {
				return err
			}
			controller.GetTipForTopic(topic, cmd.OutOrStdout())
			return nil
		},
	}
	cmd.Flags().StringVar(&topic, "topic", "", "user input string help for the topic")
	cmd.Flags().StringVar(&subtopic, "subtopic", "", "user input string help for the sub topic")

	return cmd
}

// Execute executes the root command.
func Execute(writer io.Writer) error {
	rootCmd.SetOutput(writer)
	return rootCmd.Execute()
}

func getTopic(args []string) (string, error) {
	if isValidInput(topic) {
		userInput := topic
		return userInput, nil
	}
	var validError error = errors.New("key length should be greater than 2")
	return "", validError

}

func isValidInput(userInput string) bool {
	if len(userInput) > validLen || userInput == validArg {
		return true
	}
	return false
}
