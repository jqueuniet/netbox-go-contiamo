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

// clusterTypeRequestSlugPattern is the validation pattern for ClusterTypeRequest.Slug
var clusterTypeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ClusterTypeRequest is an object. Adds support for custom fields and tags.
type ClusterTypeRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m ClusterTypeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(clusterTypeRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m ClusterTypeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ClusterTypeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ClusterTypeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ClusterTypeRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m ClusterTypeRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ClusterTypeRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m ClusterTypeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ClusterTypeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m ClusterTypeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ClusterTypeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
