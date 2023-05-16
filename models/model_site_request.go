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

// siteRequestSlugPattern is the validation pattern for SiteRequest.Slug
var siteRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// SiteRequest is an object. Adds support for custom fields and tags.
type SiteRequest struct {
	// Asns:
	Asns []int32 `json:"asns,omitempty" mapstructure:"asns,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Facility: Local facility ID or description
	Facility string `json:"facility,omitempty" mapstructure:"facility,omitempty"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedSiteGroupRequest `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Latitude: GPS coordinate in decimal format (xx.yyyyyy)
	Latitude *float64 `json:"latitude,omitempty" mapstructure:"latitude,omitempty"`
	// Longitude: GPS coordinate in decimal format (xx.yyyyyy)
	Longitude *float64 `json:"longitude,omitempty" mapstructure:"longitude,omitempty"`
	// Name: Full name of the site
	Name string `json:"name" mapstructure:"name"`
	// PhysicalAddress: Physical location of the building
	PhysicalAddress string `json:"physical_address,omitempty" mapstructure:"physical_address,omitempty"`
	// Region: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Region *NestedRegionRequest `json:"region,omitempty" mapstructure:"region,omitempty"`
	// ShippingAddress: If different from the physical address
	ShippingAddress string `json:"shipping_address,omitempty" mapstructure:"shipping_address,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Status: * `planned` - Planned
	// * `staging` - Staging
	// * `active` - Active
	// * `decommissioning` - Decommissioning
	// * `retired` - Retired
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TimeZone:
	TimeZone *string `json:"time_zone,omitempty" mapstructure:"time_zone,omitempty"`
}

// Validate implements basic validation for this model
func (m SiteRequest) Validate() error {
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
			m.Name, validation.Required, validation.Length(1, 100),
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(siteRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
	}.Filter()
}

// GetAsns returns the Asns property
func (m SiteRequest) GetAsns() []int32 {
	return m.Asns
}

// SetAsns sets the Asns property
func (m *SiteRequest) SetAsns(val []int32) {
	m.Asns = val
}

// GetComments returns the Comments property
func (m SiteRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *SiteRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m SiteRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *SiteRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m SiteRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *SiteRequest) SetDescription(val string) {
	m.Description = val
}

// GetFacility returns the Facility property
func (m SiteRequest) GetFacility() string {
	return m.Facility
}

// SetFacility sets the Facility property
func (m *SiteRequest) SetFacility(val string) {
	m.Facility = val
}

// GetGroup returns the Group property
func (m SiteRequest) GetGroup() *NestedSiteGroupRequest {
	return m.Group
}

// SetGroup sets the Group property
func (m *SiteRequest) SetGroup(val *NestedSiteGroupRequest) {
	m.Group = val
}

// GetLatitude returns the Latitude property
func (m SiteRequest) GetLatitude() *float64 {
	return m.Latitude
}

// SetLatitude sets the Latitude property
func (m *SiteRequest) SetLatitude(val *float64) {
	m.Latitude = val
}

// GetLongitude returns the Longitude property
func (m SiteRequest) GetLongitude() *float64 {
	return m.Longitude
}

// SetLongitude sets the Longitude property
func (m *SiteRequest) SetLongitude(val *float64) {
	m.Longitude = val
}

// GetName returns the Name property
func (m SiteRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *SiteRequest) SetName(val string) {
	m.Name = val
}

// GetPhysicalAddress returns the PhysicalAddress property
func (m SiteRequest) GetPhysicalAddress() string {
	return m.PhysicalAddress
}

// SetPhysicalAddress sets the PhysicalAddress property
func (m *SiteRequest) SetPhysicalAddress(val string) {
	m.PhysicalAddress = val
}

// GetRegion returns the Region property
func (m SiteRequest) GetRegion() *NestedRegionRequest {
	return m.Region
}

// SetRegion sets the Region property
func (m *SiteRequest) SetRegion(val *NestedRegionRequest) {
	m.Region = val
}

// GetShippingAddress returns the ShippingAddress property
func (m SiteRequest) GetShippingAddress() string {
	return m.ShippingAddress
}

// SetShippingAddress sets the ShippingAddress property
func (m *SiteRequest) SetShippingAddress(val string) {
	m.ShippingAddress = val
}

// GetSlug returns the Slug property
func (m SiteRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *SiteRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStatus returns the Status property
func (m SiteRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *SiteRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m SiteRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *SiteRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m SiteRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *SiteRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetTimeZone returns the TimeZone property
func (m SiteRequest) GetTimeZone() *string {
	return m.TimeZone
}

// SetTimeZone sets the TimeZone property
func (m *SiteRequest) SetTimeZone(val *string) {
	m.TimeZone = val
}
