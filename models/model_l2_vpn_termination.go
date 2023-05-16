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

// L2VPNTermination is an object. Adds support for custom fields and tags.
type L2VPNTermination struct {
	// AssignedObject:
	AssignedObject map[string]interface{} `json:"assigned_object" mapstructure:"assigned_object"`
	// AssignedObjectId:
	AssignedObjectId int64 `json:"assigned_object_id" mapstructure:"assigned_object_id"`
	// AssignedObjectType:
	AssignedObjectType string `json:"assigned_object_type" mapstructure:"assigned_object_type"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// L2vpn: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	L2vpn NestedL2VPN `json:"l2vpn" mapstructure:"l2vpn"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m L2VPNTermination) Validate() error {
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
		"l2vpn": validation.Validate(
			m.L2vpn, validation.NotNil,
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
func (m L2VPNTermination) GetAssignedObject() map[string]interface{} {
	return m.AssignedObject
}

// SetAssignedObject sets the AssignedObject property
func (m *L2VPNTermination) SetAssignedObject(val map[string]interface{}) {
	m.AssignedObject = val
}

// GetAssignedObjectId returns the AssignedObjectId property
func (m L2VPNTermination) GetAssignedObjectId() int64 {
	return m.AssignedObjectId
}

// SetAssignedObjectId sets the AssignedObjectId property
func (m *L2VPNTermination) SetAssignedObjectId(val int64) {
	m.AssignedObjectId = val
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m L2VPNTermination) GetAssignedObjectType() string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *L2VPNTermination) SetAssignedObjectType(val string) {
	m.AssignedObjectType = val
}

// GetCreated returns the Created property
func (m L2VPNTermination) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *L2VPNTermination) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m L2VPNTermination) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *L2VPNTermination) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDisplay returns the Display property
func (m L2VPNTermination) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *L2VPNTermination) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m L2VPNTermination) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *L2VPNTermination) SetId(val int32) {
	m.Id = val
}

// GetL2vpn returns the L2vpn property
func (m L2VPNTermination) GetL2vpn() NestedL2VPN {
	return m.L2vpn
}

// SetL2vpn sets the L2vpn property
func (m *L2VPNTermination) SetL2vpn(val NestedL2VPN) {
	m.L2vpn = val
}

// GetLastUpdated returns the LastUpdated property
func (m L2VPNTermination) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *L2VPNTermination) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetTags returns the Tags property
func (m L2VPNTermination) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *L2VPNTermination) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m L2VPNTermination) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *L2VPNTermination) SetUrl(val string) {
	m.Url = val
}
