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

// wirelessLANGroupSlugPattern is the validation pattern for WirelessLANGroup.Slug
var wirelessLANGroupSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WirelessLANGroup is an object. Extends PrimaryModelSerializer to include MPTT support.
type WirelessLANGroup struct {
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
	Parent *NestedWirelessLANGroup `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// WirelesslanCount:
	WirelesslanCount int32 `json:"wirelesslan_count" mapstructure:"wirelesslan_count"`
}

// Validate implements basic validation for this model
func (m WirelessLANGroup) Validate() error {
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(wirelessLANGroupSlugPattern),
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
func (m WirelessLANGroup) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *WirelessLANGroup) SetDepth(val int32) {
	m.Depth = val
}

// GetCreated returns the Created property
func (m WirelessLANGroup) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *WirelessLANGroup) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m WirelessLANGroup) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WirelessLANGroup) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WirelessLANGroup) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessLANGroup) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m WirelessLANGroup) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *WirelessLANGroup) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m WirelessLANGroup) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *WirelessLANGroup) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m WirelessLANGroup) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *WirelessLANGroup) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m WirelessLANGroup) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WirelessLANGroup) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m WirelessLANGroup) GetParent() *NestedWirelessLANGroup {
	return m.Parent
}

// SetParent sets the Parent property
func (m *WirelessLANGroup) SetParent(val *NestedWirelessLANGroup) {
	m.Parent = val
}

// GetSlug returns the Slug property
func (m WirelessLANGroup) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WirelessLANGroup) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m WirelessLANGroup) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WirelessLANGroup) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m WirelessLANGroup) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *WirelessLANGroup) SetUrl(val string) {
	m.Url = val
}

// GetWirelesslanCount returns the WirelesslanCount property
func (m WirelessLANGroup) GetWirelesslanCount() int32 {
	return m.WirelesslanCount
}

// SetWirelesslanCount sets the WirelesslanCount property
func (m *WirelessLANGroup) SetWirelesslanCount(val int32) {
	m.WirelesslanCount = val
}
