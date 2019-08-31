/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var (
	productVersion = ""
	rootCmd        = &cobra.Command{
		Use:   "step",
		Short: "STEP is a Simplified Toolchain for Extensible Platforms",
		Long: heredoc.Doc(`
			STEP is a simplified toolchain for extensible platforms (micro service architecture).
		`),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	rootCmdVersionFlag = rootCmd.Flags().BoolP("version", "V", false, "print version information")
)

// Execute invokes the root command.
func Execute(version string) {
	rootCmd.Version = version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
