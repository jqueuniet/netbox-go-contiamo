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

// DcimDeviceBaysRetrieveQueryParameters is an object.
type DcimDeviceBaysRetrieveQueryParameters struct {
	// Id: A unique integer value identifying this device bay.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m DcimDeviceBaysRetrieveQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m DcimDeviceBaysRetrieveQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimDeviceBaysRetrieveQueryParameters) SetId(val int32) {
	m.Id = val
}
