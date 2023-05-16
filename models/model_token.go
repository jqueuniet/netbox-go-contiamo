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
)

// Token is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type Token struct {
	// Created:
	Created time.Time `json:"created" mapstructure:"created"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Expires:
	Expires *time.Time `json:"expires,omitempty" mapstructure:"expires,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Key:
	Key string `json:"key,omitempty" mapstructure:"key,omitempty"`
	// LastUsed:
	LastUsed *time.Time `json:"last_used,omitempty" mapstructure:"last_used,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// User: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	User NestedUser `json:"user" mapstructure:"user"`
	// WriteEnabled: Permit create/update/delete operations using this key
	WriteEnabled bool `json:"write_enabled,omitempty" mapstructure:"write_enabled,omitempty"`
}

// Validate implements basic validation for this model
func (m Token) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"key": validation.Validate(
			m.Key, validation.Length(40, 40),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"user": validation.Validate(
			m.User, validation.NotNil,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m Token) GetCreated() time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Token) SetCreated(val time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m Token) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Token) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Token) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Token) SetDisplay(val string) {
	m.Display = val
}

// GetExpires returns the Expires property
func (m Token) GetExpires() *time.Time {
	return m.Expires
}

// SetExpires sets the Expires property
func (m *Token) SetExpires(val *time.Time) {
	m.Expires = val
}

// GetId returns the Id property
func (m Token) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Token) SetId(val int32) {
	m.Id = val
}

// GetKey returns the Key property
func (m Token) GetKey() string {
	return m.Key
}

// SetKey sets the Key property
func (m *Token) SetKey(val string) {
	m.Key = val
}

// GetLastUsed returns the LastUsed property
func (m Token) GetLastUsed() *time.Time {
	return m.LastUsed
}

// SetLastUsed sets the LastUsed property
func (m *Token) SetLastUsed(val *time.Time) {
	m.LastUsed = val
}

// GetUrl returns the Url property
func (m Token) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Token) SetUrl(val string) {
	m.Url = val
}

// GetUser returns the User property
func (m Token) GetUser() NestedUser {
	return m.User
}

// SetUser sets the User property
func (m *Token) SetUser(val NestedUser) {
	m.User = val
}

// GetWriteEnabled returns the WriteEnabled property
func (m Token) GetWriteEnabled() bool {
	return m.WriteEnabled
}

// SetWriteEnabled sets the WriteEnabled property
func (m *Token) SetWriteEnabled(val bool) {
	m.WriteEnabled = val
}
