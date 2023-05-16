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

// ProviderNetwork is an object. Adds support for custom fields and tags.
type ProviderNetwork struct {
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
	Name string `json:"name" mapstructure:"name"`
	// Provider: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Provider NestedProvider `json:"provider" mapstructure:"provider"`
	// ServiceId:
	ServiceId string `json:"service_id,omitempty" mapstructure:"service_id,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ProviderNetwork) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"provider": validation.Validate(
			m.Provider, validation.NotNil,
		),
		"serviceId": validation.Validate(
			m.ServiceId, validation.Length(0, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m ProviderNetwork) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ProviderNetwork) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m ProviderNetwork) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ProviderNetwork) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ProviderNetwork) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ProviderNetwork) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ProviderNetwork) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ProviderNetwork) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ProviderNetwork) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ProviderNetwork) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ProviderNetwork) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ProviderNetwork) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ProviderNetwork) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ProviderNetwork) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ProviderNetwork) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ProviderNetwork) SetName(val string) {
	m.Name = val
}

// GetProvider returns the Provider property
func (m ProviderNetwork) GetProvider() NestedProvider {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *ProviderNetwork) SetProvider(val NestedProvider) {
	m.Provider = val
}

// GetServiceId returns the ServiceId property
func (m ProviderNetwork) GetServiceId() string {
	return m.ServiceId
}

// SetServiceId sets the ServiceId property
func (m *ProviderNetwork) SetServiceId(val string) {
	m.ServiceId = val
}

// GetTags returns the Tags property
func (m ProviderNetwork) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ProviderNetwork) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m ProviderNetwork) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ProviderNetwork) SetUrl(val string) {
	m.Url = val
}
