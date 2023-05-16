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

// IpamIpRangesPartialUpdateQueryParameters is an object.
type IpamIpRangesPartialUpdateQueryParameters struct {
	// Id: A unique integer value identifying this IP range.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m IpamIpRangesPartialUpdateQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m IpamIpRangesPartialUpdateQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamIpRangesPartialUpdateQueryParameters) SetId(val int32) {
	m.Id = val
}
