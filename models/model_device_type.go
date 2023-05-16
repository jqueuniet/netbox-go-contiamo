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

	"regexp"
)

// deviceTypeSlugPattern is the validation pattern for DeviceType.Slug
var deviceTypeSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// DeviceType is an object. Adds support for custom fields and tags.
type DeviceType struct {
	// Airflow:
	Airflow *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"airflow,omitempty" mapstructure:"airflow,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DefaultPlatform: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DefaultPlatform *NestedPlatform `json:"default_platform,omitempty" mapstructure:"default_platform,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceCount:
	DeviceCount int32 `json:"device_count" mapstructure:"device_count"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// FrontImage:
	FrontImage string `json:"front_image,omitempty" mapstructure:"front_image,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IsFullDepth: Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty" mapstructure:"is_full_depth,omitempty"`
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
	// RearImage:
	RearImage string `json:"rear_image,omitempty" mapstructure:"rear_image,omitempty"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// SubdeviceRole:
	SubdeviceRole *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"subdevice_role,omitempty" mapstructure:"subdevice_role,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// UHeight:
	UHeight float64 `json:"u_height,omitempty" mapstructure:"u_height,omitempty"`
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
func (m DeviceType) Validate() error {
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
		"frontImage": validation.Validate(
			m.FrontImage, is.RequestURI,
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
		"rearImage": validation.Validate(
			m.RearImage, is.RequestURI,
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(deviceTypeSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"uHeight": validation.Validate(
			m.UHeight, validation.Min(float64(-5e-324)), validation.Max(float64(1000)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(*float64(-1e+06)), validation.Max(*float64(1e+06)),
		),
	}.Filter()
}

// GetAirflow returns the Airflow property
func (m DeviceType) GetAirflow() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *DeviceType) SetAirflow(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Airflow = val
}

// GetComments returns the Comments property
func (m DeviceType) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *DeviceType) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m DeviceType) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DeviceType) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m DeviceType) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *DeviceType) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDefaultPlatform returns the DefaultPlatform property
func (m DeviceType) GetDefaultPlatform() *NestedPlatform {
	return m.DefaultPlatform
}

// SetDefaultPlatform sets the DefaultPlatform property
func (m *DeviceType) SetDefaultPlatform(val *NestedPlatform) {
	m.DefaultPlatform = val
}

// GetDescription returns the Description property
func (m DeviceType) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceType) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m DeviceType) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *DeviceType) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m DeviceType) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *DeviceType) SetDisplay(val string) {
	m.Display = val
}

// GetFrontImage returns the FrontImage property
func (m DeviceType) GetFrontImage() string {
	return m.FrontImage
}

// SetFrontImage sets the FrontImage property
func (m *DeviceType) SetFrontImage(val string) {
	m.FrontImage = val
}

// GetId returns the Id property
func (m DeviceType) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DeviceType) SetId(val int32) {
	m.Id = val
}

// GetIsFullDepth returns the IsFullDepth property
func (m DeviceType) GetIsFullDepth() bool {
	return m.IsFullDepth
}

// SetIsFullDepth sets the IsFullDepth property
func (m *DeviceType) SetIsFullDepth(val bool) {
	m.IsFullDepth = val
}

// GetLastUpdated returns the LastUpdated property
func (m DeviceType) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DeviceType) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetManufacturer returns the Manufacturer property
func (m DeviceType) GetManufacturer() NestedManufacturer {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DeviceType) SetManufacturer(val NestedManufacturer) {
	m.Manufacturer = val
}

// GetModel returns the Model property
func (m DeviceType) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *DeviceType) SetModel(val string) {
	m.Model = val
}

// GetPartNumber returns the PartNumber property
func (m DeviceType) GetPartNumber() string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *DeviceType) SetPartNumber(val string) {
	m.PartNumber = val
}

// GetRearImage returns the RearImage property
func (m DeviceType) GetRearImage() string {
	return m.RearImage
}

// SetRearImage sets the RearImage property
func (m *DeviceType) SetRearImage(val string) {
	m.RearImage = val
}

// GetSlug returns the Slug property
func (m DeviceType) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *DeviceType) SetSlug(val string) {
	m.Slug = val
}

// GetSubdeviceRole returns the SubdeviceRole property
func (m DeviceType) GetSubdeviceRole() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.SubdeviceRole
}

// SetSubdeviceRole sets the SubdeviceRole property
func (m *DeviceType) SetSubdeviceRole(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.SubdeviceRole = val
}

// GetTags returns the Tags property
func (m DeviceType) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *DeviceType) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUHeight returns the UHeight property
func (m DeviceType) GetUHeight() float64 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *DeviceType) SetUHeight(val float64) {
	m.UHeight = val
}

// GetUrl returns the Url property
func (m DeviceType) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *DeviceType) SetUrl(val string) {
	m.Url = val
}

// GetWeight returns the Weight property
func (m DeviceType) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *DeviceType) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m DeviceType) GetWeightUnit() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *DeviceType) SetWeightUnit(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.WeightUnit = val
}
