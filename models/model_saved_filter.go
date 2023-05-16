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

// savedFilterSlugPattern is the validation pattern for SavedFilter.Slug
var savedFilterSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// SavedFilter is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type SavedFilter struct {
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parameters:
	Parameters map[string]interface{} `json:"parameters" mapstructure:"parameters"`
	// Shared:
	Shared bool `json:"shared,omitempty" mapstructure:"shared,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// User:
	User *int32 `json:"user,omitempty" mapstructure:"user,omitempty"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m SavedFilter) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"parameters": validation.Validate(
			m.Parameters, validation.NotNil,
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(savedFilterSlugPattern),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetContentTypes returns the ContentTypes property
func (m SavedFilter) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *SavedFilter) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetCreated returns the Created property
func (m SavedFilter) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *SavedFilter) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m SavedFilter) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *SavedFilter) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m SavedFilter) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *SavedFilter) SetDisplay(val string) {
	m.Display = val
}

// GetEnabled returns the Enabled property
func (m SavedFilter) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *SavedFilter) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m SavedFilter) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *SavedFilter) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m SavedFilter) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *SavedFilter) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m SavedFilter) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *SavedFilter) SetName(val string) {
	m.Name = val
}

// GetParameters returns the Parameters property
func (m SavedFilter) GetParameters() map[string]interface{} {
	return m.Parameters
}

// SetParameters sets the Parameters property
func (m *SavedFilter) SetParameters(val map[string]interface{}) {
	m.Parameters = val
}

// GetShared returns the Shared property
func (m SavedFilter) GetShared() bool {
	return m.Shared
}

// SetShared sets the Shared property
func (m *SavedFilter) SetShared(val bool) {
	m.Shared = val
}

// GetSlug returns the Slug property
func (m SavedFilter) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *SavedFilter) SetSlug(val string) {
	m.Slug = val
}

// GetUrl returns the Url property
func (m SavedFilter) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *SavedFilter) SetUrl(val string) {
	m.Url = val
}

// GetUser returns the User property
func (m SavedFilter) GetUser() *int32 {
	return m.User
}

// SetUser sets the User property
func (m *SavedFilter) SetUser(val *int32) {
	m.User = val
}

// GetWeight returns the Weight property
func (m SavedFilter) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *SavedFilter) SetWeight(val int32) {
	m.Weight = val
}
