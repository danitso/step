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
	initCmdEncryptionKeyFlag      = initCmd.Flags().String("encryption-key", "", "Encryption key")
	initCmdRepositoryFlag         = initCmd.Flags().String("repository", "", "Repository URL")
	initCmdRepositoryPasswordFlag = initCmd.Flags().String("repository-password", "", "Repository password")
	initCmdRepositoryUsernameFlag = initCmd.Flags().String("repository-username", "", "Repository username")
)

func init() {
	rootCmd.AddCommand(initCmd)
}
