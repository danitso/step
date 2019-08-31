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
	initCmdEncryptionKeyFlag         = initCmd.Flags().StringP("encryption-key", "k", "", "the encryption key")
	initCmdGitRepositoryFlag         = initCmd.Flags().StringP("git-repository", "r", "", "the git repository URL")
	initCmdGitRepositoryPasswordFlag = initCmd.Flags().StringP("git-repository-password", "p", "", "the git password")
	initCmdGitRepositoryUsernameFlag = initCmd.Flags().StringP("git-repository-username", "u", "", "the git username")
)

func init() {
	rootCmd.AddCommand(initCmd)
}
