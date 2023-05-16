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

// DcimInventoryItemsListQueryParameters is an object.
type DcimInventoryItemsListQueryParameters struct {
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
	// ComponentId:
	ComponentId []int32 `json:"component_id,omitempty" mapstructure:"component_id,omitempty"`
	// ComponentIdGt:
	ComponentIdGt []int32 `json:"component_id__gt,omitempty" mapstructure:"component_id__gt,omitempty"`
	// ComponentIdGte:
	ComponentIdGte []int32 `json:"component_id__gte,omitempty" mapstructure:"component_id__gte,omitempty"`
	// ComponentIdLt:
	ComponentIdLt []int32 `json:"component_id__lt,omitempty" mapstructure:"component_id__lt,omitempty"`
	// ComponentIdLte:
	ComponentIdLte []int32 `json:"component_id__lte,omitempty" mapstructure:"component_id__lte,omitempty"`
	// ComponentIdN:
	ComponentIdN []int32 `json:"component_id__n,omitempty" mapstructure:"component_id__n,omitempty"`
	// ComponentType:
	ComponentType string `json:"component_type,omitempty" mapstructure:"component_type,omitempty"`
	// ComponentTypeN:
	ComponentTypeN string `json:"component_type__n,omitempty" mapstructure:"component_type__n,omitempty"`
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
	// Device: Device (name)
	Device []*string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceN: Device (name)
	DeviceN []*string `json:"device__n,omitempty" mapstructure:"device__n,omitempty"`
	// DeviceId: Device (ID)
	DeviceId []int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
	// DeviceIdN: Device (ID)
	DeviceIdN []int32 `json:"device_id__n,omitempty" mapstructure:"device_id__n,omitempty"`
	// Discovered:
	Discovered bool `json:"discovered,omitempty" mapstructure:"discovered,omitempty"`
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
	// Label:
	Label []string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LabelEmpty:
	LabelEmpty []string `json:"label__empty,omitempty" mapstructure:"label__empty,omitempty"`
	// LabelIc:
	LabelIc []string `json:"label__ic,omitempty" mapstructure:"label__ic,omitempty"`
	// LabelIe:
	LabelIe []string `json:"label__ie,omitempty" mapstructure:"label__ie,omitempty"`
	// LabelIew:
	LabelIew []string `json:"label__iew,omitempty" mapstructure:"label__iew,omitempty"`
	// LabelIsw:
	LabelIsw []string `json:"label__isw,omitempty" mapstructure:"label__isw,omitempty"`
	// LabelN:
	LabelN []string `json:"label__n,omitempty" mapstructure:"label__n,omitempty"`
	// LabelNic:
	LabelNic []string `json:"label__nic,omitempty" mapstructure:"label__nic,omitempty"`
	// LabelNie:
	LabelNie []string `json:"label__nie,omitempty" mapstructure:"label__nie,omitempty"`
	// LabelNiew:
	LabelNiew []string `json:"label__niew,omitempty" mapstructure:"label__niew,omitempty"`
	// LabelNisw:
	LabelNisw []string `json:"label__nisw,omitempty" mapstructure:"label__nisw,omitempty"`
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
	// Location: Location (slug)
	Location []string `json:"location,omitempty" mapstructure:"location,omitempty"`
	// LocationN: Location (slug)
	LocationN []string `json:"location__n,omitempty" mapstructure:"location__n,omitempty"`
	// LocationId: Location (ID)
	LocationId []int32 `json:"location_id,omitempty" mapstructure:"location_id,omitempty"`
	// LocationIdN: Location (ID)
	LocationIdN []int32 `json:"location_id__n,omitempty" mapstructure:"location_id__n,omitempty"`
	// Manufacturer: Manufacturer (slug)
	Manufacturer []string `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// ManufacturerN: Manufacturer (slug)
	ManufacturerN []string `json:"manufacturer__n,omitempty" mapstructure:"manufacturer__n,omitempty"`
	// ManufacturerId: Manufacturer (ID)
	ManufacturerId []*int32 `json:"manufacturer_id,omitempty" mapstructure:"manufacturer_id,omitempty"`
	// ManufacturerIdN: Manufacturer (ID)
	ManufacturerIdN []*int32 `json:"manufacturer_id__n,omitempty" mapstructure:"manufacturer_id__n,omitempty"`
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
	// ParentId: Parent inventory item (ID)
	ParentId []*int32 `json:"parent_id,omitempty" mapstructure:"parent_id,omitempty"`
	// ParentIdN: Parent inventory item (ID)
	ParentIdN []*int32 `json:"parent_id__n,omitempty" mapstructure:"parent_id__n,omitempty"`
	// PartId:
	PartId []string `json:"part_id,omitempty" mapstructure:"part_id,omitempty"`
	// PartIdEmpty:
	PartIdEmpty []string `json:"part_id__empty,omitempty" mapstructure:"part_id__empty,omitempty"`
	// PartIdIc:
	PartIdIc []string `json:"part_id__ic,omitempty" mapstructure:"part_id__ic,omitempty"`
	// PartIdIe:
	PartIdIe []string `json:"part_id__ie,omitempty" mapstructure:"part_id__ie,omitempty"`
	// PartIdIew:
	PartIdIew []string `json:"part_id__iew,omitempty" mapstructure:"part_id__iew,omitempty"`
	// PartIdIsw:
	PartIdIsw []string `json:"part_id__isw,omitempty" mapstructure:"part_id__isw,omitempty"`
	// PartIdN:
	PartIdN []string `json:"part_id__n,omitempty" mapstructure:"part_id__n,omitempty"`
	// PartIdNic:
	PartIdNic []string `json:"part_id__nic,omitempty" mapstructure:"part_id__nic,omitempty"`
	// PartIdNie:
	PartIdNie []string `json:"part_id__nie,omitempty" mapstructure:"part_id__nie,omitempty"`
	// PartIdNiew:
	PartIdNiew []string `json:"part_id__niew,omitempty" mapstructure:"part_id__niew,omitempty"`
	// PartIdNisw:
	PartIdNisw []string `json:"part_id__nisw,omitempty" mapstructure:"part_id__nisw,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Rack: Rack (name)
	Rack []string `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// RackN: Rack (name)
	RackN []string `json:"rack__n,omitempty" mapstructure:"rack__n,omitempty"`
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
	RoleId []*int32 `json:"role_id,omitempty" mapstructure:"role_id,omitempty"`
	// RoleIdN: Role (ID)
	RoleIdN []*int32 `json:"role_id__n,omitempty" mapstructure:"role_id__n,omitempty"`
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
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// VirtualChassis: Virtual Chassis
	VirtualChassis []string `json:"virtual_chassis,omitempty" mapstructure:"virtual_chassis,omitempty"`
	// VirtualChassisN: Virtual Chassis
	VirtualChassisN []string `json:"virtual_chassis__n,omitempty" mapstructure:"virtual_chassis__n,omitempty"`
	// VirtualChassisId: Virtual Chassis (ID)
	VirtualChassisId []int32 `json:"virtual_chassis_id,omitempty" mapstructure:"virtual_chassis_id,omitempty"`
	// VirtualChassisIdN: Virtual Chassis (ID)
	VirtualChassisIdN []int32 `json:"virtual_chassis_id__n,omitempty" mapstructure:"virtual_chassis_id__n,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimInventoryItemsListQueryParameters) Validate() error {
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
		"componentId": validation.Validate(
			m.ComponentId,
		),
		"componentIdGt": validation.Validate(
			m.ComponentIdGt,
		),
		"componentIdGte": validation.Validate(
			m.ComponentIdGte,
		),
		"componentIdLt": validation.Validate(
			m.ComponentIdLt,
		),
		"componentIdLte": validation.Validate(
			m.ComponentIdLte,
		),
		"componentIdN": validation.Validate(
			m.ComponentIdN,
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
		"device": validation.Validate(
			m.Device,
		),
		"deviceN": validation.Validate(
			m.DeviceN,
		),
		"deviceId": validation.Validate(
			m.DeviceId,
		),
		"deviceIdN": validation.Validate(
			m.DeviceIdN,
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
		"label": validation.Validate(
			m.Label,
		),
		"labelEmpty": validation.Validate(
			m.LabelEmpty,
		),
		"labelIc": validation.Validate(
			m.LabelIc,
		),
		"labelIe": validation.Validate(
			m.LabelIe,
		),
		"labelIew": validation.Validate(
			m.LabelIew,
		),
		"labelIsw": validation.Validate(
			m.LabelIsw,
		),
		"labelN": validation.Validate(
			m.LabelN,
		),
		"labelNic": validation.Validate(
			m.LabelNic,
		),
		"labelNie": validation.Validate(
			m.LabelNie,
		),
		"labelNiew": validation.Validate(
			m.LabelNiew,
		),
		"labelNisw": validation.Validate(
			m.LabelNisw,
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
		"location": validation.Validate(
			m.Location,
		),
		"locationN": validation.Validate(
			m.LocationN,
		),
		"locationId": validation.Validate(
			m.LocationId,
		),
		"locationIdN": validation.Validate(
			m.LocationIdN,
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
		"parentId": validation.Validate(
			m.ParentId,
		),
		"parentIdN": validation.Validate(
			m.ParentIdN,
		),
		"partId": validation.Validate(
			m.PartId,
		),
		"partIdEmpty": validation.Validate(
			m.PartIdEmpty,
		),
		"partIdIc": validation.Validate(
			m.PartIdIc,
		),
		"partIdIe": validation.Validate(
			m.PartIdIe,
		),
		"partIdIew": validation.Validate(
			m.PartIdIew,
		),
		"partIdIsw": validation.Validate(
			m.PartIdIsw,
		),
		"partIdN": validation.Validate(
			m.PartIdN,
		),
		"partIdNic": validation.Validate(
			m.PartIdNic,
		),
		"partIdNie": validation.Validate(
			m.PartIdNie,
		),
		"partIdNiew": validation.Validate(
			m.PartIdNiew,
		),
		"partIdNisw": validation.Validate(
			m.PartIdNisw,
		),
		"rack": validation.Validate(
			m.Rack,
		),
		"rackN": validation.Validate(
			m.RackN,
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
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"virtualChassis": validation.Validate(
			m.VirtualChassis,
		),
		"virtualChassisN": validation.Validate(
			m.VirtualChassisN,
		),
		"virtualChassisId": validation.Validate(
			m.VirtualChassisId,
		),
		"virtualChassisIdN": validation.Validate(
			m.VirtualChassisIdN,
		),
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m DcimInventoryItemsListQueryParameters) GetAssetTag() []string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTag(val []string) {
	m.AssetTag = val
}

// GetAssetTagEmpty returns the AssetTagEmpty property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagEmpty() []string {
	return m.AssetTagEmpty
}

// SetAssetTagEmpty sets the AssetTagEmpty property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagEmpty(val []string) {
	m.AssetTagEmpty = val
}

// GetAssetTagIc returns the AssetTagIc property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagIc() []string {
	return m.AssetTagIc
}

// SetAssetTagIc sets the AssetTagIc property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagIc(val []string) {
	m.AssetTagIc = val
}

// GetAssetTagIe returns the AssetTagIe property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagIe() []string {
	return m.AssetTagIe
}

