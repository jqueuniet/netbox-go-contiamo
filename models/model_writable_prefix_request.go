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

// WritablePrefixRequest is an object. Adds support for custom fields and tags.
type WritablePrefixRequest struct {
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
	// Role: The primary function of this prefix
	Role *int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site:
	Site *int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: Operational status of this prefix
	//
	// * `container` - Container
	// * `active` - Active
	// * `reserved` - Reserved
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vlan:
	Vlan *int32 `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
	// Vrf:
	Vrf *int32 `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m WritablePrefixRequest) Validate() error {
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
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritablePrefixRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritablePrefixRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritablePrefixRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritablePrefixRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritablePrefixRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritablePrefixRequest) SetDescription(val string) {
	m.Description = val
}

// GetIsPool returns the IsPool property
func (m WritablePrefixRequest) GetIsPool() bool {
	return m.IsPool
}

// SetIsPool sets the IsPool property
func (m *WritablePrefixRequest) SetIsPool(val bool) {
	m.IsPool = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m WritablePrefixRequest) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *WritablePrefixRequest) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetPrefix returns the Prefix property
func (m WritablePrefixRequest) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *WritablePrefixRequest) SetPrefix(val string) {
	m.Prefix = val
}

// GetRole returns the Role property
func (m WritablePrefixRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *WritablePrefixRequest) SetRole(val *int32) {
	m.Role = val
}

// GetSite returns the Site property
func (m WritablePrefixRequest) GetSite() *int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *WritablePrefixRequest) SetSite(val *int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m WritablePrefixRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WritablePrefixRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WritablePrefixRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritablePrefixRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritablePrefixRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritablePrefixRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVlan returns the Vlan property
func (m WritablePrefixRequest) GetVlan() *int32 {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *WritablePrefixRequest) SetVlan(val *int32) {
	m.Vlan = val
}

// GetVrf returns the Vrf property
func (m WritablePrefixRequest) GetVrf() *int32 {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *WritablePrefixRequest) SetVrf(val *int32) {
	m.Vrf = val
}
