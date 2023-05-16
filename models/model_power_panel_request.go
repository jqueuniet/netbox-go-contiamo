// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// PowerPanelRequest is an object. Adds support for custom fields and tags.
type PowerPanelRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Location: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Location *NestedLocationRequest `json:"location,omitempty" mapstructure:"location,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site NestedSiteRequest `json:"site" mapstructure:"site"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PowerPanelRequest) Validate() error {
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
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"site": validation.Validate(
			m.Site, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PowerPanelRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PowerPanelRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PowerPanelRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PowerPanelRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PowerPanelRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PowerPanelRequest) SetDescription(val string) {
	m.Description = val
}

// GetLocation returns the Location property
func (m PowerPanelRequest) GetLocation() *NestedLocationRequest {
	return m.Location
}

// SetLocation sets the Location property
func (m *PowerPanelRequest) SetLocation(val *NestedLocationRequest) {
	m.Location = val
}

// GetName returns the Name property
func (m PowerPanelRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PowerPanelRequest) SetName(val string) {
	m.Name = val
}

// GetSite returns the Site property
func (m PowerPanelRequest) GetSite() NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *PowerPanelRequest) SetSite(val NestedSiteRequest) {
	m.Site = val
}

// GetTags returns the Tags property
func (m PowerPanelRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PowerPanelRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
