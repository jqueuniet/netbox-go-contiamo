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

	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// l2VPNSlugPattern is the validation pattern for L2VPN.Slug
var l2VPNSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// L2VPN is an object. Adds support for custom fields and tags.
type L2VPN struct {
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
	// ExportTargets:
	ExportTargets []int32 `json:"export_targets,omitempty" mapstructure:"export_targets,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Identifier:
	Identifier *int64 `json:"identifier,omitempty" mapstructure:"identifier,omitempty"`
	// ImportTargets:
	ImportTargets []int32 `json:"import_targets,omitempty" mapstructure:"import_targets,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type:
	Type *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type,omitempty" mapstructure:"type,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m L2VPN) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(l2VPNSlugPattern),
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
func (m L2VPN) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *L2VPN) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m L2VPN) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *L2VPN) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m L2VPN) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *L2VPN) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m L2VPN) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *L2VPN) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m L2VPN) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *L2VPN) SetDisplay(val string) {
	m.Display = val
}

// GetExportTargets returns the ExportTargets property
func (m L2VPN) GetExportTargets() []int32 {
	return m.ExportTargets
}

// SetExportTargets sets the ExportTargets property
func (m *L2VPN) SetExportTargets(val []int32) {
	m.ExportTargets = val
}

// GetId returns the Id property
func (m L2VPN) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *L2VPN) SetId(val int32) {
	m.Id = val
}

// GetIdentifier returns the Identifier property
func (m L2VPN) GetIdentifier() *int64 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *L2VPN) SetIdentifier(val *int64) {
	m.Identifier = val
}

// GetImportTargets returns the ImportTargets property
func (m L2VPN) GetImportTargets() []int32 {
	return m.ImportTargets
}

// SetImportTargets sets the ImportTargets property
func (m *L2VPN) SetImportTargets(val []int32) {
	m.ImportTargets = val
}

// GetLastUpdated returns the LastUpdated property
func (m L2VPN) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *L2VPN) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m L2VPN) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *L2VPN) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m L2VPN) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *L2VPN) SetSlug(val string) {
	m.Slug = val
}

// GetTags returns the Tags property
func (m L2VPN) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *L2VPN) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m L2VPN) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *L2VPN) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetType returns the Type property
func (m L2VPN) GetType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *L2VPN) SetType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m L2VPN) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *L2VPN) SetUrl(val string) {
	m.Url = val
}
