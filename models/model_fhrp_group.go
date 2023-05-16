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

// FHRPGroup is an object. Adds support for custom fields and tags.
type FHRPGroup struct {
	// AuthKey:
	AuthKey string `json:"auth_key,omitempty" mapstructure:"auth_key,omitempty"`
	// AuthType: * `plaintext` - Plaintext
	// * `md5` - MD5
	AuthType string `json:"auth_type,omitempty" mapstructure:"auth_type,omitempty"`
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
	// GroupId:
	GroupId int32 `json:"group_id" mapstructure:"group_id"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IpAddresses:
	IpAddresses []NestedIPAddress `json:"ip_addresses" mapstructure:"ip_addresses"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Protocol: * `vrrp2` - VRRPv2
	// * `vrrp3` - VRRPv3
	// * `carp` - CARP
	// * `clusterxl` - ClusterXL
	// * `hsrp` - HSRP
	// * `glbp` - GLBP
	// * `other` - Other
	Protocol string `json:"protocol" mapstructure:"protocol"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m FHRPGroup) Validate() error {
	return validation.Errors{
		"authKey": validation.Validate(
			m.AuthKey, validation.Length(0, 255),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"groupId": validation.Validate(
			m.GroupId, validation.NotNil, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"ipAddresses": validation.Validate(
			m.IpAddresses, validation.NotNil,
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAuthKey returns the AuthKey property
func (m FHRPGroup) GetAuthKey() string {
	return m.AuthKey
}

// SetAuthKey sets the AuthKey property
func (m *FHRPGroup) SetAuthKey(val string) {
	m.AuthKey = val
}

// GetAuthType returns the AuthType property
func (m FHRPGroup) GetAuthType() string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *FHRPGroup) SetAuthType(val string) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m FHRPGroup) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *FHRPGroup) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m FHRPGroup) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *FHRPGroup) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m FHRPGroup) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *FHRPGroup) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m FHRPGroup) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *FHRPGroup) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m FHRPGroup) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *FHRPGroup) SetDisplay(val string) {
	m.Display = val
}

// GetGroupId returns the GroupId property
func (m FHRPGroup) GetGroupId() int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *FHRPGroup) SetGroupId(val int32) {
	m.GroupId = val
}

// GetId returns the Id property
func (m FHRPGroup) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *FHRPGroup) SetId(val int32) {
	m.Id = val
}

// GetIpAddresses returns the IpAddresses property
func (m FHRPGroup) GetIpAddresses() []NestedIPAddress {
	return m.IpAddresses
}

// SetIpAddresses sets the IpAddresses property
func (m *FHRPGroup) SetIpAddresses(val []NestedIPAddress) {
	m.IpAddresses = val
}

// GetLastUpdated returns the LastUpdated property
func (m FHRPGroup) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *FHRPGroup) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m FHRPGroup) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *FHRPGroup) SetName(val string) {
	m.Name = val
}

// GetProtocol returns the Protocol property
func (m FHRPGroup) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *FHRPGroup) SetProtocol(val string) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m FHRPGroup) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *FHRPGroup) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m FHRPGroup) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *FHRPGroup) SetUrl(val string) {
	m.Url = val
}
