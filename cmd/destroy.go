/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	destroyCmd = &cobra.Command{
		Use:   "destroy",
		Short: "Destroy a resource",
		Long:  "Destroy a resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	destroyCmdNameFlag = destroyCmd.PersistentFlags().String("name", "", "The resource name")
)

func init() {
	rootCmd.AddCommand(destroyCmd)
	changeHelpUsageText(destroyCmd)
}
