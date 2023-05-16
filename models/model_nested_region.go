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

// nestedRegionSlugPattern is the validation pattern for NestedRegion.Slug
var nestedRegionSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// NestedRegion is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedRegion struct {
	// Depth:
	Depth int32 `json:"_depth" mapstructure:"_depth"`
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
func (m NestedRegion) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(nestedRegionSlugPattern),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDepth returns the Depth property
func (m NestedRegion) GetDepth() int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *NestedRegion) SetDepth(val int32) {
	m.Depth = val
}

// GetDisplay returns the Display property
func (m NestedRegion) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedRegion) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedRegion) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedRegion) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m NestedRegion) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedRegion) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m NestedRegion) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedRegion) SetSlug(val string) {
	m.Slug = val
}

// GetUrl returns the Url property
func (m NestedRegion) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedRegion) SetUrl(val string) {
	m.Url = val
}