// SetAssetTagIe sets the AssetTagIe property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagIe(val []string) {
	m.AssetTagIe = val
}

// GetAssetTagIew returns the AssetTagIew property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagIew() []string {
	return m.AssetTagIew
}

// SetAssetTagIew sets the AssetTagIew property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagIew(val []string) {
	m.AssetTagIew = val
}

// GetAssetTagIsw returns the AssetTagIsw property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagIsw() []string {
	return m.AssetTagIsw
}

// SetAssetTagIsw sets the AssetTagIsw property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagIsw(val []string) {
	m.AssetTagIsw = val
}

// GetAssetTagN returns the AssetTagN property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagN() []string {
	return m.AssetTagN
}

// SetAssetTagN sets the AssetTagN property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagN(val []string) {
	m.AssetTagN = val
}

// GetAssetTagNic returns the AssetTagNic property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagNic() []string {
	return m.AssetTagNic
}

// SetAssetTagNic sets the AssetTagNic property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagNic(val []string) {
	m.AssetTagNic = val
}

// GetAssetTagNie returns the AssetTagNie property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagNie() []string {
	return m.AssetTagNie
}

// SetAssetTagNie sets the AssetTagNie property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagNie(val []string) {
	m.AssetTagNie = val
}

// GetAssetTagNiew returns the AssetTagNiew property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagNiew() []string {
	return m.AssetTagNiew
}

