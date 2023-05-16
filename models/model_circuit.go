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

// Circuit is an object. Adds support for custom fields and tags.
type Circuit struct {
	// Cid: Unique circuit ID
	Cid string `json:"cid" mapstructure:"cid"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CommitRate: Committed rate
	CommitRate *int32 `json:"commit_rate,omitempty" mapstructure:"commit_rate,omitempty"`
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
	// InstallDate:
	InstallDate *string `json:"install_date,omitempty" mapstructure:"install_date,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Provider: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Provider NestedProvider `json:"provider" mapstructure:"provider"`
	// ProviderAccount: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ProviderAccount *NestedProviderAccount `json:"provider_account,omitempty" mapstructure:"provider_account,omitempty"`
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
	// TerminationA: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	TerminationA *CircuitCircuitTermination `json:"termination_a" mapstructure:"termination_a"`
	// TerminationDate:
	TerminationDate *string `json:"termination_date,omitempty" mapstructure:"termination_date,omitempty"`
	// TerminationZ: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	TerminationZ *CircuitCircuitTermination `json:"termination_z" mapstructure:"termination_z"`
	// Type: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Type NestedCircuitType `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Circuit) Validate() error {
	return validation.Errors{
		"cid": validation.Validate(
			m.Cid, validation.NotNil, validation.Length(0, 100),
		),
		"commitRate": validation.Validate(
			m.CommitRate, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"provider": validation.Validate(
			m.Provider, validation.NotNil,
		),
		"providerAccount": validation.Validate(
			m.ProviderAccount,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"terminationA": validation.Validate(
			m.TerminationA,
		),
		"terminationZ": validation.Validate(
			m.TerminationZ,
		),
		"type": validation.Validate(
			m.Type, validation.NotNil,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCid returns the Cid property
func (m Circuit) GetCid() string {
	return m.Cid
}

// SetCid sets the Cid property
func (m *Circuit) SetCid(val string) {
	m.Cid = val
}

// GetComments returns the Comments property
func (m Circuit) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Circuit) SetComments(val string) {
	m.Comments = val
}

// GetCommitRate returns the CommitRate property
func (m Circuit) GetCommitRate() *int32 {
	return m.CommitRate
}

// SetCommitRate sets the CommitRate property
func (m *Circuit) SetCommitRate(val *int32) {
	m.CommitRate = val
}

// GetCreated returns the Created property
func (m Circuit) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Circuit) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Circuit) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Circuit) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Circuit) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Circuit) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Circuit) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Circuit) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Circuit) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Circuit) SetId(val int32) {
	m.Id = val
}

// GetInstallDate returns the InstallDate property
func (m Circuit) GetInstallDate() *string {
	return m.InstallDate
}

// SetInstallDate sets the InstallDate property
func (m *Circuit) SetInstallDate(val *string) {
	m.InstallDate = val
}

// GetLastUpdated returns the LastUpdated property
func (m Circuit) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Circuit) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetProvider returns the Provider property
func (m Circuit) GetProvider() NestedProvider {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *Circuit) SetProvider(val NestedProvider) {
	m.Provider = val
}

// GetProviderAccount returns the ProviderAccount property
func (m Circuit) GetProviderAccount() *NestedProviderAccount {
	return m.ProviderAccount
}

// SetProviderAccount sets the ProviderAccount property
func (m *Circuit) SetProviderAccount(val *NestedProviderAccount) {
	m.ProviderAccount = val
}

// GetStatus returns the Status property
func (m Circuit) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Circuit) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Circuit) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Circuit) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Circuit) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Circuit) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetTerminationA returns the TerminationA property
func (m Circuit) GetTerminationA() *CircuitCircuitTermination {
	return m.TerminationA
}

// SetTerminationA sets the TerminationA property
func (m *Circuit) SetTerminationA(val *CircuitCircuitTermination) {
	m.TerminationA = val
}

// GetTerminationDate returns the TerminationDate property
func (m Circuit) GetTerminationDate() *string {
	return m.TerminationDate
}

// SetTerminationDate sets the TerminationDate property
func (m *Circuit) SetTerminationDate(val *string) {
	m.TerminationDate = val
}

// GetTerminationZ returns the TerminationZ property
func (m Circuit) GetTerminationZ() *CircuitCircuitTermination {
	return m.TerminationZ
}

// SetTerminationZ sets the TerminationZ property
func (m *Circuit) SetTerminationZ(val *CircuitCircuitTermination) {
	m.TerminationZ = val
}

// GetType returns the Type property
func (m Circuit) GetType() NestedCircuitType {
	return m.Type
}

// SetType sets the Type property
func (m *Circuit) SetType(val NestedCircuitType) {
	m.Type = val
}

// GetUrl returns the Url property
func (m Circuit) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Circuit) SetUrl(val string) {
	m.Url = val
}
