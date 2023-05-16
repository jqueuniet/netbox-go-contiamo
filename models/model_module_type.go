// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ModuleType is an object. Adds support for custom fields and tags.
type ModuleType struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer NestedManufacturer `json:"manufacturer" mapstructure:"manufacturer"`
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// PartNumber: Discrete part number (optional)
	PartNumber string `json:"part_number,omitempty" mapstructure:"part_number,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Weight:
	Weight *float64 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightUnit:
	WeightUnit *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
}

// Validate implements basic validation for this model
func (m ModuleType) Validate() error {
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
			m.Model, validation.NotNil, validation.Length(0, 100),
		),
		"partNumber": validation.Validate(
			m.PartNumber, validation.Length(0, 50),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(*float64(-1e+06)), validation.Max(*float64(1e+06)),
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m ModuleType) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ModuleType) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m ModuleType) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ModuleType) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ModuleType) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ModuleType) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ModuleType) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ModuleType) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ModuleType) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ModuleType) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ModuleType) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ModuleType) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ModuleType) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ModuleType) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetManufacturer returns the Manufacturer property
func (m ModuleType) GetManufacturer() NestedManufacturer {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *ModuleType) SetManufacturer(val NestedManufacturer) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m ModuleType) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *ModuleType) SetModel(val string) {
	m.Model = val
}

// GetPartNumber returns the PartNumber property
func (m ModuleType) GetPartNumber() string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *ModuleType) SetPartNumber(val string) {
	m.PartNumber = val
}

// GetTags returns the Tags property
func (m ModuleType) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ModuleType) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m ModuleType) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ModuleType) SetUrl(val string) {
	m.Url = val
}

// GetWeight returns the Weight property
func (m ModuleType) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *ModuleType) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m ModuleType) GetWeightUnit() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *ModuleType) SetWeightUnit(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.WeightUnit = val
}
