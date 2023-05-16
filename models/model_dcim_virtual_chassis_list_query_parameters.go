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

// DcimVirtualChassisListQueryParameters is an object.
type DcimVirtualChassisListQueryParameters struct {
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
	// Domain:
	Domain []string `json:"domain,omitempty" mapstructure:"domain,omitempty"`
	// DomainEmpty:
	DomainEmpty []string `json:"domain__empty,omitempty" mapstructure:"domain__empty,omitempty"`
	// DomainIc:
	DomainIc []string `json:"domain__ic,omitempty" mapstructure:"domain__ic,omitempty"`
	// DomainIe:
	DomainIe []string `json:"domain__ie,omitempty" mapstructure:"domain__ie,omitempty"`
	// DomainIew:
	DomainIew []string `json:"domain__iew,omitempty" mapstructure:"domain__iew,omitempty"`
	// DomainIsw:
	DomainIsw []string `json:"domain__isw,omitempty" mapstructure:"domain__isw,omitempty"`
	// DomainN:
	DomainN []string `json:"domain__n,omitempty" mapstructure:"domain__n,omitempty"`
	// DomainNic:
	DomainNic []string `json:"domain__nic,omitempty" mapstructure:"domain__nic,omitempty"`
	// DomainNie:
	DomainNie []string `json:"domain__nie,omitempty" mapstructure:"domain__nie,omitempty"`
	// DomainNiew:
	DomainNiew []string `json:"domain__niew,omitempty" mapstructure:"domain__niew,omitempty"`
	// DomainNisw:
	DomainNisw []string `json:"domain__nisw,omitempty" mapstructure:"domain__nisw,omitempty"`
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
	// Master: Master (name)
	Master []*string `json:"master,omitempty" mapstructure:"master,omitempty"`
	// MasterN: Master (name)
	MasterN []*string `json:"master__n,omitempty" mapstructure:"master__n,omitempty"`
	// MasterId: Master (ID)
	MasterId []*int32 `json:"master_id,omitempty" mapstructure:"master_id,omitempty"`
	// MasterIdN: Master (ID)
	MasterIdN []*int32 `json:"master_id__n,omitempty" mapstructure:"master_id__n,omitempty"`
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
	// Tenant: Tenant (slug)
	Tenant []string `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TenantN: Tenant (slug)
	TenantN []string `json:"tenant__n,omitempty" mapstructure:"tenant__n,omitempty"`
	// TenantId: Tenant (ID)
	TenantId []int32 `json:"tenant_id,omitempty" mapstructure:"tenant_id,omitempty"`
	// TenantIdN: Tenant (ID)
	TenantIdN []int32 `json:"tenant_id__n,omitempty" mapstructure:"tenant_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimVirtualChassisListQueryParameters) Validate() error {
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
		"domain": validation.Validate(
			m.Domain,
		),
		"domainEmpty": validation.Validate(
			m.DomainEmpty,
		),
		"domainIc": validation.Validate(
			m.DomainIc,
		),
		"domainIe": validation.Validate(
			m.DomainIe,
		),
		"domainIew": validation.Validate(
			m.DomainIew,
		),
		"domainIsw": validation.Validate(
			m.DomainIsw,
		),
		"domainN": validation.Validate(
			m.DomainN,
		),
		"domainNic": validation.Validate(
			m.DomainNic,
		),
		"domainNie": validation.Validate(
			m.DomainNie,
		),
		"domainNiew": validation.Validate(
			m.DomainNiew,
		),
		"domainNisw": validation.Validate(
			m.DomainNisw,
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
		"master": validation.Validate(
			m.Master,
		),
		"masterN": validation.Validate(
			m.MasterN,
		),
		"masterId": validation.Validate(
			m.MasterId,
		),
		"masterIdN": validation.Validate(
			m.MasterIdN,
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
		"tenantId": validation.Validate(
			m.TenantId,
		),
		"tenantIdN": validation.Validate(
			m.TenantIdN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m DcimVirtualChassisListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimVirtualChassisListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimVirtualChassisListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimVirtualChassisListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimVirtualChassisListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimVirtualChassisListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimVirtualChassisListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimVirtualChassisListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimVirtualChassisListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimVirtualChassisListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimVirtualChassisListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimVirtualChassisListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimVirtualChassisListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimVirtualChassisListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDomain returns the Domain property
func (m DcimVirtualChassisListQueryParameters) GetDomain() []string {
	return m.Domain
}

// SetDomain sets the Domain property
func (m *DcimVirtualChassisListQueryParameters) SetDomain(val []string) {
	m.Domain = val
}

// GetDomainEmpty returns the DomainEmpty property
func (m DcimVirtualChassisListQueryParameters) GetDomainEmpty() []string {
	return m.DomainEmpty
}

// SetDomainEmpty sets the DomainEmpty property
func (m *DcimVirtualChassisListQueryParameters) SetDomainEmpty(val []string) {
	m.DomainEmpty = val
}

// GetDomainIc returns the DomainIc property
func (m DcimVirtualChassisListQueryParameters) GetDomainIc() []string {
	return m.DomainIc
}

// SetDomainIc sets the DomainIc property
func (m *DcimVirtualChassisListQueryParameters) SetDomainIc(val []string) {
	m.DomainIc = val
}

// GetDomainIe returns the DomainIe property
func (m DcimVirtualChassisListQueryParameters) GetDomainIe() []string {
	return m.DomainIe
}

// SetDomainIe sets the DomainIe property
func (m *DcimVirtualChassisListQueryParameters) SetDomainIe(val []string) {
	m.DomainIe = val
}

// GetDomainIew returns the DomainIew property
func (m DcimVirtualChassisListQueryParameters) GetDomainIew() []string {
	return m.DomainIew
}

// SetDomainIew sets the DomainIew property
func (m *DcimVirtualChassisListQueryParameters) SetDomainIew(val []string) {
	m.DomainIew = val
}

// GetDomainIsw returns the DomainIsw property
func (m DcimVirtualChassisListQueryParameters) GetDomainIsw() []string {
	return m.DomainIsw
}

// SetDomainIsw sets the DomainIsw property
func (m *DcimVirtualChassisListQueryParameters) SetDomainIsw(val []string) {
	m.DomainIsw = val
}

// GetDomainN returns the DomainN property
func (m DcimVirtualChassisListQueryParameters) GetDomainN() []string {
	return m.DomainN
}

// SetDomainN sets the DomainN property
func (m *DcimVirtualChassisListQueryParameters) SetDomainN(val []string) {
	m.DomainN = val
}

// GetDomainNic returns the DomainNic property
func (m DcimVirtualChassisListQueryParameters) GetDomainNic() []string {
	return m.DomainNic
}

// SetDomainNic sets the DomainNic property
func (m *DcimVirtualChassisListQueryParameters) SetDomainNic(val []string) {
	m.DomainNic = val
}

// GetDomainNie returns the DomainNie property
func (m DcimVirtualChassisListQueryParameters) GetDomainNie() []string {
	return m.DomainNie
}

// SetDomainNie sets the DomainNie property
func (m *DcimVirtualChassisListQueryParameters) SetDomainNie(val []string) {
	m.DomainNie = val
}

// GetDomainNiew returns the DomainNiew property
func (m DcimVirtualChassisListQueryParameters) GetDomainNiew() []string {
	return m.DomainNiew
}

// SetDomainNiew sets the DomainNiew property
func (m *DcimVirtualChassisListQueryParameters) SetDomainNiew(val []string) {
	m.DomainNiew = val
}

// GetDomainNisw returns the DomainNisw property
func (m DcimVirtualChassisListQueryParameters) GetDomainNisw() []string {
	return m.DomainNisw
}

// SetDomainNisw sets the DomainNisw property
func (m *DcimVirtualChassisListQueryParameters) SetDomainNisw(val []string) {
	m.DomainNisw = val
}

// GetId returns the Id property
func (m DcimVirtualChassisListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimVirtualChassisListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimVirtualChassisListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimVirtualChassisListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimVirtualChassisListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimVirtualChassisListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimVirtualChassisListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimVirtualChassisListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimVirtualChassisListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimVirtualChassisListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimVirtualChassisListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimVirtualChassisListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimVirtualChassisListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimVirtualChassisListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimVirtualChassisListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimVirtualChassisListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimVirtualChassisListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimVirtualChassisListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimVirtualChassisListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimVirtualChassisListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimVirtualChassisListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimVirtualChassisListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimVirtualChassisListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimVirtualChassisListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimVirtualChassisListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimVirtualChassisListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetMaster returns the Master property
func (m DcimVirtualChassisListQueryParameters) GetMaster() []*string {
	return m.Master
}

// SetMaster sets the Master property
func (m *DcimVirtualChassisListQueryParameters) SetMaster(val []*string) {
	m.Master = val
}

// GetMasterN returns the MasterN property
func (m DcimVirtualChassisListQueryParameters) GetMasterN() []*string {
	return m.MasterN
}

// SetMasterN sets the MasterN property
func (m *DcimVirtualChassisListQueryParameters) SetMasterN(val []*string) {
	m.MasterN = val
}

// GetMasterId returns the MasterId property
func (m DcimVirtualChassisListQueryParameters) GetMasterId() []*int32 {
	return m.MasterId
}

// SetMasterId sets the MasterId property
func (m *DcimVirtualChassisListQueryParameters) SetMasterId(val []*int32) {
	m.MasterId = val
}

// GetMasterIdN returns the MasterIdN property
func (m DcimVirtualChassisListQueryParameters) GetMasterIdN() []*int32 {
	return m.MasterIdN
}

// SetMasterIdN sets the MasterIdN property
func (m *DcimVirtualChassisListQueryParameters) SetMasterIdN(val []*int32) {
	m.MasterIdN = val
}

// GetName returns the Name property
func (m DcimVirtualChassisListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimVirtualChassisListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimVirtualChassisListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimVirtualChassisListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimVirtualChassisListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimVirtualChassisListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimVirtualChassisListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimVirtualChassisListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimVirtualChassisListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimVirtualChassisListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimVirtualChassisListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimVirtualChassisListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimVirtualChassisListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimVirtualChassisListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimVirtualChassisListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimVirtualChassisListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimVirtualChassisListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimVirtualChassisListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimVirtualChassisListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimVirtualChassisListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimVirtualChassisListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimVirtualChassisListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimVirtualChassisListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimVirtualChassisListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimVirtualChassisListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimVirtualChassisListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimVirtualChassisListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimVirtualChassisListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m DcimVirtualChassisListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimVirtualChassisListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimVirtualChassisListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimVirtualChassisListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimVirtualChassisListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimVirtualChassisListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimVirtualChassisListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimVirtualChassisListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSite returns the Site property
func (m DcimVirtualChassisListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimVirtualChassisListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimVirtualChassisListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimVirtualChassisListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimVirtualChassisListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimVirtualChassisListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimVirtualChassisListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimVirtualChassisListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimVirtualChassisListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimVirtualChassisListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimVirtualChassisListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimVirtualChassisListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimVirtualChassisListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimVirtualChassisListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimVirtualChassisListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimVirtualChassisListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetTag returns the Tag property
func (m DcimVirtualChassisListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimVirtualChassisListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimVirtualChassisListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimVirtualChassisListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m DcimVirtualChassisListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DcimVirtualChassisListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m DcimVirtualChassisListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *DcimVirtualChassisListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantId returns the TenantId property
func (m DcimVirtualChassisListQueryParameters) GetTenantId() []int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *DcimVirtualChassisListQueryParameters) SetTenantId(val []int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m DcimVirtualChassisListQueryParameters) GetTenantIdN() []int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *DcimVirtualChassisListQueryParameters) SetTenantIdN(val []int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimVirtualChassisListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimVirtualChassisListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
