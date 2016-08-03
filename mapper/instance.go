/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"net"
	"strconv"

	"github.com/r3labs/aws-definition-mapper/definition"
	"github.com/r3labs/aws-definition-mapper/output"
)

// MapInstances : Maps the instances for the input payload on a ernest internal format
func MapInstances(d definition.Definition) []output.Instance {
	var instances []output.Instance

	for _, instance := range d.Instances {
		ip := make(net.IP, net.IPv4len)
		copy(ip, instance.StartIP.To4())

		for i := 0; i < instance.Count; i++ {
			newInstance := output.Instance{
				Name:           d.GeneratedName() + instance.Name + "-" + strconv.Itoa(i+1),
				Type:           instance.Type,
				Image:          instance.Image,
				Network:        d.GeneratedName() + instance.Network,
				IP:             net.ParseIP(ip.String()),
				SecurityGroups: instance.SecurityGroups,
			}

			instances = append(instances, newInstance)

			// Increment IP address
			ip[3]++
		}
	}
	return instances
}
