/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	pushCmd = &cobra.Command{
		Use:   "push",
		Short: "Push changes for the toolchain",
		Long:  "Push changes for the toolchain",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	pushCmdForceFlag = pushCmd.Flags().BoolP("force", "f", false, "Forcefully push local changes by overwriting remote changes")
)

func init() {
	rootCmd.AddCommand(pushCmd)
	changeHelpUsageText(pushCmd)
}
