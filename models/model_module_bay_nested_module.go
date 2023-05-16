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

// ModuleBayNestedModule is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type ModuleBayNestedModule struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ModuleBayNestedModule) Validate() error {
	return validation.Errors{
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m ModuleBayNestedModule) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ModuleBayNestedModule) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ModuleBayNestedModule) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ModuleBayNestedModule) SetId(val int32) {
	m.Id = val
}

// GetSerial returns the Serial property
func (m ModuleBayNestedModule) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *ModuleBayNestedModule) SetSerial(val string) {
	m.Serial = val
}

// GetUrl returns the Url property
func (m ModuleBayNestedModule) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ModuleBayNestedModule) SetUrl(val string) {
	m.Url = val
}
