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

// NestedL2VPNTermination is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedL2VPNTermination struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// L2vpn: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	L2vpn NestedL2VPN `json:"l2vpn" mapstructure:"l2vpn"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedL2VPNTermination) Validate() error {
	return validation.Errors{
		"l2vpn": validation.Validate(
			m.L2vpn, validation.NotNil,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedL2VPNTermination) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedL2VPNTermination) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedL2VPNTermination) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedL2VPNTermination) SetId(val int32) {
	m.Id = val
}

// GetL2vpn returns the L2vpn property
func (m NestedL2VPNTermination) GetL2vpn() NestedL2VPN {
	return m.L2vpn
}

// SetL2vpn sets the L2vpn property
func (m *NestedL2VPNTermination) SetL2vpn(val NestedL2VPN) {
	m.L2vpn = val
}

// GetUrl returns the Url property
func (m NestedL2VPNTermination) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedL2VPNTermination) SetUrl(val string) {
	m.Url = val
}
