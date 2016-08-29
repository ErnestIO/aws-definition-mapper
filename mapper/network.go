/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ErnestIO/aws-definition-mapper/definition"
	"github.com/ErnestIO/aws-definition-mapper/output"
)

// MapNetworks : Maps the networks from a given input payload.
func MapNetworks(d definition.Definition) []output.Network {
	var networks []output.Network

	for _, network := range d.Networks {

		n := output.Network{
			Name:   d.GeneratedName() + network.Name,
			Subnet: network.Subnet,
		}

		networks = append(networks, n)
	}

	return networks
}
