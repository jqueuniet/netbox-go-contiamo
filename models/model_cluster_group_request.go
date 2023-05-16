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

// clusterGroupRequestSlugPattern is the validation pattern for ClusterGroupRequest.Slug
var clusterGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ClusterGroupRequest is an object. Adds support for custom fields and tags.
type ClusterGroupRequest struct {
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
func (m ClusterGroupRequest) Validate() error {
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(clusterGroupRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m ClusterGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ClusterGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ClusterGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ClusterGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m ClusterGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ClusterGroupRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m ClusterGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ClusterGroupRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m ClusterGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ClusterGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
