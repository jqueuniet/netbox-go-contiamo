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

// siteGroupSlugPattern is the validation pattern for SiteGroup.Slug
var siteGroupSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// SiteGroup is an object. Extends PrimaryModelSerializer to include MPTT support.
type SiteGroup struct {
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
	Parent *NestedSiteGroup `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// SiteCount:
	SiteCount int32 `json:"site_count" mapstructure:"site_count"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m SiteGroup) Validate() error {
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(siteGroupSlugPattern),
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
func (m SiteGroup) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *SiteGroup) SetDepth(val int32) {
	m.Depth = val
}

// GetCreated returns the Created property
func (m SiteGroup) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *SiteGroup) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m SiteGroup) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *SiteGroup) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m SiteGroup) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *SiteGroup) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m SiteGroup) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *SiteGroup) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m SiteGroup) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *SiteGroup) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m SiteGroup) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *SiteGroup) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m SiteGroup) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *SiteGroup) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m SiteGroup) GetParent() *NestedSiteGroup {
	return m.Parent
}

// SetParent sets the Parent property
func (m *SiteGroup) SetParent(val *NestedSiteGroup) {
	m.Parent = val
}

// GetSiteCount returns the SiteCount property
func (m SiteGroup) GetSiteCount() int32 {
	return m.SiteCount
}

// SetSiteCount sets the SiteCount property
func (m *SiteGroup) SetSiteCount(val int32) {
	m.SiteCount = val
}

// GetSlug returns the Slug property
func (m SiteGroup) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *SiteGroup) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m SiteGroup) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *SiteGroup) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m SiteGroup) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *SiteGroup) SetUrl(val string) {
	m.Url = val
}
