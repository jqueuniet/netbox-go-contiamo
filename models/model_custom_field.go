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

	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// customFieldNamePattern is the validation pattern for CustomField.Name
var customFieldNamePattern = regexp.MustCompile(`^[a-z0-9_]+$`)

// CustomField is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomField struct {
	// Choices: Comma-separated list of available choices (for selection fields)
	Choices []string `json:"choices,omitempty" mapstructure:"choices,omitempty"`
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// DataType:
	DataType string `json:"data_type" mapstructure:"data_type"`
	// Default: Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. "Foo").
	Default map[string]interface{} `json:"default,omitempty" mapstructure:"default,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// FilterLogic:
	FilterLogic *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"filter_logic,omitempty" mapstructure:"filter_logic,omitempty"`
	// GroupName: Custom fields within the same group will be displayed together
	GroupName string `json:"group_name,omitempty" mapstructure:"group_name,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IsCloneable: Replicate this value when cloning objects
	IsCloneable bool `json:"is_cloneable,omitempty" mapstructure:"is_cloneable,omitempty"`
	// Label: Name of the field as displayed to users (if not provided, the field's name will be used)
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name: Internal field name
	Name string `json:"name" mapstructure:"name"`
	// ObjectType:
	ObjectType string `json:"object_type,omitempty" mapstructure:"object_type,omitempty"`
	// Required: If true, this field is required when creating new objects or editing an existing object.
	Required bool `json:"required,omitempty" mapstructure:"required,omitempty"`
	// SearchWeight: Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored.
	SearchWeight int32 `json:"search_weight,omitempty" mapstructure:"search_weight,omitempty"`
	// Type:
	Type struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type" mapstructure:"type"`
	// UiVisibility:
	UiVisibility *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"ui_visibility,omitempty" mapstructure:"ui_visibility,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// ValidationMaximum: Maximum allowed value (for numeric fields)
	ValidationMaximum *int32 `json:"validation_maximum,omitempty" mapstructure:"validation_maximum,omitempty"`
	// ValidationMinimum: Minimum allowed value (for numeric fields)
	ValidationMinimum *int32 `json:"validation_minimum,omitempty" mapstructure:"validation_minimum,omitempty"`
	// ValidationRegex: Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters.
	ValidationRegex string `json:"validation_regex,omitempty" mapstructure:"validation_regex,omitempty"`
	// Weight: Fields with higher weights appear lower in a form.
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m CustomField) Validate() error {
	return validation.Errors{
		"choices": validation.Validate(
			m.Choices,
		),
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
		),
		"default": validation.Validate(
			m.Default,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"groupName": validation.Validate(
			m.GroupName, validation.Length(0, 50),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 50),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 50), validation.Match(customFieldNamePattern),
		),
		"searchWeight": validation.Validate(
			m.SearchWeight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"validationMaximum": validation.Validate(
			m.ValidationMaximum, validation.Min(*int32(-2.147483648e+09)), validation.Max(*int32(2.147483647e+09)),
		),
		"validationMinimum": validation.Validate(
			m.ValidationMinimum, validation.Min(*int32(-2.147483648e+09)), validation.Max(*int32(2.147483647e+09)),
		),
		"validationRegex": validation.Validate(
			m.ValidationRegex, validation.Length(0, 500),
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetChoices returns the Choices property
func (m CustomField) GetChoices() []string {
	return m.Choices
}

// SetChoices sets the Choices property
func (m *CustomField) SetChoices(val []string) {
	m.Choices = val
}

// GetContentTypes returns the ContentTypes property
func (m CustomField) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *CustomField) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetCreated returns the Created property
func (m CustomField) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CustomField) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDataType returns the DataType property
func (m CustomField) GetDataType() string {
	return m.DataType
}

// SetDataType sets the DataType property
func (m *CustomField) SetDataType(val string) {
	m.DataType = val
}

// GetDefault returns the Default property
func (m CustomField) GetDefault() map[string]interface{} {
	return m.Default
}

// SetDefault sets the Default property
func (m *CustomField) SetDefault(val map[string]interface{}) {
	m.Default = val
}

// GetDescription returns the Description property
func (m CustomField) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CustomField) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m CustomField) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *CustomField) SetDisplay(val string) {
	m.Display = val
}

// GetFilterLogic returns the FilterLogic property
func (m CustomField) GetFilterLogic() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.FilterLogic
}

// SetFilterLogic sets the FilterLogic property
func (m *CustomField) SetFilterLogic(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.FilterLogic = val
}

// GetGroupName returns the GroupName property
func (m CustomField) GetGroupName() string {
	return m.GroupName
}

// SetGroupName sets the GroupName property
func (m *CustomField) SetGroupName(val string) {
	m.GroupName = val
}

// GetId returns the Id property
func (m CustomField) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CustomField) SetId(val int32) {
	m.Id = val
}

// GetIsCloneable returns the IsCloneable property
func (m CustomField) GetIsCloneable() bool {
	return m.IsCloneable
}

// SetIsCloneable sets the IsCloneable property
func (m *CustomField) SetIsCloneable(val bool) {
	m.IsCloneable = val
}

// GetLabel returns the Label property
func (m CustomField) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *CustomField) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m CustomField) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CustomField) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m CustomField) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *CustomField) SetName(val string) {
	m.Name = val
}

// GetObjectType returns the ObjectType property
func (m CustomField) GetObjectType() string {
	return m.ObjectType
}

// SetObjectType sets the ObjectType property
func (m *CustomField) SetObjectType(val string) {
	m.ObjectType = val
}

// GetRequired returns the Required property
func (m CustomField) GetRequired() bool {
	return m.Required
}

// SetRequired sets the Required property
func (m *CustomField) SetRequired(val bool) {
	m.Required = val
}

// GetSearchWeight returns the SearchWeight property
func (m CustomField) GetSearchWeight() int32 {
	return m.SearchWeight
}

// SetSearchWeight sets the SearchWeight property
func (m *CustomField) SetSearchWeight(val int32) {
	m.SearchWeight = val
}

// GetType returns the Type property
func (m CustomField) GetType() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *CustomField) SetType(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUiVisibility returns the UiVisibility property
func (m CustomField) GetUiVisibility() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.UiVisibility
}

// SetUiVisibility sets the UiVisibility property
func (m *CustomField) SetUiVisibility(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.UiVisibility = val
}

// GetUrl returns the Url property
func (m CustomField) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *CustomField) SetUrl(val string) {
	m.Url = val
}

// GetValidationMaximum returns the ValidationMaximum property
func (m CustomField) GetValidationMaximum() *int32 {
	return m.ValidationMaximum
}

// SetValidationMaximum sets the ValidationMaximum property
func (m *CustomField) SetValidationMaximum(val *int32) {
	m.ValidationMaximum = val
}

// GetValidationMinimum returns the ValidationMinimum property
func (m CustomField) GetValidationMinimum() *int32 {
	return m.ValidationMinimum
}

// SetValidationMinimum sets the ValidationMinimum property
func (m *CustomField) SetValidationMinimum(val *int32) {
	m.ValidationMinimum = val
}

// GetValidationRegex returns the ValidationRegex property
func (m CustomField) GetValidationRegex() string {
	return m.ValidationRegex
}

// SetValidationRegex sets the ValidationRegex property
func (m *CustomField) SetValidationRegex(val string) {
	m.ValidationRegex = val
}

// GetWeight returns the Weight property
func (m CustomField) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *CustomField) SetWeight(val int32) {
	m.Weight = val
}
