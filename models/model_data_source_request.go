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

// DataSourceRequest is an object. Adds support for custom fields and tags.
type DataSourceRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// IgnoreRules: Patterns (one per line) matching files to ignore when syncing
	IgnoreRules string `json:"ignore_rules,omitempty" mapstructure:"ignore_rules,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parameters:
	Parameters map[string]interface{} `json:"parameters,omitempty" mapstructure:"parameters,omitempty"`
	// SourceUrl:
	SourceUrl string `json:"source_url" mapstructure:"source_url"`
	// Type: * `local` - Local
	// * `git` - Git
	// * `amazon-s3` - Amazon S3
	Type string `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m DataSourceRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"parameters": validation.Validate(
			m.Parameters,
		),
		"sourceUrl": validation.Validate(
			m.SourceUrl, validation.Required, validation.Length(1, 200),
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m DataSourceRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *DataSourceRequest) SetComments(val string) {
	m.Comments = val
}

// GetDescription returns the Description property
func (m DataSourceRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DataSourceRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnabled returns the Enabled property
func (m DataSourceRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *DataSourceRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetIgnoreRules returns the IgnoreRules property
func (m DataSourceRequest) GetIgnoreRules() string {
	return m.IgnoreRules
}

// SetIgnoreRules sets the IgnoreRules property
func (m *DataSourceRequest) SetIgnoreRules(val string) {
	m.IgnoreRules = val
}

// GetName returns the Name property
func (m DataSourceRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *DataSourceRequest) SetName(val string) {
	m.Name = val
}

// GetParameters returns the Parameters property
func (m DataSourceRequest) GetParameters() map[string]interface{} {
	return m.Parameters
}

// SetParameters sets the Parameters property
func (m *DataSourceRequest) SetParameters(val map[string]interface{}) {
	m.Parameters = val
}

// GetSourceUrl returns the SourceUrl property
func (m DataSourceRequest) GetSourceUrl() string {
	return m.SourceUrl
}

// SetSourceUrl sets the SourceUrl property
func (m *DataSourceRequest) SetSourceUrl(val string) {
	m.SourceUrl = val
}

// GetType returns the Type property
func (m DataSourceRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *DataSourceRequest) SetType(val string) {
	m.Type = val
}
