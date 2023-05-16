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

// IpamRouteTargetsListQueryParameters is an object.
type IpamRouteTargetsListQueryParameters struct {
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
	// ExportingVrf: Export VRF (RD)
	ExportingVrf []*string `json:"exporting_vrf,omitempty" mapstructure:"exporting_vrf,omitempty"`
	// ExportingVrfN: Export VRF (RD)
	ExportingVrfN []*string `json:"exporting_vrf__n,omitempty" mapstructure:"exporting_vrf__n,omitempty"`
	// ExportingVrfId: Exporting VRF
	ExportingVrfId []int32 `json:"exporting_vrf_id,omitempty" mapstructure:"exporting_vrf_id,omitempty"`
	// ExportingVrfIdN: Exporting VRF
	ExportingVrfIdN []int32 `json:"exporting_vrf_id__n,omitempty" mapstructure:"exporting_vrf_id__n,omitempty"`
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
	// ImportingVrf: Import VRF (RD)
	ImportingVrf []*string `json:"importing_vrf,omitempty" mapstructure:"importing_vrf,omitempty"`
	// ImportingVrfN: Import VRF (RD)
	ImportingVrfN []*string `json:"importing_vrf__n,omitempty" mapstructure:"importing_vrf__n,omitempty"`
	// ImportingVrfId: Importing VRF
	ImportingVrfId []int32 `json:"importing_vrf_id,omitempty" mapstructure:"importing_vrf_id,omitempty"`
	// ImportingVrfIdN: Importing VRF
	ImportingVrfIdN []int32 `json:"importing_vrf_id__n,omitempty" mapstructure:"importing_vrf_id__n,omitempty"`
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
}

