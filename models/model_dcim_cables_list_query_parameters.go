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

// DcimCablesListQueryParameters is an object.
type DcimCablesListQueryParameters struct {
	// Color:
	Color []string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// ColorN:
	ColorN []string `json:"color__n,omitempty" mapstructure:"color__n,omitempty"`
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
	// Device:
	Device []string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceId:
	DeviceId []int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
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
	// Length:
	Length []float64 `json:"length,omitempty" mapstructure:"length,omitempty"`
	// LengthGt:
	LengthGt []float64 `json:"length__gt,omitempty" mapstructure:"length__gt,omitempty"`
	// LengthGte:
	LengthGte []float64 `json:"length__gte,omitempty" mapstructure:"length__gte,omitempty"`
	// LengthLt:
	LengthLt []float64 `json:"length__lt,omitempty" mapstructure:"length__lt,omitempty"`
	// LengthLte:
	LengthLte []float64 `json:"length__lte,omitempty" mapstructure:"length__lte,omitempty"`
	// LengthN:
	LengthN []float64 `json:"length__n,omitempty" mapstructure:"length__n,omitempty"`
	// LengthUnit: * `km` - Kilometers
	// * `m` - Meters
	// * `cm` - Centimeters
	// * `mi` - Miles
	// * `ft` - Feet
	// * `in` - Inches
	LengthUnit string `json:"length_unit,omitempty" mapstructure:"length_unit,omitempty"`
	// LengthUnitN: * `km` - Kilometers
	// * `m` - Meters
	// * `cm` - Centimeters
	// * `mi` - Miles
	// * `ft` - Feet
	// * `in` - Inches
	LengthUnitN string `json:"length_unit__n,omitempty" mapstructure:"length_unit__n,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// Location:
	Location []string `json:"location,omitempty" mapstructure:"location,omitempty"`
	// LocationId:
	LocationId []int32 `json:"location_id,omitempty" mapstructure:"location_id,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Rack:
	Rack []string `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// RackId:
	RackId []int32 `json:"rack_id,omitempty" mapstructure:"rack_id,omitempty"`
	// Site:
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteId:
	SiteId []int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
	// Status: * `connected` - Connected
	// * `planned` - Planned
	// * `decommissioning` - Decommissioning
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `connected` - Connected
	// * `planned` - Planned
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
	// TerminationAId:
	TerminationAId []int32 `json:"termination_a_id,omitempty" mapstructure:"termination_a_id,omitempty"`
	// TerminationAType:
	TerminationAType string `json:"termination_a_type,omitempty" mapstructure:"termination_a_type,omitempty"`
	// TerminationATypeN:
	TerminationATypeN string `json:"termination_a_type__n,omitempty" mapstructure:"termination_a_type__n,omitempty"`
	// TerminationBId:
	TerminationBId []int32 `json:"termination_b_id,omitempty" mapstructure:"termination_b_id,omitempty"`
	// TerminationBType:
	TerminationBType string `json:"termination_b_type,omitempty" mapstructure:"termination_b_type,omitempty"`
	// TerminationBTypeN:
	TerminationBTypeN string `json:"termination_b_type__n,omitempty" mapstructure:"termination_b_type__n,omitempty"`
	// Type: * `cat3` - CAT3
	// * `cat5` - CAT5
	// * `cat5e` - CAT5e
	// * `cat6` - CAT6
	// * `cat6a` - CAT6a
	// * `cat7` - CAT7
	// * `cat7a` - CAT7a
	// * `cat8` - CAT8
	// * `dac-active` - Direct Attach Copper (Active)
	// * `dac-passive` - Direct Attach Copper (Passive)
	// * `mrj21-trunk` - MRJ21 Trunk
	// * `coaxial` - Coaxial
	// * `mmf` - Multimode Fiber
	// * `mmf-om1` - Multimode Fiber (OM1)
	// * `mmf-om2` - Multimode Fiber (OM2)
	// * `mmf-om3` - Multimode Fiber (OM3)
	// * `mmf-om4` - Multimode Fiber (OM4)
	// * `mmf-om5` - Multimode Fiber (OM5)
	// * `smf` - Singlemode Fiber
	// * `smf-os1` - Singlemode Fiber (OS1)
	// * `smf-os2` - Singlemode Fiber (OS2)
	// * `aoc` - Active Optical Cabling (AOC)
	// * `power` - Power
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `cat3` - CAT3
	// * `cat5` - CAT5
	// * `cat5e` - CAT5e
	// * `cat6` - CAT6
	// * `cat6a` - CAT6a
	// * `cat7` - CAT7
	// * `cat7a` - CAT7a
	// * `cat8` - CAT8
	// * `dac-active` - Direct Attach Copper (Active)
	// * `dac-passive` - Direct Attach Copper (Passive)
	// * `mrj21-trunk` - MRJ21 Trunk
	// * `coaxial` - Coaxial
	// * `mmf` - Multimode Fiber
	// * `mmf-om1` - Multimode Fiber (OM1)
	// * `mmf-om2` - Multimode Fiber (OM2)
	// * `mmf-om3` - Multimode Fiber (OM3)
	// * `mmf-om4` - Multimode Fiber (OM4)
	// * `mmf-om5` - Multimode Fiber (OM5)
	// * `smf` - Singlemode Fiber
	// * `smf-os1` - Singlemode Fiber (OS1)
	// * `smf-os2` - Singlemode Fiber (OS2)
	// * `aoc` - Active Optical Cabling (AOC)
	// * `power` - Power
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimCablesListQueryParameters) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color,
		),
		"colorN": validation.Validate(
			m.ColorN,
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
		"deviceId": validation.Validate(
			m.DeviceId,
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
		"length": validation.Validate(
			m.Length,
		),
		"lengthGt": validation.Validate(
			m.LengthGt,
		),
		"lengthGte": validation.Validate(
			m.LengthGte,
		),
		"lengthLt": validation.Validate(
			m.LengthLt,
		),
		"lengthLte": validation.Validate(
			m.LengthLte,
		),
		"lengthN": validation.Validate(
			m.LengthN,
		),
		"location": validation.Validate(
			m.Location,
		),
		"locationId": validation.Validate(
			m.LocationId,
		),
		"rack": validation.Validate(
			m.Rack,
		),
		"rackId": validation.Validate(
			m.RackId,
		),
		"site": validation.Validate(
			m.Site,
		),
		"siteId": validation.Validate(
			m.SiteId,
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
		"terminationAId": validation.Validate(
			m.TerminationAId,
		),
		"terminationBId": validation.Validate(
			m.TerminationBId,
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
	}.Filter()
}

// GetColor returns the Color property
func (m DcimCablesListQueryParameters) GetColor() []string {
	return m.Color
}

// SetColor sets the Color property
func (m *DcimCablesListQueryParameters) SetColor(val []string) {
	m.Color = val
}

// GetColorN returns the ColorN property
func (m DcimCablesListQueryParameters) GetColorN() []string {
	return m.ColorN
}

// SetColorN sets the ColorN property
func (m *DcimCablesListQueryParameters) SetColorN(val []string) {
	m.ColorN = val
}

// GetCreated returns the Created property
func (m DcimCablesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimCablesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimCablesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimCablesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimCablesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimCablesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimCablesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimCablesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimCablesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimCablesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimCablesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimCablesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimCablesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimCablesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevice returns the Device property
func (m DcimCablesListQueryParameters) GetDevice() []string {
	return m.Device
}

// SetDevice sets the Device property
func (m *DcimCablesListQueryParameters) SetDevice(val []string) {
	m.Device = val
}

// GetDeviceId returns the DeviceId property
func (m DcimCablesListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *DcimCablesListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetId returns the Id property
func (m DcimCablesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimCablesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimCablesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimCablesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimCablesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimCablesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimCablesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimCablesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimCablesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimCablesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimCablesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimCablesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLabel returns the Label property
func (m DcimCablesListQueryParameters) GetLabel() []string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DcimCablesListQueryParameters) SetLabel(val []string) {
	m.Label = val
}

// GetLabelEmpty returns the LabelEmpty property
func (m DcimCablesListQueryParameters) GetLabelEmpty() []string {
	return m.LabelEmpty
}

// SetLabelEmpty sets the LabelEmpty property
func (m *DcimCablesListQueryParameters) SetLabelEmpty(val []string) {
	m.LabelEmpty = val
}

// GetLabelIc returns the LabelIc property
func (m DcimCablesListQueryParameters) GetLabelIc() []string {
	return m.LabelIc
}

// SetLabelIc sets the LabelIc property
func (m *DcimCablesListQueryParameters) SetLabelIc(val []string) {
	m.LabelIc = val
}

// GetLabelIe returns the LabelIe property
func (m DcimCablesListQueryParameters) GetLabelIe() []string {
	return m.LabelIe
}

// SetLabelIe sets the LabelIe property
func (m *DcimCablesListQueryParameters) SetLabelIe(val []string) {
	m.LabelIe = val
}

// GetLabelIew returns the LabelIew property
func (m DcimCablesListQueryParameters) GetLabelIew() []string {
	return m.LabelIew
}

// SetLabelIew sets the LabelIew property
func (m *DcimCablesListQueryParameters) SetLabelIew(val []string) {
	m.LabelIew = val
}

// GetLabelIsw returns the LabelIsw property
func (m DcimCablesListQueryParameters) GetLabelIsw() []string {
	return m.LabelIsw
}

// SetLabelIsw sets the LabelIsw property
func (m *DcimCablesListQueryParameters) SetLabelIsw(val []string) {
	m.LabelIsw = val
}

// GetLabelN returns the LabelN property
func (m DcimCablesListQueryParameters) GetLabelN() []string {
	return m.LabelN
}

// SetLabelN sets the LabelN property
func (m *DcimCablesListQueryParameters) SetLabelN(val []string) {
	m.LabelN = val
}

// GetLabelNic returns the LabelNic property
func (m DcimCablesListQueryParameters) GetLabelNic() []string {
	return m.LabelNic
}

// SetLabelNic sets the LabelNic property
func (m *DcimCablesListQueryParameters) SetLabelNic(val []string) {
	m.LabelNic = val
}

// GetLabelNie returns the LabelNie property
func (m DcimCablesListQueryParameters) GetLabelNie() []string {
	return m.LabelNie
}

// SetLabelNie sets the LabelNie property
func (m *DcimCablesListQueryParameters) SetLabelNie(val []string) {
	m.LabelNie = val
}

// GetLabelNiew returns the LabelNiew property
func (m DcimCablesListQueryParameters) GetLabelNiew() []string {
	return m.LabelNiew
}

// SetLabelNiew sets the LabelNiew property
func (m *DcimCablesListQueryParameters) SetLabelNiew(val []string) {
	m.LabelNiew = val
}

// GetLabelNisw returns the LabelNisw property
func (m DcimCablesListQueryParameters) GetLabelNisw() []string {
	return m.LabelNisw
}

// SetLabelNisw sets the LabelNisw property
func (m *DcimCablesListQueryParameters) SetLabelNisw(val []string) {
	m.LabelNisw = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimCablesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimCablesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimCablesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimCablesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimCablesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimCablesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimCablesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimCablesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimCablesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimCablesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimCablesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimCablesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLength returns the Length property
func (m DcimCablesListQueryParameters) GetLength() []float64 {
	return m.Length
}

// SetLength sets the Length property
func (m *DcimCablesListQueryParameters) SetLength(val []float64) {
	m.Length = val
}

// GetLengthGt returns the LengthGt property
func (m DcimCablesListQueryParameters) GetLengthGt() []float64 {
	return m.LengthGt
}

// SetLengthGt sets the LengthGt property
func (m *DcimCablesListQueryParameters) SetLengthGt(val []float64) {
	m.LengthGt = val
}

// GetLengthGte returns the LengthGte property
func (m DcimCablesListQueryParameters) GetLengthGte() []float64 {
	return m.LengthGte
}

// SetLengthGte sets the LengthGte property
func (m *DcimCablesListQueryParameters) SetLengthGte(val []float64) {
	m.LengthGte = val
}

// GetLengthLt returns the LengthLt property
func (m DcimCablesListQueryParameters) GetLengthLt() []float64 {
	return m.LengthLt
}

// SetLengthLt sets the LengthLt property
func (m *DcimCablesListQueryParameters) SetLengthLt(val []float64) {
	m.LengthLt = val
}

// GetLengthLte returns the LengthLte property
func (m DcimCablesListQueryParameters) GetLengthLte() []float64 {
	return m.LengthLte
}

// SetLengthLte sets the LengthLte property
func (m *DcimCablesListQueryParameters) SetLengthLte(val []float64) {
	m.LengthLte = val
}

// GetLengthN returns the LengthN property
func (m DcimCablesListQueryParameters) GetLengthN() []float64 {
	return m.LengthN
}

// SetLengthN sets the LengthN property
func (m *DcimCablesListQueryParameters) SetLengthN(val []float64) {
	m.LengthN = val
}

// GetLengthUnit returns the LengthUnit property
func (m DcimCablesListQueryParameters) GetLengthUnit() string {
	return m.LengthUnit
}

// SetLengthUnit sets the LengthUnit property
func (m *DcimCablesListQueryParameters) SetLengthUnit(val string) {
	m.LengthUnit = val
}

// GetLengthUnitN returns the LengthUnitN property
func (m DcimCablesListQueryParameters) GetLengthUnitN() string {
	return m.LengthUnitN
}

// SetLengthUnitN sets the LengthUnitN property
func (m *DcimCablesListQueryParameters) SetLengthUnitN(val string) {
	m.LengthUnitN = val
}

// GetLimit returns the Limit property
func (m DcimCablesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimCablesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m DcimCablesListQueryParameters) GetLocation() []string {
	return m.Location
}

// SetLocation sets the Location property
func (m *DcimCablesListQueryParameters) SetLocation(val []string) {
	m.Location = val
}

// GetLocationId returns the LocationId property
func (m DcimCablesListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimCablesListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetOffset returns the Offset property
func (m DcimCablesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimCablesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimCablesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimCablesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimCablesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimCablesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRack returns the Rack property
func (m DcimCablesListQueryParameters) GetRack() []string {
	return m.Rack
}

// SetRack sets the Rack property
func (m *DcimCablesListQueryParameters) SetRack(val []string) {
	m.Rack = val
}

// GetRackId returns the RackId property
func (m DcimCablesListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimCablesListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetSite returns the Site property
func (m DcimCablesListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimCablesListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteId returns the SiteId property
func (m DcimCablesListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimCablesListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetStatus returns the Status property
func (m DcimCablesListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *DcimCablesListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m DcimCablesListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *DcimCablesListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m DcimCablesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimCablesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimCablesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimCablesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m DcimCablesListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DcimCablesListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m DcimCablesListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *DcimCablesListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m DcimCablesListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *DcimCablesListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m DcimCablesListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *DcimCablesListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m DcimCablesListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *DcimCablesListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m DcimCablesListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *DcimCablesListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m DcimCablesListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *DcimCablesListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m DcimCablesListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *DcimCablesListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetTerminationAId returns the TerminationAId property
func (m DcimCablesListQueryParameters) GetTerminationAId() []int32 {
	return m.TerminationAId
}

// SetTerminationAId sets the TerminationAId property
func (m *DcimCablesListQueryParameters) SetTerminationAId(val []int32) {
	m.TerminationAId = val
}

// GetTerminationAType returns the TerminationAType property
func (m DcimCablesListQueryParameters) GetTerminationAType() string {
	return m.TerminationAType
}

// SetTerminationAType sets the TerminationAType property
func (m *DcimCablesListQueryParameters) SetTerminationAType(val string) {
	m.TerminationAType = val
}

// GetTerminationATypeN returns the TerminationATypeN property
func (m DcimCablesListQueryParameters) GetTerminationATypeN() string {
	return m.TerminationATypeN
}

// SetTerminationATypeN sets the TerminationATypeN property
func (m *DcimCablesListQueryParameters) SetTerminationATypeN(val string) {
	m.TerminationATypeN = val
}

// GetTerminationBId returns the TerminationBId property
func (m DcimCablesListQueryParameters) GetTerminationBId() []int32 {
	return m.TerminationBId
}

// SetTerminationBId sets the TerminationBId property
func (m *DcimCablesListQueryParameters) SetTerminationBId(val []int32) {
	m.TerminationBId = val
}

// GetTerminationBType returns the TerminationBType property
func (m DcimCablesListQueryParameters) GetTerminationBType() string {
	return m.TerminationBType
}

// SetTerminationBType sets the TerminationBType property
func (m *DcimCablesListQueryParameters) SetTerminationBType(val string) {
	m.TerminationBType = val
}

// GetTerminationBTypeN returns the TerminationBTypeN property
func (m DcimCablesListQueryParameters) GetTerminationBTypeN() string {
	return m.TerminationBTypeN
}

// SetTerminationBTypeN sets the TerminationBTypeN property
func (m *DcimCablesListQueryParameters) SetTerminationBTypeN(val string) {
	m.TerminationBTypeN = val
}

// GetType returns the Type property
func (m DcimCablesListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimCablesListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimCablesListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimCablesListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimCablesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimCablesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
