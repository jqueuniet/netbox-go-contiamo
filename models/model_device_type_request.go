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

// deviceTypeRequestSlugPattern is the validation pattern for DeviceTypeRequest.Slug
var deviceTypeRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// DeviceTypeRequest is an object. Adds support for custom fields and tags.
type DeviceTypeRequest struct {
	// Airflow: * `front-to-rear` - Front to rear
	// * `rear-to-front` - Rear to front
	// * `left-to-right` - Left to right
	// * `right-to-left` - Right to left
	// * `side-to-rear` - Side to rear
	// * `passive` - Passive
	// * `mixed` - Mixed
	Airflow *string `json:"airflow,omitempty" mapstructure:"airflow,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DefaultPlatform: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DefaultPlatform *NestedPlatformRequest `json:"default_platform,omitempty" mapstructure:"default_platform,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// FrontImage:
	FrontImage string `json:"front_image,omitempty" mapstructure:"front_image,omitempty"`
	// IsFullDepth: Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty" mapstructure:"is_full_depth,omitempty"`
	// Manufacturer: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Manufacturer NestedManufacturerRequest `json:"manufacturer" mapstructure:"manufacturer"`
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// PartNumber: Discrete part number (optional)
	PartNumber string `json:"part_number,omitempty" mapstructure:"part_number,omitempty"`
	// RearImage:
	RearImage string `json:"rear_image,omitempty" mapstructure:"rear_image,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// SubdeviceRole: * `parent` - Parent
	// * `child` - Child
	SubdeviceRole *string `json:"subdevice_role,omitempty" mapstructure:"subdevice_role,omitempty"`
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
	WeightUnit *string `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
}

// Validate implements basic validation for this model
func (m DeviceTypeRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"defaultPlatform": validation.Validate(
			m.DefaultPlatform,
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
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(deviceTypeRequestSlugPattern),
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
func (m DeviceTypeRequest) GetAirflow() *string {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *DeviceTypeRequest) SetAirflow(val *string) {
	m.Airflow = val
}

// GetComments returns the Comments property
func (m DeviceTypeRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *DeviceTypeRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m DeviceTypeRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *DeviceTypeRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDefaultPlatform returns the DefaultPlatform property
func (m DeviceTypeRequest) GetDefaultPlatform() *NestedPlatformRequest {
	return m.DefaultPlatform
}

// SetDefaultPlatform sets the DefaultPlatform property
func (m *DeviceTypeRequest) SetDefaultPlatform(val *NestedPlatformRequest) {
	m.DefaultPlatform = val
}

// GetDescription returns the Description property
func (m DeviceTypeRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceTypeRequest) SetDescription(val string) {
	m.Description = val
}

// GetFrontImage returns the FrontImage property
func (m DeviceTypeRequest) GetFrontImage() string {
	return m.FrontImage
}

// SetFrontImage sets the FrontImage property
func (m *DeviceTypeRequest) SetFrontImage(val string) {
	m.FrontImage = val
}

// GetIsFullDepth returns the IsFullDepth property
func (m DeviceTypeRequest) GetIsFullDepth() bool {
	return m.IsFullDepth
}

// SetIsFullDepth sets the IsFullDepth property
func (m *DeviceTypeRequest) SetIsFullDepth(val bool) {
	m.IsFullDepth = val
}

// GetManufacturer returns the Manufacturer property
func (m DeviceTypeRequest) GetManufacturer() NestedManufacturerRequest {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DeviceTypeRequest) SetManufacturer(val NestedManufacturerRequest) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m DeviceTypeRequest) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *DeviceTypeRequest) SetModel(val string) {
	m.Model = val
}

// GetPartNumber returns the PartNumber property
func (m DeviceTypeRequest) GetPartNumber() string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *DeviceTypeRequest) SetPartNumber(val string) {
	m.PartNumber = val
}

// GetRearImage returns the RearImage property
func (m DeviceTypeRequest) GetRearImage() string {
	return m.RearImage
}

// SetRearImage sets the RearImage property
func (m *DeviceTypeRequest) SetRearImage(val string) {
	m.RearImage = val
}

// GetSlug returns the Slug property
func (m DeviceTypeRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *DeviceTypeRequest) SetSlug(val string) {
	m.Slug = val
}

// GetSubdeviceRole returns the SubdeviceRole property
func (m DeviceTypeRequest) GetSubdeviceRole() *string {
	return m.SubdeviceRole
}

// SetSubdeviceRole sets the SubdeviceRole property
func (m *DeviceTypeRequest) SetSubdeviceRole(val *string) {
	m.SubdeviceRole = val
}

// GetTags returns the Tags property
func (m DeviceTypeRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *DeviceTypeRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetUHeight returns the UHeight property
func (m DeviceTypeRequest) GetUHeight() float64 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *DeviceTypeRequest) SetUHeight(val float64) {
	m.UHeight = val
}

// GetWeight returns the Weight property
func (m DeviceTypeRequest) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *DeviceTypeRequest) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m DeviceTypeRequest) GetWeightUnit() *string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *DeviceTypeRequest) SetWeightUnit(val *string) {
	m.WeightUnit = val
}
