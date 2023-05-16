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

// DcimFrontPortsListQueryParameters is an object.
type DcimFrontPortsListQueryParameters struct {
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end,omitempty" mapstructure:"cable_end,omitempty"`
	// CableEndN: * `A` - A
	// * `B` - B
	CableEndN string `json:"cable_end__n,omitempty" mapstructure:"cable_end__n,omitempty"`
	// Cabled:
	Cabled bool `json:"cabled,omitempty" mapstructure:"cabled,omitempty"`
	// Color:
	Color []string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// ColorEmpty:
	ColorEmpty []string `json:"color__empty,omitempty" mapstructure:"color__empty,omitempty"`
	// ColorIc:
	ColorIc []string `json:"color__ic,omitempty" mapstructure:"color__ic,omitempty"`
	// ColorIe:
	ColorIe []string `json:"color__ie,omitempty" mapstructure:"color__ie,omitempty"`
	// ColorIew:
	ColorIew []string `json:"color__iew,omitempty" mapstructure:"color__iew,omitempty"`
	// ColorIsw:
	ColorIsw []string `json:"color__isw,omitempty" mapstructure:"color__isw,omitempty"`
	// ColorN:
	ColorN []string `json:"color__n,omitempty" mapstructure:"color__n,omitempty"`
	// ColorNic:
	ColorNic []string `json:"color__nic,omitempty" mapstructure:"color__nic,omitempty"`
	// ColorNie:
	ColorNie []string `json:"color__nie,omitempty" mapstructure:"color__nie,omitempty"`
	// ColorNiew:
	ColorNiew []string `json:"color__niew,omitempty" mapstructure:"color__niew,omitempty"`
	// ColorNisw:
	ColorNisw []string `json:"color__nisw,omitempty" mapstructure:"color__nisw,omitempty"`
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
	// Description:
	Description []string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DescriptionEmpty:
	DescriptionEmpty []string `json:"description__empty,omitempty" mapstructure:"description__empty,omitempty"`
	// DescriptionIc:
	DescriptionIc []string `json:"description__ic,omitempty" mapstructure:"description__ic,omitempty"`
	// DescriptionIe:
	DescriptionIe []string `json:"description__ie,omitempty" mapstructure:"description__ie,omitempty"`
	// DescriptionIew:
	DescriptionIew []string `json:"description__iew,omitempty" mapstructure:"description__iew,omitempty"`
	// DescriptionIsw:
	DescriptionIsw []string `json:"description__isw,omitempty" mapstructure:"description__isw,omitempty"`
	// DescriptionN:
	DescriptionN []string `json:"description__n,omitempty" mapstructure:"description__n,omitempty"`
	// DescriptionNic:
	DescriptionNic []string `json:"description__nic,omitempty" mapstructure:"description__nic,omitempty"`
	// DescriptionNie:
	DescriptionNie []string `json:"description__nie,omitempty" mapstructure:"description__nie,omitempty"`
	// DescriptionNiew:
	DescriptionNiew []string `json:"description__niew,omitempty" mapstructure:"description__niew,omitempty"`
	// DescriptionNisw:
	DescriptionNisw []string `json:"description__nisw,omitempty" mapstructure:"description__nisw,omitempty"`
	// Device: Device (name)
	Device []*string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceN: Device (name)
	DeviceN []*string `json:"device__n,omitempty" mapstructure:"device__n,omitempty"`
	// DeviceId: Device (ID)
	DeviceId []int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
	// DeviceIdN: Device (ID)
	DeviceIdN []int32 `json:"device_id__n,omitempty" mapstructure:"device_id__n,omitempty"`
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
	// ModuleId: Module (ID)
	ModuleId []*int32 `json:"module_id,omitempty" mapstructure:"module_id,omitempty"`
	// ModuleIdN: Module (ID)
	ModuleIdN []*int32 `json:"module_id__n,omitempty" mapstructure:"module_id__n,omitempty"`
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
	// Occupied:
	Occupied bool `json:"occupied,omitempty" mapstructure:"occupied,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
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
	// Type: * `8p8c` - 8P8C
	// * `8p6c` - 8P6C
	// * `8p4c` - 8P4C
	// * `8p2c` - 8P2C
	// * `6p6c` - 6P6C
	// * `6p4c` - 6P4C
	// * `6p2c` - 6P2C
	// * `4p4c` - 4P4C
	// * `4p2c` - 4P2C
	// * `gg45` - GG45
	// * `tera-4p` - TERA 4P
	// * `tera-2p` - TERA 2P
	// * `tera-1p` - TERA 1P
	// * `110-punch` - 110 Punch
	// * `bnc` - BNC
	// * `f` - F Connector
	// * `n` - N Connector
	// * `mrj21` - MRJ21
	// * `fc` - FC
	// * `lc` - LC
	// * `lc-pc` - LC/PC
	// * `lc-upc` - LC/UPC
	// * `lc-apc` - LC/APC
	// * `lsh` - LSH
	// * `lsh-pc` - LSH/PC
	// * `lsh-upc` - LSH/UPC
	// * `lsh-apc` - LSH/APC
	// * `mpo` - MPO
	// * `mtrj` - MTRJ
	// * `sc` - SC
	// * `sc-pc` - SC/PC
	// * `sc-upc` - SC/UPC
	// * `sc-apc` - SC/APC
	// * `st` - ST
	// * `cs` - CS
	// * `sn` - SN
	// * `sma-905` - SMA 905
	// * `sma-906` - SMA 906
	// * `urm-p2` - URM-P2
	// * `urm-p4` - URM-P4
	// * `urm-p8` - URM-P8
	// * `splice` - Splice
	// * `other` - Other
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `8p8c` - 8P8C
	// * `8p6c` - 8P6C
	// * `8p4c` - 8P4C
	// * `8p2c` - 8P2C
	// * `6p6c` - 6P6C
	// * `6p4c` - 6P4C
	// * `6p2c` - 6P2C
	// * `4p4c` - 4P4C
	// * `4p2c` - 4P2C
	// * `gg45` - GG45
	// * `tera-4p` - TERA 4P
	// * `tera-2p` - TERA 2P
	// * `tera-1p` - TERA 1P
	// * `110-punch` - 110 Punch
	// * `bnc` - BNC
	// * `f` - F Connector
	// * `n` - N Connector
	// * `mrj21` - MRJ21
	// * `fc` - FC
	// * `lc` - LC
	// * `lc-pc` - LC/PC
	// * `lc-upc` - LC/UPC
	// * `lc-apc` - LC/APC
	// * `lsh` - LSH
	// * `lsh-pc` - LSH/PC
	// * `lsh-upc` - LSH/UPC
	// * `lsh-apc` - LSH/APC
	// * `mpo` - MPO
	// * `mtrj` - MTRJ
	// * `sc` - SC
	// * `sc-pc` - SC/PC
	// * `sc-upc` - SC/UPC
	// * `sc-apc` - SC/APC
	// * `st` - ST
	// * `cs` - CS
	// * `sn` - SN
	// * `sma-905` - SMA 905
	// * `sma-906` - SMA 906
	// * `urm-p2` - URM-P2
	// * `urm-p4` - URM-P4
	// * `urm-p8` - URM-P8
	// * `splice` - Splice
	// * `other` - Other
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
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
func (m DcimFrontPortsListQueryParameters) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color,
		),
		"colorEmpty": validation.Validate(
			m.ColorEmpty,
		),
		"colorIc": validation.Validate(
			m.ColorIc,
		),
		"colorIe": validation.Validate(
			m.ColorIe,
		),
		"colorIew": validation.Validate(
			m.ColorIew,
		),
		"colorIsw": validation.Validate(
			m.ColorIsw,
		),
		"colorN": validation.Validate(
			m.ColorN,
		),
		"colorNic": validation.Validate(
			m.ColorNic,
		),
		"colorNie": validation.Validate(
			m.ColorNie,
		),
		"colorNiew": validation.Validate(
			m.ColorNiew,
		),
		"colorNisw": validation.Validate(
			m.ColorNisw,
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
		"description": validation.Validate(
			m.Description,
		),
		"descriptionEmpty": validation.Validate(
			m.DescriptionEmpty,
		),
		"descriptionIc": validation.Validate(
			m.DescriptionIc,
		),
		"descriptionIe": validation.Validate(
			m.DescriptionIe,
		),
		"descriptionIew": validation.Validate(
			m.DescriptionIew,
		),
		"descriptionIsw": validation.Validate(
			m.DescriptionIsw,
		),
		"descriptionN": validation.Validate(
			m.DescriptionN,
		),
		"descriptionNic": validation.Validate(
			m.DescriptionNic,
		),
		"descriptionNie": validation.Validate(
			m.DescriptionNie,
		),
		"descriptionNiew": validation.Validate(
			m.DescriptionNiew,
		),
		"descriptionNisw": validation.Validate(
			m.DescriptionNisw,
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
		"moduleId": validation.Validate(
			m.ModuleId,
		),
		"moduleIdN": validation.Validate(
			m.ModuleIdN,
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
		"type": validation.Validate(
			m.Type,
		),
		"typeN": validation.Validate(
			m.TypeN,
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

// GetCableEnd returns the CableEnd property
func (m DcimFrontPortsListQueryParameters) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *DcimFrontPortsListQueryParameters) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCableEndN returns the CableEndN property
func (m DcimFrontPortsListQueryParameters) GetCableEndN() string {
	return m.CableEndN
}

// SetCableEndN sets the CableEndN property
func (m *DcimFrontPortsListQueryParameters) SetCableEndN(val string) {
	m.CableEndN = val
}

// GetCabled returns the Cabled property
func (m DcimFrontPortsListQueryParameters) GetCabled() bool {
	return m.Cabled
}

// SetCabled sets the Cabled property
func (m *DcimFrontPortsListQueryParameters) SetCabled(val bool) {
	m.Cabled = val
}

// GetColor returns the Color property
func (m DcimFrontPortsListQueryParameters) GetColor() []string {
	return m.Color
}

// SetColor sets the Color property
func (m *DcimFrontPortsListQueryParameters) SetColor(val []string) {
	m.Color = val
}

// GetColorEmpty returns the ColorEmpty property
func (m DcimFrontPortsListQueryParameters) GetColorEmpty() []string {
	return m.ColorEmpty
}

// SetColorEmpty sets the ColorEmpty property
func (m *DcimFrontPortsListQueryParameters) SetColorEmpty(val []string) {
	m.ColorEmpty = val
}

// GetColorIc returns the ColorIc property
func (m DcimFrontPortsListQueryParameters) GetColorIc() []string {
	return m.ColorIc
}

// SetColorIc sets the ColorIc property
func (m *DcimFrontPortsListQueryParameters) SetColorIc(val []string) {
	m.ColorIc = val
}

// GetColorIe returns the ColorIe property
func (m DcimFrontPortsListQueryParameters) GetColorIe() []string {
	return m.ColorIe
}

// SetColorIe sets the ColorIe property
func (m *DcimFrontPortsListQueryParameters) SetColorIe(val []string) {
	m.ColorIe = val
}

// GetColorIew returns the ColorIew property
func (m DcimFrontPortsListQueryParameters) GetColorIew() []string {
	return m.ColorIew
}

// SetColorIew sets the ColorIew property
func (m *DcimFrontPortsListQueryParameters) SetColorIew(val []string) {
	m.ColorIew = val
}

// GetColorIsw returns the ColorIsw property
func (m DcimFrontPortsListQueryParameters) GetColorIsw() []string {
	return m.ColorIsw
}

// SetColorIsw sets the ColorIsw property
func (m *DcimFrontPortsListQueryParameters) SetColorIsw(val []string) {
	m.ColorIsw = val
}

// GetColorN returns the ColorN property
func (m DcimFrontPortsListQueryParameters) GetColorN() []string {
	return m.ColorN
}

// SetColorN sets the ColorN property
func (m *DcimFrontPortsListQueryParameters) SetColorN(val []string) {
	m.ColorN = val
}

// GetColorNic returns the ColorNic property
func (m DcimFrontPortsListQueryParameters) GetColorNic() []string {
	return m.ColorNic
}

// SetColorNic sets the ColorNic property
func (m *DcimFrontPortsListQueryParameters) SetColorNic(val []string) {
	m.ColorNic = val
}

// GetColorNie returns the ColorNie property
func (m DcimFrontPortsListQueryParameters) GetColorNie() []string {
	return m.ColorNie
}

// SetColorNie sets the ColorNie property
func (m *DcimFrontPortsListQueryParameters) SetColorNie(val []string) {
	m.ColorNie = val
}

// GetColorNiew returns the ColorNiew property
func (m DcimFrontPortsListQueryParameters) GetColorNiew() []string {
	return m.ColorNiew
}

// SetColorNiew sets the ColorNiew property
func (m *DcimFrontPortsListQueryParameters) SetColorNiew(val []string) {
	m.ColorNiew = val
}

// GetColorNisw returns the ColorNisw property
func (m DcimFrontPortsListQueryParameters) GetColorNisw() []string {
	return m.ColorNisw
}

// SetColorNisw sets the ColorNisw property
func (m *DcimFrontPortsListQueryParameters) SetColorNisw(val []string) {
	m.ColorNisw = val
}

// GetCreated returns the Created property
func (m DcimFrontPortsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimFrontPortsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimFrontPortsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimFrontPortsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimFrontPortsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimFrontPortsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimFrontPortsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimFrontPortsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimFrontPortsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimFrontPortsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimFrontPortsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimFrontPortsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimFrontPortsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimFrontPortsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m DcimFrontPortsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DcimFrontPortsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m DcimFrontPortsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m DcimFrontPortsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m DcimFrontPortsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m DcimFrontPortsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m DcimFrontPortsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m DcimFrontPortsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m DcimFrontPortsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m DcimFrontPortsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m DcimFrontPortsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m DcimFrontPortsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *DcimFrontPortsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetDevice returns the Device property
func (m DcimFrontPortsListQueryParameters) GetDevice() []*string {
	return m.Device
}

// SetDevice sets the Device property
func (m *DcimFrontPortsListQueryParameters) SetDevice(val []*string) {
	m.Device = val
}

// GetDeviceN returns the DeviceN property
func (m DcimFrontPortsListQueryParameters) GetDeviceN() []*string {
	return m.DeviceN
}

// SetDeviceN sets the DeviceN property
func (m *DcimFrontPortsListQueryParameters) SetDeviceN(val []*string) {
	m.DeviceN = val
}

// GetDeviceId returns the DeviceId property
func (m DcimFrontPortsListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *DcimFrontPortsListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetDeviceIdN returns the DeviceIdN property
func (m DcimFrontPortsListQueryParameters) GetDeviceIdN() []int32 {
	return m.DeviceIdN
}

// SetDeviceIdN sets the DeviceIdN property
func (m *DcimFrontPortsListQueryParameters) SetDeviceIdN(val []int32) {
	m.DeviceIdN = val
}

// GetId returns the Id property
func (m DcimFrontPortsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimFrontPortsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimFrontPortsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimFrontPortsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimFrontPortsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimFrontPortsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimFrontPortsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimFrontPortsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimFrontPortsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimFrontPortsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimFrontPortsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimFrontPortsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLabel returns the Label property
func (m DcimFrontPortsListQueryParameters) GetLabel() []string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DcimFrontPortsListQueryParameters) SetLabel(val []string) {
	m.Label = val
}

// GetLabelEmpty returns the LabelEmpty property
func (m DcimFrontPortsListQueryParameters) GetLabelEmpty() []string {
	return m.LabelEmpty
}

// SetLabelEmpty sets the LabelEmpty property
func (m *DcimFrontPortsListQueryParameters) SetLabelEmpty(val []string) {
	m.LabelEmpty = val
}

// GetLabelIc returns the LabelIc property
func (m DcimFrontPortsListQueryParameters) GetLabelIc() []string {
	return m.LabelIc
}

// SetLabelIc sets the LabelIc property
func (m *DcimFrontPortsListQueryParameters) SetLabelIc(val []string) {
	m.LabelIc = val
}

// GetLabelIe returns the LabelIe property
func (m DcimFrontPortsListQueryParameters) GetLabelIe() []string {
	return m.LabelIe
}

// SetLabelIe sets the LabelIe property
func (m *DcimFrontPortsListQueryParameters) SetLabelIe(val []string) {
	m.LabelIe = val
}

// GetLabelIew returns the LabelIew property
func (m DcimFrontPortsListQueryParameters) GetLabelIew() []string {
	return m.LabelIew
}

// SetLabelIew sets the LabelIew property
func (m *DcimFrontPortsListQueryParameters) SetLabelIew(val []string) {
	m.LabelIew = val
}

// GetLabelIsw returns the LabelIsw property
func (m DcimFrontPortsListQueryParameters) GetLabelIsw() []string {
	return m.LabelIsw
}

// SetLabelIsw sets the LabelIsw property
func (m *DcimFrontPortsListQueryParameters) SetLabelIsw(val []string) {
	m.LabelIsw = val
}

// GetLabelN returns the LabelN property
func (m DcimFrontPortsListQueryParameters) GetLabelN() []string {
	return m.LabelN
}

// SetLabelN sets the LabelN property
func (m *DcimFrontPortsListQueryParameters) SetLabelN(val []string) {
	m.LabelN = val
}

// GetLabelNic returns the LabelNic property
func (m DcimFrontPortsListQueryParameters) GetLabelNic() []string {
	return m.LabelNic
}

// SetLabelNic sets the LabelNic property
func (m *DcimFrontPortsListQueryParameters) SetLabelNic(val []string) {
	m.LabelNic = val
}

// GetLabelNie returns the LabelNie property
func (m DcimFrontPortsListQueryParameters) GetLabelNie() []string {
	return m.LabelNie
}

// SetLabelNie sets the LabelNie property
func (m *DcimFrontPortsListQueryParameters) SetLabelNie(val []string) {
	m.LabelNie = val
}

// GetLabelNiew returns the LabelNiew property
func (m DcimFrontPortsListQueryParameters) GetLabelNiew() []string {
	return m.LabelNiew
}

// SetLabelNiew sets the LabelNiew property
func (m *DcimFrontPortsListQueryParameters) SetLabelNiew(val []string) {
	m.LabelNiew = val
}

// GetLabelNisw returns the LabelNisw property
func (m DcimFrontPortsListQueryParameters) GetLabelNisw() []string {
	return m.LabelNisw
}

// SetLabelNisw sets the LabelNisw property
func (m *DcimFrontPortsListQueryParameters) SetLabelNisw(val []string) {
	m.LabelNisw = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimFrontPortsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimFrontPortsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimFrontPortsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimFrontPortsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimFrontPortsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimFrontPortsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimFrontPortsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimFrontPortsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimFrontPortsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimFrontPortsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimFrontPortsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimFrontPortsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimFrontPortsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimFrontPortsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m DcimFrontPortsListQueryParameters) GetLocation() []string {
	return m.Location
}

// SetLocation sets the Location property
func (m *DcimFrontPortsListQueryParameters) SetLocation(val []string) {
	m.Location = val
}

// GetLocationN returns the LocationN property
func (m DcimFrontPortsListQueryParameters) GetLocationN() []string {
	return m.LocationN
}

// SetLocationN sets the LocationN property
func (m *DcimFrontPortsListQueryParameters) SetLocationN(val []string) {
	m.LocationN = val
}

// GetLocationId returns the LocationId property
func (m DcimFrontPortsListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimFrontPortsListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m DcimFrontPortsListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *DcimFrontPortsListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetModuleId returns the ModuleId property
func (m DcimFrontPortsListQueryParameters) GetModuleId() []*int32 {
	return m.ModuleId
}

// SetModuleId sets the ModuleId property
func (m *DcimFrontPortsListQueryParameters) SetModuleId(val []*int32) {
	m.ModuleId = val
}

// GetModuleIdN returns the ModuleIdN property
func (m DcimFrontPortsListQueryParameters) GetModuleIdN() []*int32 {
	return m.ModuleIdN
}

// SetModuleIdN sets the ModuleIdN property
func (m *DcimFrontPortsListQueryParameters) SetModuleIdN(val []*int32) {
	m.ModuleIdN = val
}

// GetName returns the Name property
func (m DcimFrontPortsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimFrontPortsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimFrontPortsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimFrontPortsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimFrontPortsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimFrontPortsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimFrontPortsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimFrontPortsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimFrontPortsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimFrontPortsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimFrontPortsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimFrontPortsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimFrontPortsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimFrontPortsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimFrontPortsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimFrontPortsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimFrontPortsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimFrontPortsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimFrontPortsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimFrontPortsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimFrontPortsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimFrontPortsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOccupied returns the Occupied property
func (m DcimFrontPortsListQueryParameters) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *DcimFrontPortsListQueryParameters) SetOccupied(val bool) {
	m.Occupied = val
}

// GetOffset returns the Offset property
func (m DcimFrontPortsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimFrontPortsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimFrontPortsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimFrontPortsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimFrontPortsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimFrontPortsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRack returns the Rack property
func (m DcimFrontPortsListQueryParameters) GetRack() []string {
	return m.Rack
}

// SetRack sets the Rack property
func (m *DcimFrontPortsListQueryParameters) SetRack(val []string) {
	m.Rack = val
}

// GetRackN returns the RackN property
func (m DcimFrontPortsListQueryParameters) GetRackN() []string {
	return m.RackN
}

// SetRackN sets the RackN property
func (m *DcimFrontPortsListQueryParameters) SetRackN(val []string) {
	m.RackN = val
}

// GetRackId returns the RackId property
func (m DcimFrontPortsListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimFrontPortsListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetRackIdN returns the RackIdN property
func (m DcimFrontPortsListQueryParameters) GetRackIdN() []int32 {
	return m.RackIdN
}

// SetRackIdN sets the RackIdN property
func (m *DcimFrontPortsListQueryParameters) SetRackIdN(val []int32) {
	m.RackIdN = val
}

// GetRegion returns the Region property
func (m DcimFrontPortsListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimFrontPortsListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimFrontPortsListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimFrontPortsListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimFrontPortsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimFrontPortsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimFrontPortsListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimFrontPortsListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSite returns the Site property
func (m DcimFrontPortsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimFrontPortsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimFrontPortsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimFrontPortsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimFrontPortsListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimFrontPortsListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimFrontPortsListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimFrontPortsListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimFrontPortsListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimFrontPortsListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimFrontPortsListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimFrontPortsListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimFrontPortsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimFrontPortsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimFrontPortsListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimFrontPortsListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetTag returns the Tag property
func (m DcimFrontPortsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimFrontPortsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimFrontPortsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimFrontPortsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetType returns the Type property
func (m DcimFrontPortsListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimFrontPortsListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimFrontPortsListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimFrontPortsListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimFrontPortsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimFrontPortsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualChassis returns the VirtualChassis property
func (m DcimFrontPortsListQueryParameters) GetVirtualChassis() []string {
	return m.VirtualChassis
}

// SetVirtualChassis sets the VirtualChassis property
func (m *DcimFrontPortsListQueryParameters) SetVirtualChassis(val []string) {
	m.VirtualChassis = val
}

// GetVirtualChassisN returns the VirtualChassisN property
func (m DcimFrontPortsListQueryParameters) GetVirtualChassisN() []string {
	return m.VirtualChassisN
}

// SetVirtualChassisN sets the VirtualChassisN property
func (m *DcimFrontPortsListQueryParameters) SetVirtualChassisN(val []string) {
	m.VirtualChassisN = val
}

// GetVirtualChassisId returns the VirtualChassisId property
func (m DcimFrontPortsListQueryParameters) GetVirtualChassisId() []int32 {
	return m.VirtualChassisId
}

// SetVirtualChassisId sets the VirtualChassisId property
func (m *DcimFrontPortsListQueryParameters) SetVirtualChassisId(val []int32) {
	m.VirtualChassisId = val
}

// GetVirtualChassisIdN returns the VirtualChassisIdN property
func (m DcimFrontPortsListQueryParameters) GetVirtualChassisIdN() []int32 {
	return m.VirtualChassisIdN
}

// SetVirtualChassisIdN sets the VirtualChassisIdN property
func (m *DcimFrontPortsListQueryParameters) SetVirtualChassisIdN(val []int32) {
	m.VirtualChassisIdN = val
}
