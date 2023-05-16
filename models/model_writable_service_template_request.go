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

// WritableServiceTemplateRequest is an object. Adds support for custom fields and tags.
type WritableServiceTemplateRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Ports:
	Ports []int32 `json:"ports" mapstructure:"ports"`
	// Protocol: * `tcp` - TCP
	// * `udp` - UDP
	// * `sctp` - SCTP
	Protocol string `json:"protocol" mapstructure:"protocol"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableServiceTemplateRequest) Validate() error {
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
		"ports": validation.Validate(
			m.Ports, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableServiceTemplateRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableServiceTemplateRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableServiceTemplateRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableServiceTemplateRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableServiceTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableServiceTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m WritableServiceTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableServiceTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetPorts returns the Ports property
func (m WritableServiceTemplateRequest) GetPorts() []int32 {
	return m.Ports
}

// SetPorts sets the Ports property
func (m *WritableServiceTemplateRequest) SetPorts(val []int32) {
	m.Ports = val
}

// GetProtocol returns the Protocol property
func (m WritableServiceTemplateRequest) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *WritableServiceTemplateRequest) SetProtocol(val string) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m WritableServiceTemplateRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableServiceTemplateRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
