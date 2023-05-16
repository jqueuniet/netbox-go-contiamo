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

// VirtualDeviceContextRequest is an object. Adds support for custom fields and tags.
type VirtualDeviceContextRequest struct {
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
	// Identifier: Numeric identifier unique to the parent device
	Identifier *int32 `json:"identifier,omitempty" mapstructure:"identifier,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PrimaryIp4: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp4 *NestedIPAddressRequest `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp6 *NestedIPAddressRequest `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Status: * `active` - Active
	// * `planned` - Planned
	// * `offline` - Offline
	Status string `json:"status" mapstructure:"status"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m VirtualDeviceContextRequest) Validate() error {
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
		"identifier": validation.Validate(
			m.Identifier, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"primaryIp4": validation.Validate(
			m.PrimaryIp4,
		),
		"primaryIp6": validation.Validate(
			m.PrimaryIp6,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m VirtualDeviceContextRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VirtualDeviceContextRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m VirtualDeviceContextRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VirtualDeviceContextRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VirtualDeviceContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VirtualDeviceContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m VirtualDeviceContextRequest) GetDevice() NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *VirtualDeviceContextRequest) SetDevice(val NestedDeviceRequest) {
	m.Device = val
}

// GetIdentifier returns the Identifier property
func (m VirtualDeviceContextRequest) GetIdentifier() *int32 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *VirtualDeviceContextRequest) SetIdentifier(val *int32) {
	m.Identifier = val
}

// GetName returns the Name property
func (m VirtualDeviceContextRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualDeviceContextRequest) SetName(val string) {
	m.Name = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m VirtualDeviceContextRequest) GetPrimaryIp4() *NestedIPAddressRequest {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *VirtualDeviceContextRequest) SetPrimaryIp4(val *NestedIPAddressRequest) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m VirtualDeviceContextRequest) GetPrimaryIp6() *NestedIPAddressRequest {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *VirtualDeviceContextRequest) SetPrimaryIp6(val *NestedIPAddressRequest) {
	m.PrimaryIp6 = val
}

// GetStatus returns the Status property
func (m VirtualDeviceContextRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *VirtualDeviceContextRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m VirtualDeviceContextRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VirtualDeviceContextRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VirtualDeviceContextRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VirtualDeviceContextRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
