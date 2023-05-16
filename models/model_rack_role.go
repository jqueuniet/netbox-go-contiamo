// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"

	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// rackRoleColorPattern is the validation pattern for RackRole.Color
var rackRoleColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// rackRoleSlugPattern is the validation pattern for RackRole.Slug
var rackRoleSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// RackRole is an object. Adds support for custom fields and tags.
type RackRole struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
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
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// RackCount:
	RackCount int32 `json:"rack_count" mapstructure:"rack_count"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m RackRole) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(rackRoleColorPattern),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(rackRoleSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m RackRole) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *RackRole) SetColor(val string) {
	m.Color = val
}

// GetCreated returns the Created property
func (m RackRole) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *RackRole) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m RackRole) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RackRole) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RackRole) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RackRole) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m RackRole) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *RackRole) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m RackRole) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *RackRole) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m RackRole) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *RackRole) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m RackRole) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RackRole) SetName(val string) {
	m.Name = val
}

// GetRackCount returns the RackCount property
func (m RackRole) GetRackCount() int32 {
	return m.RackCount
}

// SetRackCount sets the RackCount property
func (m *RackRole) SetRackCount(val int32) {
	m.RackCount = val
}

// GetSlug returns the Slug property
func (m RackRole) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *RackRole) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m RackRole) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RackRole) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m RackRole) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *RackRole) SetUrl(val string) {
	m.Url = val
}
