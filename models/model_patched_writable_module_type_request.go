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

// PatchedWritableModuleTypeRequest is an object. Adds support for custom fields and tags.
type PatchedWritableModuleTypeRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Manufacturer:
	Manufacturer int32 `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Model:
	Model string `json:"model,omitempty" mapstructure:"model,omitempty"`
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
	WeightUnit string `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableModuleTypeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"model": validation.Validate(
			m.Model, validation.Length(1, 100),
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
func (m PatchedWritableModuleTypeRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableModuleTypeRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableModuleTypeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableModuleTypeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableModuleTypeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableModuleTypeRequest) SetDescription(val string) {
	m.Description = val
}

// GetManufacturer returns the Manufacturer property
func (m PatchedWritableModuleTypeRequest) GetManufacturer() int32 {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *PatchedWritableModuleTypeRequest) SetManufacturer(val int32) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m PatchedWritableModuleTypeRequest) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *PatchedWritableModuleTypeRequest) SetModel(val string) {
	m.Model = val
}

// GetPartNumber returns the PartNumber property
func (m PatchedWritableModuleTypeRequest) GetPartNumber() string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *PatchedWritableModuleTypeRequest) SetPartNumber(val string) {
	m.PartNumber = val
}

// GetTags returns the Tags property
func (m PatchedWritableModuleTypeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableModuleTypeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetWeight returns the Weight property
func (m PatchedWritableModuleTypeRequest) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *PatchedWritableModuleTypeRequest) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m PatchedWritableModuleTypeRequest) GetWeightUnit() string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *PatchedWritableModuleTypeRequest) SetWeightUnit(val string) {
	m.WeightUnit = val
}
