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

// wirelessLANGroupRequestSlugPattern is the validation pattern for WirelessLANGroupRequest.Slug
var wirelessLANGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WirelessLANGroupRequest is an object. Extends PrimaryModelSerializer to include MPTT support.
type WirelessLANGroupRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedWirelessLANGroupRequest `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WirelessLANGroupRequest) Validate() error {
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
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(wirelessLANGroupRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m WirelessLANGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WirelessLANGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WirelessLANGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessLANGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m WirelessLANGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WirelessLANGroupRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m WirelessLANGroupRequest) GetParent() *NestedWirelessLANGroupRequest {
	return m.Parent
}

// SetParent sets the Parent property
func (m *WirelessLANGroupRequest) SetParent(val *NestedWirelessLANGroupRequest) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m WirelessLANGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WirelessLANGroupRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m WirelessLANGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WirelessLANGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
