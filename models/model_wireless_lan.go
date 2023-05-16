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

// WirelessLAN is an object. Adds support for custom fields and tags.
type WirelessLAN struct {
	// AuthCipher:
	AuthCipher *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"auth_cipher,omitempty" mapstructure:"auth_cipher,omitempty"`
	// AuthPsk:
	AuthPsk string `json:"auth_psk,omitempty" mapstructure:"auth_psk,omitempty"`
	// AuthType:
	AuthType *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"auth_type,omitempty" mapstructure:"auth_type,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedWirelessLANGroup `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Ssid:
	Ssid string `json:"ssid" mapstructure:"ssid"`
	// Status:
	Status *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Vlan: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vlan *NestedVLAN `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
}

// Validate implements basic validation for this model
func (m WirelessLAN) Validate() error {
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
			m.Ssid, validation.NotNil, validation.Length(0, 32),
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
		"vlan": validation.Validate(
			m.Vlan,
		),
	}.Filter()
}

// GetAuthCipher returns the AuthCipher property
func (m WirelessLAN) GetAuthCipher() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.AuthCipher
}

// SetAuthCipher sets the AuthCipher property
func (m *WirelessLAN) SetAuthCipher(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.AuthCipher = val
}

// GetAuthPsk returns the AuthPsk property
func (m WirelessLAN) GetAuthPsk() string {
	return m.AuthPsk
}

// SetAuthPsk sets the AuthPsk property
func (m *WirelessLAN) SetAuthPsk(val string) {
	m.AuthPsk = val
}

// GetAuthType returns the AuthType property
func (m WirelessLAN) GetAuthType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *WirelessLAN) SetAuthType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m WirelessLAN) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WirelessLAN) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m WirelessLAN) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *WirelessLAN) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m WirelessLAN) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WirelessLAN) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WirelessLAN) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessLAN) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m WirelessLAN) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *WirelessLAN) SetDisplay(val string) {
	m.Display = val
}

// GetGroup returns the Group property
func (m WirelessLAN) GetGroup() *NestedWirelessLANGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *WirelessLAN) SetGroup(val *NestedWirelessLANGroup) {
	m.Group = val
}

// GetId returns the Id property
func (m WirelessLAN) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *WirelessLAN) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m WirelessLAN) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *WirelessLAN) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetSsid returns the Ssid property
func (m WirelessLAN) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *WirelessLAN) SetSsid(val string) {
	m.Ssid = val
}

// GetStatus returns the Status property
func (m WirelessLAN) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *WirelessLAN) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WirelessLAN) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WirelessLAN) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WirelessLAN) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WirelessLAN) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m WirelessLAN) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *WirelessLAN) SetUrl(val string) {
	m.Url = val
}

// GetVlan returns the Vlan property
func (m WirelessLAN) GetVlan() *NestedVLAN {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *WirelessLAN) SetVlan(val *NestedVLAN) {
	m.Vlan = val
}
