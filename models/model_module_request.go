// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ModuleRequest is an object. Adds support for custom fields and tags.
type ModuleRequest struct {
	// AssetTag: A unique tag used to identify this device
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDeviceRequest `json:"device" mapstructure:"device"`
	// ModuleBay: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleBay NestedModuleBayRequest `json:"module_bay" mapstructure:"module_bay"`
	// ModuleType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleType NestedModuleTypeRequest `json:"module_type" mapstructure:"module_type"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Status: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
	// * `decommissioning` - Decommissioning
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m ModuleRequest) Validate() error {
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
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m ModuleRequest) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *ModuleRequest) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComments returns the Comments property
func (m ModuleRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ModuleRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m ModuleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ModuleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ModuleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ModuleRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m ModuleRequest) GetDevice() NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *ModuleRequest) SetDevice(val NestedDeviceRequest) {
	m.Device = val
}

// GetModuleBay returns the ModuleBay property
func (m ModuleRequest) GetModuleBay() NestedModuleBayRequest {
	return m.ModuleBay
}

// SetModuleBay sets the ModuleBay property
func (m *ModuleRequest) SetModuleBay(val NestedModuleBayRequest) {
	m.ModuleBay = val
}

// GetModuleType returns the ModuleType property
func (m ModuleRequest) GetModuleType() NestedModuleTypeRequest {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *ModuleRequest) SetModuleType(val NestedModuleTypeRequest) {
	m.ModuleType = val
}

// GetSerial returns the Serial property
func (m ModuleRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *ModuleRequest) SetSerial(val string) {
	m.Serial = val
}

// GetStatus returns the Status property
func (m ModuleRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *ModuleRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m ModuleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ModuleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
