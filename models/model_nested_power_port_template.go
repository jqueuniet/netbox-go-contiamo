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

// NestedPowerPortTemplate is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedPowerPortTemplate struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedPowerPortTemplate) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedPowerPortTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedPowerPortTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedPowerPortTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedPowerPortTemplate) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m NestedPowerPortTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedPowerPortTemplate) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m NestedPowerPortTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedPowerPortTemplate) SetUrl(val string) {
	m.Url = val
}
