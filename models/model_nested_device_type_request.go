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

// nestedDeviceTypeRequestSlugPattern is the validation pattern for NestedDeviceTypeRequest.Slug
var nestedDeviceTypeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// NestedDeviceTypeRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedDeviceTypeRequest struct {
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
}

// Validate implements basic validation for this model
func (m NestedDeviceTypeRequest) Validate() error {
	return validation.Errors{
		"model": validation.Validate(
			m.Model, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(nestedDeviceTypeRequestSlugPattern),
		),
	}.Filter()
}

// GetModel returns the Model property
func (m NestedDeviceTypeRequest) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *NestedDeviceTypeRequest) SetModel(val string) {
	m.Model = val
}

// GetSlug returns the Slug property
func (m NestedDeviceTypeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedDeviceTypeRequest) SetSlug(val string) {
	m.Slug = val
}
