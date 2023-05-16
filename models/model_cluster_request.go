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

// ClusterRequest is an object. Adds support for custom fields and tags.
type ClusterRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedClusterGroupRequest `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSiteRequest `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `planned` - Planned
	// * `staging` - Staging
	// * `active` - Active
	// * `decommissioning` - Decommissioning
	// * `offline` - Offline
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Type NestedClusterTypeRequest `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m ClusterRequest) Validate() error {
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
			m.Name, validation.Required, validation.Length(1, 100),
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
		"type": validation.Validate(
			m.Type, validation.NotNil,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m ClusterRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ClusterRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m ClusterRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ClusterRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ClusterRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ClusterRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m ClusterRequest) GetGroup() *NestedClusterGroupRequest {
	return m.Group
}

// SetGroup sets the Group property
func (m *ClusterRequest) SetGroup(val *NestedClusterGroupRequest) {
	m.Group = val
}

// GetName returns the Name property
func (m ClusterRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ClusterRequest) SetName(val string) {
	m.Name = val
}

// GetSite returns the Site property
func (m ClusterRequest) GetSite() *NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *ClusterRequest) SetSite(val *NestedSiteRequest) {
	m.Site = val
}

// GetStatus returns the Status property
func (m ClusterRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *ClusterRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m ClusterRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ClusterRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m ClusterRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *ClusterRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetType returns the Type property
func (m ClusterRequest) GetType() NestedClusterTypeRequest {
	return m.Type
}

// SetType sets the Type property
func (m *ClusterRequest) SetType(val NestedClusterTypeRequest) {
	m.Type = val
}
