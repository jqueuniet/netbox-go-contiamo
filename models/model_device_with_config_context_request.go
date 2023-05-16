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

// DeviceWithConfigContextRequest is an object. Adds support for custom fields and tags.
type DeviceWithConfigContextRequest struct {
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
	// DeviceRole: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceRole NestedDeviceRoleRequest `json:"device_role" mapstructure:"device_role"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType NestedDeviceTypeRequest `json:"device_type" mapstructure:"device_type"`
	// Face: * `front` - Front
	// * `rear` - Rear
	Face string `json:"face,omitempty" mapstructure:"face,omitempty"`
	// LocalContextData: Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// Location: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Location *NestedLocationRequest `json:"location,omitempty" mapstructure:"location,omitempty"`
	// Name:
	Name *string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Platform: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Platform *NestedPlatformRequest `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// Position:
	Position *float64 `json:"position,omitempty" mapstructure:"position,omitempty"`
	// PrimaryIp4: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp4 *NestedIPAddressRequest `json:"primary_ip4,omitempty" mapstructure:"primary_ip4,omitempty"`
	// PrimaryIp6: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PrimaryIp6 *NestedIPAddressRequest `json:"primary_ip6,omitempty" mapstructure:"primary_ip6,omitempty"`
	// Rack: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rack *NestedRackRequest `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// Serial: Chassis serial number, assigned by the manufacturer
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site NestedSiteRequest `json:"site" mapstructure:"site"`
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
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenantRequest `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// VcPosition:
	VcPosition *int32 `json:"vc_position,omitempty" mapstructure:"vc_position,omitempty"`
	// VcPriority: Virtual chassis master election priority
	VcPriority *int32 `json:"vc_priority,omitempty" mapstructure:"vc_priority,omitempty"`
	// VirtualChassis: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	VirtualChassis *NestedVirtualChassisRequest `json:"virtual_chassis,omitempty" mapstructure:"virtual_chassis,omitempty"`
}

// Validate implements basic validation for this model
func (m DeviceWithConfigContextRequest) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag, validation.Length(0, 50),
		),
		"cluster": validation.Validate(
			m.Cluster,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"deviceRole": validation.Validate(
			m.DeviceRole, validation.NotNil,
		),
		"deviceType": validation.Validate(
			m.DeviceType, validation.NotNil,
		),
		"localContextData": validation.Validate(
			m.LocalContextData,
		),
		"location": validation.Validate(
			m.Location,
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 64),
		),
		"platform": validation.Validate(
			m.Platform,
		),
		"position": validation.Validate(
			m.Position, validation.Min(*float64(0.5)), validation.Max(*float64(1000)),
		),
		"primaryIp4": validation.Validate(
			m.PrimaryIp4,
		),
		"primaryIp6": validation.Validate(
			m.PrimaryIp6,
		),
		"rack": validation.Validate(
			m.Rack,
		),
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
		"site": validation.Validate(
			m.Site, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"vcPosition": validation.Validate(
			m.VcPosition, validation.Min(*int32(0)), validation.Max(*int32(255)),
		),
		"vcPriority": validation.Validate(
			m.VcPriority, validation.Min(*int32(0)), validation.Max(*int32(255)),
		),
		"virtualChassis": validation.Validate(
			m.VirtualChassis,
		),
	}.Filter()
}

// GetAirflow returns the Airflow property
func (m DeviceWithConfigContextRequest) GetAirflow() string {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *DeviceWithConfigContextRequest) SetAirflow(val string) {
	m.Airflow = val
}

// GetAssetTag returns the AssetTag property
func (m DeviceWithConfigContextRequest) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *DeviceWithConfigContextRequest) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetCluster returns the Cluster property
func (m DeviceWithConfigContextRequest) GetCluster() *NestedClusterRequest {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *DeviceWithConfigContextRequest) SetCluster(val *NestedClusterRequest) {
	m.Cluster = val
}

// GetComments returns the Comments property
func (m DeviceWithConfigContextRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *DeviceWithConfigContextRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m DeviceWithConfigContextRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *DeviceWithConfigContextRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m DeviceWithConfigContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceWithConfigContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceRole returns the DeviceRole property
func (m DeviceWithConfigContextRequest) GetDeviceRole() NestedDeviceRoleRequest {
	return m.DeviceRole
}

// SetDeviceRole sets the DeviceRole property
func (m *DeviceWithConfigContextRequest) SetDeviceRole(val NestedDeviceRoleRequest) {
	m.DeviceRole = val
}

// GetDeviceType returns the DeviceType property
func (m DeviceWithConfigContextRequest) GetDeviceType() NestedDeviceTypeRequest {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *DeviceWithConfigContextRequest) SetDeviceType(val NestedDeviceTypeRequest) {
	m.DeviceType = val
}

// GetFace returns the Face property
func (m DeviceWithConfigContextRequest) GetFace() string {
	return m.Face
}

// SetFace sets the Face property
func (m *DeviceWithConfigContextRequest) SetFace(val string) {
	m.Face = val
}

// GetLocalContextData returns the LocalContextData property
func (m DeviceWithConfigContextRequest) GetLocalContextData() map[string]interface{} {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *DeviceWithConfigContextRequest) SetLocalContextData(val map[string]interface{}) {
	m.LocalContextData = val
}

// GetLocation returns the Location property
func (m DeviceWithConfigContextRequest) GetLocation() *NestedLocationRequest {
	return m.Location
}

// SetLocation sets the Location property
func (m *DeviceWithConfigContextRequest) SetLocation(val *NestedLocationRequest) {
	m.Location = val
}

// GetName returns the Name property
func (m DeviceWithConfigContextRequest) GetName() *string {
	return m.Name
}

// SetName sets the Name property
func (m *DeviceWithConfigContextRequest) SetName(val *string) {
	m.Name = val
}

// GetPlatform returns the Platform property
func (m DeviceWithConfigContextRequest) GetPlatform() *NestedPlatformRequest {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *DeviceWithConfigContextRequest) SetPlatform(val *NestedPlatformRequest) {
	m.Platform = val
}

// GetPosition returns the Position property
func (m DeviceWithConfigContextRequest) GetPosition() *float64 {
	return m.Position
}

// SetPosition sets the Position property
func (m *DeviceWithConfigContextRequest) SetPosition(val *float64) {
	m.Position = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m DeviceWithConfigContextRequest) GetPrimaryIp4() *NestedIPAddressRequest {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *DeviceWithConfigContextRequest) SetPrimaryIp4(val *NestedIPAddressRequest) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m DeviceWithConfigContextRequest) GetPrimaryIp6() *NestedIPAddressRequest {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *DeviceWithConfigContextRequest) SetPrimaryIp6(val *NestedIPAddressRequest) {
	m.PrimaryIp6 = val
}

// GetRack returns the Rack property
func (m DeviceWithConfigContextRequest) GetRack() *NestedRackRequest {
	return m.Rack
}

// SetRack sets the Rack property
func (m *DeviceWithConfigContextRequest) SetRack(val *NestedRackRequest) {
	m.Rack = val
}

// GetSerial returns the Serial property
func (m DeviceWithConfigContextRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *DeviceWithConfigContextRequest) SetSerial(val string) {
	m.Serial = val
}

// GetSite returns the Site property
func (m DeviceWithConfigContextRequest) GetSite() NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *DeviceWithConfigContextRequest) SetSite(val NestedSiteRequest) {
	m.Site = val
}

// GetStatus returns the Status property
func (m DeviceWithConfigContextRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *DeviceWithConfigContextRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m DeviceWithConfigContextRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *DeviceWithConfigContextRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m DeviceWithConfigContextRequest) GetTenant() *NestedTenantRequest {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DeviceWithConfigContextRequest) SetTenant(val *NestedTenantRequest) {
	m.Tenant = val
}

// GetVcPosition returns the VcPosition property
func (m DeviceWithConfigContextRequest) GetVcPosition() *int32 {
	return m.VcPosition
}

// SetVcPosition sets the VcPosition property
func (m *DeviceWithConfigContextRequest) SetVcPosition(val *int32) {
	m.VcPosition = val
}

// GetVcPriority returns the VcPriority property
func (m DeviceWithConfigContextRequest) GetVcPriority() *int32 {
	return m.VcPriority
}

// SetVcPriority sets the VcPriority property
func (m *DeviceWithConfigContextRequest) SetVcPriority(val *int32) {
	m.VcPriority = val
}

// GetVirtualChassis returns the VirtualChassis property
func (m DeviceWithConfigContextRequest) GetVirtualChassis() *NestedVirtualChassisRequest {
	return m.VirtualChassis
}

// SetVirtualChassis sets the VirtualChassis property
func (m *DeviceWithConfigContextRequest) SetVirtualChassis(val *NestedVirtualChassisRequest) {
	m.VirtualChassis = val
}
