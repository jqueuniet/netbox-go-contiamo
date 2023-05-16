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

// WritableFHRPGroupAssignmentRequest is an object. Adds support for custom fields and tags.
type WritableFHRPGroupAssignmentRequest struct {
	// Group:
	Group int32 `json:"group" mapstructure:"group"`
	// InterfaceId:
	InterfaceId int64 `json:"interface_id" mapstructure:"interface_id"`
	// InterfaceType:
	InterfaceType string `json:"interface_type" mapstructure:"interface_type"`
	// Priority:
	Priority int32 `json:"priority" mapstructure:"priority"`
}

// Validate implements basic validation for this model
func (m WritableFHRPGroupAssignmentRequest) Validate() error {
	return validation.Errors{
		"interfaceId": validation.Validate(
			m.InterfaceId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"priority": validation.Validate(
			m.Priority, validation.NotNil, validation.Min(int32(0)), validation.Max(int32(255)),
		),
	}.Filter()
}

// GetGroup returns the Group property
func (m WritableFHRPGroupAssignmentRequest) GetGroup() int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *WritableFHRPGroupAssignmentRequest) SetGroup(val int32) {
	m.Group = val
}

// GetInterfaceId returns the InterfaceId property
func (m WritableFHRPGroupAssignmentRequest) GetInterfaceId() int64 {
	return m.InterfaceId
}

// SetInterfaceId sets the InterfaceId property
func (m *WritableFHRPGroupAssignmentRequest) SetInterfaceId(val int64) {
	m.InterfaceId = val
}

// GetInterfaceType returns the InterfaceType property
func (m WritableFHRPGroupAssignmentRequest) GetInterfaceType() string {
	return m.InterfaceType
}

// SetInterfaceType sets the InterfaceType property
func (m *WritableFHRPGroupAssignmentRequest) SetInterfaceType(val string) {
	m.InterfaceType = val
}

// GetPriority returns the Priority property
func (m WritableFHRPGroupAssignmentRequest) GetPriority() int32 {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *WritableFHRPGroupAssignmentRequest) SetPriority(val int32) {
	m.Priority = val
}
