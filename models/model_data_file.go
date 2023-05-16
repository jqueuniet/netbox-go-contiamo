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

// DataFile is an object. Adds support for custom fields and tags.
type DataFile struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Hash: SHA256 hash of the file data
	Hash string `json:"hash" mapstructure:"hash"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Path: File path relative to the data source's root
	Path string `json:"path" mapstructure:"path"`
	// Size:
	Size int32 `json:"size" mapstructure:"size"`
	// Source: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Source NestedDataSource `json:"source" mapstructure:"source"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m DataFile) Validate() error {
	return validation.Errors{
		"source": validation.Validate(
			m.Source, validation.NotNil,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m DataFile) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *DataFile) SetDisplay(val string) {
	m.Display = val
}

// GetHash returns the Hash property
func (m DataFile) GetHash() string {
	return m.Hash
}

// SetHash sets the Hash property
func (m *DataFile) SetHash(val string) {
	m.Hash = val
}

// GetId returns the Id property
func (m DataFile) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DataFile) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m DataFile) GetLastUpdated() time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DataFile) SetLastUpdated(val time.Time) {
	m.LastUpdated = val
}

// GetPath returns the Path property
func (m DataFile) GetPath() string {
	return m.Path
}

// SetPath sets the Path property
func (m *DataFile) SetPath(val string) {
	m.Path = val
}

// GetSize returns the Size property
func (m DataFile) GetSize() int32 {
	return m.Size
}

// SetSize sets the Size property
func (m *DataFile) SetSize(val int32) {
	m.Size = val
}

// GetSource returns the Source property
func (m DataFile) GetSource() NestedDataSource {
	return m.Source
}

// SetSource sets the Source property
func (m *DataFile) SetSource(val NestedDataSource) {
	m.Source = val
}

// GetUrl returns the Url property
func (m DataFile) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *DataFile) SetUrl(val string) {
	m.Url = val
}
