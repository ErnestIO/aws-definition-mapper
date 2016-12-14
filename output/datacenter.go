/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Datacenter : Mapping for a datacenter component
type Datacenter struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Region          string `json:"region"`
	Type            string `json:"type"`
	ExternalNetwork string `json:"external_network"`
	AccessKeyID     string `json:"aws_access_key_id"`
	SecretAccessKey string `json:"aws_secret_access_key"`
	VCloudURL       string `json:"vcloud_url"`
	VseURL          string `json:"vse_url"`
	Status          string `json:"status"`
}

// HasChanged diff's the two items and returns true if there have been any changes
func (d *Datacenter) HasChanged(od *Datacenter) bool {
	if d.Name == od.Name &&
		d.Username == od.Username &&
		d.Password == od.Password &&
		d.Region == od.Region &&
		d.Type == od.Type &&
		d.ExternalNetwork == od.ExternalNetwork &&
		d.AccessKeyID == od.AccessKeyID &&
		d.SecretAccessKey == od.SecretAccessKey &&
		d.VCloudURL == od.VCloudURL &&
		d.VseURL == od.VseURL {
		return false
	}
	return true
}