// SetAssetTagNiew sets the AssetTagNiew property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagNiew(val []string) {
	m.AssetTagNiew = val
}

// GetAssetTagNisw returns the AssetTagNisw property
func (m DcimInventoryItemsListQueryParameters) GetAssetTagNisw() []string {
	return m.AssetTagNisw
}

// SetAssetTagNisw sets the AssetTagNisw property
func (m *DcimInventoryItemsListQueryParameters) SetAssetTagNisw(val []string) {
	m.AssetTagNisw = val
}

// GetComponentId returns the ComponentId property
func (m DcimInventoryItemsListQueryParameters) GetComponentId() []int32 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *DcimInventoryItemsListQueryParameters) SetComponentId(val []int32) {
	m.ComponentId = val
}

// GetComponentIdGt returns the ComponentIdGt property
func (m DcimInventoryItemsListQueryParameters) GetComponentIdGt() []int32 {
	return m.ComponentIdGt
}

// SetComponentIdGt sets the ComponentIdGt property
func (m *DcimInventoryItemsListQueryParameters) SetComponentIdGt(val []int32) {
	m.ComponentIdGt = val
}

// GetComponentIdGte returns the ComponentIdGte property
func (m DcimInventoryItemsListQueryParameters) GetComponentIdGte() []int32 {
	return m.ComponentIdGte
}

