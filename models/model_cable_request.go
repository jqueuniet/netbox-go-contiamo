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

// cableRequestColorPattern is the validation pattern for CableRequest.Color
var cableRequestColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// CableRequest is an object. Adds support for custom fields and tags.
type CableRequest struct {
	// ATerminations:
	ATerminations []GenericObjectRequest `json:"a_terminations,omitempty" mapstructure:"a_terminations,omitempty"`
	// BTerminations:
	BTerminations []GenericObjectRequest `json:"b_terminations,omitempty" mapstructure:"b_terminations,omitempty"`
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Label:
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Length:
	Length *float64 `json:"length,omitempty" mapstructure:"length,omitempty"`
	// LengthUnit: * `km` - Kilometers
	// * `m` - Meters
	// * `cm` - Centimeters
	// * `mi` - Miles
	// * `ft` - Feet
	// * `in` - Inches
	LengthUnit *string `json:"length_unit,omitempty" mapstructure:"length_unit,omitempty"`
	// Status: * `connected` - Connected
	// * `planned` - Planned
	// * `decommissioning` - Decommissioning
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type: * `cat3` - CAT3
	// * `cat5` - CAT5
	// * `cat5e` - CAT5e
	// * `cat6` - CAT6
	// * `cat6a` - CAT6a
	// * `cat7` - CAT7
	// * `cat7a` - CAT7a
	// * `cat8` - CAT8
	// * `dac-active` - Direct Attach Copper (Active)
	// * `dac-passive` - Direct Attach Copper (Passive)
	// * `mrj21-trunk` - MRJ21 Trunk
	// * `coaxial` - Coaxial
	// * `mmf` - Multimode Fiber
	// * `mmf-om1` - Multimode Fiber (OM1)
	// * `mmf-om2` - Multimode Fiber (OM2)
	// * `mmf-om3` - Multimode Fiber (OM3)
	// * `mmf-om4` - Multimode Fiber (OM4)
	// * `mmf-om5` - Multimode Fiber (OM5)
	// * `smf` - Singlemode Fiber
	// * `smf-os1` - Singlemode Fiber (OS1)
	// * `smf-os2` - Singlemode Fiber (OS2)
	// * `aoc` - Active Optical Cabling (AOC)
	// * `power` - Power
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m CableRequest) Validate() error {
	return validation.Errors{
		"aTerminations": validation.Validate(
			m.ATerminations,
		),
		"bTerminations": validation.Validate(
			m.BTerminations,
		),
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(cableRequestColorPattern),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 100),
		),
		"length": validation.Validate(
			m.Length, validation.Min(*float64(-1e+06)), validation.Max(*float64(1e+06)),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
	}.Filter()
}

// GetATerminations returns the ATerminations property
func (m CableRequest) GetATerminations() []GenericObjectRequest {
	return m.ATerminations
}

// SetATerminations sets the ATerminations property
func (m *CableRequest) SetATerminations(val []GenericObjectRequest) {
	m.ATerminations = val
}

// GetBTerminations returns the BTerminations property
func (m CableRequest) GetBTerminations() []GenericObjectRequest {
	return m.BTerminations
}

// SetBTerminations sets the BTerminations property
func (m *CableRequest) SetBTerminations(val []GenericObjectRequest) {
	m.BTerminations = val
}

// GetColor returns the Color property
func (m CableRequest) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *CableRequest) SetColor(val string) {
	m.Color = val
}

// GetComments returns the Comments property
func (m CableRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *CableRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m CableRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *CableRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m CableRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CableRequest) SetDescription(val string) {
	m.Description = val
}

// GetLabel returns the Label property
func (m CableRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *CableRequest) SetLabel(val string) {
	m.Label = val
}

// GetLength returns the Length property
func (m CableRequest) GetLength() *float64 {
	return m.Length
}

// SetLength sets the Length property
func (m *CableRequest) SetLength(val *float64) {
	m.Length = val
}

// GetLengthUnit returns the LengthUnit property
func (m CableRequest) GetLengthUnit() *string {
	return m.LengthUnit
}

// SetLengthUnit sets the LengthUnit property
func (m *CableRequest) SetLengthUnit(val *string) {
	m.LengthUnit = val
}

// GetStatus returns the Status property
func (m CableRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *CableRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m CableRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *CableRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m CableRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *CableRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetType returns the Type property
func (m CableRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *CableRequest) SetType(val string) {
	m.Type = val
}
