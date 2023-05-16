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

// patchedWritableLocationRequestSlugPattern is the validation pattern for PatchedWritableLocationRequest.Slug
var patchedWritableLocationRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedWritableLocationRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type PatchedWritableLocationRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Site:
	Site int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
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
func (m PatchedWritableLocationRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedWritableLocationRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableLocationRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableLocationRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableLocationRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableLocationRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m PatchedWritableLocationRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableLocationRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m PatchedWritableLocationRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *PatchedWritableLocationRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetSite returns the Site property
func (m PatchedWritableLocationRequest) GetSite() int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *PatchedWritableLocationRequest) SetSite(val int32) {
	m.Site = val
}

// GetSlug returns the Slug property
func (m PatchedWritableLocationRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedWritableLocationRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStatus returns the Status property
func (m PatchedWritableLocationRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableLocationRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableLocationRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableLocationRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableLocationRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableLocationRequest) SetTenant(val *int32) {
	m.Tenant = val
}
