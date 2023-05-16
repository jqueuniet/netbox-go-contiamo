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

// DcimConsoleServerPortsListQueryParameters is an object.
type DcimConsoleServerPortsListQueryParameters struct {
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end,omitempty" mapstructure:"cable_end,omitempty"`
	// CableEndN: * `A` - A
	// * `B` - B
	CableEndN string `json:"cable_end__n,omitempty" mapstructure:"cable_end__n,omitempty"`
	// Cabled:
	Cabled bool `json:"cabled,omitempty" mapstructure:"cabled,omitempty"`
	// Connected:
	Connected bool `json:"connected,omitempty" mapstructure:"connected,omitempty"`
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
	// Type: Physical port type
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: Physical port type
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
func (m DcimConsoleServerPortsListQueryParameters) Validate() error {
	return validation.Errors{
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
func (m DcimConsoleServerPortsListQueryParameters) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *DcimConsoleServerPortsListQueryParameters) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCableEndN returns the CableEndN property
func (m DcimConsoleServerPortsListQueryParameters) GetCableEndN() string {
	return m.CableEndN
}

// SetCableEndN sets the CableEndN property
func (m *DcimConsoleServerPortsListQueryParameters) SetCableEndN(val string) {
	m.CableEndN = val
}

// GetCabled returns the Cabled property
func (m DcimConsoleServerPortsListQueryParameters) GetCabled() bool {
	return m.Cabled
}

// SetCabled sets the Cabled property
func (m *DcimConsoleServerPortsListQueryParameters) SetCabled(val bool) {
	m.Cabled = val
}

// GetConnected returns the Connected property
func (m DcimConsoleServerPortsListQueryParameters) GetConnected() bool {
	return m.Connected
}

// SetConnected sets the Connected property
func (m *DcimConsoleServerPortsListQueryParameters) SetConnected(val bool) {
	m.Connected = val
}

// GetCreated returns the Created property
func (m DcimConsoleServerPortsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimConsoleServerPortsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimConsoleServerPortsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimConsoleServerPortsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimConsoleServerPortsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimConsoleServerPortsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimConsoleServerPortsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimConsoleServerPortsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimConsoleServerPortsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimConsoleServerPortsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimConsoleServerPortsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimConsoleServerPortsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimConsoleServerPortsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimConsoleServerPortsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m DcimConsoleServerPortsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m DcimConsoleServerPortsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *DcimConsoleServerPortsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetDevice returns the Device property
func (m DcimConsoleServerPortsListQueryParameters) GetDevice() []*string {
	return m.Device
}

// SetDevice sets the Device property
func (m *DcimConsoleServerPortsListQueryParameters) SetDevice(val []*string) {
	m.Device = val
}

// GetDeviceN returns the DeviceN property
func (m DcimConsoleServerPortsListQueryParameters) GetDeviceN() []*string {
	return m.DeviceN
}

// SetDeviceN sets the DeviceN property
func (m *DcimConsoleServerPortsListQueryParameters) SetDeviceN(val []*string) {
	m.DeviceN = val
}

// GetDeviceId returns the DeviceId property
func (m DcimConsoleServerPortsListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *DcimConsoleServerPortsListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetDeviceIdN returns the DeviceIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetDeviceIdN() []int32 {
	return m.DeviceIdN
}

// SetDeviceIdN sets the DeviceIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetDeviceIdN(val []int32) {
	m.DeviceIdN = val
}

// GetId returns the Id property
func (m DcimConsoleServerPortsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimConsoleServerPortsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimConsoleServerPortsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimConsoleServerPortsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimConsoleServerPortsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimConsoleServerPortsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimConsoleServerPortsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimConsoleServerPortsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimConsoleServerPortsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimConsoleServerPortsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimConsoleServerPortsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLabel returns the Label property
func (m DcimConsoleServerPortsListQueryParameters) GetLabel() []string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabel(val []string) {
	m.Label = val
}

// GetLabelEmpty returns the LabelEmpty property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelEmpty() []string {
	return m.LabelEmpty
}

// SetLabelEmpty sets the LabelEmpty property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelEmpty(val []string) {
	m.LabelEmpty = val
}

// GetLabelIc returns the LabelIc property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelIc() []string {
	return m.LabelIc
}

// SetLabelIc sets the LabelIc property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelIc(val []string) {
	m.LabelIc = val
}

