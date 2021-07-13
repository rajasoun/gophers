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
var (
	validError = errors.New("key length should be greater than '2 ' ")
)

const (
	validLen int    = 2
	validArg string = "git"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tips",
		Short: "tips for command line interface function",
		Long:  "tips provides help for docker and git cli commands ",
		RunE: func(cmd *cobra.Command, args []string) error {
			topic, err := getTopic(topic)
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

func getTopic(userInput string) (string, error) {
	if isValidInput(userInput) {
		return userInput, nil
	}
	return "", validError
}

func isValidInput(userInput string) bool {
	if len(userInput) > validLen || topic == validArg {
		return true
	}
	return false
}
