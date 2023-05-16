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

// ModuleBay is an object. Adds support for custom fields and tags.
type ModuleBay struct {
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDevice `json:"device" mapstructure:"device"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// InstalledModule: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	InstalledModule *ModuleBayNestedModule `json:"installed_module,omitempty" mapstructure:"installed_module,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Position: Identifier to reference when renaming installed components
	Position string `json:"position,omitempty" mapstructure:"position,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ModuleBay) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"installedModule": validation.Validate(
			m.InstalledModule,
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
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m ModuleBay) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ModuleBay) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ModuleBay) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ModuleBay) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ModuleBay) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ModuleBay) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m ModuleBay) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *ModuleBay) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m ModuleBay) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ModuleBay) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ModuleBay) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ModuleBay) SetId(val int32) {
	m.Id = val
}

// GetInstalledModule returns the InstalledModule property
func (m ModuleBay) GetInstalledModule() *ModuleBayNestedModule {
	return m.InstalledModule
}

// SetInstalledModule sets the InstalledModule property
func (m *ModuleBay) SetInstalledModule(val *ModuleBayNestedModule) {
	m.InstalledModule = val
}

// GetLabel returns the Label property
func (m ModuleBay) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *ModuleBay) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m ModuleBay) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ModuleBay) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ModuleBay) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ModuleBay) SetName(val string) {
	m.Name = val
}

// GetPosition returns the Position property
func (m ModuleBay) GetPosition() string {
	return m.Position
}

// SetPosition sets the Position property
func (m *ModuleBay) SetPosition(val string) {
	m.Position = val
}

// GetTags returns the Tags property
func (m ModuleBay) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ModuleBay) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m ModuleBay) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ModuleBay) SetUrl(val string) {
	m.Url = val
}
