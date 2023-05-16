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

// DcimPowerOutletTemplatesUpdateQueryParameters is an object.
type DcimPowerOutletTemplatesUpdateQueryParameters struct {
	// Id: A unique integer value identifying this power outlet template.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m DcimPowerOutletTemplatesUpdateQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m DcimPowerOutletTemplatesUpdateQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimPowerOutletTemplatesUpdateQueryParameters) SetId(val int32) {
	m.Id = val
}
