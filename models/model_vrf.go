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

// VRF is an object. Adds support for custom fields and tags.
type VRF struct {
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
	// EnforceUnique: Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique bool `json:"enforce_unique,omitempty" mapstructure:"enforce_unique,omitempty"`
	// ExportTargets:
	ExportTargets []int32 `json:"export_targets,omitempty" mapstructure:"export_targets,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// ImportTargets:
	ImportTargets []int32 `json:"import_targets,omitempty" mapstructure:"import_targets,omitempty"`
	// IpaddressCount:
	IpaddressCount int32 `json:"ipaddress_count" mapstructure:"ipaddress_count"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PrefixCount:
	PrefixCount int32 `json:"prefix_count" mapstructure:"prefix_count"`
	// Rd: Unique route distinguisher (as defined in RFC 4364)
	Rd *string `json:"rd,omitempty" mapstructure:"rd,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m VRF) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"exportTargets": validation.Validate(
			m.ExportTargets,
		),
		"importTargets": validation.Validate(
			m.ImportTargets,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"rd": validation.Validate(
			m.Rd, validation.Length(0, 21),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m VRF) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VRF) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m VRF) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VRF) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m VRF) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VRF) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VRF) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VRF) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m VRF) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *VRF) SetDisplay(val string) {
	m.Display = val
}

// GetEnforceUnique returns the EnforceUnique property
func (m VRF) GetEnforceUnique() bool {
	return m.EnforceUnique
}

// SetEnforceUnique sets the EnforceUnique property
func (m *VRF) SetEnforceUnique(val bool) {
	m.EnforceUnique = val
}

// GetExportTargets returns the ExportTargets property
func (m VRF) GetExportTargets() []int32 {
	return m.ExportTargets
}

// SetExportTargets sets the ExportTargets property
func (m *VRF) SetExportTargets(val []int32) {
	m.ExportTargets = val
}

// GetId returns the Id property
func (m VRF) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VRF) SetId(val int32) {
	m.Id = val
}

// GetImportTargets returns the ImportTargets property
func (m VRF) GetImportTargets() []int32 {
	return m.ImportTargets
}

// SetImportTargets sets the ImportTargets property
func (m *VRF) SetImportTargets(val []int32) {
	m.ImportTargets = val
}

// GetIpaddressCount returns the IpaddressCount property
func (m VRF) GetIpaddressCount() int32 {
	return m.IpaddressCount
}

// SetIpaddressCount sets the IpaddressCount property
func (m *VRF) SetIpaddressCount(val int32) {
	m.IpaddressCount = val
}

// GetLastUpdated returns the LastUpdated property
func (m VRF) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VRF) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m VRF) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VRF) SetName(val string) {
	m.Name = val
}

// GetPrefixCount returns the PrefixCount property
func (m VRF) GetPrefixCount() int32 {
	return m.PrefixCount
}

// SetPrefixCount sets the PrefixCount property
func (m *VRF) SetPrefixCount(val int32) {
	m.PrefixCount = val
}

// GetRd returns the Rd property
func (m VRF) GetRd() *string {
	return m.Rd
}

// SetRd sets the Rd property
func (m *VRF) SetRd(val *string) {
	m.Rd = val
}

// GetTags returns the Tags property
func (m VRF) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VRF) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VRF) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VRF) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m VRF) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *VRF) SetUrl(val string) {
	m.Url = val
}
