/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"net"
	"strconv"

	"github.com/ernestio/aws-definition-mapper/definition"
	"github.com/ernestio/aws-definition-mapper/output"
)

const SSMBootstrap = `
#!/bin/bash
if which dpkg >/dev/null; then
cd /tmp
curl https://amazon-ssm-region.s3.amazonaws.com/latest/debian_amd64/amazon-ssm-agent.deb -o amazon-ssm-agent.deb
dpkg -i amazon-ssm-agent.deb
else
cd /tmp
curl https://amazon-ssm-region.s3.amazonaws.com/latest/linux_amd64/amazon-ssm-agent.rpm -o amazon-ssm-agent.rpm
yum install -y amazon-ssm-agent.rpm
fi`

// MapInstances : Maps the instances for the input payload on a ernest internal format
func MapInstances(d definition.Definition) []output.Instance {
	var instances []output.Instance

	for _, instance := range d.Instances {
		ip := make(net.IP, net.IPv4len)
		copy(ip, instance.StartIP.To4())

		for i := 0; i < instance.Count; i++ {
			var sgroups []string
			for _, sg := range instance.SecurityGroups {
				sgroups = append(sgroups, d.GeneratedName()+sg)
			}

			newInstance := output.Instance{
				Name:            d.GeneratedName() + instance.Name + "-" + strconv.Itoa(i+1),
				Type:            instance.Type,
				Image:           instance.Image,
				Network:         d.GeneratedName() + instance.Network,
				IP:              net.ParseIP(ip.String()),
				KeyPair:         instance.KeyPair,
				AssignElasticIP: instance.ElasticIP,
				SecurityGroups:  sgroups,
			}

			if d.Bootstrapping == "ssm" {
				newInstance.UserData = SSMBootstrap
			}

			instances = append(instances, newInstance)

			// Increment IP address
			ip[3]++
		}
	}
	return instances
}
