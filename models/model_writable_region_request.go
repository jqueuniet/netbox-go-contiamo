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

// writableRegionRequestSlugPattern is the validation pattern for WritableRegionRequest.Slug
var writableRegionRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WritableRegionRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type WritableRegionRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableRegionRequest) Validate() error {
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(writableRegionRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m WritableRegionRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableRegionRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableRegionRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableRegionRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m WritableRegionRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableRegionRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m WritableRegionRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *WritableRegionRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m WritableRegionRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WritableRegionRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m WritableRegionRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableRegionRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
