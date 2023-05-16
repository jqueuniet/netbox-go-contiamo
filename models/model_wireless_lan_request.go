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

// WirelessLANRequest is an object. Adds support for custom fields and tags.
type WirelessLANRequest struct {
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
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedWirelessLANGroupRequest `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Ssid:
	Ssid string `json:"ssid" mapstructure:"ssid"`
	// Status: * `active` - Active
	// * `reserved` - Reserved
	// * `disabled` - Disabled
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vlan: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vlan *NestedVLANRequest `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
}

// Validate implements basic validation for this model
func (m WirelessLANRequest) Validate() error {
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
		"group": validation.Validate(
			m.Group,
		),
		"ssid": validation.Validate(
			m.Ssid, validation.Required, validation.Length(1, 32),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"vlan": validation.Validate(
			m.Vlan,
		),
	}.Filter()
}

// GetAuthCipher returns the AuthCipher property
func (m WirelessLANRequest) GetAuthCipher() string {
	return m.AuthCipher
}

// SetAuthCipher sets the AuthCipher property
func (m *WirelessLANRequest) SetAuthCipher(val string) {
	m.AuthCipher = val
}

// GetAuthPsk returns the AuthPsk property
func (m WirelessLANRequest) GetAuthPsk() string {
	return m.AuthPsk
}

// SetAuthPsk sets the AuthPsk property
func (m *WirelessLANRequest) SetAuthPsk(val string) {
	m.AuthPsk = val
}

// GetAuthType returns the AuthType property
func (m WirelessLANRequest) GetAuthType() string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *WirelessLANRequest) SetAuthType(val string) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m WirelessLANRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WirelessLANRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WirelessLANRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WirelessLANRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WirelessLANRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessLANRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m WirelessLANRequest) GetGroup() *NestedWirelessLANGroupRequest {
	return m.Group
}

// SetGroup sets the Group property
func (m *WirelessLANRequest) SetGroup(val *NestedWirelessLANGroupRequest) {
	m.Group = val
}

// GetSsid returns the Ssid property
func (m WirelessLANRequest) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *WirelessLANRequest) SetSsid(val string) {
	m.Ssid = val
}

// GetStatus returns the Status property
func (m WirelessLANRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WirelessLANRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WirelessLANRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WirelessLANRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WirelessLANRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WirelessLANRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetVlan returns the Vlan property
func (m WirelessLANRequest) GetVlan() *NestedVLANRequest {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *WirelessLANRequest) SetVlan(val *NestedVLANRequest) {
	m.Vlan = val
}
