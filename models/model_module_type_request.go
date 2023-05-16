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

// ModuleTypeRequest is an object. Adds support for custom fields and tags.
type ModuleTypeRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer NestedManufacturerRequest `json:"manufacturer" mapstructure:"manufacturer"`
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// PartNumber: Discrete part number (optional)
	PartNumber string `json:"part_number,omitempty" mapstructure:"part_number,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Weight:
	Weight *float64 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightUnit: * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	WeightUnit *string `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
}

// Validate implements basic validation for this model
func (m ModuleTypeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"manufacturer": validation.Validate(
			m.Manufacturer, validation.NotNil,
		),
		"model": validation.Validate(
			m.Model, validation.Required, validation.Length(1, 100),
		),
		"partNumber": validation.Validate(
			m.PartNumber, validation.Length(0, 50),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(*float64(-1e+06)), validation.Max(*float64(1e+06)),
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m ModuleTypeRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ModuleTypeRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m ModuleTypeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ModuleTypeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ModuleTypeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ModuleTypeRequest) SetDescription(val string) {
	m.Description = val
}

// GetManufacturer returns the Manufacturer property
func (m ModuleTypeRequest) GetManufacturer() NestedManufacturerRequest {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *ModuleTypeRequest) SetManufacturer(val NestedManufacturerRequest) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m ModuleTypeRequest) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *ModuleTypeRequest) SetModel(val string) {
	m.Model = val
}

// GetPartNumber returns the PartNumber property
func (m ModuleTypeRequest) GetPartNumber() string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *ModuleTypeRequest) SetPartNumber(val string) {
	m.PartNumber = val
}

// GetTags returns the Tags property
func (m ModuleTypeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ModuleTypeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetWeight returns the Weight property
func (m ModuleTypeRequest) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *ModuleTypeRequest) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m ModuleTypeRequest) GetWeightUnit() *string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *ModuleTypeRequest) SetWeightUnit(val *string) {
	m.WeightUnit = val
}
