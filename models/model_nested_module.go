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

// NestedModule is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedModule struct {
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDevice `json:"device" mapstructure:"device"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// ModuleBay: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleBay ModuleNestedModuleBay `json:"module_bay" mapstructure:"module_bay"`
	// ModuleType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleType NestedModuleType `json:"module_type" mapstructure:"module_type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedModule) Validate() error {
	return validation.Errors{
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"moduleBay": validation.Validate(
			m.ModuleBay, validation.NotNil,
		),
		"moduleType": validation.Validate(
			m.ModuleType, validation.NotNil,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDevice returns the Device property
func (m NestedModule) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *NestedModule) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m NestedModule) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedModule) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedModule) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedModule) SetId(val int32) {
	m.Id = val
}

// GetModuleBay returns the ModuleBay property
func (m NestedModule) GetModuleBay() ModuleNestedModuleBay {
	return m.ModuleBay
}

// SetModuleBay sets the ModuleBay property
func (m *NestedModule) SetModuleBay(val ModuleNestedModuleBay) {
	m.ModuleBay = val
}

// GetModuleType returns the ModuleType property
func (m NestedModule) GetModuleType() NestedModuleType {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *NestedModule) SetModuleType(val NestedModuleType) {
	m.ModuleType = val
}

// GetUrl returns the Url property
func (m NestedModule) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedModule) SetUrl(val string) {
	m.Url = val
}
