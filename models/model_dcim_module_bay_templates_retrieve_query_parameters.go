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

// DcimModuleBayTemplatesRetrieveQueryParameters is an object.
type DcimModuleBayTemplatesRetrieveQueryParameters struct {
	// Id: A unique integer value identifying this module bay template.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m DcimModuleBayTemplatesRetrieveQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m DcimModuleBayTemplatesRetrieveQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimModuleBayTemplatesRetrieveQueryParameters) SetId(val int32) {
	m.Id = val
}
