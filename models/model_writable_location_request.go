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

// writableLocationRequestSlugPattern is the validation pattern for WritableLocationRequest.Slug
var writableLocationRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WritableLocationRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type WritableLocationRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Site:
	Site int32 `json:"site" mapstructure:"site"`
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
}

// Validate implements basic validation for this model
func (m WritableLocationRequest) Validate() error {
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
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(writableLocationRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m WritableLocationRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableLocationRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableLocationRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableLocationRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m WritableLocationRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableLocationRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m WritableLocationRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *WritableLocationRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetSite returns the Site property
func (m WritableLocationRequest) GetSite() int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *WritableLocationRequest) SetSite(val int32) {
	m.Site = val
}

// GetSlug returns the Slug property
func (m WritableLocationRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WritableLocationRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStatus returns the Status property
func (m WritableLocationRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WritableLocationRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m WritableLocationRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableLocationRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritableLocationRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritableLocationRequest) SetTenant(val *int32) {
	m.Tenant = val
}
