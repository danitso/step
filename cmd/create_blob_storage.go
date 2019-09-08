/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	createBlobStorageCmd = &cobra.Command{
		Use:     "blob-storage",
		Aliases: []string{"blob"},
		Short:   "Create a blob storage resource",
		Long:    "Create a blob storage resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	createBlobStorageCmdEnvironmentFlag = createBlobStorageCmd.PersistentFlags().String("environment", "", "The name of the environment to create the blob storage for")
)

func init() {
	createCmd.AddCommand(createBlobStorageCmd)
	changeHelpUsageText(createBlobStorageCmd)
}
