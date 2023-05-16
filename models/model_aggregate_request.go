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

// AggregateRequest is an object. Adds support for custom fields and tags.
type AggregateRequest struct {
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
	// Rir: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rir NestedRIRRequest `json:"rir" mapstructure:"rir"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m AggregateRequest) Validate() error {
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
		"rir": validation.Validate(
			m.Rir, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m AggregateRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *AggregateRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m AggregateRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *AggregateRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDateAdded returns the DateAdded property
func (m AggregateRequest) GetDateAdded() *string {
	return m.DateAdded
}

// SetDateAdded sets the DateAdded property
func (m *AggregateRequest) SetDateAdded(val *string) {
	m.DateAdded = val
}

// GetDescription returns the Description property
func (m AggregateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *AggregateRequest) SetDescription(val string) {
	m.Description = val
}

// GetPrefix returns the Prefix property
func (m AggregateRequest) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *AggregateRequest) SetPrefix(val string) {
	m.Prefix = val
}

// GetRir returns the Rir property
func (m AggregateRequest) GetRir() NestedRIRRequest {
	return m.Rir
}

// SetRir sets the Rir property
func (m *AggregateRequest) SetRir(val NestedRIRRequest) {
	m.Rir = val
}

// GetTags returns the Tags property
func (m AggregateRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *AggregateRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m AggregateRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *AggregateRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
