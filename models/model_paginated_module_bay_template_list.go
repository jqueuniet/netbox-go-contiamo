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

// PaginatedModuleBayTemplateList is an object.
type PaginatedModuleBayTemplateList struct {
	// Count:
	Count int32 `json:"count,omitempty" mapstructure:"count,omitempty"`
	// Next:
	Next *string `json:"next,omitempty" mapstructure:"next,omitempty"`
	// Previous:
	Previous *string `json:"previous,omitempty" mapstructure:"previous,omitempty"`
	// Results:
	Results []ModuleBayTemplate `json:"results,omitempty" mapstructure:"results,omitempty"`
}

// Validate implements basic validation for this model
func (m PaginatedModuleBayTemplateList) Validate() error {
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
func (m PaginatedModuleBayTemplateList) GetCount() int32 {
	return m.Count
}

// SetCount sets the Count property
func (m *PaginatedModuleBayTemplateList) SetCount(val int32) {
	m.Count = val
}

// GetNext returns the Next property
func (m PaginatedModuleBayTemplateList) GetNext() *string {
	return m.Next
}

// SetNext sets the Next property
func (m *PaginatedModuleBayTemplateList) SetNext(val *string) {
	m.Next = val
}

// GetPrevious returns the Previous property
func (m PaginatedModuleBayTemplateList) GetPrevious() *string {
	return m.Previous
}

// SetPrevious sets the Previous property
func (m *PaginatedModuleBayTemplateList) SetPrevious(val *string) {
	m.Previous = val
}

// GetResults returns the Results property
func (m PaginatedModuleBayTemplateList) GetResults() []ModuleBayTemplate {
	return m.Results
}

// SetResults sets the Results property
func (m *PaginatedModuleBayTemplateList) SetResults(val []ModuleBayTemplate) {
	m.Results = val
}
