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

// PatchedWritableWirelessLinkRequest is an object. Adds support for custom fields and tags.
type PatchedWritableWirelessLinkRequest struct {
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
	// InterfaceA:
	InterfaceA int32 `json:"interface_a,omitempty" mapstructure:"interface_a,omitempty"`
	// InterfaceB:
	InterfaceB int32 `json:"interface_b,omitempty" mapstructure:"interface_b,omitempty"`
	// Ssid:
	Ssid string `json:"ssid,omitempty" mapstructure:"ssid,omitempty"`
	// Status: * `connected` - Connected
	// * `planned` - Planned
	// * `decommissioning` - Decommissioning
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableWirelessLinkRequest) Validate() error {
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
			m.Ssid, validation.Length(0, 32),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAuthCipher returns the AuthCipher property
func (m PatchedWritableWirelessLinkRequest) GetAuthCipher() string {
	return m.AuthCipher
}

// SetAuthCipher sets the AuthCipher property
func (m *PatchedWritableWirelessLinkRequest) SetAuthCipher(val string) {
	m.AuthCipher = val
}

// GetAuthPsk returns the AuthPsk property
func (m PatchedWritableWirelessLinkRequest) GetAuthPsk() string {
	return m.AuthPsk
}

// SetAuthPsk sets the AuthPsk property
func (m *PatchedWritableWirelessLinkRequest) SetAuthPsk(val string) {
	m.AuthPsk = val
}

// GetAuthType returns the AuthType property
func (m PatchedWritableWirelessLinkRequest) GetAuthType() string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *PatchedWritableWirelessLinkRequest) SetAuthType(val string) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m PatchedWritableWirelessLinkRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableWirelessLinkRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableWirelessLinkRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableWirelessLinkRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableWirelessLinkRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableWirelessLinkRequest) SetDescription(val string) {
	m.Description = val
}

// GetInterfaceA returns the InterfaceA property
func (m PatchedWritableWirelessLinkRequest) GetInterfaceA() int32 {
	return m.InterfaceA
}

// SetInterfaceA sets the InterfaceA property
func (m *PatchedWritableWirelessLinkRequest) SetInterfaceA(val int32) {
	m.InterfaceA = val
}

// GetInterfaceB returns the InterfaceB property
func (m PatchedWritableWirelessLinkRequest) GetInterfaceB() int32 {
	return m.InterfaceB
}

// SetInterfaceB sets the InterfaceB property
func (m *PatchedWritableWirelessLinkRequest) SetInterfaceB(val int32) {
	m.InterfaceB = val
}

// GetSsid returns the Ssid property
func (m PatchedWritableWirelessLinkRequest) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *PatchedWritableWirelessLinkRequest) SetSsid(val string) {
	m.Ssid = val
}

// GetStatus returns the Status property
func (m PatchedWritableWirelessLinkRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableWirelessLinkRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableWirelessLinkRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableWirelessLinkRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableWirelessLinkRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableWirelessLinkRequest) SetTenant(val *int32) {
	m.Tenant = val
}
