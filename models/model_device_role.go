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

	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// deviceRoleColorPattern is the validation pattern for DeviceRole.Color
var deviceRoleColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// deviceRoleSlugPattern is the validation pattern for DeviceRole.Slug
var deviceRoleSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// DeviceRole is an object. Adds support for custom fields and tags.
type DeviceRole struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// ConfigTemplate: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ConfigTemplate *NestedConfigTemplate `json:"config_template,omitempty" mapstructure:"config_template,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceCount:
	DeviceCount int32 `json:"device_count" mapstructure:"device_count"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VirtualmachineCount:
	VirtualmachineCount int32 `json:"virtualmachine_count" mapstructure:"virtualmachine_count"`
	// VmRole: Virtual machines may be assigned to this role
	VmRole bool `json:"vm_role,omitempty" mapstructure:"vm_role,omitempty"`
}

// Validate implements basic validation for this model
func (m DeviceRole) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(deviceRoleColorPattern),
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
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(deviceRoleSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m DeviceRole) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *DeviceRole) SetColor(val string) {
	m.Color = val
}

// GetConfigTemplate returns the ConfigTemplate property
func (m DeviceRole) GetConfigTemplate() *NestedConfigTemplate {
	return m.ConfigTemplate
}

// SetConfigTemplate sets the ConfigTemplate property
func (m *DeviceRole) SetConfigTemplate(val *NestedConfigTemplate) {
	m.ConfigTemplate = val
}

// GetCreated returns the Created property
func (m DeviceRole) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DeviceRole) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m DeviceRole) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *DeviceRole) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m DeviceRole) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceRole) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m DeviceRole) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *DeviceRole) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m DeviceRole) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *DeviceRole) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m DeviceRole) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DeviceRole) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m DeviceRole) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DeviceRole) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m DeviceRole) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *DeviceRole) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m DeviceRole) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *DeviceRole) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m DeviceRole) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *DeviceRole) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m DeviceRole) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *DeviceRole) SetUrl(val string) {
	m.Url = val
}

// GetVirtualmachineCount returns the VirtualmachineCount property
func (m DeviceRole) GetVirtualmachineCount() int32 {
	return m.VirtualmachineCount
}

// SetVirtualmachineCount sets the VirtualmachineCount property
func (m *DeviceRole) SetVirtualmachineCount(val int32) {
	m.VirtualmachineCount = val
}

// GetVmRole returns the VmRole property
func (m DeviceRole) GetVmRole() bool {
	return m.VmRole
}

// SetVmRole sets the VmRole property
func (m *DeviceRole) SetVmRole(val bool) {
	m.VmRole = val
}
