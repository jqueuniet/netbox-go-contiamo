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

// roleRequestSlugPattern is the validation pattern for RoleRequest.Slug
var roleRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// RoleRequest is an object. Adds support for custom fields and tags.
type RoleRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m RoleRequest) Validate() error {
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(roleRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m RoleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RoleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RoleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RoleRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m RoleRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RoleRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m RoleRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *RoleRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m RoleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RoleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetWeight returns the Weight property
func (m RoleRequest) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *RoleRequest) SetWeight(val int32) {
	m.Weight = val
}
