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

// NestedCable is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedCable struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label:
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedCable) Validate() error {
	return validation.Errors{
		"label": validation.Validate(
			m.Label, validation.Length(0, 100),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedCable) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedCable) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedCable) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedCable) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m NestedCable) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *NestedCable) SetLabel(val string) {
	m.Label = val
}

// GetUrl returns the Url property
func (m NestedCable) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedCable) SetUrl(val string) {
	m.Url = val
}
