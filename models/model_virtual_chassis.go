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

// VirtualChassis is an object. Adds support for custom fields and tags.
type VirtualChassis struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Domain:
	Domain string `json:"domain,omitempty" mapstructure:"domain,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Master: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Master *NestedDevice `json:"master,omitempty" mapstructure:"master,omitempty"`
	// MemberCount:
	MemberCount int32 `json:"member_count" mapstructure:"member_count"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m VirtualChassis) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"domain": validation.Validate(
			m.Domain, validation.Length(0, 30),
		),
		"master": validation.Validate(
			m.Master,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m VirtualChassis) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VirtualChassis) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m VirtualChassis) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VirtualChassis) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m VirtualChassis) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VirtualChassis) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VirtualChassis) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VirtualChassis) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m VirtualChassis) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *VirtualChassis) SetDisplay(val string) {
	m.Display = val
}

// GetDomain returns the Domain property
func (m VirtualChassis) GetDomain() string {
	return m.Domain
}

// SetDomain sets the Domain property
func (m *VirtualChassis) SetDomain(val string) {
	m.Domain = val
}

// GetId returns the Id property
func (m VirtualChassis) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VirtualChassis) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m VirtualChassis) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VirtualChassis) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMaster returns the Master property
func (m VirtualChassis) GetMaster() *NestedDevice {
	return m.Master
}

// SetMaster sets the Master property
func (m *VirtualChassis) SetMaster(val *NestedDevice) {
	m.Master = val
}

// GetMemberCount returns the MemberCount property
func (m VirtualChassis) GetMemberCount() int32 {
	return m.MemberCount
}

// SetMemberCount sets the MemberCount property
func (m *VirtualChassis) SetMemberCount(val int32) {
	m.MemberCount = val
}

// GetName returns the Name property
func (m VirtualChassis) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualChassis) SetName(val string) {
	m.Name = val
}

// GetTags returns the Tags property
func (m VirtualChassis) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VirtualChassis) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m VirtualChassis) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *VirtualChassis) SetUrl(val string) {
	m.Url = val
}
