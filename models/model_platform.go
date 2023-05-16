// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// platformSlugPattern is the validation pattern for Platform.Slug
var platformSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// Platform is an object. Adds support for custom fields and tags.
type Platform struct {
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
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer *NestedManufacturer `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// NapalmArgs: Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs map[string]interface{} `json:"napalm_args,omitempty" mapstructure:"napalm_args,omitempty"`
	// NapalmDriver: The name of the NAPALM driver to use when interacting with devices
	NapalmDriver string `json:"napalm_driver,omitempty" mapstructure:"napalm_driver,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VirtualmachineCount:
	VirtualmachineCount int32 `json:"virtualmachine_count" mapstructure:"virtualmachine_count"`
}

// Validate implements basic validation for this model
func (m Platform) Validate() error {
	return validation.Errors{
		"configTemplate": validation.Validate(
			m.ConfigTemplate,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"manufacturer": validation.Validate(
			m.Manufacturer,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"napalmArgs": validation.Validate(
			m.NapalmArgs,
		),
		"napalmDriver": validation.Validate(
			m.NapalmDriver, validation.Length(0, 50),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(platformSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetConfigTemplate returns the ConfigTemplate property
func (m Platform) GetConfigTemplate() *NestedConfigTemplate {
	return m.ConfigTemplate
}

// SetConfigTemplate sets the ConfigTemplate property
func (m *Platform) SetConfigTemplate(val *NestedConfigTemplate) {
	m.ConfigTemplate = val
}

// GetCreated returns the Created property
func (m Platform) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Platform) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Platform) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Platform) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Platform) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Platform) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m Platform) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *Platform) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m Platform) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Platform) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Platform) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Platform) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Platform) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Platform) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetManufacturer returns the Manufacturer property
func (m Platform) GetManufacturer() *NestedManufacturer {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *Platform) SetManufacturer(val *NestedManufacturer) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m Platform) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Platform) SetName(val string) {
	m.Name = val
}

// GetNapalmArgs returns the NapalmArgs property
func (m Platform) GetNapalmArgs() map[string]interface{} {
	return m.NapalmArgs
}

// SetNapalmArgs sets the NapalmArgs property
func (m *Platform) SetNapalmArgs(val map[string]interface{}) {
	m.NapalmArgs = val
}

// GetNapalmDriver returns the NapalmDriver property
func (m Platform) GetNapalmDriver() string {
	return m.NapalmDriver
}

// SetNapalmDriver sets the NapalmDriver property
func (m *Platform) SetNapalmDriver(val string) {
	m.NapalmDriver = val
}

// GetSlug returns the Slug property
func (m Platform) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Platform) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m Platform) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Platform) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m Platform) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Platform) SetUrl(val string) {
	m.Url = val
}

// GetVirtualmachineCount returns the VirtualmachineCount property
func (m Platform) GetVirtualmachineCount() int32 {
	return m.VirtualmachineCount
}

// SetVirtualmachineCount sets the VirtualmachineCount property
func (m *Platform) SetVirtualmachineCount(val int32) {
	m.VirtualmachineCount = val
}
