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

// ContentType is an object.
type ContentType struct {
	// AppLabel:
	AppLabel string `json:"app_label" mapstructure:"app_label"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Model:
	Model string `json:"model" mapstructure:"model"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ContentType) Validate() error {
	return validation.Errors{
		"appLabel": validation.Validate(
			m.AppLabel, validation.NotNil, validation.Length(0, 100),
		),
		"model": validation.Validate(
			m.Model, validation.NotNil, validation.Length(0, 100),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAppLabel returns the AppLabel property
func (m ContentType) GetAppLabel() string {
	return m.AppLabel
}

// SetAppLabel sets the AppLabel property
func (m *ContentType) SetAppLabel(val string) {
	m.AppLabel = val
}

// GetDisplay returns the Display property
func (m ContentType) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ContentType) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ContentType) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ContentType) SetId(val int32) {
	m.Id = val
}

// GetModel returns the Model property
func (m ContentType) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *ContentType) SetModel(val string) {
	m.Model = val
}

// GetUrl returns the Url property
func (m ContentType) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ContentType) SetUrl(val string) {
	m.Url = val
}
