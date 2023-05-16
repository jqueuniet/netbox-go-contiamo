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

// tenantGroupRequestSlugPattern is the validation pattern for TenantGroupRequest.Slug
var tenantGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// TenantGroupRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type TenantGroupRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedTenantGroupRequest `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m TenantGroupRequest) Validate() error {
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
		"parent": validation.Validate(
			m.Parent,
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(tenantGroupRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m TenantGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *TenantGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m TenantGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *TenantGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m TenantGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *TenantGroupRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m TenantGroupRequest) GetParent() *NestedTenantGroupRequest {
	return m.Parent
}

// SetParent sets the Parent property
func (m *TenantGroupRequest) SetParent(val *NestedTenantGroupRequest) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m TenantGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *TenantGroupRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m TenantGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *TenantGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
