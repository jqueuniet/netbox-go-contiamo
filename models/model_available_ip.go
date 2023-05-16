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

// AvailableIP is an object. Representation of an IP address which does not exist in the database.
type AvailableIP struct {
	// Address:
	Address string `json:"address" mapstructure:"address"`
	// Family:
	Family int32 `json:"family" mapstructure:"family"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf NestedVRF `json:"vrf" mapstructure:"vrf"`
}

// Validate implements basic validation for this model
func (m AvailableIP) Validate() error {
	return validation.Errors{
		"vrf": validation.Validate(
			m.Vrf, validation.NotNil,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m AvailableIP) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *AvailableIP) SetAddress(val string) {
	m.Address = val
}

// GetFamily returns the Family property
func (m AvailableIP) GetFamily() int32 {
	return m.Family
}

// SetFamily sets the Family property
func (m *AvailableIP) SetFamily(val int32) {
	m.Family = val
}

// GetVrf returns the Vrf property
func (m AvailableIP) GetVrf() NestedVRF {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *AvailableIP) SetVrf(val NestedVRF) {
	m.Vrf = val
}
