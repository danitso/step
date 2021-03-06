/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:     "config",
		Aliases: []string{"cfg", "conf"},
		Short:   "Configure the toolchain",
		Long:    "Configure the toolchain",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	configCmdBinPathGitFlag          = configCmd.Flags().String("bin-path-git", "", "The absolute path to Git executable")
	configCmdBinPathTerraformFlag    = configCmd.Flags().String("bin-path-terraform", "", "The absolute path to Terraform executable")
	configCmdAddressFlag             = configCmd.Flags().String("bind-address", "0.0.0.0", "The address to bind to")
	configCmdPortFlag                = configCmd.Flags().Uint16("bind-port", 8868, "The port number to bind to")
	configCmdClientIPRangesFlag      = configCmd.Flags().StringArray("client-ip-ranges", []string{}, "The IP ranges (CIDR blocks) for client connections")
	configCmdHookIPRangesFlag        = configCmd.Flags().StringArray("hook-ip-ranges", []string{}, "The IP ranges (CIDR blocks) for hook services")
	configCmdLogFileFlag             = configCmd.Flags().String("log-file", "", "The absolute path to log file")
	configCmdPIDFileFlag             = configCmd.Flags().String("pid-file", "", "The absolute path to PID file")
	configCmdSMTPEnabledFlag         = configCmd.Flags().Bool("smtp-enabled", false, "Whether to enable the SMTP features")
	configCmdSMTPHostnameFlag        = configCmd.Flags().String("smtp-hostname", "", "The hostname for the SMTP server")
	configCmdSMTPPasswordFlag        = configCmd.Flags().String("smtp-password", "", "The password for the SMTP server")
	configCmdSMTPPortFlag            = configCmd.Flags().Uint16("smtp-port", 587, "The port number for the SMTP server")
	configCmdSMTPTLSFlag             = configCmd.Flags().Bool("smtp-tls", true, "Whether to enable TLS for SMTP connections")
	configCmdSMTPUsernameFlag        = configCmd.Flags().String("smtp-username", "", "The username for the SMTP server")
	configCmdTLSCertificateFlag      = configCmd.Flags().String("tls-certificate", "", "The absolute path to the TLS (SSL) certificate or Base64 encoded string")
	configCmdTLSCertificateChainFlag = configCmd.Flags().String("tls-certificate-chain", "", "The absolute path to the TLS (SSL) certificate chain or Base64 encoded string")
	configCmdTLSCertificateKeyFlag   = configCmd.Flags().String("tls-certificate-key", "", "The absolute path to the TLS (SSL) certificate key or Base64 encoded string")
)

func init() {
	rootCmd.AddCommand(configCmd)
	changeHelpUsageText(configCmd)
}
