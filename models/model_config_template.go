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

// ConfigTemplate is an object. Introduces support for Tag assignment. Adds `tags` serialization, and handles tag assignment
// on create() and update().
type ConfigTemplate struct {
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// DataFile: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DataFile NestedDataFile `json:"data_file" mapstructure:"data_file"`
	// DataPath: Path to remote file (relative to data source root)
	DataPath string `json:"data_path" mapstructure:"data_path"`
	// DataSource: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DataSource *NestedDataSource `json:"data_source,omitempty" mapstructure:"data_source,omitempty"`
	// DataSynced:
	DataSynced *time.Time `json:"data_synced" mapstructure:"data_synced"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// EnvironmentParams: Any <a href="https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment">additional parameters</a> to pass when constructing the Jinja2 environment.
	EnvironmentParams map[string]interface{} `json:"environment_params,omitempty" mapstructure:"environment_params,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TemplateCode: Jinja2 template code.
	TemplateCode string `json:"template_code" mapstructure:"template_code"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ConfigTemplate) Validate() error {
	return validation.Errors{
		"dataFile": validation.Validate(
			m.DataFile, validation.NotNil,
		),
		"dataSource": validation.Validate(
			m.DataSource,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"environmentParams": validation.Validate(
			m.EnvironmentParams,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m ConfigTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ConfigTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDataFile returns the DataFile property
func (m ConfigTemplate) GetDataFile() NestedDataFile {
	return m.DataFile
}

// SetDataFile sets the DataFile property
func (m *ConfigTemplate) SetDataFile(val NestedDataFile) {
	m.DataFile = val
}

// GetDataPath returns the DataPath property
func (m ConfigTemplate) GetDataPath() string {
	return m.DataPath
}

// SetDataPath sets the DataPath property
func (m *ConfigTemplate) SetDataPath(val string) {
	m.DataPath = val
}

// GetDataSource returns the DataSource property
func (m ConfigTemplate) GetDataSource() *NestedDataSource {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *ConfigTemplate) SetDataSource(val *NestedDataSource) {
	m.DataSource = val
}

// GetDataSynced returns the DataSynced property
func (m ConfigTemplate) GetDataSynced() *time.Time {
	return m.DataSynced
}

// SetDataSynced sets the DataSynced property
func (m *ConfigTemplate) SetDataSynced(val *time.Time) {
	m.DataSynced = val
}

// GetDescription returns the Description property
func (m ConfigTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ConfigTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ConfigTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ConfigTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetEnvironmentParams returns the EnvironmentParams property
func (m ConfigTemplate) GetEnvironmentParams() map[string]interface{} {
	return m.EnvironmentParams
}

// SetEnvironmentParams sets the EnvironmentParams property
func (m *ConfigTemplate) SetEnvironmentParams(val map[string]interface{}) {
	m.EnvironmentParams = val
}

// GetId returns the Id property
func (m ConfigTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ConfigTemplate) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ConfigTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ConfigTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ConfigTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ConfigTemplate) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m ConfigTemplate) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ConfigTemplate) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTemplateCode returns the TemplateCode property
func (m ConfigTemplate) GetTemplateCode() string {
	return m.TemplateCode
}

// SetTemplateCode sets the TemplateCode property
func (m *ConfigTemplate) SetTemplateCode(val string) {
	m.TemplateCode = val
}

// GetUrl returns the Url property
func (m ConfigTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ConfigTemplate) SetUrl(val string) {
	m.Url = val
}
