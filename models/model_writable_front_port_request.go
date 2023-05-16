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

// writableFrontPortRequestColorPattern is the validation pattern for WritableFrontPortRequest.Color
var writableFrontPortRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// WritableFrontPortRequest is an object. Adds support for custom fields and tags.
type WritableFrontPortRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
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
	// RearPort:
	RearPort int32 `json:"rear_port" mapstructure:"rear_port"`
	// RearPortPosition: Mapped position on corresponding rear port
	RearPortPosition int32 `json:"rear_port_position,omitempty" mapstructure:"rear_port_position,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type: * `8p8c` - 8P8C
	// * `8p6c` - 8P6C
	// * `8p4c` - 8P4C
	// * `8p2c` - 8P2C
	// * `6p6c` - 6P6C
	// * `6p4c` - 6P4C
	// * `6p2c` - 6P2C
	// * `4p4c` - 4P4C
	// * `4p2c` - 4P2C
	// * `gg45` - GG45
	// * `tera-4p` - TERA 4P
	// * `tera-2p` - TERA 2P
	// * `tera-1p` - TERA 1P
	// * `110-punch` - 110 Punch
	// * `bnc` - BNC
	// * `f` - F Connector
	// * `n` - N Connector
	// * `mrj21` - MRJ21
	// * `fc` - FC
	// * `lc` - LC
	// * `lc-pc` - LC/PC
	// * `lc-upc` - LC/UPC
	// * `lc-apc` - LC/APC
	// * `lsh` - LSH
	// * `lsh-pc` - LSH/PC
	// * `lsh-upc` - LSH/UPC
	// * `lsh-apc` - LSH/APC
	// * `mpo` - MPO
	// * `mtrj` - MTRJ
	// * `sc` - SC
	// * `sc-pc` - SC/PC
	// * `sc-upc` - SC/UPC
	// * `sc-apc` - SC/APC
	// * `st` - ST
	// * `cs` - CS
	// * `sn` - SN
	// * `sma-905` - SMA 905
	// * `sma-906` - SMA 906
	// * `urm-p2` - URM-P2
	// * `urm-p4` - URM-P4
	// * `urm-p8` - URM-P8
	// * `splice` - Splice
	// * `other` - Other
	Type string `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m WritableFrontPortRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(writableFrontPortRequestColorPattern),
		),
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
		"rearPortPosition": validation.Validate(
			m.RearPortPosition, validation.Min(int32(1)), validation.Max(int32(1024)),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m WritableFrontPortRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *WritableFrontPortRequest) SetColor(val string) {
	m.Color = val
}

// GetCustomFields returns the CustomFields property
func (m WritableFrontPortRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableFrontPortRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableFrontPortRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableFrontPortRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m WritableFrontPortRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *WritableFrontPortRequest) SetDevice(val int32) {
	m.Device = val
}

// GetLabel returns the Label property
func (m WritableFrontPortRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *WritableFrontPortRequest) SetLabel(val string) {
	m.Label = val
}

// GetMarkConnected returns the MarkConnected property
func (m WritableFrontPortRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *WritableFrontPortRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetModule returns the Module property
func (m WritableFrontPortRequest) GetModule() *int32 {
	return m.Module
}

// SetModule sets the Module property
func (m *WritableFrontPortRequest) SetModule(val *int32) {
	m.Module = val
}

// GetName returns the Name property
func (m WritableFrontPortRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableFrontPortRequest) SetName(val string) {
	m.Name = val
}

// GetRearPort returns the RearPort property
func (m WritableFrontPortRequest) GetRearPort() int32 {
	return m.RearPort
}

// SetRearPort sets the RearPort property
func (m *WritableFrontPortRequest) SetRearPort(val int32) {
	m.RearPort = val
}

// GetRearPortPosition returns the RearPortPosition property
func (m WritableFrontPortRequest) GetRearPortPosition() int32 {
	return m.RearPortPosition
}

// SetRearPortPosition sets the RearPortPosition property
func (m *WritableFrontPortRequest) SetRearPortPosition(val int32) {
	m.RearPortPosition = val
}

// GetTags returns the Tags property
func (m WritableFrontPortRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableFrontPortRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetType returns the Type property
func (m WritableFrontPortRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *WritableFrontPortRequest) SetType(val string) {
	m.Type = val
}
