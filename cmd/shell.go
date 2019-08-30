/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	shellCmd = &cobra.Command{
		Use:     "shell",
		Aliases: []string{"sh"},
		Short:   "Open a resource shell",
		Long:    "Open a resource shell",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(shellCmd)
}
