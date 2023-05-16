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

// ExtrasContentTypesListQueryParameters is an object.
type ExtrasContentTypesListQueryParameters struct {
	// AppLabel:
	AppLabel string `json:"app_label,omitempty" mapstructure:"app_label,omitempty"`
	// Id:
	Id int32 `json:"id,omitempty" mapstructure:"id,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// Model:
	Model string `json:"model,omitempty" mapstructure:"model,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasContentTypesListQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAppLabel returns the AppLabel property
func (m ExtrasContentTypesListQueryParameters) GetAppLabel() string {
	return m.AppLabel
}

// SetAppLabel sets the AppLabel property
func (m *ExtrasContentTypesListQueryParameters) SetAppLabel(val string) {
	m.AppLabel = val
}

// GetId returns the Id property
func (m ExtrasContentTypesListQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasContentTypesListQueryParameters) SetId(val int32) {
	m.Id = val
}

// GetLimit returns the Limit property
func (m ExtrasContentTypesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasContentTypesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetModel returns the Model property
func (m ExtrasContentTypesListQueryParameters) GetModel() string {
	return m.Model
}

// SetModel sets the Model property
func (m *ExtrasContentTypesListQueryParameters) SetModel(val string) {
	m.Model = val
}

// GetOffset returns the Offset property
func (m ExtrasContentTypesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasContentTypesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasContentTypesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasContentTypesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasContentTypesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasContentTypesListQueryParameters) SetQ(val string) {
	m.Q = val
}
