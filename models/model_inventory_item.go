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

// InventoryItem is an object. Adds support for custom fields and tags.
type InventoryItem struct {
	// Depth:
	Depth int32 `json:"_depth" mapstructure:"_depth"`
	// AssetTag: A unique tag used to identify this item
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// Component:
	Component map[string]interface{} `json:"component" mapstructure:"component"`
	// ComponentId:
	ComponentId *int64 `json:"component_id,omitempty" mapstructure:"component_id,omitempty"`
	// ComponentType:
	ComponentType *string `json:"component_type,omitempty" mapstructure:"component_type,omitempty"`
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
	// Discovered: This item was automatically discovered
	Discovered bool `json:"discovered,omitempty" mapstructure:"discovered,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer *NestedManufacturer `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// PartId: Manufacturer-assigned part identifier
	PartId string `json:"part_id,omitempty" mapstructure:"part_id,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedInventoryItemRole `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m InventoryItem) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag, validation.Length(0, 50),
		),
		"component": validation.Validate(
			m.Component,
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
			m.Name, validation.NotNil, validation.Length(0, 64),
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDepth returns the Depth property
func (m InventoryItem) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *InventoryItem) SetDepth(val int32) {
	m.Depth = val
}

// GetAssetTag returns the AssetTag property
func (m InventoryItem) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *InventoryItem) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComponent returns the Component property
func (m InventoryItem) GetComponent() map[string]interface{} {
	return m.Component
}

// SetComponent sets the Component property
func (m *InventoryItem) SetComponent(val map[string]interface{}) {
	m.Component = val
}

// GetComponentId returns the ComponentId property
func (m InventoryItem) GetComponentId() *int64 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *InventoryItem) SetComponentId(val *int64) {
	m.ComponentId = val
}

// GetComponentType returns the ComponentType property
func (m InventoryItem) GetComponentType() *string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *InventoryItem) SetComponentType(val *string) {
	m.ComponentType = val
}

// GetCreated returns the Created property
func (m InventoryItem) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *InventoryItem) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m InventoryItem) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *InventoryItem) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m InventoryItem) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *InventoryItem) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m InventoryItem) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *InventoryItem) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDiscovered returns the Discovered property
func (m InventoryItem) GetDiscovered() bool {
	return m.Discovered
}

// SetDiscovered sets the Discovered property
func (m *InventoryItem) SetDiscovered(val bool) {
	m.Discovered = val
}

// GetDisplay returns the Display property
func (m InventoryItem) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *InventoryItem) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m InventoryItem) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *InventoryItem) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m InventoryItem) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *InventoryItem) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m InventoryItem) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *InventoryItem) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetManufacturer returns the Manufacturer property
func (m InventoryItem) GetManufacturer() *NestedManufacturer {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *InventoryItem) SetManufacturer(val *NestedManufacturer) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m InventoryItem) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *InventoryItem) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m InventoryItem) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *InventoryItem) SetParent(val *int32) {
	m.Parent = val
}

// GetPartId returns the PartId property
func (m InventoryItem) GetPartId() string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *InventoryItem) SetPartId(val string) {
	m.PartId = val
}

// GetRole returns the Role property
func (m InventoryItem) GetRole() *NestedInventoryItemRole {
	return m.Role
}

// SetRole sets the Role property
func (m *InventoryItem) SetRole(val *NestedInventoryItemRole) {
	m.Role = val
}

// GetSerial returns the Serial property
func (m InventoryItem) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *InventoryItem) SetSerial(val string) {
	m.Serial = val
}

// GetTags returns the Tags property
func (m InventoryItem) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *InventoryItem) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m InventoryItem) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *InventoryItem) SetUrl(val string) {
	m.Url = val
}
