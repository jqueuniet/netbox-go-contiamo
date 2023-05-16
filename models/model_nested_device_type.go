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

// nestedDeviceTypeSlugPattern is the validation pattern for NestedDeviceType.Slug
var nestedDeviceTypeSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// NestedDeviceType is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedDeviceType struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer NestedManufacturer `json:"manufacturer" mapstructure:"manufacturer"`
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedDeviceType) Validate() error {
	return validation.Errors{
		"manufacturer": validation.Validate(
			m.Manufacturer, validation.NotNil,
		),
		"model": validation.Validate(
			m.Model, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(nestedDeviceTypeSlugPattern),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedDeviceType) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedDeviceType) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedDeviceType) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedDeviceType) SetId(val int32) {
	m.Id = val
}

// GetManufacturer returns the Manufacturer property
func (m NestedDeviceType) GetManufacturer() NestedManufacturer {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *NestedDeviceType) SetManufacturer(val NestedManufacturer) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m NestedDeviceType) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *NestedDeviceType) SetModel(val string) {
	m.Model = val
}

// GetSlug returns the Slug property
func (m NestedDeviceType) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedDeviceType) SetSlug(val string) {
	m.Slug = val
}

// GetUrl returns the Url property
func (m NestedDeviceType) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedDeviceType) SetUrl(val string) {
	m.Url = val
}
