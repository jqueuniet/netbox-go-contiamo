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

// NestedIPAddressRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedIPAddressRequest struct {
	// Address:
	Address string `json:"address" mapstructure:"address"`
}

// Validate implements basic validation for this model
func (m NestedIPAddressRequest) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address, validation.Required, validation.Length(1, 0),
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m NestedIPAddressRequest) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *NestedIPAddressRequest) SetAddress(val string) {
	m.Address = val
}
