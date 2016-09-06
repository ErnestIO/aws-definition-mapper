/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// NatGateway ...
type NatGateway struct {
	Name          string `json:"name"`
	PublicNetwork string `json:"public_network"`
}

// Validate checks if a Network is valid
func (n *NatGateway) Validate(networks []Network) error {
	if n.Name == "" {
		return errors.New("Nat Gateway name should not be null")
	}

	// Check if network name is > 50 characters
	if utf8.RuneCountInString(n.Name) > AWSMAXNAME {
		return fmt.Errorf("Nat Gateway name can't be greater than %d characters", AWSMAXNAME)
	}

	if n.PublicNetwork == "" {
		return errors.New("Nat Gateway should specify a public network")
	}

	for _, nw := range networks {
		if nw.Name == n.PublicNetwork && nw.Public {
			return nil
		}
	}

	return errors.New("Nat Gateway public network is not defined")
}
