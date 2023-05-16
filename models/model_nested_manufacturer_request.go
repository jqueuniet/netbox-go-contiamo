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

// nestedManufacturerRequestSlugPattern is the validation pattern for NestedManufacturerRequest.Slug
var nestedManufacturerRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// NestedManufacturerRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedManufacturerRequest struct {
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
}

// Validate implements basic validation for this model
func (m NestedManufacturerRequest) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(nestedManufacturerRequestSlugPattern),
		),
	}.Filter()
}

// GetName returns the Name property
func (m NestedManufacturerRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedManufacturerRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m NestedManufacturerRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedManufacturerRequest) SetSlug(val string) {
	m.Slug = val
}
