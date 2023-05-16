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

// frontPortRequestColorPattern is the validation pattern for FrontPortRequest.Color
var frontPortRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// FrontPortRequest is an object. Adds support for custom fields and tags.
type FrontPortRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDeviceRequest `json:"device" mapstructure:"device"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// Module: Used by device component serializers.
	Module *ComponentNestedModuleRequest `json:"module,omitempty" mapstructure:"module,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// RearPort: NestedRearPortSerializer but with parent device omitted (since front and rear ports must belong to same device)
	RearPort FrontPortRearPortRequest `json:"rear_port" mapstructure:"rear_port"`
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
func (m FrontPortRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(frontPortRequestColorPattern),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"module": validation.Validate(
			m.Module,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"rearPort": validation.Validate(
			m.RearPort, validation.NotNil,
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
func (m FrontPortRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *FrontPortRequest) SetColor(val string) {
	m.Color = val
}

// GetCustomFields returns the CustomFields property
func (m FrontPortRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *FrontPortRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m FrontPortRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *FrontPortRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m FrontPortRequest) GetDevice() NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *FrontPortRequest) SetDevice(val NestedDeviceRequest) {
	m.Device = val
}

// GetLabel returns the Label property
func (m FrontPortRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *FrontPortRequest) SetLabel(val string) {
	m.Label = val
}

// GetMarkConnected returns the MarkConnected property
func (m FrontPortRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *FrontPortRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetModule returns the Module property
func (m FrontPortRequest) GetModule() *ComponentNestedModuleRequest {
	return m.Module
}

// SetModule sets the Module property
func (m *FrontPortRequest) SetModule(val *ComponentNestedModuleRequest) {
	m.Module = val
}

// GetName returns the Name property
func (m FrontPortRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *FrontPortRequest) SetName(val string) {
	m.Name = val
}

// GetRearPort returns the RearPort property
func (m FrontPortRequest) GetRearPort() FrontPortRearPortRequest {
	return m.RearPort
}

// SetRearPort sets the RearPort property
func (m *FrontPortRequest) SetRearPort(val FrontPortRearPortRequest) {
	m.RearPort = val
}

// GetRearPortPosition returns the RearPortPosition property
func (m FrontPortRequest) GetRearPortPosition() int32 {
	return m.RearPortPosition
}

// SetRearPortPosition sets the RearPortPosition property
func (m *FrontPortRequest) SetRearPortPosition(val int32) {
	m.RearPortPosition = val
}

// GetTags returns the Tags property
func (m FrontPortRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *FrontPortRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetType returns the Type property
func (m FrontPortRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *FrontPortRequest) SetType(val string) {
	m.Type = val
}
