/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

import (
	"net"
	"reflect"
)

// Instance ...
type Instance struct {
	InstanceAWSID   string   `json:"instance_aws_id"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	Image           string   `json:"reference_image"`
	Network         string   `json:"network_name"`
	IP              net.IP   `json:"ip"`
	AssignElasticIP bool     `json:"assign_elastic_ip"`
	KeyPair         string   `json:"key_pair"`
	SecurityGroups  []string `json:"security_groups"`
	Exists          bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (i *Instance) HasChanged(oi *Instance) bool {
	if i.Type != oi.Type {
		return true
	}

	return !reflect.DeepEqual(i.SecurityGroups, oi.SecurityGroups)
}
