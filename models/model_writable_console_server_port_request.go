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

// WritableConsoleServerPortRequest is an object. Adds support for custom fields and tags.
type WritableConsoleServerPortRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device" mapstructure:"device"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// Module:
	Module *int32 `json:"module,omitempty" mapstructure:"module,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Speed: Port speed in bits per second
	//
	// * `1200` - 1200 bps
	// * `2400` - 2400 bps
	// * `4800` - 4800 bps
	// * `9600` - 9600 bps
	// * `19200` - 19.2 kbps
	// * `38400` - 38.4 kbps
	// * `57600` - 57.6 kbps
	// * `115200` - 115.2 kbps
	Speed *int32 `json:"speed,omitempty" mapstructure:"speed,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type: Physical port type
	//
	// * `de-9` - DE-9
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
func (m WritableConsoleServerPortRequest) Validate() error {
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
		"speed": validation.Validate(
			m.Speed, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m WritableConsoleServerPortRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableConsoleServerPortRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableConsoleServerPortRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableConsoleServerPortRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m WritableConsoleServerPortRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *WritableConsoleServerPortRequest) SetDevice(val int32) {
	m.Device = val
}

// GetLabel returns the Label property
func (m WritableConsoleServerPortRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *WritableConsoleServerPortRequest) SetLabel(val string) {
	m.Label = val
}

// GetMarkConnected returns the MarkConnected property
func (m WritableConsoleServerPortRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *WritableConsoleServerPortRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetModule returns the Module property
func (m WritableConsoleServerPortRequest) GetModule() *int32 {
	return m.Module
}

// SetModule sets the Module property
func (m *WritableConsoleServerPortRequest) SetModule(val *int32) {
	m.Module = val
}

// GetName returns the Name property
func (m WritableConsoleServerPortRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableConsoleServerPortRequest) SetName(val string) {
	m.Name = val
}

// GetSpeed returns the Speed property
func (m WritableConsoleServerPortRequest) GetSpeed() *int32 {
	return m.Speed
}

// SetSpeed sets the Speed property
func (m *WritableConsoleServerPortRequest) SetSpeed(val *int32) {
	m.Speed = val
}

// GetTags returns the Tags property
func (m WritableConsoleServerPortRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableConsoleServerPortRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetType returns the Type property
func (m WritableConsoleServerPortRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *WritableConsoleServerPortRequest) SetType(val string) {
	m.Type = val
}
