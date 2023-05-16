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

// ExtrasImageAttachmentsDestroyQueryParameters is an object.
type ExtrasImageAttachmentsDestroyQueryParameters struct {
	// Id: A unique integer value identifying this image attachment.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m ExtrasImageAttachmentsDestroyQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m ExtrasImageAttachmentsDestroyQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasImageAttachmentsDestroyQueryParameters) SetId(val int32) {
	m.Id = val
}
