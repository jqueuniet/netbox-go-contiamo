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

// PatchedWritableAggregateRequest is an object. Adds support for custom fields and tags.
type PatchedWritableAggregateRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DateAdded:
	DateAdded *string `json:"date_added,omitempty" mapstructure:"date_added,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Prefix:
	Prefix string `json:"prefix,omitempty" mapstructure:"prefix,omitempty"`
	// Rir: Regional Internet Registry responsible for this IP space
	Rir int32 `json:"rir,omitempty" mapstructure:"rir,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableAggregateRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"prefix": validation.Validate(
			m.Prefix, validation.Length(1, 0),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableAggregateRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableAggregateRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableAggregateRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableAggregateRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDateAdded returns the DateAdded property
func (m PatchedWritableAggregateRequest) GetDateAdded() *string {
	return m.DateAdded
}

// SetDateAdded sets the DateAdded property
func (m *PatchedWritableAggregateRequest) SetDateAdded(val *string) {
	m.DateAdded = val
}

// GetDescription returns the Description property
func (m PatchedWritableAggregateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableAggregateRequest) SetDescription(val string) {
	m.Description = val
}

// GetPrefix returns the Prefix property
func (m PatchedWritableAggregateRequest) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *PatchedWritableAggregateRequest) SetPrefix(val string) {
	m.Prefix = val
}

// GetRir returns the Rir property
func (m PatchedWritableAggregateRequest) GetRir() int32 {
	return m.Rir
}

// SetRir sets the Rir property
func (m *PatchedWritableAggregateRequest) SetRir(val int32) {
	m.Rir = val
}

// GetTags returns the Tags property
func (m PatchedWritableAggregateRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableAggregateRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableAggregateRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableAggregateRequest) SetTenant(val *int32) {
	m.Tenant = val
}
