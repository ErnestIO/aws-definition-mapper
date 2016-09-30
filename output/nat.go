/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import "reflect"

// Nat : mapping of a nat component
type Nat struct {
	NatAWSID       string   `json:"nat_gateway_aws_id"`
	Name           string   `json:"name"`
	Service        string   `json:"service"`
	PublicNetwork  string   `json:"public_network"`
	RoutedNetworks []string `json:"routed_networks"`
	Status         string   `json:"status"`
}

// HasChanged diff's the two items and returns true if there have been any changes
func (n *Nat) HasChanged(on *Nat) bool {
	if !reflect.DeepEqual(n.RoutedNetworks, on.RoutedNetworks) {
		return true
	}

	return false
}

func hasNetwork(networks []string, name string) bool {
	for _, network := range networks {
		if network == name {
			return true
		}
	}

	return false
}
