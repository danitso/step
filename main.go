/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"github.com/danitso/step/cmd"
	"github.com/gobuffalo/packr"
)

func main() {
	// Create the box for the Terraform templates.
	templatesBox := packr.NewBox("./templates")

	// Run the command line utility.
	cmd.Execute(&templatesBox)
}
