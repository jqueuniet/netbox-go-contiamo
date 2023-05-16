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

// vLANGroupRequestSlugPattern is the validation pattern for VLANGroupRequest.Slug
var vLANGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// VLANGroupRequest is an object. Adds support for custom fields and tags.
type VLANGroupRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// MaxVid: Highest permissible ID of a child VLAN
	MaxVid int32 `json:"max_vid,omitempty" mapstructure:"max_vid,omitempty"`
	// MinVid: Lowest permissible ID of a child VLAN
	MinVid int32 `json:"min_vid,omitempty" mapstructure:"min_vid,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// ScopeId:
	ScopeId *int32 `json:"scope_id,omitempty" mapstructure:"scope_id,omitempty"`
	// ScopeType:
	ScopeType *string `json:"scope_type,omitempty" mapstructure:"scope_type,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m VLANGroupRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"maxVid": validation.Validate(
			m.MaxVid, validation.Min(int32(1)), validation.Max(int32(4094)),
		),
		"minVid": validation.Validate(
			m.MinVid, validation.Min(int32(1)), validation.Max(int32(4094)),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(vLANGroupRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m VLANGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VLANGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VLANGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VLANGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetMaxVid returns the MaxVid property
func (m VLANGroupRequest) GetMaxVid() int32 {
	return m.MaxVid
}

// SetMaxVid sets the MaxVid property
func (m *VLANGroupRequest) SetMaxVid(val int32) {
	m.MaxVid = val
}

// GetMinVid returns the MinVid property
func (m VLANGroupRequest) GetMinVid() int32 {
	return m.MinVid
}

// SetMinVid sets the MinVid property
func (m *VLANGroupRequest) SetMinVid(val int32) {
	m.MinVid = val
}

// GetName returns the Name property
func (m VLANGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VLANGroupRequest) SetName(val string) {
	m.Name = val
}

// GetScopeId returns the ScopeId property
func (m VLANGroupRequest) GetScopeId() *int32 {
	return m.ScopeId
}

// SetScopeId sets the ScopeId property
func (m *VLANGroupRequest) SetScopeId(val *int32) {
	m.ScopeId = val
}

// GetScopeType returns the ScopeType property
func (m VLANGroupRequest) GetScopeType() *string {
	return m.ScopeType
}

// SetScopeType sets the ScopeType property
func (m *VLANGroupRequest) SetScopeType(val *string) {
	m.ScopeType = val
}

// GetSlug returns the Slug property
func (m VLANGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *VLANGroupRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m VLANGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VLANGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
