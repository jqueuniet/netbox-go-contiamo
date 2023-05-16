// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// NestedDataFile is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedDataFile struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Path: File path relative to the data source's root
	Path string `json:"path" mapstructure:"path"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedDataFile) Validate() error {
	return validation.Errors{
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedDataFile) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedDataFile) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedDataFile) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedDataFile) SetId(val int32) {
	m.Id = val
}

// GetPath returns the Path property
func (m NestedDataFile) GetPath() string {
	return m.Path
}

// SetPath sets the Path property
func (m *NestedDataFile) SetPath(val string) {
	m.Path = val
}

// GetUrl returns the Url property
func (m NestedDataFile) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedDataFile) SetUrl(val string) {
	m.Url = val
}
