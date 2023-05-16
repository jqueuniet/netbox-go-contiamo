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

// SchemaRetrieveQueryParameters is an object.
type SchemaRetrieveQueryParameters struct {
	// Format:
	Format string `json:"format,omitempty" mapstructure:"format,omitempty"`
}

// Validate implements basic validation for this model
func (m SchemaRetrieveQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetFormat returns the Format property
func (m SchemaRetrieveQueryParameters) GetFormat() string {
	return m.Format
}

// SetFormat sets the Format property
func (m *SchemaRetrieveQueryParameters) SetFormat(val string) {
	m.Format = val
}
