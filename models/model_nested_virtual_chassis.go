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

// NestedVirtualChassis is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedVirtualChassis struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Master: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Master NestedDevice `json:"master" mapstructure:"master"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedVirtualChassis) Validate() error {
	return validation.Errors{
		"master": validation.Validate(
			m.Master, validation.NotNil,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedVirtualChassis) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedVirtualChassis) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedVirtualChassis) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedVirtualChassis) SetId(val int32) {
	m.Id = val
}

// GetMaster returns the Master property
func (m NestedVirtualChassis) GetMaster() NestedDevice {
	return m.Master
}

// SetMaster sets the Master property
func (m *NestedVirtualChassis) SetMaster(val NestedDevice) {
	m.Master = val
}

// GetName returns the Name property
func (m NestedVirtualChassis) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedVirtualChassis) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m NestedVirtualChassis) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedVirtualChassis) SetUrl(val string) {
	m.Url = val
}
