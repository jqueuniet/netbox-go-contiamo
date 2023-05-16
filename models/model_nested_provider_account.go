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

// NestedProviderAccount is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedProviderAccount struct {
	// Account:
	Account string `json:"account" mapstructure:"account"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedProviderAccount) Validate() error {
	return validation.Errors{
		"account": validation.Validate(
			m.Account, validation.NotNil, validation.Length(0, 100),
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 100),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAccount returns the Account property
func (m NestedProviderAccount) GetAccount() string {
	return m.Account
}

// SetAccount sets the Account property
func (m *NestedProviderAccount) SetAccount(val string) {
	m.Account = val
}

// GetDisplay returns the Display property
func (m NestedProviderAccount) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedProviderAccount) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedProviderAccount) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedProviderAccount) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m NestedProviderAccount) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedProviderAccount) SetName(val string) {
	m.Name = val
}

// GetUrl returns the Url property
func (m NestedProviderAccount) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedProviderAccount) SetUrl(val string) {
	m.Url = val
}
