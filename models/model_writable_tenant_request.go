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

// writableTenantRequestSlugPattern is the validation pattern for WritableTenantRequest.Slug
var writableTenantRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WritableTenantRequest is an object. Adds support for custom fields and tags.
type WritableTenantRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Group:
	Group *int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableTenantRequest) Validate() error {
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(writableTenantRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableTenantRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableTenantRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableTenantRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableTenantRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableTenantRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableTenantRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m WritableTenantRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *WritableTenantRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetName returns the Name property
func (m WritableTenantRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableTenantRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m WritableTenantRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WritableTenantRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m WritableTenantRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableTenantRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
