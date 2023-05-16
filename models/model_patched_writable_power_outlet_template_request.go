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

// PatchedWritablePowerOutletTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritablePowerOutletTemplateRequest struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType:
	DeviceType *int32 `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// FeedLeg: Phase (for three-phase feeds)
	//
	// * `A` - A
	// * `B` - B
	// * `C` - C
	FeedLeg string `json:"feed_leg,omitempty" mapstructure:"feed_leg,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// ModuleType:
	ModuleType *int32 `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// PowerPort:
	PowerPort *int32 `json:"power_port,omitempty" mapstructure:"power_port,omitempty"`
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
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritablePowerOutletTemplateRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 64),
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m PatchedWritablePowerOutletTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritablePowerOutletTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m PatchedWritablePowerOutletTemplateRequest) GetDeviceType() *int32 {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *PatchedWritablePowerOutletTemplateRequest) SetDeviceType(val *int32) {
	m.DeviceType = val
}

// GetFeedLeg returns the FeedLeg property
func (m PatchedWritablePowerOutletTemplateRequest) GetFeedLeg() string {
	return m.FeedLeg
}

// SetFeedLeg sets the FeedLeg property
func (m *PatchedWritablePowerOutletTemplateRequest) SetFeedLeg(val string) {
	m.FeedLeg = val
}

// GetLabel returns the Label property
func (m PatchedWritablePowerOutletTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PatchedWritablePowerOutletTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetModuleType returns the ModuleType property
func (m PatchedWritablePowerOutletTemplateRequest) GetModuleType() *int32 {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *PatchedWritablePowerOutletTemplateRequest) SetModuleType(val *int32) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m PatchedWritablePowerOutletTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritablePowerOutletTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetPowerPort returns the PowerPort property
func (m PatchedWritablePowerOutletTemplateRequest) GetPowerPort() *int32 {
	return m.PowerPort
}

// SetPowerPort sets the PowerPort property
func (m *PatchedWritablePowerOutletTemplateRequest) SetPowerPort(val *int32) {
	m.PowerPort = val
}

// GetType returns the Type property
func (m PatchedWritablePowerOutletTemplateRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritablePowerOutletTemplateRequest) SetType(val string) {
	m.Type = val
}
