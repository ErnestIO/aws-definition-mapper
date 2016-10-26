/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// FirewallRule ...
type FirewallRule struct {
	IP       string `json:"ip"`
	From     int    `json:"from_port"`
	To       int    `json:"to_port"`
	Protocol string `json:"protocol"`
}
