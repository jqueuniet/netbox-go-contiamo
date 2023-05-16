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

// PatchedWritableVirtualMachineWithConfigContextRequest is an object. Adds support for custom fields and tags.
type PatchedWritableVirtualMachineWithConfigContextRequest struct {
	// Cluster:
	Cluster *int32 `json:"cluster,omitempty" mapstructure:"cluster,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device *int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Disk:
	Disk *int32 `json:"disk,omitempty" mapstructure:"disk,omitempty"`
	// LocalContextData: Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// Memory:
	Memory *int32 `json:"memory,omitempty" mapstructure:"memory,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Platform:
	Platform *int32 `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// PrimaryIp4:
	PrimaryIp4 *int32 `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6:
	PrimaryIp6 *int32 `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Role:
	Role *int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site:
	Site *int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
	// * `decommissioning` - Decommissioning
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vcpus:
	Vcpus *float64 `json:"vcpus,omitempty" mapstructure:"vcpus,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableVirtualMachineWithConfigContextRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"disk": validation.Validate(
			m.Disk, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"localContextData": validation.Validate(
			m.LocalContextData,
		),
		"memory": validation.Validate(
			m.Memory, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"vcpus": validation.Validate(
			m.Vcpus, validation.Min(*float64(0.01)), validation.Max(*float64(10000)),
		),
	}.Filter()
}

// GetCluster returns the Cluster property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetCluster() *int32 {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetCluster(val *int32) {
	m.Cluster = val
}

// GetComments returns the Comments property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetDevice() *int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetDevice(val *int32) {
	m.Device = val
}

// GetDisk returns the Disk property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetDisk() *int32 {
	return m.Disk
}

// SetDisk sets the Disk property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetDisk(val *int32) {
	m.Disk = val
}

// GetLocalContextData returns the LocalContextData property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetLocalContextData() map[string]interface{} {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetLocalContextData(val map[string]interface{}) {
	m.LocalContextData = val
}

// GetMemory returns the Memory property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetMemory() *int32 {
	return m.Memory
}

// SetMemory sets the Memory property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetMemory(val *int32) {
	m.Memory = val
}

// GetName returns the Name property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetName(val string) {
	m.Name = val
}

// GetPlatform returns the Platform property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetPlatform() *int32 {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetPlatform(val *int32) {
	m.Platform = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetPrimaryIp4() *int32 {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetPrimaryIp4(val *int32) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetPrimaryIp6() *int32 {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetPrimaryIp6(val *int32) {
	m.PrimaryIp6 = val
}

// GetRole returns the Role property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetRole(val *int32) {
	m.Role = val
}

// GetSite returns the Site property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetSite() *int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetSite(val *int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVcpus returns the Vcpus property
func (m PatchedWritableVirtualMachineWithConfigContextRequest) GetVcpus() *float64 {
	return m.Vcpus
}

// SetVcpus sets the Vcpus property
func (m *PatchedWritableVirtualMachineWithConfigContextRequest) SetVcpus(val *float64) {
	m.Vcpus = val
}
