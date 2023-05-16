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

// Cluster is an object. Adds support for custom fields and tags.
type Cluster struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceCount:
	DeviceCount int32 `json:"device_count" mapstructure:"device_count"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedClusterGroup `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSite `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status:
	Status *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Type NestedClusterType `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VirtualmachineCount:
	VirtualmachineCount int32 `json:"virtualmachine_count" mapstructure:"virtualmachine_count"`
}

// Validate implements basic validation for this model
func (m Cluster) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"group": validation.Validate(
			m.Group,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"site": validation.Validate(
			m.Site,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"type": validation.Validate(
			m.Type, validation.NotNil,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m Cluster) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Cluster) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Cluster) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Cluster) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Cluster) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Cluster) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Cluster) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Cluster) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m Cluster) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *Cluster) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m Cluster) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Cluster) SetDisplay(val string) {
	m.Display = val
}

// GetGroup returns the Group property
func (m Cluster) GetGroup() *NestedClusterGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *Cluster) SetGroup(val *NestedClusterGroup) {
	m.Group = val
}

// GetId returns the Id property
func (m Cluster) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Cluster) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Cluster) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Cluster) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Cluster) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Cluster) SetName(val string) {
	m.Name = val
}

// GetSite returns the Site property
func (m Cluster) GetSite() *NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *Cluster) SetSite(val *NestedSite) {
	m.Site = val
}

// GetStatus returns the Status property
func (m Cluster) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Cluster) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Cluster) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Cluster) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Cluster) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Cluster) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetType returns the Type property
func (m Cluster) GetType() NestedClusterType {
	return m.Type
}

// SetType sets the Type property
func (m *Cluster) SetType(val NestedClusterType) {
	m.Type = val
}

// GetUrl returns the Url property
func (m Cluster) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Cluster) SetUrl(val string) {
	m.Url = val
}

// GetVirtualmachineCount returns the VirtualmachineCount property
func (m Cluster) GetVirtualmachineCount() int32 {
	return m.VirtualmachineCount
}

// SetVirtualmachineCount sets the VirtualmachineCount property
func (m *Cluster) SetVirtualmachineCount(val int32) {
	m.VirtualmachineCount = val
}
