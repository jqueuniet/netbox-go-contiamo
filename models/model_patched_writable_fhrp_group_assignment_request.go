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

// PatchedWritableFHRPGroupAssignmentRequest is an object. Adds support for custom fields and tags.
type PatchedWritableFHRPGroupAssignmentRequest struct {
	// Group:
	Group int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// InterfaceId:
	InterfaceId int64 `json:"interface_id,omitempty" mapstructure:"interface_id,omitempty"`
	// InterfaceType:
	InterfaceType string `json:"interface_type,omitempty" mapstructure:"interface_type,omitempty"`
	// Priority:
	Priority int32 `json:"priority,omitempty" mapstructure:"priority,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableFHRPGroupAssignmentRequest) Validate() error {
	return validation.Errors{
		"interfaceId": validation.Validate(
			m.InterfaceId, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"priority": validation.Validate(
			m.Priority, validation.Min(int32(0)), validation.Max(int32(255)),
		),
	}.Filter()
}

// GetGroup returns the Group property
func (m PatchedWritableFHRPGroupAssignmentRequest) GetGroup() int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *PatchedWritableFHRPGroupAssignmentRequest) SetGroup(val int32) {
	m.Group = val
}

// GetInterfaceId returns the InterfaceId property
func (m PatchedWritableFHRPGroupAssignmentRequest) GetInterfaceId() int64 {
	return m.InterfaceId
}

// SetInterfaceId sets the InterfaceId property
func (m *PatchedWritableFHRPGroupAssignmentRequest) SetInterfaceId(val int64) {
	m.InterfaceId = val
}

// GetInterfaceType returns the InterfaceType property
func (m PatchedWritableFHRPGroupAssignmentRequest) GetInterfaceType() string {
	return m.InterfaceType
}

// SetInterfaceType sets the InterfaceType property
func (m *PatchedWritableFHRPGroupAssignmentRequest) SetInterfaceType(val string) {
	m.InterfaceType = val
}

// GetPriority returns the Priority property
func (m PatchedWritableFHRPGroupAssignmentRequest) GetPriority() int32 {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *PatchedWritableFHRPGroupAssignmentRequest) SetPriority(val int32) {
	m.Priority = val
}
