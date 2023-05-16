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

// patchedTagRequestColorPattern is the validation pattern for PatchedTagRequest.Color
var patchedTagRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// patchedTagRequestSlugPattern is the validation pattern for PatchedTagRequest.Slug
var patchedTagRequestSlugPattern = regexp.MustCompile(`^[-\w]+$`)

// PatchedTagRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedTagRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedTagRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(patchedTagRequestColorPattern),
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedTagRequestSlugPattern),
		),
	}.Filter()
}

// GetColor returns the Color property
func (m PatchedTagRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *PatchedTagRequest) SetColor(val string) {
	m.Color = val
}

// GetDescription returns the Description property
func (m PatchedTagRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedTagRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m PatchedTagRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedTagRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m PatchedTagRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedTagRequest) SetSlug(val string) {
	m.Slug = val
}
