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

// rearPortTemplateColorPattern is the validation pattern for RearPortTemplate.Color
var rearPortTemplateColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// RearPortTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type RearPortTemplate struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType *NestedDeviceType `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// ModuleType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleType *NestedModuleType `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Positions:
	Positions int32 `json:"positions,omitempty" mapstructure:"positions,omitempty"`
	// Type:
	Type struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m RearPortTemplate) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(rearPortTemplateColorPattern),
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
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"positions": validation.Validate(
			m.Positions, validation.Min(int32(1)), validation.Max(int32(1024)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m RearPortTemplate) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *RearPortTemplate) SetColor(val string) {
	m.Color = val
}

// GetCreated returns the Created property
func (m RearPortTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *RearPortTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m RearPortTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RearPortTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m RearPortTemplate) GetDeviceType() *NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *RearPortTemplate) SetDeviceType(val *NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m RearPortTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *RearPortTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m RearPortTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *RearPortTemplate) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m RearPortTemplate) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *RearPortTemplate) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m RearPortTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *RearPortTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetModuleType returns the ModuleType property
func (m RearPortTemplate) GetModuleType() *NestedModuleType {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *RearPortTemplate) SetModuleType(val *NestedModuleType) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m RearPortTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RearPortTemplate) SetName(val string) {
	m.Name = val
}

// GetPositions returns the Positions property
func (m RearPortTemplate) GetPositions() int32 {
	return m.Positions
}

// SetPositions sets the Positions property
func (m *RearPortTemplate) SetPositions(val int32) {
	m.Positions = val
}

// GetType returns the Type property
func (m RearPortTemplate) GetType() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *RearPortTemplate) SetType(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m RearPortTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *RearPortTemplate) SetUrl(val string) {
	m.Url = val
}
