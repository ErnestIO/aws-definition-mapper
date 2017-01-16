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

// ComponentNamesFromIDs : Returns all names for a given input of component id's
func ComponentNamesFromIDs(components interface{}, ids []string) []string {
	var names []string

	for _, id := range ids {
		c := ComponentByID(components, id)
		names = append(names, c.ComponentName())
	}

	return names
}

// ComponentByID : Get a component by its provider ID
func ComponentByID(components interface{}, id string) output.Component {
	cs, ok := components.([]output.Component)
	if ok != true {
		return nil
	}

	for _, c := range cs {
		if c.ProviderID() == id {
			return c
		}
	}

	return nil
}

func appendStringUnique(s []string, item string) []string {
	for _, i := range s {
		if i == item {
			return s
		}
	}

	return append(s, item)
}
