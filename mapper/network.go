/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapNetworks : Maps the networks from a given input payload.
func MapNetworks(d definition.Definition) []output.Network {
	var networks []output.Network

	for _, network := range d.Networks {

		n := output.Network{
			Name:               d.GeneratedName() + network.Name,
			Subnet:             network.Subnet,
			IsPublic:           network.Public,
			RouterName:         "$(routers.items.0.name)",
			RouterType:         "$(routers.items.0.type)",
			RouterIP:           "$(routers.items.0.ip)",
			ClientName:         "$(client_name)",
			DatacenterType:     "$(datacenters.items.0.type)",
			DatacenterName:     "$(datacenters.items.0.name)",
			DatacenterUsername: "$(datacenters.items.0.username)",
			DatacenterPassword: "$(datacenters.items.0.password)",
			DatacenterRegion:   "$(datacenters.items.0.region)",
			NetworkType:        "$(datacenters.items.0.type)",
			VpcID:              "$(datacentes.items.0.vpc_id)",
		}

		networks = append(networks, n)
	}

	return networks
}
