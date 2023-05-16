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

// PatchedWritableVirtualChassisRequest is an object. Adds support for custom fields and tags.
type PatchedWritableVirtualChassisRequest struct {
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
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableVirtualChassisRequest) Validate() error {
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
			m.Name, validation.Length(1, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableVirtualChassisRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableVirtualChassisRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableVirtualChassisRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableVirtualChassisRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableVirtualChassisRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableVirtualChassisRequest) SetDescription(val string) {
	m.Description = val
}

// GetDomain returns the Domain property
func (m PatchedWritableVirtualChassisRequest) GetDomain() string {
	return m.Domain
}

// SetDomain sets the Domain property
func (m *PatchedWritableVirtualChassisRequest) SetDomain(val string) {
	m.Domain = val
}

// GetMaster returns the Master property
func (m PatchedWritableVirtualChassisRequest) GetMaster() *int32 {
	return m.Master
}

// SetMaster sets the Master property
func (m *PatchedWritableVirtualChassisRequest) SetMaster(val *int32) {
	m.Master = val
}

// GetName returns the Name property
func (m PatchedWritableVirtualChassisRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableVirtualChassisRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m PatchedWritableVirtualChassisRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableVirtualChassisRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
