/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"reflect"
	"strings"

	"github.com/ernestio/aws-definition-mapper/output"
)

// ComponentsByTag : Filters a given component array by a tag
func ComponentsByTag(components interface{}, key, value string) []output.Component {
	var c []output.Component

	switch reflect.TypeOf(components).Kind() {
	case reflect.Slice:
		cs := reflect.ValueOf(components)

		for i := 0; i < cs.Len(); i++ {
			x := cs.Index(i).Interface()

			cp, ok := x.(output.Component)
			if ok != true {
				return nil
			}

			tags := cp.GetTags()
			if tags[key] == value {
				c = append(c, cp)
			}
		}
	}

	return c
}

// ComponentGroups : Lists all component group names
func ComponentGroups(components interface{}, key string) []string {
	var groups []string

	switch reflect.TypeOf(components).Kind() {
	case reflect.Slice:
		cs := reflect.ValueOf(components)

		for i := 0; i < cs.Len(); i++ {
			x := cs.Index(i).Interface()

			c, ok := x.(output.Component)
			if ok != true {
				return nil
			}

			tags := c.GetTags()
			groups = appendStringUnique(groups, tags[key])
		}
	}

	return groups
}

// ComponentNamesFromIDs : Returns all names for a given input of component id's
func ComponentNamesFromIDs(components interface{}, ids []string) []string {
	var names []string

	for _, id := range ids {
		c := ComponentByID(components, id)
		if c != nil {
			names = append(names, c.ComponentName())
		}
	}

	return names
}

// ComponentsByIDs : Get components by their provider id's
func ComponentsByIDs(components interface{}, ids []string) []output.Component {
	var cs []output.Component

	for _, id := range ids {
		c := ComponentByID(components, id)
		if c != nil {
			cs = append(cs, c)
		}
	}

	return cs
}

// ComponentByID : Get a component by its provider id
func ComponentByID(components interface{}, id string) output.Component {
	switch reflect.TypeOf(components).Kind() {
	case reflect.Slice:
		cs := reflect.ValueOf(components)

		for i := 0; i < cs.Len(); i++ {
			x := cs.Index(i).Interface()

			c, ok := x.(output.Component)
			if ok != true {
				return nil
			}

			if c.ProviderID() == id {
				return c
			}
		}
	}

	return nil
}

// ShortNames removes a prefix from a list of components full names
func ShortNames(names []string, prefix string) []string {
	for i := 0; i < len(names); i++ {
		names[i] = ShortName(names[i], prefix)
	}

	return names
}

// ShortName removes a prefix from a components full name
func ShortName(name string, prefix string) string {
	return strings.Replace(name, prefix, "", -1)
}

func appendStringUnique(s []string, item string) []string {
	for _, i := range s {
		if i == item {
			return s
		}
	}

	return append(s, item)
}
