// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// PatchedWritableASNRequest is an object. Adds support for custom fields and tags.
type PatchedWritableASNRequest struct {
	// Asn: 16- or 32-bit autonomous system number
	Asn int64 `json:"asn,omitempty" mapstructure:"asn,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Rir: Regional Internet Registry responsible for this AS number space
	Rir int32 `json:"rir,omitempty" mapstructure:"rir,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableASNRequest) Validate() error {
	return validation.Errors{
		"asn": validation.Validate(
			m.Asn, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAsn returns the Asn property
func (m PatchedWritableASNRequest) GetAsn() int64 {
	return m.Asn
}

// SetAsn sets the Asn property
func (m *PatchedWritableASNRequest) SetAsn(val int64) {
	m.Asn = val
}

// GetComments returns the Comments property
func (m PatchedWritableASNRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableASNRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableASNRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableASNRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableASNRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableASNRequest) SetDescription(val string) {
	m.Description = val
}

// GetRir returns the Rir property
func (m PatchedWritableASNRequest) GetRir() int32 {
	return m.Rir
}

// SetRir sets the Rir property
func (m *PatchedWritableASNRequest) SetRir(val int32) {
	m.Rir = val
}

// GetTags returns the Tags property
func (m PatchedWritableASNRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableASNRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableASNRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableASNRequest) SetTenant(val *int32) {
	m.Tenant = val
}
