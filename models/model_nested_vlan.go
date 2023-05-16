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

// NestedVLAN is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedVLAN struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Vid: Numeric VLAN ID (1-4094)
	Vid int32 `json:"vid" mapstructure:"vid"`
}

// Validate implements basic validation for this model
func (m NestedVLAN) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"vid": validation.Validate(
			m.Vid, validation.Required, validation.Min(int32(1)), validation.Max(int32(4094)),
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedVLAN) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedVLAN) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedVLAN) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedVLAN) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m NestedVLAN) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedVLAN) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m NestedVLAN) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedVLAN) SetUrl(val string) {
	m.Url = val
}

// GetVid returns the Vid property
func (m NestedVLAN) GetVid() int32 {
	return m.Vid
}

// SetVid sets the Vid property
func (m *NestedVLAN) SetVid(val int32) {
	m.Vid = val
}
