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

// VLANRequest is an object. Adds support for custom fields and tags.
type VLANRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedVLANGroupRequest `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedRoleRequest `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSiteRequest `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `active` - Active
	// * `reserved` - Reserved
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vid: Numeric VLAN ID (1-4094)
	Vid int32 `json:"vid" mapstructure:"vid"`
}

// Validate implements basic validation for this model
func (m VLANRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"group": validation.Validate(
			m.Group,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"role": validation.Validate(
			m.Role,
		),
		"site": validation.Validate(
			m.Site,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"vid": validation.Validate(
			m.Vid, validation.Required, validation.Min(int32(1)), validation.Max(int32(4094)),
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m VLANRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VLANRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m VLANRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VLANRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VLANRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VLANRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m VLANRequest) GetGroup() *NestedVLANGroupRequest {
	return m.Group
}

// SetGroup sets the Group property
func (m *VLANRequest) SetGroup(val *NestedVLANGroupRequest) {
	m.Group = val
}

// GetName returns the Name property
func (m VLANRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VLANRequest) SetName(val string) {
	m.Name = val
}

// GetRole returns the Role property
func (m VLANRequest) GetRole() *NestedRoleRequest {
	return m.Role
}

// SetRole sets the Role property
func (m *VLANRequest) SetRole(val *NestedRoleRequest) {
	m.Role = val
}

// GetSite returns the Site property
func (m VLANRequest) GetSite() *NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *VLANRequest) SetSite(val *NestedSiteRequest) {
	m.Site = val
}

// GetStatus returns the Status property
func (m VLANRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *VLANRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m VLANRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VLANRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VLANRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VLANRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetVid returns the Vid property
func (m VLANRequest) GetVid() int32 {
	return m.Vid
}

// SetVid sets the Vid property
func (m *VLANRequest) SetVid(val int32) {
	m.Vid = val
}
