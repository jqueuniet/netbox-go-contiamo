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

// VirtualChassisRequest is an object. Adds support for custom fields and tags.
type VirtualChassisRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Domain:
	Domain string `json:"domain,omitempty" mapstructure:"domain,omitempty"`
	// Master: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Master *NestedDeviceRequest `json:"master,omitempty" mapstructure:"master,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m VirtualChassisRequest) Validate() error {
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
		"master": validation.Validate(
			m.Master,
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
func (m VirtualChassisRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VirtualChassisRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m VirtualChassisRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VirtualChassisRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VirtualChassisRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VirtualChassisRequest) SetDescription(val string) {
	m.Description = val
}

// GetDomain returns the Domain property
func (m VirtualChassisRequest) GetDomain() string {
	return m.Domain
}

// SetDomain sets the Domain property
func (m *VirtualChassisRequest) SetDomain(val string) {
	m.Domain = val
}

// GetMaster returns the Master property
func (m VirtualChassisRequest) GetMaster() *NestedDeviceRequest {
	return m.Master
}

// SetMaster sets the Master property
func (m *VirtualChassisRequest) SetMaster(val *NestedDeviceRequest) {
	m.Master = val
}

// GetName returns the Name property
func (m VirtualChassisRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualChassisRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m VirtualChassisRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VirtualChassisRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
