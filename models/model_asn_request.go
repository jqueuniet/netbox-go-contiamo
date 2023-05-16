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

// ASNRequest is an object. Adds support for custom fields and tags.
type ASNRequest struct {
	// Asn: 16- or 32-bit autonomous system number
	Asn int64 `json:"asn" mapstructure:"asn"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Rir: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rir *NestedRIRRequest `json:"rir,omitempty" mapstructure:"rir,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m ASNRequest) Validate() error {
	return validation.Errors{
		"asn": validation.Validate(
			m.Asn, validation.Required, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"rir": validation.Validate(
			m.Rir,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
	}.Filter()
}

// GetAsn returns the Asn property
func (m ASNRequest) GetAsn() int64 {
	return m.Asn
}

// SetAsn sets the Asn property
func (m *ASNRequest) SetAsn(val int64) {
	m.Asn = val
}

// GetComments returns the Comments property
func (m ASNRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ASNRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m ASNRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ASNRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ASNRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ASNRequest) SetDescription(val string) {
	m.Description = val
}

// GetRir returns the Rir property
func (m ASNRequest) GetRir() *NestedRIRRequest {
	return m.Rir
}

// SetRir sets the Rir property
func (m *ASNRequest) SetRir(val *NestedRIRRequest) {
	m.Rir = val
}

// GetTags returns the Tags property
func (m ASNRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ASNRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m ASNRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *ASNRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
