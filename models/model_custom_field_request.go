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

// customFieldRequestNamePattern is the validation pattern for CustomFieldRequest.Name
var customFieldRequestNamePattern = regexp.MustCompile(`^[a-z0-9_]+$`)

// CustomFieldRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomFieldRequest struct {
	// Choices: Comma-separated list of available choices (for selection fields)
	Choices []string `json:"choices,omitempty" mapstructure:"choices,omitempty"`
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// Default: Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. "Foo").
	Default map[string]interface{} `json:"default,omitempty" mapstructure:"default,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// FilterLogic: * `disabled` - Disabled
	// * `loose` - Loose
	// * `exact` - Exact
	FilterLogic string `json:"filter_logic,omitempty" mapstructure:"filter_logic,omitempty"`
	// GroupName: Custom fields within the same group will be displayed together
	GroupName string `json:"group_name,omitempty" mapstructure:"group_name,omitempty"`
	// IsCloneable: Replicate this value when cloning objects
	IsCloneable bool `json:"is_cloneable,omitempty" mapstructure:"is_cloneable,omitempty"`
	// Label: Name of the field as displayed to users (if not provided, the field's name will be used)
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Name: Internal field name
	Name string `json:"name" mapstructure:"name"`
	// ObjectType:
	ObjectType string `json:"object_type,omitempty" mapstructure:"object_type,omitempty"`
	// Required: If true, this field is required when creating new objects or editing an existing object.
	Required bool `json:"required,omitempty" mapstructure:"required,omitempty"`
	// SearchWeight: Weighting for search. Lower values are considered more important. Fields with a search weight of zero will be ignored.
	SearchWeight int32 `json:"search_weight,omitempty" mapstructure:"search_weight,omitempty"`
	// Type: * `text` - Text
	// * `longtext` - Text (long)
	// * `integer` - Integer
	// * `decimal` - Decimal
	// * `boolean` - Boolean (true/false)
	// * `date` - Date
	// * `datetime` - Date & time
	// * `url` - URL
	// * `json` - JSON
	// * `select` - Selection
	// * `multiselect` - Multiple selection
	// * `object` - Object
	// * `multiobject` - Multiple objects
	Type string `json:"type" mapstructure:"type"`
	// UiVisibility: * `read-write` - Read/Write
	// * `read-only` - Read-only
	// * `hidden` - Hidden
	UiVisibility string `json:"ui_visibility,omitempty" mapstructure:"ui_visibility,omitempty"`
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
func (m CustomFieldRequest) Validate() error {
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
			m.Name, validation.Required, validation.Length(1, 50), validation.Match(customFieldRequestNamePattern),
		),
		"searchWeight": validation.Validate(
			m.SearchWeight, validation.Min(int32(0)), validation.Max(int32(32767)),
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
func (m CustomFieldRequest) GetChoices() []string {
	return m.Choices
}

// SetChoices sets the Choices property
func (m *CustomFieldRequest) SetChoices(val []string) {
	m.Choices = val
}

// GetContentTypes returns the ContentTypes property
func (m CustomFieldRequest) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *CustomFieldRequest) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetDefault returns the Default property
func (m CustomFieldRequest) GetDefault() map[string]interface{} {
	return m.Default
}

// SetDefault sets the Default property
func (m *CustomFieldRequest) SetDefault(val map[string]interface{}) {
	m.Default = val
}

// GetDescription returns the Description property
func (m CustomFieldRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CustomFieldRequest) SetDescription(val string) {
	m.Description = val
}

// GetFilterLogic returns the FilterLogic property
func (m CustomFieldRequest) GetFilterLogic() string {
	return m.FilterLogic
}

// SetFilterLogic sets the FilterLogic property
func (m *CustomFieldRequest) SetFilterLogic(val string) {
	m.FilterLogic = val
}

// GetGroupName returns the GroupName property
func (m CustomFieldRequest) GetGroupName() string {
	return m.GroupName
}

// SetGroupName sets the GroupName property
func (m *CustomFieldRequest) SetGroupName(val string) {
	m.GroupName = val
}

// GetIsCloneable returns the IsCloneable property
func (m CustomFieldRequest) GetIsCloneable() bool {
	return m.IsCloneable
}

// SetIsCloneable sets the IsCloneable property
func (m *CustomFieldRequest) SetIsCloneable(val bool) {
	m.IsCloneable = val
}

// GetLabel returns the Label property
func (m CustomFieldRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *CustomFieldRequest) SetLabel(val string) {
	m.Label = val
}

// GetName returns the Name property
func (m CustomFieldRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *CustomFieldRequest) SetName(val string) {
	m.Name = val
}

// GetObjectType returns the ObjectType property
func (m CustomFieldRequest) GetObjectType() string {
	return m.ObjectType
}

// SetObjectType sets the ObjectType property
func (m *CustomFieldRequest) SetObjectType(val string) {
	m.ObjectType = val
}

// GetRequired returns the Required property
func (m CustomFieldRequest) GetRequired() bool {
	return m.Required
}

// SetRequired sets the Required property
func (m *CustomFieldRequest) SetRequired(val bool) {
	m.Required = val
}

// GetSearchWeight returns the SearchWeight property
func (m CustomFieldRequest) GetSearchWeight() int32 {
	return m.SearchWeight
}

// SetSearchWeight sets the SearchWeight property
func (m *CustomFieldRequest) SetSearchWeight(val int32) {
	m.SearchWeight = val
}

// GetType returns the Type property
func (m CustomFieldRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *CustomFieldRequest) SetType(val string) {
	m.Type = val
}

// GetUiVisibility returns the UiVisibility property
func (m CustomFieldRequest) GetUiVisibility() string {
	return m.UiVisibility
}

// SetUiVisibility sets the UiVisibility property
func (m *CustomFieldRequest) SetUiVisibility(val string) {
	m.UiVisibility = val
}

// GetValidationMaximum returns the ValidationMaximum property
func (m CustomFieldRequest) GetValidationMaximum() *int32 {
	return m.ValidationMaximum
}

// SetValidationMaximum sets the ValidationMaximum property
func (m *CustomFieldRequest) SetValidationMaximum(val *int32) {
	m.ValidationMaximum = val
}

// GetValidationMinimum returns the ValidationMinimum property
func (m CustomFieldRequest) GetValidationMinimum() *int32 {
	return m.ValidationMinimum
}

// SetValidationMinimum sets the ValidationMinimum property
func (m *CustomFieldRequest) SetValidationMinimum(val *int32) {
	m.ValidationMinimum = val
}

// GetValidationRegex returns the ValidationRegex property
func (m CustomFieldRequest) GetValidationRegex() string {
	return m.ValidationRegex
}

// SetValidationRegex sets the ValidationRegex property
func (m *CustomFieldRequest) SetValidationRegex(val string) {
	m.ValidationRegex = val
}

// GetWeight returns the Weight property
func (m CustomFieldRequest) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *CustomFieldRequest) SetWeight(val int32) {
	m.Weight = val
}
