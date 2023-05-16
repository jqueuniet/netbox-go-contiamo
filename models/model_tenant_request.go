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

// tenantRequestSlugPattern is the validation pattern for TenantRequest.Slug
var tenantRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// TenantRequest is an object. Adds support for custom fields and tags.
type TenantRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedTenantGroupRequest `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m TenantRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"group": validation.Validate(
			m.Group,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(tenantRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m TenantRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *TenantRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m TenantRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *TenantRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m TenantRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *TenantRequest) SetDescription(val string) {
	m.Description = val
}

// GetGroup returns the Group property
func (m TenantRequest) GetGroup() *NestedTenantGroupRequest {
	return m.Group
}

// SetGroup sets the Group property
func (m *TenantRequest) SetGroup(val *NestedTenantGroupRequest) {
	m.Group = val
}

// GetName returns the Name property
func (m TenantRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *TenantRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m TenantRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *TenantRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m TenantRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *TenantRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
