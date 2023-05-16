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

// PowerOutletTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PowerOutletTemplate struct {
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
	// FeedLeg:
	FeedLeg *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"feed_leg,omitempty" mapstructure:"feed_leg,omitempty"`
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
	// PowerPort: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PowerPort *NestedPowerPortTemplate `json:"power_port,omitempty" mapstructure:"power_port,omitempty"`
	// Type:
	Type *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type,omitempty" mapstructure:"type,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m PowerOutletTemplate) Validate() error {
	return validation.Errors{
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
		"powerPort": validation.Validate(
			m.PowerPort,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m PowerOutletTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *PowerOutletTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m PowerOutletTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PowerOutletTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m PowerOutletTemplate) GetDeviceType() *NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *PowerOutletTemplate) SetDeviceType(val *NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m PowerOutletTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *PowerOutletTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetFeedLeg returns the FeedLeg property
func (m PowerOutletTemplate) GetFeedLeg() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.FeedLeg
}

// SetFeedLeg sets the FeedLeg property
func (m *PowerOutletTemplate) SetFeedLeg(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.FeedLeg = val
}

// GetId returns the Id property
func (m PowerOutletTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *PowerOutletTemplate) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m PowerOutletTemplate) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PowerOutletTemplate) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m PowerOutletTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *PowerOutletTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetModuleType returns the ModuleType property
func (m PowerOutletTemplate) GetModuleType() *NestedModuleType {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *PowerOutletTemplate) SetModuleType(val *NestedModuleType) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m PowerOutletTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PowerOutletTemplate) SetName(val string) {
	m.Name = val
}

// GetPowerPort returns the PowerPort property
func (m PowerOutletTemplate) GetPowerPort() *NestedPowerPortTemplate {
	return m.PowerPort
}

// SetPowerPort sets the PowerPort property
func (m *PowerOutletTemplate) SetPowerPort(val *NestedPowerPortTemplate) {
	m.PowerPort = val
}

// GetType returns the Type property
func (m PowerOutletTemplate) GetType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *PowerOutletTemplate) SetType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m PowerOutletTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *PowerOutletTemplate) SetUrl(val string) {
	m.Url = val
}
