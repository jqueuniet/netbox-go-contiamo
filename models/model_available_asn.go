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

// AvailableASN is an object. Representation of an ASN which does not exist in the database.
type AvailableASN struct {
	// Asn:
	Asn int32 `json:"asn" mapstructure:"asn"`
}

// Validate implements basic validation for this model
func (m AvailableASN) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAsn returns the Asn property
func (m AvailableASN) GetAsn() int32 {
	return m.Asn
}

// SetAsn sets the Asn property
func (m *AvailableASN) SetAsn(val int32) {
	m.Asn = val
}
