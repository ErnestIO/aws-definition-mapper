/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// EBSVolume ...
type EBSVolume struct {
	ProviderType     string `json:"_type"`
	VPCID            string `json:"vpc_id"`
	DatacenterName   string `json:"datacenter_name"`
	DatacenterType   string `json:"datacenter_type"`
	DatacenterRegion string `json:"datacenter_region"`
	AccessKeyID      string `json:"aws_access_key_id"`
	SecretAccessKey  string `json:"aws_secret_access_key"`
	VolumeAWSID      string `json:"volume_aws_id"`
	Name             string `json:"name"`
	AvailabilityZone string `json:"availability_zone"`
	VolumeType       string `json:"volume_type"`
	Size             *int64 `json:"size"`
	Iops             *int64 `json:"iops"`
	Encrypted        bool   `json:"encrypted"`
	EncryptionKeyID  string `json:"encryption_key_id"`
	Status           string `json:"status"`
	Exists           bool
}

// HasChanged diff's the two items and returns true if there have been any changes
func (v *EBSVolume) HasChanged(ov *EBSVolume) bool {
	return false
}
