// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GroupRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type GroupRequest struct {
	// Name:
	Name string `json:"name" mapstructure:"name"`
}

// Validate implements basic validation for this model
func (m GroupRequest) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 150),
		),
	}.Filter()
}

// GetName returns the Name property
func (m GroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *GroupRequest) SetName(val string) {
	m.Name = val
}
