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

// PaginatedContentTypeList is an object.
type PaginatedContentTypeList struct {
	// Count:
	Count int32 `json:"count,omitempty" mapstructure:"count,omitempty"`
	// Next:
	Next *string `json:"next,omitempty" mapstructure:"next,omitempty"`
	// Previous:
	Previous *string `json:"previous,omitempty" mapstructure:"previous,omitempty"`
	// Results:
	Results []ContentType `json:"results,omitempty" mapstructure:"results,omitempty"`
}

// Validate implements basic validation for this model
func (m PaginatedContentTypeList) Validate() error {
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
func (m PaginatedContentTypeList) GetCount() int32 {
	return m.Count
}

// SetCount sets the Count property
func (m *PaginatedContentTypeList) SetCount(val int32) {
	m.Count = val
}

// GetNext returns the Next property
func (m PaginatedContentTypeList) GetNext() *string {
	return m.Next
}

// SetNext sets the Next property
func (m *PaginatedContentTypeList) SetNext(val *string) {
	m.Next = val
}

// GetPrevious returns the Previous property
func (m PaginatedContentTypeList) GetPrevious() *string {
	return m.Previous
}

// SetPrevious sets the Previous property
func (m *PaginatedContentTypeList) SetPrevious(val *string) {
	m.Previous = val
}

// GetResults returns the Results property
func (m PaginatedContentTypeList) GetResults() []ContentType {
	return m.Results
}

// SetResults sets the Results property
func (m *PaginatedContentTypeList) SetResults(val []ContentType) {
	m.Results = val
}
