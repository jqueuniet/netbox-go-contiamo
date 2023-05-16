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

// WritableVLANRequest is an object. Adds support for custom fields and tags.
type WritableVLANRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Group: VLAN group (optional)
	Group *int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Role: The primary function of this VLAN
	Role *int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site: The specific site to which this VLAN is assigned (if any)
	Site *int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: Operational status of this VLAN
	//
	// * `active` - Active
	// * `reserved` - Reserved
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vid: Numeric VLAN ID (1-4094)
	Vid int32 `json:"vid" mapstructure:"vid"`
}

// Validate implements basic validation for this model
func (m WritableVLANRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"vid": validation.Validate(
			m.Vid, validation.Required, validation.Min(int32(1)), validation.Max(int32(4094)),
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableVLANRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableVLANRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableVLANRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableVLANRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableVLANRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableVLANRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m WritableVLANRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *WritableVLANRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetName returns the Name property
func (m WritableVLANRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableVLANRequest) SetName(val string) {
	m.Name = val
}

// GetRole returns the Role property
func (m WritableVLANRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *WritableVLANRequest) SetRole(val *int32) {
	m.Role = val
}

// GetSite returns the Site property
func (m WritableVLANRequest) GetSite() *int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *WritableVLANRequest) SetSite(val *int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m WritableVLANRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WritableVLANRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WritableVLANRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableVLANRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritableVLANRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritableVLANRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVid returns the Vid property
func (m WritableVLANRequest) GetVid() int32 {
	return m.Vid
}

// SetVid sets the Vid property
func (m *WritableVLANRequest) SetVid(val int32) {
	m.Vid = val
}
