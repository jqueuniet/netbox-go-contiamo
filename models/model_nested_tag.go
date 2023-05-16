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

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// nestedTagColorPattern is the validation pattern for NestedTag.Color
var nestedTagColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// nestedTagSlugPattern is the validation pattern for NestedTag.Slug
var nestedTagSlugPattern = regexp.MustCompile(`^[-\w]+$`)

// NestedTag is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedTag struct {
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedTag) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(nestedTagColorPattern),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(nestedTagSlugPattern),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m NestedTag) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *NestedTag) SetColor(val string) {
	m.Color = val
}

// GetDisplay returns the Display property
func (m NestedTag) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedTag) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedTag) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedTag) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m NestedTag) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedTag) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m NestedTag) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedTag) SetSlug(val string) {
	m.Slug = val
}

// GetUrl returns the Url property
func (m NestedTag) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedTag) SetUrl(val string) {
	m.Url = val
}
