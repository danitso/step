/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func changeHelpUsageText(cmd *cobra.Command) {
	currentCmd := cmd
	currentCmdParent := currentCmd.Parent()
	commandNames := []string{}

	for currentCmdParent != nil {
		commandNames = append([]string{currentCmd.Name()}, commandNames...)
		currentCmd = currentCmd.Parent()
		currentCmdParent = currentCmd.Parent()
	}

	cmd.InitDefaultHelpFlag()

	if len(commandNames) == 0 {
		cmd.Flags().Lookup("help").Usage = fmt.Sprintf("Display help text for the '%s' utility", rootCmd.Name())
	} else if len(commandNames) == 1 {
		cmd.Flags().Lookup("help").Usage = fmt.Sprintf("Display help text for the '%s' command", strings.Join(commandNames, " "))
	} else {
		cmd.Flags().Lookup("help").Usage = fmt.Sprintf("Display help text for the '%s' subcommand", strings.Join(commandNames, " "))
	}
}
