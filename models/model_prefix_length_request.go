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

// PrefixLengthRequest is an object.
type PrefixLengthRequest struct {
	// PrefixLength:
	PrefixLength int32 `json:"prefix_length" mapstructure:"prefix_length"`
}

// Validate implements basic validation for this model
func (m PrefixLengthRequest) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPrefixLength returns the PrefixLength property
func (m PrefixLengthRequest) GetPrefixLength() int32 {
	return m.PrefixLength
}

// SetPrefixLength sets the PrefixLength property
func (m *PrefixLengthRequest) SetPrefixLength(val int32) {
	m.PrefixLength = val
}
