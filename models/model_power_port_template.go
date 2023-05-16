// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// PowerPortTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PowerPortTemplate struct {
	// AllocatedDraw: Allocated power draw (watts)
	AllocatedDraw *int32 `json:"allocated_draw,omitempty" mapstructure:"allocated_draw,omitempty"`
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
	// MaximumDraw: Maximum power draw (watts)
	MaximumDraw *int32 `json:"maximum_draw,omitempty" mapstructure:"maximum_draw,omitempty"`
	// ModuleType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleType *NestedModuleType `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Type:
	Type *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type,omitempty" mapstructure:"type,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m PowerPortTemplate) Validate() error {
	return validation.Errors{
		"allocatedDraw": validation.Validate(
			m.AllocatedDraw, validation.Min(*int32(1)), validation.Max(*int32(32767)),
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
		"maximumDraw": validation.Validate(
			m.MaximumDraw, validation.Min(*int32(1)), validation.Max(*int32(32767)),
		),
		"moduleType": validation.Validate(
			m.ModuleType,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAllocatedDraw returns the AllocatedDraw property
func (m PowerPortTemplate) GetAllocatedDraw() *int32 {
	return m.AllocatedDraw
}

// SetAllocatedDraw sets the AllocatedDraw property
func (m *PowerPortTemplate) SetAllocatedDraw(val *int32) {
	m.AllocatedDraw = val
}

// GetCreated returns the Created property
func (m PowerPortTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *PowerPortTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m PowerPortTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PowerPortTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m PowerPortTemplate) GetDeviceType() *NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *PowerPortTemplate) SetDeviceType(val *NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m PowerPortTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *PowerPortTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m PowerPortTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *PowerPortTemplate) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m PowerPortTemplate) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PowerPortTemplate) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m PowerPortTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *PowerPortTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMaximumDraw returns the MaximumDraw property
func (m PowerPortTemplate) GetMaximumDraw() *int32 {
	return m.MaximumDraw
}

// SetMaximumDraw sets the MaximumDraw property
func (m *PowerPortTemplate) SetMaximumDraw(val *int32) {
	m.MaximumDraw = val
}

// GetModuleType returns the ModuleType property
func (m PowerPortTemplate) GetModuleType() *NestedModuleType {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *PowerPortTemplate) SetModuleType(val *NestedModuleType) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m PowerPortTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PowerPortTemplate) SetName(val string) {
	m.Name = val
}

// GetType returns the Type property
func (m PowerPortTemplate) GetType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *PowerPortTemplate) SetType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m PowerPortTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *PowerPortTemplate) SetUrl(val string) {
	m.Url = val
}
