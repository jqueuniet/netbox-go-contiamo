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

// IPRange is an object. Adds support for custom fields and tags.
type IPRange struct {
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
	// EndAddress:
	EndAddress string `json:"end_address" mapstructure:"end_address"`
	// Family:
	Family struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"family" mapstructure:"family"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// MarkUtilized: Treat as 100% utilized
	MarkUtilized bool `json:"mark_utilized,omitempty" mapstructure:"mark_utilized,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedRole `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Size:
	Size int32 `json:"size" mapstructure:"size"`
	// StartAddress:
	StartAddress string `json:"start_address" mapstructure:"start_address"`
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
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRF `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m IPRange) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"role": validation.Validate(
			m.Role,
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
		"vrf": validation.Validate(
			m.Vrf,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m IPRange) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *IPRange) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m IPRange) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IPRange) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m IPRange) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *IPRange) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m IPRange) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IPRange) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m IPRange) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *IPRange) SetDisplay(val string) {
	m.Display = val
}

// GetEndAddress returns the EndAddress property
func (m IPRange) GetEndAddress() string {
	return m.EndAddress
}

// SetEndAddress sets the EndAddress property
func (m *IPRange) SetEndAddress(val string) {
	m.EndAddress = val
}

// GetFamily returns the Family property
func (m IPRange) GetFamily() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Family
}

// SetFamily sets the Family property
func (m *IPRange) SetFamily(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Family = val
}

// GetId returns the Id property
func (m IPRange) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IPRange) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m IPRange) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IPRange) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m IPRange) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *IPRange) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetRole returns the Role property
func (m IPRange) GetRole() *NestedRole {
	return m.Role
}

// SetRole sets the Role property
func (m *IPRange) SetRole(val *NestedRole) {
	m.Role = val
}

// GetSize returns the Size property
func (m IPRange) GetSize() int32 {
	return m.Size
}

// SetSize sets the Size property
func (m *IPRange) SetSize(val int32) {
	m.Size = val
}

// GetStartAddress returns the StartAddress property
func (m IPRange) GetStartAddress() string {
	return m.StartAddress
}

// SetStartAddress sets the StartAddress property
func (m *IPRange) SetStartAddress(val string) {
	m.StartAddress = val
}

// GetStatus returns the Status property
func (m IPRange) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *IPRange) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m IPRange) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *IPRange) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m IPRange) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IPRange) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m IPRange) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *IPRange) SetUrl(val string) {
	m.Url = val
}

// GetVrf returns the Vrf property
func (m IPRange) GetVrf() *NestedVRF {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *IPRange) SetVrf(val *NestedVRF) {
	m.Vrf = val
}
