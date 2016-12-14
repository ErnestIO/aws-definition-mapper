/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import "errors"

// EBSVolume ...
type EBSVolume struct {
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Size             *int64  `json:"size"`
	Iops             *int64  `json:"iops"`
	Count            int     `json:"count"`
	Encrypted        bool    `json:"encrypted"`
	EncryptionKeyID  *string `json:"encryption_key_id"`
	AvailabilityZone string  `json:"availability_zone"`
}

// Validate the ebs volume
func (v *EBSVolume) Validate() error {
	if v.Name == "" {
		return errors.New("EBS Volume name should not be null")
	}

	if v.AvailabilityZone == "" {
		return errors.New("EBS Volume availability zone name should not be null")
	}

	if v.Type == "" {
		return errors.New("EBS Volume type should not be null")
	}

	if v.Encrypted && v.EncryptionKeyID == nil {
		return errors.New("EBS Volume encryption key id (KMS key id) should be set if volume is encrypted")
	}

	if v.Type != "io1" && v.Iops != nil {
		return errors.New("EBS Volume type must be 'io1' when specifying iops")
	}

	if v.Size != nil {
		if *v.Size < 1 || *v.Size > 16384 {
			return errors.New("EBS Volume size should be between 1 - 16385 (GB)")
		}
	}

	if v.Count < 1 {
		return errors.New("EBS volume count should not be less than 1")
	}

	return nil
}
