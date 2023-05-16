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

// InventoryItemTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type InventoryItemTemplateRequest struct {
	// ComponentId:
	ComponentId *int64 `json:"component_id,omitempty" mapstructure:"component_id,omitempty"`
	// ComponentType:
	ComponentType *string `json:"component_type,omitempty" mapstructure:"component_type,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType NestedDeviceTypeRequest `json:"device_type" mapstructure:"device_type"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer *NestedManufacturerRequest `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
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
	Role *NestedInventoryItemRoleRequest `json:"role,omitempty" mapstructure:"role,omitempty"`
}

// Validate implements basic validation for this model
func (m InventoryItemTemplateRequest) Validate() error {
	return validation.Errors{
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
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"partId": validation.Validate(
			m.PartId, validation.Length(0, 50),
		),
		"role": validation.Validate(
			m.Role,
		),
	}.Filter()
}

// GetComponentId returns the ComponentId property
func (m InventoryItemTemplateRequest) GetComponentId() *int64 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *InventoryItemTemplateRequest) SetComponentId(val *int64) {
	m.ComponentId = val
}

// GetComponentType returns the ComponentType property
func (m InventoryItemTemplateRequest) GetComponentType() *string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *InventoryItemTemplateRequest) SetComponentType(val *string) {
	m.ComponentType = val
}

// GetDescription returns the Description property
func (m InventoryItemTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *InventoryItemTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m InventoryItemTemplateRequest) GetDeviceType() NestedDeviceTypeRequest {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *InventoryItemTemplateRequest) SetDeviceType(val NestedDeviceTypeRequest) {
	m.DeviceType = val
}

// GetLabel returns the Label property
func (m InventoryItemTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *InventoryItemTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetManufacturer returns the Manufacturer property
func (m InventoryItemTemplateRequest) GetManufacturer() *NestedManufacturerRequest {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *InventoryItemTemplateRequest) SetManufacturer(val *NestedManufacturerRequest) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m InventoryItemTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *InventoryItemTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m InventoryItemTemplateRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *InventoryItemTemplateRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetPartId returns the PartId property
func (m InventoryItemTemplateRequest) GetPartId() string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *InventoryItemTemplateRequest) SetPartId(val string) {
	m.PartId = val
}

// GetRole returns the Role property
func (m InventoryItemTemplateRequest) GetRole() *NestedInventoryItemRoleRequest {
	return m.Role
}

// SetRole sets the Role property
func (m *InventoryItemTemplateRequest) SetRole(val *NestedInventoryItemRoleRequest) {
	m.Role = val
}
