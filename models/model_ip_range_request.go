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

// IPRangeRequest is an object. Adds support for custom fields and tags.
type IPRangeRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// EndAddress:
	EndAddress string `json:"end_address" mapstructure:"end_address"`
	// MarkUtilized: Treat as 100% utilized
	MarkUtilized bool `json:"mark_utilized,omitempty" mapstructure:"mark_utilized,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedRoleRequest `json:"role,omitempty" mapstructure:"role,omitempty"`
	// StartAddress:
	StartAddress string `json:"start_address" mapstructure:"start_address"`
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
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRFRequest `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m IPRangeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"endAddress": validation.Validate(
			m.EndAddress, validation.Required, validation.Length(1, 0),
		),
		"role": validation.Validate(
			m.Role,
		),
		"startAddress": validation.Validate(
			m.StartAddress, validation.Required, validation.Length(1, 0),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m IPRangeRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *IPRangeRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m IPRangeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *IPRangeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m IPRangeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IPRangeRequest) SetDescription(val string) {
	m.Description = val
}

// GetEndAddress returns the EndAddress property
func (m IPRangeRequest) GetEndAddress() string {
	return m.EndAddress
}

// SetEndAddress sets the EndAddress property
func (m *IPRangeRequest) SetEndAddress(val string) {
	m.EndAddress = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m IPRangeRequest) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *IPRangeRequest) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetRole returns the Role property
func (m IPRangeRequest) GetRole() *NestedRoleRequest {
	return m.Role
}

// SetRole sets the Role property
func (m *IPRangeRequest) SetRole(val *NestedRoleRequest) {
	m.Role = val
}

// GetStartAddress returns the StartAddress property
func (m IPRangeRequest) GetStartAddress() string {
	return m.StartAddress
}

// SetStartAddress sets the StartAddress property
func (m *IPRangeRequest) SetStartAddress(val string) {
	m.StartAddress = val
}

// GetStatus returns the Status property
func (m IPRangeRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *IPRangeRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m IPRangeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *IPRangeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m IPRangeRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IPRangeRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetVrf returns the Vrf property
func (m IPRangeRequest) GetVrf() *NestedVRFRequest {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *IPRangeRequest) SetVrf(val *NestedVRFRequest) {
	m.Vrf = val
}
