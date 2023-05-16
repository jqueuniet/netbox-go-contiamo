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

// DeviceWithConfigContext is an object. Adds support for custom fields and tags.
type DeviceWithConfigContext struct {
	// Airflow:
	Airflow *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"airflow,omitempty" mapstructure:"airflow,omitempty"`
	// AssetTag: A unique tag used to identify this device
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
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
	// DeviceRole: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceRole NestedDeviceRole `json:"device_role" mapstructure:"device_role"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType NestedDeviceType `json:"device_type" mapstructure:"device_type"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Face:
	Face *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"face,omitempty" mapstructure:"face,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// LocalContextData: Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData map[string]interface{} `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// Location: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Location *NestedLocation `json:"location,omitempty" mapstructure:"location,omitempty"`
	// Name:
	Name *string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// ParentDevice: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ParentDevice NestedDevice `json:"parent_device" mapstructure:"parent_device"`
	// Platform: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Platform *NestedPlatform `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// Position:
	Position *float64 `json:"position,omitempty" mapstructure:"position,omitempty"`
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
	// Rack: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rack *NestedRack `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// Serial: Chassis serial number, assigned by the manufacturer
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site NestedSite `json:"site" mapstructure:"site"`
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
	// VcPosition:
	VcPosition *int32 `json:"vc_position,omitempty" mapstructure:"vc_position,omitempty"`
	// VcPriority: Virtual chassis master election priority
	VcPriority *int32 `json:"vc_priority,omitempty" mapstructure:"vc_priority,omitempty"`
	// VirtualChassis: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	VirtualChassis *NestedVirtualChassis `json:"virtual_chassis,omitempty" mapstructure:"virtual_chassis,omitempty"`
}

// Validate implements basic validation for this model
func (m DeviceWithConfigContext) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag, validation.Length(0, 50),
		),
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
		"parentDevice": validation.Validate(
			m.ParentDevice, validation.NotNil,
		),
		"platform": validation.Validate(
			m.Platform,
		),
		"position": validation.Validate(
			m.Position, validation.Min(*float64(0.5)), validation.Max(*float64(1000)),
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
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
func (m DeviceWithConfigContext) GetAirflow() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *DeviceWithConfigContext) SetAirflow(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Airflow = val
}

// GetAssetTag returns the AssetTag property
func (m DeviceWithConfigContext) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *DeviceWithConfigContext) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetCluster returns the Cluster property
func (m DeviceWithConfigContext) GetCluster() *NestedCluster {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *DeviceWithConfigContext) SetCluster(val *NestedCluster) {
	m.Cluster = val
}

// GetComments returns the Comments property
func (m DeviceWithConfigContext) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *DeviceWithConfigContext) SetComments(val string) {
	m.Comments = val
}

// GetConfigContext returns the ConfigContext property
func (m DeviceWithConfigContext) GetConfigContext() map[string]interface{} {
	return m.ConfigContext
}

// SetConfigContext sets the ConfigContext property
func (m *DeviceWithConfigContext) SetConfigContext(val map[string]interface{}) {
	m.ConfigContext = val
}

// GetCreated returns the Created property
func (m DeviceWithConfigContext) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DeviceWithConfigContext) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m DeviceWithConfigContext) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *DeviceWithConfigContext) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m DeviceWithConfigContext) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DeviceWithConfigContext) SetDescription(val string) {
	m.Description = val
}

// GetDeviceRole returns the DeviceRole property
func (m DeviceWithConfigContext) GetDeviceRole() NestedDeviceRole {
	return m.DeviceRole
}

// SetDeviceRole sets the DeviceRole property
func (m *DeviceWithConfigContext) SetDeviceRole(val NestedDeviceRole) {
	m.DeviceRole = val
}

// GetDeviceType returns the DeviceType property
func (m DeviceWithConfigContext) GetDeviceType() NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *DeviceWithConfigContext) SetDeviceType(val NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m DeviceWithConfigContext) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *DeviceWithConfigContext) SetDisplay(val string) {
	m.Display = val
}

// GetFace returns the Face property
func (m DeviceWithConfigContext) GetFace() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Face
}

// SetFace sets the Face property
func (m *DeviceWithConfigContext) SetFace(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Face = val
}

// GetId returns the Id property
func (m DeviceWithConfigContext) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DeviceWithConfigContext) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m DeviceWithConfigContext) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DeviceWithConfigContext) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLocalContextData returns the LocalContextData property
func (m DeviceWithConfigContext) GetLocalContextData() map[string]interface{} {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *DeviceWithConfigContext) SetLocalContextData(val map[string]interface{}) {
	m.LocalContextData = val
}

