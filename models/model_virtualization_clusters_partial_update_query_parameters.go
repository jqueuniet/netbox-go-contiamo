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

// VirtualizationClustersPartialUpdateQueryParameters is an object.
type VirtualizationClustersPartialUpdateQueryParameters struct {
	// Id: A unique integer value identifying this cluster.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m VirtualizationClustersPartialUpdateQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m VirtualizationClustersPartialUpdateQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VirtualizationClustersPartialUpdateQueryParameters) SetId(val int32) {
	m.Id = val
}
