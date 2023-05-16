// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// RackReservation is an object. Adds support for custom fields and tags.
type RackReservation struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description" mapstructure:"description"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Rack: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rack NestedRack `json:"rack" mapstructure:"rack"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Units:
	Units []int32 `json:"units" mapstructure:"units"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// User: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	User NestedUser `json:"user" mapstructure:"user"`
}

// Validate implements basic validation for this model
func (m RackReservation) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.NotNil, validation.Length(0, 200),
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"user": validation.Validate(
			m.User, validation.NotNil,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m RackReservation) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *RackReservation) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m RackReservation) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *RackReservation) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m RackReservation) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RackReservation) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RackReservation) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RackReservation) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m RackReservation) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *RackReservation) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m RackReservation) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *RackReservation) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m RackReservation) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *RackReservation) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetRack returns the Rack property
func (m RackReservation) GetRack() NestedRack {
	return m.Rack
}

// SetRack sets the Rack property
func (m *RackReservation) SetRack(val NestedRack) {
	m.Rack = val
}

// GetTags returns the Tags property
func (m RackReservation) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RackReservation) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m RackReservation) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *RackReservation) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUnits returns the Units property
func (m RackReservation) GetUnits() []int32 {
	return m.Units
}

// SetUnits sets the Units property
func (m *RackReservation) SetUnits(val []int32) {
	m.Units = val
}

// GetUrl returns the Url property
func (m RackReservation) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *RackReservation) SetUrl(val string) {
	m.Url = val
}

// GetUser returns the User property
func (m RackReservation) GetUser() NestedUser {
	return m.User
}

// SetUser sets the User property
func (m *RackReservation) SetUser(val NestedUser) {
	m.User = val
}
