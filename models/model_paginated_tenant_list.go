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

// PaginatedTenantList is an object.
type PaginatedTenantList struct {
	// Count:
	Count int32 `json:"count,omitempty" mapstructure:"count,omitempty"`
	// Next:
	Next *string `json:"next,omitempty" mapstructure:"next,omitempty"`
	// Previous:
	Previous *string `json:"previous,omitempty" mapstructure:"previous,omitempty"`
	// Results:
	Results []Tenant `json:"results,omitempty" mapstructure:"results,omitempty"`
}

// Validate implements basic validation for this model
func (m PaginatedTenantList) Validate() error {
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
func (m PaginatedTenantList) GetCount() int32 {
	return m.Count
}

// SetCount sets the Count property
func (m *PaginatedTenantList) SetCount(val int32) {
	m.Count = val
}

// GetNext returns the Next property
func (m PaginatedTenantList) GetNext() *string {
	return m.Next
}

// SetNext sets the Next property
func (m *PaginatedTenantList) SetNext(val *string) {
	m.Next = val
}

// GetPrevious returns the Previous property
func (m PaginatedTenantList) GetPrevious() *string {
	return m.Previous
}

// SetPrevious sets the Previous property
func (m *PaginatedTenantList) SetPrevious(val *string) {
	m.Previous = val
}

// GetResults returns the Results property
func (m PaginatedTenantList) GetResults() []Tenant {
	return m.Results
}

// SetResults sets the Results property
func (m *PaginatedTenantList) SetResults(val []Tenant) {
	m.Results = val
}
