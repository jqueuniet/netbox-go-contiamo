// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// nestedUserRequestUsernamePattern is the validation pattern for NestedUserRequest.Username
var nestedUserRequestUsernamePattern = regexp.MustCompile(`^[\w.@+-]+$`)

// NestedUserRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedUserRequest struct {
	// Username: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username" mapstructure:"username"`
}

// Validate implements basic validation for this model
func (m NestedUserRequest) Validate() error {
	return validation.Errors{
		"username": validation.Validate(
			m.Username, validation.Required, validation.Length(1, 150), validation.Match(nestedUserRequestUsernamePattern),
		),
	}.Filter()
}

// GetUsername returns the Username property
func (m NestedUserRequest) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *NestedUserRequest) SetUsername(val string) {
	m.Username = val
}
