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

// writableASNRangeRequestSlugPattern is the validation pattern for WritableASNRangeRequest.Slug
var writableASNRangeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WritableASNRangeRequest is an object. Adds support for custom fields and tags.
type WritableASNRangeRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// End:
	End int64 `json:"end" mapstructure:"end"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Rir:
	Rir int32 `json:"rir" mapstructure:"rir"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Start:
	Start int64 `json:"start" mapstructure:"start"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableASNRangeRequest) Validate() error {
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
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(writableASNRangeRequestSlugPattern),
		),
		"start": validation.Validate(
			m.Start, validation.Required, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m WritableASNRangeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableASNRangeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableASNRangeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableASNRangeRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnd returns the End property
func (m WritableASNRangeRequest) GetEnd() int64 {
	return m.End
}

// SetEnd sets the End property
func (m *WritableASNRangeRequest) SetEnd(val int64) {
	m.End = val
}

// GetName returns the Name property
func (m WritableASNRangeRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableASNRangeRequest) SetName(val string) {
	m.Name = val
}

// GetRir returns the Rir property
func (m WritableASNRangeRequest) GetRir() int32 {
	return m.Rir
}

// SetRir sets the Rir property
func (m *WritableASNRangeRequest) SetRir(val int32) {
	m.Rir = val
}

// GetSlug returns the Slug property
func (m WritableASNRangeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WritableASNRangeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStart returns the Start property
func (m WritableASNRangeRequest) GetStart() int64 {
	return m.Start
}

// SetStart sets the Start property
func (m *WritableASNRangeRequest) SetStart(val int64) {
	m.Start = val
}

// GetTags returns the Tags property
func (m WritableASNRangeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableASNRangeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritableASNRangeRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritableASNRangeRequest) SetTenant(val *int32) {
	m.Tenant = val
}
