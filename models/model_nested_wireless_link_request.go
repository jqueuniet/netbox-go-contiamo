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

// NestedWirelessLinkRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedWirelessLinkRequest struct {
	// Ssid:
	Ssid string `json:"ssid,omitempty" mapstructure:"ssid,omitempty"`
}

// Validate implements basic validation for this model
func (m NestedWirelessLinkRequest) Validate() error {
	return validation.Errors{
		"ssid": validation.Validate(
			m.Ssid, validation.Length(0, 32),
		),
	}.Filter()
}

// GetSsid returns the Ssid property
func (m NestedWirelessLinkRequest) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *NestedWirelessLinkRequest) SetSsid(val string) {
	m.Ssid = val
}
