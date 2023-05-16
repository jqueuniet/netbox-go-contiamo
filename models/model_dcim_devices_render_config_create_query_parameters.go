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

// DcimDevicesRenderConfigCreateQueryParameters is an object.
type DcimDevicesRenderConfigCreateQueryParameters struct {
	// Format:
	Format string `json:"format,omitempty" mapstructure:"format,omitempty"`
	// Id: A unique integer value identifying this device.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m DcimDevicesRenderConfigCreateQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetFormat returns the Format property
func (m DcimDevicesRenderConfigCreateQueryParameters) GetFormat() string {
	return m.Format
}

// SetFormat sets the Format property
func (m *DcimDevicesRenderConfigCreateQueryParameters) SetFormat(val string) {
	m.Format = val
}

// GetId returns the Id property
func (m DcimDevicesRenderConfigCreateQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimDevicesRenderConfigCreateQueryParameters) SetId(val int32) {
	m.Id = val
}
