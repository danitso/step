/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"github.com/danitso/step/cmd"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	// Create the boxes for the Docker images and Terraform templates.
	imagesBox := packr.New("images", "./images")
	templatesBox := packr.New("templates", "./templates")

	// Run the command line utility.
	cmd.Execute(imagesBox, templatesBox)
}
