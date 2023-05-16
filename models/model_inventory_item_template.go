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

// InventoryItemTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type InventoryItemTemplate struct {
	// Depth:
	Depth int32 `json:"_depth" mapstructure:"_depth"`
	// Component:
	Component map[string]interface{} `json:"component" mapstructure:"component"`
	// ComponentId:
	ComponentId *int64 `json:"component_id,omitempty" mapstructure:"component_id,omitempty"`
	// ComponentType:
	ComponentType *string `json:"component_type,omitempty" mapstructure:"component_type,omitempty"`
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
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer *NestedManufacturer `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// PartId: Manufacturer-assigned part identifier
	PartId string `json:"part_id,omitempty" mapstructure:"part_id,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedInventoryItemRole `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m InventoryItemTemplate) Validate() error {
	return validation.Errors{
		"component": validation.Validate(
			m.Component,
		),
		"componentId": validation.Validate(
			m.ComponentId, validation.Min(*int64(0)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"deviceType": validation.Validate(
			m.DeviceType, validation.NotNil,
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDepth returns the Depth property
func (m InventoryItemTemplate) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *InventoryItemTemplate) SetDepth(val int32) {
	m.Depth = val
}

// GetComponent returns the Component property
func (m InventoryItemTemplate) GetComponent() map[string]interface{} {
	return m.Component
}

// SetComponent sets the Component property
func (m *InventoryItemTemplate) SetComponent(val map[string]interface{}) {
	m.Component = val
}

// GetComponentId returns the ComponentId property
func (m InventoryItemTemplate) GetComponentId() *int64 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *InventoryItemTemplate) SetComponentId(val *int64) {
	m.ComponentId = val
}

// GetComponentType returns the ComponentType property
func (m InventoryItemTemplate) GetComponentType() *string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *InventoryItemTemplate) SetComponentType(val *string) {
	m.ComponentType = val
}

// GetCreated returns the Created property
func (m InventoryItemTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *InventoryItemTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m InventoryItemTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *InventoryItemTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m InventoryItemTemplate) GetDeviceType() NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *InventoryItemTemplate) SetDeviceType(val NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m InventoryItemTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *InventoryItemTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m InventoryItemTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *InventoryItemTemplate) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m InventoryItemTemplate) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *InventoryItemTemplate) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m InventoryItemTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *InventoryItemTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetManufacturer returns the Manufacturer property
func (m InventoryItemTemplate) GetManufacturer() *NestedManufacturer {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *InventoryItemTemplate) SetManufacturer(val *NestedManufacturer) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m InventoryItemTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *InventoryItemTemplate) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m InventoryItemTemplate) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *InventoryItemTemplate) SetParent(val *int32) {
	m.Parent = val
}

// GetPartId returns the PartId property
func (m InventoryItemTemplate) GetPartId() string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *InventoryItemTemplate) SetPartId(val string) {
	m.PartId = val
}

// GetRole returns the Role property
func (m InventoryItemTemplate) GetRole() *NestedInventoryItemRole {
	return m.Role
}

// SetRole sets the Role property
func (m *InventoryItemTemplate) SetRole(val *NestedInventoryItemRole) {
	m.Role = val
}

// GetUrl returns the Url property
func (m InventoryItemTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *InventoryItemTemplate) SetUrl(val string) {
	m.Url = val
}
