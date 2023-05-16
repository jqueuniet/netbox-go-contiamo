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

// rIRRequestSlugPattern is the validation pattern for RIRRequest.Slug
var rIRRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// RIRRequest is an object. Adds support for custom fields and tags.
type RIRRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// IsPrivate: IP space managed by this RIR is considered private
	IsPrivate bool `json:"is_private,omitempty" mapstructure:"is_private,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m RIRRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(rIRRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m RIRRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RIRRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RIRRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RIRRequest) SetDescription(val string) {
	m.Description = val
}

// GetIsPrivate returns the IsPrivate property
func (m RIRRequest) GetIsPrivate() bool {
	return m.IsPrivate
}

// SetIsPrivate sets the IsPrivate property
func (m *RIRRequest) SetIsPrivate(val bool) {
	m.IsPrivate = val
}

// GetName returns the Name property
func (m RIRRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RIRRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m RIRRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *RIRRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m RIRRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RIRRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
