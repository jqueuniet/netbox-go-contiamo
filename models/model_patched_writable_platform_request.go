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

// patchedWritablePlatformRequestSlugPattern is the validation pattern for PatchedWritablePlatformRequest.Slug
var patchedWritablePlatformRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedWritablePlatformRequest is an object. Adds support for custom fields and tags.
type PatchedWritablePlatformRequest struct {
	// ConfigTemplate:
	ConfigTemplate *int32 `json:"config_template,omitempty" mapstructure:"config_template,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Manufacturer: Optionally limit this platform to devices of a certain manufacturer
	Manufacturer *int32 `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// NapalmArgs: Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs map[string]interface{} `json:"napalm_args,omitempty" mapstructure:"napalm_args,omitempty"`
	// NapalmDriver: The name of the NAPALM driver to use when interacting with devices
	NapalmDriver string `json:"napalm_driver,omitempty" mapstructure:"napalm_driver,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritablePlatformRequest) Validate() error {
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
		"napalmArgs": validation.Validate(
			m.NapalmArgs,
		),
		"napalmDriver": validation.Validate(
			m.NapalmDriver, validation.Length(0, 50),
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedWritablePlatformRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetConfigTemplate returns the ConfigTemplate property
func (m PatchedWritablePlatformRequest) GetConfigTemplate() *int32 {
	return m.ConfigTemplate
}

// SetConfigTemplate sets the ConfigTemplate property
func (m *PatchedWritablePlatformRequest) SetConfigTemplate(val *int32) {
	m.ConfigTemplate = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritablePlatformRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritablePlatformRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritablePlatformRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritablePlatformRequest) SetDescription(val string) {
	m.Description = val
}

// GetManufacturer returns the Manufacturer property
func (m PatchedWritablePlatformRequest) GetManufacturer() *int32 {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *PatchedWritablePlatformRequest) SetManufacturer(val *int32) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m PatchedWritablePlatformRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritablePlatformRequest) SetName(val string) {
	m.Name = val
}

// GetNapalmArgs returns the NapalmArgs property
func (m PatchedWritablePlatformRequest) GetNapalmArgs() map[string]interface{} {
	return m.NapalmArgs
}

// SetNapalmArgs sets the NapalmArgs property
func (m *PatchedWritablePlatformRequest) SetNapalmArgs(val map[string]interface{}) {
	m.NapalmArgs = val
}

// GetNapalmDriver returns the NapalmDriver property
func (m PatchedWritablePlatformRequest) GetNapalmDriver() string {
	return m.NapalmDriver
}

// SetNapalmDriver sets the NapalmDriver property
func (m *PatchedWritablePlatformRequest) SetNapalmDriver(val string) {
	m.NapalmDriver = val
}

// GetSlug returns the Slug property
func (m PatchedWritablePlatformRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedWritablePlatformRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedWritablePlatformRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritablePlatformRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
