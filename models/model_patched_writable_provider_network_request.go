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

// PatchedWritableProviderNetworkRequest is an object. Adds support for custom fields and tags.
type PatchedWritableProviderNetworkRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Provider:
	Provider int32 `json:"provider,omitempty" mapstructure:"provider,omitempty"`
	// ServiceId:
	ServiceId string `json:"service_id,omitempty" mapstructure:"service_id,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableProviderNetworkRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
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
func (m PatchedWritableProviderNetworkRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableProviderNetworkRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableProviderNetworkRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableProviderNetworkRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableProviderNetworkRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableProviderNetworkRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m PatchedWritableProviderNetworkRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableProviderNetworkRequest) SetName(val string) {
	m.Name = val
}

// GetProvider returns the Provider property
func (m PatchedWritableProviderNetworkRequest) GetProvider() int32 {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *PatchedWritableProviderNetworkRequest) SetProvider(val int32) {
	m.Provider = val
}

// GetServiceId returns the ServiceId property
func (m PatchedWritableProviderNetworkRequest) GetServiceId() string {
	return m.ServiceId
}

// SetServiceId sets the ServiceId property
func (m *PatchedWritableProviderNetworkRequest) SetServiceId(val string) {
	m.ServiceId = val
}

// GetTags returns the Tags property
func (m PatchedWritableProviderNetworkRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableProviderNetworkRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
