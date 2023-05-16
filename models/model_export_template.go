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

// ExportTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ExportTemplate struct {
	// AsAttachment: Download file as attachment
	AsAttachment bool `json:"as_attachment,omitempty" mapstructure:"as_attachment,omitempty"`
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
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
	// FileExtension: Extension to append to the rendered filename
	FileExtension string `json:"file_extension,omitempty" mapstructure:"file_extension,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// MimeType: Defaults to <code>text/plain; charset=utf-8</code>
	MimeType string `json:"mime_type,omitempty" mapstructure:"mime_type,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// TemplateCode: Jinja2 template code. The list of objects being exported is passed as a context variable named <code>queryset</code>.
	TemplateCode string `json:"template_code" mapstructure:"template_code"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ExportTemplate) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
		),
		"dataFile": validation.Validate(
			m.DataFile, validation.NotNil,
		),
		"dataSource": validation.Validate(
			m.DataSource,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"fileExtension": validation.Validate(
			m.FileExtension, validation.Length(0, 15),
		),
		"mimeType": validation.Validate(
			m.MimeType, validation.Length(0, 50),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAsAttachment returns the AsAttachment property
func (m ExportTemplate) GetAsAttachment() bool {
	return m.AsAttachment
}

// SetAsAttachment sets the AsAttachment property
func (m *ExportTemplate) SetAsAttachment(val bool) {
	m.AsAttachment = val
}

// GetContentTypes returns the ContentTypes property
func (m ExportTemplate) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *ExportTemplate) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetCreated returns the Created property
func (m ExportTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ExportTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDataFile returns the DataFile property
func (m ExportTemplate) GetDataFile() NestedDataFile {
	return m.DataFile
}

// SetDataFile sets the DataFile property
func (m *ExportTemplate) SetDataFile(val NestedDataFile) {
	m.DataFile = val
}

// GetDataPath returns the DataPath property
func (m ExportTemplate) GetDataPath() string {
	return m.DataPath
}

// SetDataPath sets the DataPath property
func (m *ExportTemplate) SetDataPath(val string) {
	m.DataPath = val
}

// GetDataSource returns the DataSource property
func (m ExportTemplate) GetDataSource() *NestedDataSource {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *ExportTemplate) SetDataSource(val *NestedDataSource) {
	m.DataSource = val
}

// GetDataSynced returns the DataSynced property
func (m ExportTemplate) GetDataSynced() *time.Time {
	return m.DataSynced
}

// SetDataSynced sets the DataSynced property
func (m *ExportTemplate) SetDataSynced(val *time.Time) {
	m.DataSynced = val
}

// GetDescription returns the Description property
func (m ExportTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ExportTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ExportTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ExportTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetFileExtension returns the FileExtension property
func (m ExportTemplate) GetFileExtension() string {
	return m.FileExtension
}

// SetFileExtension sets the FileExtension property
func (m *ExportTemplate) SetFileExtension(val string) {
	m.FileExtension = val
}

// GetId returns the Id property
func (m ExportTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExportTemplate) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ExportTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ExportTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMimeType returns the MimeType property
func (m ExportTemplate) GetMimeType() string {
	return m.MimeType
}

// SetMimeType sets the MimeType property
func (m *ExportTemplate) SetMimeType(val string) {
	m.MimeType = val
}

// GetName returns the Name property
func (m ExportTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ExportTemplate) SetName(val string) {
	m.Name = val
}

// GetTemplateCode returns the TemplateCode property
func (m ExportTemplate) GetTemplateCode() string {
	return m.TemplateCode
}

// SetTemplateCode sets the TemplateCode property
func (m *ExportTemplate) SetTemplateCode(val string) {
	m.TemplateCode = val
}

// GetUrl returns the Url property
func (m ExportTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ExportTemplate) SetUrl(val string) {
	m.Url = val
}
