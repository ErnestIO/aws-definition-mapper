/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"math/rand"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

// MapVPCs : Maps input vpc to an ernest formatted vpc
func MapVPCs(p *definition.Payload) []output.VPC {
	var vpcs []output.VPC

	return append(vpcs, output.VPC{
		DatacenterName:   p.Datacenter.Name,
		DatacenterRegion: p.Datacenter.Region,
		DatacenterToken:  p.Datacenter.Token,
		DatacenterSecret: p.Datacenter.Secret,
		VpcID:            p.Service.VpcID,
		VpcSubnet:        p.Service.VpcSubnet,
		Type:             `$(datacenters.items.0.type)`,
	})
}

func randStr(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}