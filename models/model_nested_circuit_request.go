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

// NestedCircuitRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedCircuitRequest struct {
	// Cid: Unique circuit ID
	Cid string `json:"cid" mapstructure:"cid"`
}

// Validate implements basic validation for this model
func (m NestedCircuitRequest) Validate() error {
	return validation.Errors{
		"cid": validation.Validate(
			m.Cid, validation.Required, validation.Length(1, 100),
		),
	}.Filter()
}

// GetCid returns the Cid property
func (m NestedCircuitRequest) GetCid() string {
	return m.Cid
}

// SetCid sets the Cid property
func (m *NestedCircuitRequest) SetCid(val string) {
	m.Cid = val
}
