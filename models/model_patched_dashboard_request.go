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

// PatchedDashboardRequest is an object.
type PatchedDashboardRequest struct {
	// Config:
	Config map[string]interface{} `json:"config,omitempty" mapstructure:"config,omitempty"`
	// Layout:
	Layout map[string]interface{} `json:"layout,omitempty" mapstructure:"layout,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedDashboardRequest) Validate() error {
	return validation.Errors{
		"config": validation.Validate(
			m.Config,
		),
		"layout": validation.Validate(
			m.Layout,
		),
	}.Filter()
}

// GetConfig returns the Config property
func (m PatchedDashboardRequest) GetConfig() map[string]interface{} {
	return m.Config
}

// SetConfig sets the Config property
func (m *PatchedDashboardRequest) SetConfig(val map[string]interface{}) {
	m.Config = val
}

// GetLayout returns the Layout property
func (m PatchedDashboardRequest) GetLayout() map[string]interface{} {
	return m.Layout
}

// SetLayout sets the Layout property
func (m *PatchedDashboardRequest) SetLayout(val map[string]interface{}) {
	m.Layout = val
}
