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

// VirtualMachineWithConfigContext is an object. Adds support for custom fields and tags.
type VirtualMachineWithConfigContext struct {
	// Cluster: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Cluster *NestedCluster `json:"cluster,omitempty" mapstructure:"cluster,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// ConfigContext:
	ConfigContext map[string]interface{} `json:"config_context" mapstructure:"config_context"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device *NestedDevice `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Disk:
	Disk *int32 `json:"disk,omitempty" mapstructure:"disk,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// LocalContextData: Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// Memory:
	Memory *int32 `json:"memory,omitempty" mapstructure:"memory,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Platform: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Platform *NestedPlatform `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// PrimaryIp: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp NestedIPAddress `json:"primary_ip" mapstructure:"primary_ip"`
	// PrimaryIp4: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedDeviceRole `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSite `json:"site,omitempty" mapstructure:"site,omitempty"`
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
	// Vcpus:
	Vcpus *float64 `json:"vcpus,omitempty" mapstructure:"vcpus,omitempty"`
}

// Validate implements basic validation for this model
func (m VirtualMachineWithConfigContext) Validate() error {
	return validation.Errors{
		"cluster": validation.Validate(
			m.Cluster,
		),
		"configContext": validation.Validate(
			m.ConfigContext,
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
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"platform": validation.Validate(
			m.Platform,
		),
		"primaryIp": validation.Validate(
			m.PrimaryIp, validation.NotNil,
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"vcpus": validation.Validate(
			m.Vcpus, validation.Min(*float64(0.01)), validation.Max(*float64(10000)),
		),
	}.Filter()
}

// GetCluster returns the Cluster property
func (m VirtualMachineWithConfigContext) GetCluster() *NestedCluster {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *VirtualMachineWithConfigContext) SetCluster(val *NestedCluster) {
	m.Cluster = val
}

// GetComments returns the Comments property
func (m VirtualMachineWithConfigContext) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *VirtualMachineWithConfigContext) SetComments(val string) {
	m.Comments = val
}

// GetConfigContext returns the ConfigContext property
func (m VirtualMachineWithConfigContext) GetConfigContext() map[string]interface{} {
	return m.ConfigContext
}

// SetConfigContext sets the ConfigContext property
func (m *VirtualMachineWithConfigContext) SetConfigContext(val map[string]interface{}) {
	m.ConfigContext = val
}

// GetCreated returns the Created property
func (m VirtualMachineWithConfigContext) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VirtualMachineWithConfigContext) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m VirtualMachineWithConfigContext) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VirtualMachineWithConfigContext) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VirtualMachineWithConfigContext) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VirtualMachineWithConfigContext) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m VirtualMachineWithConfigContext) GetDevice() *NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *VirtualMachineWithConfigContext) SetDevice(val *NestedDevice) {
	m.Device = val
}

// GetDisk returns the Disk property
func (m VirtualMachineWithConfigContext) GetDisk() *int32 {
	return m.Disk
}

// SetDisk sets the Disk property
func (m *VirtualMachineWithConfigContext) SetDisk(val *int32) {
	m.Disk = val
}

// GetDisplay returns the Display property
func (m VirtualMachineWithConfigContext) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *VirtualMachineWithConfigContext) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m VirtualMachineWithConfigContext) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VirtualMachineWithConfigContext) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m VirtualMachineWithConfigContext) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VirtualMachineWithConfigContext) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLocalContextData returns the LocalContextData property
func (m VirtualMachineWithConfigContext) GetLocalContextData() map[string]interface{} {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *VirtualMachineWithConfigContext) SetLocalContextData(val map[string]interface{}) {
	m.LocalContextData = val
}

// GetMemory returns the Memory property
func (m VirtualMachineWithConfigContext) GetMemory() *int32 {
	return m.Memory
}

// SetMemory sets the Memory property
func (m *VirtualMachineWithConfigContext) SetMemory(val *int32) {
	m.Memory = val
}

// GetName returns the Name property
func (m VirtualMachineWithConfigContext) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualMachineWithConfigContext) SetName(val string) {
	m.Name = val
}

// GetPlatform returns the Platform property
func (m VirtualMachineWithConfigContext) GetPlatform() *NestedPlatform {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *VirtualMachineWithConfigContext) SetPlatform(val *NestedPlatform) {
	m.Platform = val
}

// GetPrimaryIp returns the PrimaryIp property
func (m VirtualMachineWithConfigContext) GetPrimaryIp() NestedIPAddress {
	return m.PrimaryIp
}

// SetPrimaryIp sets the PrimaryIp property
func (m *VirtualMachineWithConfigContext) SetPrimaryIp(val NestedIPAddress) {
	m.PrimaryIp = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m VirtualMachineWithConfigContext) GetPrimaryIp4() *NestedIPAddress {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *VirtualMachineWithConfigContext) SetPrimaryIp4(val *NestedIPAddress) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m VirtualMachineWithConfigContext) GetPrimaryIp6() *NestedIPAddress {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *VirtualMachineWithConfigContext) SetPrimaryIp6(val *NestedIPAddress) {
	m.PrimaryIp6 = val
}

// GetRole returns the Role property
func (m VirtualMachineWithConfigContext) GetRole() *NestedDeviceRole {
	return m.Role
}

// SetRole sets the Role property
func (m *VirtualMachineWithConfigContext) SetRole(val *NestedDeviceRole) {
	m.Role = val
}

// GetSite returns the Site property
func (m VirtualMachineWithConfigContext) GetSite() *NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *VirtualMachineWithConfigContext) SetSite(val *NestedSite) {
	m.Site = val
}

// GetStatus returns the Status property
func (m VirtualMachineWithConfigContext) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *VirtualMachineWithConfigContext) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m VirtualMachineWithConfigContext) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VirtualMachineWithConfigContext) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m VirtualMachineWithConfigContext) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VirtualMachineWithConfigContext) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m VirtualMachineWithConfigContext) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *VirtualMachineWithConfigContext) SetUrl(val string) {
	m.Url = val
}

// GetVcpus returns the Vcpus property
func (m VirtualMachineWithConfigContext) GetVcpus() *float64 {
	return m.Vcpus
}

// SetVcpus sets the Vcpus property
func (m *VirtualMachineWithConfigContext) SetVcpus(val *float64) {
	m.Vcpus = val
}
