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

// ModuleBayRequest is an object. Adds support for custom fields and tags.
type ModuleBayRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDeviceRequest `json:"device" mapstructure:"device"`
	// InstalledModule: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	InstalledModule *ModuleBayNestedModuleRequest `json:"installed_module,omitempty" mapstructure:"installed_module,omitempty"`
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
func (m ModuleBayRequest) Validate() error {
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
		"installedModule": validation.Validate(
			m.InstalledModule,
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
func (m ModuleBayRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ModuleBayRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ModuleBayRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ModuleBayRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m ModuleBayRequest) GetDevice() NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *ModuleBayRequest) SetDevice(val NestedDeviceRequest) {
	m.Device = val
}

// GetInstalledModule returns the InstalledModule property
func (m ModuleBayRequest) GetInstalledModule() *ModuleBayNestedModuleRequest {
	return m.InstalledModule
}

// SetInstalledModule sets the InstalledModule property
func (m *ModuleBayRequest) SetInstalledModule(val *ModuleBayNestedModuleRequest) {
	m.InstalledModule = val
}

// GetLabel returns the Label property
func (m ModuleBayRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *ModuleBayRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m ModuleBayRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ModuleBayRequest) SetName(val string) {
	m.Name = val
}

// GetPosition returns the Position property
func (m ModuleBayRequest) GetPosition() string {
	return m.Position
}

// SetPosition sets the Position property
func (m *ModuleBayRequest) SetPosition(val string) {
	m.Position = val
}

// GetTags returns the Tags property
func (m ModuleBayRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ModuleBayRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