// Validate implements basic validation for this model
func (m IpamRouteTargetsListQueryParameters) Validate() error {
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
		"exportingVrf": validation.Validate(
			m.ExportingVrf,
		),
		"exportingVrfN": validation.Validate(
			m.ExportingVrfN,
		),
		"exportingVrfId": validation.Validate(
			m.ExportingVrfId,
		),
		"exportingVrfIdN": validation.Validate(
			m.ExportingVrfIdN,
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
		"importingVrf": validation.Validate(
			m.ImportingVrf,
		),
		"importingVrfN": validation.Validate(
			m.ImportingVrfN,
		),
		"importingVrfId": validation.Validate(
			m.ImportingVrfId,
		),
		"importingVrfIdN": validation.Validate(
			m.ImportingVrfIdN,
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
	}.Filter()
}

// GetCreated returns the Created property
func (m IpamRouteTargetsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamRouteTargetsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamRouteTargetsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamRouteTargetsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamRouteTargetsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamRouteTargetsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamRouteTargetsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamRouteTargetsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamRouteTargetsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamRouteTargetsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamRouteTargetsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamRouteTargetsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamRouteTargetsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamRouteTargetsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m IpamRouteTargetsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IpamRouteTargetsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m IpamRouteTargetsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *IpamRouteTargetsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetExportingVrf returns the ExportingVrf property
func (m IpamRouteTargetsListQueryParameters) GetExportingVrf() []*string {
	return m.ExportingVrf
}

// SetExportingVrf sets the ExportingVrf property
func (m *IpamRouteTargetsListQueryParameters) SetExportingVrf(val []*string) {
	m.ExportingVrf = val
}

// GetExportingVrfN returns the ExportingVrfN property
func (m IpamRouteTargetsListQueryParameters) GetExportingVrfN() []*string {
	return m.ExportingVrfN
}

// SetExportingVrfN sets the ExportingVrfN property
func (m *IpamRouteTargetsListQueryParameters) SetExportingVrfN(val []*string) {
	m.ExportingVrfN = val
}

// GetExportingVrfId returns the ExportingVrfId property
func (m IpamRouteTargetsListQueryParameters) GetExportingVrfId() []int32 {
	return m.ExportingVrfId
}

// SetExportingVrfId sets the ExportingVrfId property
func (m *IpamRouteTargetsListQueryParameters) SetExportingVrfId(val []int32) {
	m.ExportingVrfId = val
}

// GetExportingVrfIdN returns the ExportingVrfIdN property
func (m IpamRouteTargetsListQueryParameters) GetExportingVrfIdN() []int32 {
	return m.ExportingVrfIdN
}

// SetExportingVrfIdN sets the ExportingVrfIdN property
func (m *IpamRouteTargetsListQueryParameters) SetExportingVrfIdN(val []int32) {
	m.ExportingVrfIdN = val
}

// GetId returns the Id property
func (m IpamRouteTargetsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamRouteTargetsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamRouteTargetsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamRouteTargetsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamRouteTargetsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamRouteTargetsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamRouteTargetsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamRouteTargetsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamRouteTargetsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamRouteTargetsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamRouteTargetsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamRouteTargetsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetImportingVrf returns the ImportingVrf property
func (m IpamRouteTargetsListQueryParameters) GetImportingVrf() []*string {
	return m.ImportingVrf
}

// SetImportingVrf sets the ImportingVrf property
func (m *IpamRouteTargetsListQueryParameters) SetImportingVrf(val []*string) {
	m.ImportingVrf = val
}

// GetImportingVrfN returns the ImportingVrfN property
func (m IpamRouteTargetsListQueryParameters) GetImportingVrfN() []*string {
	return m.ImportingVrfN
}

// SetImportingVrfN sets the ImportingVrfN property
func (m *IpamRouteTargetsListQueryParameters) SetImportingVrfN(val []*string) {
	m.ImportingVrfN = val
}

// GetImportingVrfId returns the ImportingVrfId property
func (m IpamRouteTargetsListQueryParameters) GetImportingVrfId() []int32 {
	return m.ImportingVrfId
}

// SetImportingVrfId sets the ImportingVrfId property
func (m *IpamRouteTargetsListQueryParameters) SetImportingVrfId(val []int32) {
	m.ImportingVrfId = val
}

// GetImportingVrfIdN returns the ImportingVrfIdN property
func (m IpamRouteTargetsListQueryParameters) GetImportingVrfIdN() []int32 {
	return m.ImportingVrfIdN
}

// SetImportingVrfIdN sets the ImportingVrfIdN property
func (m *IpamRouteTargetsListQueryParameters) SetImportingVrfIdN(val []int32) {
	m.ImportingVrfIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamRouteTargetsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamRouteTargetsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamRouteTargetsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamRouteTargetsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamRouteTargetsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamRouteTargetsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamRouteTargetsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamRouteTargetsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamRouteTargetsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamRouteTargetsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamRouteTargetsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamRouteTargetsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamRouteTargetsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamRouteTargetsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m IpamRouteTargetsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *IpamRouteTargetsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m IpamRouteTargetsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *IpamRouteTargetsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m IpamRouteTargetsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *IpamRouteTargetsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m IpamRouteTargetsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *IpamRouteTargetsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m IpamRouteTargetsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *IpamRouteTargetsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m IpamRouteTargetsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *IpamRouteTargetsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m IpamRouteTargetsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *IpamRouteTargetsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m IpamRouteTargetsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *IpamRouteTargetsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m IpamRouteTargetsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *IpamRouteTargetsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m IpamRouteTargetsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *IpamRouteTargetsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m IpamRouteTargetsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *IpamRouteTargetsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m IpamRouteTargetsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamRouteTargetsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamRouteTargetsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamRouteTargetsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m IpamRouteTargetsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamRouteTargetsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m IpamRouteTargetsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamRouteTargetsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamRouteTargetsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamRouteTargetsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m IpamRouteTargetsListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IpamRouteTargetsListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m IpamRouteTargetsListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *IpamRouteTargetsListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m IpamRouteTargetsListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *IpamRouteTargetsListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m IpamRouteTargetsListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *IpamRouteTargetsListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m IpamRouteTargetsListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *IpamRouteTargetsListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m IpamRouteTargetsListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *IpamRouteTargetsListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m IpamRouteTargetsListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *IpamRouteTargetsListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m IpamRouteTargetsListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *IpamRouteTargetsListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamRouteTargetsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamRouteTargetsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
