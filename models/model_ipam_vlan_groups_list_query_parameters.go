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

// IpamVlanGroupsListQueryParameters is an object.
type IpamVlanGroupsListQueryParameters struct {
	// Cluster:
	Cluster int32 `json:"cluster,omitempty" mapstructure:"cluster,omitempty"`
	// Clustergroup:
	Clustergroup float32 `json:"clustergroup,omitempty" mapstructure:"clustergroup,omitempty"`
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
	// Location:
	Location int32 `json:"location,omitempty" mapstructure:"location,omitempty"`
	// MaxVid:
	MaxVid []int32 `json:"max_vid,omitempty" mapstructure:"max_vid,omitempty"`
	// MaxVidGt:
	MaxVidGt []int32 `json:"max_vid__gt,omitempty" mapstructure:"max_vid__gt,omitempty"`
	// MaxVidGte:
	MaxVidGte []int32 `json:"max_vid__gte,omitempty" mapstructure:"max_vid__gte,omitempty"`
	// MaxVidLt:
	MaxVidLt []int32 `json:"max_vid__lt,omitempty" mapstructure:"max_vid__lt,omitempty"`
	// MaxVidLte:
	MaxVidLte []int32 `json:"max_vid__lte,omitempty" mapstructure:"max_vid__lte,omitempty"`
	// MaxVidN:
	MaxVidN []int32 `json:"max_vid__n,omitempty" mapstructure:"max_vid__n,omitempty"`
	// MinVid:
	MinVid []int32 `json:"min_vid,omitempty" mapstructure:"min_vid,omitempty"`
	// MinVidGt:
	MinVidGt []int32 `json:"min_vid__gt,omitempty" mapstructure:"min_vid__gt,omitempty"`
	// MinVidGte:
	MinVidGte []int32 `json:"min_vid__gte,omitempty" mapstructure:"min_vid__gte,omitempty"`
	// MinVidLt:
	MinVidLt []int32 `json:"min_vid__lt,omitempty" mapstructure:"min_vid__lt,omitempty"`
	// MinVidLte:
	MinVidLte []int32 `json:"min_vid__lte,omitempty" mapstructure:"min_vid__lte,omitempty"`
	// MinVidN:
	MinVidN []int32 `json:"min_vid__n,omitempty" mapstructure:"min_vid__n,omitempty"`
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
	// Rack:
	Rack int32 `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// Region:
	Region int32 `json:"region,omitempty" mapstructure:"region,omitempty"`
	// ScopeId:
	ScopeId []int32 `json:"scope_id,omitempty" mapstructure:"scope_id,omitempty"`
	// ScopeIdGt:
	ScopeIdGt []int32 `json:"scope_id__gt,omitempty" mapstructure:"scope_id__gt,omitempty"`
	// ScopeIdGte:
	ScopeIdGte []int32 `json:"scope_id__gte,omitempty" mapstructure:"scope_id__gte,omitempty"`
	// ScopeIdLt:
	ScopeIdLt []int32 `json:"scope_id__lt,omitempty" mapstructure:"scope_id__lt,omitempty"`
	// ScopeIdLte:
	ScopeIdLte []int32 `json:"scope_id__lte,omitempty" mapstructure:"scope_id__lte,omitempty"`
	// ScopeIdN:
	ScopeIdN []int32 `json:"scope_id__n,omitempty" mapstructure:"scope_id__n,omitempty"`
	// ScopeType:
	ScopeType string `json:"scope_type,omitempty" mapstructure:"scope_type,omitempty"`
	// ScopeTypeN:
	ScopeTypeN string `json:"scope_type__n,omitempty" mapstructure:"scope_type__n,omitempty"`
	// Site:
	Site int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Sitegroup:
	Sitegroup float32 `json:"sitegroup,omitempty" mapstructure:"sitegroup,omitempty"`
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
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamVlanGroupsListQueryParameters) Validate() error {
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
		"maxVid": validation.Validate(
			m.MaxVid,
		),
		"maxVidGt": validation.Validate(
			m.MaxVidGt,
		),
		"maxVidGte": validation.Validate(
			m.MaxVidGte,
		),
		"maxVidLt": validation.Validate(
			m.MaxVidLt,
		),
		"maxVidLte": validation.Validate(
			m.MaxVidLte,
		),
		"maxVidN": validation.Validate(
			m.MaxVidN,
		),
		"minVid": validation.Validate(
			m.MinVid,
		),
		"minVidGt": validation.Validate(
			m.MinVidGt,
		),
		"minVidGte": validation.Validate(
			m.MinVidGte,
		),
		"minVidLt": validation.Validate(
			m.MinVidLt,
		),
		"minVidLte": validation.Validate(
			m.MinVidLte,
		),
		"minVidN": validation.Validate(
			m.MinVidN,
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
		"scopeId": validation.Validate(
			m.ScopeId,
		),
		"scopeIdGt": validation.Validate(
			m.ScopeIdGt,
		),
		"scopeIdGte": validation.Validate(
			m.ScopeIdGte,
		),
		"scopeIdLt": validation.Validate(
			m.ScopeIdLt,
		),
		"scopeIdLte": validation.Validate(
			m.ScopeIdLte,
		),
		"scopeIdN": validation.Validate(
			m.ScopeIdN,
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
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetCluster returns the Cluster property
func (m IpamVlanGroupsListQueryParameters) GetCluster() int32 {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *IpamVlanGroupsListQueryParameters) SetCluster(val int32) {
	m.Cluster = val
}

// GetClustergroup returns the Clustergroup property
func (m IpamVlanGroupsListQueryParameters) GetClustergroup() float32 {
	return m.Clustergroup
}

// SetClustergroup sets the Clustergroup property
func (m *IpamVlanGroupsListQueryParameters) SetClustergroup(val float32) {
	m.Clustergroup = val
}

// GetCreated returns the Created property
func (m IpamVlanGroupsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamVlanGroupsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamVlanGroupsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamVlanGroupsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamVlanGroupsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamVlanGroupsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamVlanGroupsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamVlanGroupsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamVlanGroupsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamVlanGroupsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamVlanGroupsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamVlanGroupsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamVlanGroupsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamVlanGroupsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m IpamVlanGroupsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IpamVlanGroupsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m IpamVlanGroupsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *IpamVlanGroupsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m IpamVlanGroupsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamVlanGroupsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamVlanGroupsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamVlanGroupsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamVlanGroupsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamVlanGroupsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamVlanGroupsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamVlanGroupsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamVlanGroupsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamVlanGroupsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamVlanGroupsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamVlanGroupsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamVlanGroupsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamVlanGroupsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamVlanGroupsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamVlanGroupsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamVlanGroupsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamVlanGroupsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamVlanGroupsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamVlanGroupsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamVlanGroupsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamVlanGroupsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamVlanGroupsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamVlanGroupsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamVlanGroupsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamVlanGroupsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m IpamVlanGroupsListQueryParameters) GetLocation() int32 {
	return m.Location
}

// SetLocation sets the Location property
func (m *IpamVlanGroupsListQueryParameters) SetLocation(val int32) {
	m.Location = val
}

// GetMaxVid returns the MaxVid property
func (m IpamVlanGroupsListQueryParameters) GetMaxVid() []int32 {
	return m.MaxVid
}

// SetMaxVid sets the MaxVid property
func (m *IpamVlanGroupsListQueryParameters) SetMaxVid(val []int32) {
	m.MaxVid = val
}

// GetMaxVidGt returns the MaxVidGt property
func (m IpamVlanGroupsListQueryParameters) GetMaxVidGt() []int32 {
	return m.MaxVidGt
}

// SetMaxVidGt sets the MaxVidGt property
func (m *IpamVlanGroupsListQueryParameters) SetMaxVidGt(val []int32) {
	m.MaxVidGt = val
}

// GetMaxVidGte returns the MaxVidGte property
func (m IpamVlanGroupsListQueryParameters) GetMaxVidGte() []int32 {
	return m.MaxVidGte
}

// SetMaxVidGte sets the MaxVidGte property
func (m *IpamVlanGroupsListQueryParameters) SetMaxVidGte(val []int32) {
	m.MaxVidGte = val
}

// GetMaxVidLt returns the MaxVidLt property
func (m IpamVlanGroupsListQueryParameters) GetMaxVidLt() []int32 {
	return m.MaxVidLt
}

// SetMaxVidLt sets the MaxVidLt property
func (m *IpamVlanGroupsListQueryParameters) SetMaxVidLt(val []int32) {
	m.MaxVidLt = val
}

// GetMaxVidLte returns the MaxVidLte property
func (m IpamVlanGroupsListQueryParameters) GetMaxVidLte() []int32 {
	return m.MaxVidLte
}

// SetMaxVidLte sets the MaxVidLte property
func (m *IpamVlanGroupsListQueryParameters) SetMaxVidLte(val []int32) {
	m.MaxVidLte = val
}

// GetMaxVidN returns the MaxVidN property
func (m IpamVlanGroupsListQueryParameters) GetMaxVidN() []int32 {
	return m.MaxVidN
}

// SetMaxVidN sets the MaxVidN property
func (m *IpamVlanGroupsListQueryParameters) SetMaxVidN(val []int32) {
	m.MaxVidN = val
}

// GetMinVid returns the MinVid property
func (m IpamVlanGroupsListQueryParameters) GetMinVid() []int32 {
	return m.MinVid
}

// SetMinVid sets the MinVid property
func (m *IpamVlanGroupsListQueryParameters) SetMinVid(val []int32) {
	m.MinVid = val
}

// GetMinVidGt returns the MinVidGt property
func (m IpamVlanGroupsListQueryParameters) GetMinVidGt() []int32 {
	return m.MinVidGt
}

// SetMinVidGt sets the MinVidGt property
func (m *IpamVlanGroupsListQueryParameters) SetMinVidGt(val []int32) {
	m.MinVidGt = val
}

// GetMinVidGte returns the MinVidGte property
func (m IpamVlanGroupsListQueryParameters) GetMinVidGte() []int32 {
	return m.MinVidGte
}

// SetMinVidGte sets the MinVidGte property
func (m *IpamVlanGroupsListQueryParameters) SetMinVidGte(val []int32) {
	m.MinVidGte = val
}

// GetMinVidLt returns the MinVidLt property
func (m IpamVlanGroupsListQueryParameters) GetMinVidLt() []int32 {
	return m.MinVidLt
}

// SetMinVidLt sets the MinVidLt property
func (m *IpamVlanGroupsListQueryParameters) SetMinVidLt(val []int32) {
	m.MinVidLt = val
}

// GetMinVidLte returns the MinVidLte property
func (m IpamVlanGroupsListQueryParameters) GetMinVidLte() []int32 {
	return m.MinVidLte
}

// SetMinVidLte sets the MinVidLte property
func (m *IpamVlanGroupsListQueryParameters) SetMinVidLte(val []int32) {
	m.MinVidLte = val
}

// GetMinVidN returns the MinVidN property
func (m IpamVlanGroupsListQueryParameters) GetMinVidN() []int32 {
	return m.MinVidN
}

// SetMinVidN sets the MinVidN property
func (m *IpamVlanGroupsListQueryParameters) SetMinVidN(val []int32) {
	m.MinVidN = val
}

// GetName returns the Name property
func (m IpamVlanGroupsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *IpamVlanGroupsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m IpamVlanGroupsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *IpamVlanGroupsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m IpamVlanGroupsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *IpamVlanGroupsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m IpamVlanGroupsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *IpamVlanGroupsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m IpamVlanGroupsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *IpamVlanGroupsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m IpamVlanGroupsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *IpamVlanGroupsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m IpamVlanGroupsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *IpamVlanGroupsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m IpamVlanGroupsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *IpamVlanGroupsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m IpamVlanGroupsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *IpamVlanGroupsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m IpamVlanGroupsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *IpamVlanGroupsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m IpamVlanGroupsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *IpamVlanGroupsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m IpamVlanGroupsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamVlanGroupsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamVlanGroupsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamVlanGroupsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m IpamVlanGroupsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamVlanGroupsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRack returns the Rack property
func (m IpamVlanGroupsListQueryParameters) GetRack() int32 {
	return m.Rack
}

// SetRack sets the Rack property
func (m *IpamVlanGroupsListQueryParameters) SetRack(val int32) {
	m.Rack = val
}

// GetRegion returns the Region property
func (m IpamVlanGroupsListQueryParameters) GetRegion() int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *IpamVlanGroupsListQueryParameters) SetRegion(val int32) {
	m.Region = val
}

// GetScopeId returns the ScopeId property
func (m IpamVlanGroupsListQueryParameters) GetScopeId() []int32 {
	return m.ScopeId
}

// SetScopeId sets the ScopeId property
func (m *IpamVlanGroupsListQueryParameters) SetScopeId(val []int32) {
	m.ScopeId = val
}

// GetScopeIdGt returns the ScopeIdGt property
func (m IpamVlanGroupsListQueryParameters) GetScopeIdGt() []int32 {
	return m.ScopeIdGt
}

// SetScopeIdGt sets the ScopeIdGt property
func (m *IpamVlanGroupsListQueryParameters) SetScopeIdGt(val []int32) {
	m.ScopeIdGt = val
}

// GetScopeIdGte returns the ScopeIdGte property
func (m IpamVlanGroupsListQueryParameters) GetScopeIdGte() []int32 {
	return m.ScopeIdGte
}

// SetScopeIdGte sets the ScopeIdGte property
func (m *IpamVlanGroupsListQueryParameters) SetScopeIdGte(val []int32) {
	m.ScopeIdGte = val
}

// GetScopeIdLt returns the ScopeIdLt property
func (m IpamVlanGroupsListQueryParameters) GetScopeIdLt() []int32 {
	return m.ScopeIdLt
}

// SetScopeIdLt sets the ScopeIdLt property
func (m *IpamVlanGroupsListQueryParameters) SetScopeIdLt(val []int32) {
	m.ScopeIdLt = val
}

// GetScopeIdLte returns the ScopeIdLte property
func (m IpamVlanGroupsListQueryParameters) GetScopeIdLte() []int32 {
	return m.ScopeIdLte
}

// SetScopeIdLte sets the ScopeIdLte property
func (m *IpamVlanGroupsListQueryParameters) SetScopeIdLte(val []int32) {
	m.ScopeIdLte = val
}

// GetScopeIdN returns the ScopeIdN property
func (m IpamVlanGroupsListQueryParameters) GetScopeIdN() []int32 {
	return m.ScopeIdN
}

// SetScopeIdN sets the ScopeIdN property
func (m *IpamVlanGroupsListQueryParameters) SetScopeIdN(val []int32) {
	m.ScopeIdN = val
}

// GetScopeType returns the ScopeType property
func (m IpamVlanGroupsListQueryParameters) GetScopeType() string {
	return m.ScopeType
}

// SetScopeType sets the ScopeType property
func (m *IpamVlanGroupsListQueryParameters) SetScopeType(val string) {
	m.ScopeType = val
}

// GetScopeTypeN returns the ScopeTypeN property
func (m IpamVlanGroupsListQueryParameters) GetScopeTypeN() string {
	return m.ScopeTypeN
}

// SetScopeTypeN sets the ScopeTypeN property
func (m *IpamVlanGroupsListQueryParameters) SetScopeTypeN(val string) {
	m.ScopeTypeN = val
}

// GetSite returns the Site property
func (m IpamVlanGroupsListQueryParameters) GetSite() int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *IpamVlanGroupsListQueryParameters) SetSite(val int32) {
	m.Site = val
}

// GetSitegroup returns the Sitegroup property
func (m IpamVlanGroupsListQueryParameters) GetSitegroup() float32 {
	return m.Sitegroup
}

// SetSitegroup sets the Sitegroup property
func (m *IpamVlanGroupsListQueryParameters) SetSitegroup(val float32) {
	m.Sitegroup = val
}

// GetSlug returns the Slug property
func (m IpamVlanGroupsListQueryParameters) GetSlug() []string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *IpamVlanGroupsListQueryParameters) SetSlug(val []string) {
	m.Slug = val
}

// GetSlugEmpty returns the SlugEmpty property
func (m IpamVlanGroupsListQueryParameters) GetSlugEmpty() []string {
	return m.SlugEmpty
}

// SetSlugEmpty sets the SlugEmpty property
func (m *IpamVlanGroupsListQueryParameters) SetSlugEmpty(val []string) {
	m.SlugEmpty = val
}

// GetSlugIc returns the SlugIc property
func (m IpamVlanGroupsListQueryParameters) GetSlugIc() []string {
	return m.SlugIc
}

// SetSlugIc sets the SlugIc property
func (m *IpamVlanGroupsListQueryParameters) SetSlugIc(val []string) {
	m.SlugIc = val
}

// GetSlugIe returns the SlugIe property
func (m IpamVlanGroupsListQueryParameters) GetSlugIe() []string {
	return m.SlugIe
}

// SetSlugIe sets the SlugIe property
func (m *IpamVlanGroupsListQueryParameters) SetSlugIe(val []string) {
	m.SlugIe = val
}

// GetSlugIew returns the SlugIew property
func (m IpamVlanGroupsListQueryParameters) GetSlugIew() []string {
	return m.SlugIew
}

// SetSlugIew sets the SlugIew property
func (m *IpamVlanGroupsListQueryParameters) SetSlugIew(val []string) {
	m.SlugIew = val
}

// GetSlugIsw returns the SlugIsw property
func (m IpamVlanGroupsListQueryParameters) GetSlugIsw() []string {
	return m.SlugIsw
}

// SetSlugIsw sets the SlugIsw property
func (m *IpamVlanGroupsListQueryParameters) SetSlugIsw(val []string) {
	m.SlugIsw = val
}

// GetSlugN returns the SlugN property
func (m IpamVlanGroupsListQueryParameters) GetSlugN() []string {
	return m.SlugN
}

// SetSlugN sets the SlugN property
func (m *IpamVlanGroupsListQueryParameters) SetSlugN(val []string) {
	m.SlugN = val
}

// GetSlugNic returns the SlugNic property
func (m IpamVlanGroupsListQueryParameters) GetSlugNic() []string {
	return m.SlugNic
}

// SetSlugNic sets the SlugNic property
func (m *IpamVlanGroupsListQueryParameters) SetSlugNic(val []string) {
	m.SlugNic = val
}

// GetSlugNie returns the SlugNie property
func (m IpamVlanGroupsListQueryParameters) GetSlugNie() []string {
	return m.SlugNie
}

// SetSlugNie sets the SlugNie property
func (m *IpamVlanGroupsListQueryParameters) SetSlugNie(val []string) {
	m.SlugNie = val
}

// GetSlugNiew returns the SlugNiew property
func (m IpamVlanGroupsListQueryParameters) GetSlugNiew() []string {
	return m.SlugNiew
}

// SetSlugNiew sets the SlugNiew property
func (m *IpamVlanGroupsListQueryParameters) SetSlugNiew(val []string) {
	m.SlugNiew = val
}

// GetSlugNisw returns the SlugNisw property
func (m IpamVlanGroupsListQueryParameters) GetSlugNisw() []string {
	return m.SlugNisw
}

// SetSlugNisw sets the SlugNisw property
func (m *IpamVlanGroupsListQueryParameters) SetSlugNisw(val []string) {
	m.SlugNisw = val
}

// GetTag returns the Tag property
func (m IpamVlanGroupsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamVlanGroupsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamVlanGroupsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamVlanGroupsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamVlanGroupsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamVlanGroupsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
