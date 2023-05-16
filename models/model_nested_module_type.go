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

// NestedModuleType is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedModuleType struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer NestedManufacturer `json:"manufacturer" mapstructure:"manufacturer"`
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedModuleType) Validate() error {
	return validation.Errors{
		"manufacturer": validation.Validate(
			m.Manufacturer, validation.NotNil,
		),
		"model": validation.Validate(
			m.Model, validation.NotNil, validation.Length(0, 100),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedModuleType) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedModuleType) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedModuleType) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedModuleType) SetId(val int32) {
	m.Id = val
}

// GetManufacturer returns the Manufacturer property
func (m NestedModuleType) GetManufacturer() NestedManufacturer {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *NestedModuleType) SetManufacturer(val NestedManufacturer) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m NestedModuleType) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *NestedModuleType) SetModel(val string) {
	m.Model = val
}

// GetUrl returns the Url property
func (m NestedModuleType) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedModuleType) SetUrl(val string) {
	m.Url = val
}
