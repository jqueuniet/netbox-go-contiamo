// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// providerSlugPattern is the validation pattern for Provider.Slug
var providerSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// Provider is an object. Adds support for custom fields and tags.
type Provider struct {
	// Accounts:
	Accounts []int32 `json:"accounts,omitempty" mapstructure:"accounts,omitempty"`
	// Asns:
	Asns []int32 `json:"asns,omitempty" mapstructure:"asns,omitempty"`
	// CircuitCount:
	CircuitCount int32 `json:"circuit_count" mapstructure:"circuit_count"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name: Full name of the provider
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Provider) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(providerSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAccounts returns the Accounts property
func (m Provider) GetAccounts() []int32 {
	return m.Accounts
}

// SetAccounts sets the Accounts property
func (m *Provider) SetAccounts(val []int32) {
	m.Accounts = val
}

// GetAsns returns the Asns property
func (m Provider) GetAsns() []int32 {
	return m.Asns
}

// SetAsns sets the Asns property
func (m *Provider) SetAsns(val []int32) {
	m.Asns = val
}

// GetCircuitCount returns the CircuitCount property
func (m Provider) GetCircuitCount() int32 {
	return m.CircuitCount
}

// SetCircuitCount sets the CircuitCount property
func (m *Provider) SetCircuitCount(val int32) {
	m.CircuitCount = val
}

// GetComments returns the Comments property
func (m Provider) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Provider) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Provider) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Provider) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Provider) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Provider) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Provider) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Provider) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Provider) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Provider) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Provider) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Provider) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Provider) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Provider) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Provider) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Provider) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m Provider) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *Provider) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m Provider) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Provider) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m Provider) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Provider) SetUrl(val string) {
	m.Url = val
}
