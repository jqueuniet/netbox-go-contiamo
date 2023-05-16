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

// WritableProviderNetworkRequest is an object. Adds support for custom fields and tags.
type WritableProviderNetworkRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Provider:
	Provider int32 `json:"provider" mapstructure:"provider"`
	// ServiceId:
	ServiceId string `json:"service_id,omitempty" mapstructure:"service_id,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableProviderNetworkRequest) Validate() error {
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
		"serviceId": validation.Validate(
			m.ServiceId, validation.Length(0, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableProviderNetworkRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableProviderNetworkRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableProviderNetworkRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableProviderNetworkRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableProviderNetworkRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableProviderNetworkRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m WritableProviderNetworkRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableProviderNetworkRequest) SetName(val string) {
	m.Name = val
}

// GetProvider returns the Provider property
func (m WritableProviderNetworkRequest) GetProvider() int32 {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *WritableProviderNetworkRequest) SetProvider(val int32) {
	m.Provider = val
}

// GetServiceId returns the ServiceId property
func (m WritableProviderNetworkRequest) GetServiceId() string {
	return m.ServiceId
}

// SetServiceId sets the ServiceId property
func (m *WritableProviderNetworkRequest) SetServiceId(val string) {
	m.ServiceId = val
}

// GetTags returns the Tags property
func (m WritableProviderNetworkRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableProviderNetworkRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
