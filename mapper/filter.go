/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import "github.com/ernestio/aws-definition-mapper/output"

// ComponentsByTag : Filters a given component array by a tag
func ComponentsByTag(components interface{}, key, value string) []output.Component {
	var c []output.Component

	cs, ok := components.([]output.Component)
	if ok != true {
		return c
	}

	for _, cp := range cs {
		tags := cp.GetTags()
		if tags[key] == value {
			c = append(c, cp)
		}
	}

	return c
}

// ComponentGroups : Lists all component group names
func ComponentGroups(components interface{}, key string) []string {
	var groups []string

	cs, ok := components.([]output.Component)
	if ok != true {
		return groups
	}

	for _, c := range cs {
		tags := c.GetTags()
		groups = appendStringUnique(groups, tags[key])
	}

	return groups
}

func appendStringUnique(s []string, item string) []string {
	for _, i := range s {
		if i == item {
			return s
		}
	}

	return append(s, item)
}
