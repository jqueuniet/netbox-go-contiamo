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

// rearPortTemplateRequestColorPattern is the validation pattern for RearPortTemplateRequest.Color
var rearPortTemplateRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// RearPortTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type RearPortTemplateRequest struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType *NestedDeviceTypeRequest `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// ModuleType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleType *NestedModuleTypeRequest `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Positions:
	Positions int32 `json:"positions,omitempty" mapstructure:"positions,omitempty"`
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
func (m RearPortTemplateRequest) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(rearPortTemplateRequestColorPattern),
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"deviceType": validation.Validate(
			m.DeviceType,
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"moduleType": validation.Validate(
			m.ModuleType,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"positions": validation.Validate(
			m.Positions, validation.Min(int32(1)), validation.Max(int32(1024)),
		),
	}.Filter()
}

// GetColor returns the Color property
func (m RearPortTemplateRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *RearPortTemplateRequest) SetColor(val string) {
	m.Color = val
}

// GetDescription returns the Description property
func (m RearPortTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RearPortTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m RearPortTemplateRequest) GetDeviceType() *NestedDeviceTypeRequest {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *RearPortTemplateRequest) SetDeviceType(val *NestedDeviceTypeRequest) {
	m.DeviceType = val
}

// GetLabel returns the Label property
func (m RearPortTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *RearPortTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetModuleType returns the ModuleType property
func (m RearPortTemplateRequest) GetModuleType() *NestedModuleTypeRequest {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *RearPortTemplateRequest) SetModuleType(val *NestedModuleTypeRequest) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m RearPortTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RearPortTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetPositions returns the Positions property
func (m RearPortTemplateRequest) GetPositions() int32 {
	return m.Positions
}

// SetPositions sets the Positions property
func (m *RearPortTemplateRequest) SetPositions(val int32) {
	m.Positions = val
}

// GetType returns the Type property
func (m RearPortTemplateRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *RearPortTemplateRequest) SetType(val string) {
	m.Type = val
}
