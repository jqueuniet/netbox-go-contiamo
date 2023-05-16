// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// providerRequestSlugPattern is the validation pattern for ProviderRequest.Slug
var providerRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ProviderRequest is an object. Adds support for custom fields and tags.
type ProviderRequest struct {
	// Accounts:
	Accounts []int32 `json:"accounts,omitempty" mapstructure:"accounts,omitempty"`
	// Asns:
	Asns []int32 `json:"asns,omitempty" mapstructure:"asns,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Name: Full name of the provider
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m ProviderRequest) Validate() error {
	return validation.Errors{
		"accounts": validation.Validate(
			m.Accounts,
		),
		"asns": validation.Validate(
			m.Asns,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(providerRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAccounts returns the Accounts property
func (m ProviderRequest) GetAccounts() []int32 {
	return m.Accounts
}

// SetAccounts sets the Accounts property
func (m *ProviderRequest) SetAccounts(val []int32) {
	m.Accounts = val
}

// GetAsns returns the Asns property
func (m ProviderRequest) GetAsns() []int32 {
	return m.Asns
}

// SetAsns sets the Asns property
func (m *ProviderRequest) SetAsns(val []int32) {
	m.Asns = val
}

// GetComments returns the Comments property
func (m ProviderRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ProviderRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m ProviderRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ProviderRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ProviderRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ProviderRequest) SetDescription(val string) {
	m.Description = val
}

// GetName returns the Name property
func (m ProviderRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ProviderRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m ProviderRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ProviderRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m ProviderRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ProviderRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
