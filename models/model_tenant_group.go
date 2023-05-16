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

// tenantGroupSlugPattern is the validation pattern for TenantGroup.Slug
var tenantGroupSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// TenantGroup is an object. Extends PrimaryModelSerializer to include MPTT support.
type TenantGroup struct {
	// Depth:
	Depth int32 `json:"_depth" mapstructure:"_depth"`
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
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedTenantGroup `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TenantCount:
	TenantCount int32 `json:"tenant_count" mapstructure:"tenant_count"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m TenantGroup) Validate() error {
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
		"parent": validation.Validate(
			m.Parent,
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(tenantGroupSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDepth returns the Depth property
func (m TenantGroup) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *TenantGroup) SetDepth(val int32) {
	m.Depth = val
}

// GetCreated returns the Created property
func (m TenantGroup) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *TenantGroup) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m TenantGroup) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *TenantGroup) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m TenantGroup) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *TenantGroup) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m TenantGroup) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *TenantGroup) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m TenantGroup) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *TenantGroup) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m TenantGroup) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *TenantGroup) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m TenantGroup) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *TenantGroup) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m TenantGroup) GetParent() *NestedTenantGroup {
	return m.Parent
}

// SetParent sets the Parent property
func (m *TenantGroup) SetParent(val *NestedTenantGroup) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m TenantGroup) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *TenantGroup) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m TenantGroup) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *TenantGroup) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenantCount returns the TenantCount property
func (m TenantGroup) GetTenantCount() int32 {
	return m.TenantCount
}

// SetTenantCount sets the TenantCount property
func (m *TenantGroup) SetTenantCount(val int32) {
	m.TenantCount = val
}

// GetUrl returns the Url property
func (m TenantGroup) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *TenantGroup) SetUrl(val string) {
	m.Url = val
}
