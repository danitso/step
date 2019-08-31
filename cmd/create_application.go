/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

const (
	defaultBlobStorageSize    = 8192
	defaultBuildMemoryLimit   = 2048
	defaultCacheSize          = 512
	defaultCronInterval       = 5
	defaultDatabaseMemory     = 4096
	defaultDatabaseSize       = 16384
	defaultRuntimeMemoryLimit = 128
	defaultRuntimeWorkerLimit = 8
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
	createApplicationCmdBlobStorageSizeProduction     = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-production", defaultBlobStorageSize, "The blob storage size for the production tier")
	createApplicationCmdBlobStorageSizeReview         = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-review", defaultBlobStorageSize, "The blob storage size for the review tier")
	createApplicationCmdBlobStorageSizeStaging        = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-staging", defaultBlobStorageSize, "The blob storage size for the staging tier")
	createApplicationCmdBlobStorageSizeTesting        = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-testing", defaultBlobStorageSize, "The blob storage size for the testing tier")
	createApplicationCmdBlobStorageSystem             = createApplicationCmd.PersistentFlags().String("blob-storage-system", "", "The blob storage system (s3)")
	createApplicationCmdBuildMemoryLimit              = createApplicationCmd.PersistentFlags().Uint16("build-memory-limit", defaultBuildMemoryLimit, "The memory allocation per build job")
	createApplicationCmdCacheEnabled                  = createApplicationCmd.PersistentFlags().Bool("cache-enabled", false, "Whether to enable a cache system")
	createApplicationCmdCacheSizeProduction           = createApplicationCmd.PersistentFlags().Uint16("cache-size-production", defaultCacheSize, "The cache size for the production tier")
	createApplicationCmdCacheSizeReview               = createApplicationCmd.PersistentFlags().Uint16("cache-size-review", defaultCacheSize, "The cache size for the review tier")
	createApplicationCmdCacheSizeStaging              = createApplicationCmd.PersistentFlags().Uint16("cache-size-staging", defaultCacheSize, "The cache size for the staging tier")
	createApplicationCmdCacheSizeTesting              = createApplicationCmd.PersistentFlags().Uint16("cache-size-testing", defaultCacheSize, "The cache size for the testing tier")
	createApplicationCmdCacheSystem                   = createApplicationCmd.PersistentFlags().String("cache-system", "", "The cache system (memcached or redis)")
	createApplicationCmdContinuousDeploymentsFlag     = createApplicationCmd.PersistentFlags().Bool("continuous-deployments", false, "Whether to enable continuous deployments")
	createApplicationCmdCronEnabledProduction         = createApplicationCmd.PersistentFlags().Bool("cron-enabled-production", false, "Whether to enable cron for the production tier")
	createApplicationCmdCronEnabledReview             = createApplicationCmd.PersistentFlags().Bool("cron-enabled-review", false, "Whether to enable cron for the review tier")
	createApplicationCmdCronEnabledStaging            = createApplicationCmd.PersistentFlags().Bool("cron-enabled-staging", false, "Whether to enable cron for the staging tier")
	createApplicationCmdCronEnabledTesting            = createApplicationCmd.PersistentFlags().Bool("cron-enabled-testing", false, "Whether to enable cron for the testing tier")
	createApplicationCmdCronIntervalProduction        = createApplicationCmd.PersistentFlags().Uint8("cron-interval-production", defaultCronInterval, "The delay in minutes between each cron run on the production tier")
	createApplicationCmdCronIntervalReview            = createApplicationCmd.PersistentFlags().Uint8("cron-interval-review", defaultCronInterval, "The delay in minutes between each cron run on the review tier")
	createApplicationCmdCronIntervalStaging           = createApplicationCmd.PersistentFlags().Uint8("cron-interval-staging", defaultCronInterval, "The delay in minutes between each cron run on the staging tier")
	createApplicationCmdCronIntervalTesting           = createApplicationCmd.PersistentFlags().Uint8("cron-interval-testing", defaultCronInterval, "The delay in minutes between each cron run on the testing tier")
	createApplicationCmdCronMemoryLimitProduction     = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-production", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the production tier")
	createApplicationCmdCronMemoryLimitReview         = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-review", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the review tier")
	createApplicationCmdCronMemoryLimitStaging        = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-staging", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the staging tier")
	createApplicationCmdCronMemoryLimitTesting        = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-testing", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the testing tier")
	createApplicationCmdDatabaseEnabled               = createApplicationCmd.PersistentFlags().Bool("database-enabled", false, "Whether to enable a database system")
	createApplicationCmdDatabaseMemoryLimitProduction = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-production", defaultDatabaseMemory, "The memory allocation per database for the production tier")
	createApplicationCmdDatabaseMemoryLimitReview     = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-review", defaultDatabaseMemory, "The memory allocation per database for the review tier")
	createApplicationCmdDatabaseMemoryLimitStaging    = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-staging", defaultDatabaseMemory, "The memory allocation per database for the staging tier")
	createApplicationCmdDatabaseMemoryLimitTesting    = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-testing", defaultDatabaseMemory, "The memory allocation per database for the testing tier")
	createApplicationCmdDatabaseSizeProduction        = createApplicationCmd.PersistentFlags().Uint32("database-size-production", defaultDatabaseSize, "The database size for the production tier")
	createApplicationCmdDatabaseSizeReview            = createApplicationCmd.PersistentFlags().Uint32("database-size-review", defaultDatabaseSize, "The database size for the review tier")
	createApplicationCmdDatabaseSizeStaging           = createApplicationCmd.PersistentFlags().Uint32("database-size-staging", defaultDatabaseSize, "The database size for the staging tier")
	createApplicationCmdDatabaseSizeTesting           = createApplicationCmd.PersistentFlags().Uint32("database-size-testing", defaultDatabaseSize, "The database size for the testing tier")
	createApplicationCmdDatabaseSystem                = createApplicationCmd.PersistentFlags().String("database-system", "", "The database system (mongodb, mssql or mysql)")
	createApplicationCmdGitRepositoryFlag             = createApplicationCmd.PersistentFlags().StringP("git-repository", "r", "", "The git repository URL")
	createApplicationCmdGitRepositoryPasswordFlag     = createApplicationCmd.PersistentFlags().StringP("git-repository-password", "p", "", "The git password or token")
	createApplicationCmdGitRepositoryUsernameFlag     = createApplicationCmd.PersistentFlags().StringP("git-repository-username", "u", "", "The git username")
	createApplicationCmdGroupFlag                     = createApplicationCmd.PersistentFlags().String("group", "", "The name of the group the application belongs to")
	createApplicationCmdPackageRepositoryFlag         = createApplicationCmd.PersistentFlags().StringP("package-repository", "R", "", "The package repository URL")
	createApplicationCmdPackageRepositoryPasswordFlag = createApplicationCmd.PersistentFlags().StringP("package-repository-password", "P", "", "The package repository password")
	createApplicationCmdPackageRepositoryUsernameFlag = createApplicationCmd.PersistentFlags().StringP("package-repository-username", "U", "", "The package repository username")
	createApplicationCmdReviewDeploymentsFlag         = createApplicationCmd.PersistentFlags().Bool("review-deployments", false, "Whether to enable review deployments")
	createApplicationCmdRuntimeMemoryLimitProduction  = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-production", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the production tier")
	createApplicationCmdRuntimeMemoryLimitReview      = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-review", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the review tier")
	createApplicationCmdRuntimeMemoryLimitStaging     = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-staging", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the staging tier")
	createApplicationCmdRuntimeMemoryLimitTesting     = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-testing", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the testing tier")
	createApplicationCmdRuntimeSystem                 = createApplicationCmd.PersistentFlags().String("runtime-system", "", "The runtime system (dotnet, nodejs or php)")
	createApplicationCmdRuntimeWorkerLimitProduction  = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-production", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the production tier")
	createApplicationCmdRuntimeWorkerLimitReview      = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-review", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the review tier")
	createApplicationCmdRuntimeWorkerLimitStaging     = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-staging", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the staging tier")
	createApplicationCmdRuntimeWorkerLimitTesting     = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-testing", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the testing tier")
)

func init() {
	createCmd.AddCommand(createApplicationCmd)
	changeHelpUsageText(createApplicationCmd)
}
