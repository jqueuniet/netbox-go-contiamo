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

// WirelessWirelessLanGroupsListQueryParameters is an object.
type WirelessWirelessLanGroupsListQueryParameters struct {
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
	// Parent:
	Parent []string `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// ParentN:
	ParentN []string `json:"parent__n,omitempty" mapstructure:"parent__n,omitempty"`
	// ParentId:
	ParentId []*int32 `json:"parent_id,omitempty" mapstructure:"parent_id,omitempty"`
	// ParentIdN:
	ParentIdN []*int32 `json:"parent_id__n,omitempty" mapstructure:"parent_id__n,omitempty"`
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
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m WirelessWirelessLanGroupsListQueryParameters) Validate() error {
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
		"parent": validation.Validate(
			m.Parent,
		),
		"parentN": validation.Validate(
			m.ParentN,
		),
		"parentId": validation.Validate(
			m.ParentId,
		),
		"parentIdN": validation.Validate(
			m.ParentIdN,
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

// GetCreated returns the Created property
func (m WirelessWirelessLanGroupsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m WirelessWirelessLanGroupsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m WirelessWirelessLanGroupsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m WirelessWirelessLanGroupsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m WirelessWirelessLanGroupsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m WirelessWirelessLanGroupsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m WirelessWirelessLanGroupsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m WirelessWirelessLanGroupsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m WirelessWirelessLanGroupsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m WirelessWirelessLanGroupsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m WirelessWirelessLanGroupsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m WirelessWirelessLanGroupsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m WirelessWirelessLanGroupsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m WirelessWirelessLanGroupsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m WirelessWirelessLanGroupsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m WirelessWirelessLanGroupsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m WirelessWirelessLanGroupsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m WirelessWirelessLanGroupsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m WirelessWirelessLanGroupsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m WirelessWirelessLanGroupsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m WirelessWirelessLanGroupsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m WirelessWirelessLanGroupsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetParent returns the Parent property
func (m WirelessWirelessLanGroupsListQueryParameters) GetParent() []string {
	return m.Parent
}

// SetParent sets the Parent property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetParent(val []string) {
	m.Parent = val
}

// GetParentN returns the ParentN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetParentN() []string {
	return m.ParentN
}

// SetParentN sets the ParentN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetParentN(val []string) {
	m.ParentN = val
}

// GetParentId returns the ParentId property
func (m WirelessWirelessLanGroupsListQueryParameters) GetParentId() []*int32 {
	return m.ParentId
}

// SetParentId sets the ParentId property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetParentId(val []*int32) {
	m.ParentId = val
}

// GetParentIdN returns the ParentIdN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetParentIdN() []*int32 {
	return m.ParentIdN
}

// SetParentIdN sets the ParentIdN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetParentIdN(val []*int32) {
	m.ParentIdN = val
}

// GetQ returns the Q property
func (m WirelessWirelessLanGroupsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSlug returns the Slug property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlug() []string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlug(val []string) {
	m.Slug = val
}

// GetSlugEmpty returns the SlugEmpty property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugEmpty() []string {
	return m.SlugEmpty
}

// SetSlugEmpty sets the SlugEmpty property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugEmpty(val []string) {
	m.SlugEmpty = val
}

// GetSlugIc returns the SlugIc property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugIc() []string {
	return m.SlugIc
}

// SetSlugIc sets the SlugIc property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugIc(val []string) {
	m.SlugIc = val
}

// GetSlugIe returns the SlugIe property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugIe() []string {
	return m.SlugIe
}

// SetSlugIe sets the SlugIe property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugIe(val []string) {
	m.SlugIe = val
}

// GetSlugIew returns the SlugIew property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugIew() []string {
	return m.SlugIew
}

// SetSlugIew sets the SlugIew property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugIew(val []string) {
	m.SlugIew = val
}

// GetSlugIsw returns the SlugIsw property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugIsw() []string {
	return m.SlugIsw
}

// SetSlugIsw sets the SlugIsw property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugIsw(val []string) {
	m.SlugIsw = val
}

// GetSlugN returns the SlugN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugN() []string {
	return m.SlugN
}

// SetSlugN sets the SlugN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugN(val []string) {
	m.SlugN = val
}

// GetSlugNic returns the SlugNic property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugNic() []string {
	return m.SlugNic
}

// SetSlugNic sets the SlugNic property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugNic(val []string) {
	m.SlugNic = val
}

// GetSlugNie returns the SlugNie property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugNie() []string {
	return m.SlugNie
}

// SetSlugNie sets the SlugNie property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugNie(val []string) {
	m.SlugNie = val
}

// GetSlugNiew returns the SlugNiew property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugNiew() []string {
	return m.SlugNiew
}

// SetSlugNiew sets the SlugNiew property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugNiew(val []string) {
	m.SlugNiew = val
}

// GetSlugNisw returns the SlugNisw property
func (m WirelessWirelessLanGroupsListQueryParameters) GetSlugNisw() []string {
	return m.SlugNisw
}

// SetSlugNisw sets the SlugNisw property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetSlugNisw(val []string) {
	m.SlugNisw = val
}

// GetTag returns the Tag property
func (m WirelessWirelessLanGroupsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m WirelessWirelessLanGroupsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m WirelessWirelessLanGroupsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *WirelessWirelessLanGroupsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
