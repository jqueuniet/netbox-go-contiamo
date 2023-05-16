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

// contactGroupSlugPattern is the validation pattern for ContactGroup.Slug
var contactGroupSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ContactGroup is an object. Extends PrimaryModelSerializer to include MPTT support.
type ContactGroup struct {
	// Depth:
	Depth int32 `json:"_depth" mapstructure:"_depth"`
	// ContactCount:
	ContactCount int32 `json:"contact_count" mapstructure:"contact_count"`
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
	Parent *NestedContactGroup `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ContactGroup) Validate() error {
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(contactGroupSlugPattern),
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
func (m ContactGroup) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *ContactGroup) SetDepth(val int32) {
	m.Depth = val
}

// GetContactCount returns the ContactCount property
func (m ContactGroup) GetContactCount() int32 {
	return m.ContactCount
}

// SetContactCount sets the ContactCount property
func (m *ContactGroup) SetContactCount(val int32) {
	m.ContactCount = val
}

// GetCreated returns the Created property
func (m ContactGroup) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ContactGroup) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ContactGroup) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ContactGroup) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ContactGroup) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ContactGroup) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ContactGroup) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ContactGroup) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ContactGroup) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ContactGroup) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ContactGroup) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ContactGroup) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ContactGroup) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ContactGroup) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m ContactGroup) GetParent() *NestedContactGroup {
	return m.Parent
}

// SetParent sets the Parent property
func (m *ContactGroup) SetParent(val *NestedContactGroup) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m ContactGroup) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ContactGroup) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m ContactGroup) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ContactGroup) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m ContactGroup) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ContactGroup) SetUrl(val string) {
	m.Url = val
}
