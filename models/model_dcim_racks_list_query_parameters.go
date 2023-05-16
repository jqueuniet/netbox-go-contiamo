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

// DcimRacksListQueryParameters is an object.
type DcimRacksListQueryParameters struct {
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
	// DescUnits:
	DescUnits bool `json:"desc_units,omitempty" mapstructure:"desc_units,omitempty"`
	// FacilityId:
	FacilityId []string `json:"facility_id,omitempty" mapstructure:"facility_id,omitempty"`
	// FacilityIdEmpty:
	FacilityIdEmpty []string `json:"facility_id__empty,omitempty" mapstructure:"facility_id__empty,omitempty"`
	// FacilityIdIc:
	FacilityIdIc []string `json:"facility_id__ic,omitempty" mapstructure:"facility_id__ic,omitempty"`
	// FacilityIdIe:
	FacilityIdIe []string `json:"facility_id__ie,omitempty" mapstructure:"facility_id__ie,omitempty"`
	// FacilityIdIew:
	FacilityIdIew []string `json:"facility_id__iew,omitempty" mapstructure:"facility_id__iew,omitempty"`
	// FacilityIdIsw:
	FacilityIdIsw []string `json:"facility_id__isw,omitempty" mapstructure:"facility_id__isw,omitempty"`
	// FacilityIdN:
	FacilityIdN []string `json:"facility_id__n,omitempty" mapstructure:"facility_id__n,omitempty"`
	// FacilityIdNic:
	FacilityIdNic []string `json:"facility_id__nic,omitempty" mapstructure:"facility_id__nic,omitempty"`
	// FacilityIdNie:
	FacilityIdNie []string `json:"facility_id__nie,omitempty" mapstructure:"facility_id__nie,omitempty"`
	// FacilityIdNiew:
	FacilityIdNiew []string `json:"facility_id__niew,omitempty" mapstructure:"facility_id__niew,omitempty"`
	// FacilityIdNisw:
	FacilityIdNisw []string `json:"facility_id__nisw,omitempty" mapstructure:"facility_id__nisw,omitempty"`
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
	Location []int32 `json:"location,omitempty" mapstructure:"location,omitempty"`
	// LocationN: Location (slug)
	LocationN []int32 `json:"location__n,omitempty" mapstructure:"location__n,omitempty"`
	// LocationId: Location (ID)
	LocationId []int32 `json:"location_id,omitempty" mapstructure:"location_id,omitempty"`
	// LocationIdN: Location (ID)
	LocationIdN []int32 `json:"location_id__n,omitempty" mapstructure:"location_id__n,omitempty"`
	// MaxWeight:
	MaxWeight []int32 `json:"max_weight,omitempty" mapstructure:"max_weight,omitempty"`
	// MaxWeightGt:
	MaxWeightGt []int32 `json:"max_weight__gt,omitempty" mapstructure:"max_weight__gt,omitempty"`
	// MaxWeightGte:
	MaxWeightGte []int32 `json:"max_weight__gte,omitempty" mapstructure:"max_weight__gte,omitempty"`
	// MaxWeightLt:
	MaxWeightLt []int32 `json:"max_weight__lt,omitempty" mapstructure:"max_weight__lt,omitempty"`
	// MaxWeightLte:
	MaxWeightLte []int32 `json:"max_weight__lte,omitempty" mapstructure:"max_weight__lte,omitempty"`
	// MaxWeightN:
	MaxWeightN []int32 `json:"max_weight__n,omitempty" mapstructure:"max_weight__n,omitempty"`
	// MountingDepth:
	MountingDepth []int32 `json:"mounting_depth,omitempty" mapstructure:"mounting_depth,omitempty"`
	// MountingDepthGt:
	MountingDepthGt []int32 `json:"mounting_depth__gt,omitempty" mapstructure:"mounting_depth__gt,omitempty"`
	// MountingDepthGte:
	MountingDepthGte []int32 `json:"mounting_depth__gte,omitempty" mapstructure:"mounting_depth__gte,omitempty"`
	// MountingDepthLt:
	MountingDepthLt []int32 `json:"mounting_depth__lt,omitempty" mapstructure:"mounting_depth__lt,omitempty"`
	// MountingDepthLte:
	MountingDepthLte []int32 `json:"mounting_depth__lte,omitempty" mapstructure:"mounting_depth__lte,omitempty"`
	// MountingDepthN:
	MountingDepthN []int32 `json:"mounting_depth__n,omitempty" mapstructure:"mounting_depth__n,omitempty"`
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
	// OuterDepth:
	OuterDepth []int32 `json:"outer_depth,omitempty" mapstructure:"outer_depth,omitempty"`
	// OuterDepthGt:
	OuterDepthGt []int32 `json:"outer_depth__gt,omitempty" mapstructure:"outer_depth__gt,omitempty"`
	// OuterDepthGte:
	OuterDepthGte []int32 `json:"outer_depth__gte,omitempty" mapstructure:"outer_depth__gte,omitempty"`
	// OuterDepthLt:
	OuterDepthLt []int32 `json:"outer_depth__lt,omitempty" mapstructure:"outer_depth__lt,omitempty"`
	// OuterDepthLte:
	OuterDepthLte []int32 `json:"outer_depth__lte,omitempty" mapstructure:"outer_depth__lte,omitempty"`
	// OuterDepthN:
	OuterDepthN []int32 `json:"outer_depth__n,omitempty" mapstructure:"outer_depth__n,omitempty"`
	// OuterUnit: * `mm` - Millimeters
	// * `in` - Inches
	OuterUnit string `json:"outer_unit,omitempty" mapstructure:"outer_unit,omitempty"`
	// OuterUnitN: * `mm` - Millimeters
	// * `in` - Inches
	OuterUnitN string `json:"outer_unit__n,omitempty" mapstructure:"outer_unit__n,omitempty"`
	// OuterWidth:
	OuterWidth []int32 `json:"outer_width,omitempty" mapstructure:"outer_width,omitempty"`
	// OuterWidthGt:
	OuterWidthGt []int32 `json:"outer_width__gt,omitempty" mapstructure:"outer_width__gt,omitempty"`
	// OuterWidthGte:
	OuterWidthGte []int32 `json:"outer_width__gte,omitempty" mapstructure:"outer_width__gte,omitempty"`
	// OuterWidthLt:
	OuterWidthLt []int32 `json:"outer_width__lt,omitempty" mapstructure:"outer_width__lt,omitempty"`
	// OuterWidthLte:
	OuterWidthLte []int32 `json:"outer_width__lte,omitempty" mapstructure:"outer_width__lte,omitempty"`
	// OuterWidthN:
	OuterWidthN []int32 `json:"outer_width__n,omitempty" mapstructure:"outer_width__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
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
	// Site: Site (slug)
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteN: Site (slug)
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
	// Status: * `reserved` - Reserved
	// * `available` - Available
	// * `planned` - Planned
	// * `active` - Active
	// * `deprecated` - Deprecated
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `reserved` - Reserved
	// * `available` - Available
	// * `planned` - Planned
	// * `active` - Active
	// * `deprecated` - Deprecated
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
	// Type: * `2-post-frame` - 2-post frame
	// * `4-post-frame` - 4-post frame
	// * `4-post-cabinet` - 4-post cabinet
	// * `wall-frame` - Wall-mounted frame
	// * `wall-frame-vertical` - Wall-mounted frame (vertical)
	// * `wall-cabinet` - Wall-mounted cabinet
	// * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `2-post-frame` - 2-post frame
	// * `4-post-frame` - 4-post frame
	// * `4-post-cabinet` - 4-post cabinet
	// * `wall-frame` - Wall-mounted frame
	// * `wall-frame-vertical` - Wall-mounted frame (vertical)
	// * `wall-cabinet` - Wall-mounted cabinet
	// * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UHeight:
	UHeight []int32 `json:"u_height,omitempty" mapstructure:"u_height,omitempty"`
	// UHeightGt:
	UHeightGt []int32 `json:"u_height__gt,omitempty" mapstructure:"u_height__gt,omitempty"`
	// UHeightGte:
	UHeightGte []int32 `json:"u_height__gte,omitempty" mapstructure:"u_height__gte,omitempty"`
	// UHeightLt:
	UHeightLt []int32 `json:"u_height__lt,omitempty" mapstructure:"u_height__lt,omitempty"`
	// UHeightLte:
	UHeightLte []int32 `json:"u_height__lte,omitempty" mapstructure:"u_height__lte,omitempty"`
	// UHeightN:
	UHeightN []int32 `json:"u_height__n,omitempty" mapstructure:"u_height__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// Weight:
	Weight []float64 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightGt:
	WeightGt []float64 `json:"weight__gt,omitempty" mapstructure:"weight__gt,omitempty"`
	// WeightGte:
	WeightGte []float64 `json:"weight__gte,omitempty" mapstructure:"weight__gte,omitempty"`
	// WeightLt:
	WeightLt []float64 `json:"weight__lt,omitempty" mapstructure:"weight__lt,omitempty"`
	// WeightLte:
	WeightLte []float64 `json:"weight__lte,omitempty" mapstructure:"weight__lte,omitempty"`
	// WeightN:
	WeightN []float64 `json:"weight__n,omitempty" mapstructure:"weight__n,omitempty"`
	// WeightUnit: * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	WeightUnit string `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
	// WeightUnitN: * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	WeightUnitN string `json:"weight_unit__n,omitempty" mapstructure:"weight_unit__n,omitempty"`
	// Width: Rail-to-rail width
	Width []int32 `json:"width,omitempty" mapstructure:"width,omitempty"`
	// WidthN: Rail-to-rail width
	WidthN []int32 `json:"width__n,omitempty" mapstructure:"width__n,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimRacksListQueryParameters) Validate() error {
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
		"facilityId": validation.Validate(
			m.FacilityId,
		),
		"facilityIdEmpty": validation.Validate(
			m.FacilityIdEmpty,
		),
		"facilityIdIc": validation.Validate(
			m.FacilityIdIc,
		),
		"facilityIdIe": validation.Validate(
			m.FacilityIdIe,
		),
		"facilityIdIew": validation.Validate(
			m.FacilityIdIew,
		),
		"facilityIdIsw": validation.Validate(
			m.FacilityIdIsw,
		),
		"facilityIdN": validation.Validate(
			m.FacilityIdN,
		),
		"facilityIdNic": validation.Validate(
			m.FacilityIdNic,
		),
		"facilityIdNie": validation.Validate(
			m.FacilityIdNie,
		),
		"facilityIdNiew": validation.Validate(
			m.FacilityIdNiew,
		),
		"facilityIdNisw": validation.Validate(
			m.FacilityIdNisw,
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
		"maxWeight": validation.Validate(
			m.MaxWeight,
		),
		"maxWeightGt": validation.Validate(
			m.MaxWeightGt,
		),
		"maxWeightGte": validation.Validate(
			m.MaxWeightGte,
		),
		"maxWeightLt": validation.Validate(
			m.MaxWeightLt,
		),
		"maxWeightLte": validation.Validate(
			m.MaxWeightLte,
		),
		"maxWeightN": validation.Validate(
			m.MaxWeightN,
		),
		"mountingDepth": validation.Validate(
			m.MountingDepth,
		),
		"mountingDepthGt": validation.Validate(
			m.MountingDepthGt,
		),
		"mountingDepthGte": validation.Validate(
			m.MountingDepthGte,
		),
		"mountingDepthLt": validation.Validate(
			m.MountingDepthLt,
		),
		"mountingDepthLte": validation.Validate(
			m.MountingDepthLte,
		),
		"mountingDepthN": validation.Validate(
			m.MountingDepthN,
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
		"outerDepth": validation.Validate(
			m.OuterDepth,
		),
		"outerDepthGt": validation.Validate(
			m.OuterDepthGt,
		),
		"outerDepthGte": validation.Validate(
			m.OuterDepthGte,
		),
		"outerDepthLt": validation.Validate(
			m.OuterDepthLt,
		),
		"outerDepthLte": validation.Validate(
			m.OuterDepthLte,
		),
		"outerDepthN": validation.Validate(
			m.OuterDepthN,
		),
		"outerWidth": validation.Validate(
			m.OuterWidth,
		),
		"outerWidthGt": validation.Validate(
			m.OuterWidthGt,
		),
		"outerWidthGte": validation.Validate(
			m.OuterWidthGte,
		),
		"outerWidthLt": validation.Validate(
			m.OuterWidthLt,
		),
		"outerWidthLte": validation.Validate(
			m.OuterWidthLte,
		),
		"outerWidthN": validation.Validate(
			m.OuterWidthN,
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
		"type": validation.Validate(
			m.Type,
		),
		"typeN": validation.Validate(
			m.TypeN,
		),
		"uHeight": validation.Validate(
			m.UHeight,
		),
		"uHeightGt": validation.Validate(
			m.UHeightGt,
		),
		"uHeightGte": validation.Validate(
			m.UHeightGte,
		),
		"uHeightLt": validation.Validate(
			m.UHeightLt,
		),
		"uHeightLte": validation.Validate(
			m.UHeightLte,
		),
		"uHeightN": validation.Validate(
			m.UHeightN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"weight": validation.Validate(
			m.Weight,
		),
		"weightGt": validation.Validate(
			m.WeightGt,
		),
		"weightGte": validation.Validate(
			m.WeightGte,
		),
		"weightLt": validation.Validate(
			m.WeightLt,
		),
		"weightLte": validation.Validate(
			m.WeightLte,
		),
		"weightN": validation.Validate(
			m.WeightN,
		),
		"width": validation.Validate(
			m.Width,
		),
		"widthN": validation.Validate(
			m.WidthN,
		),
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m DcimRacksListQueryParameters) GetAssetTag() []string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *DcimRacksListQueryParameters) SetAssetTag(val []string) {
	m.AssetTag = val
}

// GetAssetTagEmpty returns the AssetTagEmpty property
func (m DcimRacksListQueryParameters) GetAssetTagEmpty() []string {
	return m.AssetTagEmpty
}

// SetAssetTagEmpty sets the AssetTagEmpty property
func (m *DcimRacksListQueryParameters) SetAssetTagEmpty(val []string) {
	m.AssetTagEmpty = val
}

// GetAssetTagIc returns the AssetTagIc property
func (m DcimRacksListQueryParameters) GetAssetTagIc() []string {
	return m.AssetTagIc
}

// SetAssetTagIc sets the AssetTagIc property
func (m *DcimRacksListQueryParameters) SetAssetTagIc(val []string) {
	m.AssetTagIc = val
}

// GetAssetTagIe returns the AssetTagIe property
func (m DcimRacksListQueryParameters) GetAssetTagIe() []string {
	return m.AssetTagIe
}

// SetAssetTagIe sets the AssetTagIe property
func (m *DcimRacksListQueryParameters) SetAssetTagIe(val []string) {
	m.AssetTagIe = val
}

// GetAssetTagIew returns the AssetTagIew property
func (m DcimRacksListQueryParameters) GetAssetTagIew() []string {
	return m.AssetTagIew
}

// SetAssetTagIew sets the AssetTagIew property
func (m *DcimRacksListQueryParameters) SetAssetTagIew(val []string) {
	m.AssetTagIew = val
}

// GetAssetTagIsw returns the AssetTagIsw property
func (m DcimRacksListQueryParameters) GetAssetTagIsw() []string {
	return m.AssetTagIsw
}

// SetAssetTagIsw sets the AssetTagIsw property
func (m *DcimRacksListQueryParameters) SetAssetTagIsw(val []string) {
	m.AssetTagIsw = val
}

// GetAssetTagN returns the AssetTagN property
func (m DcimRacksListQueryParameters) GetAssetTagN() []string {
	return m.AssetTagN
}

// SetAssetTagN sets the AssetTagN property
func (m *DcimRacksListQueryParameters) SetAssetTagN(val []string) {
	m.AssetTagN = val
}

// GetAssetTagNic returns the AssetTagNic property
func (m DcimRacksListQueryParameters) GetAssetTagNic() []string {
	return m.AssetTagNic
}

// SetAssetTagNic sets the AssetTagNic property
func (m *DcimRacksListQueryParameters) SetAssetTagNic(val []string) {
	m.AssetTagNic = val
}

// GetAssetTagNie returns the AssetTagNie property
func (m DcimRacksListQueryParameters) GetAssetTagNie() []string {
	return m.AssetTagNie
}

// SetAssetTagNie sets the AssetTagNie property
func (m *DcimRacksListQueryParameters) SetAssetTagNie(val []string) {
	m.AssetTagNie = val
}

// GetAssetTagNiew returns the AssetTagNiew property
func (m DcimRacksListQueryParameters) GetAssetTagNiew() []string {
	return m.AssetTagNiew
}

// SetAssetTagNiew sets the AssetTagNiew property
func (m *DcimRacksListQueryParameters) SetAssetTagNiew(val []string) {
	m.AssetTagNiew = val
}

// GetAssetTagNisw returns the AssetTagNisw property
func (m DcimRacksListQueryParameters) GetAssetTagNisw() []string {
	return m.AssetTagNisw
}

// SetAssetTagNisw sets the AssetTagNisw property
func (m *DcimRacksListQueryParameters) SetAssetTagNisw(val []string) {
	m.AssetTagNisw = val
}

// GetContact returns the Contact property
func (m DcimRacksListQueryParameters) GetContact() []int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *DcimRacksListQueryParameters) SetContact(val []int32) {
	m.Contact = val
}

// GetContactN returns the ContactN property
func (m DcimRacksListQueryParameters) GetContactN() []int32 {
	return m.ContactN
}

// SetContactN sets the ContactN property
func (m *DcimRacksListQueryParameters) SetContactN(val []int32) {
	m.ContactN = val
}

// GetContactGroup returns the ContactGroup property
func (m DcimRacksListQueryParameters) GetContactGroup() []int32 {
	return m.ContactGroup
}

// SetContactGroup sets the ContactGroup property
func (m *DcimRacksListQueryParameters) SetContactGroup(val []int32) {
	m.ContactGroup = val
}

// GetContactGroupN returns the ContactGroupN property
func (m DcimRacksListQueryParameters) GetContactGroupN() []int32 {
	return m.ContactGroupN
}

// SetContactGroupN sets the ContactGroupN property
func (m *DcimRacksListQueryParameters) SetContactGroupN(val []int32) {
	m.ContactGroupN = val
}

// GetContactRole returns the ContactRole property
func (m DcimRacksListQueryParameters) GetContactRole() []int32 {
	return m.ContactRole
}

// SetContactRole sets the ContactRole property
func (m *DcimRacksListQueryParameters) SetContactRole(val []int32) {
	m.ContactRole = val
}

// GetContactRoleN returns the ContactRoleN property
func (m DcimRacksListQueryParameters) GetContactRoleN() []int32 {
	return m.ContactRoleN
}

// SetContactRoleN sets the ContactRoleN property
func (m *DcimRacksListQueryParameters) SetContactRoleN(val []int32) {
	m.ContactRoleN = val
}

// GetCreated returns the Created property
func (m DcimRacksListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimRacksListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimRacksListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimRacksListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimRacksListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimRacksListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimRacksListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimRacksListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimRacksListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimRacksListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimRacksListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimRacksListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimRacksListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimRacksListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescUnits returns the DescUnits property
func (m DcimRacksListQueryParameters) GetDescUnits() bool {
	return m.DescUnits
}

// SetDescUnits sets the DescUnits property
func (m *DcimRacksListQueryParameters) SetDescUnits(val bool) {
	m.DescUnits = val
}

// GetFacilityId returns the FacilityId property
func (m DcimRacksListQueryParameters) GetFacilityId() []string {
	return m.FacilityId
}

// SetFacilityId sets the FacilityId property
func (m *DcimRacksListQueryParameters) SetFacilityId(val []string) {
	m.FacilityId = val
}

// GetFacilityIdEmpty returns the FacilityIdEmpty property
func (m DcimRacksListQueryParameters) GetFacilityIdEmpty() []string {
	return m.FacilityIdEmpty
}

// SetFacilityIdEmpty sets the FacilityIdEmpty property
func (m *DcimRacksListQueryParameters) SetFacilityIdEmpty(val []string) {
	m.FacilityIdEmpty = val
}

// GetFacilityIdIc returns the FacilityIdIc property
func (m DcimRacksListQueryParameters) GetFacilityIdIc() []string {
	return m.FacilityIdIc
}

// SetFacilityIdIc sets the FacilityIdIc property
func (m *DcimRacksListQueryParameters) SetFacilityIdIc(val []string) {
	m.FacilityIdIc = val
}

// GetFacilityIdIe returns the FacilityIdIe property
func (m DcimRacksListQueryParameters) GetFacilityIdIe() []string {
	return m.FacilityIdIe
}

// SetFacilityIdIe sets the FacilityIdIe property
func (m *DcimRacksListQueryParameters) SetFacilityIdIe(val []string) {
	m.FacilityIdIe = val
}

// GetFacilityIdIew returns the FacilityIdIew property
func (m DcimRacksListQueryParameters) GetFacilityIdIew() []string {
	return m.FacilityIdIew
}

// SetFacilityIdIew sets the FacilityIdIew property
func (m *DcimRacksListQueryParameters) SetFacilityIdIew(val []string) {
	m.FacilityIdIew = val
}

// GetFacilityIdIsw returns the FacilityIdIsw property
func (m DcimRacksListQueryParameters) GetFacilityIdIsw() []string {
	return m.FacilityIdIsw
}

// SetFacilityIdIsw sets the FacilityIdIsw property
func (m *DcimRacksListQueryParameters) SetFacilityIdIsw(val []string) {
	m.FacilityIdIsw = val
}

// GetFacilityIdN returns the FacilityIdN property
func (m DcimRacksListQueryParameters) GetFacilityIdN() []string {
	return m.FacilityIdN
}

// SetFacilityIdN sets the FacilityIdN property
func (m *DcimRacksListQueryParameters) SetFacilityIdN(val []string) {
	m.FacilityIdN = val
}

// GetFacilityIdNic returns the FacilityIdNic property
func (m DcimRacksListQueryParameters) GetFacilityIdNic() []string {
	return m.FacilityIdNic
}

// SetFacilityIdNic sets the FacilityIdNic property
func (m *DcimRacksListQueryParameters) SetFacilityIdNic(val []string) {
	m.FacilityIdNic = val
}

// GetFacilityIdNie returns the FacilityIdNie property
func (m DcimRacksListQueryParameters) GetFacilityIdNie() []string {
	return m.FacilityIdNie
}

// SetFacilityIdNie sets the FacilityIdNie property
func (m *DcimRacksListQueryParameters) SetFacilityIdNie(val []string) {
	m.FacilityIdNie = val
}

// GetFacilityIdNiew returns the FacilityIdNiew property
func (m DcimRacksListQueryParameters) GetFacilityIdNiew() []string {
	return m.FacilityIdNiew
}

// SetFacilityIdNiew sets the FacilityIdNiew property
func (m *DcimRacksListQueryParameters) SetFacilityIdNiew(val []string) {
	m.FacilityIdNiew = val
}

// GetFacilityIdNisw returns the FacilityIdNisw property
func (m DcimRacksListQueryParameters) GetFacilityIdNisw() []string {
	return m.FacilityIdNisw
}

// SetFacilityIdNisw sets the FacilityIdNisw property
func (m *DcimRacksListQueryParameters) SetFacilityIdNisw(val []string) {
	m.FacilityIdNisw = val
}

// GetId returns the Id property
func (m DcimRacksListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimRacksListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimRacksListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimRacksListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimRacksListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimRacksListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimRacksListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimRacksListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimRacksListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimRacksListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimRacksListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimRacksListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimRacksListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimRacksListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimRacksListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimRacksListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimRacksListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimRacksListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimRacksListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimRacksListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimRacksListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimRacksListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimRacksListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimRacksListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimRacksListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimRacksListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m DcimRacksListQueryParameters) GetLocation() []int32 {
	return m.Location
}

// SetLocation sets the Location property
func (m *DcimRacksListQueryParameters) SetLocation(val []int32) {
	m.Location = val
}

// GetLocationN returns the LocationN property
func (m DcimRacksListQueryParameters) GetLocationN() []int32 {
	return m.LocationN
}

// SetLocationN sets the LocationN property
func (m *DcimRacksListQueryParameters) SetLocationN(val []int32) {
	m.LocationN = val
}

// GetLocationId returns the LocationId property
func (m DcimRacksListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimRacksListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m DcimRacksListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *DcimRacksListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetMaxWeight returns the MaxWeight property
func (m DcimRacksListQueryParameters) GetMaxWeight() []int32 {
	return m.MaxWeight
}

// SetMaxWeight sets the MaxWeight property
func (m *DcimRacksListQueryParameters) SetMaxWeight(val []int32) {
	m.MaxWeight = val
}

// GetMaxWeightGt returns the MaxWeightGt property
func (m DcimRacksListQueryParameters) GetMaxWeightGt() []int32 {
	return m.MaxWeightGt
}

// SetMaxWeightGt sets the MaxWeightGt property
func (m *DcimRacksListQueryParameters) SetMaxWeightGt(val []int32) {
	m.MaxWeightGt = val
}

// GetMaxWeightGte returns the MaxWeightGte property
func (m DcimRacksListQueryParameters) GetMaxWeightGte() []int32 {
	return m.MaxWeightGte
}

// SetMaxWeightGte sets the MaxWeightGte property
func (m *DcimRacksListQueryParameters) SetMaxWeightGte(val []int32) {
	m.MaxWeightGte = val
}

// GetMaxWeightLt returns the MaxWeightLt property
func (m DcimRacksListQueryParameters) GetMaxWeightLt() []int32 {
	return m.MaxWeightLt
}

// SetMaxWeightLt sets the MaxWeightLt property
func (m *DcimRacksListQueryParameters) SetMaxWeightLt(val []int32) {
	m.MaxWeightLt = val
}

// GetMaxWeightLte returns the MaxWeightLte property
func (m DcimRacksListQueryParameters) GetMaxWeightLte() []int32 {
	return m.MaxWeightLte
}

// SetMaxWeightLte sets the MaxWeightLte property
func (m *DcimRacksListQueryParameters) SetMaxWeightLte(val []int32) {
	m.MaxWeightLte = val
}

// GetMaxWeightN returns the MaxWeightN property
func (m DcimRacksListQueryParameters) GetMaxWeightN() []int32 {
	return m.MaxWeightN
}

// SetMaxWeightN sets the MaxWeightN property
func (m *DcimRacksListQueryParameters) SetMaxWeightN(val []int32) {
	m.MaxWeightN = val
}

// GetMountingDepth returns the MountingDepth property
func (m DcimRacksListQueryParameters) GetMountingDepth() []int32 {
	return m.MountingDepth
}

// SetMountingDepth sets the MountingDepth property
func (m *DcimRacksListQueryParameters) SetMountingDepth(val []int32) {
	m.MountingDepth = val
}

// GetMountingDepthGt returns the MountingDepthGt property
func (m DcimRacksListQueryParameters) GetMountingDepthGt() []int32 {
	return m.MountingDepthGt
}

// SetMountingDepthGt sets the MountingDepthGt property
func (m *DcimRacksListQueryParameters) SetMountingDepthGt(val []int32) {
	m.MountingDepthGt = val
}

// GetMountingDepthGte returns the MountingDepthGte property
func (m DcimRacksListQueryParameters) GetMountingDepthGte() []int32 {
	return m.MountingDepthGte
}

// SetMountingDepthGte sets the MountingDepthGte property
func (m *DcimRacksListQueryParameters) SetMountingDepthGte(val []int32) {
	m.MountingDepthGte = val
}

// GetMountingDepthLt returns the MountingDepthLt property
func (m DcimRacksListQueryParameters) GetMountingDepthLt() []int32 {
	return m.MountingDepthLt
}

// SetMountingDepthLt sets the MountingDepthLt property
func (m *DcimRacksListQueryParameters) SetMountingDepthLt(val []int32) {
	m.MountingDepthLt = val
}

// GetMountingDepthLte returns the MountingDepthLte property
func (m DcimRacksListQueryParameters) GetMountingDepthLte() []int32 {
	return m.MountingDepthLte
}

// SetMountingDepthLte sets the MountingDepthLte property
func (m *DcimRacksListQueryParameters) SetMountingDepthLte(val []int32) {
	m.MountingDepthLte = val
}

// GetMountingDepthN returns the MountingDepthN property
func (m DcimRacksListQueryParameters) GetMountingDepthN() []int32 {
	return m.MountingDepthN
}

// SetMountingDepthN sets the MountingDepthN property
func (m *DcimRacksListQueryParameters) SetMountingDepthN(val []int32) {
	m.MountingDepthN = val
}

// GetName returns the Name property
func (m DcimRacksListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimRacksListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimRacksListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimRacksListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimRacksListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimRacksListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimRacksListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimRacksListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimRacksListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimRacksListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimRacksListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimRacksListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimRacksListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimRacksListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimRacksListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimRacksListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimRacksListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimRacksListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimRacksListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimRacksListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimRacksListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimRacksListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimRacksListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimRacksListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimRacksListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimRacksListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetOuterDepth returns the OuterDepth property
func (m DcimRacksListQueryParameters) GetOuterDepth() []int32 {
	return m.OuterDepth
}

// SetOuterDepth sets the OuterDepth property
func (m *DcimRacksListQueryParameters) SetOuterDepth(val []int32) {
	m.OuterDepth = val
}

// GetOuterDepthGt returns the OuterDepthGt property
func (m DcimRacksListQueryParameters) GetOuterDepthGt() []int32 {
	return m.OuterDepthGt
}

// SetOuterDepthGt sets the OuterDepthGt property
func (m *DcimRacksListQueryParameters) SetOuterDepthGt(val []int32) {
	m.OuterDepthGt = val
}

// GetOuterDepthGte returns the OuterDepthGte property
func (m DcimRacksListQueryParameters) GetOuterDepthGte() []int32 {
	return m.OuterDepthGte
}

// SetOuterDepthGte sets the OuterDepthGte property
func (m *DcimRacksListQueryParameters) SetOuterDepthGte(val []int32) {
	m.OuterDepthGte = val
}

// GetOuterDepthLt returns the OuterDepthLt property
func (m DcimRacksListQueryParameters) GetOuterDepthLt() []int32 {
	return m.OuterDepthLt
}

// SetOuterDepthLt sets the OuterDepthLt property
func (m *DcimRacksListQueryParameters) SetOuterDepthLt(val []int32) {
	m.OuterDepthLt = val
}

// GetOuterDepthLte returns the OuterDepthLte property
func (m DcimRacksListQueryParameters) GetOuterDepthLte() []int32 {
	return m.OuterDepthLte
}

// SetOuterDepthLte sets the OuterDepthLte property
func (m *DcimRacksListQueryParameters) SetOuterDepthLte(val []int32) {
	m.OuterDepthLte = val
}

// GetOuterDepthN returns the OuterDepthN property
func (m DcimRacksListQueryParameters) GetOuterDepthN() []int32 {
	return m.OuterDepthN
}

// SetOuterDepthN sets the OuterDepthN property
func (m *DcimRacksListQueryParameters) SetOuterDepthN(val []int32) {
	m.OuterDepthN = val
}

// GetOuterUnit returns the OuterUnit property
func (m DcimRacksListQueryParameters) GetOuterUnit() string {
	return m.OuterUnit
}

// SetOuterUnit sets the OuterUnit property
func (m *DcimRacksListQueryParameters) SetOuterUnit(val string) {
	m.OuterUnit = val
}

// GetOuterUnitN returns the OuterUnitN property
func (m DcimRacksListQueryParameters) GetOuterUnitN() string {
	return m.OuterUnitN
}

// SetOuterUnitN sets the OuterUnitN property
func (m *DcimRacksListQueryParameters) SetOuterUnitN(val string) {
	m.OuterUnitN = val
}

// GetOuterWidth returns the OuterWidth property
func (m DcimRacksListQueryParameters) GetOuterWidth() []int32 {
	return m.OuterWidth
}

// SetOuterWidth sets the OuterWidth property
func (m *DcimRacksListQueryParameters) SetOuterWidth(val []int32) {
	m.OuterWidth = val
}

// GetOuterWidthGt returns the OuterWidthGt property
func (m DcimRacksListQueryParameters) GetOuterWidthGt() []int32 {
	return m.OuterWidthGt
}

// SetOuterWidthGt sets the OuterWidthGt property
func (m *DcimRacksListQueryParameters) SetOuterWidthGt(val []int32) {
	m.OuterWidthGt = val
}

// GetOuterWidthGte returns the OuterWidthGte property
func (m DcimRacksListQueryParameters) GetOuterWidthGte() []int32 {
	return m.OuterWidthGte
}

// SetOuterWidthGte sets the OuterWidthGte property
func (m *DcimRacksListQueryParameters) SetOuterWidthGte(val []int32) {
	m.OuterWidthGte = val
}

// GetOuterWidthLt returns the OuterWidthLt property
func (m DcimRacksListQueryParameters) GetOuterWidthLt() []int32 {
	return m.OuterWidthLt
}

// SetOuterWidthLt sets the OuterWidthLt property
func (m *DcimRacksListQueryParameters) SetOuterWidthLt(val []int32) {
	m.OuterWidthLt = val
}

// GetOuterWidthLte returns the OuterWidthLte property
func (m DcimRacksListQueryParameters) GetOuterWidthLte() []int32 {
	return m.OuterWidthLte
}

// SetOuterWidthLte sets the OuterWidthLte property
func (m *DcimRacksListQueryParameters) SetOuterWidthLte(val []int32) {
	m.OuterWidthLte = val
}

// GetOuterWidthN returns the OuterWidthN property
func (m DcimRacksListQueryParameters) GetOuterWidthN() []int32 {
	return m.OuterWidthN
}

// SetOuterWidthN sets the OuterWidthN property
func (m *DcimRacksListQueryParameters) SetOuterWidthN(val []int32) {
	m.OuterWidthN = val
}

// GetQ returns the Q property
func (m DcimRacksListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimRacksListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m DcimRacksListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimRacksListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimRacksListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimRacksListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimRacksListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimRacksListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimRacksListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimRacksListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetRole returns the Role property
func (m DcimRacksListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *DcimRacksListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m DcimRacksListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *DcimRacksListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m DcimRacksListQueryParameters) GetRoleId() []*int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *DcimRacksListQueryParameters) SetRoleId(val []*int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m DcimRacksListQueryParameters) GetRoleIdN() []*int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *DcimRacksListQueryParameters) SetRoleIdN(val []*int32) {
	m.RoleIdN = val
}

// GetSerial returns the Serial property
func (m DcimRacksListQueryParameters) GetSerial() []string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *DcimRacksListQueryParameters) SetSerial(val []string) {
	m.Serial = val
}

// GetSerialEmpty returns the SerialEmpty property
func (m DcimRacksListQueryParameters) GetSerialEmpty() []string {
	return m.SerialEmpty
}

// SetSerialEmpty sets the SerialEmpty property
func (m *DcimRacksListQueryParameters) SetSerialEmpty(val []string) {
	m.SerialEmpty = val
}

// GetSerialIc returns the SerialIc property
func (m DcimRacksListQueryParameters) GetSerialIc() []string {
	return m.SerialIc
}

// SetSerialIc sets the SerialIc property
func (m *DcimRacksListQueryParameters) SetSerialIc(val []string) {
	m.SerialIc = val
}

// GetSerialIe returns the SerialIe property
func (m DcimRacksListQueryParameters) GetSerialIe() []string {
	return m.SerialIe
}

// SetSerialIe sets the SerialIe property
func (m *DcimRacksListQueryParameters) SetSerialIe(val []string) {
	m.SerialIe = val
}

// GetSerialIew returns the SerialIew property
func (m DcimRacksListQueryParameters) GetSerialIew() []string {
	return m.SerialIew
}

// SetSerialIew sets the SerialIew property
func (m *DcimRacksListQueryParameters) SetSerialIew(val []string) {
	m.SerialIew = val
}

// GetSerialIsw returns the SerialIsw property
func (m DcimRacksListQueryParameters) GetSerialIsw() []string {
	return m.SerialIsw
}

// SetSerialIsw sets the SerialIsw property
func (m *DcimRacksListQueryParameters) SetSerialIsw(val []string) {
	m.SerialIsw = val
}

// GetSerialN returns the SerialN property
func (m DcimRacksListQueryParameters) GetSerialN() []string {
	return m.SerialN
}

// SetSerialN sets the SerialN property
func (m *DcimRacksListQueryParameters) SetSerialN(val []string) {
	m.SerialN = val
}

// GetSerialNic returns the SerialNic property
func (m DcimRacksListQueryParameters) GetSerialNic() []string {
	return m.SerialNic
}

// SetSerialNic sets the SerialNic property
func (m *DcimRacksListQueryParameters) SetSerialNic(val []string) {
	m.SerialNic = val
}

// GetSerialNie returns the SerialNie property
func (m DcimRacksListQueryParameters) GetSerialNie() []string {
	return m.SerialNie
}

// SetSerialNie sets the SerialNie property
func (m *DcimRacksListQueryParameters) SetSerialNie(val []string) {
	m.SerialNie = val
}

// GetSerialNiew returns the SerialNiew property
func (m DcimRacksListQueryParameters) GetSerialNiew() []string {
	return m.SerialNiew
}

// SetSerialNiew sets the SerialNiew property
func (m *DcimRacksListQueryParameters) SetSerialNiew(val []string) {
	m.SerialNiew = val
}

// GetSerialNisw returns the SerialNisw property
func (m DcimRacksListQueryParameters) GetSerialNisw() []string {
	return m.SerialNisw
}

// SetSerialNisw sets the SerialNisw property
func (m *DcimRacksListQueryParameters) SetSerialNisw(val []string) {
	m.SerialNisw = val
}

// GetSite returns the Site property
func (m DcimRacksListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimRacksListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimRacksListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimRacksListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimRacksListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimRacksListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimRacksListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimRacksListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimRacksListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimRacksListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimRacksListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimRacksListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimRacksListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimRacksListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimRacksListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimRacksListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetStatus returns the Status property
func (m DcimRacksListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *DcimRacksListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m DcimRacksListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *DcimRacksListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m DcimRacksListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimRacksListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimRacksListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimRacksListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m DcimRacksListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DcimRacksListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m DcimRacksListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *DcimRacksListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m DcimRacksListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *DcimRacksListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m DcimRacksListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *DcimRacksListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m DcimRacksListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *DcimRacksListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m DcimRacksListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *DcimRacksListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m DcimRacksListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *DcimRacksListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m DcimRacksListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *DcimRacksListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetType returns the Type property
func (m DcimRacksListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimRacksListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimRacksListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimRacksListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUHeight returns the UHeight property
func (m DcimRacksListQueryParameters) GetUHeight() []int32 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *DcimRacksListQueryParameters) SetUHeight(val []int32) {
	m.UHeight = val
}

// GetUHeightGt returns the UHeightGt property
func (m DcimRacksListQueryParameters) GetUHeightGt() []int32 {
	return m.UHeightGt
}

// SetUHeightGt sets the UHeightGt property
func (m *DcimRacksListQueryParameters) SetUHeightGt(val []int32) {
	m.UHeightGt = val
}

// GetUHeightGte returns the UHeightGte property
func (m DcimRacksListQueryParameters) GetUHeightGte() []int32 {
	return m.UHeightGte
}

// SetUHeightGte sets the UHeightGte property
func (m *DcimRacksListQueryParameters) SetUHeightGte(val []int32) {
	m.UHeightGte = val
}

// GetUHeightLt returns the UHeightLt property
func (m DcimRacksListQueryParameters) GetUHeightLt() []int32 {
	return m.UHeightLt
}

// SetUHeightLt sets the UHeightLt property
func (m *DcimRacksListQueryParameters) SetUHeightLt(val []int32) {
	m.UHeightLt = val
}

// GetUHeightLte returns the UHeightLte property
func (m DcimRacksListQueryParameters) GetUHeightLte() []int32 {
	return m.UHeightLte
}

// SetUHeightLte sets the UHeightLte property
func (m *DcimRacksListQueryParameters) SetUHeightLte(val []int32) {
	m.UHeightLte = val
}

// GetUHeightN returns the UHeightN property
func (m DcimRacksListQueryParameters) GetUHeightN() []int32 {
	return m.UHeightN
}

// SetUHeightN sets the UHeightN property
func (m *DcimRacksListQueryParameters) SetUHeightN(val []int32) {
	m.UHeightN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimRacksListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimRacksListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetWeight returns the Weight property
func (m DcimRacksListQueryParameters) GetWeight() []float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *DcimRacksListQueryParameters) SetWeight(val []float64) {
	m.Weight = val
}

// GetWeightGt returns the WeightGt property
func (m DcimRacksListQueryParameters) GetWeightGt() []float64 {
	return m.WeightGt
}

// SetWeightGt sets the WeightGt property
func (m *DcimRacksListQueryParameters) SetWeightGt(val []float64) {
	m.WeightGt = val
}

// GetWeightGte returns the WeightGte property
func (m DcimRacksListQueryParameters) GetWeightGte() []float64 {
	return m.WeightGte
}

// SetWeightGte sets the WeightGte property
func (m *DcimRacksListQueryParameters) SetWeightGte(val []float64) {
	m.WeightGte = val
}

// GetWeightLt returns the WeightLt property
func (m DcimRacksListQueryParameters) GetWeightLt() []float64 {
	return m.WeightLt
}

// SetWeightLt sets the WeightLt property
func (m *DcimRacksListQueryParameters) SetWeightLt(val []float64) {
	m.WeightLt = val
}

// GetWeightLte returns the WeightLte property
func (m DcimRacksListQueryParameters) GetWeightLte() []float64 {
	return m.WeightLte
}

// SetWeightLte sets the WeightLte property
func (m *DcimRacksListQueryParameters) SetWeightLte(val []float64) {
	m.WeightLte = val
}

// GetWeightN returns the WeightN property
func (m DcimRacksListQueryParameters) GetWeightN() []float64 {
	return m.WeightN
}

// SetWeightN sets the WeightN property
func (m *DcimRacksListQueryParameters) SetWeightN(val []float64) {
	m.WeightN = val
}

// GetWeightUnit returns the WeightUnit property
func (m DcimRacksListQueryParameters) GetWeightUnit() string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *DcimRacksListQueryParameters) SetWeightUnit(val string) {
	m.WeightUnit = val
}

// GetWeightUnitN returns the WeightUnitN property
func (m DcimRacksListQueryParameters) GetWeightUnitN() string {
	return m.WeightUnitN
}

// SetWeightUnitN sets the WeightUnitN property
func (m *DcimRacksListQueryParameters) SetWeightUnitN(val string) {
	m.WeightUnitN = val
}

// GetWidth returns the Width property
func (m DcimRacksListQueryParameters) GetWidth() []int32 {
	return m.Width
}

// SetWidth sets the Width property
func (m *DcimRacksListQueryParameters) SetWidth(val []int32) {
	m.Width = val
}

// GetWidthN returns the WidthN property
func (m DcimRacksListQueryParameters) GetWidthN() []int32 {
	return m.WidthN
}

// SetWidthN sets the WidthN property
func (m *DcimRacksListQueryParameters) SetWidthN(val []int32) {
	m.WidthN = val
}
