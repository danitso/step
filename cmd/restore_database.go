/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	restoreCmd.AddCommand(restoreDatabaseCmd)
}

var restoreDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Restore the data for a database resource",
	Long:  "Restore the data for a database resource",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}
