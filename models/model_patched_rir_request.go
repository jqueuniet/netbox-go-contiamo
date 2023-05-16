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

// patchedRIRRequestSlugPattern is the validation pattern for PatchedRIRRequest.Slug
var patchedRIRRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedRIRRequest is an object. Adds support for custom fields and tags.
type PatchedRIRRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// IsPrivate: IP space managed by this RIR is considered private
	IsPrivate bool `json:"is_private,omitempty" mapstructure:"is_private,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedRIRRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedRIRRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PatchedRIRRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedRIRRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedRIRRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedRIRRequest) SetDescription(val string) {
	m.Description = val
}

// GetIsPrivate returns the IsPrivate property
func (m PatchedRIRRequest) GetIsPrivate() bool {
	return m.IsPrivate
}

// SetIsPrivate sets the IsPrivate property
func (m *PatchedRIRRequest) SetIsPrivate(val bool) {
	m.IsPrivate = val
}

// GetName returns the Name property
func (m PatchedRIRRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedRIRRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m PatchedRIRRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedRIRRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedRIRRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedRIRRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
