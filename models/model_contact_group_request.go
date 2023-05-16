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

// contactGroupRequestSlugPattern is the validation pattern for ContactGroupRequest.Slug
var contactGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ContactGroupRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type ContactGroupRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedContactGroupRequest `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m ContactGroupRequest) Validate() error {
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(contactGroupRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m ContactGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ContactGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ContactGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ContactGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m ContactGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ContactGroupRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m ContactGroupRequest) GetParent() *NestedContactGroupRequest {
	return m.Parent
}

// SetParent sets the Parent property
func (m *ContactGroupRequest) SetParent(val *NestedContactGroupRequest) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m ContactGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ContactGroupRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m ContactGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ContactGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