// GetLocation returns the Location property
func (m DeviceWithConfigContext) GetLocation() *NestedLocation {
	return m.Location
}

// SetLocation sets the Location property
func (m *DeviceWithConfigContext) SetLocation(val *NestedLocation) {
	m.Location = val
}

// GetName returns the Name property
func (m DeviceWithConfigContext) GetName() *string {
	return m.Name
}

// SetName sets the Name property
func (m *DeviceWithConfigContext) SetName(val *string) {
	m.Name = val
}

// GetParentDevice returns the ParentDevice property
func (m DeviceWithConfigContext) GetParentDevice() NestedDevice {
	return m.ParentDevice
}

// SetParentDevice sets the ParentDevice property
func (m *DeviceWithConfigContext) SetParentDevice(val NestedDevice) {
	m.ParentDevice = val
}

// GetPlatform returns the Platform property
func (m DeviceWithConfigContext) GetPlatform() *NestedPlatform {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *DeviceWithConfigContext) SetPlatform(val *NestedPlatform) {
	m.Platform = val
}

// GetPosition returns the Position property
func (m DeviceWithConfigContext) GetPosition() *float64 {
	return m.Position
}

// SetPosition sets the Position property
func (m *DeviceWithConfigContext) SetPosition(val *float64) {
	m.Position = val
}

// GetPrimaryIp returns the PrimaryIp property
func (m DeviceWithConfigContext) GetPrimaryIp() NestedIPAddress {
	return m.PrimaryIp
}

// SetPrimaryIp sets the PrimaryIp property
func (m *DeviceWithConfigContext) SetPrimaryIp(val NestedIPAddress) {
	m.PrimaryIp = val
}

// GetPrimaryIp4 returns the PrimaryIp4 property
func (m DeviceWithConfigContext) GetPrimaryIp4() *NestedIPAddress {
	return m.PrimaryIp4
}

// SetPrimaryIp4 sets the PrimaryIp4 property
func (m *DeviceWithConfigContext) SetPrimaryIp4(val *NestedIPAddress) {
	m.PrimaryIp4 = val
}

// GetPrimaryIp6 returns the PrimaryIp6 property
func (m DeviceWithConfigContext) GetPrimaryIp6() *NestedIPAddress {
	return m.PrimaryIp6
}

// SetPrimaryIp6 sets the PrimaryIp6 property
func (m *DeviceWithConfigContext) SetPrimaryIp6(val *NestedIPAddress) {
	m.PrimaryIp6 = val
}

// GetRack returns the Rack property
func (m DeviceWithConfigContext) GetRack() *NestedRack {
	return m.Rack
}

// SetRack sets the Rack property
func (m *DeviceWithConfigContext) SetRack(val *NestedRack) {
	m.Rack = val
}

// GetSerial returns the Serial property
func (m DeviceWithConfigContext) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *DeviceWithConfigContext) SetSerial(val string) {
	m.Serial = val
}

// GetSite returns the Site property
func (m DeviceWithConfigContext) GetSite() NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *DeviceWithConfigContext) SetSite(val NestedSite) {
	m.Site = val
}

// GetStatus returns the Status property
func (m DeviceWithConfigContext) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *DeviceWithConfigContext) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m DeviceWithConfigContext) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *DeviceWithConfigContext) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m DeviceWithConfigContext) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DeviceWithConfigContext) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetUrl returns the Url property
func (m DeviceWithConfigContext) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *DeviceWithConfigContext) SetUrl(val string) {
	m.Url = val
}

// GetVcPosition returns the VcPosition property
func (m DeviceWithConfigContext) GetVcPosition() *int32 {
	return m.VcPosition
}

// SetVcPosition sets the VcPosition property
func (m *DeviceWithConfigContext) SetVcPosition(val *int32) {
	m.VcPosition = val
}

// GetVcPriority returns the VcPriority property
func (m DeviceWithConfigContext) GetVcPriority() *int32 {
	return m.VcPriority
}

// SetVcPriority sets the VcPriority property
func (m *DeviceWithConfigContext) SetVcPriority(val *int32) {
	m.VcPriority = val
}

// GetVirtualChassis returns the VirtualChassis property
func (m DeviceWithConfigContext) GetVirtualChassis() *NestedVirtualChassis {
	return m.VirtualChassis
}

// SetVirtualChassis sets the VirtualChassis property
func (m *DeviceWithConfigContext) SetVirtualChassis(val *NestedVirtualChassis) {
	m.VirtualChassis = val
}
