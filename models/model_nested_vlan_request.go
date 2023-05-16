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

// NestedVLANRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedVLANRequest struct {
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Vid: Numeric VLAN ID (1-4094)
	Vid int32 `json:"vid" mapstructure:"vid"`
}

// Validate implements basic validation for this model
func (m NestedVLANRequest) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"vid": validation.Validate(
			m.Vid, validation.Required, validation.Min(int32(1)), validation.Max(int32(4094)),
		),
	}.Filter()
}

// GetName returns the Name property
func (m NestedVLANRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedVLANRequest) SetName(val string) {
	m.Name = val
}

// GetVid returns the Vid property
func (m NestedVLANRequest) GetVid() int32 {
	return m.Vid
}

// SetVid sets the Vid property
func (m *NestedVLANRequest) SetVid(val int32) {
	m.Vid = val
}
