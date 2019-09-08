/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	createCacheCmd = &cobra.Command{
		Use:   "cache",
		Short: "Create a cache resource",
		Long:  "Create a cache resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	createCacheCmdEnvironmentFlag = createCacheCmd.PersistentFlags().String("environment", "", "The name of the environment to create the cache for")
)

func init() {
	createCmd.AddCommand(createCacheCmd)
	changeHelpUsageText(createCacheCmd)
}
