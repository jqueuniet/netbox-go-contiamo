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

// PatchedWritableRackReservationRequest is an object. Adds support for custom fields and tags.
type PatchedWritableRackReservationRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Rack:
	Rack int32 `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Units:
	Units []int32 `json:"units,omitempty" mapstructure:"units,omitempty"`
	// User:
	User int32 `json:"user,omitempty" mapstructure:"user,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableRackReservationRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(1, 200),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"units": validation.Validate(
			m.Units,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableRackReservationRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableRackReservationRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableRackReservationRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableRackReservationRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableRackReservationRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableRackReservationRequest) SetDescription(val string) {
	m.Description = val
}

// GetRack returns the Rack property
func (m PatchedWritableRackReservationRequest) GetRack() int32 {
	return m.Rack
}

// SetRack sets the Rack property
func (m *PatchedWritableRackReservationRequest) SetRack(val int32) {
	m.Rack = val
}

// GetTags returns the Tags property
func (m PatchedWritableRackReservationRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableRackReservationRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableRackReservationRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableRackReservationRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetUnits returns the Units property
func (m PatchedWritableRackReservationRequest) GetUnits() []int32 {
	return m.Units
}

// SetUnits sets the Units property
func (m *PatchedWritableRackReservationRequest) SetUnits(val []int32) {
	m.Units = val
}

// GetUser returns the User property
func (m PatchedWritableRackReservationRequest) GetUser() int32 {
	return m.User
}

// SetUser sets the User property
func (m *PatchedWritableRackReservationRequest) SetUser(val int32) {
	m.User = val
}
