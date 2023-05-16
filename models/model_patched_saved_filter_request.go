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

// patchedSavedFilterRequestSlugPattern is the validation pattern for PatchedSavedFilterRequest.Slug
var patchedSavedFilterRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedSavedFilterRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedSavedFilterRequest struct {
	// ContentTypes:
	ContentTypes []string `json:"content_types,omitempty" mapstructure:"content_types,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Parameters:
	Parameters map[string]interface{} `json:"parameters,omitempty" mapstructure:"parameters,omitempty"`
	// Shared:
	Shared bool `json:"shared,omitempty" mapstructure:"shared,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// User:
	User *int32 `json:"user,omitempty" mapstructure:"user,omitempty"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedSavedFilterRequest) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"parameters": validation.Validate(
			m.Parameters,
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedSavedFilterRequestSlugPattern),
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetContentTypes returns the ContentTypes property
func (m PatchedSavedFilterRequest) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *PatchedSavedFilterRequest) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetDescription returns the Description property
func (m PatchedSavedFilterRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedSavedFilterRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnabled returns the Enabled property
func (m PatchedSavedFilterRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *PatchedSavedFilterRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetName returns the Name property
func (m PatchedSavedFilterRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedSavedFilterRequest) SetName(val string) {
	m.Name = val
}

// GetParameters returns the Parameters property
func (m PatchedSavedFilterRequest) GetParameters() map[string]interface{} {
	return m.Parameters
}

// SetParameters sets the Parameters property
func (m *PatchedSavedFilterRequest) SetParameters(val map[string]interface{}) {
	m.Parameters = val
}

// GetShared returns the Shared property
func (m PatchedSavedFilterRequest) GetShared() bool {
	return m.Shared
}

// SetShared sets the Shared property
func (m *PatchedSavedFilterRequest) SetShared(val bool) {
	m.Shared = val
}

// GetSlug returns the Slug property
func (m PatchedSavedFilterRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedSavedFilterRequest) SetSlug(val string) {
	m.Slug = val
}

// GetUser returns the User property
func (m PatchedSavedFilterRequest) GetUser() *int32 {
	return m.User
}

// SetUser sets the User property
func (m *PatchedSavedFilterRequest) SetUser(val *int32) {
	m.User = val
}

// GetWeight returns the Weight property
func (m PatchedSavedFilterRequest) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *PatchedSavedFilterRequest) SetWeight(val int32) {
	m.Weight = val
}
