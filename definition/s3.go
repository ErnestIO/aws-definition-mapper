/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed With this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package definition

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// S3Grantee ...
type S3Grantee struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Permissions string `json:"permissions"`
}

// S3 ...
type S3 struct {
	Name           string      `json:"name"`
	ACL            string      `json:"acl"`
	BucketLocation string      `json:"bucket_location"`
	Grantees       []S3Grantee `json:"grantees"`
}

// Validate checks if a Network is valid
func (s *S3) Validate() error {
	granteeTypes := []string{"id", "emailaddress", "uri", "canonicaluser"}
	permissionTypes := []string{"full_control", "write", "write_acp", "read", "read_acp"}
	aclTypes := []string{"private", "public-read", "public-read-write", "aws-exec-read", "authenticated-read", "log-delivery-write"}

	if s.Name == "" {
		return errors.New("S3 bucket name should not be null")
	}

	// Check if s3 bucket name is > 50 characters
	if utf8.RuneCountInString(s.Name) > AWSMAXNAME {
		return fmt.Errorf("S3 bucket name can't be greater than %d characters", AWSMAXNAME)
	}

	if s.BucketLocation == "" {
		return errors.New("S3 bucket location should not be null")
	}

	if s.ACL != "" && len(s.Grantees) > 0 {
		return errors.New("S3 bucket must specify either acl or grantees, not both")
	}

	if s.ACL != "" && isOneOf(aclTypes, s.ACL) == false {
		return fmt.Errorf("S3 bucket ACL (%s) is not valid. Must be one of [%s]", s.ACL, strings.Join(aclTypes, " | "))
	}

	for _, g := range s.Grantees {
		if isOneOf(granteeTypes, g.Type) == false {
			return fmt.Errorf("S3 grantee type (%s) is invalid", g.Type)
		}

		if g.ID == "" {
			return fmt.Errorf("S3 grantee id should not be null")
		}

		if isOneOf(permissionTypes, g.Permissions) == false {
			return fmt.Errorf("S3 grantee permissions (%s) is not valid. Must be one of [%s]", s.ACL, strings.Join(permissionTypes, " | "))
		}
	}

	return nil
}
