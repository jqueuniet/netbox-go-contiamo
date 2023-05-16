// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// DcimDevicesListQueryParameters is an object.
type DcimDevicesListQueryParameters struct {
	// Airflow: * `front-to-rear` - Front to rear
	// * `rear-to-front` - Rear to front
	// * `left-to-right` - Left to right
	// * `right-to-left` - Right to left
	// * `side-to-rear` - Side to rear
	// * `passive` - Passive
	// * `mixed` - Mixed
	Airflow string `json:"airflow,omitempty" mapstructure:"airflow,omitempty"`
	// AirflowN: * `front-to-rear` - Front to rear
	// * `rear-to-front` - Rear to front
	// * `left-to-right` - Left to right
	// * `right-to-left` - Right to left
	// * `side-to-rear` - Side to rear
	// * `passive` - Passive
	// * `mixed` - Mixed
	AirflowN string `json:"airflow__n,omitempty" mapstructure:"airflow__n,omitempty"`
	// AssetTag:
	AssetTag []string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// AssetTagEmpty:
	AssetTagEmpty []string `json:"asset_tag__empty,omitempty" mapstructure:"asset_tag__empty,omitempty"`
	// AssetTagIc:
	AssetTagIc []string `json:"asset_tag__ic,omitempty" mapstructure:"asset_tag__ic,omitempty"`
	// AssetTagIe:
	AssetTagIe []string `json:"asset_tag__ie,omitempty" mapstructure:"asset_tag__ie,omitempty"`
	// AssetTagIew:
	AssetTagIew []string `json:"asset_tag__iew,omitempty" mapstructure:"asset_tag__iew,omitempty"`
	// AssetTagIsw:
	AssetTagIsw []string `json:"asset_tag__isw,omitempty" mapstructure:"asset_tag__isw,omitempty"`
	// AssetTagN:
	AssetTagN []string `json:"asset_tag__n,omitempty" mapstructure:"asset_tag__n,omitempty"`
	// AssetTagNic:
	AssetTagNic []string `json:"asset_tag__nic,omitempty" mapstructure:"asset_tag__nic,omitempty"`
	// AssetTagNie:
	AssetTagNie []string `json:"asset_tag__nie,omitempty" mapstructure:"asset_tag__nie,omitempty"`
	// AssetTagNiew:
	AssetTagNiew []string `json:"asset_tag__niew,omitempty" mapstructure:"asset_tag__niew,omitempty"`
	// AssetTagNisw:
	AssetTagNisw []string `json:"asset_tag__nisw,omitempty" mapstructure:"asset_tag__nisw,omitempty"`
	// ClusterId: VM cluster (ID)
	ClusterId []*int32 `json:"cluster_id,omitempty" mapstructure:"cluster_id,omitempty"`
	// ClusterIdN: VM cluster (ID)
	ClusterIdN []*int32 `json:"cluster_id__n,omitempty" mapstructure:"cluster_id__n,omitempty"`
	// ConfigTemplateId: Config template (ID)
	ConfigTemplateId []*int32 `json:"config_template_id,omitempty" mapstructure:"config_template_id,omitempty"`
	// ConfigTemplateIdN: Config template (ID)
	ConfigTemplateIdN []*int32 `json:"config_template_id__n,omitempty" mapstructure:"config_template_id__n,omitempty"`
	// ConsolePorts: Has console ports
	ConsolePorts bool `json:"console_ports,omitempty" mapstructure:"console_ports,omitempty"`
	// ConsoleServerPorts: Has console server ports
	ConsoleServerPorts bool `json:"console_server_ports,omitempty" mapstructure:"console_server_ports,omitempty"`
	// Contact: Contact
	Contact []int32 `json:"contact,omitempty" mapstructure:"contact,omitempty"`
	// ContactN: Contact
	ContactN []int32 `json:"contact__n,omitempty" mapstructure:"contact__n,omitempty"`
	// ContactGroup: Contact group
	ContactGroup []int32 `json:"contact_group,omitempty" mapstructure:"contact_group,omitempty"`
	// ContactGroupN: Contact group
	ContactGroupN []int32 `json:"contact_group__n,omitempty" mapstructure:"contact_group__n,omitempty"`
	// ContactRole: Contact Role
	ContactRole []int32 `json:"contact_role,omitempty" mapstructure:"contact_role,omitempty"`
	// ContactRoleN: Contact Role
	ContactRoleN []int32 `json:"contact_role__n,omitempty" mapstructure:"contact_role__n,omitempty"`
	// Created:
	Created []time.Time `json:"created,omitempty" mapstructure:"created,omitempty"`
	// CreatedGt:
	CreatedGt []time.Time `json:"created__gt,omitempty" mapstructure:"created__gt,omitempty"`
	// CreatedGte:
	CreatedGte []time.Time `json:"created__gte,omitempty" mapstructure:"created__gte,omitempty"`
	// CreatedLt:
	CreatedLt []time.Time `json:"created__lt,omitempty" mapstructure:"created__lt,omitempty"`
	// CreatedLte:
	CreatedLte []time.Time `json:"created__lte,omitempty" mapstructure:"created__lte,omitempty"`
	// CreatedN:
	CreatedN []time.Time `json:"created__n,omitempty" mapstructure:"created__n,omitempty"`
	// CreatedByRequest:
	CreatedByRequest string `json:"created_by_request,omitempty" mapstructure:"created_by_request,omitempty"`
	// DeviceBays: Has device bays
	DeviceBays bool `json:"device_bays,omitempty" mapstructure:"device_bays,omitempty"`
	// DeviceType: Device type (slug)
	DeviceType []string `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// DeviceTypeN: Device type (slug)
	DeviceTypeN []string `json:"device_type__n,omitempty" mapstructure:"device_type__n,omitempty"`
	// DeviceTypeId: Device type (ID)
	DeviceTypeId []int32 `json:"device_type_id,omitempty" mapstructure:"device_type_id,omitempty"`
	// DeviceTypeIdN: Device type (ID)
	DeviceTypeIdN []int32 `json:"device_type_id__n,omitempty" mapstructure:"device_type_id__n,omitempty"`
	// Face: * `front` - Front
	// * `rear` - Rear
	Face string `json:"face,omitempty" mapstructure:"face,omitempty"`
	// FaceN: * `front` - Front
	// * `rear` - Rear
	FaceN string `json:"face__n,omitempty" mapstructure:"face__n,omitempty"`
	// HasPrimaryIp: Has a primary IP
	HasPrimaryIp bool `json:"has_primary_ip,omitempty" mapstructure:"has_primary_ip,omitempty"`
	// Id:
	Id []int32 `json:"id,omitempty" mapstructure:"id,omitempty"`
	// IdGt:
	IdGt []int32 `json:"id__gt,omitempty" mapstructure:"id__gt,omitempty"`
	// IdGte:
	IdGte []int32 `json:"id__gte,omitempty" mapstructure:"id__gte,omitempty"`
	// IdLt:
	IdLt []int32 `json:"id__lt,omitempty" mapstructure:"id__lt,omitempty"`
	// IdLte:
	IdLte []int32 `json:"id__lte,omitempty" mapstructure:"id__lte,omitempty"`
	// IdN:
	IdN []int32 `json:"id__n,omitempty" mapstructure:"id__n,omitempty"`
	// Interfaces: Has interfaces
	Interfaces bool `json:"interfaces,omitempty" mapstructure:"interfaces,omitempty"`
	// IsFullDepth: Is full depth
	IsFullDepth bool `json:"is_full_depth,omitempty" mapstructure:"is_full_depth,omitempty"`
	// LastUpdated:
	LastUpdated []time.Time `json:"last_updated,omitempty" mapstructure:"last_updated,omitempty"`
	// LastUpdatedGt:
	LastUpdatedGt []time.Time `json:"last_updated__gt,omitempty" mapstructure:"last_updated__gt,omitempty"`
	// LastUpdatedGte:
	LastUpdatedGte []time.Time `json:"last_updated__gte,omitempty" mapstructure:"last_updated__gte,omitempty"`
	// LastUpdatedLt:
	LastUpdatedLt []time.Time `json:"last_updated__lt,omitempty" mapstructure:"last_updated__lt,omitempty"`
	// LastUpdatedLte:
	LastUpdatedLte []time.Time `json:"last_updated__lte,omitempty" mapstructure:"last_updated__lte,omitempty"`
	// LastUpdatedN:
	LastUpdatedN []time.Time `json:"last_updated__n,omitempty" mapstructure:"last_updated__n,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// LocalContextData: Has local config context data
	LocalContextData bool `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// LocationId: Location (ID)
	LocationId []int32 `json:"location_id,omitempty" mapstructure:"location_id,omitempty"`
	// LocationIdN: Location (ID)
	LocationIdN []int32 `json:"location_id__n,omitempty" mapstructure:"location_id__n,omitempty"`
	// MacAddress:
	MacAddress []string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// MacAddressIc:
	MacAddressIc []string `json:"mac_address__ic,omitempty" mapstructure:"mac_address__ic,omitempty"`
	// MacAddressIe:
	MacAddressIe []string `json:"mac_address__ie,omitempty" mapstructure:"mac_address__ie,omitempty"`
	// MacAddressIew:
	MacAddressIew []string `json:"mac_address__iew,omitempty" mapstructure:"mac_address__iew,omitempty"`
	// MacAddressIsw:
	MacAddressIsw []string `json:"mac_address__isw,omitempty" mapstructure:"mac_address__isw,omitempty"`
	// MacAddressN:
	MacAddressN []string `json:"mac_address__n,omitempty" mapstructure:"mac_address__n,omitempty"`
	// MacAddressNic:
	MacAddressNic []string `json:"mac_address__nic,omitempty" mapstructure:"mac_address__nic,omitempty"`
	// MacAddressNie:
	MacAddressNie []string `json:"mac_address__nie,omitempty" mapstructure:"mac_address__nie,omitempty"`
	// MacAddressNiew:
	MacAddressNiew []string `json:"mac_address__niew,omitempty" mapstructure:"mac_address__niew,omitempty"`
	// MacAddressNisw:
	MacAddressNisw []string `json:"mac_address__nisw,omitempty" mapstructure:"mac_address__nisw,omitempty"`
	// Manufacturer: Manufacturer (slug)
	Manufacturer []string `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// ManufacturerN: Manufacturer (slug)
	ManufacturerN []string `json:"manufacturer__n,omitempty" mapstructure:"manufacturer__n,omitempty"`
	// ManufacturerId: Manufacturer (ID)
	ManufacturerId []int32 `json:"manufacturer_id,omitempty" mapstructure:"manufacturer_id,omitempty"`
	// ManufacturerIdN: Manufacturer (ID)
	ManufacturerIdN []int32 `json:"manufacturer_id__n,omitempty" mapstructure:"manufacturer_id__n,omitempty"`
	// Model: Device model (slug)
	Model []string `json:"model,omitempty" mapstructure:"model,omitempty"`
	// ModelN: Device model (slug)
	ModelN []string `json:"model__n,omitempty" mapstructure:"model__n,omitempty"`
	// ModuleBays: Has module bays
	ModuleBays bool `json:"module_bays,omitempty" mapstructure:"module_bays,omitempty"`
	// Name:
	Name []string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// NameEmpty:
	NameEmpty []string `json:"name__empty,omitempty" mapstructure:"name__empty,omitempty"`
	// NameIc:
	NameIc []string `json:"name__ic,omitempty" mapstructure:"name__ic,omitempty"`
	// NameIe:
	NameIe []string `json:"name__ie,omitempty" mapstructure:"name__ie,omitempty"`
	// NameIew:
	NameIew []string `json:"name__iew,omitempty" mapstructure:"name__iew,omitempty"`
	// NameIsw:
	NameIsw []string `json:"name__isw,omitempty" mapstructure:"name__isw,omitempty"`
	// NameN:
	NameN []string `json:"name__n,omitempty" mapstructure:"name__n,omitempty"`
	// NameNic:
	NameNic []string `json:"name__nic,omitempty" mapstructure:"name__nic,omitempty"`
	// NameNie:
	NameNie []string `json:"name__nie,omitempty" mapstructure:"name__nie,omitempty"`
	// NameNiew:
	NameNiew []string `json:"name__niew,omitempty" mapstructure:"name__niew,omitempty"`
	// NameNisw:
	NameNisw []string `json:"name__nisw,omitempty" mapstructure:"name__nisw,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// ParentDeviceId: Parent Device (ID)
	ParentDeviceId []int32 `json:"parent_device_id,omitempty" mapstructure:"parent_device_id,omitempty"`
	// ParentDeviceIdN: Parent Device (ID)
	ParentDeviceIdN []int32 `json:"parent_device_id__n,omitempty" mapstructure:"parent_device_id__n,omitempty"`
	// PassThroughPorts: Has pass-through ports
	PassThroughPorts bool `json:"pass_through_ports,omitempty" mapstructure:"pass_through_ports,omitempty"`
	// Platform: Platform (slug)
	Platform []string `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// PlatformN: Platform (slug)
	PlatformN []string `json:"platform__n,omitempty" mapstructure:"platform__n,omitempty"`
	// PlatformId: Platform (ID)
	PlatformId []*int32 `json:"platform_id,omitempty" mapstructure:"platform_id,omitempty"`
	// PlatformIdN: Platform (ID)
	PlatformIdN []*int32 `json:"platform_id__n,omitempty" mapstructure:"platform_id__n,omitempty"`
	// Position:
	Position []float64 `json:"position,omitempty" mapstructure:"position,omitempty"`
	// PositionGt:
	PositionGt []float64 `json:"position__gt,omitempty" mapstructure:"position__gt,omitempty"`
	// PositionGte:
	PositionGte []float64 `json:"position__gte,omitempty" mapstructure:"position__gte,omitempty"`
	// PositionLt:
	PositionLt []float64 `json:"position__lt,omitempty" mapstructure:"position__lt,omitempty"`
	// PositionLte:
	PositionLte []float64 `json:"position__lte,omitempty" mapstructure:"position__lte,omitempty"`
	// PositionN:
	PositionN []float64 `json:"position__n,omitempty" mapstructure:"position__n,omitempty"`
	// PowerOutlets: Has power outlets
	PowerOutlets bool `json:"power_outlets,omitempty" mapstructure:"power_outlets,omitempty"`
	// PowerPorts: Has power ports
	PowerPorts bool `json:"power_ports,omitempty" mapstructure:"power_ports,omitempty"`
	// PrimaryIp4Id: Primary IPv4 (ID)
	PrimaryIp4Id []int32 `json:"primary_ip4_id,omitempty" mapstructure:"primary_ip4_id,omitempty"`
	// PrimaryIp4IdN: Primary IPv4 (ID)
	PrimaryIp4IdN []int32 `json:"primary_ip4_id__n,omitempty" mapstructure:"primary_ip4_id__n,omitempty"`
	// PrimaryIp6Id: Primary IPv6 (ID)
	PrimaryIp6Id []int32 `json:"primary_ip6_id,omitempty" mapstructure:"primary_ip6_id,omitempty"`
	// PrimaryIp6IdN: Primary IPv6 (ID)
	PrimaryIp6IdN []int32 `json:"primary_ip6_id__n,omitempty" mapstructure:"primary_ip6_id__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// RackId: Rack (ID)
	RackId []int32 `json:"rack_id,omitempty" mapstructure:"rack_id,omitempty"`
	// RackIdN: Rack (ID)
	RackIdN []int32 `json:"rack_id__n,omitempty" mapstructure:"rack_id__n,omitempty"`
	// Region: Region (slug)
	Region []int32 `json:"region,omitempty" mapstructure:"region,omitempty"`
	// RegionN: Region (slug)
	RegionN []int32 `json:"region__n,omitempty" mapstructure:"region__n,omitempty"`
	// RegionId: Region (ID)
	RegionId []int32 `json:"region_id,omitempty" mapstructure:"region_id,omitempty"`
	// RegionIdN: Region (ID)
	RegionIdN []int32 `json:"region_id__n,omitempty" mapstructure:"region_id__n,omitempty"`
	// Role: Role (slug)
	Role []string `json:"role,omitempty" mapstructure:"role,omitempty"`
	// RoleN: Role (slug)
	RoleN []string `json:"role__n,omitempty" mapstructure:"role__n,omitempty"`
	// RoleId: Role (ID)
	RoleId []int32 `json:"role_id,omitempty" mapstructure:"role_id,omitempty"`
	// RoleIdN: Role (ID)
	RoleIdN []int32 `json:"role_id__n,omitempty" mapstructure:"role_id__n,omitempty"`
	// Serial:
	Serial []string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// SerialEmpty:
	SerialEmpty []string `json:"serial__empty,omitempty" mapstructure:"serial__empty,omitempty"`
	// SerialIc:
	SerialIc []string `json:"serial__ic,omitempty" mapstructure:"serial__ic,omitempty"`
	// SerialIe:
	SerialIe []string `json:"serial__ie,omitempty" mapstructure:"serial__ie,omitempty"`
	// SerialIew:
	SerialIew []string `json:"serial__iew,omitempty" mapstructure:"serial__iew,omitempty"`
	// SerialIsw:
	SerialIsw []string `json:"serial__isw,omitempty" mapstructure:"serial__isw,omitempty"`
	// SerialN:
	SerialN []string `json:"serial__n,omitempty" mapstructure:"serial__n,omitempty"`
	// SerialNic:
	SerialNic []string `json:"serial__nic,omitempty" mapstructure:"serial__nic,omitempty"`
	// SerialNie:
	SerialNie []string `json:"serial__nie,omitempty" mapstructure:"serial__nie,omitempty"`
	// SerialNiew:
	SerialNiew []string `json:"serial__niew,omitempty" mapstructure:"serial__niew,omitempty"`
	// SerialNisw:
	SerialNisw []string `json:"serial__nisw,omitempty" mapstructure:"serial__nisw,omitempty"`
	// Site: Site name (slug)
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteN: Site name (slug)
	SiteN []string `json:"site__n,omitempty" mapstructure:"site__n,omitempty"`
	// SiteGroup: Site group (slug)
	SiteGroup []int32 `json:"site_group,omitempty" mapstructure:"site_group,omitempty"`
	// SiteGroupN: Site group (slug)
	SiteGroupN []int32 `json:"site_group__n,omitempty" mapstructure:"site_group__n,omitempty"`
	// SiteGroupId: Site group (ID)
	SiteGroupId []int32 `json:"site_group_id,omitempty" mapstructure:"site_group_id,omitempty"`
	// SiteGroupIdN: Site group (ID)
	SiteGroupIdN []int32 `json:"site_group_id__n,omitempty" mapstructure:"site_group_id__n,omitempty"`
	// SiteId: Site (ID)
	SiteId []int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
	// SiteIdN: Site (ID)
	SiteIdN []int32 `json:"site_id__n,omitempty" mapstructure:"site_id__n,omitempty"`
	// Status: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
	// * `inventory` - Inventory
	// * `decommissioning` - Decommissioning
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
	// * `inventory` - Inventory
	// * `decommissioning` - Decommissioning
	StatusN []string `json:"status__n,omitempty" mapstructure:"status__n,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// Tenant: Tenant (slug)
	Tenant []string `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TenantN: Tenant (slug)
	TenantN []string `json:"tenant__n,omitempty" mapstructure:"tenant__n,omitempty"`
	// TenantGroup: Tenant Group (slug)
	TenantGroup []int32 `json:"tenant_group,omitempty" mapstructure:"tenant_group,omitempty"`
	// TenantGroupN: Tenant Group (slug)
	TenantGroupN []int32 `json:"tenant_group__n,omitempty" mapstructure:"tenant_group__n,omitempty"`
	// TenantGroupId: Tenant Group (ID)
	TenantGroupId []int32 `json:"tenant_group_id,omitempty" mapstructure:"tenant_group_id,omitempty"`
	// TenantGroupIdN: Tenant Group (ID)
	TenantGroupIdN []int32 `json:"tenant_group_id__n,omitempty" mapstructure:"tenant_group_id__n,omitempty"`
	// TenantId: Tenant (ID)
	TenantId []*int32 `json:"tenant_id,omitempty" mapstructure:"tenant_id,omitempty"`
	// TenantIdN: Tenant (ID)
	TenantIdN []*int32 `json:"tenant_id__n,omitempty" mapstructure:"tenant_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// VcPosition:
	VcPosition []int32 `json:"vc_position,omitempty" mapstructure:"vc_position,omitempty"`
	// VcPositionGt:
	VcPositionGt []int32 `json:"vc_position__gt,omitempty" mapstructure:"vc_position__gt,omitempty"`
	// VcPositionGte:
	VcPositionGte []int32 `json:"vc_position__gte,omitempty" mapstructure:"vc_position__gte,omitempty"`
	// VcPositionLt:
	VcPositionLt []int32 `json:"vc_position__lt,omitempty" mapstructure:"vc_position__lt,omitempty"`
	// VcPositionLte:
	VcPositionLte []int32 `json:"vc_position__lte,omitempty" mapstructure:"vc_position__lte,omitempty"`
	// VcPositionN:
	VcPositionN []int32 `json:"vc_position__n,omitempty" mapstructure:"vc_position__n,omitempty"`
	// VcPriority:
	VcPriority []int32 `json:"vc_priority,omitempty" mapstructure:"vc_priority,omitempty"`
	// VcPriorityGt:
	VcPriorityGt []int32 `json:"vc_priority__gt,omitempty" mapstructure:"vc_priority__gt,omitempty"`
	// VcPriorityGte:
	VcPriorityGte []int32 `json:"vc_priority__gte,omitempty" mapstructure:"vc_priority__gte,omitempty"`
	// VcPriorityLt:
	VcPriorityLt []int32 `json:"vc_priority__lt,omitempty" mapstructure:"vc_priority__lt,omitempty"`
	// VcPriorityLte:
	VcPriorityLte []int32 `json:"vc_priority__lte,omitempty" mapstructure:"vc_priority__lte,omitempty"`
	// VcPriorityN:
	VcPriorityN []int32 `json:"vc_priority__n,omitempty" mapstructure:"vc_priority__n,omitempty"`
	// VirtualChassisId: Virtual chassis (ID)
	VirtualChassisId []int32 `json:"virtual_chassis_id,omitempty" mapstructure:"virtual_chassis_id,omitempty"`
	// VirtualChassisIdN: Virtual chassis (ID)
	VirtualChassisIdN []int32 `json:"virtual_chassis_id__n,omitempty" mapstructure:"virtual_chassis_id__n,omitempty"`
	// VirtualChassisMember: Is a virtual chassis member
	VirtualChassisMember bool `json:"virtual_chassis_member,omitempty" mapstructure:"virtual_chassis_member,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimDevicesListQueryParameters) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag,
		),
		"assetTagEmpty": validation.Validate(
			m.AssetTagEmpty,
		),
		"assetTagIc": validation.Validate(
			m.AssetTagIc,
		),
		"assetTagIe": validation.Validate(
			m.AssetTagIe,
		),
		"assetTagIew": validation.Validate(
			m.AssetTagIew,
		),
		"assetTagIsw": validation.Validate(
			m.AssetTagIsw,
		),
		"assetTagN": validation.Validate(
			m.AssetTagN,
		),
		"assetTagNic": validation.Validate(
			m.AssetTagNic,
		),
		"assetTagNie": validation.Validate(
			m.AssetTagNie,
		),
		"assetTagNiew": validation.Validate(
			m.AssetTagNiew,
		),
		"assetTagNisw": validation.Validate(
			m.AssetTagNisw,
		),
		"clusterId": validation.Validate(
			m.ClusterId,
		),
		"clusterIdN": validation.Validate(
			m.ClusterIdN,
		),
		"configTemplateId": validation.Validate(
			m.ConfigTemplateId,
		),
		"configTemplateIdN": validation.Validate(
			m.ConfigTemplateIdN,
		),
		"contact": validation.Validate(
			m.Contact,
		),
		"contactN": validation.Validate(
			m.ContactN,
		),
		"contactGroup": validation.Validate(
			m.ContactGroup,
		),
		"contactGroupN": validation.Validate(
			m.ContactGroupN,
		),
		"contactRole": validation.Validate(
			m.ContactRole,
		),
		"contactRoleN": validation.Validate(
			m.ContactRoleN,
		),
		"created": validation.Validate(
			m.Created,
		),
		"createdGt": validation.Validate(
			m.CreatedGt,
		),
		"createdGte": validation.Validate(
			m.CreatedGte,
		),
		"createdLt": validation.Validate(
			m.CreatedLt,
		),
		"createdLte": validation.Validate(
			m.CreatedLte,
		),
		"createdN": validation.Validate(
			m.CreatedN,
		),
		"createdByRequest": validation.Validate(
			m.CreatedByRequest, is.UUID,
		),
		"deviceType": validation.Validate(
			m.DeviceType,
		),
		"deviceTypeN": validation.Validate(
			m.DeviceTypeN,
		),
		"deviceTypeId": validation.Validate(
			m.DeviceTypeId,
		),
		"deviceTypeIdN": validation.Validate(
			m.DeviceTypeIdN,
		),
		"id": validation.Validate(
			m.Id,
		),
		"idGt": validation.Validate(
			m.IdGt,
		),
		"idGte": validation.Validate(
			m.IdGte,
		),
		"idLt": validation.Validate(
			m.IdLt,
		),
		"idLte": validation.Validate(
			m.IdLte,
		),
		"idN": validation.Validate(
			m.IdN,
		),
		"lastUpdated": validation.Validate(
			m.LastUpdated,
		),
		"lastUpdatedGt": validation.Validate(
			m.LastUpdatedGt,
		),
		"lastUpdatedGte": validation.Validate(
			m.LastUpdatedGte,
		),
		"lastUpdatedLt": validation.Validate(
			m.LastUpdatedLt,
		),
		"lastUpdatedLte": validation.Validate(
			m.LastUpdatedLte,
		),
		"lastUpdatedN": validation.Validate(
			m.LastUpdatedN,
		),
		"locationId": validation.Validate(
			m.LocationId,
		),
		"locationIdN": validation.Validate(
			m.LocationIdN,
		),
		"macAddress": validation.Validate(
			m.MacAddress,
		),
		"macAddressIc": validation.Validate(
			m.MacAddressIc,
		),
		"macAddressIe": validation.Validate(
			m.MacAddressIe,
		),
		"macAddressIew": validation.Validate(
			m.MacAddressIew,
		),
		"macAddressIsw": validation.Validate(
			m.MacAddressIsw,
		),
		"macAddressN": validation.Validate(
			m.MacAddressN,
		),
		"macAddressNic": validation.Validate(
			m.MacAddressNic,
		),
		"macAddressNie": validation.Validate(
			m.MacAddressNie,
		),
		"macAddressNiew": validation.Validate(
			m.MacAddressNiew,
		),
		"macAddressNisw": validation.Validate(
			m.MacAddressNisw,
		),
		"manufacturer": validation.Validate(
			m.Manufacturer,
		),
		"manufacturerN": validation.Validate(
			m.ManufacturerN,
		),
		"manufacturerId": validation.Validate(
			m.ManufacturerId,
		),
		"manufacturerIdN": validation.Validate(
			m.ManufacturerIdN,
		),
		"model": validation.Validate(
			m.Model,
		),
		"modelN": validation.Validate(
			m.ModelN,
		),
		"name": validation.Validate(
			m.Name,
		),
		"nameEmpty": validation.Validate(
			m.NameEmpty,
		),
		"nameIc": validation.Validate(
			m.NameIc,
		),
		"nameIe": validation.Validate(
			m.NameIe,
		),
		"nameIew": validation.Validate(
			m.NameIew,
		),
		"nameIsw": validation.Validate(
			m.NameIsw,
		),
		"nameN": validation.Validate(
			m.NameN,
		),
		"nameNic": validation.Validate(
			m.NameNic,
		),
		"nameNie": validation.Validate(
			m.NameNie,
		),
		"nameNiew": validation.Validate(
			m.NameNiew,
		),
		"nameNisw": validation.Validate(
			m.NameNisw,
		),
		"parentDeviceId": validation.Validate(
			m.ParentDeviceId,
		),
		"parentDeviceIdN": validation.Validate(
			m.ParentDeviceIdN,
		),
		"platform": validation.Validate(
			m.Platform,
		),
		"platformN": validation.Validate(
			m.PlatformN,
		),
		"platformId": validation.Validate(
			m.PlatformId,
		),
		"platformIdN": validation.Validate(
			m.PlatformIdN,
		),
		"position": validation.Validate(
			m.Position,
		),
		"positionGt": validation.Validate(
			m.PositionGt,
		),
		"positionGte": validation.Validate(
			m.PositionGte,
		),
		"positionLt": validation.Validate(
			m.PositionLt,
		),
		"positionLte": validation.Validate(
			m.PositionLte,
		),
		"positionN": validation.Validate(
			m.PositionN,
		),
		"primaryIp4Id": validation.Validate(
			m.PrimaryIp4Id,
		),
		"primaryIp4IdN": validation.Validate(
			m.PrimaryIp4IdN,
		),
		"primaryIp6Id": validation.Validate(
			m.PrimaryIp6Id,
		),
		"primaryIp6IdN": validation.Validate(
			m.PrimaryIp6IdN,
		),
		"rackId": validation.Validate(
			m.RackId,
		),
		"rackIdN": validation.Validate(
			m.RackIdN,
		),
		"region": validation.Validate(
			m.Region,
		),
		"regionN": validation.Validate(
			m.RegionN,
		),
		"regionId": validation.Validate(
			m.RegionId,
		),
		"regionIdN": validation.Validate(
			m.RegionIdN,
		),
		"role": validation.Validate(
			m.Role,
		),
		"roleN": validation.Validate(
			m.RoleN,
		),
		"roleId": validation.Validate(
			m.RoleId,
		),
		"roleIdN": validation.Validate(
			m.RoleIdN,
		),
		"serial": validation.Validate(
			m.Serial,
		),
		"serialEmpty": validation.Validate(
			m.SerialEmpty,
		),
		"serialIc": validation.Validate(
			m.SerialIc,
		),
		"serialIe": validation.Validate(
			m.SerialIe,
		),
		"serialIew": validation.Validate(
			m.SerialIew,
		),
		"serialIsw": validation.Validate(
			m.SerialIsw,
		),
		"serialN": validation.Validate(
			m.SerialN,
		),
		"serialNic": validation.Validate(
			m.SerialNic,
		),
		"serialNie": validation.Validate(
			m.SerialNie,
		),
		"serialNiew": validation.Validate(
			m.SerialNiew,
		),
		"serialNisw": validation.Validate(
			m.SerialNisw,
		),
		"site": validation.Validate(
			m.Site,
		),
		"siteN": validation.Validate(
			m.SiteN,
		),
		"siteGroup": validation.Validate(
			m.SiteGroup,
		),
		"siteGroupN": validation.Validate(
			m.SiteGroupN,
		),
		"siteGroupId": validation.Validate(
			m.SiteGroupId,
		),
		"siteGroupIdN": validation.Validate(
			m.SiteGroupIdN,
		),
		"siteId": validation.Validate(
			m.SiteId,
		),
		"siteIdN": validation.Validate(
			m.SiteIdN,
		),
		"status": validation.Validate(
			m.Status,
		),
		"statusN": validation.Validate(
			m.StatusN,
		),
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"tenantN": validation.Validate(
			m.TenantN,
		),
		"tenantGroup": validation.Validate(
			m.TenantGroup,
		),
		"tenantGroupN": validation.Validate(
			m.TenantGroupN,
		),
		"tenantGroupId": validation.Validate(
			m.TenantGroupId,
		),
		"tenantGroupIdN": validation.Validate(
			m.TenantGroupIdN,
		),
		"tenantId": validation.Validate(
			m.TenantId,
		),
		"tenantIdN": validation.Validate(
			m.TenantIdN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"vcPosition": validation.Validate(
			m.VcPosition,
		),
		"vcPositionGt": validation.Validate(
			m.VcPositionGt,
		),
		"vcPositionGte": validation.Validate(
			m.VcPositionGte,
		),
		"vcPositionLt": validation.Validate(
			m.VcPositionLt,
		),
		"vcPositionLte": validation.Validate(
			m.VcPositionLte,
		),
		"vcPositionN": validation.Validate(
			m.VcPositionN,
		),
		"vcPriority": validation.Validate(
			m.VcPriority,
		),
		"vcPriorityGt": validation.Validate(
			m.VcPriorityGt,
		),
		"vcPriorityGte": validation.Validate(
			m.VcPriorityGte,
		),
		"vcPriorityLt": validation.Validate(
			m.VcPriorityLt,
		),
		"vcPriorityLte": validation.Validate(
			m.VcPriorityLte,
		),
		"vcPriorityN": validation.Validate(
			m.VcPriorityN,
		),
		"virtualChassisId": validation.Validate(
			m.VirtualChassisId,
		),
		"virtualChassisIdN": validation.Validate(
			m.VirtualChassisIdN,
		),
	}.Filter()
}

// GetAirflow returns the Airflow property
func (m DcimDevicesListQueryParameters) GetAirflow() string {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *DcimDevicesListQueryParameters) SetAirflow(val string) {
	m.Airflow = val
}

// GetAirflowN returns the AirflowN property
func (m DcimDevicesListQueryParameters) GetAirflowN() string {
	return m.AirflowN
}

// SetAirflowN sets the AirflowN property
func (m *DcimDevicesListQueryParameters) SetAirflowN(val string) {
	m.AirflowN = val
}

// GetAssetTag returns the AssetTag property
func (m DcimDevicesListQueryParameters) GetAssetTag() []string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *DcimDevicesListQueryParameters) SetAssetTag(val []string) {
	m.AssetTag = val
}

// GetAssetTagEmpty returns the AssetTagEmpty property
func (m DcimDevicesListQueryParameters) GetAssetTagEmpty() []string {
	return m.AssetTagEmpty
}

// SetAssetTagEmpty sets the AssetTagEmpty property
func (m *DcimDevicesListQueryParameters) SetAssetTagEmpty(val []string) {
	m.AssetTagEmpty = val
}

// GetAssetTagIc returns the AssetTagIc property
func (m DcimDevicesListQueryParameters) GetAssetTagIc() []string {
	return m.AssetTagIc
}

// SetAssetTagIc sets the AssetTagIc property
func (m *DcimDevicesListQueryParameters) SetAssetTagIc(val []string) {
	m.AssetTagIc = val
}

// GetAssetTagIe returns the AssetTagIe property
func (m DcimDevicesListQueryParameters) GetAssetTagIe() []string {
	return m.AssetTagIe
}

// SetAssetTagIe sets the AssetTagIe property
func (m *DcimDevicesListQueryParameters) SetAssetTagIe(val []string) {
	m.AssetTagIe = val
}

// GetAssetTagIew returns the AssetTagIew property
func (m DcimDevicesListQueryParameters) GetAssetTagIew() []string {
	return m.AssetTagIew
}

// SetAssetTagIew sets the AssetTagIew property
func (m *DcimDevicesListQueryParameters) SetAssetTagIew(val []string) {
	m.AssetTagIew = val
}

// GetAssetTagIsw returns the AssetTagIsw property
func (m DcimDevicesListQueryParameters) GetAssetTagIsw() []string {
	return m.AssetTagIsw
}

// SetAssetTagIsw sets the AssetTagIsw property
func (m *DcimDevicesListQueryParameters) SetAssetTagIsw(val []string) {
	m.AssetTagIsw = val
}

// GetAssetTagN returns the AssetTagN property
func (m DcimDevicesListQueryParameters) GetAssetTagN() []string {
	return m.AssetTagN
}

// SetAssetTagN sets the AssetTagN property
func (m *DcimDevicesListQueryParameters) SetAssetTagN(val []string) {
	m.AssetTagN = val
}

// GetAssetTagNic returns the AssetTagNic property
func (m DcimDevicesListQueryParameters) GetAssetTagNic() []string {
	return m.AssetTagNic
}

// SetAssetTagNic sets the AssetTagNic property
func (m *DcimDevicesListQueryParameters) SetAssetTagNic(val []string) {
	m.AssetTagNic = val
}

// GetAssetTagNie returns the AssetTagNie property
func (m DcimDevicesListQueryParameters) GetAssetTagNie() []string {
	return m.AssetTagNie
}

// SetAssetTagNie sets the AssetTagNie property
func (m *DcimDevicesListQueryParameters) SetAssetTagNie(val []string) {
	m.AssetTagNie = val
}

// GetAssetTagNiew returns the AssetTagNiew property
func (m DcimDevicesListQueryParameters) GetAssetTagNiew() []string {
	return m.AssetTagNiew
}

// SetAssetTagNiew sets the AssetTagNiew property
func (m *DcimDevicesListQueryParameters) SetAssetTagNiew(val []string) {
	m.AssetTagNiew = val
}

// GetAssetTagNisw returns the AssetTagNisw property
func (m DcimDevicesListQueryParameters) GetAssetTagNisw() []string {
	return m.AssetTagNisw
}

// SetAssetTagNisw sets the AssetTagNisw property
func (m *DcimDevicesListQueryParameters) SetAssetTagNisw(val []string) {
	m.AssetTagNisw = val
}

// GetClusterId returns the ClusterId property
func (m DcimDevicesListQueryParameters) GetClusterId() []*int32 {
	return m.ClusterId
}

// SetClusterId sets the ClusterId property
func (m *DcimDevicesListQueryParameters) SetClusterId(val []*int32) {
	m.ClusterId = val
}

// GetClusterIdN returns the ClusterIdN property
func (m DcimDevicesListQueryParameters) GetClusterIdN() []*int32 {
	return m.ClusterIdN
}

// SetClusterIdN sets the ClusterIdN property
func (m *DcimDevicesListQueryParameters) SetClusterIdN(val []*int32) {
	m.ClusterIdN = val
}

// GetConfigTemplateId returns the ConfigTemplateId property
func (m DcimDevicesListQueryParameters) GetConfigTemplateId() []*int32 {
	return m.ConfigTemplateId
}

// SetConfigTemplateId sets the ConfigTemplateId property
func (m *DcimDevicesListQueryParameters) SetConfigTemplateId(val []*int32) {
	m.ConfigTemplateId = val
}

// GetConfigTemplateIdN returns the ConfigTemplateIdN property
func (m DcimDevicesListQueryParameters) GetConfigTemplateIdN() []*int32 {
	return m.ConfigTemplateIdN
}

// SetConfigTemplateIdN sets the ConfigTemplateIdN property
func (m *DcimDevicesListQueryParameters) SetConfigTemplateIdN(val []*int32) {
	m.ConfigTemplateIdN = val
}

// GetConsolePorts returns the ConsolePorts property
func (m DcimDevicesListQueryParameters) GetConsolePorts() bool {
	return m.ConsolePorts
}

// SetConsolePorts sets the ConsolePorts property
func (m *DcimDevicesListQueryParameters) SetConsolePorts(val bool) {
	m.ConsolePorts = val
}

// GetConsoleServerPorts returns the ConsoleServerPorts property
func (m DcimDevicesListQueryParameters) GetConsoleServerPorts() bool {
	return m.ConsoleServerPorts
}

// SetConsoleServerPorts sets the ConsoleServerPorts property
func (m *DcimDevicesListQueryParameters) SetConsoleServerPorts(val bool) {
	m.ConsoleServerPorts = val
}

// GetContact returns the Contact property
func (m DcimDevicesListQueryParameters) GetContact() []int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *DcimDevicesListQueryParameters) SetContact(val []int32) {
	m.Contact = val
}

// GetContactN returns the ContactN property
func (m DcimDevicesListQueryParameters) GetContactN() []int32 {
	return m.ContactN
}

// SetContactN sets the ContactN property
func (m *DcimDevicesListQueryParameters) SetContactN(val []int32) {
	m.ContactN = val
}

// GetContactGroup returns the ContactGroup property
func (m DcimDevicesListQueryParameters) GetContactGroup() []int32 {
	return m.ContactGroup
}

// SetContactGroup sets the ContactGroup property
func (m *DcimDevicesListQueryParameters) SetContactGroup(val []int32) {
	m.ContactGroup = val
}

// GetContactGroupN returns the ContactGroupN property
func (m DcimDevicesListQueryParameters) GetContactGroupN() []int32 {
	return m.ContactGroupN
}

// SetContactGroupN sets the ContactGroupN property
func (m *DcimDevicesListQueryParameters) SetContactGroupN(val []int32) {
	m.ContactGroupN = val
}

// GetContactRole returns the ContactRole property
func (m DcimDevicesListQueryParameters) GetContactRole() []int32 {
	return m.ContactRole
}

// SetContactRole sets the ContactRole property
func (m *DcimDevicesListQueryParameters) SetContactRole(val []int32) {
	m.ContactRole = val
}

// GetContactRoleN returns the ContactRoleN property
func (m DcimDevicesListQueryParameters) GetContactRoleN() []int32 {
	return m.ContactRoleN
}

// SetContactRoleN sets the ContactRoleN property
func (m *DcimDevicesListQueryParameters) SetContactRoleN(val []int32) {
	m.ContactRoleN = val
}

// GetCreated returns the Created property
func (m DcimDevicesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimDevicesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimDevicesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimDevicesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimDevicesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimDevicesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimDevicesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimDevicesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimDevicesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimDevicesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimDevicesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimDevicesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimDevicesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimDevicesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDeviceBays returns the DeviceBays property
func (m DcimDevicesListQueryParameters) GetDeviceBays() bool {
	return m.DeviceBays
}

// SetDeviceBays sets the DeviceBays property
func (m *DcimDevicesListQueryParameters) SetDeviceBays(val bool) {
	m.DeviceBays = val
}

// GetDeviceType returns the DeviceType property
func (m DcimDevicesListQueryParameters) GetDeviceType() []string {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *DcimDevicesListQueryParameters) SetDeviceType(val []string) {
	m.DeviceType = val
}

// GetDeviceTypeN returns the DeviceTypeN property
func (m DcimDevicesListQueryParameters) GetDeviceTypeN() []string {
	return m.DeviceTypeN
}

// SetDeviceTypeN sets the DeviceTypeN property
func (m *DcimDevicesListQueryParameters) SetDeviceTypeN(val []string) {
	m.DeviceTypeN = val
}

// GetDeviceTypeId returns the DeviceTypeId property
func (m DcimDevicesListQueryParameters) GetDeviceTypeId() []int32 {
	return m.DeviceTypeId
}

// SetDeviceTypeId sets the DeviceTypeId property
func (m *DcimDevicesListQueryParameters) SetDeviceTypeId(val []int32) {
	m.DeviceTypeId = val
}

// GetDeviceTypeIdN returns the DeviceTypeIdN property
func (m DcimDevicesListQueryParameters) GetDeviceTypeIdN() []int32 {
	return m.DeviceTypeIdN
}

// SetDeviceTypeIdN sets the DeviceTypeIdN property
func (m *DcimDevicesListQueryParameters) SetDeviceTypeIdN(val []int32) {
	m.DeviceTypeIdN = val
}

// GetFace returns the Face property
func (m DcimDevicesListQueryParameters) GetFace() string {
	return m.Face
}

// SetFace sets the Face property
func (m *DcimDevicesListQueryParameters) SetFace(val string) {
	m.Face = val
}

// GetFaceN returns the FaceN property
func (m DcimDevicesListQueryParameters) GetFaceN() string {
	return m.FaceN
}

// SetFaceN sets the FaceN property
func (m *DcimDevicesListQueryParameters) SetFaceN(val string) {
	m.FaceN = val
}

// GetHasPrimaryIp returns the HasPrimaryIp property
func (m DcimDevicesListQueryParameters) GetHasPrimaryIp() bool {
	return m.HasPrimaryIp
}

// SetHasPrimaryIp sets the HasPrimaryIp property
func (m *DcimDevicesListQueryParameters) SetHasPrimaryIp(val bool) {
	m.HasPrimaryIp = val
}

// GetId returns the Id property
func (m DcimDevicesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimDevicesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimDevicesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimDevicesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimDevicesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimDevicesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimDevicesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimDevicesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimDevicesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimDevicesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimDevicesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimDevicesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterfaces returns the Interfaces property
func (m DcimDevicesListQueryParameters) GetInterfaces() bool {
	return m.Interfaces
}

// SetInterfaces sets the Interfaces property
func (m *DcimDevicesListQueryParameters) SetInterfaces(val bool) {
	m.Interfaces = val
}

// GetIsFullDepth returns the IsFullDepth property
func (m DcimDevicesListQueryParameters) GetIsFullDepth() bool {
	return m.IsFullDepth
}

// SetIsFullDepth sets the IsFullDepth property
func (m *DcimDevicesListQueryParameters) SetIsFullDepth(val bool) {
	m.IsFullDepth = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimDevicesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimDevicesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimDevicesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimDevicesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimDevicesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimDevicesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimDevicesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimDevicesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimDevicesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimDevicesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimDevicesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimDevicesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimDevicesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimDevicesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocalContextData returns the LocalContextData property
func (m DcimDevicesListQueryParameters) GetLocalContextData() bool {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *DcimDevicesListQueryParameters) SetLocalContextData(val bool) {
	m.LocalContextData = val
}

// GetLocationId returns the LocationId property
func (m DcimDevicesListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimDevicesListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m DcimDevicesListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *DcimDevicesListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetMacAddress returns the MacAddress property
func (m DcimDevicesListQueryParameters) GetMacAddress() []string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *DcimDevicesListQueryParameters) SetMacAddress(val []string) {
	m.MacAddress = val
}

// GetMacAddressIc returns the MacAddressIc property
func (m DcimDevicesListQueryParameters) GetMacAddressIc() []string {
	return m.MacAddressIc
}

// SetMacAddressIc sets the MacAddressIc property
func (m *DcimDevicesListQueryParameters) SetMacAddressIc(val []string) {
	m.MacAddressIc = val
}

// GetMacAddressIe returns the MacAddressIe property
func (m DcimDevicesListQueryParameters) GetMacAddressIe() []string {
	return m.MacAddressIe
}

// SetMacAddressIe sets the MacAddressIe property
func (m *DcimDevicesListQueryParameters) SetMacAddressIe(val []string) {
	m.MacAddressIe = val
}

// GetMacAddressIew returns the MacAddressIew property
func (m DcimDevicesListQueryParameters) GetMacAddressIew() []string {
	return m.MacAddressIew
}

// SetMacAddressIew sets the MacAddressIew property
func (m *DcimDevicesListQueryParameters) SetMacAddressIew(val []string) {
	m.MacAddressIew = val
}

// GetMacAddressIsw returns the MacAddressIsw property
func (m DcimDevicesListQueryParameters) GetMacAddressIsw() []string {
	return m.MacAddressIsw
}

// SetMacAddressIsw sets the MacAddressIsw property
func (m *DcimDevicesListQueryParameters) SetMacAddressIsw(val []string) {
	m.MacAddressIsw = val
}

// GetMacAddressN returns the MacAddressN property
func (m DcimDevicesListQueryParameters) GetMacAddressN() []string {
	return m.MacAddressN
}

// SetMacAddressN sets the MacAddressN property
func (m *DcimDevicesListQueryParameters) SetMacAddressN(val []string) {
	m.MacAddressN = val
}

// GetMacAddressNic returns the MacAddressNic property
func (m DcimDevicesListQueryParameters) GetMacAddressNic() []string {
	return m.MacAddressNic
}

// SetMacAddressNic sets the MacAddressNic property
func (m *DcimDevicesListQueryParameters) SetMacAddressNic(val []string) {
	m.MacAddressNic = val
}

// GetMacAddressNie returns the MacAddressNie property
func (m DcimDevicesListQueryParameters) GetMacAddressNie() []string {
	return m.MacAddressNie
}

// SetMacAddressNie sets the MacAddressNie property
func (m *DcimDevicesListQueryParameters) SetMacAddressNie(val []string) {
	m.MacAddressNie = val
}

// GetMacAddressNiew returns the MacAddressNiew property
func (m DcimDevicesListQueryParameters) GetMacAddressNiew() []string {
	return m.MacAddressNiew
}

// SetMacAddressNiew sets the MacAddressNiew property
func (m *DcimDevicesListQueryParameters) SetMacAddressNiew(val []string) {
	m.MacAddressNiew = val
}

// GetMacAddressNisw returns the MacAddressNisw property
func (m DcimDevicesListQueryParameters) GetMacAddressNisw() []string {
	return m.MacAddressNisw
}

// SetMacAddressNisw sets the MacAddressNisw property
func (m *DcimDevicesListQueryParameters) SetMacAddressNisw(val []string) {
	m.MacAddressNisw = val
}

// GetManufacturer returns the Manufacturer property
func (m DcimDevicesListQueryParameters) GetManufacturer() []string {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DcimDevicesListQueryParameters) SetManufacturer(val []string) {
	m.Manufacturer = val
}

// GetManufacturerN returns the ManufacturerN property
func (m DcimDevicesListQueryParameters) GetManufacturerN() []string {
	return m.ManufacturerN
}

// SetManufacturerN sets the ManufacturerN property
func (m *DcimDevicesListQueryParameters) SetManufacturerN(val []string) {
	m.ManufacturerN = val
}

// GetManufacturerId returns the ManufacturerId property
func (m DcimDevicesListQueryParameters) GetManufacturerId() []int32 {
	return m.ManufacturerId
}

// SetManufacturerId sets the ManufacturerId property
func (m *DcimDevicesListQueryParameters) SetManufacturerId(val []int32) {
	m.ManufacturerId = val
}

// GetManufacturerIdN returns the ManufacturerIdN property
func (m DcimDevicesListQueryParameters) GetManufacturerIdN() []int32 {
	return m.ManufacturerIdN
}

// SetManufacturerIdN sets the ManufacturerIdN property
func (m *DcimDevicesListQueryParameters) SetManufacturerIdN(val []int32) {
	m.ManufacturerIdN = val
}

// GetModel returns the Model property
func (m DcimDevicesListQueryParameters) GetModel() []string {
	return m.Model
}

// SetModel sets the Model property
func (m *DcimDevicesListQueryParameters) SetModel(val []string) {
	m.Model = val
}

// GetModelN returns the ModelN property
func (m DcimDevicesListQueryParameters) GetModelN() []string {
	return m.ModelN
}

// SetModelN sets the ModelN property
func (m *DcimDevicesListQueryParameters) SetModelN(val []string) {
	m.ModelN = val
}

// GetModuleBays returns the ModuleBays property
func (m DcimDevicesListQueryParameters) GetModuleBays() bool {
	return m.ModuleBays
}

// SetModuleBays sets the ModuleBays property
func (m *DcimDevicesListQueryParameters) SetModuleBays(val bool) {
	m.ModuleBays = val
}

// GetName returns the Name property
func (m DcimDevicesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimDevicesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimDevicesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimDevicesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimDevicesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimDevicesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimDevicesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimDevicesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimDevicesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimDevicesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimDevicesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimDevicesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimDevicesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimDevicesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimDevicesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimDevicesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimDevicesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimDevicesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimDevicesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimDevicesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimDevicesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimDevicesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimDevicesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimDevicesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimDevicesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimDevicesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetParentDeviceId returns the ParentDeviceId property
func (m DcimDevicesListQueryParameters) GetParentDeviceId() []int32 {
	return m.ParentDeviceId
}

// SetParentDeviceId sets the ParentDeviceId property
func (m *DcimDevicesListQueryParameters) SetParentDeviceId(val []int32) {
	m.ParentDeviceId = val
}

// GetParentDeviceIdN returns the ParentDeviceIdN property
func (m DcimDevicesListQueryParameters) GetParentDeviceIdN() []int32 {
	return m.ParentDeviceIdN
}

// SetParentDeviceIdN sets the ParentDeviceIdN property
func (m *DcimDevicesListQueryParameters) SetParentDeviceIdN(val []int32) {
	m.ParentDeviceIdN = val
}

// GetPassThroughPorts returns the PassThroughPorts property
func (m DcimDevicesListQueryParameters) GetPassThroughPorts() bool {
	return m.PassThroughPorts
}

// SetPassThroughPorts sets the PassThroughPorts property
func (m *DcimDevicesListQueryParameters) SetPassThroughPorts(val bool) {
	m.PassThroughPorts = val
}

// GetPlatform returns the Platform property
func (m DcimDevicesListQueryParameters) GetPlatform() []string {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *DcimDevicesListQueryParameters) SetPlatform(val []string) {
	m.Platform = val
}

// GetPlatformN returns the PlatformN property
func (m DcimDevicesListQueryParameters) GetPlatformN() []string {
	return m.PlatformN
}

// SetPlatformN sets the PlatformN property
func (m *DcimDevicesListQueryParameters) SetPlatformN(val []string) {
	m.PlatformN = val
}

// GetPlatformId returns the PlatformId property
func (m DcimDevicesListQueryParameters) GetPlatformId() []*int32 {
	return m.PlatformId
}

// SetPlatformId sets the PlatformId property
func (m *DcimDevicesListQueryParameters) SetPlatformId(val []*int32) {
	m.PlatformId = val
}

// GetPlatformIdN returns the PlatformIdN property
func (m DcimDevicesListQueryParameters) GetPlatformIdN() []*int32 {
	return m.PlatformIdN
}

// SetPlatformIdN sets the PlatformIdN property
func (m *DcimDevicesListQueryParameters) SetPlatformIdN(val []*int32) {
	m.PlatformIdN = val
}

// GetPosition returns the Position property
func (m DcimDevicesListQueryParameters) GetPosition() []float64 {
	return m.Position
}

// SetPosition sets the Position property
func (m *DcimDevicesListQueryParameters) SetPosition(val []float64) {
	m.Position = val
}

// GetPositionGt returns the PositionGt property
func (m DcimDevicesListQueryParameters) GetPositionGt() []float64 {
	return m.PositionGt
}

// SetPositionGt sets the PositionGt property
func (m *DcimDevicesListQueryParameters) SetPositionGt(val []float64) {
	m.PositionGt = val
}

// GetPositionGte returns the PositionGte property
func (m DcimDevicesListQueryParameters) GetPositionGte() []float64 {
	return m.PositionGte
}

// SetPositionGte sets the PositionGte property
func (m *DcimDevicesListQueryParameters) SetPositionGte(val []float64) {
	m.PositionGte = val
}

// GetPositionLt returns the PositionLt property
func (m DcimDevicesListQueryParameters) GetPositionLt() []float64 {
	return m.PositionLt
}

// SetPositionLt sets the PositionLt property
func (m *DcimDevicesListQueryParameters) SetPositionLt(val []float64) {
	m.PositionLt = val
}

// GetPositionLte returns the PositionLte property
func (m DcimDevicesListQueryParameters) GetPositionLte() []float64 {
	return m.PositionLte
}

// SetPositionLte sets the PositionLte property
func (m *DcimDevicesListQueryParameters) SetPositionLte(val []float64) {
	m.PositionLte = val
}

// GetPositionN returns the PositionN property
func (m DcimDevicesListQueryParameters) GetPositionN() []float64 {
	return m.PositionN
}

// SetPositionN sets the PositionN property
func (m *DcimDevicesListQueryParameters) SetPositionN(val []float64) {
	m.PositionN = val
}

// GetPowerOutlets returns the PowerOutlets property
func (m DcimDevicesListQueryParameters) GetPowerOutlets() bool {
	return m.PowerOutlets
}

// SetPowerOutlets sets the PowerOutlets property
func (m *DcimDevicesListQueryParameters) SetPowerOutlets(val bool) {
	m.PowerOutlets = val
}

// GetPowerPorts returns the PowerPorts property
func (m DcimDevicesListQueryParameters) GetPowerPorts() bool {
	return m.PowerPorts
}

// SetPowerPorts sets the PowerPorts property
func (m *DcimDevicesListQueryParameters) SetPowerPorts(val bool) {
	m.PowerPorts = val
}

// GetPrimaryIp4Id returns the PrimaryIp4Id property
func (m DcimDevicesListQueryParameters) GetPrimaryIp4Id() []int32 {
	return m.PrimaryIp4Id
}

// SetPrimaryIp4Id sets the PrimaryIp4Id property
func (m *DcimDevicesListQueryParameters) SetPrimaryIp4Id(val []int32) {
	m.PrimaryIp4Id = val
}

// GetPrimaryIp4IdN returns the PrimaryIp4IdN property
func (m DcimDevicesListQueryParameters) GetPrimaryIp4IdN() []int32 {
	return m.PrimaryIp4IdN
}

// SetPrimaryIp4IdN sets the PrimaryIp4IdN property
func (m *DcimDevicesListQueryParameters) SetPrimaryIp4IdN(val []int32) {
	m.PrimaryIp4IdN = val
}

// GetPrimaryIp6Id returns the PrimaryIp6Id property
func (m DcimDevicesListQueryParameters) GetPrimaryIp6Id() []int32 {
	return m.PrimaryIp6Id
}

// SetPrimaryIp6Id sets the PrimaryIp6Id property
func (m *DcimDevicesListQueryParameters) SetPrimaryIp6Id(val []int32) {
	m.PrimaryIp6Id = val
}

// GetPrimaryIp6IdN returns the PrimaryIp6IdN property
func (m DcimDevicesListQueryParameters) GetPrimaryIp6IdN() []int32 {
	return m.PrimaryIp6IdN
}

// SetPrimaryIp6IdN sets the PrimaryIp6IdN property
func (m *DcimDevicesListQueryParameters) SetPrimaryIp6IdN(val []int32) {
	m.PrimaryIp6IdN = val
}

// GetQ returns the Q property
func (m DcimDevicesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimDevicesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRackId returns the RackId property
func (m DcimDevicesListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimDevicesListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetRackIdN returns the RackIdN property
func (m DcimDevicesListQueryParameters) GetRackIdN() []int32 {
	return m.RackIdN
}

// SetRackIdN sets the RackIdN property
func (m *DcimDevicesListQueryParameters) SetRackIdN(val []int32) {
	m.RackIdN = val
}

// GetRegion returns the Region property
func (m DcimDevicesListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimDevicesListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimDevicesListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimDevicesListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimDevicesListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimDevicesListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimDevicesListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimDevicesListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetRole returns the Role property
func (m DcimDevicesListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *DcimDevicesListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m DcimDevicesListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *DcimDevicesListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m DcimDevicesListQueryParameters) GetRoleId() []int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *DcimDevicesListQueryParameters) SetRoleId(val []int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m DcimDevicesListQueryParameters) GetRoleIdN() []int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *DcimDevicesListQueryParameters) SetRoleIdN(val []int32) {
	m.RoleIdN = val
}

// GetSerial returns the Serial property
func (m DcimDevicesListQueryParameters) GetSerial() []string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *DcimDevicesListQueryParameters) SetSerial(val []string) {
	m.Serial = val
}

// GetSerialEmpty returns the SerialEmpty property
func (m DcimDevicesListQueryParameters) GetSerialEmpty() []string {
	return m.SerialEmpty
}

// SetSerialEmpty sets the SerialEmpty property
func (m *DcimDevicesListQueryParameters) SetSerialEmpty(val []string) {
	m.SerialEmpty = val
}

// GetSerialIc returns the SerialIc property
func (m DcimDevicesListQueryParameters) GetSerialIc() []string {
	return m.SerialIc
}

// SetSerialIc sets the SerialIc property
func (m *DcimDevicesListQueryParameters) SetSerialIc(val []string) {
	m.SerialIc = val
}

// GetSerialIe returns the SerialIe property
func (m DcimDevicesListQueryParameters) GetSerialIe() []string {
	return m.SerialIe
}

// SetSerialIe sets the SerialIe property
func (m *DcimDevicesListQueryParameters) SetSerialIe(val []string) {
	m.SerialIe = val
}

// GetSerialIew returns the SerialIew property
func (m DcimDevicesListQueryParameters) GetSerialIew() []string {
	return m.SerialIew
}

// SetSerialIew sets the SerialIew property
func (m *DcimDevicesListQueryParameters) SetSerialIew(val []string) {
	m.SerialIew = val
}

// GetSerialIsw returns the SerialIsw property
func (m DcimDevicesListQueryParameters) GetSerialIsw() []string {
	return m.SerialIsw
}

// SetSerialIsw sets the SerialIsw property
func (m *DcimDevicesListQueryParameters) SetSerialIsw(val []string) {
	m.SerialIsw = val
}

// GetSerialN returns the SerialN property
func (m DcimDevicesListQueryParameters) GetSerialN() []string {
	return m.SerialN
}

// SetSerialN sets the SerialN property
func (m *DcimDevicesListQueryParameters) SetSerialN(val []string) {
	m.SerialN = val
}

// GetSerialNic returns the SerialNic property
func (m DcimDevicesListQueryParameters) GetSerialNic() []string {
	return m.SerialNic
}

// SetSerialNic sets the SerialNic property
func (m *DcimDevicesListQueryParameters) SetSerialNic(val []string) {
	m.SerialNic = val
}

// GetSerialNie returns the SerialNie property
func (m DcimDevicesListQueryParameters) GetSerialNie() []string {
	return m.SerialNie
}

// SetSerialNie sets the SerialNie property
func (m *DcimDevicesListQueryParameters) SetSerialNie(val []string) {
	m.SerialNie = val
}

// GetSerialNiew returns the SerialNiew property
func (m DcimDevicesListQueryParameters) GetSerialNiew() []string {
	return m.SerialNiew
}

// SetSerialNiew sets the SerialNiew property
func (m *DcimDevicesListQueryParameters) SetSerialNiew(val []string) {
	m.SerialNiew = val
}

// GetSerialNisw returns the SerialNisw property
func (m DcimDevicesListQueryParameters) GetSerialNisw() []string {
	return m.SerialNisw
}

// SetSerialNisw sets the SerialNisw property
func (m *DcimDevicesListQueryParameters) SetSerialNisw(val []string) {
	m.SerialNisw = val
}

// GetSite returns the Site property
func (m DcimDevicesListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimDevicesListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimDevicesListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimDevicesListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimDevicesListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimDevicesListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimDevicesListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimDevicesListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimDevicesListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimDevicesListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimDevicesListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimDevicesListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimDevicesListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimDevicesListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimDevicesListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimDevicesListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetStatus returns the Status property
func (m DcimDevicesListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *DcimDevicesListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m DcimDevicesListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *DcimDevicesListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m DcimDevicesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimDevicesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimDevicesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimDevicesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m DcimDevicesListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DcimDevicesListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m DcimDevicesListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *DcimDevicesListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m DcimDevicesListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *DcimDevicesListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m DcimDevicesListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *DcimDevicesListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m DcimDevicesListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *DcimDevicesListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m DcimDevicesListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *DcimDevicesListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m DcimDevicesListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *DcimDevicesListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m DcimDevicesListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *DcimDevicesListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimDevicesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimDevicesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVcPosition returns the VcPosition property
func (m DcimDevicesListQueryParameters) GetVcPosition() []int32 {
	return m.VcPosition
}

// SetVcPosition sets the VcPosition property
func (m *DcimDevicesListQueryParameters) SetVcPosition(val []int32) {
	m.VcPosition = val
}

// GetVcPositionGt returns the VcPositionGt property
func (m DcimDevicesListQueryParameters) GetVcPositionGt() []int32 {
	return m.VcPositionGt
}

// SetVcPositionGt sets the VcPositionGt property
func (m *DcimDevicesListQueryParameters) SetVcPositionGt(val []int32) {
	m.VcPositionGt = val
}

// GetVcPositionGte returns the VcPositionGte property
func (m DcimDevicesListQueryParameters) GetVcPositionGte() []int32 {
	return m.VcPositionGte
}

// SetVcPositionGte sets the VcPositionGte property
func (m *DcimDevicesListQueryParameters) SetVcPositionGte(val []int32) {
	m.VcPositionGte = val
}

// GetVcPositionLt returns the VcPositionLt property
func (m DcimDevicesListQueryParameters) GetVcPositionLt() []int32 {
	return m.VcPositionLt
}

// SetVcPositionLt sets the VcPositionLt property
func (m *DcimDevicesListQueryParameters) SetVcPositionLt(val []int32) {
	m.VcPositionLt = val
}

// GetVcPositionLte returns the VcPositionLte property
func (m DcimDevicesListQueryParameters) GetVcPositionLte() []int32 {
	return m.VcPositionLte
}

// SetVcPositionLte sets the VcPositionLte property
func (m *DcimDevicesListQueryParameters) SetVcPositionLte(val []int32) {
	m.VcPositionLte = val
}

// GetVcPositionN returns the VcPositionN property
func (m DcimDevicesListQueryParameters) GetVcPositionN() []int32 {
	return m.VcPositionN
}

// SetVcPositionN sets the VcPositionN property
func (m *DcimDevicesListQueryParameters) SetVcPositionN(val []int32) {
	m.VcPositionN = val
}

// GetVcPriority returns the VcPriority property
func (m DcimDevicesListQueryParameters) GetVcPriority() []int32 {
	return m.VcPriority
}

// SetVcPriority sets the VcPriority property
func (m *DcimDevicesListQueryParameters) SetVcPriority(val []int32) {
	m.VcPriority = val
}

// GetVcPriorityGt returns the VcPriorityGt property
func (m DcimDevicesListQueryParameters) GetVcPriorityGt() []int32 {
	return m.VcPriorityGt
}

// SetVcPriorityGt sets the VcPriorityGt property
func (m *DcimDevicesListQueryParameters) SetVcPriorityGt(val []int32) {
	m.VcPriorityGt = val
}

// GetVcPriorityGte returns the VcPriorityGte property
func (m DcimDevicesListQueryParameters) GetVcPriorityGte() []int32 {
	return m.VcPriorityGte
}

// SetVcPriorityGte sets the VcPriorityGte property
func (m *DcimDevicesListQueryParameters) SetVcPriorityGte(val []int32) {
	m.VcPriorityGte = val
}

// GetVcPriorityLt returns the VcPriorityLt property
func (m DcimDevicesListQueryParameters) GetVcPriorityLt() []int32 {
	return m.VcPriorityLt
}

// SetVcPriorityLt sets the VcPriorityLt property
func (m *DcimDevicesListQueryParameters) SetVcPriorityLt(val []int32) {
	m.VcPriorityLt = val
}

// GetVcPriorityLte returns the VcPriorityLte property
func (m DcimDevicesListQueryParameters) GetVcPriorityLte() []int32 {
	return m.VcPriorityLte
}

// SetVcPriorityLte sets the VcPriorityLte property
func (m *DcimDevicesListQueryParameters) SetVcPriorityLte(val []int32) {
	m.VcPriorityLte = val
}

// GetVcPriorityN returns the VcPriorityN property
func (m DcimDevicesListQueryParameters) GetVcPriorityN() []int32 {
	return m.VcPriorityN
}

// SetVcPriorityN sets the VcPriorityN property
func (m *DcimDevicesListQueryParameters) SetVcPriorityN(val []int32) {
	m.VcPriorityN = val
}

// GetVirtualChassisId returns the VirtualChassisId property
func (m DcimDevicesListQueryParameters) GetVirtualChassisId() []int32 {
	return m.VirtualChassisId
}

// SetVirtualChassisId sets the VirtualChassisId property
func (m *DcimDevicesListQueryParameters) SetVirtualChassisId(val []int32) {
	m.VirtualChassisId = val
}

// GetVirtualChassisIdN returns the VirtualChassisIdN property
func (m DcimDevicesListQueryParameters) GetVirtualChassisIdN() []int32 {
	return m.VirtualChassisIdN
}

// SetVirtualChassisIdN sets the VirtualChassisIdN property
func (m *DcimDevicesListQueryParameters) SetVirtualChassisIdN(val []int32) {
	m.VirtualChassisIdN = val
}

// GetVirtualChassisMember returns the VirtualChassisMember property
func (m DcimDevicesListQueryParameters) GetVirtualChassisMember() bool {
	return m.VirtualChassisMember
}

// SetVirtualChassisMember sets the VirtualChassisMember property
func (m *DcimDevicesListQueryParameters) SetVirtualChassisMember(val bool) {
	m.VirtualChassisMember = val
}
