/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "step",
	Short: "STEP is a Simplified Toolchain for Extensible Platforms",
	Long:  "STEP is a simplified all-in-one toolchain for extensible web platforms (micro service architecture). The goal is to provide a standardized workflow, which is capable of supporting both simple and complex web platforms powered by some of the most popular runtime solutions.",
	Run: func(cmd *cobra.Command, args []string) {
		// Not implemented.
	},
}

// Execute invokes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
