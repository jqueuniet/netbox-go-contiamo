// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	"regexp"
)

// userRequestUsernamePattern is the validation pattern for UserRequest.Username
var userRequestUsernamePattern = regexp.MustCompile(`^[\w.@+-]+$`)

// UserRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type UserRequest struct {
	// DateJoined:
	DateJoined time.Time `json:"date_joined,omitempty" mapstructure:"date_joined,omitempty"`
	// Email:
	Email string `json:"email,omitempty" mapstructure:"email,omitempty"`
	// FirstName:
	FirstName string `json:"first_name,omitempty" mapstructure:"first_name,omitempty"`
	// Groups:
	Groups []int32 `json:"groups,omitempty" mapstructure:"groups,omitempty"`
	// IsActive: Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive bool `json:"is_active,omitempty" mapstructure:"is_active,omitempty"`
	// IsStaff: Designates whether the user can log into this admin site.
	IsStaff bool `json:"is_staff,omitempty" mapstructure:"is_staff,omitempty"`
	// LastName:
	LastName string `json:"last_name,omitempty" mapstructure:"last_name,omitempty"`
	// Password:
	Password string `json:"password" mapstructure:"password"`
	// Username: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username" mapstructure:"username"`
}

// Validate implements basic validation for this model
func (m UserRequest) Validate() error {
	return validation.Errors{
		"email": validation.Validate(
			m.Email, validation.Length(0, 254), is.EmailFormat,
		),
		"firstName": validation.Validate(
			m.FirstName, validation.Length(0, 150),
		),
		"groups": validation.Validate(
			m.Groups,
		),
		"lastName": validation.Validate(
			m.LastName, validation.Length(0, 150),
		),
		"password": validation.Validate(
			m.Password, validation.Required, validation.Length(1, 128),
		),
		"username": validation.Validate(
			m.Username, validation.Required, validation.Length(1, 150), validation.Match(userRequestUsernamePattern),
		),
	}.Filter()
}

// GetDateJoined returns the DateJoined property
func (m UserRequest) GetDateJoined() time.Time {
	return m.DateJoined
}

// SetDateJoined sets the DateJoined property
func (m *UserRequest) SetDateJoined(val time.Time) {
	m.DateJoined = val
}

// GetEmail returns the Email property
func (m UserRequest) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *UserRequest) SetEmail(val string) {
	m.Email = val
}

// GetFirstName returns the FirstName property
func (m UserRequest) GetFirstName() string {
	return m.FirstName
}

// SetFirstName sets the FirstName property
func (m *UserRequest) SetFirstName(val string) {
	m.FirstName = val
}

// GetGroups returns the Groups property
func (m UserRequest) GetGroups() []int32 {
	return m.Groups
}

// SetGroups sets the Groups property
func (m *UserRequest) SetGroups(val []int32) {
	m.Groups = val
}

// GetIsActive returns the IsActive property
func (m UserRequest) GetIsActive() bool {
	return m.IsActive
}

// SetIsActive sets the IsActive property
func (m *UserRequest) SetIsActive(val bool) {
	m.IsActive = val
}

// GetIsStaff returns the IsStaff property
func (m UserRequest) GetIsStaff() bool {
	return m.IsStaff
}

// SetIsStaff sets the IsStaff property
func (m *UserRequest) SetIsStaff(val bool) {
	m.IsStaff = val
}

// GetLastName returns the LastName property
func (m UserRequest) GetLastName() string {
	return m.LastName
}

// SetLastName sets the LastName property
func (m *UserRequest) SetLastName(val string) {
	m.LastName = val
}

// GetPassword returns the Password property
func (m UserRequest) GetPassword() string {
	return m.Password
}

// SetPassword sets the Password property
func (m *UserRequest) SetPassword(val string) {
	m.Password = val
}

// GetUsername returns the Username property
func (m UserRequest) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *UserRequest) SetUsername(val string) {
	m.Username = val
}
