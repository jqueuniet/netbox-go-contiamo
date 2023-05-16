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

// RouteTarget is an object. Adds support for custom fields and tags.
type RouteTarget struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name: Route target value (formatted in accordance with RFC 4360)
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m RouteTarget) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 21),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m RouteTarget) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *RouteTarget) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m RouteTarget) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *RouteTarget) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m RouteTarget) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RouteTarget) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RouteTarget) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RouteTarget) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m RouteTarget) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *RouteTarget) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m RouteTarget) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *RouteTarget) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m RouteTarget) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *RouteTarget) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m RouteTarget) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RouteTarget) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m RouteTarget) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RouteTarget) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m RouteTarget) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *RouteTarget) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m RouteTarget) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *RouteTarget) SetUrl(val string) {
	m.Url = val
}
