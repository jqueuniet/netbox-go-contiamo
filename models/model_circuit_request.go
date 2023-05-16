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

// CircuitRequest is an object. Adds support for custom fields and tags.
type CircuitRequest struct {
	// Cid: Unique circuit ID
	Cid string `json:"cid" mapstructure:"cid"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CommitRate: Committed rate
	CommitRate *int32 `json:"commit_rate,omitempty" mapstructure:"commit_rate,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// InstallDate:
	InstallDate *string `json:"install_date,omitempty" mapstructure:"install_date,omitempty"`
	// Provider: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Provider NestedProviderRequest `json:"provider" mapstructure:"provider"`
	// ProviderAccount: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ProviderAccount *NestedProviderAccountRequest `json:"provider_account,omitempty" mapstructure:"provider_account,omitempty"`
	// Status: * `planned` - Planned
	// * `provisioning` - Provisioning
	// * `active` - Active
	// * `offline` - Offline
	// * `deprovisioning` - Deprovisioning
	// * `decommissioned` - Decommissioned
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TerminationDate:
	TerminationDate *string `json:"termination_date,omitempty" mapstructure:"termination_date,omitempty"`
	// Type: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Type NestedCircuitTypeRequest `json:"type" mapstructure:"type"`
}

// Validate implements basic validation for this model
func (m CircuitRequest) Validate() error {
	return validation.Errors{
		"cid": validation.Validate(
			m.Cid, validation.Required, validation.Length(1, 100),
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
		"type": validation.Validate(
			m.Type, validation.NotNil,
		),
	}.Filter()
}

// GetCid returns the Cid property
func (m CircuitRequest) GetCid() string {
	return m.Cid
}

// SetCid sets the Cid property
func (m *CircuitRequest) SetCid(val string) {
	m.Cid = val
}

// GetComments returns the Comments property
func (m CircuitRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *CircuitRequest) SetComments(val string) {
	m.Comments = val
}

// GetCommitRate returns the CommitRate property
func (m CircuitRequest) GetCommitRate() *int32 {
	return m.CommitRate
}

// SetCommitRate sets the CommitRate property
func (m *CircuitRequest) SetCommitRate(val *int32) {
	m.CommitRate = val
}

// GetCustomFields returns the CustomFields property
func (m CircuitRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *CircuitRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m CircuitRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitRequest) SetDescription(val string) {
	m.Description = val
}

// GetInstallDate returns the InstallDate property
func (m CircuitRequest) GetInstallDate() *string {
	return m.InstallDate
}

// SetInstallDate sets the InstallDate property
func (m *CircuitRequest) SetInstallDate(val *string) {
	m.InstallDate = val
}

// GetProvider returns the Provider property
func (m CircuitRequest) GetProvider() NestedProviderRequest {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *CircuitRequest) SetProvider(val NestedProviderRequest) {
	m.Provider = val
}

// GetProviderAccount returns the ProviderAccount property
func (m CircuitRequest) GetProviderAccount() *NestedProviderAccountRequest {
	return m.ProviderAccount
}

// SetProviderAccount sets the ProviderAccount property
func (m *CircuitRequest) SetProviderAccount(val *NestedProviderAccountRequest) {
	m.ProviderAccount = val
}

// GetStatus returns the Status property
func (m CircuitRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *CircuitRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m CircuitRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *CircuitRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m CircuitRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *CircuitRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetTerminationDate returns the TerminationDate property
func (m CircuitRequest) GetTerminationDate() *string {
	return m.TerminationDate
}

// SetTerminationDate sets the TerminationDate property
func (m *CircuitRequest) SetTerminationDate(val *string) {
	m.TerminationDate = val
}

// GetType returns the Type property
func (m CircuitRequest) GetType() NestedCircuitTypeRequest {
	return m.Type
}

// SetType sets the Type property
func (m *CircuitRequest) SetType(val NestedCircuitTypeRequest) {
	m.Type = val
}
