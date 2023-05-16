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

// PatchedWritableClusterRequest is an object. Adds support for custom fields and tags.
type PatchedWritableClusterRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Group:
	Group *int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Site:
	Site *int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `planned` - Planned
	// * `staging` - Staging
	// * `active` - Active
	// * `decommissioning` - Decommissioning
	// * `offline` - Offline
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type:
	Type int32 `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableClusterRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableClusterRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableClusterRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableClusterRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableClusterRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableClusterRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableClusterRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m PatchedWritableClusterRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *PatchedWritableClusterRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetName returns the Name property
func (m PatchedWritableClusterRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableClusterRequest) SetName(val string) {
	m.Name = val
}

// GetSite returns the Site property
func (m PatchedWritableClusterRequest) GetSite() *int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *PatchedWritableClusterRequest) SetSite(val *int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m PatchedWritableClusterRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableClusterRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableClusterRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableClusterRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableClusterRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableClusterRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetType returns the Type property
func (m PatchedWritableClusterRequest) GetType() int32 {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritableClusterRequest) SetType(val int32) {
	m.Type = val
}
