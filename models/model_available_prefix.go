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

// AvailablePrefix is an object. Representation of a prefix which does not exist in the database.
type AvailablePrefix struct {
	// Family:
	Family int32 `json:"family" mapstructure:"family"`
	// Prefix:
	Prefix string `json:"prefix" mapstructure:"prefix"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf NestedVRF `json:"vrf" mapstructure:"vrf"`
}

// Validate implements basic validation for this model
func (m AvailablePrefix) Validate() error {
	return validation.Errors{
		"vrf": validation.Validate(
			m.Vrf, validation.NotNil,
		),
	}.Filter()
}

// GetFamily returns the Family property
func (m AvailablePrefix) GetFamily() int32 {
	return m.Family
}

// SetFamily sets the Family property
func (m *AvailablePrefix) SetFamily(val int32) {
	m.Family = val
}

// GetPrefix returns the Prefix property
func (m AvailablePrefix) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *AvailablePrefix) SetPrefix(val string) {
	m.Prefix = val
}

// GetVrf returns the Vrf property
func (m AvailablePrefix) GetVrf() NestedVRF {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *AvailablePrefix) SetVrf(val NestedVRF) {
	m.Vrf = val
}
