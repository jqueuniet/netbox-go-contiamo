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

// NestedInterfaceTemplateRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedInterfaceTemplateRequest struct {
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
}

// Validate implements basic validation for this model
func (m NestedInterfaceTemplateRequest) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
	}.Filter()
}

// GetName returns the Name property
func (m NestedInterfaceTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedInterfaceTemplateRequest) SetName(val string) {
	m.Name = val
}
