package cmd

import (
	"errors"
	"github/gophers/tips/controller"
	"io"

	"github.com/spf13/cobra"
)

var rootCmd = NewRootCmd()
var topic string

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tips",
		Short: "tips for command line function",
		Long:  "tips provides help for docker and git cli commands ",

		RunE: func(cmd *cobra.Command, args []string) error {
			user_input, _ := cmd.Flags().GetString("topic")
			if len(user_input) > 3 {
				controller.GetTipForTopic(topic, cmd.OutOrStdout())
				return nil
			}
			return errors.New("key length should be greater than 3 and not be empty")
		},
	}
	cmd.Flags().StringVar(&topic, "topic", "", "User Input String help for the topic")
	return cmd
}

// Execute executes the root command.
func Execute(writer io.Writer) error {
	rootCmd.SetOutput(writer)
	return rootCmd.Execute()
}
