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

// locationRequestSlugPattern is the validation pattern for LocationRequest.Slug
var locationRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// LocationRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type LocationRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedLocationRequest `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site NestedSiteRequest `json:"site" mapstructure:"site"`
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
}

// Validate implements basic validation for this model
func (m LocationRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"parent": validation.Validate(
			m.Parent,
		),
		"site": validation.Validate(
			m.Site, validation.NotNil,
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(locationRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m LocationRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *LocationRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m LocationRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *LocationRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m LocationRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *LocationRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m LocationRequest) GetParent() *NestedLocationRequest {
	return m.Parent
}

// SetParent sets the Parent property
func (m *LocationRequest) SetParent(val *NestedLocationRequest) {
	m.Parent = val
}

// GetSite returns the Site property
func (m LocationRequest) GetSite() NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *LocationRequest) SetSite(val NestedSiteRequest) {
	m.Site = val
}

// GetSlug returns the Slug property
func (m LocationRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *LocationRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStatus returns the Status property
func (m LocationRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *LocationRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m LocationRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *LocationRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m LocationRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *LocationRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
