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

// NestedVirtualChassisRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedVirtualChassisRequest struct {
	// Master: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Master NestedDeviceRequest `json:"master" mapstructure:"master"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
}

// Validate implements basic validation for this model
func (m NestedVirtualChassisRequest) Validate() error {
	return validation.Errors{
		"master": validation.Validate(
			m.Master, validation.NotNil,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
	}.Filter()
}

// GetMaster returns the Master property
func (m NestedVirtualChassisRequest) GetMaster() NestedDeviceRequest {
	return m.Master
}

// SetMaster sets the Master property
func (m *NestedVirtualChassisRequest) SetMaster(val NestedDeviceRequest) {
	m.Master = val
}

// GetName returns the Name property
func (m NestedVirtualChassisRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedVirtualChassisRequest) SetName(val string) {
	m.Name = val
}
