// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// NestedWirelessLink is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedWirelessLink struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Ssid:
	Ssid string `json:"ssid,omitempty" mapstructure:"ssid,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedWirelessLink) Validate() error {
	return validation.Errors{
		"ssid": validation.Validate(
			m.Ssid, validation.Length(0, 32),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedWirelessLink) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedWirelessLink) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedWirelessLink) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedWirelessLink) SetId(val int32) {
	m.Id = val
}

// GetSsid returns the Ssid property
func (m NestedWirelessLink) GetSsid() string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *NestedWirelessLink) SetSsid(val string) {
	m.Ssid = val
}

// GetUrl returns the Url property
func (m NestedWirelessLink) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedWirelessLink) SetUrl(val string) {
	m.Url = val
}
