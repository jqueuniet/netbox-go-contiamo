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

// IpamAggregatesListQueryParameters is an object.
type IpamAggregatesListQueryParameters struct {
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
	// DateAdded:
	DateAdded []string `json:"date_added,omitempty" mapstructure:"date_added,omitempty"`
	// DateAddedGt:
	DateAddedGt []string `json:"date_added__gt,omitempty" mapstructure:"date_added__gt,omitempty"`
	// DateAddedGte:
	DateAddedGte []string `json:"date_added__gte,omitempty" mapstructure:"date_added__gte,omitempty"`
	// DateAddedLt:
	DateAddedLt []string `json:"date_added__lt,omitempty" mapstructure:"date_added__lt,omitempty"`
	// DateAddedLte:
	DateAddedLte []string `json:"date_added__lte,omitempty" mapstructure:"date_added__lte,omitempty"`
	// DateAddedN:
	DateAddedN []string `json:"date_added__n,omitempty" mapstructure:"date_added__n,omitempty"`
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
	// Family:
	Family float32 `json:"family,omitempty" mapstructure:"family,omitempty"`
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
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Prefix: Prefix
	Prefix string `json:"prefix,omitempty" mapstructure:"prefix,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Rir: RIR (slug)
	Rir []string `json:"rir,omitempty" mapstructure:"rir,omitempty"`
	// RirN: RIR (slug)
	RirN []string `json:"rir__n,omitempty" mapstructure:"rir__n,omitempty"`
	// RirId: RIR (ID)
	RirId []int32 `json:"rir_id,omitempty" mapstructure:"rir_id,omitempty"`
	// RirIdN: RIR (ID)
	RirIdN []int32 `json:"rir_id__n,omitempty" mapstructure:"rir_id__n,omitempty"`
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
func (m IpamAggregatesListQueryParameters) Validate() error {
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
		"dateAdded": validation.Validate(
			m.DateAdded,
		),
		"dateAddedGt": validation.Validate(
			m.DateAddedGt,
		),
		"dateAddedGte": validation.Validate(
			m.DateAddedGte,
		),
		"dateAddedLt": validation.Validate(
			m.DateAddedLt,
		),
		"dateAddedLte": validation.Validate(
			m.DateAddedLte,
		),
		"dateAddedN": validation.Validate(
			m.DateAddedN,
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
		"rir": validation.Validate(
			m.Rir,
		),
		"rirN": validation.Validate(
			m.RirN,
		),
		"rirId": validation.Validate(
			m.RirId,
		),
		"rirIdN": validation.Validate(
			m.RirIdN,
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
func (m IpamAggregatesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamAggregatesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamAggregatesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamAggregatesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamAggregatesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamAggregatesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamAggregatesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamAggregatesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamAggregatesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamAggregatesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamAggregatesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamAggregatesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamAggregatesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamAggregatesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDateAdded returns the DateAdded property
func (m IpamAggregatesListQueryParameters) GetDateAdded() []string {
	return m.DateAdded
}

// SetDateAdded sets the DateAdded property
func (m *IpamAggregatesListQueryParameters) SetDateAdded(val []string) {
	m.DateAdded = val
}

// GetDateAddedGt returns the DateAddedGt property
func (m IpamAggregatesListQueryParameters) GetDateAddedGt() []string {
	return m.DateAddedGt
}

// SetDateAddedGt sets the DateAddedGt property
func (m *IpamAggregatesListQueryParameters) SetDateAddedGt(val []string) {
	m.DateAddedGt = val
}

// GetDateAddedGte returns the DateAddedGte property
func (m IpamAggregatesListQueryParameters) GetDateAddedGte() []string {
	return m.DateAddedGte
}

// SetDateAddedGte sets the DateAddedGte property
func (m *IpamAggregatesListQueryParameters) SetDateAddedGte(val []string) {
	m.DateAddedGte = val
}

// GetDateAddedLt returns the DateAddedLt property
func (m IpamAggregatesListQueryParameters) GetDateAddedLt() []string {
	return m.DateAddedLt
}

// SetDateAddedLt sets the DateAddedLt property
func (m *IpamAggregatesListQueryParameters) SetDateAddedLt(val []string) {
	m.DateAddedLt = val
}

// GetDateAddedLte returns the DateAddedLte property
func (m IpamAggregatesListQueryParameters) GetDateAddedLte() []string {
	return m.DateAddedLte
}

// SetDateAddedLte sets the DateAddedLte property
func (m *IpamAggregatesListQueryParameters) SetDateAddedLte(val []string) {
	m.DateAddedLte = val
}

// GetDateAddedN returns the DateAddedN property
func (m IpamAggregatesListQueryParameters) GetDateAddedN() []string {
	return m.DateAddedN
}

// SetDateAddedN sets the DateAddedN property
func (m *IpamAggregatesListQueryParameters) SetDateAddedN(val []string) {
	m.DateAddedN = val
}

// GetDescription returns the Description property
func (m IpamAggregatesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IpamAggregatesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m IpamAggregatesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *IpamAggregatesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m IpamAggregatesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *IpamAggregatesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m IpamAggregatesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *IpamAggregatesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m IpamAggregatesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *IpamAggregatesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m IpamAggregatesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *IpamAggregatesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m IpamAggregatesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *IpamAggregatesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m IpamAggregatesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *IpamAggregatesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m IpamAggregatesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *IpamAggregatesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m IpamAggregatesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *IpamAggregatesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m IpamAggregatesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *IpamAggregatesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetFamily returns the Family property
func (m IpamAggregatesListQueryParameters) GetFamily() float32 {
	return m.Family
}

// SetFamily sets the Family property
func (m *IpamAggregatesListQueryParameters) SetFamily(val float32) {
	m.Family = val
}

// GetId returns the Id property
func (m IpamAggregatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamAggregatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamAggregatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamAggregatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamAggregatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamAggregatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamAggregatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamAggregatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamAggregatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamAggregatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamAggregatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamAggregatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamAggregatesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamAggregatesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamAggregatesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamAggregatesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamAggregatesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamAggregatesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamAggregatesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamAggregatesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamAggregatesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamAggregatesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamAggregatesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamAggregatesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamAggregatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamAggregatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m IpamAggregatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamAggregatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamAggregatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamAggregatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPrefix returns the Prefix property
func (m IpamAggregatesListQueryParameters) GetPrefix() string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *IpamAggregatesListQueryParameters) SetPrefix(val string) {
	m.Prefix = val
}

// GetQ returns the Q property
func (m IpamAggregatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamAggregatesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRir returns the Rir property
func (m IpamAggregatesListQueryParameters) GetRir() []string {
	return m.Rir
}

// SetRir sets the Rir property
func (m *IpamAggregatesListQueryParameters) SetRir(val []string) {
	m.Rir = val
}

// GetRirN returns the RirN property
func (m IpamAggregatesListQueryParameters) GetRirN() []string {
	return m.RirN
}

// SetRirN sets the RirN property
func (m *IpamAggregatesListQueryParameters) SetRirN(val []string) {
	m.RirN = val
}

// GetRirId returns the RirId property
func (m IpamAggregatesListQueryParameters) GetRirId() []int32 {
	return m.RirId
}

// SetRirId sets the RirId property
func (m *IpamAggregatesListQueryParameters) SetRirId(val []int32) {
	m.RirId = val
}

// GetRirIdN returns the RirIdN property
func (m IpamAggregatesListQueryParameters) GetRirIdN() []int32 {
	return m.RirIdN
}

// SetRirIdN sets the RirIdN property
func (m *IpamAggregatesListQueryParameters) SetRirIdN(val []int32) {
	m.RirIdN = val
}

// GetTag returns the Tag property
func (m IpamAggregatesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamAggregatesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamAggregatesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamAggregatesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m IpamAggregatesListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IpamAggregatesListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m IpamAggregatesListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *IpamAggregatesListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m IpamAggregatesListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *IpamAggregatesListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m IpamAggregatesListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *IpamAggregatesListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m IpamAggregatesListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *IpamAggregatesListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m IpamAggregatesListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *IpamAggregatesListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m IpamAggregatesListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *IpamAggregatesListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m IpamAggregatesListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *IpamAggregatesListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamAggregatesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamAggregatesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
