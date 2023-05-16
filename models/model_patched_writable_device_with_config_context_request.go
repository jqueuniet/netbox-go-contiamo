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

// PatchedWritableDeviceWithConfigContextRequest is an object. Adds support for custom fields and tags.
type PatchedWritableDeviceWithConfigContextRequest struct {
	// Airflow: * `front-to-rear` - Front to rear
	// * `rear-to-front` - Rear to front
	// * `left-to-right` - Left to right
	// * `right-to-left` - Right to left
	// * `side-to-rear` - Side to rear
	// * `passive` - Passive
	// * `mixed` - Mixed
	Airflow string `json:"airflow,omitempty" mapstructure:"airflow,omitempty"`
	// AssetTag: A unique tag used to identify this device
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// Cluster:
	Cluster *int32 `json:"cluster,omitempty" mapstructure:"cluster,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceRole: The function this device serves
	DeviceRole int32 `json:"device_role,omitempty" mapstructure:"device_role,omitempty"`
	// DeviceType:
	DeviceType int32 `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// Face: * `front` - Front
	// * `rear` - Rear
	Face string `json:"face,omitempty" mapstructure:"face,omitempty"`
	// LocalContextData: Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// Location:
	Location *int32 `json:"location,omitempty" mapstructure:"location,omitempty"`
	// Name:
	Name *string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Platform:
	Platform *int32 `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// Position:
	Position *float64 `json:"position,omitempty" mapstructure:"position,omitempty"`
	// PrimaryIp4:
	PrimaryIp4 *int32 `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6:
	PrimaryIp6 *int32 `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Rack:
	Rack *int32 `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// Serial: Chassis serial number, assigned by the manufacturer
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Site:
	Site int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
	// * `inventory` - Inventory
	// * `decommissioning` - Decommissioning
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// VcPosition:
	VcPosition *int32 `json:"vc_position,omitempty" mapstructure:"vc_position,omitempty"`
	// VcPriority: Virtual chassis master election priority
	VcPriority *int32 `json:"vc_priority,omitempty" mapstructure:"vc_priority,omitempty"`
	// VirtualChassis:
	VirtualChassis *int32 `json:"virtual_chassis,omitempty" mapstructure:"virtual_chassis,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableDeviceWithConfigContextRequest) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag, validation.Length(0, 50),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"localContextData": validation.Validate(
			m.LocalContextData,
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 64),
		),
		"position": validation.Validate(
			m.Position, validation.Min(*float64(0.5)), validation.Max(*float64(1000)),
		),
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"vcPosition": validation.Validate(
			m.VcPosition, validation.Min(*int32(0)), validation.Max(*int32(255)),
		),
		"vcPriority": validation.Validate(
			m.VcPriority, validation.Min(*int32(0)), validation.Max(*int32(255)),
		),
	}.Filter()
}

// GetAirflow returns the Airflow property
func (m PatchedWritableDeviceWithConfigContextRequest) GetAirflow() string {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetAirflow(val string) {
	m.Airflow = val
}

// GetAssetTag returns the AssetTag property
func (m PatchedWritableDeviceWithConfigContextRequest) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetCluster returns the Cluster property
func (m PatchedWritableDeviceWithConfigContextRequest) GetCluster() *int32 {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetCluster(val *int32) {
	m.Cluster = val
}

// GetComments returns the Comments property
func (m PatchedWritableDeviceWithConfigContextRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableDeviceWithConfigContextRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableDeviceWithConfigContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceRole returns the DeviceRole property
func (m PatchedWritableDeviceWithConfigContextRequest) GetDeviceRole() int32 {
	return m.DeviceRole
}

// SetDeviceRole sets the DeviceRole property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetDeviceRole(val int32) {
	m.DeviceRole = val
}

// GetDeviceType returns the DeviceType property
func (m PatchedWritableDeviceWithConfigContextRequest) GetDeviceType() int32 {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetDeviceType(val int32) {
	m.DeviceType = val
}

// GetFace returns the Face property
func (m PatchedWritableDeviceWithConfigContextRequest) GetFace() string {
	return m.Face
}

// SetFace sets the Face property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetFace(val string) {
	m.Face = val
}

// GetLocalContextData returns the LocalContextData property
func (m PatchedWritableDeviceWithConfigContextRequest) GetLocalContextData() map[string]interface{} {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetLocalContextData(val map[string]interface{}) {
	m.LocalContextData = val
}

// GetLocation returns the Location property
func (m PatchedWritableDeviceWithConfigContextRequest) GetLocation() *int32 {
	return m.Location
}

// SetLocation sets the Location property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetLocation(val *int32) {
	m.Location = val
}

// GetName returns the Name property
func (m PatchedWritableDeviceWithConfigContextRequest) GetName() *string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetName(val *string) {
	m.Name = val
}

// GetPlatform returns the Platform property
func (m PatchedWritableDeviceWithConfigContextRequest) GetPlatform() *int32 {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetPlatform(val *int32) {
	m.Platform = val
}

// GetPosition returns the Position property
func (m PatchedWritableDeviceWithConfigContextRequest) GetPosition() *float64 {
	return m.Position
}

// SetPosition sets the Position property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetPosition(val *float64) {
	m.Position = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp4() *int32 {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp4(val *int32) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp6() *int32 {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp6(val *int32) {
	m.PrimaryIp6 = val
}

// GetRack returns the Rack property
func (m PatchedWritableDeviceWithConfigContextRequest) GetRack() *int32 {
	return m.Rack
}

// SetRack sets the Rack property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetRack(val *int32) {
	m.Rack = val
}

// GetSerial returns the Serial property
func (m PatchedWritableDeviceWithConfigContextRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetSerial(val string) {
	m.Serial = val
}

// GetSite returns the Site property
func (m PatchedWritableDeviceWithConfigContextRequest) GetSite() int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetSite(val int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m PatchedWritableDeviceWithConfigContextRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableDeviceWithConfigContextRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableDeviceWithConfigContextRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetVcPosition returns the VcPosition property
func (m PatchedWritableDeviceWithConfigContextRequest) GetVcPosition() *int32 {
	return m.VcPosition
}

// SetVcPosition sets the VcPosition property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetVcPosition(val *int32) {
	m.VcPosition = val
}

// GetVcPriority returns the VcPriority property
func (m PatchedWritableDeviceWithConfigContextRequest) GetVcPriority() *int32 {
	return m.VcPriority
}

// SetVcPriority sets the VcPriority property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetVcPriority(val *int32) {
	m.VcPriority = val
}

// GetVirtualChassis returns the VirtualChassis property
func (m PatchedWritableDeviceWithConfigContextRequest) GetVirtualChassis() *int32 {
	return m.VirtualChassis
}

// SetVirtualChassis sets the VirtualChassis property
func (m *PatchedWritableDeviceWithConfigContextRequest) SetVirtualChassis(val *int32) {
	m.VirtualChassis = val
}