// GetLabelIe returns the LabelIe property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelIe() []string {
	return m.LabelIe
}

// SetLabelIe sets the LabelIe property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelIe(val []string) {
	m.LabelIe = val
}

// GetLabelIew returns the LabelIew property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelIew() []string {
	return m.LabelIew
}

// SetLabelIew sets the LabelIew property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelIew(val []string) {
	m.LabelIew = val
}

// GetLabelIsw returns the LabelIsw property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelIsw() []string {
	return m.LabelIsw
}

// SetLabelIsw sets the LabelIsw property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelIsw(val []string) {
	m.LabelIsw = val
}

// GetLabelN returns the LabelN property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelN() []string {
	return m.LabelN
}

// SetLabelN sets the LabelN property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelN(val []string) {
	m.LabelN = val
}

// GetLabelNic returns the LabelNic property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelNic() []string {
	return m.LabelNic
}

// SetLabelNic sets the LabelNic property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelNic(val []string) {
	m.LabelNic = val
}

// GetLabelNie returns the LabelNie property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelNie() []string {
	return m.LabelNie
}

// SetLabelNie sets the LabelNie property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelNie(val []string) {
	m.LabelNie = val
}

// GetLabelNiew returns the LabelNiew property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelNiew() []string {
	return m.LabelNiew
}

// SetLabelNiew sets the LabelNiew property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelNiew(val []string) {
	m.LabelNiew = val
}

// GetLabelNisw returns the LabelNisw property
func (m DcimConsoleServerPortsListQueryParameters) GetLabelNisw() []string {
	return m.LabelNisw
}

// SetLabelNisw sets the LabelNisw property
func (m *DcimConsoleServerPortsListQueryParameters) SetLabelNisw(val []string) {
	m.LabelNisw = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimConsoleServerPortsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimConsoleServerPortsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimConsoleServerPortsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimConsoleServerPortsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimConsoleServerPortsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimConsoleServerPortsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimConsoleServerPortsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimConsoleServerPortsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimConsoleServerPortsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimConsoleServerPortsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimConsoleServerPortsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimConsoleServerPortsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimConsoleServerPortsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimConsoleServerPortsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m DcimConsoleServerPortsListQueryParameters) GetLocation() []string {
	return m.Location
}

// SetLocation sets the Location property
func (m *DcimConsoleServerPortsListQueryParameters) SetLocation(val []string) {
	m.Location = val
}

// GetLocationN returns the LocationN property
func (m DcimConsoleServerPortsListQueryParameters) GetLocationN() []string {
	return m.LocationN
}

// SetLocationN sets the LocationN property
func (m *DcimConsoleServerPortsListQueryParameters) SetLocationN(val []string) {
	m.LocationN = val
}

// GetLocationId returns the LocationId property
func (m DcimConsoleServerPortsListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimConsoleServerPortsListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetModuleId returns the ModuleId property
func (m DcimConsoleServerPortsListQueryParameters) GetModuleId() []*int32 {
	return m.ModuleId
}

// SetModuleId sets the ModuleId property
func (m *DcimConsoleServerPortsListQueryParameters) SetModuleId(val []*int32) {
	m.ModuleId = val
}

// GetModuleIdN returns the ModuleIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetModuleIdN() []*int32 {
	return m.ModuleIdN
}

// SetModuleIdN sets the ModuleIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetModuleIdN(val []*int32) {
	m.ModuleIdN = val
}

