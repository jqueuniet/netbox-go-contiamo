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

// rIRSlugPattern is the validation pattern for RIR.Slug
var rIRSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// RIR is an object. Adds support for custom fields and tags.
type RIR struct {
	// AggregateCount:
	AggregateCount int32 `json:"aggregate_count" mapstructure:"aggregate_count"`
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
	// IsPrivate: IP space managed by this RIR is considered private
	IsPrivate bool `json:"is_private,omitempty" mapstructure:"is_private,omitempty"`
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
func (m RIR) Validate() error {
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(rIRSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAggregateCount returns the AggregateCount property
func (m RIR) GetAggregateCount() int32 {
	return m.AggregateCount
}

// SetAggregateCount sets the AggregateCount property
func (m *RIR) SetAggregateCount(val int32) {
	m.AggregateCount = val
}

// GetCreated returns the Created property
func (m RIR) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *RIR) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m RIR) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *RIR) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m RIR) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *RIR) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m RIR) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *RIR) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m RIR) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *RIR) SetId(val int32) {
	m.Id = val
}

// GetIsPrivate returns the IsPrivate property
func (m RIR) GetIsPrivate() bool {
	return m.IsPrivate
}

// SetIsPrivate sets the IsPrivate property
func (m *RIR) SetIsPrivate(val bool) {
	m.IsPrivate = val
}

// GetLastUpdated returns the LastUpdated property
func (m RIR) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *RIR) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m RIR) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *RIR) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m RIR) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *RIR) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m RIR) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *RIR) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m RIR) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *RIR) SetUrl(val string) {
	m.Url = val
}
