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

// WritableExportTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableExportTemplateRequest struct {
	// AsAttachment: Download file as attachment
	AsAttachment bool `json:"as_attachment,omitempty" mapstructure:"as_attachment,omitempty"`
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// DataSource: Remote data source
	DataSource *int32 `json:"data_source,omitempty" mapstructure:"data_source,omitempty"`
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
func (m WritableExportTemplateRequest) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
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
func (m WritableExportTemplateRequest) GetAsAttachment() bool {
	return m.AsAttachment
}

// SetAsAttachment sets the AsAttachment property
func (m *WritableExportTemplateRequest) SetAsAttachment(val bool) {
	m.AsAttachment = val
}

// GetContentTypes returns the ContentTypes property
func (m WritableExportTemplateRequest) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *WritableExportTemplateRequest) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetDataSource returns the DataSource property
func (m WritableExportTemplateRequest) GetDataSource() *int32 {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *WritableExportTemplateRequest) SetDataSource(val *int32) {
	m.DataSource = val
}

// GetDescription returns the Description property
func (m WritableExportTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableExportTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetFileExtension returns the FileExtension property
func (m WritableExportTemplateRequest) GetFileExtension() string {
	return m.FileExtension
}

// SetFileExtension sets the FileExtension property
func (m *WritableExportTemplateRequest) SetFileExtension(val string) {
	m.FileExtension = val
}

// GetMimeType returns the MimeType property
func (m WritableExportTemplateRequest) GetMimeType() string {
	return m.MimeType
}

// SetMimeType sets the MimeType property
func (m *WritableExportTemplateRequest) SetMimeType(val string) {
	m.MimeType = val
}

// GetName returns the Name property
func (m WritableExportTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableExportTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetTemplateCode returns the TemplateCode property
func (m WritableExportTemplateRequest) GetTemplateCode() string {
	return m.TemplateCode
}

// SetTemplateCode sets the TemplateCode property
func (m *WritableExportTemplateRequest) SetTemplateCode(val string) {
	m.TemplateCode = val
}
