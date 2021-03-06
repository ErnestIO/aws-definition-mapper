/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package output

// Component : Generic interface for any type of aws component
type Component interface {
	GetTags() map[string]string
	ProviderID() string
	ComponentName() string
}
