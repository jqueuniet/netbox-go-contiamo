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

// PatchedWritableL2VPNTerminationRequest is an object. Adds support for custom fields and tags.
type PatchedWritableL2VPNTerminationRequest struct {
	// AssignedObjectId:
	AssignedObjectId int64 `json:"assigned_object_id,omitempty" mapstructure:"assigned_object_id,omitempty"`
	// AssignedObjectType:
	AssignedObjectType string `json:"assigned_object_type,omitempty" mapstructure:"assigned_object_type,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// L2vpn:
	L2vpn int32 `json:"l2vpn,omitempty" mapstructure:"l2vpn,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableL2VPNTerminationRequest) Validate() error {
	return validation.Errors{
		"assignedObjectId": validation.Validate(
			m.AssignedObjectId, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetAssignedObjectId returns the AssignedObjectId property
func (m PatchedWritableL2VPNTerminationRequest) GetAssignedObjectId() int64 {
	return m.AssignedObjectId
}

// SetAssignedObjectId sets the AssignedObjectId property
func (m *PatchedWritableL2VPNTerminationRequest) SetAssignedObjectId(val int64) {
	m.AssignedObjectId = val
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m PatchedWritableL2VPNTerminationRequest) GetAssignedObjectType() string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *PatchedWritableL2VPNTerminationRequest) SetAssignedObjectType(val string) {
	m.AssignedObjectType = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableL2VPNTerminationRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableL2VPNTerminationRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetL2vpn returns the L2vpn property
func (m PatchedWritableL2VPNTerminationRequest) GetL2vpn() int32 {
	return m.L2vpn
}

// SetL2vpn sets the L2vpn property
func (m *PatchedWritableL2VPNTerminationRequest) SetL2vpn(val int32) {
	m.L2vpn = val
}

// GetTags returns the Tags property
func (m PatchedWritableL2VPNTerminationRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableL2VPNTerminationRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
