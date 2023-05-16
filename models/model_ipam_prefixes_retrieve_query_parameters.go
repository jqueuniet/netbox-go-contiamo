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

// IpamPrefixesRetrieveQueryParameters is an object.
type IpamPrefixesRetrieveQueryParameters struct {
	// Id: A unique integer value identifying this prefix.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m IpamPrefixesRetrieveQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m IpamPrefixesRetrieveQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamPrefixesRetrieveQueryParameters) SetId(val int32) {
	m.Id = val
}
