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
)

// savedFilterRequestSlugPattern is the validation pattern for SavedFilterRequest.Slug
var savedFilterRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// SavedFilterRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type SavedFilterRequest struct {
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parameters:
	Parameters map[string]interface{} `json:"parameters" mapstructure:"parameters"`
	// Shared:
	Shared bool `json:"shared,omitempty" mapstructure:"shared,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// User:
	User *int32 `json:"user,omitempty" mapstructure:"user,omitempty"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m SavedFilterRequest) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"parameters": validation.Validate(
			m.Parameters, validation.NotNil,
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(savedFilterRequestSlugPattern),
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetContentTypes returns the ContentTypes property
func (m SavedFilterRequest) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *SavedFilterRequest) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetDescription returns the Description property
func (m SavedFilterRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *SavedFilterRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnabled returns the Enabled property
func (m SavedFilterRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *SavedFilterRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetName returns the Name property
func (m SavedFilterRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *SavedFilterRequest) SetName(val string) {
	m.Name = val
}

// GetParameters returns the Parameters property
func (m SavedFilterRequest) GetParameters() map[string]interface{} {
	return m.Parameters
}

// SetParameters sets the Parameters property
func (m *SavedFilterRequest) SetParameters(val map[string]interface{}) {
	m.Parameters = val
}

// GetShared returns the Shared property
func (m SavedFilterRequest) GetShared() bool {
	return m.Shared
}

// SetShared sets the Shared property
func (m *SavedFilterRequest) SetShared(val bool) {
	m.Shared = val
}

// GetSlug returns the Slug property
func (m SavedFilterRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *SavedFilterRequest) SetSlug(val string) {
	m.Slug = val
}

// GetUser returns the User property
func (m SavedFilterRequest) GetUser() *int32 {
	return m.User
}

// SetUser sets the User property
func (m *SavedFilterRequest) SetUser(val *int32) {
	m.User = val
}

// GetWeight returns the Weight property
func (m SavedFilterRequest) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *SavedFilterRequest) SetWeight(val int32) {
	m.Weight = val
}