// SetComponentIdGte sets the ComponentIdGte property
func (m *DcimInventoryItemsListQueryParameters) SetComponentIdGte(val []int32) {
	m.ComponentIdGte = val
}

// GetComponentIdLt returns the ComponentIdLt property
func (m DcimInventoryItemsListQueryParameters) GetComponentIdLt() []int32 {
	return m.ComponentIdLt
}

// SetComponentIdLt sets the ComponentIdLt property
func (m *DcimInventoryItemsListQueryParameters) SetComponentIdLt(val []int32) {
	m.ComponentIdLt = val
}

// GetComponentIdLte returns the ComponentIdLte property
func (m DcimInventoryItemsListQueryParameters) GetComponentIdLte() []int32 {
	return m.ComponentIdLte
}

// SetComponentIdLte sets the ComponentIdLte property
func (m *DcimInventoryItemsListQueryParameters) SetComponentIdLte(val []int32) {
	m.ComponentIdLte = val
}

// GetComponentIdN returns the ComponentIdN property
func (m DcimInventoryItemsListQueryParameters) GetComponentIdN() []int32 {
	return m.ComponentIdN
}

// SetComponentIdN sets the ComponentIdN property
func (m *DcimInventoryItemsListQueryParameters) SetComponentIdN(val []int32) {
	m.ComponentIdN = val
}

// GetComponentType returns the ComponentType property
func (m DcimInventoryItemsListQueryParameters) GetComponentType() string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *DcimInventoryItemsListQueryParameters) SetComponentType(val string) {
	m.ComponentType = val
}

// GetComponentTypeN returns the ComponentTypeN property
func (m DcimInventoryItemsListQueryParameters) GetComponentTypeN() string {
	return m.ComponentTypeN
}

// SetComponentTypeN sets the ComponentTypeN property
func (m *DcimInventoryItemsListQueryParameters) SetComponentTypeN(val string) {
	m.ComponentTypeN = val
}

