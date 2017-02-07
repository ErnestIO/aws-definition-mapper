/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapDatacenters : Maps input datacenter to an ernest formatted datacenter
func MapDatacenters(dat definition.Datacenter) []output.Datacenter {
	var datacenters []output.Datacenter

	datacenters = append(datacenters, output.Datacenter{
		Name:            dat.Name,
		Region:          dat.Region,
		Type:            dat.Type,
		AccessKeyID:     dat.AccessKeyID,
		SecretAccessKey: dat.SecretAccessKey,
	})

	return datacenters
}

// MapDefinitionDatacenter : Maps an output datacenter into a definition datacenter
func MapDefinitionDatacenter(datacenters []output.Datacenter) definition.Datacenter {
	return definition.Datacenter{
		Name:            datacenters[0].Name,
		Type:            datacenters[0].Type,
		Region:          datacenters[0].Region,
		AccessKeyID:     datacenters[0].AccessKeyID,
		SecretAccessKey: datacenters[0].SecretAccessKey,
	}
}