// GetName returns the Name property
func (m DcimConsoleServerPortsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimConsoleServerPortsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimConsoleServerPortsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimConsoleServerPortsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimConsoleServerPortsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimConsoleServerPortsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimConsoleServerPortsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimConsoleServerPortsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimConsoleServerPortsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimConsoleServerPortsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimConsoleServerPortsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimConsoleServerPortsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimConsoleServerPortsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOccupied returns the Occupied property
func (m DcimConsoleServerPortsListQueryParameters) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *DcimConsoleServerPortsListQueryParameters) SetOccupied(val bool) {
	m.Occupied = val
}

// GetOffset returns the Offset property
func (m DcimConsoleServerPortsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimConsoleServerPortsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimConsoleServerPortsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimConsoleServerPortsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimConsoleServerPortsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimConsoleServerPortsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRack returns the Rack property
func (m DcimConsoleServerPortsListQueryParameters) GetRack() []string {
	return m.Rack
}

// SetRack sets the Rack property
func (m *DcimConsoleServerPortsListQueryParameters) SetRack(val []string) {
	m.Rack = val
}

// GetRackN returns the RackN property
func (m DcimConsoleServerPortsListQueryParameters) GetRackN() []string {
	return m.RackN
}

// SetRackN sets the RackN property
func (m *DcimConsoleServerPortsListQueryParameters) SetRackN(val []string) {
	m.RackN = val
}

// GetRackId returns the RackId property
func (m DcimConsoleServerPortsListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimConsoleServerPortsListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetRackIdN returns the RackIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetRackIdN() []int32 {
	return m.RackIdN
}

// SetRackIdN sets the RackIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetRackIdN(val []int32) {
	m.RackIdN = val
}

// GetRegion returns the Region property
func (m DcimConsoleServerPortsListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimConsoleServerPortsListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimConsoleServerPortsListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimConsoleServerPortsListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimConsoleServerPortsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimConsoleServerPortsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSite returns the Site property
func (m DcimConsoleServerPortsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimConsoleServerPortsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimConsoleServerPortsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimConsoleServerPortsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimConsoleServerPortsListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimConsoleServerPortsListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimConsoleServerPortsListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimConsoleServerPortsListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimConsoleServerPortsListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimConsoleServerPortsListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimConsoleServerPortsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimConsoleServerPortsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetTag returns the Tag property
func (m DcimConsoleServerPortsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimConsoleServerPortsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimConsoleServerPortsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimConsoleServerPortsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetType returns the Type property
func (m DcimConsoleServerPortsListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimConsoleServerPortsListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimConsoleServerPortsListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimConsoleServerPortsListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimConsoleServerPortsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimConsoleServerPortsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualChassis returns the VirtualChassis property
func (m DcimConsoleServerPortsListQueryParameters) GetVirtualChassis() []string {
	return m.VirtualChassis
}

// SetVirtualChassis sets the VirtualChassis property
func (m *DcimConsoleServerPortsListQueryParameters) SetVirtualChassis(val []string) {
	m.VirtualChassis = val
}

// GetVirtualChassisN returns the VirtualChassisN property
func (m DcimConsoleServerPortsListQueryParameters) GetVirtualChassisN() []string {
	return m.VirtualChassisN
}

// SetVirtualChassisN sets the VirtualChassisN property
func (m *DcimConsoleServerPortsListQueryParameters) SetVirtualChassisN(val []string) {
	m.VirtualChassisN = val
}

// GetVirtualChassisId returns the VirtualChassisId property
func (m DcimConsoleServerPortsListQueryParameters) GetVirtualChassisId() []int32 {
	return m.VirtualChassisId
}

// SetVirtualChassisId sets the VirtualChassisId property
func (m *DcimConsoleServerPortsListQueryParameters) SetVirtualChassisId(val []int32) {
	m.VirtualChassisId = val
}

// GetVirtualChassisIdN returns the VirtualChassisIdN property
func (m DcimConsoleServerPortsListQueryParameters) GetVirtualChassisIdN() []int32 {
	return m.VirtualChassisIdN
}

// SetVirtualChassisIdN sets the VirtualChassisIdN property
func (m *DcimConsoleServerPortsListQueryParameters) SetVirtualChassisIdN(val []int32) {
	m.VirtualChassisIdN = val
}
