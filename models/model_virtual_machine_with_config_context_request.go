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

// VirtualMachineWithConfigContextRequest is an object. Adds support for custom fields and tags.
type VirtualMachineWithConfigContextRequest struct {
	// Cluster: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Cluster *NestedClusterRequest `json:"cluster,omitempty" mapstructure:"cluster,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device *NestedDeviceRequest `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Disk:
	Disk *int32 `json:"disk,omitempty" mapstructure:"disk,omitempty"`
	// LocalContextData: Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// Memory:
	Memory *int32 `json:"memory,omitempty" mapstructure:"memory,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Platform: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Platform *NestedPlatformRequest `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// PrimaryIp4: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp4 *NestedIPAddressRequest `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp6 *NestedIPAddressRequest `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedDeviceRoleRequest `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSiteRequest `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
	// * `decommissioning` - Decommissioning
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Vcpus:
	Vcpus *float64 `json:"vcpus,omitempty" mapstructure:"vcpus,omitempty"`
}

// Validate implements basic validation for this model
func (m VirtualMachineWithConfigContextRequest) Validate() error {
	return validation.Errors{
		"cluster": validation.Validate(
			m.Cluster,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"device": validation.Validate(
			m.Device,
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
			m.Name, validation.Required, validation.Length(1, 64),
		),
		"platform": validation.Validate(
			m.Platform,
		),
		"primaryIp4": validation.Validate(
			m.PrimaryIp4,
		),
		"primaryIp6": validation.Validate(
			m.PrimaryIp6,
		),
		"role": validation.Validate(
			m.Role,
		),
		"site": validation.Validate(
			m.Site,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"vcpus": validation.Validate(
			m.Vcpus, validation.Min(*float64(0.01)), validation.Max(*float64(10000)),
		),
	}.Filter()
}

// GetCluster returns the Cluster property
func (m VirtualMachineWithConfigContextRequest) GetCluster() *NestedClusterRequest {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *VirtualMachineWithConfigContextRequest) SetCluster(val *NestedClusterRequest) {
	m.Cluster = val
}

// GetComments returns the Comments property
func (m VirtualMachineWithConfigContextRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VirtualMachineWithConfigContextRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m VirtualMachineWithConfigContextRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VirtualMachineWithConfigContextRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VirtualMachineWithConfigContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VirtualMachineWithConfigContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m VirtualMachineWithConfigContextRequest) GetDevice() *NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *VirtualMachineWithConfigContextRequest) SetDevice(val *NestedDeviceRequest) {
	m.Device = val
}

// GetDisk returns the Disk property
func (m VirtualMachineWithConfigContextRequest) GetDisk() *int32 {
	return m.Disk
}

// SetDisk sets the Disk property
func (m *VirtualMachineWithConfigContextRequest) SetDisk(val *int32) {
	m.Disk = val
}

// GetLocalContextData returns the LocalContextData property
func (m VirtualMachineWithConfigContextRequest) GetLocalContextData() map[string]interface{} {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *VirtualMachineWithConfigContextRequest) SetLocalContextData(val map[string]interface{}) {
	m.LocalContextData = val
}

// GetMemory returns the Memory property
func (m VirtualMachineWithConfigContextRequest) GetMemory() *int32 {
	return m.Memory
}

// SetMemory sets the Memory property
func (m *VirtualMachineWithConfigContextRequest) SetMemory(val *int32) {
	m.Memory = val
}

// GetName returns the Name property
func (m VirtualMachineWithConfigContextRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualMachineWithConfigContextRequest) SetName(val string) {
	m.Name = val
}

// GetPlatform returns the Platform property
func (m VirtualMachineWithConfigContextRequest) GetPlatform() *NestedPlatformRequest {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *VirtualMachineWithConfigContextRequest) SetPlatform(val *NestedPlatformRequest) {
	m.Platform = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m VirtualMachineWithConfigContextRequest) GetPrimaryIp4() *NestedIPAddressRequest {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *VirtualMachineWithConfigContextRequest) SetPrimaryIp4(val *NestedIPAddressRequest) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m VirtualMachineWithConfigContextRequest) GetPrimaryIp6() *NestedIPAddressRequest {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *VirtualMachineWithConfigContextRequest) SetPrimaryIp6(val *NestedIPAddressRequest) {
	m.PrimaryIp6 = val
}

// GetRole returns the Role property
func (m VirtualMachineWithConfigContextRequest) GetRole() *NestedDeviceRoleRequest {
	return m.Role
}

// SetRole sets the Role property
func (m *VirtualMachineWithConfigContextRequest) SetRole(val *NestedDeviceRoleRequest) {
	m.Role = val
}

// GetSite returns the Site property
func (m VirtualMachineWithConfigContextRequest) GetSite() *NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *VirtualMachineWithConfigContextRequest) SetSite(val *NestedSiteRequest) {
	m.Site = val
}

// GetStatus returns the Status property
func (m VirtualMachineWithConfigContextRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *VirtualMachineWithConfigContextRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m VirtualMachineWithConfigContextRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VirtualMachineWithConfigContextRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VirtualMachineWithConfigContextRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VirtualMachineWithConfigContextRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetVcpus returns the Vcpus property
func (m VirtualMachineWithConfigContextRequest) GetVcpus() *float64 {
	return m.Vcpus
}

// SetVcpus sets the Vcpus property
func (m *VirtualMachineWithConfigContextRequest) SetVcpus(val *float64) {
	m.Vcpus = val
}
