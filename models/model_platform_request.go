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

// platformRequestSlugPattern is the validation pattern for PlatformRequest.Slug
var platformRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PlatformRequest is an object. Adds support for custom fields and tags.
type PlatformRequest struct {
	// ConfigTemplate: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ConfigTemplate *NestedConfigTemplateRequest `json:"config_template,omitempty" mapstructure:"config_template,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer *NestedManufacturerRequest `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// NapalmArgs: Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs map[string]interface{} `json:"napalm_args,omitempty" mapstructure:"napalm_args,omitempty"`
	// NapalmDriver: The name of the NAPALM driver to use when interacting with devices
	NapalmDriver string `json:"napalm_driver,omitempty" mapstructure:"napalm_driver,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PlatformRequest) Validate() error {
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
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"napalmArgs": validation.Validate(
			m.NapalmArgs,
		),
		"napalmDriver": validation.Validate(
			m.NapalmDriver, validation.Length(0, 50),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(platformRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetConfigTemplate returns the ConfigTemplate property
func (m PlatformRequest) GetConfigTemplate() *NestedConfigTemplateRequest {
	return m.ConfigTemplate
}

// SetConfigTemplate sets the ConfigTemplate property
func (m *PlatformRequest) SetConfigTemplate(val *NestedConfigTemplateRequest) {
	m.ConfigTemplate = val
}

// GetCustomFields returns the CustomFields property
func (m PlatformRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PlatformRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PlatformRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PlatformRequest) SetDescription(val string) {
	m.Description = val
}

// GetManufacturer returns the Manufacturer property
func (m PlatformRequest) GetManufacturer() *NestedManufacturerRequest {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *PlatformRequest) SetManufacturer(val *NestedManufacturerRequest) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m PlatformRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PlatformRequest) SetName(val string) {
	m.Name = val
}

// GetNapalmArgs returns the NapalmArgs property
func (m PlatformRequest) GetNapalmArgs() map[string]interface{} {
	return m.NapalmArgs
}

// SetNapalmArgs sets the NapalmArgs property
func (m *PlatformRequest) SetNapalmArgs(val map[string]interface{}) {
	m.NapalmArgs = val
}

// GetNapalmDriver returns the NapalmDriver property
func (m PlatformRequest) GetNapalmDriver() string {
	return m.NapalmDriver
}

// SetNapalmDriver sets the NapalmDriver property
func (m *PlatformRequest) SetNapalmDriver(val string) {
	m.NapalmDriver = val
}

// GetSlug returns the Slug property
func (m PlatformRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PlatformRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PlatformRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PlatformRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
