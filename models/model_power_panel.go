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

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// PowerPanel is an object. Adds support for custom fields and tags.
type PowerPanel struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
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
	// Location: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Location *NestedLocation `json:"location,omitempty" mapstructure:"location,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PowerfeedCount:
	PowerfeedCount int32 `json:"powerfeed_count" mapstructure:"powerfeed_count"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site NestedSite `json:"site" mapstructure:"site"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m PowerPanel) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"location": validation.Validate(
			m.Location,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"site": validation.Validate(
			m.Site, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PowerPanel) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PowerPanel) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m PowerPanel) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *PowerPanel) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m PowerPanel) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PowerPanel) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PowerPanel) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PowerPanel) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m PowerPanel) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *PowerPanel) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m PowerPanel) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *PowerPanel) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m PowerPanel) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *PowerPanel) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLocation returns the Location property
func (m PowerPanel) GetLocation() *NestedLocation {
	return m.Location
}

// SetLocation sets the Location property
func (m *PowerPanel) SetLocation(val *NestedLocation) {
	m.Location = val
}

// GetName returns the Name property
func (m PowerPanel) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PowerPanel) SetName(val string) {
	m.Name = val
}

// GetPowerfeedCount returns the PowerfeedCount property
func (m PowerPanel) GetPowerfeedCount() int32 {
	return m.PowerfeedCount
}

// SetPowerfeedCount sets the PowerfeedCount property
func (m *PowerPanel) SetPowerfeedCount(val int32) {
	m.PowerfeedCount = val
}

// GetSite returns the Site property
func (m PowerPanel) GetSite() NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *PowerPanel) SetSite(val NestedSite) {
	m.Site = val
}

// GetTags returns the Tags property
func (m PowerPanel) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PowerPanel) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m PowerPanel) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *PowerPanel) SetUrl(val string) {
	m.Url = val
}
