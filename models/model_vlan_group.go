// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// vLANGroupSlugPattern is the validation pattern for VLANGroup.Slug
var vLANGroupSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// VLANGroup is an object. Adds support for custom fields and tags.
type VLANGroup struct {
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// MaxVid: Highest permissible ID of a child VLAN
	MaxVid int32 `json:"max_vid,omitempty" mapstructure:"max_vid,omitempty"`
	// MinVid: Lowest permissible ID of a child VLAN
	MinVid int32 `json:"min_vid,omitempty" mapstructure:"min_vid,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Scope:
	Scope map[string]interface{} `json:"scope" mapstructure:"scope"`
	// ScopeId:
	ScopeId *int32 `json:"scope_id,omitempty" mapstructure:"scope_id,omitempty"`
	// ScopeType:
	ScopeType *string `json:"scope_type,omitempty" mapstructure:"scope_type,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VlanCount:
	VlanCount int32 `json:"vlan_count" mapstructure:"vlan_count"`
}

// Validate implements basic validation for this model
func (m VLANGroup) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"scope": validation.Validate(
			m.Scope,
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(vLANGroupSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m VLANGroup) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VLANGroup) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m VLANGroup) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VLANGroup) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VLANGroup) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VLANGroup) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m VLANGroup) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *VLANGroup) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m VLANGroup) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VLANGroup) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m VLANGroup) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VLANGroup) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMaxVid returns the MaxVid property
func (m VLANGroup) GetMaxVid() int32 {
	return m.MaxVid
}

// SetMaxVid sets the MaxVid property
func (m *VLANGroup) SetMaxVid(val int32) {
	m.MaxVid = val
}

// GetMinVid returns the MinVid property
func (m VLANGroup) GetMinVid() int32 {
	return m.MinVid
}

// SetMinVid sets the MinVid property
func (m *VLANGroup) SetMinVid(val int32) {
	m.MinVid = val
}

// GetName returns the Name property
func (m VLANGroup) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VLANGroup) SetName(val string) {
	m.Name = val
}

// GetScope returns the Scope property
func (m VLANGroup) GetScope() map[string]interface{} {
	return m.Scope
}

// SetScope sets the Scope property
func (m *VLANGroup) SetScope(val map[string]interface{}) {
	m.Scope = val
}

// GetScopeId returns the ScopeId property
func (m VLANGroup) GetScopeId() *int32 {
	return m.ScopeId
}

// SetScopeId sets the ScopeId property
func (m *VLANGroup) SetScopeId(val *int32) {
	m.ScopeId = val
}

// GetScopeType returns the ScopeType property
func (m VLANGroup) GetScopeType() *string {
	return m.ScopeType
}

// SetScopeType sets the ScopeType property
func (m *VLANGroup) SetScopeType(val *string) {
	m.ScopeType = val
}

// GetSlug returns the Slug property
func (m VLANGroup) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *VLANGroup) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m VLANGroup) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VLANGroup) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m VLANGroup) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *VLANGroup) SetUrl(val string) {
	m.Url = val
}

// GetVlanCount returns the VlanCount property
func (m VLANGroup) GetVlanCount() int32 {
	return m.VlanCount
}

// SetVlanCount sets the VlanCount property
func (m *VLANGroup) SetVlanCount(val int32) {
	m.VlanCount = val
}
