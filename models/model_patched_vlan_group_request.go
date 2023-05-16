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

// patchedVLANGroupRequestSlugPattern is the validation pattern for PatchedVLANGroupRequest.Slug
var patchedVLANGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedVLANGroupRequest is an object. Adds support for custom fields and tags.
type PatchedVLANGroupRequest struct {
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// MaxVid: Highest permissible ID of a child VLAN
	MaxVid int32 `json:"max_vid,omitempty" mapstructure:"max_vid,omitempty"`
	// MinVid: Lowest permissible ID of a child VLAN
	MinVid int32 `json:"min_vid,omitempty" mapstructure:"min_vid,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// ScopeId:
	ScopeId *int32 `json:"scope_id,omitempty" mapstructure:"scope_id,omitempty"`
	// ScopeType:
	ScopeType *string `json:"scope_type,omitempty" mapstructure:"scope_type,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedVLANGroupRequest) Validate() error {
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
			m.Name, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedVLANGroupRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCustomFields returns the CustomFields property
func (m PatchedVLANGroupRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedVLANGroupRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedVLANGroupRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedVLANGroupRequest) SetDescription(val string) {
	m.Description = val
}

// GetMaxVid returns the MaxVid property
func (m PatchedVLANGroupRequest) GetMaxVid() int32 {
	return m.MaxVid
}

// SetMaxVid sets the MaxVid property
func (m *PatchedVLANGroupRequest) SetMaxVid(val int32) {
	m.MaxVid = val
}

// GetMinVid returns the MinVid property
func (m PatchedVLANGroupRequest) GetMinVid() int32 {
	return m.MinVid
}

// SetMinVid sets the MinVid property
func (m *PatchedVLANGroupRequest) SetMinVid(val int32) {
	m.MinVid = val
}

// GetName returns the Name property
func (m PatchedVLANGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedVLANGroupRequest) SetName(val string) {
	m.Name = val
}

// GetScopeId returns the ScopeId property
func (m PatchedVLANGroupRequest) GetScopeId() *int32 {
	return m.ScopeId
}

// SetScopeId sets the ScopeId property
func (m *PatchedVLANGroupRequest) SetScopeId(val *int32) {
	m.ScopeId = val
}

// GetScopeType returns the ScopeType property
func (m PatchedVLANGroupRequest) GetScopeType() *string {
	return m.ScopeType
}

// SetScopeType sets the ScopeType property
func (m *PatchedVLANGroupRequest) SetScopeType(val *string) {
	m.ScopeType = val
}

// GetSlug returns the Slug property
func (m PatchedVLANGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedVLANGroupRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedVLANGroupRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedVLANGroupRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
