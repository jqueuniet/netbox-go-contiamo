// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// patchedWritableL2VPNRequestSlugPattern is the validation pattern for PatchedWritableL2VPNRequest.Slug
var patchedWritableL2VPNRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// PatchedWritableL2VPNRequest is an object. Adds support for custom fields and tags.
type PatchedWritableL2VPNRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// ExportTargets:
	ExportTargets []int32 `json:"export_targets,omitempty" mapstructure:"export_targets,omitempty"`
	// Identifier:
	Identifier *int64 `json:"identifier,omitempty" mapstructure:"identifier,omitempty"`
	// ImportTargets:
	ImportTargets []int32 `json:"import_targets,omitempty" mapstructure:"import_targets,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Slug:
	Slug string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type: * `vpws` - VPWS
	// * `vpls` - VPLS
	// * `vxlan` - VXLAN
	// * `vxlan-evpn` - VXLAN-EVPN
	// * `mpls-evpn` - MPLS EVPN
	// * `pbb-evpn` - PBB EVPN
	// * `epl` - EPL
	// * `evpl` - EVPL
	// * `ep-lan` - Ethernet Private LAN
	// * `evp-lan` - Ethernet Virtual Private LAN
	// * `ep-tree` - Ethernet Private Tree
	// * `evp-tree` - Ethernet Virtual Private Tree
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableL2VPNRequest) Validate() error {
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
		"identifier": validation.Validate(
			m.Identifier, validation.Min(*int64(-9.223372036854776e+18)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"importTargets": validation.Validate(
			m.ImportTargets,
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Length(1, 100), validation.Match(patchedWritableL2VPNRequestSlugPattern),
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableL2VPNRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableL2VPNRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableL2VPNRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableL2VPNRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableL2VPNRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableL2VPNRequest) SetDescription(val string) {
	m.Description = val
}

// GetExportTargets returns the ExportTargets property
func (m PatchedWritableL2VPNRequest) GetExportTargets() []int32 {
	return m.ExportTargets
}

// SetExportTargets sets the ExportTargets property
func (m *PatchedWritableL2VPNRequest) SetExportTargets(val []int32) {
	m.ExportTargets = val
}

// GetIdentifier returns the Identifier property
func (m PatchedWritableL2VPNRequest) GetIdentifier() *int64 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *PatchedWritableL2VPNRequest) SetIdentifier(val *int64) {
	m.Identifier = val
}

// GetImportTargets returns the ImportTargets property
func (m PatchedWritableL2VPNRequest) GetImportTargets() []int32 {
	return m.ImportTargets
}

// SetImportTargets sets the ImportTargets property
func (m *PatchedWritableL2VPNRequest) SetImportTargets(val []int32) {
	m.ImportTargets = val
}

// GetName returns the Name property
func (m PatchedWritableL2VPNRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableL2VPNRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m PatchedWritableL2VPNRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *PatchedWritableL2VPNRequest) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m PatchedWritableL2VPNRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableL2VPNRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableL2VPNRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableL2VPNRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetType returns the Type property
func (m PatchedWritableL2VPNRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritableL2VPNRequest) SetType(val string) {
	m.Type = val
}
