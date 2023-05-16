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

// ComponentNestedModuleRequest is an object. Used by device component serializers.
type ComponentNestedModuleRequest struct {
	// Device:
	Device int32 `json:"device" mapstructure:"device"`
}

// Validate implements basic validation for this model
func (m ComponentNestedModuleRequest) Validate() error {
	return validation.Errors{}.Filter()
}

// GetDevice returns the Device property
func (m ComponentNestedModuleRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *ComponentNestedModuleRequest) SetDevice(val int32) {
	m.Device = val
}
