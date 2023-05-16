// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// iPAddressDnsNamePattern is the validation pattern for IPAddress.DnsName
var iPAddressDnsNamePattern = regexp.MustCompile(`^([0-9A-Za-z_-]+|\*)(\.[0-9A-Za-z_-]+)*\.?$`)

// IPAddress is an object. Adds support for custom fields and tags.
type IPAddress struct {
	// Address:
	Address string `json:"address" mapstructure:"address"`
	// AssignedObject:
	AssignedObject map[string]interface{} `json:"assigned_object" mapstructure:"assigned_object"`
	// AssignedObjectId:
	AssignedObjectId *int64 `json:"assigned_object_id,omitempty" mapstructure:"assigned_object_id,omitempty"`
	// AssignedObjectType:
	AssignedObjectType *string `json:"assigned_object_type,omitempty" mapstructure:"assigned_object_type,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// DnsName: Hostname or FQDN (not case-sensitive)
	DnsName string `json:"dns_name,omitempty" mapstructure:"dns_name,omitempty"`
	// Family:
	Family struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"family" mapstructure:"family"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// NatInside: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	NatInside *NestedIPAddress `json:"nat_inside,omitempty" mapstructure:"nat_inside,omitempty"`
	// NatOutside:
	NatOutside []NestedIPAddress `json:"nat_outside" mapstructure:"nat_outside"`
	// Role:
	Role *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Status:
	Status *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRF `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m IPAddress) Validate() error {
	return validation.Errors{
		"assignedObject": validation.Validate(
			m.AssignedObject,
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
			m.DnsName, validation.Length(0, 255), validation.Match(iPAddressDnsNamePattern),
		),
		"natInside": validation.Validate(
			m.NatInside,
		),
		"natOutside": validation.Validate(
			m.NatOutside, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m IPAddress) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *IPAddress) SetAddress(val string) {
	m.Address = val
}

// GetAssignedObject returns the AssignedObject property
func (m IPAddress) GetAssignedObject() map[string]interface{} {
	return m.AssignedObject
}

// SetAssignedObject sets the AssignedObject property
func (m *IPAddress) SetAssignedObject(val map[string]interface{}) {
	m.AssignedObject = val
}

// GetAssignedObjectId returns the AssignedObjectId property
func (m IPAddress) GetAssignedObjectId() *int64 {
	return m.AssignedObjectId
}

// SetAssignedObjectId sets the AssignedObjectId property
func (m *IPAddress) SetAssignedObjectId(val *int64) {
	m.AssignedObjectId = val
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m IPAddress) GetAssignedObjectType() *string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *IPAddress) SetAssignedObjectType(val *string) {
	m.AssignedObjectType = val
}

// GetComments returns the Comments property
func (m IPAddress) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *IPAddress) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m IPAddress) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IPAddress) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m IPAddress) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *IPAddress) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m IPAddress) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IPAddress) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m IPAddress) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *IPAddress) SetDisplay(val string) {
	m.Display = val
}

// GetDnsName returns the DnsName property
func (m IPAddress) GetDnsName() string {
	return m.DnsName
}

// SetDnsName sets the DnsName property
func (m *IPAddress) SetDnsName(val string) {
	m.DnsName = val
}

// GetFamily returns the Family property
func (m IPAddress) GetFamily() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Family
}

// SetFamily sets the Family property
func (m *IPAddress) SetFamily(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Family = val
}

// GetId returns the Id property
func (m IPAddress) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IPAddress) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m IPAddress) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IPAddress) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetNatInside returns the NatInside property
func (m IPAddress) GetNatInside() *NestedIPAddress {
	return m.NatInside
}

// SetNatInside sets the NatInside property
func (m *IPAddress) SetNatInside(val *NestedIPAddress) {
	m.NatInside = val
}

// GetNatOutside returns the NatOutside property
func (m IPAddress) GetNatOutside() []NestedIPAddress {
	return m.NatOutside
}

// SetNatOutside sets the NatOutside property
func (m *IPAddress) SetNatOutside(val []NestedIPAddress) {
	m.NatOutside = val
}

// GetRole returns the Role property
func (m IPAddress) GetRole() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Role
}

// SetRole sets the Role property
func (m *IPAddress) SetRole(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Role = val
}

// GetStatus returns the Status property
func (m IPAddress) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *IPAddress) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m IPAddress) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *IPAddress) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m IPAddress) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IPAddress) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m IPAddress) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *IPAddress) SetUrl(val string) {
	m.Url = val
}

// GetVrf returns the Vrf property
func (m IPAddress) GetVrf() *NestedVRF {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *IPAddress) SetVrf(val *NestedVRF) {
	m.Vrf = val
}
