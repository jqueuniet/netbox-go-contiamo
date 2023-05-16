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

// PatchedWritableInventoryItemRequest is an object. Adds support for custom fields and tags.
type PatchedWritableInventoryItemRequest struct {
	// AssetTag: A unique tag used to identify this item
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// ComponentId:
	ComponentId *int64 `json:"component_id,omitempty" mapstructure:"component_id,omitempty"`
	// ComponentType:
	ComponentType *string `json:"component_type,omitempty" mapstructure:"component_type,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Discovered: This item was automatically discovered
	Discovered bool `json:"discovered,omitempty" mapstructure:"discovered,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Manufacturer:
	Manufacturer *int32 `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// PartId: Manufacturer-assigned part identifier
	PartId string `json:"part_id,omitempty" mapstructure:"part_id,omitempty"`
	// Role:
	Role *int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableInventoryItemRequest) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag, validation.Length(0, 50),
		),
		"componentId": validation.Validate(
			m.ComponentId, validation.Min(*int64(0)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 64),
		),
		"partId": validation.Validate(
			m.PartId, validation.Length(0, 50),
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
func (m PatchedWritableInventoryItemRequest) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *PatchedWritableInventoryItemRequest) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComponentId returns the ComponentId property
func (m PatchedWritableInventoryItemRequest) GetComponentId() *int64 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *PatchedWritableInventoryItemRequest) SetComponentId(val *int64) {
	m.ComponentId = val
}

// GetComponentType returns the ComponentType property
func (m PatchedWritableInventoryItemRequest) GetComponentType() *string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *PatchedWritableInventoryItemRequest) SetComponentType(val *string) {
	m.ComponentType = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableInventoryItemRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableInventoryItemRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableInventoryItemRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableInventoryItemRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PatchedWritableInventoryItemRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *PatchedWritableInventoryItemRequest) SetDevice(val int32) {
	m.Device = val
}

// GetDiscovered returns the Discovered property
func (m PatchedWritableInventoryItemRequest) GetDiscovered() bool {
	return m.Discovered
}

// SetDiscovered sets the Discovered property
func (m *PatchedWritableInventoryItemRequest) SetDiscovered(val bool) {
	m.Discovered = val
}

// GetLabel returns the Label property
func (m PatchedWritableInventoryItemRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PatchedWritableInventoryItemRequest) SetLabel(val string) {
	m.Label = val
}

// GetManufacturer returns the Manufacturer property
func (m PatchedWritableInventoryItemRequest) GetManufacturer() *int32 {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *PatchedWritableInventoryItemRequest) SetManufacturer(val *int32) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m PatchedWritableInventoryItemRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableInventoryItemRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m PatchedWritableInventoryItemRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *PatchedWritableInventoryItemRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetPartId returns the PartId property
func (m PatchedWritableInventoryItemRequest) GetPartId() string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *PatchedWritableInventoryItemRequest) SetPartId(val string) {
	m.PartId = val
}

// GetRole returns the Role property
func (m PatchedWritableInventoryItemRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *PatchedWritableInventoryItemRequest) SetRole(val *int32) {
	m.Role = val
}

// GetSerial returns the Serial property
func (m PatchedWritableInventoryItemRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *PatchedWritableInventoryItemRequest) SetSerial(val string) {
	m.Serial = val
}

// GetTags returns the Tags property
func (m PatchedWritableInventoryItemRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableInventoryItemRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
