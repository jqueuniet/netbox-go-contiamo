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

// PatchedCableTerminationRequest is an object. Adds support for custom fields and tags.
type PatchedCableTerminationRequest struct {
	// Cable:
	Cable int32 `json:"cable,omitempty" mapstructure:"cable,omitempty"`
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end,omitempty" mapstructure:"cable_end,omitempty"`
	// TerminationId:
	TerminationId int64 `json:"termination_id,omitempty" mapstructure:"termination_id,omitempty"`
	// TerminationType:
	TerminationType string `json:"termination_type,omitempty" mapstructure:"termination_type,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedCableTerminationRequest) Validate() error {
	return validation.Errors{
		"terminationId": validation.Validate(
			m.TerminationId, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
	}.Filter()
}

// GetCable returns the Cable property
func (m PatchedCableTerminationRequest) GetCable() int32 {
	return m.Cable
}

// SetCable sets the Cable property
func (m *PatchedCableTerminationRequest) SetCable(val int32) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m PatchedCableTerminationRequest) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *PatchedCableTerminationRequest) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetTerminationId returns the TerminationId property
func (m PatchedCableTerminationRequest) GetTerminationId() int64 {
	return m.TerminationId
}

// SetTerminationId sets the TerminationId property
func (m *PatchedCableTerminationRequest) SetTerminationId(val int64) {
	m.TerminationId = val
}

// GetTerminationType returns the TerminationType property
func (m PatchedCableTerminationRequest) GetTerminationType() string {
	return m.TerminationType
}

// SetTerminationType sets the TerminationType property
func (m *PatchedCableTerminationRequest) SetTerminationType(val string) {
	m.TerminationType = val
}
