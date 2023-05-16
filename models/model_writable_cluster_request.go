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

// WritableClusterRequest is an object. Adds support for custom fields and tags.
type WritableClusterRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Group:
	Group *int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
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
	Type int32 `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m WritableClusterRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableClusterRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableClusterRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableClusterRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableClusterRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableClusterRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableClusterRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m WritableClusterRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *WritableClusterRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetName returns the Name property
func (m WritableClusterRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableClusterRequest) SetName(val string) {
	m.Name = val
}

// GetSite returns the Site property
func (m WritableClusterRequest) GetSite() *int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *WritableClusterRequest) SetSite(val *int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m WritableClusterRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WritableClusterRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WritableClusterRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableClusterRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritableClusterRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritableClusterRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetType returns the Type property
func (m WritableClusterRequest) GetType() int32 {
	return m.Type
}

// SetType sets the Type property
func (m *WritableClusterRequest) SetType(val int32) {
	m.Type = val
}
