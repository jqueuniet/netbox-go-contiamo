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

// nestedTagRequestColorPattern is the validation pattern for NestedTagRequest.Color
var nestedTagRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// nestedTagRequestSlugPattern is the validation pattern for NestedTagRequest.Slug
var nestedTagRequestSlugPattern = regexp.MustCompile(`^[-\w]+$`)

// NestedTagRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedTagRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
}

// Validate implements basic validation for this model
func (m NestedTagRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(nestedTagRequestColorPattern),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(nestedTagRequestSlugPattern),
		),
	}.Filter()
}

// GetColor returns the Color property
func (m NestedTagRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *NestedTagRequest) SetColor(val string) {
	m.Color = val
}

// GetName returns the Name property
func (m NestedTagRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedTagRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m NestedTagRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedTagRequest) SetSlug(val string) {
	m.Slug = val
}
