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

// NestedL2VPNTerminationRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedL2VPNTerminationRequest struct {
	// L2vpn: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	L2vpn NestedL2VPNRequest `json:"l2vpn" mapstructure:"l2vpn"`
}

// Validate implements basic validation for this model
func (m NestedL2VPNTerminationRequest) Validate() error {
	return validation.Errors{
		"l2vpn": validation.Validate(
			m.L2vpn, validation.NotNil,
		),
	}.Filter()
}

// GetL2vpn returns the L2vpn property
func (m NestedL2VPNTerminationRequest) GetL2vpn() NestedL2VPNRequest {
	return m.L2vpn
}

// SetL2vpn sets the L2vpn property
func (m *NestedL2VPNTerminationRequest) SetL2vpn(val NestedL2VPNRequest) {
	m.L2vpn = val
}
