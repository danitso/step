/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	createEnvironmentCmd = &cobra.Command{
		Use:     "environment",
		Aliases: []string{"env"},
		Short:   "Create an environment resource",
		Long:    "Create an environment resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	createEnvironmentCmdApplicationFlag = createEnvironmentCmd.PersistentFlags().String("application", "", "The name of the application to create the environment for")
	createEnvironmentCmdDevelopmentFlag = createEnvironmentCmd.PersistentFlags().Bool("development", false, "Whether to provision a development environment")
	createEnvironmentCmdProductionFlag  = createEnvironmentCmd.PersistentFlags().Bool("production", false, "Whether to provision a production environment")
	createEnvironmentCmdReviewFlag      = createEnvironmentCmd.PersistentFlags().Bool("review", false, "Whether to provision a review environment")
	createEnvironmentCmdStagingFlag     = createEnvironmentCmd.PersistentFlags().Bool("staging", false, "Whether to provision a staging environment")
	createEnvironmentCmdTestingFlag     = createEnvironmentCmd.PersistentFlags().Bool("testing", false, "Whether to provision a testing environment")
)

func init() {
	createCmd.AddCommand(createEnvironmentCmd)
	changeHelpUsageText(createEnvironmentCmd)
}
