// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// patchedWritableIPAddressRequestDnsNamePattern is the validation pattern for PatchedWritableIPAddressRequest.DnsName
var patchedWritableIPAddressRequestDnsNamePattern = regexp.MustCompile(`^([0-9A-Za-z_-]+|\*)(\.[0-9A-Za-z_-]+)*\.?$`)

// PatchedWritableIPAddressRequest is an object. Adds support for custom fields and tags.
type PatchedWritableIPAddressRequest struct {
	// Address:
	Address string `json:"address,omitempty" mapstructure:"address,omitempty"`
	// AssignedObjectId:
	AssignedObjectId *int64 `json:"assigned_object_id,omitempty" mapstructure:"assigned_object_id,omitempty"`
	// AssignedObjectType:
	AssignedObjectType *string `json:"assigned_object_type,omitempty" mapstructure:"assigned_object_type,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DnsName: Hostname or FQDN (not case-sensitive)
	DnsName string `json:"dns_name,omitempty" mapstructure:"dns_name,omitempty"`
	// NatInside: The IP for which this address is the "outside" IP
	NatInside *int32 `json:"nat_inside,omitempty" mapstructure:"nat_inside,omitempty"`
	// Role: The functional role of this IP
	//
	// * `loopback` - Loopback
	// * `secondary` - Secondary
	// * `anycast` - Anycast
	// * `vip` - VIP
	// * `vrrp` - VRRP
	// * `hsrp` - HSRP
	// * `glbp` - GLBP
	// * `carp` - CARP
	Role string `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Status: The operational status of this IP
	//
	// * `active` - Active
	// * `reserved` - Reserved
	// * `deprecated` - Deprecated
	// * `dhcp` - DHCP
	// * `slaac` - SLAAC
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vrf:
	Vrf *int32 `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableIPAddressRequest) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address, validation.Length(1, 0),
		),
		"assignedObjectId": validation.Validate(
			m.AssignedObjectId, validation.Min(*int64(0)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"dnsName": validation.Validate(
			m.DnsName, validation.Length(0, 255), validation.Match(patchedWritableIPAddressRequestDnsNamePattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m PatchedWritableIPAddressRequest) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *PatchedWritableIPAddressRequest) SetAddress(val string) {
	m.Address = val
}

// GetAssignedObjectId returns the AssignedObjectId property
func (m PatchedWritableIPAddressRequest) GetAssignedObjectId() *int64 {
	return m.AssignedObjectId
}

// SetAssignedObjectId sets the AssignedObjectId property
func (m *PatchedWritableIPAddressRequest) SetAssignedObjectId(val *int64) {
	m.AssignedObjectId = val
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m PatchedWritableIPAddressRequest) GetAssignedObjectType() *string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *PatchedWritableIPAddressRequest) SetAssignedObjectType(val *string) {
	m.AssignedObjectType = val
}

// GetComments returns the Comments property
func (m PatchedWritableIPAddressRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableIPAddressRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableIPAddressRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableIPAddressRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableIPAddressRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableIPAddressRequest) SetDescription(val string) {
	m.Description = val
}

// GetDnsName returns the DnsName property
func (m PatchedWritableIPAddressRequest) GetDnsName() string {
	return m.DnsName
}

// SetDnsName sets the DnsName property
func (m *PatchedWritableIPAddressRequest) SetDnsName(val string) {
	m.DnsName = val
}

// GetNatInside returns the NatInside property
func (m PatchedWritableIPAddressRequest) GetNatInside() *int32 {
	return m.NatInside
}

// SetNatInside sets the NatInside property
func (m *PatchedWritableIPAddressRequest) SetNatInside(val *int32) {
	m.NatInside = val
}

// GetRole returns the Role property
func (m PatchedWritableIPAddressRequest) GetRole() string {
	return m.Role
}

// SetRole sets the Role property
func (m *PatchedWritableIPAddressRequest) SetRole(val string) {
	m.Role = val
}

// GetStatus returns the Status property
func (m PatchedWritableIPAddressRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableIPAddressRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableIPAddressRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableIPAddressRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableIPAddressRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableIPAddressRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVrf returns the Vrf property
func (m PatchedWritableIPAddressRequest) GetVrf() *int32 {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *PatchedWritableIPAddressRequest) SetVrf(val *int32) {
	m.Vrf = val
}
