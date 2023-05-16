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

// PatchedWritableCircuitRequest is an object. Adds support for custom fields and tags.
type PatchedWritableCircuitRequest struct {
	// Cid: Unique circuit ID
	Cid string `json:"cid,omitempty" mapstructure:"cid,omitempty"`
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
	// Provider:
	Provider int32 `json:"provider,omitempty" mapstructure:"provider,omitempty"`
	// ProviderAccount:
	ProviderAccount *int32 `json:"provider_account,omitempty" mapstructure:"provider_account,omitempty"`
	// Status: * `planned` - Planned
	// * `provisioning` - Provisioning
	// * `active` - Active
	// * `offline` - Offline
	// * `deprovisioning` - Deprovisioning
	// * `decommissioned` - Decommissioned
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TerminationDate:
	TerminationDate *string `json:"termination_date,omitempty" mapstructure:"termination_date,omitempty"`
	// Type:
	Type int32 `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableCircuitRequest) Validate() error {
	return validation.Errors{
		"cid": validation.Validate(
			m.Cid, validation.Length(1, 100),
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
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCid returns the Cid property
func (m PatchedWritableCircuitRequest) GetCid() string {
	return m.Cid
}

// SetCid sets the Cid property
func (m *PatchedWritableCircuitRequest) SetCid(val string) {
	m.Cid = val
}

// GetComments returns the Comments property
func (m PatchedWritableCircuitRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableCircuitRequest) SetComments(val string) {
	m.Comments = val
}

// GetCommitRate returns the CommitRate property
func (m PatchedWritableCircuitRequest) GetCommitRate() *int32 {
	return m.CommitRate
}

// SetCommitRate sets the CommitRate property
func (m *PatchedWritableCircuitRequest) SetCommitRate(val *int32) {
	m.CommitRate = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableCircuitRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableCircuitRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableCircuitRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableCircuitRequest) SetDescription(val string) {
	m.Description = val
}

// GetInstallDate returns the InstallDate property
func (m PatchedWritableCircuitRequest) GetInstallDate() *string {
	return m.InstallDate
}

// SetInstallDate sets the InstallDate property
func (m *PatchedWritableCircuitRequest) SetInstallDate(val *string) {
	m.InstallDate = val
}

// GetProvider returns the Provider property
func (m PatchedWritableCircuitRequest) GetProvider() int32 {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *PatchedWritableCircuitRequest) SetProvider(val int32) {
	m.Provider = val
}

// GetProviderAccount returns the ProviderAccount property
func (m PatchedWritableCircuitRequest) GetProviderAccount() *int32 {
	return m.ProviderAccount
}

// SetProviderAccount sets the ProviderAccount property
func (m *PatchedWritableCircuitRequest) SetProviderAccount(val *int32) {
	m.ProviderAccount = val
}

// GetStatus returns the Status property
func (m PatchedWritableCircuitRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableCircuitRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableCircuitRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableCircuitRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableCircuitRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableCircuitRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetTerminationDate returns the TerminationDate property
func (m PatchedWritableCircuitRequest) GetTerminationDate() *string {
	return m.TerminationDate
}

// SetTerminationDate sets the TerminationDate property
func (m *PatchedWritableCircuitRequest) SetTerminationDate(val *string) {
	m.TerminationDate = val
}

// GetType returns the Type property
func (m PatchedWritableCircuitRequest) GetType() int32 {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritableCircuitRequest) SetType(val int32) {
	m.Type = val
}
