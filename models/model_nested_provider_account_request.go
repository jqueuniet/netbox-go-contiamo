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

// NestedProviderAccountRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedProviderAccountRequest struct {
	// Account:
	Account string `json:"account" mapstructure:"account"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m NestedProviderAccountRequest) Validate() error {
	return validation.Errors{
		"account": validation.Validate(
			m.Account, validation.Required, validation.Length(1, 100),
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 100),
		),
	}.Filter()
}

// GetAccount returns the Account property
func (m NestedProviderAccountRequest) GetAccount() string {
	return m.Account
}

// SetAccount sets the Account property
func (m *NestedProviderAccountRequest) SetAccount(val string) {
	m.Account = val
}

// GetName returns the Name property
func (m NestedProviderAccountRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedProviderAccountRequest) SetName(val string) {
	m.Name = val
}
