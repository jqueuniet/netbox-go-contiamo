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

// PrefixRequest is an object. Adds support for custom fields and tags.
type PrefixRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// IsPool: All IP addresses within this prefix are considered usable
	IsPool bool `json:"is_pool,omitempty" mapstructure:"is_pool,omitempty"`
	// MarkUtilized: Treat as 100% utilized
	MarkUtilized bool `json:"mark_utilized,omitempty" mapstructure:"mark_utilized,omitempty"`
	// Prefix:
	Prefix string `json:"prefix" mapstructure:"prefix"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedRoleRequest `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSiteRequest `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `container` - Container
	// * `active` - Active
	// * `reserved` - Reserved
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vlan: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vlan *NestedVLANRequest `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRFRequest `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m PrefixRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"prefix": validation.Validate(
			m.Prefix, validation.Required, validation.Length(1, 0),
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
		"vlan": validation.Validate(
			m.Vlan,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PrefixRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PrefixRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PrefixRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PrefixRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PrefixRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PrefixRequest) SetDescription(val string) {
	m.Description = val
}

// GetIsPool returns the IsPool property
func (m PrefixRequest) GetIsPool() bool {
	return m.IsPool
}

// SetIsPool sets the IsPool property
func (m *PrefixRequest) SetIsPool(val bool) {
	m.IsPool = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m PrefixRequest) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *PrefixRequest) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetPrefix returns the Prefix property
func (m PrefixRequest) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *PrefixRequest) SetPrefix(val string) {
	m.Prefix = val
}

// GetRole returns the Role property
func (m PrefixRequest) GetRole() *NestedRoleRequest {
	return m.Role
}

// SetRole sets the Role property
func (m *PrefixRequest) SetRole(val *NestedRoleRequest) {
	m.Role = val
}

// GetSite returns the Site property
func (m PrefixRequest) GetSite() *NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *PrefixRequest) SetSite(val *NestedSiteRequest) {
	m.Site = val
}

// GetStatus returns the Status property
func (m PrefixRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PrefixRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PrefixRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PrefixRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PrefixRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PrefixRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetVlan returns the Vlan property
func (m PrefixRequest) GetVlan() *NestedVLANRequest {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *PrefixRequest) SetVlan(val *NestedVLANRequest) {
	m.Vlan = val
}

// GetVrf returns the Vrf property
func (m PrefixRequest) GetVrf() *NestedVRFRequest {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *PrefixRequest) SetVrf(val *NestedVRFRequest) {
	m.Vrf = val
}
