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
	createApplicationCmdBlobStorageEnabled            = createApplicationCmd.PersistentFlags().Bool("blob-storage-enabled", false, "Whether to enable a blob storage system")
	createApplicationCmdBlobStorageSizeProduction     = createApplicationCmd.PersistentFlags().Uint64("blob-storage-size-production", 8589934592, "The blob storage size for the production tier")
	createApplicationCmdBlobStorageSizeReview         = createApplicationCmd.PersistentFlags().Uint64("blob-storage-size-review", 8589934592, "The blob storage size for the review tier")
	createApplicationCmdBlobStorageSizeStaging        = createApplicationCmd.PersistentFlags().Uint64("blob-storage-size-staging", 8589934592, "The blob storage size for the staging tier")
	createApplicationCmdBlobStorageSizeTesting        = createApplicationCmd.PersistentFlags().Uint64("blob-storage-size-testing", 8589934592, "The blob storage size for the testing tier")
	createApplicationCmdBlobStorageSystem             = createApplicationCmd.PersistentFlags().String("blob-storage-system", "", "The blob storage system (s3)")
	createApplicationCmdCacheEnabled                  = createApplicationCmd.PersistentFlags().Bool("cache-enabled", false, "Whether to enable a cache system")
	createApplicationCmdCacheSizeProduction           = createApplicationCmd.PersistentFlags().Uint64("cache-size-production", 536870912, "The cache size for the production tier")
	createApplicationCmdCacheSizeReview               = createApplicationCmd.PersistentFlags().Uint64("cache-size-review", 536870912, "The cache size for the review tier")
	createApplicationCmdCacheSizeStaging              = createApplicationCmd.PersistentFlags().Uint64("cache-size-staging", 536870912, "The cache size for the staging tier")
	createApplicationCmdCacheSizeTesting              = createApplicationCmd.PersistentFlags().Uint64("cache-size-testing", 536870912, "The cache size for the testing tier")
	createApplicationCmdCacheSystem                   = createApplicationCmd.PersistentFlags().String("cache-system", "", "The cache system (memcached or redis)")
	createApplicationCmdContinuousDeploymentsFlag     = createApplicationCmd.PersistentFlags().Bool("continuous-deployments", false, "Whether to enable continuous deployments")
	createApplicationCmdDatabaseEnabled               = createApplicationCmd.PersistentFlags().Bool("database-enabled", false, "Whether to enable a database system")
	createApplicationCmdDatabaseSizeProduction        = createApplicationCmd.PersistentFlags().Uint64("database-size-production", 8589934592, "The database size for the production tier")
	createApplicationCmdDatabaseSizeReview            = createApplicationCmd.PersistentFlags().Uint64("database-size-review", 8589934592, "The database size for the review tier")
	createApplicationCmdDatabaseSizeStaging           = createApplicationCmd.PersistentFlags().Uint64("database-size-staging", 8589934592, "The database size for the staging tier")
	createApplicationCmdDatabaseSizeTesting           = createApplicationCmd.PersistentFlags().Uint64("database-size-testing", 8589934592, "The database size for the testing tier")
	createApplicationCmdDatabaseSystem                = createApplicationCmd.PersistentFlags().String("database-system", "", "The database system (mongodb, mssql or mysql)")
	createApplicationCmdGitRepositoryFlag             = createApplicationCmd.PersistentFlags().StringP("git-repository", "r", "", "The git repository URL")
	createApplicationCmdGitRepositoryPasswordFlag     = createApplicationCmd.PersistentFlags().StringP("git-repository-password", "p", "", "The git password or token")
	createApplicationCmdGitRepositoryUsernameFlag     = createApplicationCmd.PersistentFlags().StringP("git-repository-username", "u", "", "The git username")
	createApplicationCmdGroupFlag                     = createApplicationCmd.PersistentFlags().String("group", "", "The name of the group the application belongs to")
	createApplicationCmdPackageRepositoryFlag         = createApplicationCmd.PersistentFlags().StringP("package-repository", "R", "", "The package repository URL")
	createApplicationCmdPackageRepositoryPasswordFlag = createApplicationCmd.PersistentFlags().StringP("package-repository-password", "P", "", "The package repository password")
	createApplicationCmdPackageRepositoryUsernameFlag = createApplicationCmd.PersistentFlags().StringP("package-repository-username", "U", "", "The package repository username")
	createApplicationCmdReviewDeploymentsFlag         = createApplicationCmd.PersistentFlags().Bool("review-deployments", false, "Whether to enable review deployments")
	createApplicationCmdRuntimeConcurrency            = createApplicationCmd.PersistentFlags().Uint16("runtime-concurrency", 4, "The runtime concurrency (number of worker threads)")
	createApplicationCmdRuntimeSystem                 = createApplicationCmd.PersistentFlags().String("runtime-system", "", "The runtime system (dotnet, nodejs or php)")
)

func init() {
	createCmd.AddCommand(createApplicationCmd)
	changeHelpUsageText(createApplicationCmd)
}
