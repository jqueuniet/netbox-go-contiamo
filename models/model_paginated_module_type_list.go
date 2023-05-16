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

// PaginatedModuleTypeList is an object.
type PaginatedModuleTypeList struct {
	// Count:
	Count int32 `json:"count,omitempty" mapstructure:"count,omitempty"`
	// Next:
	Next *string `json:"next,omitempty" mapstructure:"next,omitempty"`
	// Previous:
	Previous *string `json:"previous,omitempty" mapstructure:"previous,omitempty"`
	// Results:
	Results []ModuleType `json:"results,omitempty" mapstructure:"results,omitempty"`
}

// Validate implements basic validation for this model
func (m PaginatedModuleTypeList) Validate() error {
	return validation.Errors{
		"next": validation.Validate(
			m.Next, is.RequestURI,
		),
		"previous": validation.Validate(
			m.Previous, is.RequestURI,
		),
		"results": validation.Validate(
			m.Results,
		),
	}.Filter()
}

// GetCount returns the Count property
func (m PaginatedModuleTypeList) GetCount() int32 {
	return m.Count
}

// SetCount sets the Count property
func (m *PaginatedModuleTypeList) SetCount(val int32) {
	m.Count = val
}

// GetNext returns the Next property
func (m PaginatedModuleTypeList) GetNext() *string {
	return m.Next
}

// SetNext sets the Next property
func (m *PaginatedModuleTypeList) SetNext(val *string) {
	m.Next = val
}

// GetPrevious returns the Previous property
func (m PaginatedModuleTypeList) GetPrevious() *string {
	return m.Previous
}

// SetPrevious sets the Previous property
func (m *PaginatedModuleTypeList) SetPrevious(val *string) {
	m.Previous = val
}

// GetResults returns the Results property
func (m PaginatedModuleTypeList) GetResults() []ModuleType {
	return m.Results
}

// SetResults sets the Results property
func (m *PaginatedModuleTypeList) SetResults(val []ModuleType) {
	m.Results = val
}
