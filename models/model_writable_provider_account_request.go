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

// WritableProviderAccountRequest is an object. Adds support for custom fields and tags.
type WritableProviderAccountRequest struct {
	// Account:
	Account string `json:"account" mapstructure:"account"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Provider:
	Provider int32 `json:"provider" mapstructure:"provider"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableProviderAccountRequest) Validate() error {
	return validation.Errors{
		"account": validation.Validate(
			m.Account, validation.Required, validation.Length(1, 100),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAccount returns the Account property
func (m WritableProviderAccountRequest) GetAccount() string {
	return m.Account
}

// SetAccount sets the Account property
func (m *WritableProviderAccountRequest) SetAccount(val string) {
	m.Account = val
}

// GetComments returns the Comments property
func (m WritableProviderAccountRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableProviderAccountRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableProviderAccountRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableProviderAccountRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableProviderAccountRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableProviderAccountRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m WritableProviderAccountRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableProviderAccountRequest) SetName(val string) {
	m.Name = val
}

// GetProvider returns the Provider property
func (m WritableProviderAccountRequest) GetProvider() int32 {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *WritableProviderAccountRequest) SetProvider(val int32) {
	m.Provider = val
}

// GetTags returns the Tags property
func (m WritableProviderAccountRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableProviderAccountRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
