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

// DataSource is an object. Adds support for custom fields and tags.
type DataSource struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// FileCount:
	FileCount int32 `json:"file_count" mapstructure:"file_count"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IgnoreRules: Patterns (one per line) matching files to ignore when syncing
	IgnoreRules string `json:"ignore_rules,omitempty" mapstructure:"ignore_rules,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parameters:
	Parameters map[string]interface{} `json:"parameters,omitempty" mapstructure:"parameters,omitempty"`
	// SourceUrl:
	SourceUrl string `json:"source_url" mapstructure:"source_url"`
	// Status:
	Status struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status" mapstructure:"status"`
	// Type:
	Type struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m DataSource) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"parameters": validation.Validate(
			m.Parameters,
		),
		"sourceUrl": validation.Validate(
			m.SourceUrl, validation.NotNil, validation.Length(0, 200),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m DataSource) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *DataSource) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m DataSource) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DataSource) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m DataSource) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DataSource) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m DataSource) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *DataSource) SetDisplay(val string) {
	m.Display = val
}

// GetEnabled returns the Enabled property
func (m DataSource) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *DataSource) SetEnabled(val bool) {
	m.Enabled = val
}

// GetFileCount returns the FileCount property
func (m DataSource) GetFileCount() int32 {
	return m.FileCount
}

// SetFileCount sets the FileCount property
func (m *DataSource) SetFileCount(val int32) {
	m.FileCount = val
}

// GetId returns the Id property
func (m DataSource) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DataSource) SetId(val int32) {
	m.Id = val
}

// GetIgnoreRules returns the IgnoreRules property
func (m DataSource) GetIgnoreRules() string {
	return m.IgnoreRules
}

// SetIgnoreRules sets the IgnoreRules property
func (m *DataSource) SetIgnoreRules(val string) {
	m.IgnoreRules = val
}

// GetLastUpdated returns the LastUpdated property
func (m DataSource) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DataSource) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m DataSource) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *DataSource) SetName(val string) {
	m.Name = val
}

// GetParameters returns the Parameters property
func (m DataSource) GetParameters() map[string]interface{} {
	return m.Parameters
}

// SetParameters sets the Parameters property
func (m *DataSource) SetParameters(val map[string]interface{}) {
	m.Parameters = val
}

// GetSourceUrl returns the SourceUrl property
func (m DataSource) GetSourceUrl() string {
	return m.SourceUrl
}

// SetSourceUrl sets the SourceUrl property
func (m *DataSource) SetSourceUrl(val string) {
	m.SourceUrl = val
}

// GetStatus returns the Status property
func (m DataSource) GetStatus() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *DataSource) SetStatus(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetType returns the Type property
func (m DataSource) GetType() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *DataSource) SetType(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m DataSource) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *DataSource) SetUrl(val string) {
	m.Url = val
}
