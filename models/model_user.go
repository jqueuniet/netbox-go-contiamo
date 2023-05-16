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

// userUsernamePattern is the validation pattern for User.Username
var userUsernamePattern = regexp.MustCompile(`^[\w.@+-]+$`)

// User is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type User struct {
	// DateJoined:
	DateJoined time.Time `json:"date_joined,omitempty" mapstructure:"date_joined,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Email:
	Email string `json:"email,omitempty" mapstructure:"email,omitempty"`
	// FirstName:
	FirstName string `json:"first_name,omitempty" mapstructure:"first_name,omitempty"`
	// Groups:
	Groups []int32 `json:"groups,omitempty" mapstructure:"groups,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IsActive: Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive bool `json:"is_active,omitempty" mapstructure:"is_active,omitempty"`
	// IsStaff: Designates whether the user can log into this admin site.
	IsStaff bool `json:"is_staff,omitempty" mapstructure:"is_staff,omitempty"`
	// LastName:
	LastName string `json:"last_name,omitempty" mapstructure:"last_name,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Username: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username" mapstructure:"username"`
}

// Validate implements basic validation for this model
func (m User) Validate() error {
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"username": validation.Validate(
			m.Username, validation.NotNil, validation.Length(0, 150), validation.Match(userUsernamePattern),
		),
	}.Filter()
}

// GetDateJoined returns the DateJoined property
func (m User) GetDateJoined() time.Time {
	return m.DateJoined
}

// SetDateJoined sets the DateJoined property
func (m *User) SetDateJoined(val time.Time) {
	m.DateJoined = val
}

// GetDisplay returns the Display property
func (m User) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *User) SetDisplay(val string) {
	m.Display = val
}

// GetEmail returns the Email property
func (m User) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *User) SetEmail(val string) {
	m.Email = val
}

// GetFirstName returns the FirstName property
func (m User) GetFirstName() string {
	return m.FirstName
}

// SetFirstName sets the FirstName property
func (m *User) SetFirstName(val string) {
	m.FirstName = val
}

// GetGroups returns the Groups property
func (m User) GetGroups() []int32 {
	return m.Groups
}

// SetGroups sets the Groups property
func (m *User) SetGroups(val []int32) {
	m.Groups = val
}

// GetId returns the Id property
func (m User) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *User) SetId(val int32) {
	m.Id = val
}

// GetIsActive returns the IsActive property
func (m User) GetIsActive() bool {
	return m.IsActive
}

// SetIsActive sets the IsActive property
func (m *User) SetIsActive(val bool) {
	m.IsActive = val
}

// GetIsStaff returns the IsStaff property
func (m User) GetIsStaff() bool {
	return m.IsStaff
}

// SetIsStaff sets the IsStaff property
func (m *User) SetIsStaff(val bool) {
	m.IsStaff = val
}

// GetLastName returns the LastName property
func (m User) GetLastName() string {
	return m.LastName
}

// SetLastName sets the LastName property
func (m *User) SetLastName(val string) {
	m.LastName = val
}

// GetUrl returns the Url property
func (m User) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *User) SetUrl(val string) {
	m.Url = val
}

// GetUsername returns the Username property
func (m User) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *User) SetUsername(val string) {
	m.Username = val
}
