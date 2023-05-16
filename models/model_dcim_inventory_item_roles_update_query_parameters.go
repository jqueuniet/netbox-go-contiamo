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

// DcimInventoryItemRolesUpdateQueryParameters is an object.
type DcimInventoryItemRolesUpdateQueryParameters struct {
	// Id: A unique integer value identifying this inventory item role.
	Id int32 `json:"id" mapstructure:"id"`
}

// Validate implements basic validation for this model
func (m DcimInventoryItemRolesUpdateQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m DcimInventoryItemRolesUpdateQueryParameters) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimInventoryItemRolesUpdateQueryParameters) SetId(val int32) {
	m.Id = val
}
