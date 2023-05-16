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
)

// Group is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type Group struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// UserCount:
	UserCount int32 `json:"user_count" mapstructure:"user_count"`
}

// Validate implements basic validation for this model
func (m Group) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 150),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m Group) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Group) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Group) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Group) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m Group) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Group) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m Group) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Group) SetUrl(val string) {
	m.Url = val
}

// GetUserCount returns the UserCount property
func (m Group) GetUserCount() int32 {
	return m.UserCount
}

// SetUserCount sets the UserCount property
func (m *Group) SetUserCount(val int32) {
	m.UserCount = val
}
