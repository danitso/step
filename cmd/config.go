/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure the toolchain",
		Long:  "Configure the toolchain",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	configCmdAddressFlag             = configCmd.Flags().String("address", "0.0.0.0", "Address to bind to")
	configCmdBinPathGitFlag          = configCmd.Flags().String("bin-path-git", "", "Absolute path to Git executable")
	configCmdBinPathTerraformFlag    = configCmd.Flags().String("bin-path-terraform", "", "Absolute path to Terraform executable")
	configCmdClientIPRangesFlag      = configCmd.Flags().StringArray("client-ip-ranges", []string{}, "IP ranges (CIDR blocks) for client connections")
	configCmdHookIPRangesFlag        = configCmd.Flags().StringArray("hook-ip-ranges", []string{}, "IP ranges (CIDR blocks) for hook services")
	configCmdLogFileFlag             = configCmd.Flags().String("log-file", "", "Absolute path to log file")
	configCmdPIDFileFlag             = configCmd.Flags().String("pid-file", "", "Absolute path to PID file")
	configCmdPortFlag                = configCmd.Flags().Uint16("port", 8868, "Port number to bind to")
	configCmdSMTPEnabledFlag         = configCmd.Flags().Bool("smtp-enabled", false, "Whether to enable the SMTP features")
	configCmdSMTPHostnameFlag        = configCmd.Flags().String("smtp-hostname", "", "Hostname for SMTP server")
	configCmdSMTPPasswordFlag        = configCmd.Flags().String("smtp-password", "", "Password for SMTP server")
	configCmdSMTPPortFlag            = configCmd.Flags().Uint16("smtp-port", 587, "Port number for SMTP server")
	configCmdSMTPSecureFlag          = configCmd.Flags().Bool("smtp-secure", true, "Whether to authenticate with the supplied credentials")
	configCmdSMTPTLSFlag             = configCmd.Flags().Bool("smtp-tls", true, "Whether to enable TLS")
	configCmdSMTPUsernameFlag        = configCmd.Flags().String("smtp-username", "", "Username for SMTP server")
	configCmdSSLCertificateFlag      = configCmd.Flags().String("ssl-certificate", "", "Absolute path to SSL certificate or Base64 encoded string")
	configCmdSSLCertificateChainFlag = configCmd.Flags().String("ssl-certificate-chain", "", "Absolute path to SSL certificate chain or Base64 encoded string")
	configCmdSSLCertificateKeyFlag   = configCmd.Flags().String("ssl-certificate-key", "", "Absolute path to SSL certificate key or Base64 encoded string")
)

func init() {
	rootCmd.AddCommand(configCmd)
}
