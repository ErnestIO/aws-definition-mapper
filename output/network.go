/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// Network ...
type Network struct {
	NetworkAWSID string `json:"network_aws_id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Subnet       string `json:"range"`
	IsPublic     bool   `json:"is_public"`
	Service      string `json:"service"`
	Exists       bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (n *Network) HasChanged(on *Network) bool {
	return !reflect.DeepEqual(*n, *on)
}
