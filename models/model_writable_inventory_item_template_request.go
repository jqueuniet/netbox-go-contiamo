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

// WritableInventoryItemTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableInventoryItemTemplateRequest struct {
	// ComponentId:
	ComponentId *int64 `json:"component_id,omitempty" mapstructure:"component_id,omitempty"`
	// ComponentType:
	ComponentType *string `json:"component_type,omitempty" mapstructure:"component_type,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType:
	DeviceType int32 `json:"device_type" mapstructure:"device_type"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Manufacturer:
	Manufacturer *int32 `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// PartId: Manufacturer-assigned part identifier
	PartId string `json:"part_id,omitempty" mapstructure:"part_id,omitempty"`
	// Role:
	Role *int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableInventoryItemTemplateRequest) Validate() error {
	return validation.Errors{
		"componentId": validation.Validate(
			m.ComponentId, validation.Min(*int64(0)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"partId": validation.Validate(
			m.PartId, validation.Length(0, 50),
		),
	}.Filter()
}

// GetComponentId returns the ComponentId property
func (m WritableInventoryItemTemplateRequest) GetComponentId() *int64 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *WritableInventoryItemTemplateRequest) SetComponentId(val *int64) {
	m.ComponentId = val
}

// GetComponentType returns the ComponentType property
func (m WritableInventoryItemTemplateRequest) GetComponentType() *string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *WritableInventoryItemTemplateRequest) SetComponentType(val *string) {
	m.ComponentType = val
}

// GetDescription returns the Description property
func (m WritableInventoryItemTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableInventoryItemTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m WritableInventoryItemTemplateRequest) GetDeviceType() int32 {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *WritableInventoryItemTemplateRequest) SetDeviceType(val int32) {
	m.DeviceType = val
}

// GetLabel returns the Label property
func (m WritableInventoryItemTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *WritableInventoryItemTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetManufacturer returns the Manufacturer property
func (m WritableInventoryItemTemplateRequest) GetManufacturer() *int32 {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *WritableInventoryItemTemplateRequest) SetManufacturer(val *int32) {
	m.Manufacturer = val
}

// GetName returns the Name property
func (m WritableInventoryItemTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableInventoryItemTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m WritableInventoryItemTemplateRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *WritableInventoryItemTemplateRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetPartId returns the PartId property
func (m WritableInventoryItemTemplateRequest) GetPartId() string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *WritableInventoryItemTemplateRequest) SetPartId(val string) {
	m.PartId = val
}

// GetRole returns the Role property
func (m WritableInventoryItemTemplateRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *WritableInventoryItemTemplateRequest) SetRole(val *int32) {
	m.Role = val
}
