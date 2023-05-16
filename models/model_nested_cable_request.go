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

// NestedCableRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedCableRequest struct {
	// Label:
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
}

// Validate implements basic validation for this model
func (m NestedCableRequest) Validate() error {
	return validation.Errors{
		"label": validation.Validate(
			m.Label, validation.Length(0, 100),
		),
	}.Filter()
}

// GetLabel returns the Label property
func (m NestedCableRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *NestedCableRequest) SetLabel(val string) {
	m.Label = val
}
