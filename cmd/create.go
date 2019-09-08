/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a resource",
		Long:  "Create a resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	createCmdDeployFlag = createCmd.PersistentFlags().Bool("deploy", false, "Whether to deploy the resource")
	createCmdNameFlag   = createCmd.PersistentFlags().String("name", "", "The resource name")
)

func init() {
	rootCmd.AddCommand(createCmd)
	changeHelpUsageText(createCmd)
}
