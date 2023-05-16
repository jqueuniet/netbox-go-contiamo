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

// WritableDeviceBayRequest is an object. Adds support for custom fields and tags.
type WritableDeviceBayRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device" mapstructure:"device"`
	// InstalledDevice:
	InstalledDevice *int32 `json:"installed_device,omitempty" mapstructure:"installed_device,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableDeviceBayRequest) Validate() error {
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
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m WritableDeviceBayRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableDeviceBayRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableDeviceBayRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableDeviceBayRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m WritableDeviceBayRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *WritableDeviceBayRequest) SetDevice(val int32) {
	m.Device = val
}

// GetInstalledDevice returns the InstalledDevice property
func (m WritableDeviceBayRequest) GetInstalledDevice() *int32 {
	return m.InstalledDevice
}

// SetInstalledDevice sets the InstalledDevice property
func (m *WritableDeviceBayRequest) SetInstalledDevice(val *int32) {
	m.InstalledDevice = val
}

// GetLabel returns the Label property
func (m WritableDeviceBayRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *WritableDeviceBayRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m WritableDeviceBayRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableDeviceBayRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m WritableDeviceBayRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableDeviceBayRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
