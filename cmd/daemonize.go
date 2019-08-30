/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	daemonizeCmd.Flags().Bool("background", false, "Whether to run in the background")
	daemonizeCmd.Flags().String("root", rootDir, "Absolute path to the toolchain directory")

	rootCmd.AddCommand(daemonizeCmd)
}

var daemonizeCmd = &cobra.Command{
	Use:   "daemonize",
	Short: "Daemonizes the toolchain",
	Long:  "Daemonizes the toolchain",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}
