/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Firewall ...
type Firewall struct {
	ID                 string         `json:"firewall_aws_id"`
	Name               string         `json:"name"`
	SecurityGroupAWSID string         `json:"security_group_aws_id"`
	Rules              []FirewallRule `json:"rules"`
}

// HasChanged diff's the two items and returns true if there have been any changes
func (f *Firewall) HasChanged(of *Firewall) bool {
	if len(f.Rules) != len(of.Rules) {
		return true
	}

	for i := 0; i < len(f.Rules); i++ {
		if f.Rules[i].DestinationPort != of.Rules[i].DestinationPort ||
			f.Rules[i].Protocol != of.Rules[i].Protocol ||
			f.Rules[i].SourceIP != of.Rules[i].SourceIP ||
			f.Rules[i].SourcePort != of.Rules[i].SourcePort {
			return true
		}
	}

	return false
}
