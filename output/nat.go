/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Nat ...
type Nat struct {
	NatAWSID string `json:"nat_aws_id"`
	Name     string `json:"name"`
	Service  string `json:"service"`
	Network  string `json:"network_name"`
}

// HasChanged diff's the two items and returns true if there have been any changes
func (n *Nat) HasChanged(on *Nat) bool {
	return false
}
