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

// PatchedWritableContactAssignmentRequest is an object. Adds support for custom fields and tags.
type PatchedWritableContactAssignmentRequest struct {
	// Contact:
	Contact int32 `json:"contact,omitempty" mapstructure:"contact,omitempty"`
	// ContentType:
	ContentType string `json:"content_type,omitempty" mapstructure:"content_type,omitempty"`
	// ObjectId:
	ObjectId int64 `json:"object_id,omitempty" mapstructure:"object_id,omitempty"`
	// Priority: * `primary` - Primary
	// * `secondary` - Secondary
	// * `tertiary` - Tertiary
	// * `inactive` - Inactive
	Priority string `json:"priority,omitempty" mapstructure:"priority,omitempty"`
	// Role:
	Role int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableContactAssignmentRequest) Validate() error {
	return validation.Errors{
		"objectId": validation.Validate(
			m.ObjectId, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
	}.Filter()
}

// GetContact returns the Contact property
func (m PatchedWritableContactAssignmentRequest) GetContact() int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *PatchedWritableContactAssignmentRequest) SetContact(val int32) {
	m.Contact = val
}

// GetContentType returns the ContentType property
func (m PatchedWritableContactAssignmentRequest) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *PatchedWritableContactAssignmentRequest) SetContentType(val string) {
	m.ContentType = val
}

// GetObjectId returns the ObjectId property
func (m PatchedWritableContactAssignmentRequest) GetObjectId() int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *PatchedWritableContactAssignmentRequest) SetObjectId(val int64) {
	m.ObjectId = val
}

// GetPriority returns the Priority property
func (m PatchedWritableContactAssignmentRequest) GetPriority() string {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *PatchedWritableContactAssignmentRequest) SetPriority(val string) {
	m.Priority = val
}

// GetRole returns the Role property
func (m PatchedWritableContactAssignmentRequest) GetRole() int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *PatchedWritableContactAssignmentRequest) SetRole(val int32) {
	m.Role = val
}
