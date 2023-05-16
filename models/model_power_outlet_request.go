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

// PowerOutletRequest is an object. Adds support for custom fields and tags.
type PowerOutletRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDeviceRequest `json:"device" mapstructure:"device"`
	// FeedLeg: * `A` - A
	// * `B` - B
	// * `C` - C
	FeedLeg *string `json:"feed_leg,omitempty" mapstructure:"feed_leg,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// Module: Used by device component serializers.
	Module *ComponentNestedModuleRequest `json:"module,omitempty" mapstructure:"module,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PowerPort: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PowerPort *NestedPowerPortRequest `json:"power_port,omitempty" mapstructure:"power_port,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type: * `iec-60320-c5` - C5
	// * `iec-60320-c7` - C7
	// * `iec-60320-c13` - C13
	// * `iec-60320-c15` - C15
	// * `iec-60320-c19` - C19
	// * `iec-60320-c21` - C21
	// * `iec-60309-p-n-e-4h` - P+N+E 4H
	// * `iec-60309-p-n-e-6h` - P+N+E 6H
	// * `iec-60309-p-n-e-9h` - P+N+E 9H
	// * `iec-60309-2p-e-4h` - 2P+E 4H
	// * `iec-60309-2p-e-6h` - 2P+E 6H
	// * `iec-60309-2p-e-9h` - 2P+E 9H
	// * `iec-60309-3p-e-4h` - 3P+E 4H
	// * `iec-60309-3p-e-6h` - 3P+E 6H
	// * `iec-60309-3p-e-9h` - 3P+E 9H
	// * `iec-60309-3p-n-e-4h` - 3P+N+E 4H
	// * `iec-60309-3p-n-e-6h` - 3P+N+E 6H
	// * `iec-60309-3p-n-e-9h` - 3P+N+E 9H
	// * `nema-1-15r` - NEMA 1-15R
	// * `nema-5-15r` - NEMA 5-15R
	// * `nema-5-20r` - NEMA 5-20R
	// * `nema-5-30r` - NEMA 5-30R
	// * `nema-5-50r` - NEMA 5-50R
	// * `nema-6-15r` - NEMA 6-15R
	// * `nema-6-20r` - NEMA 6-20R
	// * `nema-6-30r` - NEMA 6-30R
	// * `nema-6-50r` - NEMA 6-50R
	// * `nema-10-30r` - NEMA 10-30R
	// * `nema-10-50r` - NEMA 10-50R
	// * `nema-14-20r` - NEMA 14-20R
	// * `nema-14-30r` - NEMA 14-30R
	// * `nema-14-50r` - NEMA 14-50R
	// * `nema-14-60r` - NEMA 14-60R
	// * `nema-15-15r` - NEMA 15-15R
	// * `nema-15-20r` - NEMA 15-20R
	// * `nema-15-30r` - NEMA 15-30R
	// * `nema-15-50r` - NEMA 15-50R
	// * `nema-15-60r` - NEMA 15-60R
	// * `nema-l1-15r` - NEMA L1-15R
	// * `nema-l5-15r` - NEMA L5-15R
	// * `nema-l5-20r` - NEMA L5-20R
	// * `nema-l5-30r` - NEMA L5-30R
	// * `nema-l5-50r` - NEMA L5-50R
	// * `nema-l6-15r` - NEMA L6-15R
	// * `nema-l6-20r` - NEMA L6-20R
	// * `nema-l6-30r` - NEMA L6-30R
	// * `nema-l6-50r` - NEMA L6-50R
	// * `nema-l10-30r` - NEMA L10-30R
	// * `nema-l14-20r` - NEMA L14-20R
	// * `nema-l14-30r` - NEMA L14-30R
	// * `nema-l14-50r` - NEMA L14-50R
	// * `nema-l14-60r` - NEMA L14-60R
	// * `nema-l15-20r` - NEMA L15-20R
	// * `nema-l15-30r` - NEMA L15-30R
	// * `nema-l15-50r` - NEMA L15-50R
	// * `nema-l15-60r` - NEMA L15-60R
	// * `nema-l21-20r` - NEMA L21-20R
	// * `nema-l21-30r` - NEMA L21-30R
	// * `nema-l22-30r` - NEMA L22-30R
	// * `CS6360C` - CS6360C
	// * `CS6364C` - CS6364C
	// * `CS8164C` - CS8164C
	// * `CS8264C` - CS8264C
	// * `CS8364C` - CS8364C
	// * `CS8464C` - CS8464C
	// * `ita-e` - ITA Type E (CEE 7/5)
	// * `ita-f` - ITA Type F (CEE 7/3)
	// * `ita-g` - ITA Type G (BS 1363)
	// * `ita-h` - ITA Type H
	// * `ita-i` - ITA Type I
	// * `ita-j` - ITA Type J
	// * `ita-k` - ITA Type K
	// * `ita-l` - ITA Type L (CEI 23-50)
	// * `ita-m` - ITA Type M (BS 546)
	// * `ita-n` - ITA Type N
	// * `ita-o` - ITA Type O
	// * `ita-multistandard` - ITA Multistandard
	// * `usb-a` - USB Type A
	// * `usb-micro-b` - USB Micro B
	// * `usb-c` - USB Type C
	// * `dc-terminal` - DC Terminal
	// * `hdot-cx` - HDOT Cx
	// * `saf-d-grid` - Saf-D-Grid
	// * `neutrik-powercon-20a` - Neutrik powerCON (20A)
	// * `neutrik-powercon-32a` - Neutrik powerCON (32A)
	// * `neutrik-powercon-true1` - Neutrik powerCON TRUE1
	// * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP
	// * `ubiquiti-smartpower` - Ubiquiti SmartPower
	// * `hardwired` - Hardwired
	// * `other` - Other
	Type *string `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m PowerOutletRequest) Validate() error {
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
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"module": validation.Validate(
			m.Module,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"powerPort": validation.Validate(
			m.PowerPort,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PowerOutletRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PowerOutletRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PowerOutletRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PowerOutletRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PowerOutletRequest) GetDevice() NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *PowerOutletRequest) SetDevice(val NestedDeviceRequest) {
	m.Device = val
}

// GetFeedLeg returns the FeedLeg property
func (m PowerOutletRequest) GetFeedLeg() *string {
	return m.FeedLeg
}

// SetFeedLeg sets the FeedLeg property
func (m *PowerOutletRequest) SetFeedLeg(val *string) {
	m.FeedLeg = val
}

// GetLabel returns the Label property
func (m PowerOutletRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PowerOutletRequest) SetLabel(val string) {
	m.Label = val
}

// GetMarkConnected returns the MarkConnected property
func (m PowerOutletRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *PowerOutletRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetModule returns the Module property
func (m PowerOutletRequest) GetModule() *ComponentNestedModuleRequest {
	return m.Module
}

// SetModule sets the Module property
func (m *PowerOutletRequest) SetModule(val *ComponentNestedModuleRequest) {
	m.Module = val
}

// GetName returns the Name property
func (m PowerOutletRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PowerOutletRequest) SetName(val string) {
	m.Name = val
}

// GetPowerPort returns the PowerPort property
func (m PowerOutletRequest) GetPowerPort() *NestedPowerPortRequest {
	return m.PowerPort
}

// SetPowerPort sets the PowerPort property
func (m *PowerOutletRequest) SetPowerPort(val *NestedPowerPortRequest) {
	m.PowerPort = val
}

// GetTags returns the Tags property
func (m PowerOutletRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PowerOutletRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetType returns the Type property
func (m PowerOutletRequest) GetType() *string {
	return m.Type
}

// SetType sets the Type property
func (m *PowerOutletRequest) SetType(val *string) {
	m.Type = val
}
