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

// WirelessLinkRequest is an object. Adds support for custom fields and tags.
type WirelessLinkRequest struct {
	// AuthCipher: * `auto` - Auto
	// * `tkip` - TKIP
	// * `aes` - AES
	AuthCipher string `json:"auth_cipher,omitempty" mapstructure:"auth_cipher,omitempty"`
	// AuthPsk:
	AuthPsk string `json:"auth_psk,omitempty" mapstructure:"auth_psk,omitempty"`
	// AuthType: * `open` - Open
	// * `wep` - WEP
	// * `wpa-personal` - WPA Personal (PSK)
	// * `wpa-enterprise` - WPA Enterprise
	AuthType string `json:"auth_type,omitempty" mapstructure:"auth_type,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// InterfaceA: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	InterfaceA NestedInterfaceRequest `json:"interface_a" mapstructure:"interface_a"`
	// InterfaceB: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	InterfaceB NestedInterfaceRequest `json:"interface_b" mapstructure:"interface_b"`
	// Ssid:
	Ssid string `json:"ssid,omitempty" mapstructure:"ssid,omitempty"`
	// Status: * `connected` - Connected
	// * `planned` - Planned
	// * `decommissioning` - Decommissioning
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m WirelessLinkRequest) Validate() error {
	return validation.Errors{
		"authPsk": validation.Validate(
			m.AuthPsk, validation.Length(0, 64),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"interfaceA": validation.Validate(
			m.InterfaceA, validation.NotNil,
		),
		"interfaceB": validation.Validate(
			m.InterfaceB, validation.NotNil,
		),
		"ssid": validation.Validate(
			m.Ssid, validation.Length(0, 32),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
	}.Filter()
}

// GetAuthCipher returns the AuthCipher property
func (m WirelessLinkRequest) GetAuthCipher() string {
	return m.AuthCipher
}

// SetAuthCipher sets the AuthCipher property
func (m *WirelessLinkRequest) SetAuthCipher(val string) {
	m.AuthCipher = val
}

// GetAuthPsk returns the AuthPsk property
func (m WirelessLinkRequest) GetAuthPsk() string {
	return m.AuthPsk
}

// SetAuthPsk sets the AuthPsk property
func (m *WirelessLinkRequest) SetAuthPsk(val string) {
	m.AuthPsk = val
}

// GetAuthType returns the AuthType property
func (m WirelessLinkRequest) GetAuthType() string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *WirelessLinkRequest) SetAuthType(val string) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m WirelessLinkRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WirelessLinkRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WirelessLinkRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WirelessLinkRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WirelessLinkRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessLinkRequest) SetDescription(val string) {
	m.Description = val
}

// GetInterfaceA returns the InterfaceA property
func (m WirelessLinkRequest) GetInterfaceA() NestedInterfaceRequest {
	return m.InterfaceA
}

// SetInterfaceA sets the InterfaceA property
func (m *WirelessLinkRequest) SetInterfaceA(val NestedInterfaceRequest) {
	m.InterfaceA = val
}

// GetInterfaceB returns the InterfaceB property
func (m WirelessLinkRequest) GetInterfaceB() NestedInterfaceRequest {
	return m.InterfaceB
}

// SetInterfaceB sets the InterfaceB property
func (m *WirelessLinkRequest) SetInterfaceB(val NestedInterfaceRequest) {
	m.InterfaceB = val
}

// GetSsid returns the Ssid property
func (m WirelessLinkRequest) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *WirelessLinkRequest) SetSsid(val string) {
	m.Ssid = val
}

// GetStatus returns the Status property
func (m WirelessLinkRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WirelessLinkRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WirelessLinkRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WirelessLinkRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WirelessLinkRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WirelessLinkRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
