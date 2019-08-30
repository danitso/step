/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	createApplicationCmd = &cobra.Command{
		Use:     "application",
		Aliases: []string{"app"},
		Short:   "Create an application resource",
		Long:    "Create an application resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	createApplicationCmdBlobStorageSystem             = createApplicationCmd.PersistentFlags().String("blob-storage-system", "", "The blob storage system (none or s3)")
	createApplicationCmdCacheSystem                   = createApplicationCmd.PersistentFlags().String("cache-system", "", "The cache system (none, memcached or redis)")
	createApplicationCmdContinuousDeploymentsFlag     = createApplicationCmd.PersistentFlags().Bool("continuous-deployments", false, "Whether to enable continuous deployments")
	createApplicationCmdDatabaseSystem                = createApplicationCmd.PersistentFlags().String("database-system", "", "The database system (none, mongodb, mssql or mysql)")
	createApplicationCmdGitRepositoryFlag             = createApplicationCmd.PersistentFlags().StringP("git-repository", "r", "", "The git repository URL")
	createApplicationCmdGitRepositoryPasswordFlag     = createApplicationCmd.PersistentFlags().StringP("git-repository-password", "p", "", "The git password")
	createApplicationCmdGitRepositoryUsernameFlag     = createApplicationCmd.PersistentFlags().StringP("git-repository-username", "u", "", "The git username")
	createApplicationCmdGroupFlag                     = createApplicationCmd.PersistentFlags().String("group", "", "The name of the group the application belongs to")
	createApplicationCmdPackageRepositoryFlag         = createApplicationCmd.PersistentFlags().StringP("package-repository", "R", "", "The package repository URL")
	createApplicationCmdPackageRepositoryPasswordFlag = createApplicationCmd.PersistentFlags().StringP("package-repository-password", "P", "", "The package repository password")
	createApplicationCmdPackageRepositoryUsernameFlag = createApplicationCmd.PersistentFlags().StringP("package-repository-username", "U", "", "The package repository username")
	createApplicationCmdReviewDeploymentsFlag         = createApplicationCmd.PersistentFlags().Bool("review-deployments", false, "Whether to enable review deployments")
	createApplicationCmdRuntimeSystem                 = createApplicationCmd.PersistentFlags().String("runtime-system", "", "The runtime system (dotnet, nodejs or php)")
)

func init() {
	createCmd.AddCommand(createApplicationCmd)
}
