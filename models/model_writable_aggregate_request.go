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

// WritableAggregateRequest is an object. Adds support for custom fields and tags.
type WritableAggregateRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DateAdded:
	DateAdded *string `json:"date_added,omitempty" mapstructure:"date_added,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Prefix:
	Prefix string `json:"prefix" mapstructure:"prefix"`
	// Rir: Regional Internet Registry responsible for this IP space
	Rir int32 `json:"rir" mapstructure:"rir"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableAggregateRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"prefix": validation.Validate(
			m.Prefix, validation.Required, validation.Length(1, 0),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableAggregateRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableAggregateRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableAggregateRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableAggregateRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDateAdded returns the DateAdded property
func (m WritableAggregateRequest) GetDateAdded() *string {
	return m.DateAdded
}

// SetDateAdded sets the DateAdded property
func (m *WritableAggregateRequest) SetDateAdded(val *string) {
	m.DateAdded = val
}

// GetDescription returns the Description property
func (m WritableAggregateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableAggregateRequest) SetDescription(val string) {
	m.Description = val
}

// GetPrefix returns the Prefix property
func (m WritableAggregateRequest) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *WritableAggregateRequest) SetPrefix(val string) {
	m.Prefix = val
}

// GetRir returns the Rir property
func (m WritableAggregateRequest) GetRir() int32 {
	return m.Rir
}

// SetRir sets the Rir property
func (m *WritableAggregateRequest) SetRir(val int32) {
	m.Rir = val
}

// GetTags returns the Tags property
func (m WritableAggregateRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableAggregateRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m WritableAggregateRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WritableAggregateRequest) SetTenant(val *int32) {
	m.Tenant = val
}
