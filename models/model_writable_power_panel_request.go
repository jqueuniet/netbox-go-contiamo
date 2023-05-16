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

// WritablePowerPanelRequest is an object. Adds support for custom fields and tags.
type WritablePowerPanelRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Location:
	Location *int32 `json:"location,omitempty" mapstructure:"location,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Site:
	Site int32 `json:"site" mapstructure:"site"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritablePowerPanelRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritablePowerPanelRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritablePowerPanelRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritablePowerPanelRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritablePowerPanelRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritablePowerPanelRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritablePowerPanelRequest) SetDescription(val string) {
	m.Description = val
}

// GetLocation returns the Location property
func (m WritablePowerPanelRequest) GetLocation() *int32 {
	return m.Location
}

// SetLocation sets the Location property
func (m *WritablePowerPanelRequest) SetLocation(val *int32) {
	m.Location = val
}

// GetName returns the Name property
func (m WritablePowerPanelRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritablePowerPanelRequest) SetName(val string) {
	m.Name = val
}

// GetSite returns the Site property
func (m WritablePowerPanelRequest) GetSite() int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *WritablePowerPanelRequest) SetSite(val int32) {
	m.Site = val
}

// GetTags returns the Tags property
func (m WritablePowerPanelRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritablePowerPanelRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
