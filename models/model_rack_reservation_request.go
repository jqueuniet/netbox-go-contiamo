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

// RackReservationRequest is an object. Adds support for custom fields and tags.
type RackReservationRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description" mapstructure:"description"`
	// Rack: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rack NestedRackRequest `json:"rack" mapstructure:"rack"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Units:
	Units []int32 `json:"units" mapstructure:"units"`
	// User: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	User NestedUserRequest `json:"user" mapstructure:"user"`
}

// Validate implements basic validation for this model
func (m RackReservationRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Required, validation.Length(1, 200),
		),
		"rack": validation.Validate(
			m.Rack, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"units": validation.Validate(
			m.Units, validation.NotNil,
		),
		"user": validation.Validate(
			m.User, validation.NotNil,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m RackReservationRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *RackReservationRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m RackReservationRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RackReservationRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RackReservationRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RackReservationRequest) SetDescription(val string) {
	m.Description = val
}

// GetRack returns the Rack property
func (m RackReservationRequest) GetRack() NestedRackRequest {
	return m.Rack
}

// SetRack sets the Rack property
func (m *RackReservationRequest) SetRack(val NestedRackRequest) {
	m.Rack = val
}

// GetTags returns the Tags property
func (m RackReservationRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RackReservationRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m RackReservationRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *RackReservationRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetUnits returns the Units property
func (m RackReservationRequest) GetUnits() []int32 {
	return m.Units
}

// SetUnits sets the Units property
func (m *RackReservationRequest) SetUnits(val []int32) {
	m.Units = val
}

// GetUser returns the User property
func (m RackReservationRequest) GetUser() NestedUserRequest {
	return m.User
}

// SetUser sets the User property
func (m *RackReservationRequest) SetUser(val NestedUserRequest) {
	m.User = val
}
