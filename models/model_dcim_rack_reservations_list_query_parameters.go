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

// DcimRackReservationsListQueryParameters is an object.
type DcimRackReservationsListQueryParameters struct {
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
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
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
	// User: User (name)
	User []string `json:"user,omitempty" mapstructure:"user,omitempty"`
	// UserN: User (name)
	UserN []string `json:"user__n,omitempty" mapstructure:"user__n,omitempty"`
	// UserId: User (ID)
	UserId []int32 `json:"user_id,omitempty" mapstructure:"user_id,omitempty"`
	// UserIdN: User (ID)
	UserIdN []int32 `json:"user_id__n,omitempty" mapstructure:"user_id__n,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimRackReservationsListQueryParameters) Validate() error {
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
		"user": validation.Validate(
			m.User,
		),
		"userN": validation.Validate(
			m.UserN,
		),
		"userId": validation.Validate(
			m.UserId,
		),
		"userIdN": validation.Validate(
			m.UserIdN,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m DcimRackReservationsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimRackReservationsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimRackReservationsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimRackReservationsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimRackReservationsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimRackReservationsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimRackReservationsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimRackReservationsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimRackReservationsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimRackReservationsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimRackReservationsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimRackReservationsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimRackReservationsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimRackReservationsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m DcimRackReservationsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DcimRackReservationsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m DcimRackReservationsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m DcimRackReservationsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m DcimRackReservationsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m DcimRackReservationsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m DcimRackReservationsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m DcimRackReservationsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m DcimRackReservationsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m DcimRackReservationsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m DcimRackReservationsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m DcimRackReservationsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *DcimRackReservationsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m DcimRackReservationsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimRackReservationsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimRackReservationsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimRackReservationsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimRackReservationsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimRackReservationsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimRackReservationsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimRackReservationsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimRackReservationsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimRackReservationsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimRackReservationsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimRackReservationsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimRackReservationsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimRackReservationsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimRackReservationsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimRackReservationsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimRackReservationsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimRackReservationsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimRackReservationsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimRackReservationsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimRackReservationsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimRackReservationsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimRackReservationsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimRackReservationsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimRackReservationsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimRackReservationsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m DcimRackReservationsListQueryParameters) GetLocation() []int32 {
	return m.Location
}

// SetLocation sets the Location property
func (m *DcimRackReservationsListQueryParameters) SetLocation(val []int32) {
	m.Location = val
}

// GetLocationN returns the LocationN property
func (m DcimRackReservationsListQueryParameters) GetLocationN() []int32 {
	return m.LocationN
}

// SetLocationN sets the LocationN property
func (m *DcimRackReservationsListQueryParameters) SetLocationN(val []int32) {
	m.LocationN = val
}

// GetLocationId returns the LocationId property
func (m DcimRackReservationsListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimRackReservationsListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m DcimRackReservationsListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *DcimRackReservationsListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetOffset returns the Offset property
func (m DcimRackReservationsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimRackReservationsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimRackReservationsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimRackReservationsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimRackReservationsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimRackReservationsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRackId returns the RackId property
func (m DcimRackReservationsListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimRackReservationsListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetRackIdN returns the RackIdN property
func (m DcimRackReservationsListQueryParameters) GetRackIdN() []int32 {
	return m.RackIdN
}

// SetRackIdN sets the RackIdN property
func (m *DcimRackReservationsListQueryParameters) SetRackIdN(val []int32) {
	m.RackIdN = val
}

// GetRegion returns the Region property
func (m DcimRackReservationsListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimRackReservationsListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimRackReservationsListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimRackReservationsListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimRackReservationsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimRackReservationsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimRackReservationsListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimRackReservationsListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSite returns the Site property
func (m DcimRackReservationsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimRackReservationsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimRackReservationsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimRackReservationsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimRackReservationsListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimRackReservationsListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimRackReservationsListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimRackReservationsListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimRackReservationsListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimRackReservationsListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimRackReservationsListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimRackReservationsListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimRackReservationsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimRackReservationsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimRackReservationsListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimRackReservationsListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetTag returns the Tag property
func (m DcimRackReservationsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimRackReservationsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimRackReservationsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimRackReservationsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m DcimRackReservationsListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DcimRackReservationsListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m DcimRackReservationsListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *DcimRackReservationsListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m DcimRackReservationsListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *DcimRackReservationsListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m DcimRackReservationsListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *DcimRackReservationsListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m DcimRackReservationsListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *DcimRackReservationsListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m DcimRackReservationsListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *DcimRackReservationsListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m DcimRackReservationsListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *DcimRackReservationsListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m DcimRackReservationsListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *DcimRackReservationsListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimRackReservationsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimRackReservationsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetUser returns the User property
func (m DcimRackReservationsListQueryParameters) GetUser() []string {
	return m.User
}

// SetUser sets the User property
func (m *DcimRackReservationsListQueryParameters) SetUser(val []string) {
	m.User = val
}

// GetUserN returns the UserN property
func (m DcimRackReservationsListQueryParameters) GetUserN() []string {
	return m.UserN
}

// SetUserN sets the UserN property
func (m *DcimRackReservationsListQueryParameters) SetUserN(val []string) {
	m.UserN = val
}

// GetUserId returns the UserId property
func (m DcimRackReservationsListQueryParameters) GetUserId() []int32 {
	return m.UserId
}

// SetUserId sets the UserId property
func (m *DcimRackReservationsListQueryParameters) SetUserId(val []int32) {
	m.UserId = val
}

// GetUserIdN returns the UserIdN property
func (m DcimRackReservationsListQueryParameters) GetUserIdN() []int32 {
	return m.UserIdN
}

// SetUserIdN sets the UserIdN property
func (m *DcimRackReservationsListQueryParameters) SetUserIdN(val []int32) {
	m.UserIdN = val
}
