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

// PatchedWritableIPRangeRequest is an object. Adds support for custom fields and tags.
type PatchedWritableIPRangeRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// EndAddress:
	EndAddress string `json:"end_address,omitempty" mapstructure:"end_address,omitempty"`
	// MarkUtilized: Treat as 100% utilized
	MarkUtilized bool `json:"mark_utilized,omitempty" mapstructure:"mark_utilized,omitempty"`
	// Role: The primary function of this range
	Role *int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
	// StartAddress:
	StartAddress string `json:"start_address,omitempty" mapstructure:"start_address,omitempty"`
	// Status: Operational status of this range
	//
	// * `active` - Active
	// * `reserved` - Reserved
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vrf:
	Vrf *int32 `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableIPRangeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"endAddress": validation.Validate(
			m.EndAddress, validation.Length(1, 0),
		),
		"startAddress": validation.Validate(
			m.StartAddress, validation.Length(1, 0),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableIPRangeRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableIPRangeRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableIPRangeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableIPRangeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableIPRangeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableIPRangeRequest) SetDescription(val string) {
	m.Description = val
}

// GetEndAddress returns the EndAddress property
func (m PatchedWritableIPRangeRequest) GetEndAddress() string {
	return m.EndAddress
}

// SetEndAddress sets the EndAddress property
func (m *PatchedWritableIPRangeRequest) SetEndAddress(val string) {
	m.EndAddress = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m PatchedWritableIPRangeRequest) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *PatchedWritableIPRangeRequest) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetRole returns the Role property
func (m PatchedWritableIPRangeRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *PatchedWritableIPRangeRequest) SetRole(val *int32) {
	m.Role = val
}

// GetStartAddress returns the StartAddress property
func (m PatchedWritableIPRangeRequest) GetStartAddress() string {
	return m.StartAddress
}

// SetStartAddress sets the StartAddress property
func (m *PatchedWritableIPRangeRequest) SetStartAddress(val string) {
	m.StartAddress = val
}

// GetStatus returns the Status property
func (m PatchedWritableIPRangeRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableIPRangeRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableIPRangeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableIPRangeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableIPRangeRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableIPRangeRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVrf returns the Vrf property
func (m PatchedWritableIPRangeRequest) GetVrf() *int32 {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *PatchedWritableIPRangeRequest) SetVrf(val *int32) {
	m.Vrf = val
}
