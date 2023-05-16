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

// manufacturerSlugPattern is the validation pattern for Manufacturer.Slug
var manufacturerSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// Manufacturer is an object. Adds support for custom fields and tags.
type Manufacturer struct {
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DevicetypeCount:
	DevicetypeCount int32 `json:"devicetype_count" mapstructure:"devicetype_count"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// InventoryitemCount:
	InventoryitemCount int32 `json:"inventoryitem_count" mapstructure:"inventoryitem_count"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PlatformCount:
	PlatformCount int32 `json:"platform_count" mapstructure:"platform_count"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Manufacturer) Validate() error {
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(manufacturerSlugPattern),
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
func (m Manufacturer) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Manufacturer) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Manufacturer) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Manufacturer) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Manufacturer) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Manufacturer) SetDescription(val string) {
	m.Description = val
}

// GetDevicetypeCount returns the DevicetypeCount property
func (m Manufacturer) GetDevicetypeCount() int32 {
	return m.DevicetypeCount
}

// SetDevicetypeCount sets the DevicetypeCount property
func (m *Manufacturer) SetDevicetypeCount(val int32) {
	m.DevicetypeCount = val
}

// GetDisplay returns the Display property
func (m Manufacturer) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Manufacturer) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Manufacturer) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Manufacturer) SetId(val int32) {
	m.Id = val
}

// GetInventoryitemCount returns the InventoryitemCount property
func (m Manufacturer) GetInventoryitemCount() int32 {
	return m.InventoryitemCount
}

// SetInventoryitemCount sets the InventoryitemCount property
func (m *Manufacturer) SetInventoryitemCount(val int32) {
	m.InventoryitemCount = val
}

// GetLastUpdated returns the LastUpdated property
func (m Manufacturer) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Manufacturer) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Manufacturer) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Manufacturer) SetName(val string) {
	m.Name = val
}

// GetPlatformCount returns the PlatformCount property
func (m Manufacturer) GetPlatformCount() int32 {
	return m.PlatformCount
}

// SetPlatformCount sets the PlatformCount property
func (m *Manufacturer) SetPlatformCount(val int32) {
	m.PlatformCount = val
}

// GetSlug returns the Slug property
func (m Manufacturer) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Manufacturer) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m Manufacturer) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Manufacturer) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m Manufacturer) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Manufacturer) SetUrl(val string) {
	m.Url = val
}
