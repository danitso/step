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
	createApplicationCmdBlobStorageEnabledFlag             = createApplicationCmd.PersistentFlags().Bool("blob-storage-enabled", false, "Whether to enable a blob storage system")
	createApplicationCmdBlobStorageSizeDevelopmentFlag     = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-development", defaultBlobStorageSize, "The blob storage size for the development tier")
	createApplicationCmdBlobStorageSizeProductionFlag      = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-production", defaultBlobStorageSize, "The blob storage size for the production tier")
	createApplicationCmdBlobStorageSizeReviewFlag          = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-review", defaultBlobStorageSize, "The blob storage size for the review tier")
	createApplicationCmdBlobStorageSizeStagingFlag         = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-staging", defaultBlobStorageSize, "The blob storage size for the staging tier")
	createApplicationCmdBlobStorageSizeTestingFlag         = createApplicationCmd.PersistentFlags().Uint32("blob-storage-size-testing", defaultBlobStorageSize, "The blob storage size for the testing tier")
	createApplicationCmdBlobStorageSystemFlag              = createApplicationCmd.PersistentFlags().String("blob-storage-system", "", "The blob storage system (s3)")
	createApplicationCmdBuildMemoryLimitFlag               = createApplicationCmd.PersistentFlags().Uint16("build-memory-limit", defaultBuildMemoryLimit, "The memory allocation per build job")
	createApplicationCmdCacheEnabledFlag                   = createApplicationCmd.PersistentFlags().Bool("cache-enabled", false, "Whether to enable a cache system")
	createApplicationCmdCacheSizeDevelopmentFlag           = createApplicationCmd.PersistentFlags().Uint16("cache-size-development", defaultCacheSize, "The cache size for the development tier")
	createApplicationCmdCacheSizeProductionFlag            = createApplicationCmd.PersistentFlags().Uint16("cache-size-production", defaultCacheSize, "The cache size for the production tier")
	createApplicationCmdCacheSizeReviewFlag                = createApplicationCmd.PersistentFlags().Uint16("cache-size-review", defaultCacheSize, "The cache size for the review tier")
	createApplicationCmdCacheSizeStagingFlag               = createApplicationCmd.PersistentFlags().Uint16("cache-size-staging", defaultCacheSize, "The cache size for the staging tier")
	createApplicationCmdCacheSizeTestingFlag               = createApplicationCmd.PersistentFlags().Uint16("cache-size-testing", defaultCacheSize, "The cache size for the testing tier")
	createApplicationCmdCacheSystemFlag                    = createApplicationCmd.PersistentFlags().String("cache-system", "", "The cache system (memcached or redis)")
	createApplicationCmdContinuousDeploymentsFlag          = createApplicationCmd.PersistentFlags().Bool("continuous-deployments", false, "Whether to enable continuous deployments")
	createApplicationCmdCronEnabledDevelopmentFlag         = createApplicationCmd.PersistentFlags().Bool("cron-enabled-development", false, "Whether to enable cron for the development tier")
	createApplicationCmdCronEnabledProductionFlag          = createApplicationCmd.PersistentFlags().Bool("cron-enabled-production", false, "Whether to enable cron for the production tier")
	createApplicationCmdCronEnabledReviewFlag              = createApplicationCmd.PersistentFlags().Bool("cron-enabled-review", false, "Whether to enable cron for the review tier")
	createApplicationCmdCronEnabledStagingFlag             = createApplicationCmd.PersistentFlags().Bool("cron-enabled-staging", false, "Whether to enable cron for the staging tier")
	createApplicationCmdCronEnabledTestingFlag             = createApplicationCmd.PersistentFlags().Bool("cron-enabled-testing", false, "Whether to enable cron for the testing tier")
	createApplicationCmdCronIntervalDevelopmentFlag        = createApplicationCmd.PersistentFlags().Uint8("cron-interval-development", defaultCronInterval, "The delay in minutes between each cron run on the development tier")
	createApplicationCmdCronIntervalProductionFlag         = createApplicationCmd.PersistentFlags().Uint8("cron-interval-production", defaultCronInterval, "The delay in minutes between each cron run on the production tier")
	createApplicationCmdCronIntervalReviewFlag             = createApplicationCmd.PersistentFlags().Uint8("cron-interval-review", defaultCronInterval, "The delay in minutes between each cron run on the review tier")
	createApplicationCmdCronIntervalStagingFlag            = createApplicationCmd.PersistentFlags().Uint8("cron-interval-staging", defaultCronInterval, "The delay in minutes between each cron run on the staging tier")
	createApplicationCmdCronIntervalTestingFlag            = createApplicationCmd.PersistentFlags().Uint8("cron-interval-testing", defaultCronInterval, "The delay in minutes between each cron run on the testing tier")
	createApplicationCmdCronMemoryLimitDevelopmentFlag     = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-development", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the development tier")
	createApplicationCmdCronMemoryLimitProductionFlag      = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-production", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the production tier")
	createApplicationCmdCronMemoryLimitReviewFlag          = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-review", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the review tier")
	createApplicationCmdCronMemoryLimitStagingFlag         = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-staging", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the staging tier")
	createApplicationCmdCronMemoryLimitTestingFlag         = createApplicationCmd.PersistentFlags().Uint16("cron-memory-limit-testing", defaultRuntimeMemoryLimit, "The memory allocation per cron worker for the testing tier")
	createApplicationCmdDatabaseEnabledFlag                = createApplicationCmd.PersistentFlags().Bool("database-enabled", false, "Whether to enable a database system")
	createApplicationCmdDatabaseMemoryLimitDevelopmentFlag = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-development", defaultDatabaseMemory, "The memory allocation per database for the development tier")
	createApplicationCmdDatabaseMemoryLimitProductionFlag  = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-production", defaultDatabaseMemory, "The memory allocation per database for the production tier")
	createApplicationCmdDatabaseMemoryLimitReviewFlag      = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-review", defaultDatabaseMemory, "The memory allocation per database for the review tier")
	createApplicationCmdDatabaseMemoryLimitStagingFlag     = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-staging", defaultDatabaseMemory, "The memory allocation per database for the staging tier")
	createApplicationCmdDatabaseMemoryLimitTestingFlag     = createApplicationCmd.PersistentFlags().Uint16("database-memory-limit-testing", defaultDatabaseMemory, "The memory allocation per database for the testing tier")
	createApplicationCmdDatabaseSizeDevelopmentFlag        = createApplicationCmd.PersistentFlags().Uint32("database-size-development", defaultDatabaseSize, "The database size for the development tier")
	createApplicationCmdDatabaseSizeProductionFlag         = createApplicationCmd.PersistentFlags().Uint32("database-size-production", defaultDatabaseSize, "The database size for the production tier")
	createApplicationCmdDatabaseSizeReviewFlag             = createApplicationCmd.PersistentFlags().Uint32("database-size-review", defaultDatabaseSize, "The database size for the review tier")
	createApplicationCmdDatabaseSizeStagingFlag            = createApplicationCmd.PersistentFlags().Uint32("database-size-staging", defaultDatabaseSize, "The database size for the staging tier")
	createApplicationCmdDatabaseSizeTestingFlag            = createApplicationCmd.PersistentFlags().Uint32("database-size-testing", defaultDatabaseSize, "The database size for the testing tier")
	createApplicationCmdDatabaseSystemFlag                 = createApplicationCmd.PersistentFlags().String("database-system", "", "The database system (mongodb, mssql or mysql)")
	createApplicationCmdGitBranchDevelopmentFlag           = createApplicationCmd.PersistentFlags().String("git-branch-development", "develop", "The name of the development branch")
	createApplicationCmdGitBranchProductionFlag            = createApplicationCmd.PersistentFlags().String("git-branch-production", "master", "The name of the production branch")
	createApplicationCmdGitBranchStagingFlag               = createApplicationCmd.PersistentFlags().String("git-branch-staging", "stage", "The name of the staging branch")
	createApplicationCmdGitBranchTestingFlag               = createApplicationCmd.PersistentFlags().String("git-branch-testing", "test", "The name of the testing branch")
	createApplicationCmdGitPasswordFlag                    = createApplicationCmd.PersistentFlags().StringP("git-password", "p", "", "The git password or token")
	createApplicationCmdGitRepositoryFlag                  = createApplicationCmd.PersistentFlags().StringP("git-repository", "r", "", "The git repository URL")
	createApplicationCmdGitUsernameFlag                    = createApplicationCmd.PersistentFlags().StringP("git-username", "u", "", "The git username")
	createApplicationCmdGroupFlag                          = createApplicationCmd.PersistentFlags().String("group", "", "The name of the group the application belongs to")
	createApplicationCmdPackageRepositoryFlag              = createApplicationCmd.PersistentFlags().StringP("package-repository", "R", "", "The package repository URL")
	createApplicationCmdPackageRepositoryPasswordFlag      = createApplicationCmd.PersistentFlags().StringP("package-repository-password", "P", "", "The package repository password")
	createApplicationCmdPackageRepositoryUsernameFlag      = createApplicationCmd.PersistentFlags().StringP("package-repository-username", "U", "", "The package repository username")
	createApplicationCmdReviewDeploymentsFlag              = createApplicationCmd.PersistentFlags().Bool("review-deployments", false, "Whether to enable review deployments")
	createApplicationCmdRuntimeMemoryLimitDevelopmentFlag  = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-development", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the development tier")
	createApplicationCmdRuntimeMemoryLimitProductionFlag   = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-production", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the production tier")
	createApplicationCmdRuntimeMemoryLimitReviewFlag       = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-review", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the review tier")
	createApplicationCmdRuntimeMemoryLimitStagingFlag      = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-staging", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the staging tier")
	createApplicationCmdRuntimeMemoryLimitTestingFlag      = createApplicationCmd.PersistentFlags().Uint16("runtime-memory-limit-testing", defaultRuntimeMemoryLimit, "The memory allocation per runtime worker for the testing tier")
	createApplicationCmdRuntimeSystemFlag                  = createApplicationCmd.PersistentFlags().String("runtime-system", "", "The runtime system (dotnet, nodejs or php)")
	createApplicationCmdRuntimeWorkerLimitDevelopmentFlag  = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-development", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the development tier")
	createApplicationCmdRuntimeWorkerLimitProductionFlag   = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-production", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the production tier")
	createApplicationCmdRuntimeWorkerLimitReviewFlag       = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-review", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the review tier")
	createApplicationCmdRuntimeWorkerLimitStagingFlag      = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-staging", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the staging tier")
	createApplicationCmdRuntimeWorkerLimitTestingFlag      = createApplicationCmd.PersistentFlags().Uint16("runtime-worker-limit-testing", defaultRuntimeWorkerLimit, "The maximum number of runtime workers for the testing tier")
	createApplicationCmdSecretDevelopmentFlag              = createApplicationCmd.PersistentFlags().StringArray("secret-development", []string{}, "A secret for the development tier in the form of key=value")
	createApplicationCmdSecretProductionFlag               = createApplicationCmd.PersistentFlags().StringArray("secret-production", []string{}, "A secret for the production tier in the form of key=value")
	createApplicationCmdSecretReviewFlag                   = createApplicationCmd.PersistentFlags().StringArray("secret-review", []string{}, "A secret for the review tier in the form of key=value")
	createApplicationCmdSecretStagingFlag                  = createApplicationCmd.PersistentFlags().StringArray("secret-staging", []string{}, "A secret for the staging tier in the form of key=value")
	createApplicationCmdSecretTestingFlag                  = createApplicationCmd.PersistentFlags().StringArray("secret-testing", []string{}, "A secret for the testing tier in the form of key=value")
	createApplicationCmdVariableDevelopmentFlag            = createApplicationCmd.PersistentFlags().StringArray("variable-development", []string{}, "A variable for the development tier in the form of key=value")
	createApplicationCmdVariableProductionFlag             = createApplicationCmd.PersistentFlags().StringArray("variable-production", []string{}, "A variable for the production tier in the form of key=value")
	createApplicationCmdVariableReviewFlag                 = createApplicationCmd.PersistentFlags().StringArray("variable-review", []string{}, "A variable for the review tier in the form of key=value")
	createApplicationCmdVariableStagingFlag                = createApplicationCmd.PersistentFlags().StringArray("variable-staging", []string{}, "A variable for the staging tier in the form of key=value")
	createApplicationCmdVariableTestingFlag                = createApplicationCmd.PersistentFlags().StringArray("variable-testing", []string{}, "A variable for the testing tier in the form of key=value")
)

func init() {
	createCmd.AddCommand(createApplicationCmd)
	changeHelpUsageText(createApplicationCmd)
}
