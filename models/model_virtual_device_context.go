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

// VirtualDeviceContext is an object. Adds support for custom fields and tags.
type VirtualDeviceContext struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
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
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Identifier: Numeric identifier unique to the parent device
	Identifier *int32 `json:"identifier,omitempty" mapstructure:"identifier,omitempty"`
	// InterfaceCount:
	InterfaceCount int32 `json:"interface_count" mapstructure:"interface_count"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PrimaryIp: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp *NestedIPAddress `json:"primary_ip" mapstructure:"primary_ip"`
	// PrimaryIp4: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Status: * `active` - Active
	// * `planned` - Planned
	// * `offline` - Offline
	Status string `json:"status" mapstructure:"status"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m VirtualDeviceContext) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"primaryIp": validation.Validate(
			m.PrimaryIp,
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m VirtualDeviceContext) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VirtualDeviceContext) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m VirtualDeviceContext) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VirtualDeviceContext) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m VirtualDeviceContext) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VirtualDeviceContext) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VirtualDeviceContext) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VirtualDeviceContext) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m VirtualDeviceContext) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *VirtualDeviceContext) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m VirtualDeviceContext) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *VirtualDeviceContext) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m VirtualDeviceContext) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VirtualDeviceContext) SetId(val int32) {
	m.Id = val
}

// GetIdentifier returns the Identifier property
func (m VirtualDeviceContext) GetIdentifier() *int32 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *VirtualDeviceContext) SetIdentifier(val *int32) {
	m.Identifier = val
}

// GetInterfaceCount returns the InterfaceCount property
func (m VirtualDeviceContext) GetInterfaceCount() int32 {
	return m.InterfaceCount
}

// SetInterfaceCount sets the InterfaceCount property
func (m *VirtualDeviceContext) SetInterfaceCount(val int32) {
	m.InterfaceCount = val
}

// GetLastUpdated returns the LastUpdated property
func (m VirtualDeviceContext) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VirtualDeviceContext) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m VirtualDeviceContext) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualDeviceContext) SetName(val string) {
	m.Name = val
}

// GetPrimaryIp returns the PrimaryIp property
func (m VirtualDeviceContext) GetPrimaryIp() *NestedIPAddress {
	return m.PrimaryIp
}

// SetPrimaryIp sets the PrimaryIp property
func (m *VirtualDeviceContext) SetPrimaryIp(val *NestedIPAddress) {
	m.PrimaryIp = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m VirtualDeviceContext) GetPrimaryIp4() *NestedIPAddress {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *VirtualDeviceContext) SetPrimaryIp4(val *NestedIPAddress) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m VirtualDeviceContext) GetPrimaryIp6() *NestedIPAddress {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *VirtualDeviceContext) SetPrimaryIp6(val *NestedIPAddress) {
	m.PrimaryIp6 = val
}

// GetStatus returns the Status property
func (m VirtualDeviceContext) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *VirtualDeviceContext) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m VirtualDeviceContext) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VirtualDeviceContext) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VirtualDeviceContext) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VirtualDeviceContext) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m VirtualDeviceContext) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *VirtualDeviceContext) SetUrl(val string) {
	m.Url = val
}
