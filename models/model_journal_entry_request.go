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

// JournalEntryRequest is an object. Adds support for custom fields and tags.
type JournalEntryRequest struct {
	// AssignedObjectId:
	AssignedObjectId int64 `json:"assigned_object_id" mapstructure:"assigned_object_id"`
	// AssignedObjectType:
	AssignedObjectType string `json:"assigned_object_type" mapstructure:"assigned_object_type"`
	// Comments:
	Comments string `json:"comments" mapstructure:"comments"`
	// CreatedBy:
	CreatedBy *int32 `json:"created_by,omitempty" mapstructure:"created_by,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Kind: * `info` - Info
	// * `success` - Success
	// * `warning` - Warning
	// * `danger` - Danger
	Kind string `json:"kind,omitempty" mapstructure:"kind,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m JournalEntryRequest) Validate() error {
	return validation.Errors{
		"assignedObjectId": validation.Validate(
			m.AssignedObjectId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"comments": validation.Validate(
			m.Comments, validation.Required, validation.Length(1, 0),
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
func (m JournalEntryRequest) GetAssignedObjectId() int64 {
	return m.AssignedObjectId
}

// SetAssignedObjectId sets the AssignedObjectId property
func (m *JournalEntryRequest) SetAssignedObjectId(val int64) {
	m.AssignedObjectId = val
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m JournalEntryRequest) GetAssignedObjectType() string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *JournalEntryRequest) SetAssignedObjectType(val string) {
	m.AssignedObjectType = val
}

// GetComments returns the Comments property
func (m JournalEntryRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *JournalEntryRequest) SetComments(val string) {
	m.Comments = val
}

// GetCreatedBy returns the CreatedBy property
func (m JournalEntryRequest) GetCreatedBy() *int32 {
	return m.CreatedBy
}

// SetCreatedBy sets the CreatedBy property
func (m *JournalEntryRequest) SetCreatedBy(val *int32) {
	m.CreatedBy = val
}

// GetCustomFields returns the CustomFields property
func (m JournalEntryRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *JournalEntryRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetKind returns the Kind property
func (m JournalEntryRequest) GetKind() string {
	return m.Kind
}

// SetKind sets the Kind property
func (m *JournalEntryRequest) SetKind(val string) {
	m.Kind = val
}

// GetTags returns the Tags property
func (m JournalEntryRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *JournalEntryRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}
