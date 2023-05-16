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

// writableDeviceRoleRequestColorPattern is the validation pattern for WritableDeviceRoleRequest.Color
var writableDeviceRoleRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// writableDeviceRoleRequestSlugPattern is the validation pattern for WritableDeviceRoleRequest.Slug
var writableDeviceRoleRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WritableDeviceRoleRequest is an object. Adds support for custom fields and tags.
type WritableDeviceRoleRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// ConfigTemplate:
	ConfigTemplate *int32 `json:"config_template,omitempty" mapstructure:"config_template,omitempty"`
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
	// VmRole: Virtual machines may be assigned to this role
	VmRole bool `json:"vm_role,omitempty" mapstructure:"vm_role,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableDeviceRoleRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(writableDeviceRoleRequestColorPattern),
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(writableDeviceRoleRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m WritableDeviceRoleRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *WritableDeviceRoleRequest) SetColor(val string) {
	m.Color = val
}

// GetConfigTemplate returns the ConfigTemplate property
func (m WritableDeviceRoleRequest) GetConfigTemplate() *int32 {
	return m.ConfigTemplate
}

// SetConfigTemplate sets the ConfigTemplate property
func (m *WritableDeviceRoleRequest) SetConfigTemplate(val *int32) {
	m.ConfigTemplate = val
}

// GetCustomFields returns the CustomFields property
func (m WritableDeviceRoleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableDeviceRoleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableDeviceRoleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableDeviceRoleRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m WritableDeviceRoleRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableDeviceRoleRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m WritableDeviceRoleRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WritableDeviceRoleRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m WritableDeviceRoleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableDeviceRoleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetVmRole returns the VmRole property
func (m WritableDeviceRoleRequest) GetVmRole() bool {
	return m.VmRole
}

// SetVmRole sets the VmRole property
func (m *WritableDeviceRoleRequest) SetVmRole(val bool) {
	m.VmRole = val
}
