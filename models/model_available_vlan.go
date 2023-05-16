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

// AvailableVLAN is an object. Representation of a VLAN which does not exist in the database.
type AvailableVLAN struct {
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group NestedVLANGroup `json:"group" mapstructure:"group"`
	// Vid:
	Vid int32 `json:"vid" mapstructure:"vid"`
}

// Validate implements basic validation for this model
func (m AvailableVLAN) Validate() error {
	return validation.Errors{
		"group": validation.Validate(
			m.Group, validation.NotNil,
		),
	}.Filter()
}

// GetGroup returns the Group property
func (m AvailableVLAN) GetGroup() NestedVLANGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *AvailableVLAN) SetGroup(val NestedVLANGroup) {
	m.Group = val
}

// GetVid returns the Vid property
func (m AvailableVLAN) GetVid() int32 {
	return m.Vid
}

// SetVid sets the Vid property
func (m *AvailableVLAN) SetVid(val int32) {
	m.Vid = val
}
