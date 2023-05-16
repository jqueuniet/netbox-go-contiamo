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

	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// contactRoleSlugPattern is the validation pattern for ContactRole.Slug
var contactRoleSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ContactRole is an object. Adds support for custom fields and tags.
type ContactRole struct {
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
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ContactRole) Validate() error {
	return validation.Errors{
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(contactRoleSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m ContactRole) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ContactRole) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ContactRole) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ContactRole) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ContactRole) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ContactRole) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ContactRole) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ContactRole) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ContactRole) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ContactRole) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ContactRole) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ContactRole) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ContactRole) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ContactRole) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m ContactRole) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ContactRole) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m ContactRole) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ContactRole) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m ContactRole) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ContactRole) SetUrl(val string) {
	m.Url = val
}
