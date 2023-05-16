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

// PatchedWritableWirelessLANRequest is an object. Adds support for custom fields and tags.
type PatchedWritableWirelessLANRequest struct {
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
	// Group:
	Group *int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Ssid:
	Ssid string `json:"ssid,omitempty" mapstructure:"ssid,omitempty"`
	// Status: * `active` - Active
	// * `reserved` - Reserved
	// * `disabled` - Disabled
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vlan:
	Vlan *int32 `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableWirelessLANRequest) Validate() error {
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
		"ssid": validation.Validate(
			m.Ssid, validation.Length(1, 32),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAuthCipher returns the AuthCipher property
func (m PatchedWritableWirelessLANRequest) GetAuthCipher() string {
	return m.AuthCipher
}

// SetAuthCipher sets the AuthCipher property
func (m *PatchedWritableWirelessLANRequest) SetAuthCipher(val string) {
	m.AuthCipher = val
}

// GetAuthPsk returns the AuthPsk property
func (m PatchedWritableWirelessLANRequest) GetAuthPsk() string {
	return m.AuthPsk
}

// SetAuthPsk sets the AuthPsk property
func (m *PatchedWritableWirelessLANRequest) SetAuthPsk(val string) {
	m.AuthPsk = val
}

// GetAuthType returns the AuthType property
func (m PatchedWritableWirelessLANRequest) GetAuthType() string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *PatchedWritableWirelessLANRequest) SetAuthType(val string) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m PatchedWritableWirelessLANRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableWirelessLANRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableWirelessLANRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableWirelessLANRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableWirelessLANRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableWirelessLANRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m PatchedWritableWirelessLANRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *PatchedWritableWirelessLANRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetSsid returns the Ssid property
func (m PatchedWritableWirelessLANRequest) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *PatchedWritableWirelessLANRequest) SetSsid(val string) {
	m.Ssid = val
}

// GetStatus returns the Status property
func (m PatchedWritableWirelessLANRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableWirelessLANRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableWirelessLANRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableWirelessLANRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableWirelessLANRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableWirelessLANRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVlan returns the Vlan property
func (m PatchedWritableWirelessLANRequest) GetVlan() *int32 {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *PatchedWritableWirelessLANRequest) SetVlan(val *int32) {
	m.Vlan = val
}
