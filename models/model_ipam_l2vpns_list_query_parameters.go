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

// IpamL2vpnsListQueryParameters is an object.
type IpamL2vpnsListQueryParameters struct {
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
	// ExportTarget: Export target (name)
	ExportTarget []string `json:"export_target,omitempty" mapstructure:"export_target,omitempty"`
	// ExportTargetN: Export target (name)
	ExportTargetN []string `json:"export_target__n,omitempty" mapstructure:"export_target__n,omitempty"`
	// ExportTargetId: Export target
	ExportTargetId []int32 `json:"export_target_id,omitempty" mapstructure:"export_target_id,omitempty"`
	// ExportTargetIdN: Export target
	ExportTargetIdN []int32 `json:"export_target_id__n,omitempty" mapstructure:"export_target_id__n,omitempty"`
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
	// Identifier:
	Identifier []int32 `json:"identifier,omitempty" mapstructure:"identifier,omitempty"`
	// IdentifierGt:
	IdentifierGt []int32 `json:"identifier__gt,omitempty" mapstructure:"identifier__gt,omitempty"`
	// IdentifierGte:
	IdentifierGte []int32 `json:"identifier__gte,omitempty" mapstructure:"identifier__gte,omitempty"`
	// IdentifierLt:
	IdentifierLt []int32 `json:"identifier__lt,omitempty" mapstructure:"identifier__lt,omitempty"`
	// IdentifierLte:
	IdentifierLte []int32 `json:"identifier__lte,omitempty" mapstructure:"identifier__lte,omitempty"`
	// IdentifierN:
	IdentifierN []int32 `json:"identifier__n,omitempty" mapstructure:"identifier__n,omitempty"`
	// ImportTarget: Import target (name)
	ImportTarget []string `json:"import_target,omitempty" mapstructure:"import_target,omitempty"`
	// ImportTargetN: Import target (name)
	ImportTargetN []string `json:"import_target__n,omitempty" mapstructure:"import_target__n,omitempty"`
	// ImportTargetId: Import target
	ImportTargetId []int32 `json:"import_target_id,omitempty" mapstructure:"import_target_id,omitempty"`
	// ImportTargetIdN: Import target
	ImportTargetIdN []int32 `json:"import_target_id__n,omitempty" mapstructure:"import_target_id__n,omitempty"`
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
	// Slug:
	Slug []string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// SlugEmpty:
	SlugEmpty []string `json:"slug__empty,omitempty" mapstructure:"slug__empty,omitempty"`
	// SlugIc:
	SlugIc []string `json:"slug__ic,omitempty" mapstructure:"slug__ic,omitempty"`
	// SlugIe:
	SlugIe []string `json:"slug__ie,omitempty" mapstructure:"slug__ie,omitempty"`
	// SlugIew:
	SlugIew []string `json:"slug__iew,omitempty" mapstructure:"slug__iew,omitempty"`
	// SlugIsw:
	SlugIsw []string `json:"slug__isw,omitempty" mapstructure:"slug__isw,omitempty"`
	// SlugN:
	SlugN []string `json:"slug__n,omitempty" mapstructure:"slug__n,omitempty"`
	// SlugNic:
	SlugNic []string `json:"slug__nic,omitempty" mapstructure:"slug__nic,omitempty"`
	// SlugNie:
	SlugNie []string `json:"slug__nie,omitempty" mapstructure:"slug__nie,omitempty"`
	// SlugNiew:
	SlugNiew []string `json:"slug__niew,omitempty" mapstructure:"slug__niew,omitempty"`
	// SlugNisw:
	SlugNisw []string `json:"slug__nisw,omitempty" mapstructure:"slug__nisw,omitempty"`
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
	// Type: * `vpws` - VPWS
	// * `vpls` - VPLS
	// * `vxlan` - VXLAN
	// * `vxlan-evpn` - VXLAN-EVPN
	// * `mpls-evpn` - MPLS EVPN
	// * `pbb-evpn` - PBB EVPN
	// * `epl` - EPL
	// * `evpl` - EVPL
	// * `ep-lan` - Ethernet Private LAN
	// * `evp-lan` - Ethernet Virtual Private LAN
	// * `ep-tree` - Ethernet Private Tree
	// * `evp-tree` - Ethernet Virtual Private Tree
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `vpws` - VPWS
	// * `vpls` - VPLS
	// * `vxlan` - VXLAN
	// * `vxlan-evpn` - VXLAN-EVPN
	// * `mpls-evpn` - MPLS EVPN
	// * `pbb-evpn` - PBB EVPN
	// * `epl` - EPL
	// * `evpl` - EVPL
	// * `ep-lan` - Ethernet Private LAN
	// * `evp-lan` - Ethernet Virtual Private LAN
	// * `ep-tree` - Ethernet Private Tree
	// * `evp-tree` - Ethernet Virtual Private Tree
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamL2vpnsListQueryParameters) Validate() error {
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
		"exportTarget": validation.Validate(
			m.ExportTarget,
		),
		"exportTargetN": validation.Validate(
			m.ExportTargetN,
		),
		"exportTargetId": validation.Validate(
			m.ExportTargetId,
		),
		"exportTargetIdN": validation.Validate(
			m.ExportTargetIdN,
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
		"identifier": validation.Validate(
			m.Identifier,
		),
		"identifierGt": validation.Validate(
			m.IdentifierGt,
		),
		"identifierGte": validation.Validate(
			m.IdentifierGte,
		),
		"identifierLt": validation.Validate(
			m.IdentifierLt,
		),
		"identifierLte": validation.Validate(
			m.IdentifierLte,
		),
		"identifierN": validation.Validate(
			m.IdentifierN,
		),
		"importTarget": validation.Validate(
			m.ImportTarget,
		),
		"importTargetN": validation.Validate(
			m.ImportTargetN,
		),
		"importTargetId": validation.Validate(
			m.ImportTargetId,
		),
		"importTargetIdN": validation.Validate(
			m.ImportTargetIdN,
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
		"slug": validation.Validate(
			m.Slug,
		),
		"slugEmpty": validation.Validate(
			m.SlugEmpty,
		),
		"slugIc": validation.Validate(
			m.SlugIc,
		),
		"slugIe": validation.Validate(
			m.SlugIe,
		),
		"slugIew": validation.Validate(
			m.SlugIew,
		),
		"slugIsw": validation.Validate(
			m.SlugIsw,
		),
		"slugN": validation.Validate(
			m.SlugN,
		),
		"slugNic": validation.Validate(
			m.SlugNic,
		),
		"slugNie": validation.Validate(
			m.SlugNie,
		),
		"slugNiew": validation.Validate(
			m.SlugNiew,
		),
		"slugNisw": validation.Validate(
			m.SlugNisw,
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
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m IpamL2vpnsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamL2vpnsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamL2vpnsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamL2vpnsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamL2vpnsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamL2vpnsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamL2vpnsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamL2vpnsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamL2vpnsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamL2vpnsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamL2vpnsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamL2vpnsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamL2vpnsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamL2vpnsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m IpamL2vpnsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IpamL2vpnsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m IpamL2vpnsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m IpamL2vpnsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m IpamL2vpnsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m IpamL2vpnsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m IpamL2vpnsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m IpamL2vpnsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m IpamL2vpnsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m IpamL2vpnsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m IpamL2vpnsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m IpamL2vpnsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *IpamL2vpnsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetExportTarget returns the ExportTarget property
func (m IpamL2vpnsListQueryParameters) GetExportTarget() []string {
	return m.ExportTarget
}

// SetExportTarget sets the ExportTarget property
func (m *IpamL2vpnsListQueryParameters) SetExportTarget(val []string) {
	m.ExportTarget = val
}

// GetExportTargetN returns the ExportTargetN property
func (m IpamL2vpnsListQueryParameters) GetExportTargetN() []string {
	return m.ExportTargetN
}

// SetExportTargetN sets the ExportTargetN property
func (m *IpamL2vpnsListQueryParameters) SetExportTargetN(val []string) {
	m.ExportTargetN = val
}

// GetExportTargetId returns the ExportTargetId property
func (m IpamL2vpnsListQueryParameters) GetExportTargetId() []int32 {
	return m.ExportTargetId
}

// SetExportTargetId sets the ExportTargetId property
func (m *IpamL2vpnsListQueryParameters) SetExportTargetId(val []int32) {
	m.ExportTargetId = val
}

// GetExportTargetIdN returns the ExportTargetIdN property
func (m IpamL2vpnsListQueryParameters) GetExportTargetIdN() []int32 {
	return m.ExportTargetIdN
}

// SetExportTargetIdN sets the ExportTargetIdN property
func (m *IpamL2vpnsListQueryParameters) SetExportTargetIdN(val []int32) {
	m.ExportTargetIdN = val
}

// GetId returns the Id property
func (m IpamL2vpnsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamL2vpnsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamL2vpnsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamL2vpnsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamL2vpnsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamL2vpnsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamL2vpnsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamL2vpnsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamL2vpnsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamL2vpnsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamL2vpnsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamL2vpnsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetIdentifier returns the Identifier property
func (m IpamL2vpnsListQueryParameters) GetIdentifier() []int32 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *IpamL2vpnsListQueryParameters) SetIdentifier(val []int32) {
	m.Identifier = val
}

// GetIdentifierGt returns the IdentifierGt property
func (m IpamL2vpnsListQueryParameters) GetIdentifierGt() []int32 {
	return m.IdentifierGt
}

// SetIdentifierGt sets the IdentifierGt property
func (m *IpamL2vpnsListQueryParameters) SetIdentifierGt(val []int32) {
	m.IdentifierGt = val
}

// GetIdentifierGte returns the IdentifierGte property
func (m IpamL2vpnsListQueryParameters) GetIdentifierGte() []int32 {
	return m.IdentifierGte
}

// SetIdentifierGte sets the IdentifierGte property
func (m *IpamL2vpnsListQueryParameters) SetIdentifierGte(val []int32) {
	m.IdentifierGte = val
}

// GetIdentifierLt returns the IdentifierLt property
func (m IpamL2vpnsListQueryParameters) GetIdentifierLt() []int32 {
	return m.IdentifierLt
}

// SetIdentifierLt sets the IdentifierLt property
func (m *IpamL2vpnsListQueryParameters) SetIdentifierLt(val []int32) {
	m.IdentifierLt = val
}

// GetIdentifierLte returns the IdentifierLte property
func (m IpamL2vpnsListQueryParameters) GetIdentifierLte() []int32 {
	return m.IdentifierLte
}

// SetIdentifierLte sets the IdentifierLte property
func (m *IpamL2vpnsListQueryParameters) SetIdentifierLte(val []int32) {
	m.IdentifierLte = val
}

// GetIdentifierN returns the IdentifierN property
func (m IpamL2vpnsListQueryParameters) GetIdentifierN() []int32 {
	return m.IdentifierN
}

// SetIdentifierN sets the IdentifierN property
func (m *IpamL2vpnsListQueryParameters) SetIdentifierN(val []int32) {
	m.IdentifierN = val
}

// GetImportTarget returns the ImportTarget property
func (m IpamL2vpnsListQueryParameters) GetImportTarget() []string {
	return m.ImportTarget
}

// SetImportTarget sets the ImportTarget property
func (m *IpamL2vpnsListQueryParameters) SetImportTarget(val []string) {
	m.ImportTarget = val
}

// GetImportTargetN returns the ImportTargetN property
func (m IpamL2vpnsListQueryParameters) GetImportTargetN() []string {
	return m.ImportTargetN
}

// SetImportTargetN sets the ImportTargetN property
func (m *IpamL2vpnsListQueryParameters) SetImportTargetN(val []string) {
	m.ImportTargetN = val
}

// GetImportTargetId returns the ImportTargetId property
func (m IpamL2vpnsListQueryParameters) GetImportTargetId() []int32 {
	return m.ImportTargetId
}

// SetImportTargetId sets the ImportTargetId property
func (m *IpamL2vpnsListQueryParameters) SetImportTargetId(val []int32) {
	m.ImportTargetId = val
}

// GetImportTargetIdN returns the ImportTargetIdN property
func (m IpamL2vpnsListQueryParameters) GetImportTargetIdN() []int32 {
	return m.ImportTargetIdN
}

// SetImportTargetIdN sets the ImportTargetIdN property
func (m *IpamL2vpnsListQueryParameters) SetImportTargetIdN(val []int32) {
	m.ImportTargetIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamL2vpnsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamL2vpnsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamL2vpnsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamL2vpnsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamL2vpnsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamL2vpnsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamL2vpnsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamL2vpnsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamL2vpnsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamL2vpnsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamL2vpnsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamL2vpnsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamL2vpnsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamL2vpnsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m IpamL2vpnsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *IpamL2vpnsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m IpamL2vpnsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *IpamL2vpnsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m IpamL2vpnsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *IpamL2vpnsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m IpamL2vpnsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *IpamL2vpnsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m IpamL2vpnsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *IpamL2vpnsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m IpamL2vpnsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *IpamL2vpnsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m IpamL2vpnsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *IpamL2vpnsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m IpamL2vpnsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *IpamL2vpnsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m IpamL2vpnsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *IpamL2vpnsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m IpamL2vpnsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *IpamL2vpnsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m IpamL2vpnsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *IpamL2vpnsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m IpamL2vpnsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamL2vpnsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamL2vpnsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamL2vpnsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m IpamL2vpnsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamL2vpnsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSlug returns the Slug property
func (m IpamL2vpnsListQueryParameters) GetSlug() []string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *IpamL2vpnsListQueryParameters) SetSlug(val []string) {
	m.Slug = val
}

// GetSlugEmpty returns the SlugEmpty property
func (m IpamL2vpnsListQueryParameters) GetSlugEmpty() []string {
	return m.SlugEmpty
}

// SetSlugEmpty sets the SlugEmpty property
func (m *IpamL2vpnsListQueryParameters) SetSlugEmpty(val []string) {
	m.SlugEmpty = val
}

// GetSlugIc returns the SlugIc property
func (m IpamL2vpnsListQueryParameters) GetSlugIc() []string {
	return m.SlugIc
}

// SetSlugIc sets the SlugIc property
func (m *IpamL2vpnsListQueryParameters) SetSlugIc(val []string) {
	m.SlugIc = val
}

// GetSlugIe returns the SlugIe property
func (m IpamL2vpnsListQueryParameters) GetSlugIe() []string {
	return m.SlugIe
}

// SetSlugIe sets the SlugIe property
func (m *IpamL2vpnsListQueryParameters) SetSlugIe(val []string) {
	m.SlugIe = val
}

// GetSlugIew returns the SlugIew property
func (m IpamL2vpnsListQueryParameters) GetSlugIew() []string {
	return m.SlugIew
}

// SetSlugIew sets the SlugIew property
func (m *IpamL2vpnsListQueryParameters) SetSlugIew(val []string) {
	m.SlugIew = val
}

// GetSlugIsw returns the SlugIsw property
func (m IpamL2vpnsListQueryParameters) GetSlugIsw() []string {
	return m.SlugIsw
}

// SetSlugIsw sets the SlugIsw property
func (m *IpamL2vpnsListQueryParameters) SetSlugIsw(val []string) {
	m.SlugIsw = val
}

// GetSlugN returns the SlugN property
func (m IpamL2vpnsListQueryParameters) GetSlugN() []string {
	return m.SlugN
}

// SetSlugN sets the SlugN property
func (m *IpamL2vpnsListQueryParameters) SetSlugN(val []string) {
	m.SlugN = val
}

// GetSlugNic returns the SlugNic property
func (m IpamL2vpnsListQueryParameters) GetSlugNic() []string {
	return m.SlugNic
}

// SetSlugNic sets the SlugNic property
func (m *IpamL2vpnsListQueryParameters) SetSlugNic(val []string) {
	m.SlugNic = val
}

// GetSlugNie returns the SlugNie property
func (m IpamL2vpnsListQueryParameters) GetSlugNie() []string {
	return m.SlugNie
}

// SetSlugNie sets the SlugNie property
func (m *IpamL2vpnsListQueryParameters) SetSlugNie(val []string) {
	m.SlugNie = val
}

// GetSlugNiew returns the SlugNiew property
func (m IpamL2vpnsListQueryParameters) GetSlugNiew() []string {
	return m.SlugNiew
}

// SetSlugNiew sets the SlugNiew property
func (m *IpamL2vpnsListQueryParameters) SetSlugNiew(val []string) {
	m.SlugNiew = val
}

// GetSlugNisw returns the SlugNisw property
func (m IpamL2vpnsListQueryParameters) GetSlugNisw() []string {
	return m.SlugNisw
}

// SetSlugNisw sets the SlugNisw property
func (m *IpamL2vpnsListQueryParameters) SetSlugNisw(val []string) {
	m.SlugNisw = val
}

// GetTag returns the Tag property
func (m IpamL2vpnsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamL2vpnsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamL2vpnsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamL2vpnsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m IpamL2vpnsListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IpamL2vpnsListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m IpamL2vpnsListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *IpamL2vpnsListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m IpamL2vpnsListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *IpamL2vpnsListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m IpamL2vpnsListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *IpamL2vpnsListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m IpamL2vpnsListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *IpamL2vpnsListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m IpamL2vpnsListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *IpamL2vpnsListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m IpamL2vpnsListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *IpamL2vpnsListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m IpamL2vpnsListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *IpamL2vpnsListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetType returns the Type property
func (m IpamL2vpnsListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *IpamL2vpnsListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m IpamL2vpnsListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *IpamL2vpnsListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamL2vpnsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamL2vpnsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
