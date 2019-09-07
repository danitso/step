/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

const (
	// ProductVersion defines the version number.
	ProductVersion = "0.1.0"
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
	rootCmdAssumeYesFlag = rootCmd.PersistentFlags().BoolP("assume-yes", "y", false, "Silently confirm any action")
	rootCmdDebugFlag     = rootCmd.PersistentFlags().BoolP("debug", "d", false, "Display debug messages")
	rootCmdVerboseFlag   = rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Display verbose messages")
	rootCmdVersionFlag   = rootCmd.Flags().BoolP("version", "V", false, "Display version information")
)

func init() {
	rootCmd.Version = ProductVersion
}

// Execute invokes the root command.
func Execute(imagesBox, templatesBox *packr.Box) {
	changeHelpUsageText(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
