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
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	Image          string   `json:"reference_image"`
	Network        string   `json:"network_name"`
	IP             net.IP   `json:"ip"`
	SecurityGroups []string `json:"security_groups"`
	Exists         bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (i *Instance) HasChanged(oi *Instance) bool {
	return !reflect.DeepEqual(*i, *oi)
}
