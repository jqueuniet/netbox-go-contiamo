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

// ExtrasCustomLinksRetrieveQueryParameters is an object.
type ExtrasCustomLinksRetrieveQueryParameters struct {
	// Id: A unique integer value identifying this custom link.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m ExtrasCustomLinksRetrieveQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m ExtrasCustomLinksRetrieveQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasCustomLinksRetrieveQueryParameters) SetId(val int32) {
	m.Id = val
}
