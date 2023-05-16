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

// PatchedWritableVRFRequest is an object. Adds support for custom fields and tags.
type PatchedWritableVRFRequest struct {
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
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Rd: Unique route distinguisher (as defined in RFC 4364)
	Rd *string `json:"rd,omitempty" mapstructure:"rd,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableVRFRequest) Validate() error {
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
			m.Name, validation.Length(1, 100),
		),
		"rd": validation.Validate(
			m.Rd, validation.Length(0, 21),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableVRFRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableVRFRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableVRFRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableVRFRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableVRFRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableVRFRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnforceUnique returns the EnforceUnique property
func (m PatchedWritableVRFRequest) GetEnforceUnique() bool {
	return m.EnforceUnique
}

// SetEnforceUnique sets the EnforceUnique property
func (m *PatchedWritableVRFRequest) SetEnforceUnique(val bool) {
	m.EnforceUnique = val
}

// GetExportTargets returns the ExportTargets property
func (m PatchedWritableVRFRequest) GetExportTargets() []int32 {
	return m.ExportTargets
}

// SetExportTargets sets the ExportTargets property
func (m *PatchedWritableVRFRequest) SetExportTargets(val []int32) {
	m.ExportTargets = val
}

// GetImportTargets returns the ImportTargets property
func (m PatchedWritableVRFRequest) GetImportTargets() []int32 {
	return m.ImportTargets
}

// SetImportTargets sets the ImportTargets property
func (m *PatchedWritableVRFRequest) SetImportTargets(val []int32) {
	m.ImportTargets = val
}

// GetName returns the Name property
func (m PatchedWritableVRFRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableVRFRequest) SetName(val string) {
	m.Name = val
}

// GetRd returns the Rd property
func (m PatchedWritableVRFRequest) GetRd() *string {
	return m.Rd
}

// SetRd sets the Rd property
func (m *PatchedWritableVRFRequest) SetRd(val *string) {
	m.Rd = val
}

// GetTags returns the Tags property
func (m PatchedWritableVRFRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableVRFRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableVRFRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableVRFRequest) SetTenant(val *int32) {
	m.Tenant = val
}
