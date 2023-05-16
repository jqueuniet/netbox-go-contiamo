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

// WritableConsolePortTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableConsolePortTemplateRequest struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType:
	DeviceType *int32 `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// ModuleType:
	ModuleType *int32 `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Type: * `de-9` - DE-9
	// * `db-25` - DB-25
	// * `rj-11` - RJ-11
	// * `rj-12` - RJ-12
	// * `rj-45` - RJ-45
	// * `mini-din-8` - Mini-DIN 8
	// * `usb-a` - USB Type A
	// * `usb-b` - USB Type B
	// * `usb-c` - USB Type C
	// * `usb-mini-a` - USB Mini A
	// * `usb-mini-b` - USB Mini B
	// * `usb-micro-a` - USB Micro A
	// * `usb-micro-b` - USB Micro B
	// * `usb-micro-ab` - USB Micro AB
	// * `other` - Other
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableConsolePortTemplateRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
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
func (m WritableConsolePortTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableConsolePortTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m WritableConsolePortTemplateRequest) GetDeviceType() *int32 {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *WritableConsolePortTemplateRequest) SetDeviceType(val *int32) {
	m.DeviceType = val
}

// GetLabel returns the Label property
func (m WritableConsolePortTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *WritableConsolePortTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetModuleType returns the ModuleType property
func (m WritableConsolePortTemplateRequest) GetModuleType() *int32 {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *WritableConsolePortTemplateRequest) SetModuleType(val *int32) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m WritableConsolePortTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableConsolePortTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetType returns the Type property
func (m WritableConsolePortTemplateRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *WritableConsolePortTemplateRequest) SetType(val string) {
	m.Type = val
}
