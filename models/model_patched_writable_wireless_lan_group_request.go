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

// patchedWritableWirelessLANGroupRequestSlugPattern is the validation pattern for PatchedWritableWirelessLANGroupRequest.Slug
var patchedWritableWirelessLANGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedWritableWirelessLANGroupRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type PatchedWritableWirelessLANGroupRequest struct {
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
func (m PatchedWritableWirelessLANGroupRequest) Validate() error {
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
			m.Slug, validation.Length(1, 100), validation.Match(patchedWritableWirelessLANGroupRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableWirelessLANGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableWirelessLANGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableWirelessLANGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableWirelessLANGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m PatchedWritableWirelessLANGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableWirelessLANGroupRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m PatchedWritableWirelessLANGroupRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *PatchedWritableWirelessLANGroupRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m PatchedWritableWirelessLANGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedWritableWirelessLANGroupRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedWritableWirelessLANGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableWirelessLANGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
