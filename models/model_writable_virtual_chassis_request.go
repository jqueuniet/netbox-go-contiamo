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

// WritableVirtualChassisRequest is an object. Adds support for custom fields and tags.
type WritableVirtualChassisRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Domain:
	Domain string `json:"domain,omitempty" mapstructure:"domain,omitempty"`
	// Master:
	Master *int32 `json:"master,omitempty" mapstructure:"master,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableVirtualChassisRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"domain": validation.Validate(
			m.Domain, validation.Length(0, 30),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableVirtualChassisRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableVirtualChassisRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableVirtualChassisRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableVirtualChassisRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableVirtualChassisRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableVirtualChassisRequest) SetDescription(val string) {
	m.Description = val
}

// GetDomain returns the Domain property
func (m WritableVirtualChassisRequest) GetDomain() string {
	return m.Domain
}

// SetDomain sets the Domain property
func (m *WritableVirtualChassisRequest) SetDomain(val string) {
	m.Domain = val
}

// GetMaster returns the Master property
func (m WritableVirtualChassisRequest) GetMaster() *int32 {
	return m.Master
}

// SetMaster sets the Master property
func (m *WritableVirtualChassisRequest) SetMaster(val *int32) {
	m.Master = val
}

// GetName returns the Name property
func (m WritableVirtualChassisRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableVirtualChassisRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m WritableVirtualChassisRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableVirtualChassisRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
