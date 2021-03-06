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
		AccessKeyID:      p.Datacenter.AccessKeyID,
		SecretAccessKey:  p.Datacenter.SecretAccessKey,
		VpcID:            p.Service.VpcID,
		VpcSubnet:        p.Service.VpcSubnet,
		Tags:             mapTags(p.Datacenter.Name, p.Service.Name),
		Type:             `$(datacenters.items.0.type)`,
	})
}

func randStr(size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// MapDefinitionVPC : Maps output vpc into a definition defined vpc id and subnet
func MapDefinitionVPC(m *output.FSMMessage) string {
	for i := len(m.VPCs.Items) - 1; i >= 0; i-- {
		n := m.VPCs.Items[i]

		if n.VpcID == "" {
			m.VPCs.Items = append(m.VPCs.Items[:i], m.VPCs.Items[i+1:]...)
			continue
		}
	}

	if len(m.VPCs.Items) < 1 {
		return ""
	}

	return m.VPCs.Items[0].VpcID
}
