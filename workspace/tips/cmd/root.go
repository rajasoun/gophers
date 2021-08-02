package cmd

import (
	"errors"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github/gophers/tips/controller"
)

var (
	gitCmd                    = GitCommand()
	rootCmd                   = NewRootCmd()
	cmd                       *cobra.Command
	topic, arg, subarg, debug string
	about                     bool
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

			if len(args) == 0 && topic == "" && debug == "" {
				cmd.Help()
				return nil
			} else if topic != "" || debug != "" {
				//calling setUplogs func to set logger level for debugging the code
				setUpLogs(cmd.OutOrStdout(), debug)
				logrus.WithField("loglevel", debug).Debug("successfully set logger level to debug ")
				// getting topic
				input, err := getTopic(args)
				if err != nil {
					logrus.WithField("err", err).Debug("invalid user input")
					return err
				} else {
					logrus.WithField("userInput", input).Debug("successfully getting valid input ")
					controller.GetTipForTopic(input, cmd.OutOrStdout())
				}
			}
			return nil
			//return cobra.CheckErr(rootCmd.Execute())

		},
	}
	cmd.Flags().StringVarP(&topic, "topic", "c", "", "user input string help for the topic")

	return cmd
}

func GitCommand() *cobra.Command {
	var gitcmd = &cobra.Command{
		Use:   "git",
		Short: "Git is a DevOps tool used for source code management.",
		Long: ` "Git is used to tracking changes in the source code,
 enabling multiple developers to work together on non-linear development"`,
		Aliases: []string{},
		Version: "0.1v",
		Example: `tips git --arg stash / --subarg stash
"Saving current state of unstaged changes to tracked files : git stash -k" `,
		Args: cobra.MaximumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 && arg == "" && subarg == "" {
				cmd.Help()
				return nil
			} else if arg != "" || subarg != "" {
				arg := arg + " " + subarg
				controller.GetTipForTopic(arg, cmd.OutOrStdout())
			}
			return nil
		},
	}
	gitcmd.Flags().StringVarP(&arg, "arg", "c", "", "argument help for the tip")
	gitcmd.Flags().StringVarP(&subarg, "subarg", "f", "", "sub argument help for the tip")

	return gitcmd
}

// Execute executes the root command.
func Execute(writer io.Writer) error {
	rootCmd.SetOutput(writer)
	return rootCmd.Execute()
	//return cobra.CheckErr(rootCmd.Execute())
	//return err
}

// getting topic with checking validation
func getTopic(args []string) (string, error) {
	userInput := topic
	if isValidInput(userInput) {
		logrus.WithField("topic", userInput).Debug("successfully validation checked")
		return userInput, nil
	}
	var validError error = errors.New("flag needs an argument:--topic & argument should be greater than 2")
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
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	return nil
}

func init() {

	cmd.PersistentFlags().StringVarP(&debug, "debug", "", "", "verbose logging")

	cmd.PersistentFlags().MarkHidden("debug")
	rootCmd.AddCommand(gitCmd)
	rootCmd.PersistentFlags().BoolVarP(&about, "about", "a", false, "persistant flag")

	//rootCmd.SetUsageTemplate(strings.Replace(cmd.UsageString(), "tips [flags]", "tips git <order_id> [flags]", 1))

	//cmd.SetHelpCommand(cmd * Command)
	//cmd.SetHelpFunc(f func( *Command , []string))
	//cmd.SetHelpTemplate("golabal")

}
