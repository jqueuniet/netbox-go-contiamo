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

// DeviceBayRequest is an object. Adds support for custom fields and tags.
type DeviceBayRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDeviceRequest `json:"device" mapstructure:"device"`
	// InstalledDevice: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	InstalledDevice *NestedDeviceRequest `json:"installed_device,omitempty" mapstructure:"installed_device,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m DeviceBayRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"installedDevice": validation.Validate(
			m.InstalledDevice,
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
func (m DeviceBayRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *DeviceBayRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m DeviceBayRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceBayRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m DeviceBayRequest) GetDevice() NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *DeviceBayRequest) SetDevice(val NestedDeviceRequest) {
	m.Device = val
}

// GetInstalledDevice returns the InstalledDevice property
func (m DeviceBayRequest) GetInstalledDevice() *NestedDeviceRequest {
	return m.InstalledDevice
}

// SetInstalledDevice sets the InstalledDevice property
func (m *DeviceBayRequest) SetInstalledDevice(val *NestedDeviceRequest) {
	m.InstalledDevice = val
}

// GetLabel returns the Label property
func (m DeviceBayRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DeviceBayRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m DeviceBayRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *DeviceBayRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m DeviceBayRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *DeviceBayRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
