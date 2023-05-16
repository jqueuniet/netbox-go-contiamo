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

// WritableVirtualDeviceContextRequest is an object. Adds support for custom fields and tags.
type WritableVirtualDeviceContextRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device *int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Identifier: Numeric identifier unique to the parent device
	Identifier *int32 `json:"identifier,omitempty" mapstructure:"identifier,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PrimaryIp4:
	PrimaryIp4 *int32 `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6:
	PrimaryIp6 *int32 `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Status: * `active` - Active
	// * `planned` - Planned
	// * `offline` - Offline
	Status string `json:"status" mapstructure:"status"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableVirtualDeviceContextRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"identifier": validation.Validate(
			m.Identifier, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableVirtualDeviceContextRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableVirtualDeviceContextRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableVirtualDeviceContextRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableVirtualDeviceContextRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableVirtualDeviceContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableVirtualDeviceContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m WritableVirtualDeviceContextRequest) GetDevice() *int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *WritableVirtualDeviceContextRequest) SetDevice(val *int32) {
	m.Device = val
}

// GetIdentifier returns the Identifier property
func (m WritableVirtualDeviceContextRequest) GetIdentifier() *int32 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *WritableVirtualDeviceContextRequest) SetIdentifier(val *int32) {
	m.Identifier = val
}

// GetName returns the Name property
func (m WritableVirtualDeviceContextRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableVirtualDeviceContextRequest) SetName(val string) {
	m.Name = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m WritableVirtualDeviceContextRequest) GetPrimaryIp4() *int32 {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *WritableVirtualDeviceContextRequest) SetPrimaryIp4(val *int32) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m WritableVirtualDeviceContextRequest) GetPrimaryIp6() *int32 {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *WritableVirtualDeviceContextRequest) SetPrimaryIp6(val *int32) {
	m.PrimaryIp6 = val
}

// GetStatus returns the Status property
func (m WritableVirtualDeviceContextRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WritableVirtualDeviceContextRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WritableVirtualDeviceContextRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableVirtualDeviceContextRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritableVirtualDeviceContextRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritableVirtualDeviceContextRequest) SetTenant(val *int32) {
	m.Tenant = val
}
