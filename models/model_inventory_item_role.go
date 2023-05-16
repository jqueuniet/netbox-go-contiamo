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

	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// inventoryItemRoleColorPattern is the validation pattern for InventoryItemRole.Color
var inventoryItemRoleColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// inventoryItemRoleSlugPattern is the validation pattern for InventoryItemRole.Slug
var inventoryItemRoleSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// InventoryItemRole is an object. Adds support for custom fields and tags.
type InventoryItemRole struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
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
	// InventoryitemCount:
	InventoryitemCount int32 `json:"inventoryitem_count" mapstructure:"inventoryitem_count"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m InventoryItemRole) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(inventoryItemRoleColorPattern),
		),
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
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(inventoryItemRoleSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m InventoryItemRole) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *InventoryItemRole) SetColor(val string) {
	m.Color = val
}

// GetCreated returns the Created property
func (m InventoryItemRole) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *InventoryItemRole) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m InventoryItemRole) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *InventoryItemRole) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m InventoryItemRole) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *InventoryItemRole) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m InventoryItemRole) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *InventoryItemRole) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m InventoryItemRole) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *InventoryItemRole) SetId(val int32) {
	m.Id = val
}

// GetInventoryitemCount returns the InventoryitemCount property
func (m InventoryItemRole) GetInventoryitemCount() int32 {
	return m.InventoryitemCount
}

// SetInventoryitemCount sets the InventoryitemCount property
func (m *InventoryItemRole) SetInventoryitemCount(val int32) {
	m.InventoryitemCount = val
}

// GetLastUpdated returns the LastUpdated property
func (m InventoryItemRole) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *InventoryItemRole) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m InventoryItemRole) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *InventoryItemRole) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m InventoryItemRole) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *InventoryItemRole) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m InventoryItemRole) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *InventoryItemRole) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m InventoryItemRole) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *InventoryItemRole) SetUrl(val string) {
	m.Url = val
}
