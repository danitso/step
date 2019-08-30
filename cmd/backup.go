/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "Backup the data for a resource",
		Long:  "Backup the data for a resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	backupCmdPathFlag = backupCmd.PersistentFlags().String("path", "", "Absolute path to backup file")
	backupCmdPipeFlag = backupCmd.PersistentFlags().Bool("pipe", false, "Whether to write the backup data to STDOUT")
)

func init() {
	rootCmd.AddCommand(backupCmd)
}
