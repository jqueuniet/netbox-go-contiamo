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

// ContactAssignmentRequest is an object. Adds support for custom fields and tags.
type ContactAssignmentRequest struct {
	// Contact: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Contact NestedContactRequest `json:"contact" mapstructure:"contact"`
	// ContentType:
	ContentType string `json:"content_type" mapstructure:"content_type"`
	// ObjectId:
	ObjectId int64 `json:"object_id" mapstructure:"object_id"`
	// Priority: * `primary` - Primary
	// * `secondary` - Secondary
	// * `tertiary` - Tertiary
	// * `inactive` - Inactive
	Priority string `json:"priority,omitempty" mapstructure:"priority,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedContactRoleRequest `json:"role,omitempty" mapstructure:"role,omitempty"`
}

// Validate implements basic validation for this model
func (m ContactAssignmentRequest) Validate() error {
	return validation.Errors{
		"contact": validation.Validate(
			m.Contact, validation.NotNil,
		),
		"objectId": validation.Validate(
			m.ObjectId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"role": validation.Validate(
			m.Role,
		),
	}.Filter()
}

// GetContact returns the Contact property
func (m ContactAssignmentRequest) GetContact() NestedContactRequest {
	return m.Contact
}

// SetContact sets the Contact property
func (m *ContactAssignmentRequest) SetContact(val NestedContactRequest) {
	m.Contact = val
}

// GetContentType returns the ContentType property
func (m ContactAssignmentRequest) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *ContactAssignmentRequest) SetContentType(val string) {
	m.ContentType = val
}

// GetObjectId returns the ObjectId property
func (m ContactAssignmentRequest) GetObjectId() int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *ContactAssignmentRequest) SetObjectId(val int64) {
	m.ObjectId = val
}

// GetPriority returns the Priority property
func (m ContactAssignmentRequest) GetPriority() string {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *ContactAssignmentRequest) SetPriority(val string) {
	m.Priority = val
}

// GetRole returns the Role property
func (m ContactAssignmentRequest) GetRole() *NestedContactRoleRequest {
	return m.Role
}

// SetRole sets the Role property
func (m *ContactAssignmentRequest) SetRole(val *NestedContactRoleRequest) {
	m.Role = val
}
