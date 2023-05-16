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

// patchedRackRoleRequestColorPattern is the validation pattern for PatchedRackRoleRequest.Color
var patchedRackRoleRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// patchedRackRoleRequestSlugPattern is the validation pattern for PatchedRackRoleRequest.Slug
var patchedRackRoleRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedRackRoleRequest is an object. Adds support for custom fields and tags.
type PatchedRackRoleRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
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
func (m PatchedRackRoleRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(patchedRackRoleRequestColorPattern),
		),
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
			m.Slug, validation.Length(1, 100), validation.Match(patchedRackRoleRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m PatchedRackRoleRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *PatchedRackRoleRequest) SetColor(val string) {
	m.Color = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedRackRoleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedRackRoleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedRackRoleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedRackRoleRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m PatchedRackRoleRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedRackRoleRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m PatchedRackRoleRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedRackRoleRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedRackRoleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedRackRoleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
