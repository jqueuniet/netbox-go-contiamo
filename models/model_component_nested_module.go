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

// ComponentNestedModule is an object. Used by device component serializers.
type ComponentNestedModule struct {
	// Device:
	Device int32 `json:"device" mapstructure:"device"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// ModuleBay: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleBay ModuleNestedModuleBay `json:"module_bay" mapstructure:"module_bay"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ComponentNestedModule) Validate() error {
	return validation.Errors{
		"moduleBay": validation.Validate(
			m.ModuleBay, validation.NotNil,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDevice returns the Device property
func (m ComponentNestedModule) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *ComponentNestedModule) SetDevice(val int32) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m ComponentNestedModule) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ComponentNestedModule) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ComponentNestedModule) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ComponentNestedModule) SetId(val int32) {
	m.Id = val
}

// GetModuleBay returns the ModuleBay property
func (m ComponentNestedModule) GetModuleBay() ModuleNestedModuleBay {
	return m.ModuleBay
}

// SetModuleBay sets the ModuleBay property
func (m *ComponentNestedModule) SetModuleBay(val ModuleNestedModuleBay) {
	m.ModuleBay = val
}

// GetUrl returns the Url property
func (m ComponentNestedModule) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ComponentNestedModule) SetUrl(val string) {
	m.Url = val
}
