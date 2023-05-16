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

// InventoryItemRequest is an object. Adds support for custom fields and tags.
type InventoryItemRequest struct {
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
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDeviceRequest `json:"device" mapstructure:"device"`
	// Discovered: This item was automatically discovered
	Discovered bool `json:"discovered,omitempty" mapstructure:"discovered,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer *NestedManufacturerRequest `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// PartId: Manufacturer-assigned part identifier
	PartId string `json:"part_id,omitempty" mapstructure:"part_id,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedInventoryItemRoleRequest `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m InventoryItemRequest) Validate() error {
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
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"manufacturer": validation.Validate(
			m.Manufacturer,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"partId": validation.Validate(
			m.PartId, validation.Length(0, 50),
		),
		"role": validation.Validate(
			m.Role,
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
func (m InventoryItemRequest) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *InventoryItemRequest) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComponentId returns the ComponentId property
func (m InventoryItemRequest) GetComponentId() *int64 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *InventoryItemRequest) SetComponentId(val *int64) {
	m.ComponentId = val
}

// GetComponentType returns the ComponentType property
func (m InventoryItemRequest) GetComponentType() *string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *InventoryItemRequest) SetComponentType(val *string) {
	m.ComponentType = val
}

// GetCustomFields returns the CustomFields property
func (m InventoryItemRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *InventoryItemRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m InventoryItemRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *InventoryItemRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m InventoryItemRequest) GetDevice() NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *InventoryItemRequest) SetDevice(val NestedDeviceRequest) {
	m.Device = val
}

// GetDiscovered returns the Discovered property
func (m InventoryItemRequest) GetDiscovered() bool {
	return m.Discovered
}

// SetDiscovered sets the Discovered property
func (m *InventoryItemRequest) SetDiscovered(val bool) {
	m.Discovered = val
}

// GetLabel returns the Label property
func (m InventoryItemRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *InventoryItemRequest) SetLabel(val string) {
	m.Label = val
}

// GetManufacturer returns the Manufacturer property
func (m InventoryItemRequest) GetManufacturer() *NestedManufacturerRequest {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *InventoryItemRequest) SetManufacturer(val *NestedManufacturerRequest) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m InventoryItemRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *InventoryItemRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m InventoryItemRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *InventoryItemRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetPartId returns the PartId property
func (m InventoryItemRequest) GetPartId() string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *InventoryItemRequest) SetPartId(val string) {
	m.PartId = val
}

// GetRole returns the Role property
func (m InventoryItemRequest) GetRole() *NestedInventoryItemRoleRequest {
	return m.Role
}

// SetRole sets the Role property
func (m *InventoryItemRequest) SetRole(val *NestedInventoryItemRoleRequest) {
	m.Role = val
}

// GetSerial returns the Serial property
func (m InventoryItemRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *InventoryItemRequest) SetSerial(val string) {
	m.Serial = val
}

// GetTags returns the Tags property
func (m InventoryItemRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *InventoryItemRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
