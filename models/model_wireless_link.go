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

// WirelessLink is an object. Adds support for custom fields and tags.
type WirelessLink struct {
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
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// InterfaceA: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	InterfaceA NestedInterface `json:"interface_a" mapstructure:"interface_a"`
	// InterfaceB: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	InterfaceB NestedInterface `json:"interface_b" mapstructure:"interface_b"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Ssid:
	Ssid string `json:"ssid,omitempty" mapstructure:"ssid,omitempty"`
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
}

// Validate implements basic validation for this model
func (m WirelessLink) Validate() error {
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAuthCipher returns the AuthCipher property
func (m WirelessLink) GetAuthCipher() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.AuthCipher
}

// SetAuthCipher sets the AuthCipher property
func (m *WirelessLink) SetAuthCipher(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.AuthCipher = val
}

// GetAuthPsk returns the AuthPsk property
func (m WirelessLink) GetAuthPsk() string {
	return m.AuthPsk
}

// SetAuthPsk sets the AuthPsk property
func (m *WirelessLink) SetAuthPsk(val string) {
	m.AuthPsk = val
}

// GetAuthType returns the AuthType property
func (m WirelessLink) GetAuthType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *WirelessLink) SetAuthType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m WirelessLink) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WirelessLink) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m WirelessLink) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *WirelessLink) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m WirelessLink) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WirelessLink) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WirelessLink) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessLink) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m WirelessLink) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *WirelessLink) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m WirelessLink) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *WirelessLink) SetId(val int32) {
	m.Id = val
}

// GetInterfaceA returns the InterfaceA property
func (m WirelessLink) GetInterfaceA() NestedInterface {
	return m.InterfaceA
}

// SetInterfaceA sets the InterfaceA property
func (m *WirelessLink) SetInterfaceA(val NestedInterface) {
	m.InterfaceA = val
}

// GetInterfaceB returns the InterfaceB property
func (m WirelessLink) GetInterfaceB() NestedInterface {
	return m.InterfaceB
}

// SetInterfaceB sets the InterfaceB property
func (m *WirelessLink) SetInterfaceB(val NestedInterface) {
	m.InterfaceB = val
}

// GetLastUpdated returns the LastUpdated property
func (m WirelessLink) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *WirelessLink) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetSsid returns the Ssid property
func (m WirelessLink) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *WirelessLink) SetSsid(val string) {
	m.Ssid = val
}

// GetStatus returns the Status property
func (m WirelessLink) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *WirelessLink) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WirelessLink) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WirelessLink) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WirelessLink) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WirelessLink) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m WirelessLink) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *WirelessLink) SetUrl(val string) {
	m.Url = val
}