// GetCreated returns the Created property
func (m DcimInventoryItemsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimInventoryItemsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimInventoryItemsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimInventoryItemsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimInventoryItemsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimInventoryItemsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimInventoryItemsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimInventoryItemsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimInventoryItemsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimInventoryItemsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimInventoryItemsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimInventoryItemsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimInventoryItemsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimInventoryItemsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevice returns the Device property
func (m DcimInventoryItemsListQueryParameters) GetDevice() []*string {
	return m.Device
}

// SetDevice sets the Device property
func (m *DcimInventoryItemsListQueryParameters) SetDevice(val []*string) {
	m.Device = val
}

// GetDeviceN returns the DeviceN property
func (m DcimInventoryItemsListQueryParameters) GetDeviceN() []*string {
	return m.DeviceN
}

// SetDeviceN sets the DeviceN property
func (m *DcimInventoryItemsListQueryParameters) SetDeviceN(val []*string) {
	m.DeviceN = val
}

// GetDeviceId returns the DeviceId property
func (m DcimInventoryItemsListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *DcimInventoryItemsListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetDeviceIdN returns the DeviceIdN property
func (m DcimInventoryItemsListQueryParameters) GetDeviceIdN() []int32 {
	return m.DeviceIdN
}

// SetDeviceIdN sets the DeviceIdN property
func (m *DcimInventoryItemsListQueryParameters) SetDeviceIdN(val []int32) {
	m.DeviceIdN = val
}

// GetDiscovered returns the Discovered property
func (m DcimInventoryItemsListQueryParameters) GetDiscovered() bool {
	return m.Discovered
}

// SetDiscovered sets the Discovered property
func (m *DcimInventoryItemsListQueryParameters) SetDiscovered(val bool) {
	m.Discovered = val
}

// GetId returns the Id property
func (m DcimInventoryItemsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimInventoryItemsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimInventoryItemsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimInventoryItemsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimInventoryItemsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimInventoryItemsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimInventoryItemsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimInventoryItemsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimInventoryItemsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimInventoryItemsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimInventoryItemsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimInventoryItemsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLabel returns the Label property
func (m DcimInventoryItemsListQueryParameters) GetLabel() []string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DcimInventoryItemsListQueryParameters) SetLabel(val []string) {
	m.Label = val
}

// GetLabelEmpty returns the LabelEmpty property
func (m DcimInventoryItemsListQueryParameters) GetLabelEmpty() []string {
	return m.LabelEmpty
}

// SetLabelEmpty sets the LabelEmpty property
func (m *DcimInventoryItemsListQueryParameters) SetLabelEmpty(val []string) {
	m.LabelEmpty = val
}

// GetLabelIc returns the LabelIc property
func (m DcimInventoryItemsListQueryParameters) GetLabelIc() []string {
	return m.LabelIc
}

// SetLabelIc sets the LabelIc property
func (m *DcimInventoryItemsListQueryParameters) SetLabelIc(val []string) {
	m.LabelIc = val
}

// GetLabelIe returns the LabelIe property
func (m DcimInventoryItemsListQueryParameters) GetLabelIe() []string {
	return m.LabelIe
}

// SetLabelIe sets the LabelIe property
func (m *DcimInventoryItemsListQueryParameters) SetLabelIe(val []string) {
	m.LabelIe = val
}

// GetLabelIew returns the LabelIew property
func (m DcimInventoryItemsListQueryParameters) GetLabelIew() []string {
	return m.LabelIew
}

// SetLabelIew sets the LabelIew property
func (m *DcimInventoryItemsListQueryParameters) SetLabelIew(val []string) {
	m.LabelIew = val
}

// GetLabelIsw returns the LabelIsw property
func (m DcimInventoryItemsListQueryParameters) GetLabelIsw() []string {
	return m.LabelIsw
}

// SetLabelIsw sets the LabelIsw property
func (m *DcimInventoryItemsListQueryParameters) SetLabelIsw(val []string) {
	m.LabelIsw = val
}

// GetLabelN returns the LabelN property
func (m DcimInventoryItemsListQueryParameters) GetLabelN() []string {
	return m.LabelN
}

// SetLabelN sets the LabelN property
func (m *DcimInventoryItemsListQueryParameters) SetLabelN(val []string) {
	m.LabelN = val
}

// GetLabelNic returns the LabelNic property
func (m DcimInventoryItemsListQueryParameters) GetLabelNic() []string {
	return m.LabelNic
}

// SetLabelNic sets the LabelNic property
func (m *DcimInventoryItemsListQueryParameters) SetLabelNic(val []string) {
	m.LabelNic = val
}

// GetLabelNie returns the LabelNie property
func (m DcimInventoryItemsListQueryParameters) GetLabelNie() []string {
	return m.LabelNie
}

// SetLabelNie sets the LabelNie property
func (m *DcimInventoryItemsListQueryParameters) SetLabelNie(val []string) {
	m.LabelNie = val
}

// GetLabelNiew returns the LabelNiew property
func (m DcimInventoryItemsListQueryParameters) GetLabelNiew() []string {
	return m.LabelNiew
}

// SetLabelNiew sets the LabelNiew property
func (m *DcimInventoryItemsListQueryParameters) SetLabelNiew(val []string) {
	m.LabelNiew = val
}

// GetLabelNisw returns the LabelNisw property
func (m DcimInventoryItemsListQueryParameters) GetLabelNisw() []string {
	return m.LabelNisw
}

// SetLabelNisw sets the LabelNisw property
func (m *DcimInventoryItemsListQueryParameters) SetLabelNisw(val []string) {
	m.LabelNisw = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimInventoryItemsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimInventoryItemsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimInventoryItemsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimInventoryItemsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimInventoryItemsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimInventoryItemsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimInventoryItemsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimInventoryItemsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimInventoryItemsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimInventoryItemsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimInventoryItemsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimInventoryItemsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimInventoryItemsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimInventoryItemsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m DcimInventoryItemsListQueryParameters) GetLocation() []string {
	return m.Location
}

// SetLocation sets the Location property
func (m *DcimInventoryItemsListQueryParameters) SetLocation(val []string) {
	m.Location = val
}

// GetLocationN returns the LocationN property
func (m DcimInventoryItemsListQueryParameters) GetLocationN() []string {
	return m.LocationN
}

// SetLocationN sets the LocationN property
func (m *DcimInventoryItemsListQueryParameters) SetLocationN(val []string) {
	m.LocationN = val
}

// GetLocationId returns the LocationId property
func (m DcimInventoryItemsListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimInventoryItemsListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m DcimInventoryItemsListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *DcimInventoryItemsListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetManufacturer returns the Manufacturer property
func (m DcimInventoryItemsListQueryParameters) GetManufacturer() []string {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DcimInventoryItemsListQueryParameters) SetManufacturer(val []string) {
	m.Manufacturer = val
}

// GetManufacturerN returns the ManufacturerN property
func (m DcimInventoryItemsListQueryParameters) GetManufacturerN() []string {
	return m.ManufacturerN
}

// SetManufacturerN sets the ManufacturerN property
func (m *DcimInventoryItemsListQueryParameters) SetManufacturerN(val []string) {
	m.ManufacturerN = val
}

// GetManufacturerId returns the ManufacturerId property
func (m DcimInventoryItemsListQueryParameters) GetManufacturerId() []*int32 {
	return m.ManufacturerId
}

// SetManufacturerId sets the ManufacturerId property
func (m *DcimInventoryItemsListQueryParameters) SetManufacturerId(val []*int32) {
	m.ManufacturerId = val
}

// GetManufacturerIdN returns the ManufacturerIdN property
func (m DcimInventoryItemsListQueryParameters) GetManufacturerIdN() []*int32 {
	return m.ManufacturerIdN
}

// SetManufacturerIdN sets the ManufacturerIdN property
func (m *DcimInventoryItemsListQueryParameters) SetManufacturerIdN(val []*int32) {
	m.ManufacturerIdN = val
}

// GetName returns the Name property
func (m DcimInventoryItemsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimInventoryItemsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimInventoryItemsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimInventoryItemsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimInventoryItemsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimInventoryItemsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimInventoryItemsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimInventoryItemsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimInventoryItemsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimInventoryItemsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimInventoryItemsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimInventoryItemsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimInventoryItemsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimInventoryItemsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimInventoryItemsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimInventoryItemsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimInventoryItemsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimInventoryItemsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimInventoryItemsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimInventoryItemsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimInventoryItemsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimInventoryItemsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimInventoryItemsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimInventoryItemsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimInventoryItemsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimInventoryItemsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetParentId returns the ParentId property
func (m DcimInventoryItemsListQueryParameters) GetParentId() []*int32 {
	return m.ParentId
}

// SetParentId sets the ParentId property
func (m *DcimInventoryItemsListQueryParameters) SetParentId(val []*int32) {
	m.ParentId = val
}

// GetParentIdN returns the ParentIdN property
func (m DcimInventoryItemsListQueryParameters) GetParentIdN() []*int32 {
	return m.ParentIdN
}

// SetParentIdN sets the ParentIdN property
func (m *DcimInventoryItemsListQueryParameters) SetParentIdN(val []*int32) {
	m.ParentIdN = val
}

// GetPartId returns the PartId property
func (m DcimInventoryItemsListQueryParameters) GetPartId() []string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *DcimInventoryItemsListQueryParameters) SetPartId(val []string) {
	m.PartId = val
}

// GetPartIdEmpty returns the PartIdEmpty property
func (m DcimInventoryItemsListQueryParameters) GetPartIdEmpty() []string {
	return m.PartIdEmpty
}

// SetPartIdEmpty sets the PartIdEmpty property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdEmpty(val []string) {
	m.PartIdEmpty = val
}

// GetPartIdIc returns the PartIdIc property
func (m DcimInventoryItemsListQueryParameters) GetPartIdIc() []string {
	return m.PartIdIc
}

// SetPartIdIc sets the PartIdIc property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdIc(val []string) {
	m.PartIdIc = val
}

// GetPartIdIe returns the PartIdIe property
func (m DcimInventoryItemsListQueryParameters) GetPartIdIe() []string {
	return m.PartIdIe
}

// SetPartIdIe sets the PartIdIe property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdIe(val []string) {
	m.PartIdIe = val
}

