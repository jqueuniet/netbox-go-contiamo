// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ExtrasSavedFiltersListQueryParameters is an object.
type ExtrasSavedFiltersListQueryParameters struct {
	// ContentTypeId:
	ContentTypeId []int32 `json:"content_type_id,omitempty" mapstructure:"content_type_id,omitempty"`
	// ContentTypeIdGt:
	ContentTypeIdGt []int32 `json:"content_type_id__gt,omitempty" mapstructure:"content_type_id__gt,omitempty"`
	// ContentTypeIdGte:
	ContentTypeIdGte []int32 `json:"content_type_id__gte,omitempty" mapstructure:"content_type_id__gte,omitempty"`
	// ContentTypeIdLt:
	ContentTypeIdLt []int32 `json:"content_type_id__lt,omitempty" mapstructure:"content_type_id__lt,omitempty"`
	// ContentTypeIdLte:
	ContentTypeIdLte []int32 `json:"content_type_id__lte,omitempty" mapstructure:"content_type_id__lte,omitempty"`
	// ContentTypeIdN:
	ContentTypeIdN []int32 `json:"content_type_id__n,omitempty" mapstructure:"content_type_id__n,omitempty"`
	// ContentTypes:
	ContentTypes string `json:"content_types,omitempty" mapstructure:"content_types,omitempty"`
	// ContentTypesIc:
	ContentTypesIc string `json:"content_types__ic,omitempty" mapstructure:"content_types__ic,omitempty"`
	// ContentTypesIe:
	ContentTypesIe string `json:"content_types__ie,omitempty" mapstructure:"content_types__ie,omitempty"`
	// ContentTypesIew:
	ContentTypesIew string `json:"content_types__iew,omitempty" mapstructure:"content_types__iew,omitempty"`
	// ContentTypesIsw:
	ContentTypesIsw string `json:"content_types__isw,omitempty" mapstructure:"content_types__isw,omitempty"`
	// ContentTypesN:
	ContentTypesN string `json:"content_types__n,omitempty" mapstructure:"content_types__n,omitempty"`
	// ContentTypesNic:
	ContentTypesNic string `json:"content_types__nic,omitempty" mapstructure:"content_types__nic,omitempty"`
	// ContentTypesNie:
	ContentTypesNie string `json:"content_types__nie,omitempty" mapstructure:"content_types__nie,omitempty"`
	// ContentTypesNiew:
	ContentTypesNiew string `json:"content_types__niew,omitempty" mapstructure:"content_types__niew,omitempty"`
	// ContentTypesNisw:
	ContentTypesNisw string `json:"content_types__nisw,omitempty" mapstructure:"content_types__nisw,omitempty"`
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
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
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
	// Shared:
	Shared bool `json:"shared,omitempty" mapstructure:"shared,omitempty"`
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
	// Usable:
	Usable bool `json:"usable,omitempty" mapstructure:"usable,omitempty"`
	// User: User (name)
	User []string `json:"user,omitempty" mapstructure:"user,omitempty"`
	// UserN: User (name)
	UserN []string `json:"user__n,omitempty" mapstructure:"user__n,omitempty"`
	// UserId: User (ID)
	UserId []*int32 `json:"user_id,omitempty" mapstructure:"user_id,omitempty"`
	// UserIdN: User (ID)
	UserIdN []*int32 `json:"user_id__n,omitempty" mapstructure:"user_id__n,omitempty"`
	// Weight:
	Weight []int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightGt:
	WeightGt []int32 `json:"weight__gt,omitempty" mapstructure:"weight__gt,omitempty"`
	// WeightGte:
	WeightGte []int32 `json:"weight__gte,omitempty" mapstructure:"weight__gte,omitempty"`
	// WeightLt:
	WeightLt []int32 `json:"weight__lt,omitempty" mapstructure:"weight__lt,omitempty"`
	// WeightLte:
	WeightLte []int32 `json:"weight__lte,omitempty" mapstructure:"weight__lte,omitempty"`
	// WeightN:
	WeightN []int32 `json:"weight__n,omitempty" mapstructure:"weight__n,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasSavedFiltersListQueryParameters) Validate() error {
	return validation.Errors{
		"contentTypeId": validation.Validate(
			m.ContentTypeId,
		),
		"contentTypeIdGt": validation.Validate(
			m.ContentTypeIdGt,
		),
		"contentTypeIdGte": validation.Validate(
			m.ContentTypeIdGte,
		),
		"contentTypeIdLt": validation.Validate(
			m.ContentTypeIdLt,
		),
		"contentTypeIdLte": validation.Validate(
			m.ContentTypeIdLte,
		),
		"contentTypeIdN": validation.Validate(
			m.ContentTypeIdN,
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
		"weight": validation.Validate(
			m.Weight,
		),
		"weightGt": validation.Validate(
			m.WeightGt,
		),
		"weightGte": validation.Validate(
			m.WeightGte,
		),
		"weightLt": validation.Validate(
			m.WeightLt,
		),
		"weightLte": validation.Validate(
			m.WeightLte,
		),
		"weightN": validation.Validate(
			m.WeightN,
		),
	}.Filter()
}

// GetContentTypeId returns the ContentTypeId property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypeId() []int32 {
	return m.ContentTypeId
}

// SetContentTypeId sets the ContentTypeId property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypeId(val []int32) {
	m.ContentTypeId = val
}

// GetContentTypeIdGt returns the ContentTypeIdGt property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypeIdGt() []int32 {
	return m.ContentTypeIdGt
}

// SetContentTypeIdGt sets the ContentTypeIdGt property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypeIdGt(val []int32) {
	m.ContentTypeIdGt = val
}

// GetContentTypeIdGte returns the ContentTypeIdGte property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypeIdGte() []int32 {
	return m.ContentTypeIdGte
}

// SetContentTypeIdGte sets the ContentTypeIdGte property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypeIdGte(val []int32) {
	m.ContentTypeIdGte = val
}

// GetContentTypeIdLt returns the ContentTypeIdLt property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypeIdLt() []int32 {
	return m.ContentTypeIdLt
}

// SetContentTypeIdLt sets the ContentTypeIdLt property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypeIdLt(val []int32) {
	m.ContentTypeIdLt = val
}

// GetContentTypeIdLte returns the ContentTypeIdLte property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypeIdLte() []int32 {
	return m.ContentTypeIdLte
}

// SetContentTypeIdLte sets the ContentTypeIdLte property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypeIdLte(val []int32) {
	m.ContentTypeIdLte = val
}

// GetContentTypeIdN returns the ContentTypeIdN property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypeIdN() []int32 {
	return m.ContentTypeIdN
}

// SetContentTypeIdN sets the ContentTypeIdN property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypeIdN(val []int32) {
	m.ContentTypeIdN = val
}

// GetContentTypes returns the ContentTypes property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypes() string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypes(val string) {
	m.ContentTypes = val
}

// GetContentTypesIc returns the ContentTypesIc property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesIc() string {
	return m.ContentTypesIc
}

// SetContentTypesIc sets the ContentTypesIc property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesIc(val string) {
	m.ContentTypesIc = val
}

// GetContentTypesIe returns the ContentTypesIe property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesIe() string {
	return m.ContentTypesIe
}

// SetContentTypesIe sets the ContentTypesIe property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesIe(val string) {
	m.ContentTypesIe = val
}

// GetContentTypesIew returns the ContentTypesIew property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesIew() string {
	return m.ContentTypesIew
}

// SetContentTypesIew sets the ContentTypesIew property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesIew(val string) {
	m.ContentTypesIew = val
}

// GetContentTypesIsw returns the ContentTypesIsw property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesIsw() string {
	return m.ContentTypesIsw
}

// SetContentTypesIsw sets the ContentTypesIsw property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesIsw(val string) {
	m.ContentTypesIsw = val
}

// GetContentTypesN returns the ContentTypesN property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesN() string {
	return m.ContentTypesN
}

// SetContentTypesN sets the ContentTypesN property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesN(val string) {
	m.ContentTypesN = val
}

// GetContentTypesNic returns the ContentTypesNic property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesNic() string {
	return m.ContentTypesNic
}

// SetContentTypesNic sets the ContentTypesNic property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesNic(val string) {
	m.ContentTypesNic = val
}

// GetContentTypesNie returns the ContentTypesNie property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesNie() string {
	return m.ContentTypesNie
}

// SetContentTypesNie sets the ContentTypesNie property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesNie(val string) {
	m.ContentTypesNie = val
}

// GetContentTypesNiew returns the ContentTypesNiew property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesNiew() string {
	return m.ContentTypesNiew
}

// SetContentTypesNiew sets the ContentTypesNiew property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesNiew(val string) {
	m.ContentTypesNiew = val
}

// GetContentTypesNisw returns the ContentTypesNisw property
func (m ExtrasSavedFiltersListQueryParameters) GetContentTypesNisw() string {
	return m.ContentTypesNisw
}

// SetContentTypesNisw sets the ContentTypesNisw property
func (m *ExtrasSavedFiltersListQueryParameters) SetContentTypesNisw(val string) {
	m.ContentTypesNisw = val
}

// GetDescription returns the Description property
func (m ExtrasSavedFiltersListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m ExtrasSavedFiltersListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *ExtrasSavedFiltersListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetEnabled returns the Enabled property
func (m ExtrasSavedFiltersListQueryParameters) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *ExtrasSavedFiltersListQueryParameters) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m ExtrasSavedFiltersListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasSavedFiltersListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasSavedFiltersListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasSavedFiltersListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasSavedFiltersListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasSavedFiltersListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasSavedFiltersListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasSavedFiltersListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasSavedFiltersListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasSavedFiltersListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasSavedFiltersListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasSavedFiltersListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m ExtrasSavedFiltersListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasSavedFiltersListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m ExtrasSavedFiltersListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasSavedFiltersListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasSavedFiltersListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasSavedFiltersListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasSavedFiltersListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasSavedFiltersListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasSavedFiltersListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasSavedFiltersListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasSavedFiltersListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasSavedFiltersListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasSavedFiltersListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasSavedFiltersListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasSavedFiltersListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m ExtrasSavedFiltersListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasSavedFiltersListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasSavedFiltersListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasSavedFiltersListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasSavedFiltersListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasSavedFiltersListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetShared returns the Shared property
func (m ExtrasSavedFiltersListQueryParameters) GetShared() bool {
	return m.Shared
}

// SetShared sets the Shared property
func (m *ExtrasSavedFiltersListQueryParameters) SetShared(val bool) {
	m.Shared = val
}

// GetSlug returns the Slug property
func (m ExtrasSavedFiltersListQueryParameters) GetSlug() []string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlug(val []string) {
	m.Slug = val
}

// GetSlugEmpty returns the SlugEmpty property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugEmpty() []string {
	return m.SlugEmpty
}

// SetSlugEmpty sets the SlugEmpty property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugEmpty(val []string) {
	m.SlugEmpty = val
}

// GetSlugIc returns the SlugIc property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugIc() []string {
	return m.SlugIc
}

// SetSlugIc sets the SlugIc property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugIc(val []string) {
	m.SlugIc = val
}

// GetSlugIe returns the SlugIe property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugIe() []string {
	return m.SlugIe
}

// SetSlugIe sets the SlugIe property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugIe(val []string) {
	m.SlugIe = val
}

// GetSlugIew returns the SlugIew property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugIew() []string {
	return m.SlugIew
}

// SetSlugIew sets the SlugIew property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugIew(val []string) {
	m.SlugIew = val
}

// GetSlugIsw returns the SlugIsw property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugIsw() []string {
	return m.SlugIsw
}

// SetSlugIsw sets the SlugIsw property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugIsw(val []string) {
	m.SlugIsw = val
}

// GetSlugN returns the SlugN property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugN() []string {
	return m.SlugN
}

// SetSlugN sets the SlugN property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugN(val []string) {
	m.SlugN = val
}

// GetSlugNic returns the SlugNic property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugNic() []string {
	return m.SlugNic
}

// SetSlugNic sets the SlugNic property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugNic(val []string) {
	m.SlugNic = val
}

// GetSlugNie returns the SlugNie property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugNie() []string {
	return m.SlugNie
}

// SetSlugNie sets the SlugNie property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugNie(val []string) {
	m.SlugNie = val
}

// GetSlugNiew returns the SlugNiew property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugNiew() []string {
	return m.SlugNiew
}

// SetSlugNiew sets the SlugNiew property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugNiew(val []string) {
	m.SlugNiew = val
}

// GetSlugNisw returns the SlugNisw property
func (m ExtrasSavedFiltersListQueryParameters) GetSlugNisw() []string {
	return m.SlugNisw
}

// SetSlugNisw sets the SlugNisw property
func (m *ExtrasSavedFiltersListQueryParameters) SetSlugNisw(val []string) {
	m.SlugNisw = val
}

// GetUsable returns the Usable property
func (m ExtrasSavedFiltersListQueryParameters) GetUsable() bool {
	return m.Usable
}

// SetUsable sets the Usable property
func (m *ExtrasSavedFiltersListQueryParameters) SetUsable(val bool) {
	m.Usable = val
}

// GetUser returns the User property
func (m ExtrasSavedFiltersListQueryParameters) GetUser() []string {
	return m.User
}

// SetUser sets the User property
func (m *ExtrasSavedFiltersListQueryParameters) SetUser(val []string) {
	m.User = val
}

// GetUserN returns the UserN property
func (m ExtrasSavedFiltersListQueryParameters) GetUserN() []string {
	return m.UserN
}

// SetUserN sets the UserN property
func (m *ExtrasSavedFiltersListQueryParameters) SetUserN(val []string) {
	m.UserN = val
}

// GetUserId returns the UserId property
func (m ExtrasSavedFiltersListQueryParameters) GetUserId() []*int32 {
	return m.UserId
}

// SetUserId sets the UserId property
func (m *ExtrasSavedFiltersListQueryParameters) SetUserId(val []*int32) {
	m.UserId = val
}

// GetUserIdN returns the UserIdN property
func (m ExtrasSavedFiltersListQueryParameters) GetUserIdN() []*int32 {
	return m.UserIdN
}

// SetUserIdN sets the UserIdN property
func (m *ExtrasSavedFiltersListQueryParameters) SetUserIdN(val []*int32) {
	m.UserIdN = val
}

// GetWeight returns the Weight property
func (m ExtrasSavedFiltersListQueryParameters) GetWeight() []int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *ExtrasSavedFiltersListQueryParameters) SetWeight(val []int32) {
	m.Weight = val
}

// GetWeightGt returns the WeightGt property
func (m ExtrasSavedFiltersListQueryParameters) GetWeightGt() []int32 {
	return m.WeightGt
}

// SetWeightGt sets the WeightGt property
func (m *ExtrasSavedFiltersListQueryParameters) SetWeightGt(val []int32) {
	m.WeightGt = val
}

// GetWeightGte returns the WeightGte property
func (m ExtrasSavedFiltersListQueryParameters) GetWeightGte() []int32 {
	return m.WeightGte
}

// SetWeightGte sets the WeightGte property
func (m *ExtrasSavedFiltersListQueryParameters) SetWeightGte(val []int32) {
	m.WeightGte = val
}

// GetWeightLt returns the WeightLt property
func (m ExtrasSavedFiltersListQueryParameters) GetWeightLt() []int32 {
	return m.WeightLt
}

// SetWeightLt sets the WeightLt property
func (m *ExtrasSavedFiltersListQueryParameters) SetWeightLt(val []int32) {
	m.WeightLt = val
}

// GetWeightLte returns the WeightLte property
func (m ExtrasSavedFiltersListQueryParameters) GetWeightLte() []int32 {
	return m.WeightLte
}

// SetWeightLte sets the WeightLte property
func (m *ExtrasSavedFiltersListQueryParameters) SetWeightLte(val []int32) {
	m.WeightLte = val
}

// GetWeightN returns the WeightN property
func (m ExtrasSavedFiltersListQueryParameters) GetWeightN() []int32 {
	return m.WeightN
}

// SetWeightN sets the WeightN property
func (m *ExtrasSavedFiltersListQueryParameters) SetWeightN(val []int32) {
	m.WeightN = val
}
