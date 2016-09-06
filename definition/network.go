/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"net"
	"unicode/utf8"
)

// Network ...
type Network struct {
	Name       string `json:"name"`
	Subnet     string `json:"subnet"`
	Public     bool   `json:"public"`
	NatGateway string `json:"nat_gateway"`
}

// Validate checks if a Network is valid
func (n *Network) Validate() error {
	_, _, err := net.ParseCIDR(n.Subnet)
	if err != nil {
		return errors.New("Network CIDR is not valid")
	}

	if n.Name == "" {
		return errors.New("Network name should not be null")
	}

	// Check if network name is > 50 characters
	if utf8.RuneCountInString(n.Name) > AWSMAXNAME {
		return fmt.Errorf("Network name can't be greater than %d characters", AWSMAXNAME)
	}

	if n.Public && n.NatGateway != "" {
		return errors.New("Public Network should not specify a nat gateway")
	}

	return nil
}
