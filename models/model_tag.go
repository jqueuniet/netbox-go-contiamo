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

// tagColorPattern is the validation pattern for Tag.Color
var tagColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// tagSlugPattern is the validation pattern for Tag.Slug
var tagSlugPattern = regexp.MustCompile(`^[-\w]+$`)

// Tag is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type Tag struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
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
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// TaggedItems:
	TaggedItems int32 `json:"tagged_items" mapstructure:"tagged_items"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Tag) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(tagColorPattern),
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(tagSlugPattern),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m Tag) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *Tag) SetColor(val string) {
	m.Color = val
}

// GetCreated returns the Created property
func (m Tag) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Tag) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m Tag) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Tag) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Tag) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Tag) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Tag) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Tag) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Tag) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Tag) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Tag) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Tag) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m Tag) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Tag) SetSlug(val string) {
	m.Slug = val
}

// GetTaggedItems returns the TaggedItems property
func (m Tag) GetTaggedItems() int32 {
	return m.TaggedItems
}

// SetTaggedItems sets the TaggedItems property
func (m *Tag) SetTaggedItems(val int32) {
	m.TaggedItems = val
}

// GetUrl returns the Url property
func (m Tag) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Tag) SetUrl(val string) {
	m.Url = val
}
