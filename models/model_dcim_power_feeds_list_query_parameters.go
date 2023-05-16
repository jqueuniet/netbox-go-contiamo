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

// DcimPowerFeedsListQueryParameters is an object.
type DcimPowerFeedsListQueryParameters struct {
	// Amperage:
	Amperage []int32 `json:"amperage,omitempty" mapstructure:"amperage,omitempty"`
	// AmperageGt:
	AmperageGt []int32 `json:"amperage__gt,omitempty" mapstructure:"amperage__gt,omitempty"`
	// AmperageGte:
	AmperageGte []int32 `json:"amperage__gte,omitempty" mapstructure:"amperage__gte,omitempty"`
	// AmperageLt:
	AmperageLt []int32 `json:"amperage__lt,omitempty" mapstructure:"amperage__lt,omitempty"`
	// AmperageLte:
	AmperageLte []int32 `json:"amperage__lte,omitempty" mapstructure:"amperage__lte,omitempty"`
	// AmperageN:
	AmperageN []int32 `json:"amperage__n,omitempty" mapstructure:"amperage__n,omitempty"`
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
	// MaxUtilization:
	MaxUtilization []int32 `json:"max_utilization,omitempty" mapstructure:"max_utilization,omitempty"`
	// MaxUtilizationGt:
	MaxUtilizationGt []int32 `json:"max_utilization__gt,omitempty" mapstructure:"max_utilization__gt,omitempty"`
	// MaxUtilizationGte:
	MaxUtilizationGte []int32 `json:"max_utilization__gte,omitempty" mapstructure:"max_utilization__gte,omitempty"`
	// MaxUtilizationLt:
	MaxUtilizationLt []int32 `json:"max_utilization__lt,omitempty" mapstructure:"max_utilization__lt,omitempty"`
	// MaxUtilizationLte:
	MaxUtilizationLte []int32 `json:"max_utilization__lte,omitempty" mapstructure:"max_utilization__lte,omitempty"`
	// MaxUtilizationN:
	MaxUtilizationN []int32 `json:"max_utilization__n,omitempty" mapstructure:"max_utilization__n,omitempty"`
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
	// Phase: * `single-phase` - Single phase
	// * `three-phase` - Three-phase
	Phase string `json:"phase,omitempty" mapstructure:"phase,omitempty"`
	// PhaseN: * `single-phase` - Single phase
	// * `three-phase` - Three-phase
	PhaseN string `json:"phase__n,omitempty" mapstructure:"phase__n,omitempty"`
	// PowerPanelId: Power panel (ID)
	PowerPanelId []int32 `json:"power_panel_id,omitempty" mapstructure:"power_panel_id,omitempty"`
	// PowerPanelIdN: Power panel (ID)
	PowerPanelIdN []int32 `json:"power_panel_id__n,omitempty" mapstructure:"power_panel_id__n,omitempty"`
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
	// * `failed` - Failed
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `failed` - Failed
	StatusN []string `json:"status__n,omitempty" mapstructure:"status__n,omitempty"`
	// Supply: * `ac` - AC
	// * `dc` - DC
	Supply string `json:"supply,omitempty" mapstructure:"supply,omitempty"`
	// SupplyN: * `ac` - AC
	// * `dc` - DC
	SupplyN string `json:"supply__n,omitempty" mapstructure:"supply__n,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// Type: * `primary` - Primary
	// * `redundant` - Redundant
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `primary` - Primary
	// * `redundant` - Redundant
	TypeN string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// Voltage:
	Voltage []int32 `json:"voltage,omitempty" mapstructure:"voltage,omitempty"`
	// VoltageGt:
	VoltageGt []int32 `json:"voltage__gt,omitempty" mapstructure:"voltage__gt,omitempty"`
	// VoltageGte:
	VoltageGte []int32 `json:"voltage__gte,omitempty" mapstructure:"voltage__gte,omitempty"`
	// VoltageLt:
	VoltageLt []int32 `json:"voltage__lt,omitempty" mapstructure:"voltage__lt,omitempty"`
	// VoltageLte:
	VoltageLte []int32 `json:"voltage__lte,omitempty" mapstructure:"voltage__lte,omitempty"`
	// VoltageN:
	VoltageN []int32 `json:"voltage__n,omitempty" mapstructure:"voltage__n,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimPowerFeedsListQueryParameters) Validate() error {
	return validation.Errors{
		"amperage": validation.Validate(
			m.Amperage,
		),
		"amperageGt": validation.Validate(
			m.AmperageGt,
		),
		"amperageGte": validation.Validate(
			m.AmperageGte,
		),
		"amperageLt": validation.Validate(
			m.AmperageLt,
		),
		"amperageLte": validation.Validate(
			m.AmperageLte,
		),
		"amperageN": validation.Validate(
			m.AmperageN,
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
		"maxUtilization": validation.Validate(
			m.MaxUtilization,
		),
		"maxUtilizationGt": validation.Validate(
			m.MaxUtilizationGt,
		),
		"maxUtilizationGte": validation.Validate(
			m.MaxUtilizationGte,
		),
		"maxUtilizationLt": validation.Validate(
			m.MaxUtilizationLt,
		),
		"maxUtilizationLte": validation.Validate(
			m.MaxUtilizationLte,
		),
		"maxUtilizationN": validation.Validate(
			m.MaxUtilizationN,
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
		"powerPanelId": validation.Validate(
			m.PowerPanelId,
		),
		"powerPanelIdN": validation.Validate(
			m.PowerPanelIdN,
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
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"voltage": validation.Validate(
			m.Voltage,
		),
		"voltageGt": validation.Validate(
			m.VoltageGt,
		),
		"voltageGte": validation.Validate(
			m.VoltageGte,
		),
		"voltageLt": validation.Validate(
			m.VoltageLt,
		),
		"voltageLte": validation.Validate(
			m.VoltageLte,
		),
		"voltageN": validation.Validate(
			m.VoltageN,
		),
	}.Filter()
}

// GetAmperage returns the Amperage property
func (m DcimPowerFeedsListQueryParameters) GetAmperage() []int32 {
	return m.Amperage
}

// SetAmperage sets the Amperage property
func (m *DcimPowerFeedsListQueryParameters) SetAmperage(val []int32) {
	m.Amperage = val
}

// GetAmperageGt returns the AmperageGt property
func (m DcimPowerFeedsListQueryParameters) GetAmperageGt() []int32 {
	return m.AmperageGt
}

// SetAmperageGt sets the AmperageGt property
func (m *DcimPowerFeedsListQueryParameters) SetAmperageGt(val []int32) {
	m.AmperageGt = val
}

// GetAmperageGte returns the AmperageGte property
func (m DcimPowerFeedsListQueryParameters) GetAmperageGte() []int32 {
	return m.AmperageGte
}

// SetAmperageGte sets the AmperageGte property
func (m *DcimPowerFeedsListQueryParameters) SetAmperageGte(val []int32) {
	m.AmperageGte = val
}

// GetAmperageLt returns the AmperageLt property
func (m DcimPowerFeedsListQueryParameters) GetAmperageLt() []int32 {
	return m.AmperageLt
}

// SetAmperageLt sets the AmperageLt property
func (m *DcimPowerFeedsListQueryParameters) SetAmperageLt(val []int32) {
	m.AmperageLt = val
}

// GetAmperageLte returns the AmperageLte property
func (m DcimPowerFeedsListQueryParameters) GetAmperageLte() []int32 {
	return m.AmperageLte
}

// SetAmperageLte sets the AmperageLte property
func (m *DcimPowerFeedsListQueryParameters) SetAmperageLte(val []int32) {
	m.AmperageLte = val
}

// GetAmperageN returns the AmperageN property
func (m DcimPowerFeedsListQueryParameters) GetAmperageN() []int32 {
	return m.AmperageN
}

// SetAmperageN sets the AmperageN property
func (m *DcimPowerFeedsListQueryParameters) SetAmperageN(val []int32) {
	m.AmperageN = val
}

// GetCableEnd returns the CableEnd property
func (m DcimPowerFeedsListQueryParameters) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *DcimPowerFeedsListQueryParameters) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCableEndN returns the CableEndN property
func (m DcimPowerFeedsListQueryParameters) GetCableEndN() string {
	return m.CableEndN
}

// SetCableEndN sets the CableEndN property
func (m *DcimPowerFeedsListQueryParameters) SetCableEndN(val string) {
	m.CableEndN = val
}

// GetCabled returns the Cabled property
func (m DcimPowerFeedsListQueryParameters) GetCabled() bool {
	return m.Cabled
}

// SetCabled sets the Cabled property
func (m *DcimPowerFeedsListQueryParameters) SetCabled(val bool) {
	m.Cabled = val
}

// GetConnected returns the Connected property
func (m DcimPowerFeedsListQueryParameters) GetConnected() bool {
	return m.Connected
}

// SetConnected sets the Connected property
func (m *DcimPowerFeedsListQueryParameters) SetConnected(val bool) {
	m.Connected = val
}

// GetCreated returns the Created property
func (m DcimPowerFeedsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimPowerFeedsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimPowerFeedsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimPowerFeedsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimPowerFeedsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimPowerFeedsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimPowerFeedsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimPowerFeedsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimPowerFeedsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimPowerFeedsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimPowerFeedsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimPowerFeedsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimPowerFeedsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimPowerFeedsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetId returns the Id property
func (m DcimPowerFeedsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimPowerFeedsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimPowerFeedsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimPowerFeedsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimPowerFeedsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimPowerFeedsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimPowerFeedsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimPowerFeedsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimPowerFeedsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimPowerFeedsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimPowerFeedsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimPowerFeedsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimPowerFeedsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimPowerFeedsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimPowerFeedsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimPowerFeedsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimPowerFeedsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimPowerFeedsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimPowerFeedsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimPowerFeedsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimPowerFeedsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimPowerFeedsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimPowerFeedsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimPowerFeedsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimPowerFeedsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimPowerFeedsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetMaxUtilization returns the MaxUtilization property
func (m DcimPowerFeedsListQueryParameters) GetMaxUtilization() []int32 {
	return m.MaxUtilization
}

// SetMaxUtilization sets the MaxUtilization property
func (m *DcimPowerFeedsListQueryParameters) SetMaxUtilization(val []int32) {
	m.MaxUtilization = val
}

// GetMaxUtilizationGt returns the MaxUtilizationGt property
func (m DcimPowerFeedsListQueryParameters) GetMaxUtilizationGt() []int32 {
	return m.MaxUtilizationGt
}

// SetMaxUtilizationGt sets the MaxUtilizationGt property
func (m *DcimPowerFeedsListQueryParameters) SetMaxUtilizationGt(val []int32) {
	m.MaxUtilizationGt = val
}

// GetMaxUtilizationGte returns the MaxUtilizationGte property
func (m DcimPowerFeedsListQueryParameters) GetMaxUtilizationGte() []int32 {
	return m.MaxUtilizationGte
}

// SetMaxUtilizationGte sets the MaxUtilizationGte property
func (m *DcimPowerFeedsListQueryParameters) SetMaxUtilizationGte(val []int32) {
	m.MaxUtilizationGte = val
}

// GetMaxUtilizationLt returns the MaxUtilizationLt property
func (m DcimPowerFeedsListQueryParameters) GetMaxUtilizationLt() []int32 {
	return m.MaxUtilizationLt
}

// SetMaxUtilizationLt sets the MaxUtilizationLt property
func (m *DcimPowerFeedsListQueryParameters) SetMaxUtilizationLt(val []int32) {
	m.MaxUtilizationLt = val
}

// GetMaxUtilizationLte returns the MaxUtilizationLte property
func (m DcimPowerFeedsListQueryParameters) GetMaxUtilizationLte() []int32 {
	return m.MaxUtilizationLte
}

// SetMaxUtilizationLte sets the MaxUtilizationLte property
func (m *DcimPowerFeedsListQueryParameters) SetMaxUtilizationLte(val []int32) {
	m.MaxUtilizationLte = val
}

// GetMaxUtilizationN returns the MaxUtilizationN property
func (m DcimPowerFeedsListQueryParameters) GetMaxUtilizationN() []int32 {
	return m.MaxUtilizationN
}

// SetMaxUtilizationN sets the MaxUtilizationN property
func (m *DcimPowerFeedsListQueryParameters) SetMaxUtilizationN(val []int32) {
	m.MaxUtilizationN = val
}

// GetName returns the Name property
func (m DcimPowerFeedsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimPowerFeedsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimPowerFeedsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimPowerFeedsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimPowerFeedsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimPowerFeedsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimPowerFeedsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimPowerFeedsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimPowerFeedsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimPowerFeedsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimPowerFeedsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimPowerFeedsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimPowerFeedsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimPowerFeedsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimPowerFeedsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimPowerFeedsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimPowerFeedsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimPowerFeedsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimPowerFeedsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimPowerFeedsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimPowerFeedsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimPowerFeedsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOccupied returns the Occupied property
func (m DcimPowerFeedsListQueryParameters) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *DcimPowerFeedsListQueryParameters) SetOccupied(val bool) {
	m.Occupied = val
}

// GetOffset returns the Offset property
func (m DcimPowerFeedsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimPowerFeedsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimPowerFeedsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimPowerFeedsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPhase returns the Phase property
func (m DcimPowerFeedsListQueryParameters) GetPhase() string {
	return m.Phase
}

// SetPhase sets the Phase property
func (m *DcimPowerFeedsListQueryParameters) SetPhase(val string) {
	m.Phase = val
}

// GetPhaseN returns the PhaseN property
func (m DcimPowerFeedsListQueryParameters) GetPhaseN() string {
	return m.PhaseN
}

// SetPhaseN sets the PhaseN property
func (m *DcimPowerFeedsListQueryParameters) SetPhaseN(val string) {
	m.PhaseN = val
}

// GetPowerPanelId returns the PowerPanelId property
func (m DcimPowerFeedsListQueryParameters) GetPowerPanelId() []int32 {
	return m.PowerPanelId
}

// SetPowerPanelId sets the PowerPanelId property
func (m *DcimPowerFeedsListQueryParameters) SetPowerPanelId(val []int32) {
	m.PowerPanelId = val
}

// GetPowerPanelIdN returns the PowerPanelIdN property
func (m DcimPowerFeedsListQueryParameters) GetPowerPanelIdN() []int32 {
	return m.PowerPanelIdN
}

// SetPowerPanelIdN sets the PowerPanelIdN property
func (m *DcimPowerFeedsListQueryParameters) SetPowerPanelIdN(val []int32) {
	m.PowerPanelIdN = val
}

// GetQ returns the Q property
func (m DcimPowerFeedsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimPowerFeedsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRackId returns the RackId property
func (m DcimPowerFeedsListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimPowerFeedsListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetRackIdN returns the RackIdN property
func (m DcimPowerFeedsListQueryParameters) GetRackIdN() []int32 {
	return m.RackIdN
}

// SetRackIdN sets the RackIdN property
func (m *DcimPowerFeedsListQueryParameters) SetRackIdN(val []int32) {
	m.RackIdN = val
}

// GetRegion returns the Region property
func (m DcimPowerFeedsListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimPowerFeedsListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimPowerFeedsListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimPowerFeedsListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimPowerFeedsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimPowerFeedsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimPowerFeedsListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimPowerFeedsListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSite returns the Site property
func (m DcimPowerFeedsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimPowerFeedsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimPowerFeedsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimPowerFeedsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimPowerFeedsListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimPowerFeedsListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimPowerFeedsListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimPowerFeedsListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimPowerFeedsListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimPowerFeedsListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimPowerFeedsListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimPowerFeedsListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimPowerFeedsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimPowerFeedsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimPowerFeedsListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimPowerFeedsListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetStatus returns the Status property
func (m DcimPowerFeedsListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *DcimPowerFeedsListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m DcimPowerFeedsListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *DcimPowerFeedsListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetSupply returns the Supply property
func (m DcimPowerFeedsListQueryParameters) GetSupply() string {
	return m.Supply
}

// SetSupply sets the Supply property
func (m *DcimPowerFeedsListQueryParameters) SetSupply(val string) {
	m.Supply = val
}

// GetSupplyN returns the SupplyN property
func (m DcimPowerFeedsListQueryParameters) GetSupplyN() string {
	return m.SupplyN
}

// SetSupplyN sets the SupplyN property
func (m *DcimPowerFeedsListQueryParameters) SetSupplyN(val string) {
	m.SupplyN = val
}

// GetTag returns the Tag property
func (m DcimPowerFeedsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimPowerFeedsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimPowerFeedsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimPowerFeedsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetType returns the Type property
func (m DcimPowerFeedsListQueryParameters) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimPowerFeedsListQueryParameters) SetType(val string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimPowerFeedsListQueryParameters) GetTypeN() string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimPowerFeedsListQueryParameters) SetTypeN(val string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimPowerFeedsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimPowerFeedsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVoltage returns the Voltage property
func (m DcimPowerFeedsListQueryParameters) GetVoltage() []int32 {
	return m.Voltage
}

// SetVoltage sets the Voltage property
func (m *DcimPowerFeedsListQueryParameters) SetVoltage(val []int32) {
	m.Voltage = val
}

// GetVoltageGt returns the VoltageGt property
func (m DcimPowerFeedsListQueryParameters) GetVoltageGt() []int32 {
	return m.VoltageGt
}

// SetVoltageGt sets the VoltageGt property
func (m *DcimPowerFeedsListQueryParameters) SetVoltageGt(val []int32) {
	m.VoltageGt = val
}

// GetVoltageGte returns the VoltageGte property
func (m DcimPowerFeedsListQueryParameters) GetVoltageGte() []int32 {
	return m.VoltageGte
}

// SetVoltageGte sets the VoltageGte property
func (m *DcimPowerFeedsListQueryParameters) SetVoltageGte(val []int32) {
	m.VoltageGte = val
}

// GetVoltageLt returns the VoltageLt property
func (m DcimPowerFeedsListQueryParameters) GetVoltageLt() []int32 {
	return m.VoltageLt
}

// SetVoltageLt sets the VoltageLt property
func (m *DcimPowerFeedsListQueryParameters) SetVoltageLt(val []int32) {
	m.VoltageLt = val
}

// GetVoltageLte returns the VoltageLte property
func (m DcimPowerFeedsListQueryParameters) GetVoltageLte() []int32 {
	return m.VoltageLte
}

// SetVoltageLte sets the VoltageLte property
func (m *DcimPowerFeedsListQueryParameters) SetVoltageLte(val []int32) {
	m.VoltageLte = val
}

// GetVoltageN returns the VoltageN property
func (m DcimPowerFeedsListQueryParameters) GetVoltageN() []int32 {
	return m.VoltageN
}

// SetVoltageN sets the VoltageN property
func (m *DcimPowerFeedsListQueryParameters) SetVoltageN(val []int32) {
	m.VoltageN = val
}
