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

// PatchedWritableModuleBayTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableModuleBayTemplateRequest struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType:
	DeviceType int32 `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Position: Identifier to reference when renaming installed components
	Position string `json:"position,omitempty" mapstructure:"position,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableModuleBayTemplateRequest) Validate() error {
	return validation.Errors{
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
	}.Filter()
}

// GetDescription returns the Description property
func (m PatchedWritableModuleBayTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableModuleBayTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m PatchedWritableModuleBayTemplateRequest) GetDeviceType() int32 {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *PatchedWritableModuleBayTemplateRequest) SetDeviceType(val int32) {
	m.DeviceType = val
}

// GetLabel returns the Label property
func (m PatchedWritableModuleBayTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PatchedWritableModuleBayTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m PatchedWritableModuleBayTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableModuleBayTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetPosition returns the Position property
func (m PatchedWritableModuleBayTemplateRequest) GetPosition() string {
	return m.Position
}

// SetPosition sets the Position property
func (m *PatchedWritableModuleBayTemplateRequest) SetPosition(val string) {
	m.Position = val
}
