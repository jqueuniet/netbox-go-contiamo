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

// ContactAssignment is an object. Adds support for custom fields and tags.
type ContactAssignment struct {
	// Contact: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Contact NestedContact `json:"contact" mapstructure:"contact"`
	// ContentType:
	ContentType string `json:"content_type" mapstructure:"content_type"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Object:
	Object map[string]interface{} `json:"object" mapstructure:"object"`
	// ObjectId:
	ObjectId int64 `json:"object_id" mapstructure:"object_id"`
	// Priority:
	Priority *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"priority,omitempty" mapstructure:"priority,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedContactRole `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ContactAssignment) Validate() error {
	return validation.Errors{
		"contact": validation.Validate(
			m.Contact, validation.NotNil,
		),
		"object": validation.Validate(
			m.Object, validation.NotNil,
		),
		"objectId": validation.Validate(
			m.ObjectId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"role": validation.Validate(
			m.Role,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetContact returns the Contact property
func (m ContactAssignment) GetContact() NestedContact {
	return m.Contact
}

// SetContact sets the Contact property
func (m *ContactAssignment) SetContact(val NestedContact) {
	m.Contact = val
}

// GetContentType returns the ContentType property
func (m ContactAssignment) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *ContactAssignment) SetContentType(val string) {
	m.ContentType = val
}

// GetCreated returns the Created property
func (m ContactAssignment) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ContactAssignment) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDisplay returns the Display property
func (m ContactAssignment) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ContactAssignment) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ContactAssignment) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ContactAssignment) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ContactAssignment) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ContactAssignment) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetObject returns the Object property
func (m ContactAssignment) GetObject() map[string]interface{} {
	return m.Object
}

// SetObject sets the Object property
func (m *ContactAssignment) SetObject(val map[string]interface{}) {
	m.Object = val
}

// GetObjectId returns the ObjectId property
func (m ContactAssignment) GetObjectId() int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *ContactAssignment) SetObjectId(val int64) {
	m.ObjectId = val
}

// GetPriority returns the Priority property
func (m ContactAssignment) GetPriority() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *ContactAssignment) SetPriority(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Priority = val
}

// GetRole returns the Role property
func (m ContactAssignment) GetRole() *NestedContactRole {
	return m.Role
}

// SetRole sets the Role property
func (m *ContactAssignment) SetRole(val *NestedContactRole) {
	m.Role = val
}

// GetUrl returns the Url property
func (m ContactAssignment) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ContactAssignment) SetUrl(val string) {
	m.Url = val
}
