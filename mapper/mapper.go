/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/r3labs/aws-definition-mapper/definition"
	"github.com/r3labs/aws-definition-mapper/output"
)

// ConvertPayload will build an FSMMessage based on an input definition
func ConvertPayload(p *definition.Payload) *output.FSMMessage {
	m := output.FSMMessage{
		ID:          p.ServiceID,
		Service:     p.ServiceID,
		ServiceName: p.Service.Name,
		ClientName:  p.Client.Name,
		Type:        p.Datacenter.Type,
	}

	// Map datacenters
	m.Datacenters.Items = MapDatacenters(p.Datacenter)

	// Map networks
	m.Networks.Items = MapNetworks(p.Service)

	// Map instances
	m.Instances.Items = MapInstances(p.Service)

	// Map firewalls
	m.Firewalls.Items = MapSecurityGroups(p.Service)

	// Map nats/port forwarding
	m.Nats.Items = MapNats(p.Service)

	return &m
}
