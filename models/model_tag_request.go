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

// tagRequestColorPattern is the validation pattern for TagRequest.Color
var tagRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// tagRequestSlugPattern is the validation pattern for TagRequest.Slug
var tagRequestSlugPattern = regexp.MustCompile(`^[-\w]+$`)

// TagRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type TagRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
}

// Validate implements basic validation for this model
func (m TagRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(tagRequestColorPattern),
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(tagRequestSlugPattern),
		),
	}.Filter()
}

// GetColor returns the Color property
func (m TagRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *TagRequest) SetColor(val string) {
	m.Color = val
}

// GetDescription returns the Description property
func (m TagRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *TagRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m TagRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *TagRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m TagRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *TagRequest) SetSlug(val string) {
	m.Slug = val
}
