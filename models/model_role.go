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

// roleSlugPattern is the validation pattern for Role.Slug
var roleSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// Role is an object. Adds support for custom fields and tags.
type Role struct {
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
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PrefixCount:
	PrefixCount int32 `json:"prefix_count" mapstructure:"prefix_count"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VlanCount:
	VlanCount int32 `json:"vlan_count" mapstructure:"vlan_count"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m Role) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(roleSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m Role) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Role) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Role) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Role) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Role) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Role) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Role) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Role) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Role) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Role) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Role) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Role) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Role) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Role) SetName(val string) {
	m.Name = val
}

// GetPrefixCount returns the PrefixCount property
func (m Role) GetPrefixCount() int32 {
	return m.PrefixCount
}

// SetPrefixCount sets the PrefixCount property
func (m *Role) SetPrefixCount(val int32) {
	m.PrefixCount = val
}

// GetSlug returns the Slug property
func (m Role) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Role) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m Role) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Role) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m Role) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Role) SetUrl(val string) {
	m.Url = val
}

// GetVlanCount returns the VlanCount property
func (m Role) GetVlanCount() int32 {
	return m.VlanCount
}

// SetVlanCount sets the VlanCount property
func (m *Role) SetVlanCount(val int32) {
	m.VlanCount = val
}

// GetWeight returns the Weight property
func (m Role) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *Role) SetWeight(val int32) {
	m.Weight = val
}
