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

// siteSlugPattern is the validation pattern for Site.Slug
var siteSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// Site is an object. Adds support for custom fields and tags.
type Site struct {
	// Asns:
	Asns []int32 `json:"asns,omitempty" mapstructure:"asns,omitempty"`
	// CircuitCount:
	CircuitCount int32 `json:"circuit_count" mapstructure:"circuit_count"`
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
	// Facility: Local facility ID or description
	Facility string `json:"facility,omitempty" mapstructure:"facility,omitempty"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedSiteGroup `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Latitude: GPS coordinate in decimal format (xx.yyyyyy)
	Latitude *float64 `json:"latitude,omitempty" mapstructure:"latitude,omitempty"`
	// Longitude: GPS coordinate in decimal format (xx.yyyyyy)
	Longitude *float64 `json:"longitude,omitempty" mapstructure:"longitude,omitempty"`
	// Name: Full name of the site
	Name string `json:"name" mapstructure:"name"`
	// PhysicalAddress: Physical location of the building
	PhysicalAddress string `json:"physical_address,omitempty" mapstructure:"physical_address,omitempty"`
	// PrefixCount:
	PrefixCount int32 `json:"prefix_count" mapstructure:"prefix_count"`
	// RackCount:
	RackCount int32 `json:"rack_count" mapstructure:"rack_count"`
	// Region: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Region *NestedRegion `json:"region,omitempty" mapstructure:"region,omitempty"`
	// ShippingAddress: If different from the physical address
	ShippingAddress string `json:"shipping_address,omitempty" mapstructure:"shipping_address,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
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
	// TimeZone:
	TimeZone *string `json:"time_zone,omitempty" mapstructure:"time_zone,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VirtualmachineCount:
	VirtualmachineCount int32 `json:"virtualmachine_count" mapstructure:"virtualmachine_count"`
	// VlanCount:
	VlanCount int32 `json:"vlan_count" mapstructure:"vlan_count"`
}

// Validate implements basic validation for this model
func (m Site) Validate() error {
	return validation.Errors{
		"asns": validation.Validate(
			m.Asns,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"facility": validation.Validate(
			m.Facility, validation.Length(0, 50),
		),
		"group": validation.Validate(
			m.Group,
		),
		"latitude": validation.Validate(
			m.Latitude, validation.Min(*float64(-100)), validation.Max(*float64(100)),
		),
		"longitude": validation.Validate(
			m.Longitude, validation.Min(*float64(-1000)), validation.Max(*float64(1000)),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"physicalAddress": validation.Validate(
			m.PhysicalAddress, validation.Length(0, 200),
		),
		"region": validation.Validate(
			m.Region,
		),
		"shippingAddress": validation.Validate(
			m.ShippingAddress, validation.Length(0, 200),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(siteSlugPattern),
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
	}.Filter()
}

// GetAsns returns the Asns property
func (m Site) GetAsns() []int32 {
	return m.Asns
}

// SetAsns sets the Asns property
func (m *Site) SetAsns(val []int32) {
	m.Asns = val
}

// GetCircuitCount returns the CircuitCount property
func (m Site) GetCircuitCount() int32 {
	return m.CircuitCount
}

// SetCircuitCount sets the CircuitCount property
func (m *Site) SetCircuitCount(val int32) {
	m.CircuitCount = val
}

// GetComments returns the Comments property
func (m Site) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Site) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Site) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Site) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Site) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Site) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Site) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Site) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m Site) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *Site) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m Site) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Site) SetDisplay(val string) {
	m.Display = val
}

// GetFacility returns the Facility property
func (m Site) GetFacility() string {
	return m.Facility
}

// SetFacility sets the Facility property
func (m *Site) SetFacility(val string) {
	m.Facility = val
}

// GetGroup returns the Group property
func (m Site) GetGroup() *NestedSiteGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *Site) SetGroup(val *NestedSiteGroup) {
	m.Group = val
}

// GetId returns the Id property
func (m Site) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Site) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Site) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Site) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLatitude returns the Latitude property
func (m Site) GetLatitude() *float64 {
	return m.Latitude
}

// SetLatitude sets the Latitude property
func (m *Site) SetLatitude(val *float64) {
	m.Latitude = val
}

// GetLongitude returns the Longitude property
func (m Site) GetLongitude() *float64 {
	return m.Longitude
}

// SetLongitude sets the Longitude property
func (m *Site) SetLongitude(val *float64) {
	m.Longitude = val
}

// GetName returns the Name property
func (m Site) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Site) SetName(val string) {
	m.Name = val
}

// GetPhysicalAddress returns the PhysicalAddress property
func (m Site) GetPhysicalAddress() string {
	return m.PhysicalAddress
}

// SetPhysicalAddress sets the PhysicalAddress property
func (m *Site) SetPhysicalAddress(val string) {
	m.PhysicalAddress = val
}

// GetPrefixCount returns the PrefixCount property
func (m Site) GetPrefixCount() int32 {
	return m.PrefixCount
}

// SetPrefixCount sets the PrefixCount property
func (m *Site) SetPrefixCount(val int32) {
	m.PrefixCount = val
}

// GetRackCount returns the RackCount property
func (m Site) GetRackCount() int32 {
	return m.RackCount
}

// SetRackCount sets the RackCount property
func (m *Site) SetRackCount(val int32) {
	m.RackCount = val
}

// GetRegion returns the Region property
func (m Site) GetRegion() *NestedRegion {
	return m.Region
}

// SetRegion sets the Region property
func (m *Site) SetRegion(val *NestedRegion) {
	m.Region = val
}

// GetShippingAddress returns the ShippingAddress property
func (m Site) GetShippingAddress() string {
	return m.ShippingAddress
}

// SetShippingAddress sets the ShippingAddress property
func (m *Site) SetShippingAddress(val string) {
	m.ShippingAddress = val
}

// GetSlug returns the Slug property
func (m Site) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Site) SetSlug(val string) {
	m.Slug = val
}

// GetStatus returns the Status property
func (m Site) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Site) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Site) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Site) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Site) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Site) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetTimeZone returns the TimeZone property
func (m Site) GetTimeZone() *string {
	return m.TimeZone
}

// SetTimeZone sets the TimeZone property
func (m *Site) SetTimeZone(val *string) {
	m.TimeZone = val
}

// GetUrl returns the Url property
func (m Site) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Site) SetUrl(val string) {
	m.Url = val
}

// GetVirtualmachineCount returns the VirtualmachineCount property
func (m Site) GetVirtualmachineCount() int32 {
	return m.VirtualmachineCount
}

// SetVirtualmachineCount sets the VirtualmachineCount property
func (m *Site) SetVirtualmachineCount(val int32) {
	m.VirtualmachineCount = val
}

// GetVlanCount returns the VlanCount property
func (m Site) GetVlanCount() int32 {
	return m.VlanCount
}

// SetVlanCount sets the VlanCount property
func (m *Site) SetVlanCount(val int32) {
	m.VlanCount = val
}
