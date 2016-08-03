/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// FirewallRule ...
type FirewallRule struct {
	Type            string `json:"type"`
	SourceIP        string `json:"source_ip"`
	SourcePort      string `json:"source_port"`
	DestinationPort string `json:"destination_port"`
	Protocol        string `json:"protocol"`
}
