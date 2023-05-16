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

// CircuitsProviderAccountsUpdateQueryParameters is an object.
type CircuitsProviderAccountsUpdateQueryParameters struct {
	// Id: A unique integer value identifying this provider account.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m CircuitsProviderAccountsUpdateQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m CircuitsProviderAccountsUpdateQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitsProviderAccountsUpdateQueryParameters) SetId(val int32) {
	m.Id = val
}
