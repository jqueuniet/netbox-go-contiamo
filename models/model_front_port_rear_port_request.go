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

// FrontPortRearPortRequest is an object. NestedRearPortSerializer but with parent device omitted (since front and rear ports must belong to same device)
type FrontPortRearPortRequest struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
}

// Validate implements basic validation for this model
func (m FrontPortRearPortRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m FrontPortRearPortRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *FrontPortRearPortRequest) SetDescription(val string) {
	m.Description = val
}

// GetLabel returns the Label property
func (m FrontPortRearPortRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *FrontPortRearPortRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m FrontPortRearPortRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *FrontPortRearPortRequest) SetName(val string) {
	m.Name = val
}
