/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull changes for the toolchain",
		Long:  "Pull changes for the toolchain",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	pullCmdForceFlag = pullCmd.Flags().BoolP("force", "f", false, "Forcefully pull remote changes by overwriting local changes")
)

func init() {
	rootCmd.AddCommand(pullCmd)
}
