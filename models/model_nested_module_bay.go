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

// NestedModuleBay is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedModuleBay struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Module: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Module *NestedModule `json:"module" mapstructure:"module"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedModuleBay) Validate() error {
	return validation.Errors{
		"module": validation.Validate(
			m.Module,
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
func (m NestedModuleBay) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedModuleBay) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedModuleBay) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedModuleBay) SetId(val int32) {
	m.Id = val
}

// GetModule returns the Module property
func (m NestedModuleBay) GetModule() *NestedModule {
	return m.Module
}

// SetModule sets the Module property
func (m *NestedModuleBay) SetModule(val *NestedModule) {
	m.Module = val
}

// GetName returns the Name property
func (m NestedModuleBay) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedModuleBay) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m NestedModuleBay) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedModuleBay) SetUrl(val string) {
	m.Url = val
}
