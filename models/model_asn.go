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

// ASN is an object. Adds support for custom fields and tags.
type ASN struct {
	// Asn: 16- or 32-bit autonomous system number
	Asn int64 `json:"asn" mapstructure:"asn"`
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
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// ProviderCount:
	ProviderCount int32 `json:"provider_count" mapstructure:"provider_count"`
	// Rir: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rir *NestedRIR `json:"rir,omitempty" mapstructure:"rir,omitempty"`
	// SiteCount:
	SiteCount int32 `json:"site_count" mapstructure:"site_count"`
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
func (m ASN) Validate() error {
	return validation.Errors{
		"asn": validation.Validate(
			m.Asn, validation.Required, validation.Min(int64(1)), validation.Max(int64(4.294967295e+09)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"rir": validation.Validate(
			m.Rir,
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

// GetAsn returns the Asn property
func (m ASN) GetAsn() int64 {
	return m.Asn
}

// SetAsn sets the Asn property
func (m *ASN) SetAsn(val int64) {
	m.Asn = val
}

// GetComments returns the Comments property
func (m ASN) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ASN) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m ASN) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ASN) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ASN) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ASN) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ASN) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ASN) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ASN) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ASN) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ASN) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ASN) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m ASN) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ASN) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetProviderCount returns the ProviderCount property
func (m ASN) GetProviderCount() int32 {
	return m.ProviderCount
}

// SetProviderCount sets the ProviderCount property
func (m *ASN) SetProviderCount(val int32) {
	m.ProviderCount = val
}

// GetRir returns the Rir property
func (m ASN) GetRir() *NestedRIR {
	return m.Rir
}

// SetRir sets the Rir property
func (m *ASN) SetRir(val *NestedRIR) {
	m.Rir = val
}

// GetSiteCount returns the SiteCount property
func (m ASN) GetSiteCount() int32 {
	return m.SiteCount
}

// SetSiteCount sets the SiteCount property
func (m *ASN) SetSiteCount(val int32) {
	m.SiteCount = val
}

// GetTags returns the Tags property
func (m ASN) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ASN) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m ASN) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *ASN) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m ASN) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ASN) SetUrl(val string) {
	m.Url = val
}
