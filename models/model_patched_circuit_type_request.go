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

// patchedCircuitTypeRequestSlugPattern is the validation pattern for PatchedCircuitTypeRequest.Slug
var patchedCircuitTypeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedCircuitTypeRequest is an object. Adds support for custom fields and tags.
type PatchedCircuitTypeRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedCircuitTypeRequest) Validate() error {
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
			m.Slug, validation.Length(1, 100), validation.Match(patchedCircuitTypeRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PatchedCircuitTypeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedCircuitTypeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedCircuitTypeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedCircuitTypeRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m PatchedCircuitTypeRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedCircuitTypeRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m PatchedCircuitTypeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedCircuitTypeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedCircuitTypeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedCircuitTypeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
