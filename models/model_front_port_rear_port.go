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

// FrontPortRearPort is an object. NestedRearPortSerializer but with parent device omitted (since front and rear ports must belong to same device)
type FrontPortRearPort struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m FrontPortRearPort) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m FrontPortRearPort) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *FrontPortRearPort) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m FrontPortRearPort) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *FrontPortRearPort) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m FrontPortRearPort) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *FrontPortRearPort) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m FrontPortRearPort) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *FrontPortRearPort) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m FrontPortRearPort) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *FrontPortRearPort) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m FrontPortRearPort) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *FrontPortRearPort) SetUrl(val string) {
	m.Url = val
}
