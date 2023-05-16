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

// CircuitsCircuitsListQueryParameters is an object.
type CircuitsCircuitsListQueryParameters struct {
	// Cid:
	Cid []string `json:"cid,omitempty" mapstructure:"cid,omitempty"`
	// CidEmpty:
	CidEmpty []string `json:"cid__empty,omitempty" mapstructure:"cid__empty,omitempty"`
	// CidIc:
	CidIc []string `json:"cid__ic,omitempty" mapstructure:"cid__ic,omitempty"`
	// CidIe:
	CidIe []string `json:"cid__ie,omitempty" mapstructure:"cid__ie,omitempty"`
	// CidIew:
	CidIew []string `json:"cid__iew,omitempty" mapstructure:"cid__iew,omitempty"`
	// CidIsw:
	CidIsw []string `json:"cid__isw,omitempty" mapstructure:"cid__isw,omitempty"`
	// CidN:
	CidN []string `json:"cid__n,omitempty" mapstructure:"cid__n,omitempty"`
	// CidNic:
	CidNic []string `json:"cid__nic,omitempty" mapstructure:"cid__nic,omitempty"`
	// CidNie:
	CidNie []string `json:"cid__nie,omitempty" mapstructure:"cid__nie,omitempty"`
	// CidNiew:
	CidNiew []string `json:"cid__niew,omitempty" mapstructure:"cid__niew,omitempty"`
	// CidNisw:
	CidNisw []string `json:"cid__nisw,omitempty" mapstructure:"cid__nisw,omitempty"`
	// CommitRate:
	CommitRate []int32 `json:"commit_rate,omitempty" mapstructure:"commit_rate,omitempty"`
	// CommitRateGt:
	CommitRateGt []int32 `json:"commit_rate__gt,omitempty" mapstructure:"commit_rate__gt,omitempty"`
	// CommitRateGte:
	CommitRateGte []int32 `json:"commit_rate__gte,omitempty" mapstructure:"commit_rate__gte,omitempty"`
	// CommitRateLt:
	CommitRateLt []int32 `json:"commit_rate__lt,omitempty" mapstructure:"commit_rate__lt,omitempty"`
	// CommitRateLte:
	CommitRateLte []int32 `json:"commit_rate__lte,omitempty" mapstructure:"commit_rate__lte,omitempty"`
	// CommitRateN:
	CommitRateN []int32 `json:"commit_rate__n,omitempty" mapstructure:"commit_rate__n,omitempty"`
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
	// InstallDate:
	InstallDate []string `json:"install_date,omitempty" mapstructure:"install_date,omitempty"`
	// InstallDateGt:
	InstallDateGt []string `json:"install_date__gt,omitempty" mapstructure:"install_date__gt,omitempty"`
	// InstallDateGte:
	InstallDateGte []string `json:"install_date__gte,omitempty" mapstructure:"install_date__gte,omitempty"`
	// InstallDateLt:
	InstallDateLt []string `json:"install_date__lt,omitempty" mapstructure:"install_date__lt,omitempty"`
	// InstallDateLte:
	InstallDateLte []string `json:"install_date__lte,omitempty" mapstructure:"install_date__lte,omitempty"`
	// InstallDateN:
	InstallDateN []string `json:"install_date__n,omitempty" mapstructure:"install_date__n,omitempty"`
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
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Provider: Provider (slug)
	Provider []string `json:"provider,omitempty" mapstructure:"provider,omitempty"`
	// ProviderN: Provider (slug)
	ProviderN []string `json:"provider__n,omitempty" mapstructure:"provider__n,omitempty"`
	// ProviderAccountId: ProviderAccount (ID)
	ProviderAccountId []int32 `json:"provider_account_id,omitempty" mapstructure:"provider_account_id,omitempty"`
	// ProviderAccountIdN: ProviderAccount (ID)
	ProviderAccountIdN []int32 `json:"provider_account_id__n,omitempty" mapstructure:"provider_account_id__n,omitempty"`
	// ProviderId: Provider (ID)
	ProviderId []int32 `json:"provider_id,omitempty" mapstructure:"provider_id,omitempty"`
	// ProviderIdN: Provider (ID)
	ProviderIdN []int32 `json:"provider_id__n,omitempty" mapstructure:"provider_id__n,omitempty"`
	// ProviderNetworkId: ProviderNetwork (ID)
	ProviderNetworkId []int32 `json:"provider_network_id,omitempty" mapstructure:"provider_network_id,omitempty"`
	// ProviderNetworkIdN: ProviderNetwork (ID)
	ProviderNetworkIdN []int32 `json:"provider_network_id__n,omitempty" mapstructure:"provider_network_id__n,omitempty"`
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
	// Status: * `planned` - Planned
	// * `provisioning` - Provisioning
	// * `active` - Active
	// * `offline` - Offline
	// * `deprovisioning` - Deprovisioning
	// * `decommissioned` - Decommissioned
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `planned` - Planned
	// * `provisioning` - Provisioning
	// * `active` - Active
	// * `offline` - Offline
	// * `deprovisioning` - Deprovisioning
	// * `decommissioned` - Decommissioned
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
	// TerminationDate:
	TerminationDate []string `json:"termination_date,omitempty" mapstructure:"termination_date,omitempty"`
	// TerminationDateGt:
	TerminationDateGt []string `json:"termination_date__gt,omitempty" mapstructure:"termination_date__gt,omitempty"`
	// TerminationDateGte:
	TerminationDateGte []string `json:"termination_date__gte,omitempty" mapstructure:"termination_date__gte,omitempty"`
	// TerminationDateLt:
	TerminationDateLt []string `json:"termination_date__lt,omitempty" mapstructure:"termination_date__lt,omitempty"`
	// TerminationDateLte:
	TerminationDateLte []string `json:"termination_date__lte,omitempty" mapstructure:"termination_date__lte,omitempty"`
	// TerminationDateN:
	TerminationDateN []string `json:"termination_date__n,omitempty" mapstructure:"termination_date__n,omitempty"`
	// Type: Circuit type (slug)
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: Circuit type (slug)
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// TypeId: Circuit type (ID)
	TypeId []int32 `json:"type_id,omitempty" mapstructure:"type_id,omitempty"`
	// TypeIdN: Circuit type (ID)
	TypeIdN []int32 `json:"type_id__n,omitempty" mapstructure:"type_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitsCircuitsListQueryParameters) Validate() error {
	return validation.Errors{
		"cid": validation.Validate(
			m.Cid,
		),
		"cidEmpty": validation.Validate(
			m.CidEmpty,
		),
		"cidIc": validation.Validate(
			m.CidIc,
		),
		"cidIe": validation.Validate(
			m.CidIe,
		),
		"cidIew": validation.Validate(
			m.CidIew,
		),
		"cidIsw": validation.Validate(
			m.CidIsw,
		),
		"cidN": validation.Validate(
			m.CidN,
		),
		"cidNic": validation.Validate(
			m.CidNic,
		),
		"cidNie": validation.Validate(
			m.CidNie,
		),
		"cidNiew": validation.Validate(
			m.CidNiew,
		),
		"cidNisw": validation.Validate(
			m.CidNisw,
		),
		"commitRate": validation.Validate(
			m.CommitRate,
		),
		"commitRateGt": validation.Validate(
			m.CommitRateGt,
		),
		"commitRateGte": validation.Validate(
			m.CommitRateGte,
		),
		"commitRateLt": validation.Validate(
			m.CommitRateLt,
		),
		"commitRateLte": validation.Validate(
			m.CommitRateLte,
		),
		"commitRateN": validation.Validate(
			m.CommitRateN,
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
		"installDate": validation.Validate(
			m.InstallDate,
		),
		"installDateGt": validation.Validate(
			m.InstallDateGt,
		),
		"installDateGte": validation.Validate(
			m.InstallDateGte,
		),
		"installDateLt": validation.Validate(
			m.InstallDateLt,
		),
		"installDateLte": validation.Validate(
			m.InstallDateLte,
		),
		"installDateN": validation.Validate(
			m.InstallDateN,
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
		"provider": validation.Validate(
			m.Provider,
		),
		"providerN": validation.Validate(
			m.ProviderN,
		),
		"providerAccountId": validation.Validate(
			m.ProviderAccountId,
		),
		"providerAccountIdN": validation.Validate(
			m.ProviderAccountIdN,
		),
		"providerId": validation.Validate(
			m.ProviderId,
		),
		"providerIdN": validation.Validate(
			m.ProviderIdN,
		),
		"providerNetworkId": validation.Validate(
			m.ProviderNetworkId,
		),
		"providerNetworkIdN": validation.Validate(
			m.ProviderNetworkIdN,
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
		"terminationDate": validation.Validate(
			m.TerminationDate,
		),
		"terminationDateGt": validation.Validate(
			m.TerminationDateGt,
		),
		"terminationDateGte": validation.Validate(
			m.TerminationDateGte,
		),
		"terminationDateLt": validation.Validate(
			m.TerminationDateLt,
		),
		"terminationDateLte": validation.Validate(
			m.TerminationDateLte,
		),
		"terminationDateN": validation.Validate(
			m.TerminationDateN,
		),
		"type": validation.Validate(
			m.Type,
		),
		"typeN": validation.Validate(
			m.TypeN,
		),
		"typeId": validation.Validate(
			m.TypeId,
		),
		"typeIdN": validation.Validate(
			m.TypeIdN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetCid returns the Cid property
func (m CircuitsCircuitsListQueryParameters) GetCid() []string {
	return m.Cid
}

// SetCid sets the Cid property
func (m *CircuitsCircuitsListQueryParameters) SetCid(val []string) {
	m.Cid = val
}

// GetCidEmpty returns the CidEmpty property
func (m CircuitsCircuitsListQueryParameters) GetCidEmpty() []string {
	return m.CidEmpty
}

// SetCidEmpty sets the CidEmpty property
func (m *CircuitsCircuitsListQueryParameters) SetCidEmpty(val []string) {
	m.CidEmpty = val
}

// GetCidIc returns the CidIc property
func (m CircuitsCircuitsListQueryParameters) GetCidIc() []string {
	return m.CidIc
}

// SetCidIc sets the CidIc property
func (m *CircuitsCircuitsListQueryParameters) SetCidIc(val []string) {
	m.CidIc = val
}

// GetCidIe returns the CidIe property
func (m CircuitsCircuitsListQueryParameters) GetCidIe() []string {
	return m.CidIe
}

// SetCidIe sets the CidIe property
func (m *CircuitsCircuitsListQueryParameters) SetCidIe(val []string) {
	m.CidIe = val
}

// GetCidIew returns the CidIew property
func (m CircuitsCircuitsListQueryParameters) GetCidIew() []string {
	return m.CidIew
}

// SetCidIew sets the CidIew property
func (m *CircuitsCircuitsListQueryParameters) SetCidIew(val []string) {
	m.CidIew = val
}

// GetCidIsw returns the CidIsw property
func (m CircuitsCircuitsListQueryParameters) GetCidIsw() []string {
	return m.CidIsw
}

// SetCidIsw sets the CidIsw property
func (m *CircuitsCircuitsListQueryParameters) SetCidIsw(val []string) {
	m.CidIsw = val
}

// GetCidN returns the CidN property
func (m CircuitsCircuitsListQueryParameters) GetCidN() []string {
	return m.CidN
}

// SetCidN sets the CidN property
func (m *CircuitsCircuitsListQueryParameters) SetCidN(val []string) {
	m.CidN = val
}

// GetCidNic returns the CidNic property
func (m CircuitsCircuitsListQueryParameters) GetCidNic() []string {
	return m.CidNic
}

// SetCidNic sets the CidNic property
func (m *CircuitsCircuitsListQueryParameters) SetCidNic(val []string) {
	m.CidNic = val
}

// GetCidNie returns the CidNie property
func (m CircuitsCircuitsListQueryParameters) GetCidNie() []string {
	return m.CidNie
}

// SetCidNie sets the CidNie property
func (m *CircuitsCircuitsListQueryParameters) SetCidNie(val []string) {
	m.CidNie = val
}

// GetCidNiew returns the CidNiew property
func (m CircuitsCircuitsListQueryParameters) GetCidNiew() []string {
	return m.CidNiew
}

// SetCidNiew sets the CidNiew property
func (m *CircuitsCircuitsListQueryParameters) SetCidNiew(val []string) {
	m.CidNiew = val
}

// GetCidNisw returns the CidNisw property
func (m CircuitsCircuitsListQueryParameters) GetCidNisw() []string {
	return m.CidNisw
}

// SetCidNisw sets the CidNisw property
func (m *CircuitsCircuitsListQueryParameters) SetCidNisw(val []string) {
	m.CidNisw = val
}

// GetCommitRate returns the CommitRate property
func (m CircuitsCircuitsListQueryParameters) GetCommitRate() []int32 {
	return m.CommitRate
}

// SetCommitRate sets the CommitRate property
func (m *CircuitsCircuitsListQueryParameters) SetCommitRate(val []int32) {
	m.CommitRate = val
}

// GetCommitRateGt returns the CommitRateGt property
func (m CircuitsCircuitsListQueryParameters) GetCommitRateGt() []int32 {
	return m.CommitRateGt
}

// SetCommitRateGt sets the CommitRateGt property
func (m *CircuitsCircuitsListQueryParameters) SetCommitRateGt(val []int32) {
	m.CommitRateGt = val
}

// GetCommitRateGte returns the CommitRateGte property
func (m CircuitsCircuitsListQueryParameters) GetCommitRateGte() []int32 {
	return m.CommitRateGte
}

// SetCommitRateGte sets the CommitRateGte property
func (m *CircuitsCircuitsListQueryParameters) SetCommitRateGte(val []int32) {
	m.CommitRateGte = val
}

// GetCommitRateLt returns the CommitRateLt property
func (m CircuitsCircuitsListQueryParameters) GetCommitRateLt() []int32 {
	return m.CommitRateLt
}

// SetCommitRateLt sets the CommitRateLt property
func (m *CircuitsCircuitsListQueryParameters) SetCommitRateLt(val []int32) {
	m.CommitRateLt = val
}

// GetCommitRateLte returns the CommitRateLte property
func (m CircuitsCircuitsListQueryParameters) GetCommitRateLte() []int32 {
	return m.CommitRateLte
}

// SetCommitRateLte sets the CommitRateLte property
func (m *CircuitsCircuitsListQueryParameters) SetCommitRateLte(val []int32) {
	m.CommitRateLte = val
}

// GetCommitRateN returns the CommitRateN property
func (m CircuitsCircuitsListQueryParameters) GetCommitRateN() []int32 {
	return m.CommitRateN
}

// SetCommitRateN sets the CommitRateN property
func (m *CircuitsCircuitsListQueryParameters) SetCommitRateN(val []int32) {
	m.CommitRateN = val
}

// GetContact returns the Contact property
func (m CircuitsCircuitsListQueryParameters) GetContact() []int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *CircuitsCircuitsListQueryParameters) SetContact(val []int32) {
	m.Contact = val
}

// GetContactN returns the ContactN property
func (m CircuitsCircuitsListQueryParameters) GetContactN() []int32 {
	return m.ContactN
}

// SetContactN sets the ContactN property
func (m *CircuitsCircuitsListQueryParameters) SetContactN(val []int32) {
	m.ContactN = val
}

// GetContactGroup returns the ContactGroup property
func (m CircuitsCircuitsListQueryParameters) GetContactGroup() []int32 {
	return m.ContactGroup
}

// SetContactGroup sets the ContactGroup property
func (m *CircuitsCircuitsListQueryParameters) SetContactGroup(val []int32) {
	m.ContactGroup = val
}

// GetContactGroupN returns the ContactGroupN property
func (m CircuitsCircuitsListQueryParameters) GetContactGroupN() []int32 {
	return m.ContactGroupN
}

// SetContactGroupN sets the ContactGroupN property
func (m *CircuitsCircuitsListQueryParameters) SetContactGroupN(val []int32) {
	m.ContactGroupN = val
}

// GetContactRole returns the ContactRole property
func (m CircuitsCircuitsListQueryParameters) GetContactRole() []int32 {
	return m.ContactRole
}

// SetContactRole sets the ContactRole property
func (m *CircuitsCircuitsListQueryParameters) SetContactRole(val []int32) {
	m.ContactRole = val
}

// GetContactRoleN returns the ContactRoleN property
func (m CircuitsCircuitsListQueryParameters) GetContactRoleN() []int32 {
	return m.ContactRoleN
}

// SetContactRoleN sets the ContactRoleN property
func (m *CircuitsCircuitsListQueryParameters) SetContactRoleN(val []int32) {
	m.ContactRoleN = val
}

// GetCreated returns the Created property
func (m CircuitsCircuitsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CircuitsCircuitsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m CircuitsCircuitsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *CircuitsCircuitsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m CircuitsCircuitsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *CircuitsCircuitsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m CircuitsCircuitsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *CircuitsCircuitsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m CircuitsCircuitsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *CircuitsCircuitsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m CircuitsCircuitsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *CircuitsCircuitsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m CircuitsCircuitsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *CircuitsCircuitsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m CircuitsCircuitsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitsCircuitsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m CircuitsCircuitsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *CircuitsCircuitsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m CircuitsCircuitsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitsCircuitsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m CircuitsCircuitsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *CircuitsCircuitsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m CircuitsCircuitsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *CircuitsCircuitsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m CircuitsCircuitsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *CircuitsCircuitsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m CircuitsCircuitsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *CircuitsCircuitsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m CircuitsCircuitsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *CircuitsCircuitsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInstallDate returns the InstallDate property
func (m CircuitsCircuitsListQueryParameters) GetInstallDate() []string {
	return m.InstallDate
}

// SetInstallDate sets the InstallDate property
func (m *CircuitsCircuitsListQueryParameters) SetInstallDate(val []string) {
	m.InstallDate = val
}

// GetInstallDateGt returns the InstallDateGt property
func (m CircuitsCircuitsListQueryParameters) GetInstallDateGt() []string {
	return m.InstallDateGt
}

// SetInstallDateGt sets the InstallDateGt property
func (m *CircuitsCircuitsListQueryParameters) SetInstallDateGt(val []string) {
	m.InstallDateGt = val
}

// GetInstallDateGte returns the InstallDateGte property
func (m CircuitsCircuitsListQueryParameters) GetInstallDateGte() []string {
	return m.InstallDateGte
}

// SetInstallDateGte sets the InstallDateGte property
func (m *CircuitsCircuitsListQueryParameters) SetInstallDateGte(val []string) {
	m.InstallDateGte = val
}

// GetInstallDateLt returns the InstallDateLt property
func (m CircuitsCircuitsListQueryParameters) GetInstallDateLt() []string {
	return m.InstallDateLt
}

// SetInstallDateLt sets the InstallDateLt property
func (m *CircuitsCircuitsListQueryParameters) SetInstallDateLt(val []string) {
	m.InstallDateLt = val
}

// GetInstallDateLte returns the InstallDateLte property
func (m CircuitsCircuitsListQueryParameters) GetInstallDateLte() []string {
	return m.InstallDateLte
}

// SetInstallDateLte sets the InstallDateLte property
func (m *CircuitsCircuitsListQueryParameters) SetInstallDateLte(val []string) {
	m.InstallDateLte = val
}

// GetInstallDateN returns the InstallDateN property
func (m CircuitsCircuitsListQueryParameters) GetInstallDateN() []string {
	return m.InstallDateN
}

// SetInstallDateN sets the InstallDateN property
func (m *CircuitsCircuitsListQueryParameters) SetInstallDateN(val []string) {
	m.InstallDateN = val
}

// GetLastUpdated returns the LastUpdated property
func (m CircuitsCircuitsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CircuitsCircuitsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m CircuitsCircuitsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *CircuitsCircuitsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m CircuitsCircuitsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *CircuitsCircuitsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m CircuitsCircuitsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *CircuitsCircuitsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m CircuitsCircuitsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *CircuitsCircuitsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m CircuitsCircuitsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *CircuitsCircuitsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m CircuitsCircuitsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *CircuitsCircuitsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m CircuitsCircuitsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *CircuitsCircuitsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m CircuitsCircuitsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *CircuitsCircuitsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetProvider returns the Provider property
func (m CircuitsCircuitsListQueryParameters) GetProvider() []string {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *CircuitsCircuitsListQueryParameters) SetProvider(val []string) {
	m.Provider = val
}

// GetProviderN returns the ProviderN property
func (m CircuitsCircuitsListQueryParameters) GetProviderN() []string {
	return m.ProviderN
}

// SetProviderN sets the ProviderN property
func (m *CircuitsCircuitsListQueryParameters) SetProviderN(val []string) {
	m.ProviderN = val
}

// GetProviderAccountId returns the ProviderAccountId property
func (m CircuitsCircuitsListQueryParameters) GetProviderAccountId() []int32 {
	return m.ProviderAccountId
}

// SetProviderAccountId sets the ProviderAccountId property
func (m *CircuitsCircuitsListQueryParameters) SetProviderAccountId(val []int32) {
	m.ProviderAccountId = val
}

// GetProviderAccountIdN returns the ProviderAccountIdN property
func (m CircuitsCircuitsListQueryParameters) GetProviderAccountIdN() []int32 {
	return m.ProviderAccountIdN
}

// SetProviderAccountIdN sets the ProviderAccountIdN property
func (m *CircuitsCircuitsListQueryParameters) SetProviderAccountIdN(val []int32) {
	m.ProviderAccountIdN = val
}

// GetProviderId returns the ProviderId property
func (m CircuitsCircuitsListQueryParameters) GetProviderId() []int32 {
	return m.ProviderId
}

// SetProviderId sets the ProviderId property
func (m *CircuitsCircuitsListQueryParameters) SetProviderId(val []int32) {
	m.ProviderId = val
}

// GetProviderIdN returns the ProviderIdN property
func (m CircuitsCircuitsListQueryParameters) GetProviderIdN() []int32 {
	return m.ProviderIdN
}

// SetProviderIdN sets the ProviderIdN property
func (m *CircuitsCircuitsListQueryParameters) SetProviderIdN(val []int32) {
	m.ProviderIdN = val
}

// GetProviderNetworkId returns the ProviderNetworkId property
func (m CircuitsCircuitsListQueryParameters) GetProviderNetworkId() []int32 {
	return m.ProviderNetworkId
}

// SetProviderNetworkId sets the ProviderNetworkId property
func (m *CircuitsCircuitsListQueryParameters) SetProviderNetworkId(val []int32) {
	m.ProviderNetworkId = val
}

// GetProviderNetworkIdN returns the ProviderNetworkIdN property
func (m CircuitsCircuitsListQueryParameters) GetProviderNetworkIdN() []int32 {
	return m.ProviderNetworkIdN
}

// SetProviderNetworkIdN sets the ProviderNetworkIdN property
func (m *CircuitsCircuitsListQueryParameters) SetProviderNetworkIdN(val []int32) {
	m.ProviderNetworkIdN = val
}

// GetQ returns the Q property
func (m CircuitsCircuitsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *CircuitsCircuitsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m CircuitsCircuitsListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *CircuitsCircuitsListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m CircuitsCircuitsListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *CircuitsCircuitsListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m CircuitsCircuitsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *CircuitsCircuitsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m CircuitsCircuitsListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *CircuitsCircuitsListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSite returns the Site property
func (m CircuitsCircuitsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *CircuitsCircuitsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m CircuitsCircuitsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *CircuitsCircuitsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m CircuitsCircuitsListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *CircuitsCircuitsListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m CircuitsCircuitsListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *CircuitsCircuitsListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m CircuitsCircuitsListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *CircuitsCircuitsListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m CircuitsCircuitsListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *CircuitsCircuitsListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m CircuitsCircuitsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *CircuitsCircuitsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m CircuitsCircuitsListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *CircuitsCircuitsListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetStatus returns the Status property
func (m CircuitsCircuitsListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *CircuitsCircuitsListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m CircuitsCircuitsListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *CircuitsCircuitsListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m CircuitsCircuitsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *CircuitsCircuitsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m CircuitsCircuitsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *CircuitsCircuitsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m CircuitsCircuitsListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *CircuitsCircuitsListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m CircuitsCircuitsListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *CircuitsCircuitsListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m CircuitsCircuitsListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *CircuitsCircuitsListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m CircuitsCircuitsListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *CircuitsCircuitsListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m CircuitsCircuitsListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *CircuitsCircuitsListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m CircuitsCircuitsListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *CircuitsCircuitsListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m CircuitsCircuitsListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *CircuitsCircuitsListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m CircuitsCircuitsListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *CircuitsCircuitsListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetTerminationDate returns the TerminationDate property
func (m CircuitsCircuitsListQueryParameters) GetTerminationDate() []string {
	return m.TerminationDate
}

// SetTerminationDate sets the TerminationDate property
func (m *CircuitsCircuitsListQueryParameters) SetTerminationDate(val []string) {
	m.TerminationDate = val
}

// GetTerminationDateGt returns the TerminationDateGt property
func (m CircuitsCircuitsListQueryParameters) GetTerminationDateGt() []string {
	return m.TerminationDateGt
}

// SetTerminationDateGt sets the TerminationDateGt property
func (m *CircuitsCircuitsListQueryParameters) SetTerminationDateGt(val []string) {
	m.TerminationDateGt = val
}

// GetTerminationDateGte returns the TerminationDateGte property
func (m CircuitsCircuitsListQueryParameters) GetTerminationDateGte() []string {
	return m.TerminationDateGte
}

// SetTerminationDateGte sets the TerminationDateGte property
func (m *CircuitsCircuitsListQueryParameters) SetTerminationDateGte(val []string) {
	m.TerminationDateGte = val
}

// GetTerminationDateLt returns the TerminationDateLt property
func (m CircuitsCircuitsListQueryParameters) GetTerminationDateLt() []string {
	return m.TerminationDateLt
}

// SetTerminationDateLt sets the TerminationDateLt property
func (m *CircuitsCircuitsListQueryParameters) SetTerminationDateLt(val []string) {
	m.TerminationDateLt = val
}

// GetTerminationDateLte returns the TerminationDateLte property
func (m CircuitsCircuitsListQueryParameters) GetTerminationDateLte() []string {
	return m.TerminationDateLte
}

// SetTerminationDateLte sets the TerminationDateLte property
func (m *CircuitsCircuitsListQueryParameters) SetTerminationDateLte(val []string) {
	m.TerminationDateLte = val
}

// GetTerminationDateN returns the TerminationDateN property
func (m CircuitsCircuitsListQueryParameters) GetTerminationDateN() []string {
	return m.TerminationDateN
}

// SetTerminationDateN sets the TerminationDateN property
func (m *CircuitsCircuitsListQueryParameters) SetTerminationDateN(val []string) {
	m.TerminationDateN = val
}

// GetType returns the Type property
func (m CircuitsCircuitsListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *CircuitsCircuitsListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m CircuitsCircuitsListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *CircuitsCircuitsListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetTypeId returns the TypeId property
func (m CircuitsCircuitsListQueryParameters) GetTypeId() []int32 {
	return m.TypeId
}

// SetTypeId sets the TypeId property
func (m *CircuitsCircuitsListQueryParameters) SetTypeId(val []int32) {
	m.TypeId = val
}

// GetTypeIdN returns the TypeIdN property
func (m CircuitsCircuitsListQueryParameters) GetTypeIdN() []int32 {
	return m.TypeIdN
}

// SetTypeIdN sets the TypeIdN property
func (m *CircuitsCircuitsListQueryParameters) SetTypeIdN(val []int32) {
	m.TypeIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m CircuitsCircuitsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *CircuitsCircuitsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