// GetPartIdIew returns the PartIdIew property
func (m DcimInventoryItemsListQueryParameters) GetPartIdIew() []string {
	return m.PartIdIew
}

// SetPartIdIew sets the PartIdIew property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdIew(val []string) {
	m.PartIdIew = val
}

// GetPartIdIsw returns the PartIdIsw property
func (m DcimInventoryItemsListQueryParameters) GetPartIdIsw() []string {
	return m.PartIdIsw
}

// SetPartIdIsw sets the PartIdIsw property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdIsw(val []string) {
	m.PartIdIsw = val
}

// GetPartIdN returns the PartIdN property
func (m DcimInventoryItemsListQueryParameters) GetPartIdN() []string {
	return m.PartIdN
}

// SetPartIdN sets the PartIdN property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdN(val []string) {
	m.PartIdN = val
}

// GetPartIdNic returns the PartIdNic property
func (m DcimInventoryItemsListQueryParameters) GetPartIdNic() []string {
	return m.PartIdNic
}

// SetPartIdNic sets the PartIdNic property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdNic(val []string) {
	m.PartIdNic = val
}

// GetPartIdNie returns the PartIdNie property
func (m DcimInventoryItemsListQueryParameters) GetPartIdNie() []string {
	return m.PartIdNie
}

// SetPartIdNie sets the PartIdNie property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdNie(val []string) {
	m.PartIdNie = val
}

