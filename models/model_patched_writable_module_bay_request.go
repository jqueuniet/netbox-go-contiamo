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

// PatchedWritableModuleBayRequest is an object. Adds support for custom fields and tags.
type PatchedWritableModuleBayRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// InstalledModule:
	InstalledModule int32 `json:"installed_module,omitempty" mapstructure:"installed_module,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Position: Identifier to reference when renaming installed components
	Position string `json:"position,omitempty" mapstructure:"position,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableModuleBayRequest) Validate() error {
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
			m.Name, validation.Length(1, 64),
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
func (m PatchedWritableModuleBayRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableModuleBayRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableModuleBayRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableModuleBayRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PatchedWritableModuleBayRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *PatchedWritableModuleBayRequest) SetDevice(val int32) {
	m.Device = val
}

// GetInstalledModule returns the InstalledModule property
func (m PatchedWritableModuleBayRequest) GetInstalledModule() int32 {
	return m.InstalledModule
}

// SetInstalledModule sets the InstalledModule property
func (m *PatchedWritableModuleBayRequest) SetInstalledModule(val int32) {
	m.InstalledModule = val
}

// GetLabel returns the Label property
func (m PatchedWritableModuleBayRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PatchedWritableModuleBayRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m PatchedWritableModuleBayRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableModuleBayRequest) SetName(val string) {
	m.Name = val
}

// GetPosition returns the Position property
func (m PatchedWritableModuleBayRequest) GetPosition() string {
	return m.Position
}

// SetPosition sets the Position property
func (m *PatchedWritableModuleBayRequest) SetPosition(val string) {
	m.Position = val
}

// GetTags returns the Tags property
func (m PatchedWritableModuleBayRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableModuleBayRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
