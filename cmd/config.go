/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	configCmd.Flags().String("address", "0.0.0.0", "Address to bind to")
	configCmd.Flags().String("bin-path-git", "", "Absolute path to Git executable")
	configCmd.Flags().String("bin-path-terraform", "", "Absolute path to Terraform executable")
	configCmd.Flags().StringArray("client-ip-ranges", []string{}, "IP ranges (CIDR blocks) for client connections")
	configCmd.Flags().StringArray("hook-ip-ranges", []string{}, "IP ranges (CIDR blocks) for hook invokers")
	configCmd.Flags().String("log-file", "", "Absolute path to log file")
	configCmd.Flags().String("pid-file", "", "Absolute path to PID file")
	// Port number is based on ASC(S) = 83, ASC(T) = 84, ASC(E) = 69, ASC(P) = 80 -> (8)3(8)4(6)9(8)0.
	configCmd.Flags().Uint16("port", 8868, "Port number to bind to")
	configCmd.Flags().Bool("smtp-enabled", false, "Whether to enable the SMTP features")
	configCmd.Flags().String("smtp-hostname", "", "Hostname for SMTP server")
	configCmd.Flags().String("smtp-password", "", "Password for SMTP server")
	configCmd.Flags().Uint16("smtp-port", 587, "Port number for SMTP server")
	configCmd.Flags().Bool("smtp-secure", true, "Whether to authenticate with the supplied credentials")
	configCmd.Flags().Bool("smtp-tls", true, "Whether to enable TLS")
	configCmd.Flags().String("smtp-username", "", "Username for SMTP server")
	configCmd.Flags().String("ssl-certificate", "", "Absolute path to SSL certificate")
	configCmd.Flags().String("ssl-certificate-chain", "", "Absolute path to SSL certificate chain")
	configCmd.Flags().String("ssl-certificate-key", "", "Absolute path to SSL certificate key")

	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the toolchain",
	Long:  "Configure the toolchain",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}
