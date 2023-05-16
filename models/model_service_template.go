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

// ServiceTemplate is an object. Adds support for custom fields and tags.
type ServiceTemplate struct {
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
	// Ports:
	Ports []int32 `json:"ports" mapstructure:"ports"`
	// Protocol:
	Protocol *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ServiceTemplate) Validate() error {
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
		"ports": validation.Validate(
			m.Ports, validation.NotNil,
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
func (m ServiceTemplate) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ServiceTemplate) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m ServiceTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ServiceTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ServiceTemplate) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ServiceTemplate) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ServiceTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ServiceTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ServiceTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ServiceTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ServiceTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ServiceTemplate) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ServiceTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ServiceTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ServiceTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ServiceTemplate) SetName(val string) {
	m.Name = val
}

// GetPorts returns the Ports property
func (m ServiceTemplate) GetPorts() []int32 {
	return m.Ports
}

// SetPorts sets the Ports property
func (m *ServiceTemplate) SetPorts(val []int32) {
	m.Ports = val
}

// GetProtocol returns the Protocol property
func (m ServiceTemplate) GetProtocol() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *ServiceTemplate) SetProtocol(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m ServiceTemplate) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ServiceTemplate) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m ServiceTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ServiceTemplate) SetUrl(val string) {
	m.Url = val
}
