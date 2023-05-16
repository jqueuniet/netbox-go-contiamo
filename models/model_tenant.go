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

// tenantSlugPattern is the validation pattern for Tenant.Slug
var tenantSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// Tenant is an object. Adds support for custom fields and tags.
type Tenant struct {
	// CircuitCount:
	CircuitCount int32 `json:"circuit_count" mapstructure:"circuit_count"`
	// ClusterCount:
	ClusterCount int32 `json:"cluster_count" mapstructure:"cluster_count"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceCount:
	DeviceCount int32 `json:"device_count" mapstructure:"device_count"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedTenantGroup `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IpaddressCount:
	IpaddressCount int32 `json:"ipaddress_count" mapstructure:"ipaddress_count"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PrefixCount:
	PrefixCount int32 `json:"prefix_count" mapstructure:"prefix_count"`
	// RackCount:
	RackCount int32 `json:"rack_count" mapstructure:"rack_count"`
	// SiteCount:
	SiteCount int32 `json:"site_count" mapstructure:"site_count"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VirtualmachineCount:
	VirtualmachineCount int32 `json:"virtualmachine_count" mapstructure:"virtualmachine_count"`
	// VlanCount:
	VlanCount int32 `json:"vlan_count" mapstructure:"vlan_count"`
	// VrfCount:
	VrfCount int32 `json:"vrf_count" mapstructure:"vrf_count"`
}

// Validate implements basic validation for this model
func (m Tenant) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(tenantSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCircuitCount returns the CircuitCount property
func (m Tenant) GetCircuitCount() int32 {
	return m.CircuitCount
}

// SetCircuitCount sets the CircuitCount property
func (m *Tenant) SetCircuitCount(val int32) {
	m.CircuitCount = val
}

// GetClusterCount returns the ClusterCount property
func (m Tenant) GetClusterCount() int32 {
	return m.ClusterCount
}

// SetClusterCount sets the ClusterCount property
func (m *Tenant) SetClusterCount(val int32) {
	m.ClusterCount = val
}

// GetComments returns the Comments property
func (m Tenant) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Tenant) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Tenant) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Tenant) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Tenant) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Tenant) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Tenant) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Tenant) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m Tenant) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *Tenant) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m Tenant) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Tenant) SetDisplay(val string) {
	m.Display = val
}

// GetGroup returns the Group property
func (m Tenant) GetGroup() *NestedTenantGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *Tenant) SetGroup(val *NestedTenantGroup) {
	m.Group = val
}

// GetId returns the Id property
func (m Tenant) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Tenant) SetId(val int32) {
	m.Id = val
}

// GetIpaddressCount returns the IpaddressCount property
func (m Tenant) GetIpaddressCount() int32 {
	return m.IpaddressCount
}

// SetIpaddressCount sets the IpaddressCount property
func (m *Tenant) SetIpaddressCount(val int32) {
	m.IpaddressCount = val
}

// GetLastUpdated returns the LastUpdated property
func (m Tenant) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Tenant) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Tenant) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Tenant) SetName(val string) {
	m.Name = val
}

// GetPrefixCount returns the PrefixCount property
func (m Tenant) GetPrefixCount() int32 {
	return m.PrefixCount
}

// SetPrefixCount sets the PrefixCount property
func (m *Tenant) SetPrefixCount(val int32) {
	m.PrefixCount = val
}

// GetRackCount returns the RackCount property
func (m Tenant) GetRackCount() int32 {
	return m.RackCount
}

// SetRackCount sets the RackCount property
func (m *Tenant) SetRackCount(val int32) {
	m.RackCount = val
}

// GetSiteCount returns the SiteCount property
func (m Tenant) GetSiteCount() int32 {
	return m.SiteCount
}

// SetSiteCount sets the SiteCount property
func (m *Tenant) SetSiteCount(val int32) {
	m.SiteCount = val
}

// GetSlug returns the Slug property
func (m Tenant) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Tenant) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m Tenant) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Tenant) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m Tenant) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Tenant) SetUrl(val string) {
	m.Url = val
}

// GetVirtualmachineCount returns the VirtualmachineCount property
func (m Tenant) GetVirtualmachineCount() int32 {
	return m.VirtualmachineCount
}

// SetVirtualmachineCount sets the VirtualmachineCount property
func (m *Tenant) SetVirtualmachineCount(val int32) {
	m.VirtualmachineCount = val
}

// GetVlanCount returns the VlanCount property
func (m Tenant) GetVlanCount() int32 {
	return m.VlanCount
}

// SetVlanCount sets the VlanCount property
func (m *Tenant) SetVlanCount(val int32) {
	m.VlanCount = val
}

// GetVrfCount returns the VrfCount property
func (m Tenant) GetVrfCount() int32 {
	return m.VrfCount
}

// SetVrfCount sets the VrfCount property
func (m *Tenant) SetVrfCount(val int32) {
	m.VrfCount = val
}
