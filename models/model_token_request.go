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
)

// TokenRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type TokenRequest struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Expires:
	Expires *time.Time `json:"expires,omitempty" mapstructure:"expires,omitempty"`
	// Key:
	Key string `json:"key,omitempty" mapstructure:"key,omitempty"`
	// LastUsed:
	LastUsed *time.Time `json:"last_used,omitempty" mapstructure:"last_used,omitempty"`
	// User: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	User NestedUserRequest `json:"user" mapstructure:"user"`
	// WriteEnabled: Permit create/update/delete operations using this key
	WriteEnabled bool `json:"write_enabled,omitempty" mapstructure:"write_enabled,omitempty"`
}

// Validate implements basic validation for this model
func (m TokenRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"key": validation.Validate(
			m.Key, validation.Length(40, 40),
		),
		"user": validation.Validate(
			m.User, validation.NotNil,
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m TokenRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *TokenRequest) SetDescription(val string) {
	m.Description = val
}

// GetExpires returns the Expires property
func (m TokenRequest) GetExpires() *time.Time {
	return m.Expires
}

// SetExpires sets the Expires property
func (m *TokenRequest) SetExpires(val *time.Time) {
	m.Expires = val
}

// GetKey returns the Key property
func (m TokenRequest) GetKey() string {
	return m.Key
}

// SetKey sets the Key property
func (m *TokenRequest) SetKey(val string) {
	m.Key = val
}

// GetLastUsed returns the LastUsed property
func (m TokenRequest) GetLastUsed() *time.Time {
	return m.LastUsed
}

// SetLastUsed sets the LastUsed property
func (m *TokenRequest) SetLastUsed(val *time.Time) {
	m.LastUsed = val
}

// GetUser returns the User property
func (m TokenRequest) GetUser() NestedUserRequest {
	return m.User
}

// SetUser sets the User property
func (m *TokenRequest) SetUser(val NestedUserRequest) {
	m.User = val
}

// GetWriteEnabled returns the WriteEnabled property
func (m TokenRequest) GetWriteEnabled() bool {
	return m.WriteEnabled
}

// SetWriteEnabled sets the WriteEnabled property
func (m *TokenRequest) SetWriteEnabled(val bool) {
	m.WriteEnabled = val
}
