/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	restoreCmd.PersistentFlags().String("path", "", "Absolute path to backup file")
	restoreCmd.PersistentFlags().Bool("pipe", false, "Whether to retrieve the backup data from STDIN")

	rootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore the data for a resource",
	Long:  "Restore the data for a resource",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}
