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

// Prefix is an object. Adds support for custom fields and tags.
type Prefix struct {
	// Depth:
	Depth int32 `json:"_depth" mapstructure:"_depth"`
	// Children:
	Children int32 `json:"children" mapstructure:"children"`
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
	// Family:
	Family struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"family" mapstructure:"family"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IsPool: All IP addresses within this prefix are considered usable
	IsPool bool `json:"is_pool,omitempty" mapstructure:"is_pool,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// MarkUtilized: Treat as 100% utilized
	MarkUtilized bool `json:"mark_utilized,omitempty" mapstructure:"mark_utilized,omitempty"`
	// Prefix:
	Prefix string `json:"prefix" mapstructure:"prefix"`
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
	// Vlan: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vlan *NestedVLAN `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRF `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m Prefix) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
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
		"vlan": validation.Validate(
			m.Vlan,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
	}.Filter()
}

// GetDepth returns the Depth property
func (m Prefix) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *Prefix) SetDepth(val int32) {
	m.Depth = val
}

// GetChildren returns the Children property
func (m Prefix) GetChildren() int32 {
	return m.Children
}

// SetChildren sets the Children property
func (m *Prefix) SetChildren(val int32) {
	m.Children = val
}

// GetComments returns the Comments property
func (m Prefix) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Prefix) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Prefix) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Prefix) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Prefix) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Prefix) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Prefix) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Prefix) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Prefix) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Prefix) SetDisplay(val string) {
	m.Display = val
}

// GetFamily returns the Family property
func (m Prefix) GetFamily() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Family
}

// SetFamily sets the Family property
func (m *Prefix) SetFamily(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Family = val
}

// GetId returns the Id property
func (m Prefix) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Prefix) SetId(val int32) {
	m.Id = val
}

// GetIsPool returns the IsPool property
func (m Prefix) GetIsPool() bool {
	return m.IsPool
}

// SetIsPool sets the IsPool property
func (m *Prefix) SetIsPool(val bool) {
	m.IsPool = val
}

// GetLastUpdated returns the LastUpdated property
func (m Prefix) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Prefix) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m Prefix) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *Prefix) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetPrefix returns the Prefix property
func (m Prefix) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *Prefix) SetPrefix(val string) {
	m.Prefix = val
}

// GetRole returns the Role property
func (m Prefix) GetRole() *NestedRole {
	return m.Role
}

// SetRole sets the Role property
func (m *Prefix) SetRole(val *NestedRole) {
	m.Role = val
}

// GetSite returns the Site property
func (m Prefix) GetSite() *NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *Prefix) SetSite(val *NestedSite) {
	m.Site = val
}

// GetStatus returns the Status property
func (m Prefix) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Prefix) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Prefix) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Prefix) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Prefix) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Prefix) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m Prefix) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Prefix) SetUrl(val string) {
	m.Url = val
}

// GetVlan returns the Vlan property
func (m Prefix) GetVlan() *NestedVLAN {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *Prefix) SetVlan(val *NestedVLAN) {
	m.Vlan = val
}

// GetVrf returns the Vrf property
func (m Prefix) GetVrf() *NestedVRF {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *Prefix) SetVrf(val *NestedVRF) {
	m.Vrf = val
}
