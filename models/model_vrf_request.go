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

// VRFRequest is an object. Adds support for custom fields and tags.
type VRFRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// EnforceUnique: Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique bool `json:"enforce_unique,omitempty" mapstructure:"enforce_unique,omitempty"`
	// ExportTargets:
	ExportTargets []int32 `json:"export_targets,omitempty" mapstructure:"export_targets,omitempty"`
	// ImportTargets:
	ImportTargets []int32 `json:"import_targets,omitempty" mapstructure:"import_targets,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Rd: Unique route distinguisher (as defined in RFC 4364)
	Rd *string `json:"rd,omitempty" mapstructure:"rd,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m VRFRequest) Validate() error {
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
			m.Name, validation.Required, validation.Length(1, 100),
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
	}.Filter()
}

// GetComments returns the Comments property
func (m VRFRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VRFRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m VRFRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VRFRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VRFRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VRFRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnforceUnique returns the EnforceUnique property
func (m VRFRequest) GetEnforceUnique() bool {
	return m.EnforceUnique
}

// SetEnforceUnique sets the EnforceUnique property
func (m *VRFRequest) SetEnforceUnique(val bool) {
	m.EnforceUnique = val
}

// GetExportTargets returns the ExportTargets property
func (m VRFRequest) GetExportTargets() []int32 {
	return m.ExportTargets
}

// SetExportTargets sets the ExportTargets property
func (m *VRFRequest) SetExportTargets(val []int32) {
	m.ExportTargets = val
}

// GetImportTargets returns the ImportTargets property
func (m VRFRequest) GetImportTargets() []int32 {
	return m.ImportTargets
}

// SetImportTargets sets the ImportTargets property
func (m *VRFRequest) SetImportTargets(val []int32) {
	m.ImportTargets = val
}

// GetName returns the Name property
func (m VRFRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VRFRequest) SetName(val string) {
	m.Name = val
}

// GetRd returns the Rd property
func (m VRFRequest) GetRd() *string {
	return m.Rd
}

// SetRd sets the Rd property
func (m *VRFRequest) SetRd(val *string) {
	m.Rd = val
}

// GetTags returns the Tags property
func (m VRFRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VRFRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VRFRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VRFRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}
