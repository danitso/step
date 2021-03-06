/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize the toolchain",
		Long:  "Initialize the toolchain",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	initCmdEncryptionKeyFlag = initCmd.Flags().StringP("encryption-key", "k", "", "The encryption key")
	initCmdGitEmailFlag      = initCmd.Flags().StringP("git-email", "m", "noreply@step.tld", "The git user's email address")
	initCmdGitPasswordFlag   = initCmd.Flags().StringP("git-password", "p", "", "The git password or token")
	initCmdGitRepositoryFlag = initCmd.Flags().StringP("git-repository", "r", "", "The git repository URL")
	initCmdGitUsernameFlag   = initCmd.Flags().StringP("git-username", "u", "", "The git username")
)

func init() {
	rootCmd.AddCommand(initCmd)
	changeHelpUsageText(initCmd)
}
