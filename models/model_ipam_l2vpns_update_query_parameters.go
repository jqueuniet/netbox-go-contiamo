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

// IpamL2vpnsUpdateQueryParameters is an object.
type IpamL2vpnsUpdateQueryParameters struct {
	// Id: A unique integer value identifying this L2VPN.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m IpamL2vpnsUpdateQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m IpamL2vpnsUpdateQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamL2vpnsUpdateQueryParameters) SetId(val int32) {
	m.Id = val
}
