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

// circuitTypeSlugPattern is the validation pattern for CircuitType.Slug
var circuitTypeSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// CircuitType is an object. Adds support for custom fields and tags.
type CircuitType struct {
	// CircuitCount:
	CircuitCount int32 `json:"circuit_count" mapstructure:"circuit_count"`
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
func (m CircuitType) Validate() error {
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(circuitTypeSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCircuitCount returns the CircuitCount property
func (m CircuitType) GetCircuitCount() int32 {
	return m.CircuitCount
}

// SetCircuitCount sets the CircuitCount property
func (m *CircuitType) SetCircuitCount(val int32) {
	m.CircuitCount = val
}

// GetCreated returns the Created property
func (m CircuitType) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CircuitType) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m CircuitType) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *CircuitType) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m CircuitType) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitType) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m CircuitType) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *CircuitType) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m CircuitType) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitType) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m CircuitType) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CircuitType) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m CircuitType) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *CircuitType) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m CircuitType) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *CircuitType) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m CircuitType) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *CircuitType) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m CircuitType) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *CircuitType) SetUrl(val string) {
	m.Url = val
}
