// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// WritableModuleBayRequest is an object. Adds support for custom fields and tags.
type WritableModuleBayRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device" mapstructure:"device"`
	// InstalledModule:
	InstalledModule int32 `json:"installed_module" mapstructure:"installed_module"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Position: Identifier to reference when renaming installed components
	Position string `json:"position,omitempty" mapstructure:"position,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableModuleBayRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"position": validation.Validate(
			m.Position, validation.Length(0, 30),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m WritableModuleBayRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableModuleBayRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableModuleBayRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableModuleBayRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m WritableModuleBayRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *WritableModuleBayRequest) SetDevice(val int32) {
	m.Device = val
}

// GetInstalledModule returns the InstalledModule property
func (m WritableModuleBayRequest) GetInstalledModule() int32 {
	return m.InstalledModule
}

// SetInstalledModule sets the InstalledModule property
func (m *WritableModuleBayRequest) SetInstalledModule(val int32) {
	m.InstalledModule = val
}

// GetLabel returns the Label property
func (m WritableModuleBayRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *WritableModuleBayRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m WritableModuleBayRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableModuleBayRequest) SetName(val string) {
	m.Name = val
}

// GetPosition returns the Position property
func (m WritableModuleBayRequest) GetPosition() string {
	return m.Position
}

// SetPosition sets the Position property
func (m *WritableModuleBayRequest) SetPosition(val string) {
	m.Position = val
}

// GetTags returns the Tags property
func (m WritableModuleBayRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableModuleBayRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
