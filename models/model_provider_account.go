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

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ProviderAccount is an object. Adds support for custom fields and tags.
type ProviderAccount struct {
	// Account:
	Account string `json:"account" mapstructure:"account"`
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
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Provider: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Provider NestedProvider `json:"provider" mapstructure:"provider"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ProviderAccount) Validate() error {
	return validation.Errors{
		"account": validation.Validate(
			m.Account, validation.NotNil, validation.Length(0, 100),
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
		"provider": validation.Validate(
			m.Provider, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAccount returns the Account property
func (m ProviderAccount) GetAccount() string {
	return m.Account
}

// SetAccount sets the Account property
func (m *ProviderAccount) SetAccount(val string) {
	m.Account = val
}

// GetComments returns the Comments property
func (m ProviderAccount) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ProviderAccount) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m ProviderAccount) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ProviderAccount) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ProviderAccount) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ProviderAccount) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ProviderAccount) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ProviderAccount) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ProviderAccount) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ProviderAccount) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ProviderAccount) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ProviderAccount) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ProviderAccount) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ProviderAccount) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ProviderAccount) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ProviderAccount) SetName(val string) {
	m.Name = val
}

// GetProvider returns the Provider property
func (m ProviderAccount) GetProvider() NestedProvider {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *ProviderAccount) SetProvider(val NestedProvider) {
	m.Provider = val
}

// GetTags returns the Tags property
func (m ProviderAccount) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ProviderAccount) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m ProviderAccount) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ProviderAccount) SetUrl(val string) {
	m.Url = val
}
