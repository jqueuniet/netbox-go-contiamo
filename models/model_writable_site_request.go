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

// writableSiteRequestSlugPattern is the validation pattern for WritableSiteRequest.Slug
var writableSiteRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WritableSiteRequest is an object. Adds support for custom fields and tags.
type WritableSiteRequest struct {
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
	// Group:
	Group *int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Latitude: GPS coordinate in decimal format (xx.yyyyyy)
	Latitude *float64 `json:"latitude,omitempty" mapstructure:"latitude,omitempty"`
	// Longitude: GPS coordinate in decimal format (xx.yyyyyy)
	Longitude *float64 `json:"longitude,omitempty" mapstructure:"longitude,omitempty"`
	// Name: Full name of the site
	Name string `json:"name" mapstructure:"name"`
	// PhysicalAddress: Physical location of the building
	PhysicalAddress string `json:"physical_address,omitempty" mapstructure:"physical_address,omitempty"`
	// Region:
	Region *int32 `json:"region,omitempty" mapstructure:"region,omitempty"`
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
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TimeZone:
	TimeZone *string `json:"time_zone,omitempty" mapstructure:"time_zone,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableSiteRequest) Validate() error {
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
		"shippingAddress": validation.Validate(
			m.ShippingAddress, validation.Length(0, 200),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(writableSiteRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAsns returns the Asns property
func (m WritableSiteRequest) GetAsns() []int32 {
	return m.Asns
}

// SetAsns sets the Asns property
func (m *WritableSiteRequest) SetAsns(val []int32) {
	m.Asns = val
}

// GetComments returns the Comments property
func (m WritableSiteRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableSiteRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableSiteRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableSiteRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableSiteRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableSiteRequest) SetDescription(val string) {
	m.Description = val
}

// GetFacility returns the Facility property
func (m WritableSiteRequest) GetFacility() string {
	return m.Facility
}

// SetFacility sets the Facility property
func (m *WritableSiteRequest) SetFacility(val string) {
	m.Facility = val
}

// GetGroup returns the Group property
func (m WritableSiteRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *WritableSiteRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetLatitude returns the Latitude property
func (m WritableSiteRequest) GetLatitude() *float64 {
	return m.Latitude
}

// SetLatitude sets the Latitude property
func (m *WritableSiteRequest) SetLatitude(val *float64) {
	m.Latitude = val
}

// GetLongitude returns the Longitude property
func (m WritableSiteRequest) GetLongitude() *float64 {
	return m.Longitude
}

// SetLongitude sets the Longitude property
func (m *WritableSiteRequest) SetLongitude(val *float64) {
	m.Longitude = val
}

// GetName returns the Name property
func (m WritableSiteRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableSiteRequest) SetName(val string) {
	m.Name = val
}

// GetPhysicalAddress returns the PhysicalAddress property
func (m WritableSiteRequest) GetPhysicalAddress() string {
	return m.PhysicalAddress
}

// SetPhysicalAddress sets the PhysicalAddress property
func (m *WritableSiteRequest) SetPhysicalAddress(val string) {
	m.PhysicalAddress = val
}

// GetRegion returns the Region property
func (m WritableSiteRequest) GetRegion() *int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *WritableSiteRequest) SetRegion(val *int32) {
	m.Region = val
}

// GetShippingAddress returns the ShippingAddress property
func (m WritableSiteRequest) GetShippingAddress() string {
	return m.ShippingAddress
}

// SetShippingAddress sets the ShippingAddress property
func (m *WritableSiteRequest) SetShippingAddress(val string) {
	m.ShippingAddress = val
}

// GetSlug returns the Slug property
func (m WritableSiteRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WritableSiteRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStatus returns the Status property
func (m WritableSiteRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WritableSiteRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WritableSiteRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableSiteRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritableSiteRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritableSiteRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetTimeZone returns the TimeZone property
func (m WritableSiteRequest) GetTimeZone() *string {
	return m.TimeZone
}

// SetTimeZone sets the TimeZone property
func (m *WritableSiteRequest) SetTimeZone(val *string) {
	m.TimeZone = val
}
