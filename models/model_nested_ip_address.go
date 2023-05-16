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

// NestedIPAddress is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedIPAddress struct {
	// Address:
	Address string `json:"address" mapstructure:"address"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Family:
	Family int32 `json:"family" mapstructure:"family"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedIPAddress) Validate() error {
	return validation.Errors{
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m NestedIPAddress) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *NestedIPAddress) SetAddress(val string) {
	m.Address = val
}

// GetDisplay returns the Display property
func (m NestedIPAddress) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedIPAddress) SetDisplay(val string) {
	m.Display = val
}

// GetFamily returns the Family property
func (m NestedIPAddress) GetFamily() int32 {
	return m.Family
}

// SetFamily sets the Family property
func (m *NestedIPAddress) SetFamily(val int32) {
	m.Family = val
}

// GetId returns the Id property
func (m NestedIPAddress) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedIPAddress) SetId(val int32) {
	m.Id = val
}

// GetUrl returns the Url property
func (m NestedIPAddress) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedIPAddress) SetUrl(val string) {
	m.Url = val
}
