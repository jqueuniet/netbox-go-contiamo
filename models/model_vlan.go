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

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// VLAN is an object. Adds support for custom fields and tags.
type VLAN struct {
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
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedVLANGroup `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// L2vpnTermination: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	L2vpnTermination *NestedL2VPNTermination `json:"l2vpn_termination" mapstructure:"l2vpn_termination"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PrefixCount:
	PrefixCount int32 `json:"prefix_count" mapstructure:"prefix_count"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedRole `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSite `json:"site,omitempty" mapstructure:"site,omitempty"`
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
	// Vid: Numeric VLAN ID (1-4094)
	Vid int32 `json:"vid" mapstructure:"vid"`
}

// Validate implements basic validation for this model
func (m VLAN) Validate() error {
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
		"l2vpnTermination": validation.Validate(
			m.L2vpnTermination,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"vid": validation.Validate(
			m.Vid, validation.Required, validation.Min(int32(1)), validation.Max(int32(4094)),
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m VLAN) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VLAN) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m VLAN) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VLAN) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m VLAN) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VLAN) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VLAN) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VLAN) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m VLAN) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *VLAN) SetDisplay(val string) {
	m.Display = val
}

// GetGroup returns the Group property
func (m VLAN) GetGroup() *NestedVLANGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *VLAN) SetGroup(val *NestedVLANGroup) {
	m.Group = val
}

// GetId returns the Id property
func (m VLAN) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VLAN) SetId(val int32) {
	m.Id = val
}

// GetL2vpnTermination returns the L2vpnTermination property
func (m VLAN) GetL2vpnTermination() *NestedL2VPNTermination {
	return m.L2vpnTermination
}

// SetL2vpnTermination sets the L2vpnTermination property
func (m *VLAN) SetL2vpnTermination(val *NestedL2VPNTermination) {
	m.L2vpnTermination = val
}

// GetLastUpdated returns the LastUpdated property
func (m VLAN) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VLAN) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m VLAN) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VLAN) SetName(val string) {
	m.Name = val
}

// GetPrefixCount returns the PrefixCount property
func (m VLAN) GetPrefixCount() int32 {
	return m.PrefixCount
}

// SetPrefixCount sets the PrefixCount property
func (m *VLAN) SetPrefixCount(val int32) {
	m.PrefixCount = val
}

// GetRole returns the Role property
func (m VLAN) GetRole() *NestedRole {
	return m.Role
}

// SetRole sets the Role property
func (m *VLAN) SetRole(val *NestedRole) {
	m.Role = val
}

// GetSite returns the Site property
func (m VLAN) GetSite() *NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *VLAN) SetSite(val *NestedSite) {
	m.Site = val
}

// GetStatus returns the Status property
func (m VLAN) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *VLAN) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m VLAN) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VLAN) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VLAN) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VLAN) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m VLAN) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *VLAN) SetUrl(val string) {
	m.Url = val
}

// GetVid returns the Vid property
func (m VLAN) GetVid() int32 {
	return m.Vid
}

// SetVid sets the Vid property
func (m *VLAN) SetVid(val int32) {
	m.Vid = val
}
