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

// Aggregate is an object. Adds support for custom fields and tags.
type Aggregate struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DateAdded:
	DateAdded *string `json:"date_added,omitempty" mapstructure:"date_added,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Family:
	Family struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"family" mapstructure:"family"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Prefix:
	Prefix string `json:"prefix" mapstructure:"prefix"`
	// Rir: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rir NestedRIR `json:"rir" mapstructure:"rir"`
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
func (m Aggregate) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"rir": validation.Validate(
			m.Rir, validation.NotNil,
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
func (m Aggregate) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Aggregate) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Aggregate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Aggregate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Aggregate) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Aggregate) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDateAdded returns the DateAdded property
func (m Aggregate) GetDateAdded() *string {
	return m.DateAdded
}

// SetDateAdded sets the DateAdded property
func (m *Aggregate) SetDateAdded(val *string) {
	m.DateAdded = val
}

// GetDescription returns the Description property
func (m Aggregate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Aggregate) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Aggregate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Aggregate) SetDisplay(val string) {
	m.Display = val
}

// GetFamily returns the Family property
func (m Aggregate) GetFamily() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Family
}

// SetFamily sets the Family property
func (m *Aggregate) SetFamily(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Family = val
}

// GetId returns the Id property
func (m Aggregate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Aggregate) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Aggregate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Aggregate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetPrefix returns the Prefix property
func (m Aggregate) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *Aggregate) SetPrefix(val string) {
	m.Prefix = val
}

// GetRir returns the Rir property
func (m Aggregate) GetRir() NestedRIR {
	return m.Rir
}

// SetRir sets the Rir property
func (m *Aggregate) SetRir(val NestedRIR) {
	m.Rir = val
}

// GetTags returns the Tags property
func (m Aggregate) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Aggregate) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Aggregate) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Aggregate) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m Aggregate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Aggregate) SetUrl(val string) {
	m.Url = val
}
