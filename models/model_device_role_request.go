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

// deviceRoleRequestColorPattern is the validation pattern for DeviceRoleRequest.Color
var deviceRoleRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// deviceRoleRequestSlugPattern is the validation pattern for DeviceRoleRequest.Slug
var deviceRoleRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// DeviceRoleRequest is an object. Adds support for custom fields and tags.
type DeviceRoleRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// ConfigTemplate: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ConfigTemplate *NestedConfigTemplateRequest `json:"config_template,omitempty" mapstructure:"config_template,omitempty"`
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
func (m DeviceRoleRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(1, 6), validation.Match(deviceRoleRequestColorPattern),
		),
		"configTemplate": validation.Validate(
			m.ConfigTemplate,
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(deviceRoleRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m DeviceRoleRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *DeviceRoleRequest) SetColor(val string) {
	m.Color = val
}

// GetConfigTemplate returns the ConfigTemplate property
func (m DeviceRoleRequest) GetConfigTemplate() *NestedConfigTemplateRequest {
	return m.ConfigTemplate
}

// SetConfigTemplate sets the ConfigTemplate property
func (m *DeviceRoleRequest) SetConfigTemplate(val *NestedConfigTemplateRequest) {
	m.ConfigTemplate = val
}

// GetCustomFields returns the CustomFields property
func (m DeviceRoleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *DeviceRoleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m DeviceRoleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceRoleRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m DeviceRoleRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *DeviceRoleRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m DeviceRoleRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *DeviceRoleRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m DeviceRoleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *DeviceRoleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetVmRole returns the VmRole property
func (m DeviceRoleRequest) GetVmRole() bool {
	return m.VmRole
}

// SetVmRole sets the VmRole property
func (m *DeviceRoleRequest) SetVmRole(val bool) {
	m.VmRole = val
}
