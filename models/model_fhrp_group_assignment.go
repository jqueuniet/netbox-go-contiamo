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

// FHRPGroupAssignment is an object. Adds support for custom fields and tags.
type FHRPGroupAssignment struct {
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group NestedFHRPGroup `json:"group" mapstructure:"group"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Interface:
	Interface map[string]interface{} `json:"interface" mapstructure:"interface"`
	// InterfaceId:
	InterfaceId int64 `json:"interface_id" mapstructure:"interface_id"`
	// InterfaceType:
	InterfaceType string `json:"interface_type" mapstructure:"interface_type"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Priority:
	Priority int32 `json:"priority" mapstructure:"priority"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m FHRPGroupAssignment) Validate() error {
	return validation.Errors{
		"group": validation.Validate(
			m.Group, validation.NotNil,
		),
		"interface": validation.Validate(
			m.Interface,
		),
		"interfaceId": validation.Validate(
			m.InterfaceId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"priority": validation.Validate(
			m.Priority, validation.NotNil, validation.Min(int32(0)), validation.Max(int32(255)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m FHRPGroupAssignment) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *FHRPGroupAssignment) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDisplay returns the Display property
func (m FHRPGroupAssignment) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *FHRPGroupAssignment) SetDisplay(val string) {
	m.Display = val
}

// GetGroup returns the Group property
func (m FHRPGroupAssignment) GetGroup() NestedFHRPGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *FHRPGroupAssignment) SetGroup(val NestedFHRPGroup) {
	m.Group = val
}

// GetId returns the Id property
func (m FHRPGroupAssignment) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *FHRPGroupAssignment) SetId(val int32) {
	m.Id = val
}

// GetInterface returns the Interface property
func (m FHRPGroupAssignment) GetInterface() map[string]interface{} {
	return m.Interface
}

// SetInterface sets the Interface property
func (m *FHRPGroupAssignment) SetInterface(val map[string]interface{}) {
	m.Interface = val
}

// GetInterfaceId returns the InterfaceId property
func (m FHRPGroupAssignment) GetInterfaceId() int64 {
	return m.InterfaceId
}

// SetInterfaceId sets the InterfaceId property
func (m *FHRPGroupAssignment) SetInterfaceId(val int64) {
	m.InterfaceId = val
}

// GetInterfaceType returns the InterfaceType property
func (m FHRPGroupAssignment) GetInterfaceType() string {
	return m.InterfaceType
}

// SetInterfaceType sets the InterfaceType property
func (m *FHRPGroupAssignment) SetInterfaceType(val string) {
	m.InterfaceType = val
}

// GetLastUpdated returns the LastUpdated property
func (m FHRPGroupAssignment) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *FHRPGroupAssignment) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetPriority returns the Priority property
func (m FHRPGroupAssignment) GetPriority() int32 {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *FHRPGroupAssignment) SetPriority(val int32) {
	m.Priority = val
}

// GetUrl returns the Url property
func (m FHRPGroupAssignment) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *FHRPGroupAssignment) SetUrl(val string) {
	m.Url = val
}
