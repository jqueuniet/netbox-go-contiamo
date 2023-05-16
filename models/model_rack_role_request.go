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

// rackRoleRequestColorPattern is the validation pattern for RackRoleRequest.Color
var rackRoleRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// rackRoleRequestSlugPattern is the validation pattern for RackRoleRequest.Slug
var rackRoleRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// RackRoleRequest is an object. Adds support for custom fields and tags.
type RackRoleRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
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
}

// Validate implements basic validation for this model
func (m RackRoleRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(rackRoleRequestColorPattern),
		),
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(rackRoleRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m RackRoleRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *RackRoleRequest) SetColor(val string) {
	m.Color = val
}

// GetCustomFields returns the CustomFields property
func (m RackRoleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RackRoleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RackRoleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RackRoleRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m RackRoleRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RackRoleRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m RackRoleRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *RackRoleRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m RackRoleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RackRoleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
