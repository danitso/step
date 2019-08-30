/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package main

import (
	"github.com/danitso/step/cmd"
)

const (
	// ProductVersion defines the version number.
	ProductVersion = "0.1.0"
)

func main() {
	cmd.Execute(ProductVersion)
}
