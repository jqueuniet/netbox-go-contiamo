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

// inventoryItemRoleRequestColorPattern is the validation pattern for InventoryItemRoleRequest.Color
var inventoryItemRoleRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// inventoryItemRoleRequestSlugPattern is the validation pattern for InventoryItemRoleRequest.Slug
var inventoryItemRoleRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// InventoryItemRoleRequest is an object. Adds support for custom fields and tags.
type InventoryItemRoleRequest struct {
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
func (m InventoryItemRoleRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(inventoryItemRoleRequestColorPattern),
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(inventoryItemRoleRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m InventoryItemRoleRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *InventoryItemRoleRequest) SetColor(val string) {
	m.Color = val
}

// GetCustomFields returns the CustomFields property
func (m InventoryItemRoleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *InventoryItemRoleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m InventoryItemRoleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *InventoryItemRoleRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m InventoryItemRoleRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *InventoryItemRoleRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m InventoryItemRoleRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *InventoryItemRoleRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m InventoryItemRoleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *InventoryItemRoleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
