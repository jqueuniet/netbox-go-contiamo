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

// locationSlugPattern is the validation pattern for Location.Slug
var locationSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// Location is an object. Extends PrimaryModelSerializer to include MPTT support.
type Location struct {
	// Depth:
	Depth int32 `json:"_depth" mapstructure:"_depth"`
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
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedLocation `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// RackCount:
	RackCount int32 `json:"rack_count" mapstructure:"rack_count"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site NestedSite `json:"site" mapstructure:"site"`
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
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Location) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"parent": validation.Validate(
			m.Parent,
		),
		"site": validation.Validate(
			m.Site, validation.NotNil,
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(locationSlugPattern),
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

// GetDepth returns the Depth property
func (m Location) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *Location) SetDepth(val int32) {
	m.Depth = val
}

// GetCreated returns the Created property
func (m Location) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Location) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Location) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Location) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Location) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Location) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m Location) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *Location) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m Location) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Location) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Location) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Location) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Location) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Location) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Location) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Location) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m Location) GetParent() *NestedLocation {
	return m.Parent
}

// SetParent sets the Parent property
func (m *Location) SetParent(val *NestedLocation) {
	m.Parent = val
}

// GetRackCount returns the RackCount property
func (m Location) GetRackCount() int32 {
	return m.RackCount
}

// SetRackCount sets the RackCount property
func (m *Location) SetRackCount(val int32) {
	m.RackCount = val
}

// GetSite returns the Site property
func (m Location) GetSite() NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *Location) SetSite(val NestedSite) {
	m.Site = val
}

// GetSlug returns the Slug property
func (m Location) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Location) SetSlug(val string) {
	m.Slug = val
}

// GetStatus returns the Status property
func (m Location) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Location) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Location) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Location) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Location) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Location) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m Location) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Location) SetUrl(val string) {
	m.Url = val
}
