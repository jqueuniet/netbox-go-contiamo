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

// PatchedWritableConfigTemplateRequest is an object. Introduces support for Tag assignment. Adds `tags` serialization, and handles tag assignment
// on create() and update().
type PatchedWritableConfigTemplateRequest struct {
	// DataSource: Remote data source
	DataSource *int32 `json:"data_source,omitempty" mapstructure:"data_source,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// EnvironmentParams: Any <a href="https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment">additional parameters</a> to pass when constructing the Jinja2 environment.
	EnvironmentParams map[string]interface{} `json:"environment_params,omitempty" mapstructure:"environment_params,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TemplateCode: Jinja2 template code.
	TemplateCode string `json:"template_code,omitempty" mapstructure:"template_code,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableConfigTemplateRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"environmentParams": validation.Validate(
			m.EnvironmentParams,
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"templateCode": validation.Validate(
			m.TemplateCode, validation.Length(1, 0),
		),
	}.Filter()
}

// GetDataSource returns the DataSource property
func (m PatchedWritableConfigTemplateRequest) GetDataSource() *int32 {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *PatchedWritableConfigTemplateRequest) SetDataSource(val *int32) {
	m.DataSource = val
}

// GetDescription returns the Description property
func (m PatchedWritableConfigTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableConfigTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnvironmentParams returns the EnvironmentParams property
func (m PatchedWritableConfigTemplateRequest) GetEnvironmentParams() map[string]interface{} {
	return m.EnvironmentParams
}

// SetEnvironmentParams sets the EnvironmentParams property
func (m *PatchedWritableConfigTemplateRequest) SetEnvironmentParams(val map[string]interface{}) {
	m.EnvironmentParams = val
}

// GetName returns the Name property
func (m PatchedWritableConfigTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableConfigTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m PatchedWritableConfigTemplateRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableConfigTemplateRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTemplateCode returns the TemplateCode property
func (m PatchedWritableConfigTemplateRequest) GetTemplateCode() string {
	return m.TemplateCode
}

// SetTemplateCode sets the TemplateCode property
func (m *PatchedWritableConfigTemplateRequest) SetTemplateCode(val string) {
	m.TemplateCode = val
}
