/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	daemonizeCmd = &cobra.Command{
		Use:   "daemonize",
		Short: "Daemonizes the management service",
		Long:  "Daemonizes the management service",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	daemonizeCmdBackgroundFlag = daemonizeCmd.Flags().Bool("background", false, "Whether to run in the background")
	daemonizeCmdRootFlag       = daemonizeCmd.Flags().String("root", "", "The absolute path to the toolchain directory")
)

func init() {
	rootCmd.AddCommand(daemonizeCmd)
}
