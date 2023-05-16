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

// NestedDeviceRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedDeviceRequest struct {
	// Name:
	Name *string `json:"name,omitempty" mapstructure:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m NestedDeviceRequest) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Length(0, 64),
		),
	}.Filter()
}

// GetName returns the Name property
func (m NestedDeviceRequest) GetName() *string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedDeviceRequest) SetName(val *string) {
	m.Name = val
}
