package cmd

import (
	"errors"
	"github/gophers/tips/controller"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	rootCmd      = NewRootCmd()
	cmd          *cobra.Command
	topic, debug string
)

const (
	validLen int    = 2
	validArg string = "git"
)

func NewRootCmd() *cobra.Command {
	cmd = &cobra.Command{
		Use:     "tips",
		Short:   "tips for command line interface function",
		Long:    "tips provides help for docker and git cli commands ",
		Aliases: []string{},
		Version: "0.1v",
		Example: `-> tips --topic stash 
->"Saving current state of unstaged changes to tracked files : git stash -k" `,
		Args: cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			//calling setUplogs func to set logger level for debugging the code
			setUpLogs(cmd.OutOrStdout(), debug)
			logrus.WithField("loglevel", debug).Debug("successfully set logger level to debug ")

			// getting topic
			input, err := getTopic(args)
			if err != nil {
				logrus.WithField("err", err).Debug("invalid user-input,topic value length should be greater than 2")
				return err
			} else {
				logrus.WithField("userInput", input).Debug("successfully getting valid input ")
				controller.GetTipForTopic(input, cmd.OutOrStdout())
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&topic, "topic", "", "user input string help for the topic")
	return cmd
}

// Execute executes the root command.
func Execute(writer io.Writer) error {
	rootCmd.SetOutput(writer)
	return rootCmd.Execute()
}

// getting topic with checking validation
func getTopic(args []string) (string, error) {
	if isValidInput(topic) {
		userInput := topic
		logrus.WithField("topic", userInput).Debug("successfully validation checked ")
		return userInput, nil
	}
	var validError error = errors.New("key length should be greater than 2")
	return "", validError
}

//checking  userinput validation
func isValidInput(userInput string) bool {
	if len(userInput) > validLen && len(userInput) != 0 || userInput == validArg {
		return true
	}
	return false
}

//setting log level
func setUpLogs(out io.Writer, level string) error {
	logrus.SetOutput(out)
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(logLevel)
	return nil
}

func init() {
	//todo how to hide this flag from help command
	cmd.PersistentFlags().StringVarP(&debug, "debug", "d", "", "verbose logging")

}
