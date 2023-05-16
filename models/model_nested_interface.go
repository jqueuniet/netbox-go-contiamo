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

// NestedInterface is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedInterface struct {
	// Occupied:
	Occupied bool `json:"_occupied" mapstructure:"_occupied"`
	// Cable:
	Cable *int32 `json:"cable,omitempty" mapstructure:"cable,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDevice `json:"device" mapstructure:"device"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedInterface) Validate() error {
	return validation.Errors{
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetOccupied returns the Occupied property
func (m NestedInterface) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *NestedInterface) SetOccupied(val bool) {
	m.Occupied = val
}

// GetCable returns the Cable property
func (m NestedInterface) GetCable() *int32 {
	return m.Cable
}

// SetCable sets the Cable property
func (m *NestedInterface) SetCable(val *int32) {
	m.Cable = val
}

// GetDevice returns the Device property
func (m NestedInterface) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *NestedInterface) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m NestedInterface) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedInterface) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedInterface) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedInterface) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m NestedInterface) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedInterface) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m NestedInterface) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedInterface) SetUrl(val string) {
	m.Url = val
}
