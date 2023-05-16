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

// ExportTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ExportTemplateRequest struct {
	// AsAttachment: Download file as attachment
	AsAttachment bool `json:"as_attachment,omitempty" mapstructure:"as_attachment,omitempty"`
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// DataSource: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DataSource *NestedDataSourceRequest `json:"data_source,omitempty" mapstructure:"data_source,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// FileExtension: Extension to append to the rendered filename
	FileExtension string `json:"file_extension,omitempty" mapstructure:"file_extension,omitempty"`
	// MimeType: Defaults to <code>text/plain; charset=utf-8</code>
	MimeType string `json:"mime_type,omitempty" mapstructure:"mime_type,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// TemplateCode: Jinja2 template code. The list of objects being exported is passed as a context variable named <code>queryset</code>.
	TemplateCode string `json:"template_code" mapstructure:"template_code"`
}

// Validate implements basic validation for this model
func (m ExportTemplateRequest) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
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
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"templateCode": validation.Validate(
			m.TemplateCode, validation.Required, validation.Length(1, 0),
		),
	}.Filter()
}

// GetAsAttachment returns the AsAttachment property
func (m ExportTemplateRequest) GetAsAttachment() bool {
	return m.AsAttachment
}

// SetAsAttachment sets the AsAttachment property
func (m *ExportTemplateRequest) SetAsAttachment(val bool) {
	m.AsAttachment = val
}

// GetContentTypes returns the ContentTypes property
func (m ExportTemplateRequest) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *ExportTemplateRequest) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetDataSource returns the DataSource property
func (m ExportTemplateRequest) GetDataSource() *NestedDataSourceRequest {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *ExportTemplateRequest) SetDataSource(val *NestedDataSourceRequest) {
	m.DataSource = val
}

// GetDescription returns the Description property
func (m ExportTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ExportTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetFileExtension returns the FileExtension property
func (m ExportTemplateRequest) GetFileExtension() string {
	return m.FileExtension
}

// SetFileExtension sets the FileExtension property
func (m *ExportTemplateRequest) SetFileExtension(val string) {
	m.FileExtension = val
}

// GetMimeType returns the MimeType property
func (m ExportTemplateRequest) GetMimeType() string {
	return m.MimeType
}

// SetMimeType sets the MimeType property
func (m *ExportTemplateRequest) SetMimeType(val string) {
	m.MimeType = val
}

// GetName returns the Name property
func (m ExportTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ExportTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetTemplateCode returns the TemplateCode property
func (m ExportTemplateRequest) GetTemplateCode() string {
	return m.TemplateCode
}

// SetTemplateCode sets the TemplateCode property
func (m *ExportTemplateRequest) SetTemplateCode(val string) {
	m.TemplateCode = val
}
