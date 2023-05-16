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

// ModuleBayNestedModuleRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type ModuleBayNestedModuleRequest struct {
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
}

// Validate implements basic validation for this model
func (m ModuleBayNestedModuleRequest) Validate() error {
	return validation.Errors{
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
	}.Filter()
}

// GetSerial returns the Serial property
func (m ModuleBayNestedModuleRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *ModuleBayNestedModuleRequest) SetSerial(val string) {
	m.Serial = val
}