// GetPartIdNiew returns the PartIdNiew property
func (m DcimInventoryItemsListQueryParameters) GetPartIdNiew() []string {
	return m.PartIdNiew
}

// SetPartIdNiew sets the PartIdNiew property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdNiew(val []string) {
	m.PartIdNiew = val
}

// GetPartIdNisw returns the PartIdNisw property
func (m DcimInventoryItemsListQueryParameters) GetPartIdNisw() []string {
	return m.PartIdNisw
}

// SetPartIdNisw sets the PartIdNisw property
func (m *DcimInventoryItemsListQueryParameters) SetPartIdNisw(val []string) {
	m.PartIdNisw = val
}

// GetQ returns the Q property
func (m DcimInventoryItemsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimInventoryItemsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRack returns the Rack property
func (m DcimInventoryItemsListQueryParameters) GetRack() []string {
	return m.Rack
}

// SetRack sets the Rack property
func (m *DcimInventoryItemsListQueryParameters) SetRack(val []string) {
	m.Rack = val
}

// GetRackN returns the RackN property
func (m DcimInventoryItemsListQueryParameters) GetRackN() []string {
	return m.RackN
}

// SetRackN sets the RackN property
func (m *DcimInventoryItemsListQueryParameters) SetRackN(val []string) {
	m.RackN = val
}

// GetRackId returns the RackId property
func (m DcimInventoryItemsListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimInventoryItemsListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetRackIdN returns the RackIdN property
func (m DcimInventoryItemsListQueryParameters) GetRackIdN() []int32 {
	return m.RackIdN
}

// SetRackIdN sets the RackIdN property
func (m *DcimInventoryItemsListQueryParameters) SetRackIdN(val []int32) {
	m.RackIdN = val
}

// GetRegion returns the Region property
func (m DcimInventoryItemsListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimInventoryItemsListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimInventoryItemsListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimInventoryItemsListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimInventoryItemsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimInventoryItemsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimInventoryItemsListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimInventoryItemsListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetRole returns the Role property
func (m DcimInventoryItemsListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *DcimInventoryItemsListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m DcimInventoryItemsListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *DcimInventoryItemsListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m DcimInventoryItemsListQueryParameters) GetRoleId() []*int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *DcimInventoryItemsListQueryParameters) SetRoleId(val []*int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m DcimInventoryItemsListQueryParameters) GetRoleIdN() []*int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *DcimInventoryItemsListQueryParameters) SetRoleIdN(val []*int32) {
	m.RoleIdN = val
}

// GetSerial returns the Serial property
func (m DcimInventoryItemsListQueryParameters) GetSerial() []string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *DcimInventoryItemsListQueryParameters) SetSerial(val []string) {
	m.Serial = val
}

// GetSerialEmpty returns the SerialEmpty property
func (m DcimInventoryItemsListQueryParameters) GetSerialEmpty() []string {
	return m.SerialEmpty
}

// SetSerialEmpty sets the SerialEmpty property
func (m *DcimInventoryItemsListQueryParameters) SetSerialEmpty(val []string) {
	m.SerialEmpty = val
}

// GetSerialIc returns the SerialIc property
func (m DcimInventoryItemsListQueryParameters) GetSerialIc() []string {
	return m.SerialIc
}

// SetSerialIc sets the SerialIc property
func (m *DcimInventoryItemsListQueryParameters) SetSerialIc(val []string) {
	m.SerialIc = val
}

// GetSerialIe returns the SerialIe property
func (m DcimInventoryItemsListQueryParameters) GetSerialIe() []string {
	return m.SerialIe
}

// SetSerialIe sets the SerialIe property
func (m *DcimInventoryItemsListQueryParameters) SetSerialIe(val []string) {
	m.SerialIe = val
}

