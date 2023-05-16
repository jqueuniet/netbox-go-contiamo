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

// patchedWritableRegionRequestSlugPattern is the validation pattern for PatchedWritableRegionRequest.Slug
var patchedWritableRegionRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedWritableRegionRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type PatchedWritableRegionRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableRegionRequest) Validate() error {
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
			m.Slug, validation.Length(1, 100), validation.Match(patchedWritableRegionRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableRegionRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableRegionRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableRegionRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableRegionRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m PatchedWritableRegionRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableRegionRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m PatchedWritableRegionRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *PatchedWritableRegionRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m PatchedWritableRegionRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedWritableRegionRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedWritableRegionRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableRegionRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
