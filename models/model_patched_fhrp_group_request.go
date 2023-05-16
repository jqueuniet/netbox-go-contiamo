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

// PatchedFHRPGroupRequest is an object. Adds support for custom fields and tags.
type PatchedFHRPGroupRequest struct {
	// AuthKey:
	AuthKey string `json:"auth_key,omitempty" mapstructure:"auth_key,omitempty"`
	// AuthType: * `plaintext` - Plaintext
	// * `md5` - MD5
	AuthType string `json:"auth_type,omitempty" mapstructure:"auth_type,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// GroupId:
	GroupId int32 `json:"group_id,omitempty" mapstructure:"group_id,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Protocol: * `vrrp2` - VRRPv2
	// * `vrrp3` - VRRPv3
	// * `carp` - CARP
	// * `clusterxl` - ClusterXL
	// * `hsrp` - HSRP
	// * `glbp` - GLBP
	// * `other` - Other
	Protocol string `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedFHRPGroupRequest) Validate() error {
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
			m.GroupId, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAuthKey returns the AuthKey property
func (m PatchedFHRPGroupRequest) GetAuthKey() string {
	return m.AuthKey
}

// SetAuthKey sets the AuthKey property
func (m *PatchedFHRPGroupRequest) SetAuthKey(val string) {
	m.AuthKey = val
}

// GetAuthType returns the AuthType property
func (m PatchedFHRPGroupRequest) GetAuthType() string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *PatchedFHRPGroupRequest) SetAuthType(val string) {
	m.AuthType = val
}

// GetComments returns the Comments property
func (m PatchedFHRPGroupRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedFHRPGroupRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedFHRPGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedFHRPGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedFHRPGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedFHRPGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroupId returns the GroupId property
func (m PatchedFHRPGroupRequest) GetGroupId() int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *PatchedFHRPGroupRequest) SetGroupId(val int32) {
	m.GroupId = val
}

// GetName returns the Name property
func (m PatchedFHRPGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedFHRPGroupRequest) SetName(val string) {
	m.Name = val
}

// GetProtocol returns the Protocol property
func (m PatchedFHRPGroupRequest) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *PatchedFHRPGroupRequest) SetProtocol(val string) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m PatchedFHRPGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedFHRPGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
