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

// aSNRangeRequestSlugPattern is the validation pattern for ASNRangeRequest.Slug
var aSNRangeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ASNRangeRequest is an object. Adds support for custom fields and tags.
type ASNRangeRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// End:
	End int64 `json:"end" mapstructure:"end"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Rir: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rir NestedRIRRequest `json:"rir" mapstructure:"rir"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Start:
	Start int64 `json:"start" mapstructure:"start"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m ASNRangeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"end": validation.Validate(
			m.End, validation.Required, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"rir": validation.Validate(
			m.Rir, validation.NotNil,
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(aSNRangeRequestSlugPattern),
		),
		"start": validation.Validate(
			m.Start, validation.Required, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
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
func (m ASNRangeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ASNRangeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ASNRangeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ASNRangeRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnd returns the End property
func (m ASNRangeRequest) GetEnd() int64 {
	return m.End
}

// SetEnd sets the End property
func (m *ASNRangeRequest) SetEnd(val int64) {
	m.End = val
}

// GetName returns the Name property
func (m ASNRangeRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ASNRangeRequest) SetName(val string) {
	m.Name = val
}

// GetRir returns the Rir property
func (m ASNRangeRequest) GetRir() NestedRIRRequest {
	return m.Rir
}

// SetRir sets the Rir property
func (m *ASNRangeRequest) SetRir(val NestedRIRRequest) {
	m.Rir = val
}

// GetSlug returns the Slug property
func (m ASNRangeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ASNRangeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStart returns the Start property
func (m ASNRangeRequest) GetStart() int64 {
	return m.Start
}

// SetStart sets the Start property
func (m *ASNRangeRequest) SetStart(val int64) {
	m.Start = val
}

// GetTags returns the Tags property
func (m ASNRangeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ASNRangeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m ASNRangeRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *ASNRangeRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
