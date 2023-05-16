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

// WritableContactAssignmentRequest is an object. Adds support for custom fields and tags.
type WritableContactAssignmentRequest struct {
	// Contact:
	Contact int32 `json:"contact" mapstructure:"contact"`
	// ContentType:
	ContentType string `json:"content_type" mapstructure:"content_type"`
	// ObjectId:
	ObjectId int64 `json:"object_id" mapstructure:"object_id"`
	// Priority: * `primary` - Primary
	// * `secondary` - Secondary
	// * `tertiary` - Tertiary
	// * `inactive` - Inactive
	Priority string `json:"priority,omitempty" mapstructure:"priority,omitempty"`
	// Role:
	Role int32 `json:"role" mapstructure:"role"`
}

// Validate implements basic validation for this model
func (m WritableContactAssignmentRequest) Validate() error {
	return validation.Errors{
		"objectId": validation.Validate(
			m.ObjectId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
	}.Filter()
}

// GetContact returns the Contact property
func (m WritableContactAssignmentRequest) GetContact() int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *WritableContactAssignmentRequest) SetContact(val int32) {
	m.Contact = val
}

// GetContentType returns the ContentType property
func (m WritableContactAssignmentRequest) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *WritableContactAssignmentRequest) SetContentType(val string) {
	m.ContentType = val
}

// GetObjectId returns the ObjectId property
func (m WritableContactAssignmentRequest) GetObjectId() int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *WritableContactAssignmentRequest) SetObjectId(val int64) {
	m.ObjectId = val
}

// GetPriority returns the Priority property
func (m WritableContactAssignmentRequest) GetPriority() string {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *WritableContactAssignmentRequest) SetPriority(val string) {
	m.Priority = val
}

// GetRole returns the Role property
func (m WritableContactAssignmentRequest) GetRole() int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *WritableContactAssignmentRequest) SetRole(val int32) {
	m.Role = val
}
