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

// NestedVRF is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedVRF struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Rd: Unique route distinguisher (as defined in RFC 4364)
	Rd *string `json:"rd,omitempty" mapstructure:"rd,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedVRF) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"rd": validation.Validate(
			m.Rd, validation.Length(0, 21),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedVRF) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedVRF) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedVRF) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedVRF) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m NestedVRF) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedVRF) SetName(val string) {
	m.Name = val
}

// GetRd returns the Rd property
func (m NestedVRF) GetRd() *string {
	return m.Rd
}

// SetRd sets the Rd property
func (m *NestedVRF) SetRd(val *string) {
	m.Rd = val
}

// GetUrl returns the Url property
func (m NestedVRF) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedVRF) SetUrl(val string) {
	m.Url = val
}
