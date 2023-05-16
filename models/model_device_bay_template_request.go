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

// DeviceBayTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type DeviceBayTemplateRequest struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType NestedDeviceTypeRequest `json:"device_type" mapstructure:"device_type"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
}

// Validate implements basic validation for this model
func (m DeviceBayTemplateRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"deviceType": validation.Validate(
			m.DeviceType, validation.NotNil,
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m DeviceBayTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceBayTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m DeviceBayTemplateRequest) GetDeviceType() NestedDeviceTypeRequest {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *DeviceBayTemplateRequest) SetDeviceType(val NestedDeviceTypeRequest) {
	m.DeviceType = val
}

// GetLabel returns the Label property
func (m DeviceBayTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DeviceBayTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m DeviceBayTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *DeviceBayTemplateRequest) SetName(val string) {
	m.Name = val
}
