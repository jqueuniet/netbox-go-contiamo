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

	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// cableColorPattern is the validation pattern for Cable.Color
var cableColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// Cable is an object. Adds support for custom fields and tags.
type Cable struct {
	// ATerminations:
	ATerminations []GenericObject `json:"a_terminations,omitempty" mapstructure:"a_terminations,omitempty"`
	// BTerminations:
	BTerminations []GenericObject `json:"b_terminations,omitempty" mapstructure:"b_terminations,omitempty"`
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label:
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Length:
	Length *float64 `json:"length,omitempty" mapstructure:"length,omitempty"`
	// LengthUnit:
	LengthUnit *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"length_unit,omitempty" mapstructure:"length_unit,omitempty"`
	// Status:
	Status *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
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
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Cable) Validate() error {
	return validation.Errors{
		"aTerminations": validation.Validate(
			m.ATerminations,
		),
		"bTerminations": validation.Validate(
			m.BTerminations,
		),
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(cableColorPattern),
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetATerminations returns the ATerminations property
func (m Cable) GetATerminations() []GenericObject {
	return m.ATerminations
}

// SetATerminations sets the ATerminations property
func (m *Cable) SetATerminations(val []GenericObject) {
	m.ATerminations = val
}

// GetBTerminations returns the BTerminations property
func (m Cable) GetBTerminations() []GenericObject {
	return m.BTerminations
}

// SetBTerminations sets the BTerminations property
func (m *Cable) SetBTerminations(val []GenericObject) {
	m.BTerminations = val
}

// GetColor returns the Color property
func (m Cable) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *Cable) SetColor(val string) {
	m.Color = val
}

// GetComments returns the Comments property
func (m Cable) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Cable) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Cable) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Cable) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Cable) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Cable) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Cable) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Cable) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Cable) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Cable) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Cable) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Cable) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m Cable) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *Cable) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m Cable) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Cable) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLength returns the Length property
func (m Cable) GetLength() *float64 {
	return m.Length
}

// SetLength sets the Length property
func (m *Cable) SetLength(val *float64) {
	m.Length = val
}

// GetLengthUnit returns the LengthUnit property
func (m Cable) GetLengthUnit() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.LengthUnit
}

// SetLengthUnit sets the LengthUnit property
func (m *Cable) SetLengthUnit(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.LengthUnit = val
}

// GetStatus returns the Status property
func (m Cable) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Cable) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Cable) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Cable) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Cable) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Cable) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetType returns the Type property
func (m Cable) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *Cable) SetType(val string) {
	m.Type = val
}

// GetUrl returns the Url property
func (m Cable) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Cable) SetUrl(val string) {
	m.Url = val
}
