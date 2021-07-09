/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or v  implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// todo test
// gitCmd represents the git command
// var gitCmd = &cobra.Command{
// 	Use:   "git <command>",
// 	Short: "Git is a command line tool",
// 	Long:  `Github is a web-based platform used for version control.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if len(args) == 0 {
// 			logrus.WithFields(logrus.Fields{"tech": "git", "help": "tips git --help/-h"}).Info("tips git <arguments> , argument(commands i.e push ,pull...) should not be empty")
// 		} else {
// 			input := strings.NewReader(args[0])
// 			run(input, os.Stdout)
// 		}
// 	},
// }

func init() {
	rootCmd.AddCommand(gitCmd())
	gitCmd().Flags().BoolP("all", "a", false, "show all commands")
}

func gitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "git <command>",
		Short: "Git is a command line tool",
		Long:  `Github is a web-based platform used for version control.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				logrus.WithFields(logrus.Fields{"tech": "git", "help": "tips git --help/-h"}).Info("tips git <arguments> , argument(commands i.e push ,pull...) should not be empty")
			} else {
				input := strings.NewReader(args[0])
				run(input, os.Stdout)
			}
		},
	}
}
