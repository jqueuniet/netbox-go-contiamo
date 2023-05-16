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

// Module is an object. Adds support for custom fields and tags.
type Module struct {
	// AssetTag: A unique tag used to identify this device
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
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
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// ModuleBay: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleBay NestedModuleBay `json:"module_bay" mapstructure:"module_bay"`
	// ModuleType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleType NestedModuleType `json:"module_type" mapstructure:"module_type"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Status:
	Status *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Module) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag, validation.Length(0, 50),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"moduleBay": validation.Validate(
			m.ModuleBay, validation.NotNil,
		),
		"moduleType": validation.Validate(
			m.ModuleType, validation.NotNil,
		),
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m Module) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *Module) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComments returns the Comments property
func (m Module) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Module) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Module) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Module) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Module) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Module) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Module) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Module) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m Module) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *Module) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m Module) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Module) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Module) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Module) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Module) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Module) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetModuleBay returns the ModuleBay property
func (m Module) GetModuleBay() NestedModuleBay {
	return m.ModuleBay
}

// SetModuleBay sets the ModuleBay property
func (m *Module) SetModuleBay(val NestedModuleBay) {
	m.ModuleBay = val
}

// GetModuleType returns the ModuleType property
func (m Module) GetModuleType() NestedModuleType {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *Module) SetModuleType(val NestedModuleType) {
	m.ModuleType = val
}

// GetSerial returns the Serial property
func (m Module) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *Module) SetSerial(val string) {
	m.Serial = val
}

// GetStatus returns the Status property
func (m Module) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Module) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Module) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Module) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m Module) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Module) SetUrl(val string) {
	m.Url = val
}
