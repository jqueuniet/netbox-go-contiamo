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

// aSNRangeSlugPattern is the validation pattern for ASNRange.Slug
var aSNRangeSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// ASNRange is an object. Adds support for custom fields and tags.
type ASNRange struct {
	// AsnCount:
	AsnCount int32 `json:"asn_count" mapstructure:"asn_count"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// End:
	End int64 `json:"end" mapstructure:"end"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Rir: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rir NestedRIR `json:"rir" mapstructure:"rir"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Start:
	Start int64 `json:"start" mapstructure:"start"`
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
func (m ASNRange) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"end": validation.Validate(
			m.End, validation.Required, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"rir": validation.Validate(
			m.Rir, validation.NotNil,
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(aSNRangeSlugPattern),
		),
		"start": validation.Validate(
			m.Start, validation.Required, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
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

// GetAsnCount returns the AsnCount property
func (m ASNRange) GetAsnCount() int32 {
	return m.AsnCount
}

// SetAsnCount sets the AsnCount property
func (m *ASNRange) SetAsnCount(val int32) {
	m.AsnCount = val
}

// GetCreated returns the Created property
func (m ASNRange) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ASNRange) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ASNRange) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ASNRange) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ASNRange) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ASNRange) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ASNRange) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ASNRange) SetDisplay(val string) {
	m.Display = val
}

// GetEnd returns the End property
func (m ASNRange) GetEnd() int64 {
	return m.End
}

// SetEnd sets the End property
func (m *ASNRange) SetEnd(val int64) {
	m.End = val
}

// GetId returns the Id property
func (m ASNRange) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ASNRange) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ASNRange) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ASNRange) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ASNRange) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ASNRange) SetName(val string) {
	m.Name = val
}

// GetRir returns the Rir property
func (m ASNRange) GetRir() NestedRIR {
	return m.Rir
}

// SetRir sets the Rir property
func (m *ASNRange) SetRir(val NestedRIR) {
	m.Rir = val
}

// GetSlug returns the Slug property
func (m ASNRange) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ASNRange) SetSlug(val string) {
	m.Slug = val
}

// GetStart returns the Start property
func (m ASNRange) GetStart() int64 {
	return m.Start
}

// SetStart sets the Start property
func (m *ASNRange) SetStart(val int64) {
	m.Start = val
}

// GetTags returns the Tags property
func (m ASNRange) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ASNRange) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m ASNRange) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *ASNRange) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m ASNRange) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ASNRange) SetUrl(val string) {
	m.Url = val
}
