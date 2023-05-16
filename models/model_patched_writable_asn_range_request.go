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

// patchedWritableASNRangeRequestSlugPattern is the validation pattern for PatchedWritableASNRangeRequest.Slug
var patchedWritableASNRangeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedWritableASNRangeRequest is an object. Adds support for custom fields and tags.
type PatchedWritableASNRangeRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// End:
	End int64 `json:"end,omitempty" mapstructure:"end,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Rir:
	Rir int32 `json:"rir,omitempty" mapstructure:"rir,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// Start:
	Start int64 `json:"start,omitempty" mapstructure:"start,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableASNRangeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"end": validation.Validate(
			m.End, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedWritableASNRangeRequestSlugPattern),
		),
		"start": validation.Validate(
			m.Start, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableASNRangeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableASNRangeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableASNRangeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableASNRangeRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnd returns the End property
func (m PatchedWritableASNRangeRequest) GetEnd() int64 {
	return m.End
}

// SetEnd sets the End property
func (m *PatchedWritableASNRangeRequest) SetEnd(val int64) {
	m.End = val
}

// GetName returns the Name property
func (m PatchedWritableASNRangeRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableASNRangeRequest) SetName(val string) {
	m.Name = val
}

// GetRir returns the Rir property
func (m PatchedWritableASNRangeRequest) GetRir() int32 {
	return m.Rir
}

// SetRir sets the Rir property
func (m *PatchedWritableASNRangeRequest) SetRir(val int32) {
	m.Rir = val
}

// GetSlug returns the Slug property
func (m PatchedWritableASNRangeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedWritableASNRangeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetStart returns the Start property
func (m PatchedWritableASNRangeRequest) GetStart() int64 {
	return m.Start
}

// SetStart sets the Start property
func (m *PatchedWritableASNRangeRequest) SetStart(val int64) {
	m.Start = val
}

// GetTags returns the Tags property
func (m PatchedWritableASNRangeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableASNRangeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableASNRangeRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableASNRangeRequest) SetTenant(val *int32) {
	m.Tenant = val
}
