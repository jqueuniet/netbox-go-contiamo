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

// frontPortTemplateColorPattern is the validation pattern for FrontPortTemplate.Color
var frontPortTemplateColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// FrontPortTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type FrontPortTemplate struct {
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
	// RearPort: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	RearPort NestedRearPortTemplate `json:"rear_port" mapstructure:"rear_port"`
	// RearPortPosition:
	RearPortPosition int32 `json:"rear_port_position,omitempty" mapstructure:"rear_port_position,omitempty"`
	// Type:
	Type struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m FrontPortTemplate) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(frontPortTemplateColorPattern),
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
		"rearPort": validation.Validate(
			m.RearPort, validation.NotNil,
		),
		"rearPortPosition": validation.Validate(
			m.RearPortPosition, validation.Min(int32(1)), validation.Max(int32(1024)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m FrontPortTemplate) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *FrontPortTemplate) SetColor(val string) {
	m.Color = val
}

// GetCreated returns the Created property
func (m FrontPortTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *FrontPortTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m FrontPortTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *FrontPortTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m FrontPortTemplate) GetDeviceType() *NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *FrontPortTemplate) SetDeviceType(val *NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m FrontPortTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *FrontPortTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m FrontPortTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *FrontPortTemplate) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m FrontPortTemplate) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *FrontPortTemplate) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m FrontPortTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *FrontPortTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetModuleType returns the ModuleType property
func (m FrontPortTemplate) GetModuleType() *NestedModuleType {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *FrontPortTemplate) SetModuleType(val *NestedModuleType) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m FrontPortTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *FrontPortTemplate) SetName(val string) {
	m.Name = val
}

// GetRearPort returns the RearPort property
func (m FrontPortTemplate) GetRearPort() NestedRearPortTemplate {
	return m.RearPort
}

// SetRearPort sets the RearPort property
func (m *FrontPortTemplate) SetRearPort(val NestedRearPortTemplate) {
	m.RearPort = val
}

// GetRearPortPosition returns the RearPortPosition property
func (m FrontPortTemplate) GetRearPortPosition() int32 {
	return m.RearPortPosition
}

// SetRearPortPosition sets the RearPortPosition property
func (m *FrontPortTemplate) SetRearPortPosition(val int32) {
	m.RearPortPosition = val
}

// GetType returns the Type property
func (m FrontPortTemplate) GetType() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *FrontPortTemplate) SetType(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m FrontPortTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *FrontPortTemplate) SetUrl(val string) {
	m.Url = val
}
