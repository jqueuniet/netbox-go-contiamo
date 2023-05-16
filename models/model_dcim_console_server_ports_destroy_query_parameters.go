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

// DcimConsoleServerPortsDestroyQueryParameters is an object.
type DcimConsoleServerPortsDestroyQueryParameters struct {
	// Id: A unique integer value identifying this console server port.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m DcimConsoleServerPortsDestroyQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m DcimConsoleServerPortsDestroyQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimConsoleServerPortsDestroyQueryParameters) SetId(val int32) {
	m.Id = val
}