// GetSerialIew returns the SerialIew property
func (m DcimInventoryItemsListQueryParameters) GetSerialIew() []string {
	return m.SerialIew
}

// SetSerialIew sets the SerialIew property
func (m *DcimInventoryItemsListQueryParameters) SetSerialIew(val []string) {
	m.SerialIew = val
}

// GetSerialIsw returns the SerialIsw property
func (m DcimInventoryItemsListQueryParameters) GetSerialIsw() []string {
	return m.SerialIsw
}

// SetSerialIsw sets the SerialIsw property
func (m *DcimInventoryItemsListQueryParameters) SetSerialIsw(val []string) {
	m.SerialIsw = val
}

// GetSerialN returns the SerialN property
func (m DcimInventoryItemsListQueryParameters) GetSerialN() []string {
	return m.SerialN
}

// SetSerialN sets the SerialN property
func (m *DcimInventoryItemsListQueryParameters) SetSerialN(val []string) {
	m.SerialN = val
}

// GetSerialNic returns the SerialNic property
func (m DcimInventoryItemsListQueryParameters) GetSerialNic() []string {
	return m.SerialNic
}

// SetSerialNic sets the SerialNic property
func (m *DcimInventoryItemsListQueryParameters) SetSerialNic(val []string) {
	m.SerialNic = val
}

// GetSerialNie returns the SerialNie property
func (m DcimInventoryItemsListQueryParameters) GetSerialNie() []string {
	return m.SerialNie
}

// SetSerialNie sets the SerialNie property
func (m *DcimInventoryItemsListQueryParameters) SetSerialNie(val []string) {
	m.SerialNie = val
}

// GetSerialNiew returns the SerialNiew property
func (m DcimInventoryItemsListQueryParameters) GetSerialNiew() []string {
	return m.SerialNiew
}

// SetSerialNiew sets the SerialNiew property
func (m *DcimInventoryItemsListQueryParameters) SetSerialNiew(val []string) {
	m.SerialNiew = val
}

// GetSerialNisw returns the SerialNisw property
func (m DcimInventoryItemsListQueryParameters) GetSerialNisw() []string {
	return m.SerialNisw
}

// SetSerialNisw sets the SerialNisw property
func (m *DcimInventoryItemsListQueryParameters) SetSerialNisw(val []string) {
	m.SerialNisw = val
}

// GetSite returns the Site property
func (m DcimInventoryItemsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimInventoryItemsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimInventoryItemsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimInventoryItemsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimInventoryItemsListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimInventoryItemsListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimInventoryItemsListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimInventoryItemsListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimInventoryItemsListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimInventoryItemsListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimInventoryItemsListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimInventoryItemsListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimInventoryItemsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimInventoryItemsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimInventoryItemsListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimInventoryItemsListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetTag returns the Tag property
func (m DcimInventoryItemsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimInventoryItemsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimInventoryItemsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimInventoryItemsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimInventoryItemsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimInventoryItemsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualChassis returns the VirtualChassis property
func (m DcimInventoryItemsListQueryParameters) GetVirtualChassis() []string {
	return m.VirtualChassis
}

// SetVirtualChassis sets the VirtualChassis property
func (m *DcimInventoryItemsListQueryParameters) SetVirtualChassis(val []string) {
	m.VirtualChassis = val
}

// GetVirtualChassisN returns the VirtualChassisN property
func (m DcimInventoryItemsListQueryParameters) GetVirtualChassisN() []string {
	return m.VirtualChassisN
}

// SetVirtualChassisN sets the VirtualChassisN property
func (m *DcimInventoryItemsListQueryParameters) SetVirtualChassisN(val []string) {
	m.VirtualChassisN = val
}

// GetVirtualChassisId returns the VirtualChassisId property
func (m DcimInventoryItemsListQueryParameters) GetVirtualChassisId() []int32 {
	return m.VirtualChassisId
}

// SetVirtualChassisId sets the VirtualChassisId property
func (m *DcimInventoryItemsListQueryParameters) SetVirtualChassisId(val []int32) {
	m.VirtualChassisId = val
}

// GetVirtualChassisIdN returns the VirtualChassisIdN property
func (m DcimInventoryItemsListQueryParameters) GetVirtualChassisIdN() []int32 {
	return m.VirtualChassisIdN
}

// SetVirtualChassisIdN sets the VirtualChassisIdN property
func (m *DcimInventoryItemsListQueryParameters) SetVirtualChassisIdN(val []int32) {
	m.VirtualChassisIdN = val
}
