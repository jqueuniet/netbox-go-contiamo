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

// RouteTargetRequest is an object. Adds support for custom fields and tags.
type RouteTargetRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name: Route target value (formatted in accordance with RFC 4360)
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m RouteTargetRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 21),
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
func (m RouteTargetRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *RouteTargetRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m RouteTargetRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RouteTargetRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RouteTargetRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RouteTargetRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m RouteTargetRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RouteTargetRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m RouteTargetRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RouteTargetRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m RouteTargetRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *RouteTargetRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
