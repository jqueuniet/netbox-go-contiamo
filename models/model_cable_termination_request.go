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

// CableTerminationRequest is an object. Adds support for custom fields and tags.
type CableTerminationRequest struct {
	// Cable:
	Cable int32 `json:"cable" mapstructure:"cable"`
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end" mapstructure:"cable_end"`
	// TerminationId:
	TerminationId int64 `json:"termination_id" mapstructure:"termination_id"`
	// TerminationType:
	TerminationType string `json:"termination_type" mapstructure:"termination_type"`
}

// Validate implements basic validation for this model
func (m CableTerminationRequest) Validate() error {
	return validation.Errors{
		"terminationId": validation.Validate(
			m.TerminationId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
	}.Filter()
}

// GetCable returns the Cable property
func (m CableTerminationRequest) GetCable() int32 {
	return m.Cable
}

// SetCable sets the Cable property
func (m *CableTerminationRequest) SetCable(val int32) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m CableTerminationRequest) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *CableTerminationRequest) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetTerminationId returns the TerminationId property
func (m CableTerminationRequest) GetTerminationId() int64 {
	return m.TerminationId
}

// SetTerminationId sets the TerminationId property
func (m *CableTerminationRequest) SetTerminationId(val int64) {
	m.TerminationId = val
}

// GetTerminationType returns the TerminationType property
func (m CableTerminationRequest) GetTerminationType() string {
	return m.TerminationType
}

// SetTerminationType sets the TerminationType property
func (m *CableTerminationRequest) SetTerminationType(val string) {
	m.TerminationType = val
}
