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

// PatchedWritableModuleRequest is an object. Adds support for custom fields and tags.
type PatchedWritableModuleRequest struct {
	// AssetTag: A unique tag used to identify this device
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// ModuleBay:
	ModuleBay int32 `json:"module_bay,omitempty" mapstructure:"module_bay,omitempty"`
	// ModuleType:
	ModuleType int32 `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
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
func (m PatchedWritableModuleRequest) Validate() error {
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
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m PatchedWritableModuleRequest) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *PatchedWritableModuleRequest) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComments returns the Comments property
func (m PatchedWritableModuleRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableModuleRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableModuleRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableModuleRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableModuleRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableModuleRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PatchedWritableModuleRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *PatchedWritableModuleRequest) SetDevice(val int32) {
	m.Device = val
}

// GetModuleBay returns the ModuleBay property
func (m PatchedWritableModuleRequest) GetModuleBay() int32 {
	return m.ModuleBay
}

// SetModuleBay sets the ModuleBay property
func (m *PatchedWritableModuleRequest) SetModuleBay(val int32) {
	m.ModuleBay = val
}

// GetModuleType returns the ModuleType property
func (m PatchedWritableModuleRequest) GetModuleType() int32 {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *PatchedWritableModuleRequest) SetModuleType(val int32) {
	m.ModuleType = val
}

// GetSerial returns the Serial property
func (m PatchedWritableModuleRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *PatchedWritableModuleRequest) SetSerial(val string) {
	m.Serial = val
}

// GetStatus returns the Status property
func (m PatchedWritableModuleRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableModuleRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableModuleRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableModuleRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
