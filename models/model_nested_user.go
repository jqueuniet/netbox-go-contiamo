// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"regexp"
)

// nestedUserUsernamePattern is the validation pattern for NestedUser.Username
var nestedUserUsernamePattern = regexp.MustCompile(`^[\w.@+-]+$`)

// NestedUser is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedUser struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Username: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username" mapstructure:"username"`
}

// Validate implements basic validation for this model
func (m NestedUser) Validate() error {
	return validation.Errors{
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"username": validation.Validate(
			m.Username, validation.NotNil, validation.Length(0, 150), validation.Match(nestedUserUsernamePattern),
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedUser) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedUser) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedUser) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedUser) SetId(val int32) {
	m.Id = val
}

// GetUrl returns the Url property
func (m NestedUser) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedUser) SetUrl(val string) {
	m.Url = val
}

// GetUsername returns the Username property
func (m NestedUser) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *NestedUser) SetUsername(val string) {
	m.Username = val
}
