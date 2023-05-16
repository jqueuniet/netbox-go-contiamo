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

// ModuleBayTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ModuleBayTemplate struct {
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType NestedDeviceType `json:"device_type" mapstructure:"device_type"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Position: Identifier to reference when renaming installed components
	Position string `json:"position,omitempty" mapstructure:"position,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ModuleBayTemplate) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"deviceType": validation.Validate(
			m.DeviceType, validation.NotNil,
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"position": validation.Validate(
			m.Position, validation.Length(0, 30),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m ModuleBayTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ModuleBayTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m ModuleBayTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ModuleBayTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m ModuleBayTemplate) GetDeviceType() NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *ModuleBayTemplate) SetDeviceType(val NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m ModuleBayTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ModuleBayTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ModuleBayTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ModuleBayTemplate) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m ModuleBayTemplate) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *ModuleBayTemplate) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m ModuleBayTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ModuleBayTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ModuleBayTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ModuleBayTemplate) SetName(val string) {
	m.Name = val
}

// GetPosition returns the Position property
func (m ModuleBayTemplate) GetPosition() string {
	return m.Position
}

// SetPosition sets the Position property
func (m *ModuleBayTemplate) SetPosition(val string) {
	m.Position = val
}

// GetUrl returns the Url property
func (m ModuleBayTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ModuleBayTemplate) SetUrl(val string) {
	m.Url = val
}
