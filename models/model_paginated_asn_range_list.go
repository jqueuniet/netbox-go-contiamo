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

// PaginatedASNRangeList is an object.
type PaginatedASNRangeList struct {
	// Count:
	Count int32 `json:"count,omitempty" mapstructure:"count,omitempty"`
	// Next:
	Next *string `json:"next,omitempty" mapstructure:"next,omitempty"`
	// Previous:
	Previous *string `json:"previous,omitempty" mapstructure:"previous,omitempty"`
	// Results:
	Results []ASNRange `json:"results,omitempty" mapstructure:"results,omitempty"`
}

// Validate implements basic validation for this model
func (m PaginatedASNRangeList) Validate() error {
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
func (m PaginatedASNRangeList) GetCount() int32 {
	return m.Count
}

// SetCount sets the Count property
func (m *PaginatedASNRangeList) SetCount(val int32) {
	m.Count = val
}

// GetNext returns the Next property
func (m PaginatedASNRangeList) GetNext() *string {
	return m.Next
}

// SetNext sets the Next property
func (m *PaginatedASNRangeList) SetNext(val *string) {
	m.Next = val
}

// GetPrevious returns the Previous property
func (m PaginatedASNRangeList) GetPrevious() *string {
	return m.Previous
}

// SetPrevious sets the Previous property
func (m *PaginatedASNRangeList) SetPrevious(val *string) {
	m.Previous = val
}

// GetResults returns the Results property
func (m PaginatedASNRangeList) GetResults() []ASNRange {
	return m.Results
}

// SetResults sets the Results property
func (m *PaginatedASNRangeList) SetResults(val []ASNRange) {
	m.Results = val
}
