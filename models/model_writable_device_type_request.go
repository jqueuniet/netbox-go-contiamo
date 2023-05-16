// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// writableDeviceTypeRequestSlugPattern is the validation pattern for WritableDeviceTypeRequest.Slug
var writableDeviceTypeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// WritableDeviceTypeRequest is an object. Adds support for custom fields and tags.
type WritableDeviceTypeRequest struct {
	// Airflow: * `front-to-rear` - Front to rear
	// * `rear-to-front` - Rear to front
	// * `left-to-right` - Left to right
	// * `right-to-left` - Right to left
	// * `side-to-rear` - Side to rear
	// * `passive` - Passive
	// * `mixed` - Mixed
	Airflow string `json:"airflow,omitempty" mapstructure:"airflow,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DefaultPlatform:
	DefaultPlatform *int32 `json:"default_platform,omitempty" mapstructure:"default_platform,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// FrontImage:
	FrontImage string `json:"front_image,omitempty" mapstructure:"front_image,omitempty"`
	// IsFullDepth: Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty" mapstructure:"is_full_depth,omitempty"`
	// Manufacturer:
	Manufacturer int32 `json:"manufacturer" mapstructure:"manufacturer"`
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// PartNumber: Discrete part number (optional)
	PartNumber string `json:"part_number,omitempty" mapstructure:"part_number,omitempty"`
	// RearImage:
	RearImage string `json:"rear_image,omitempty" mapstructure:"rear_image,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// SubdeviceRole: Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.
	//
	// * `parent` - Parent
	// * `child` - Child
	SubdeviceRole string `json:"subdevice_role,omitempty" mapstructure:"subdevice_role,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// UHeight:
	UHeight float64 `json:"u_height,omitempty" mapstructure:"u_height,omitempty"`
	// Weight:
	Weight *float64 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightUnit: * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	WeightUnit string `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableDeviceTypeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"model": validation.Validate(
			m.Model, validation.Required, validation.Length(1, 100),
		),
		"partNumber": validation.Validate(
			m.PartNumber, validation.Length(0, 50),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(writableDeviceTypeRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"uHeight": validation.Validate(
			m.UHeight, validation.Min(float64(-5e-324)), validation.Max(float64(1000)),
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(*float64(-1e+06)), validation.Max(*float64(1e+06)),
		),
	}.Filter()
}

// GetAirflow returns the Airflow property
func (m WritableDeviceTypeRequest) GetAirflow() string {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *WritableDeviceTypeRequest) SetAirflow(val string) {
	m.Airflow = val
}

// GetComments returns the Comments property
func (m WritableDeviceTypeRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableDeviceTypeRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableDeviceTypeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableDeviceTypeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDefaultPlatform returns the DefaultPlatform property
func (m WritableDeviceTypeRequest) GetDefaultPlatform() *int32 {
	return m.DefaultPlatform
}

// SetDefaultPlatform sets the DefaultPlatform property
func (m *WritableDeviceTypeRequest) SetDefaultPlatform(val *int32) {
	m.DefaultPlatform = val
}

// GetDescription returns the Description property
func (m WritableDeviceTypeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableDeviceTypeRequest) SetDescription(val string) {
	m.Description = val
}

// GetFrontImage returns the FrontImage property
func (m WritableDeviceTypeRequest) GetFrontImage() string {
	return m.FrontImage
}

// SetFrontImage sets the FrontImage property
func (m *WritableDeviceTypeRequest) SetFrontImage(val string) {
	m.FrontImage = val
}

// GetIsFullDepth returns the IsFullDepth property
func (m WritableDeviceTypeRequest) GetIsFullDepth() bool {
	return m.IsFullDepth
}

// SetIsFullDepth sets the IsFullDepth property
func (m *WritableDeviceTypeRequest) SetIsFullDepth(val bool) {
	m.IsFullDepth = val
}

// GetManufacturer returns the Manufacturer property
func (m WritableDeviceTypeRequest) GetManufacturer() int32 {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *WritableDeviceTypeRequest) SetManufacturer(val int32) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m WritableDeviceTypeRequest) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *WritableDeviceTypeRequest) SetModel(val string) {
	m.Model = val
}

// GetPartNumber returns the PartNumber property
func (m WritableDeviceTypeRequest) GetPartNumber() string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *WritableDeviceTypeRequest) SetPartNumber(val string) {
	m.PartNumber = val
}

// GetRearImage returns the RearImage property
func (m WritableDeviceTypeRequest) GetRearImage() string {
	return m.RearImage
}

// SetRearImage sets the RearImage property
func (m *WritableDeviceTypeRequest) SetRearImage(val string) {
	m.RearImage = val
}

// GetSlug returns the Slug property
func (m WritableDeviceTypeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WritableDeviceTypeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetSubdeviceRole returns the SubdeviceRole property
func (m WritableDeviceTypeRequest) GetSubdeviceRole() string {
	return m.SubdeviceRole
}

// SetSubdeviceRole sets the SubdeviceRole property
func (m *WritableDeviceTypeRequest) SetSubdeviceRole(val string) {
	m.SubdeviceRole = val
}

// GetTags returns the Tags property
func (m WritableDeviceTypeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableDeviceTypeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetUHeight returns the UHeight property
func (m WritableDeviceTypeRequest) GetUHeight() float64 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *WritableDeviceTypeRequest) SetUHeight(val float64) {
	m.UHeight = val
}

// GetWeight returns the Weight property
func (m WritableDeviceTypeRequest) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *WritableDeviceTypeRequest) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m WritableDeviceTypeRequest) GetWeightUnit() string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *WritableDeviceTypeRequest) SetWeightUnit(val string) {
	m.WeightUnit = val
}
