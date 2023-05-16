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

// NestedModuleTypeRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedModuleTypeRequest struct {
	// Model:
	Model string `json:"model" mapstructure:"model"`
}

// Validate implements basic validation for this model
func (m NestedModuleTypeRequest) Validate() error {
	return validation.Errors{
		"model": validation.Validate(
			m.Model, validation.Required, validation.Length(1, 100),
		),
	}.Filter()
}

// GetModel returns the Model property
func (m NestedModuleTypeRequest) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *NestedModuleTypeRequest) SetModel(val string) {
	m.Model = val
}
