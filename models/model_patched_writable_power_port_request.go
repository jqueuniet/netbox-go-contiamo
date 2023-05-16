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

// PatchedWritablePowerPortRequest is an object. Adds support for custom fields and tags.
type PatchedWritablePowerPortRequest struct {
	// AllocatedDraw: Allocated power draw (watts)
	AllocatedDraw *int32 `json:"allocated_draw,omitempty" mapstructure:"allocated_draw,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// MaximumDraw: Maximum power draw (watts)
	MaximumDraw *int32 `json:"maximum_draw,omitempty" mapstructure:"maximum_draw,omitempty"`
	// Module:
	Module *int32 `json:"module,omitempty" mapstructure:"module,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type: Physical port type
	//
	// * `iec-60320-c6` - C6
	// * `iec-60320-c8` - C8
	// * `iec-60320-c14` - C14
	// * `iec-60320-c16` - C16
	// * `iec-60320-c20` - C20
	// * `iec-60320-c22` - C22
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
	// * `nema-1-15p` - NEMA 1-15P
	// * `nema-5-15p` - NEMA 5-15P
	// * `nema-5-20p` - NEMA 5-20P
	// * `nema-5-30p` - NEMA 5-30P
	// * `nema-5-50p` - NEMA 5-50P
	// * `nema-6-15p` - NEMA 6-15P
	// * `nema-6-20p` - NEMA 6-20P
	// * `nema-6-30p` - NEMA 6-30P
	// * `nema-6-50p` - NEMA 6-50P
	// * `nema-10-30p` - NEMA 10-30P
	// * `nema-10-50p` - NEMA 10-50P
	// * `nema-14-20p` - NEMA 14-20P
	// * `nema-14-30p` - NEMA 14-30P
	// * `nema-14-50p` - NEMA 14-50P
	// * `nema-14-60p` - NEMA 14-60P
	// * `nema-15-15p` - NEMA 15-15P
	// * `nema-15-20p` - NEMA 15-20P
	// * `nema-15-30p` - NEMA 15-30P
	// * `nema-15-50p` - NEMA 15-50P
	// * `nema-15-60p` - NEMA 15-60P
	// * `nema-l1-15p` - NEMA L1-15P
	// * `nema-l5-15p` - NEMA L5-15P
	// * `nema-l5-20p` - NEMA L5-20P
	// * `nema-l5-30p` - NEMA L5-30P
	// * `nema-l5-50p` - NEMA L5-50P
	// * `nema-l6-15p` - NEMA L6-15P
	// * `nema-l6-20p` - NEMA L6-20P
	// * `nema-l6-30p` - NEMA L6-30P
	// * `nema-l6-50p` - NEMA L6-50P
	// * `nema-l10-30p` - NEMA L10-30P
	// * `nema-l14-20p` - NEMA L14-20P
	// * `nema-l14-30p` - NEMA L14-30P
	// * `nema-l14-50p` - NEMA L14-50P
	// * `nema-l14-60p` - NEMA L14-60P
	// * `nema-l15-20p` - NEMA L15-20P
	// * `nema-l15-30p` - NEMA L15-30P
	// * `nema-l15-50p` - NEMA L15-50P
	// * `nema-l15-60p` - NEMA L15-60P
	// * `nema-l21-20p` - NEMA L21-20P
	// * `nema-l21-30p` - NEMA L21-30P
	// * `nema-l22-30p` - NEMA L22-30P
	// * `cs6361c` - CS6361C
	// * `cs6365c` - CS6365C
	// * `cs8165c` - CS8165C
	// * `cs8265c` - CS8265C
	// * `cs8365c` - CS8365C
	// * `cs8465c` - CS8465C
	// * `ita-c` - ITA Type C (CEE 7/16)
	// * `ita-e` - ITA Type E (CEE 7/6)
	// * `ita-f` - ITA Type F (CEE 7/4)
	// * `ita-ef` - ITA Type E/F (CEE 7/7)
	// * `ita-g` - ITA Type G (BS 1363)
	// * `ita-h` - ITA Type H
	// * `ita-i` - ITA Type I
	// * `ita-j` - ITA Type J
	// * `ita-k` - ITA Type K
	// * `ita-l` - ITA Type L (CEI 23-50)
	// * `ita-m` - ITA Type M (BS 546)
	// * `ita-n` - ITA Type N
	// * `ita-o` - ITA Type O
	// * `usb-a` - USB Type A
	// * `usb-b` - USB Type B
	// * `usb-c` - USB Type C
	// * `usb-mini-a` - USB Mini A
	// * `usb-mini-b` - USB Mini B
	// * `usb-micro-a` - USB Micro A
	// * `usb-micro-b` - USB Micro B
	// * `usb-micro-ab` - USB Micro AB
	// * `usb-3-b` - USB 3.0 Type B
	// * `usb-3-micro-b` - USB 3.0 Micro B
	// * `dc-terminal` - DC Terminal
	// * `saf-d-grid` - Saf-D-Grid
	// * `neutrik-powercon-20` - Neutrik powerCON (20A)
	// * `neutrik-powercon-32` - Neutrik powerCON (32A)
	// * `neutrik-powercon-true1` - Neutrik powerCON TRUE1
	// * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP
	// * `ubiquiti-smartpower` - Ubiquiti SmartPower
	// * `hardwired` - Hardwired
	// * `other` - Other
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritablePowerPortRequest) Validate() error {
	return validation.Errors{
		"allocatedDraw": validation.Validate(
			m.AllocatedDraw, validation.Min(*int32(1)), validation.Max(*int32(32767)),
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
		"maximumDraw": validation.Validate(
			m.MaximumDraw, validation.Min(*int32(1)), validation.Max(*int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAllocatedDraw returns the AllocatedDraw property
func (m PatchedWritablePowerPortRequest) GetAllocatedDraw() *int32 {
	return m.AllocatedDraw
}

// SetAllocatedDraw sets the AllocatedDraw property
func (m *PatchedWritablePowerPortRequest) SetAllocatedDraw(val *int32) {
	m.AllocatedDraw = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritablePowerPortRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritablePowerPortRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritablePowerPortRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritablePowerPortRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PatchedWritablePowerPortRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *PatchedWritablePowerPortRequest) SetDevice(val int32) {
	m.Device = val
}

// GetLabel returns the Label property
func (m PatchedWritablePowerPortRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PatchedWritablePowerPortRequest) SetLabel(val string) {
	m.Label = val
}

// GetMarkConnected returns the MarkConnected property
func (m PatchedWritablePowerPortRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *PatchedWritablePowerPortRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetMaximumDraw returns the MaximumDraw property
func (m PatchedWritablePowerPortRequest) GetMaximumDraw() *int32 {
	return m.MaximumDraw
}

// SetMaximumDraw sets the MaximumDraw property
func (m *PatchedWritablePowerPortRequest) SetMaximumDraw(val *int32) {
	m.MaximumDraw = val
}

// GetModule returns the Module property
func (m PatchedWritablePowerPortRequest) GetModule() *int32 {
	return m.Module
}

// SetModule sets the Module property
func (m *PatchedWritablePowerPortRequest) SetModule(val *int32) {
	m.Module = val
}

// GetName returns the Name property
func (m PatchedWritablePowerPortRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritablePowerPortRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m PatchedWritablePowerPortRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritablePowerPortRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetType returns the Type property
func (m PatchedWritablePowerPortRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritablePowerPortRequest) SetType(val string) {
	m.Type = val
}
