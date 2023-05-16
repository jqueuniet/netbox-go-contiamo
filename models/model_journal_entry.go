// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// JournalEntry is an object. Adds support for custom fields and tags.
type JournalEntry struct {
	// AssignedObject:
	AssignedObject map[string]interface{} `json:"assigned_object" mapstructure:"assigned_object"`
	// AssignedObjectId:
	AssignedObjectId int64 `json:"assigned_object_id" mapstructure:"assigned_object_id"`
	// AssignedObjectType:
	AssignedObjectType string `json:"assigned_object_type" mapstructure:"assigned_object_type"`
	// Comments:
	Comments string `json:"comments" mapstructure:"comments"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CreatedBy:
	CreatedBy *int32 `json:"created_by,omitempty" mapstructure:"created_by,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Kind:
	Kind *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"kind,omitempty" mapstructure:"kind,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m JournalEntry) Validate() error {
	return validation.Errors{
		"assignedObject": validation.Validate(
			m.AssignedObject,
		),
		"assignedObjectId": validation.Validate(
			m.AssignedObjectId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAssignedObject returns the AssignedObject property
func (m JournalEntry) GetAssignedObject() map[string]interface{} {
	return m.AssignedObject
}

// SetAssignedObject sets the AssignedObject property
func (m *JournalEntry) SetAssignedObject(val map[string]interface{}) {
	m.AssignedObject = val
}

// GetAssignedObjectId returns the AssignedObjectId property
func (m JournalEntry) GetAssignedObjectId() int64 {
	return m.AssignedObjectId
}

// SetAssignedObjectId sets the AssignedObjectId property
func (m *JournalEntry) SetAssignedObjectId(val int64) {
	m.AssignedObjectId = val
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m JournalEntry) GetAssignedObjectType() string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *JournalEntry) SetAssignedObjectType(val string) {
	m.AssignedObjectType = val
}

// GetComments returns the Comments property
func (m JournalEntry) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *JournalEntry) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m JournalEntry) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *JournalEntry) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCreatedBy returns the CreatedBy property
func (m JournalEntry) GetCreatedBy() *int32 {
	return m.CreatedBy
}

// SetCreatedBy sets the CreatedBy property
func (m *JournalEntry) SetCreatedBy(val *int32) {
	m.CreatedBy = val
}

// GetCustomFields returns the CustomFields property
func (m JournalEntry) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *JournalEntry) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDisplay returns the Display property
func (m JournalEntry) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *JournalEntry) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m JournalEntry) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *JournalEntry) SetId(val int32) {
	m.Id = val
}

// GetKind returns the Kind property
func (m JournalEntry) GetKind() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Kind
}

// SetKind sets the Kind property
func (m *JournalEntry) SetKind(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Kind = val
}

// GetLastUpdated returns the LastUpdated property
func (m JournalEntry) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *JournalEntry) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetTags returns the Tags property
func (m JournalEntry) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *JournalEntry) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m JournalEntry) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *JournalEntry) SetUrl(val string) {
	m.Url = val
}
