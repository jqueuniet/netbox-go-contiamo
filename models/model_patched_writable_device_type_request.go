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

// patchedWritableDeviceTypeRequestSlugPattern is the validation pattern for PatchedWritableDeviceTypeRequest.Slug
var patchedWritableDeviceTypeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedWritableDeviceTypeRequest is an object. Adds support for custom fields and tags.
type PatchedWritableDeviceTypeRequest struct {
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
	Manufacturer int32 `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// Model:
	Model string `json:"model,omitempty" mapstructure:"model,omitempty"`
	// PartNumber: Discrete part number (optional)
	PartNumber string `json:"part_number,omitempty" mapstructure:"part_number,omitempty"`
	// RearImage:
	RearImage string `json:"rear_image,omitempty" mapstructure:"rear_image,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
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
func (m PatchedWritableDeviceTypeRequest) Validate() error {
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
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedWritableDeviceTypeRequestSlugPattern),
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
func (m PatchedWritableDeviceTypeRequest) GetAirflow() string {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *PatchedWritableDeviceTypeRequest) SetAirflow(val string) {
	m.Airflow = val
}

// GetComments returns the Comments property
func (m PatchedWritableDeviceTypeRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableDeviceTypeRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableDeviceTypeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableDeviceTypeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDefaultPlatform returns the DefaultPlatform property
func (m PatchedWritableDeviceTypeRequest) GetDefaultPlatform() *int32 {
	return m.DefaultPlatform
}

// SetDefaultPlatform sets the DefaultPlatform property
func (m *PatchedWritableDeviceTypeRequest) SetDefaultPlatform(val *int32) {
	m.DefaultPlatform = val
}

// GetDescription returns the Description property
func (m PatchedWritableDeviceTypeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableDeviceTypeRequest) SetDescription(val string) {
	m.Description = val
}

// GetFrontImage returns the FrontImage property
func (m PatchedWritableDeviceTypeRequest) GetFrontImage() string {
	return m.FrontImage
}

// SetFrontImage sets the FrontImage property
func (m *PatchedWritableDeviceTypeRequest) SetFrontImage(val string) {
	m.FrontImage = val
}

// GetIsFullDepth returns the IsFullDepth property
func (m PatchedWritableDeviceTypeRequest) GetIsFullDepth() bool {
	return m.IsFullDepth
}

// SetIsFullDepth sets the IsFullDepth property
func (m *PatchedWritableDeviceTypeRequest) SetIsFullDepth(val bool) {
	m.IsFullDepth = val
}

// GetManufacturer returns the Manufacturer property
func (m PatchedWritableDeviceTypeRequest) GetManufacturer() int32 {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *PatchedWritableDeviceTypeRequest) SetManufacturer(val int32) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m PatchedWritableDeviceTypeRequest) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *PatchedWritableDeviceTypeRequest) SetModel(val string) {
	m.Model = val
}

// GetPartNumber returns the PartNumber property
func (m PatchedWritableDeviceTypeRequest) GetPartNumber() string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *PatchedWritableDeviceTypeRequest) SetPartNumber(val string) {
	m.PartNumber = val
}

// GetRearImage returns the RearImage property
func (m PatchedWritableDeviceTypeRequest) GetRearImage() string {
	return m.RearImage
}

// SetRearImage sets the RearImage property
func (m *PatchedWritableDeviceTypeRequest) SetRearImage(val string) {
	m.RearImage = val
}

// GetSlug returns the Slug property
func (m PatchedWritableDeviceTypeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedWritableDeviceTypeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetSubdeviceRole returns the SubdeviceRole property
func (m PatchedWritableDeviceTypeRequest) GetSubdeviceRole() string {
	return m.SubdeviceRole
}

// SetSubdeviceRole sets the SubdeviceRole property
func (m *PatchedWritableDeviceTypeRequest) SetSubdeviceRole(val string) {
	m.SubdeviceRole = val
}

// GetTags returns the Tags property
func (m PatchedWritableDeviceTypeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableDeviceTypeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetUHeight returns the UHeight property
func (m PatchedWritableDeviceTypeRequest) GetUHeight() float64 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *PatchedWritableDeviceTypeRequest) SetUHeight(val float64) {
	m.UHeight = val
}

// GetWeight returns the Weight property
func (m PatchedWritableDeviceTypeRequest) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *PatchedWritableDeviceTypeRequest) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m PatchedWritableDeviceTypeRequest) GetWeightUnit() string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *PatchedWritableDeviceTypeRequest) SetWeightUnit(val string) {
	m.WeightUnit = val
}
