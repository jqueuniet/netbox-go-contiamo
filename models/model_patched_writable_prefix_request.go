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

// PatchedWritablePrefixRequest is an object. Adds support for custom fields and tags.
type PatchedWritablePrefixRequest struct {
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
	Prefix string `json:"prefix,omitempty" mapstructure:"prefix,omitempty"`
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
func (m PatchedWritablePrefixRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"prefix": validation.Validate(
			m.Prefix, validation.Length(1, 0),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritablePrefixRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritablePrefixRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritablePrefixRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritablePrefixRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritablePrefixRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritablePrefixRequest) SetDescription(val string) {
	m.Description = val
}

// GetIsPool returns the IsPool property
func (m PatchedWritablePrefixRequest) GetIsPool() bool {
	return m.IsPool
}

// SetIsPool sets the IsPool property
func (m *PatchedWritablePrefixRequest) SetIsPool(val bool) {
	m.IsPool = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m PatchedWritablePrefixRequest) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *PatchedWritablePrefixRequest) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetPrefix returns the Prefix property
func (m PatchedWritablePrefixRequest) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *PatchedWritablePrefixRequest) SetPrefix(val string) {
	m.Prefix = val
}

// GetRole returns the Role property
func (m PatchedWritablePrefixRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *PatchedWritablePrefixRequest) SetRole(val *int32) {
	m.Role = val
}

// GetSite returns the Site property
func (m PatchedWritablePrefixRequest) GetSite() *int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *PatchedWritablePrefixRequest) SetSite(val *int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m PatchedWritablePrefixRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritablePrefixRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritablePrefixRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritablePrefixRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritablePrefixRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritablePrefixRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVlan returns the Vlan property
func (m PatchedWritablePrefixRequest) GetVlan() *int32 {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *PatchedWritablePrefixRequest) SetVlan(val *int32) {
	m.Vlan = val
}

// GetVrf returns the Vrf property
func (m PatchedWritablePrefixRequest) GetVrf() *int32 {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *PatchedWritablePrefixRequest) SetVrf(val *int32) {
	m.Vrf = val
}
