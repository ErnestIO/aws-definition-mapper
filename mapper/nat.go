/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ErnestIO/aws-definition-mapper/definition"
	"github.com/ErnestIO/aws-definition-mapper/output"
)

// MapNats : Generates necessary nats rules for input networks
func MapNats(d definition.Definition) []output.Nat {
	var nats []output.Nat

	// All Outbound Nat rules for networks
	for _, network := range d.Networks {
		if network.Public {
			continue
		}

		nats = append(nats, output.Nat{
			Name:    d.GeneratedName() + network.Name,
			Network: d.GeneratedName() + network.Name,
		})
	}

	return nats
}
