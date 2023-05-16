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

// ExtrasCustomFieldsListQueryParameters is an object.
type ExtrasCustomFieldsListQueryParameters struct {
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
	// FilterLogic: Loose matches any instance of a given string; exact matches the entire field.
	FilterLogic string `json:"filter_logic,omitempty" mapstructure:"filter_logic,omitempty"`
	// FilterLogicN: Loose matches any instance of a given string; exact matches the entire field.
	FilterLogicN string `json:"filter_logic__n,omitempty" mapstructure:"filter_logic__n,omitempty"`
	// GroupName:
	GroupName []string `json:"group_name,omitempty" mapstructure:"group_name,omitempty"`
	// GroupNameEmpty:
	GroupNameEmpty []string `json:"group_name__empty,omitempty" mapstructure:"group_name__empty,omitempty"`
	// GroupNameIc:
	GroupNameIc []string `json:"group_name__ic,omitempty" mapstructure:"group_name__ic,omitempty"`
	// GroupNameIe:
	GroupNameIe []string `json:"group_name__ie,omitempty" mapstructure:"group_name__ie,omitempty"`
	// GroupNameIew:
	GroupNameIew []string `json:"group_name__iew,omitempty" mapstructure:"group_name__iew,omitempty"`
	// GroupNameIsw:
	GroupNameIsw []string `json:"group_name__isw,omitempty" mapstructure:"group_name__isw,omitempty"`
	// GroupNameN:
	GroupNameN []string `json:"group_name__n,omitempty" mapstructure:"group_name__n,omitempty"`
	// GroupNameNic:
	GroupNameNic []string `json:"group_name__nic,omitempty" mapstructure:"group_name__nic,omitempty"`
	// GroupNameNie:
	GroupNameNie []string `json:"group_name__nie,omitempty" mapstructure:"group_name__nie,omitempty"`
	// GroupNameNiew:
	GroupNameNiew []string `json:"group_name__niew,omitempty" mapstructure:"group_name__niew,omitempty"`
	// GroupNameNisw:
	GroupNameNisw []string `json:"group_name__nisw,omitempty" mapstructure:"group_name__nisw,omitempty"`
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
	// IsCloneable:
	IsCloneable bool `json:"is_cloneable,omitempty" mapstructure:"is_cloneable,omitempty"`
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
	// Required:
	Required bool `json:"required,omitempty" mapstructure:"required,omitempty"`
	// SearchWeight:
	SearchWeight []int32 `json:"search_weight,omitempty" mapstructure:"search_weight,omitempty"`
	// SearchWeightGt:
	SearchWeightGt []int32 `json:"search_weight__gt,omitempty" mapstructure:"search_weight__gt,omitempty"`
	// SearchWeightGte:
	SearchWeightGte []int32 `json:"search_weight__gte,omitempty" mapstructure:"search_weight__gte,omitempty"`
	// SearchWeightLt:
	SearchWeightLt []int32 `json:"search_weight__lt,omitempty" mapstructure:"search_weight__lt,omitempty"`
	// SearchWeightLte:
	SearchWeightLte []int32 `json:"search_weight__lte,omitempty" mapstructure:"search_weight__lte,omitempty"`
	// SearchWeightN:
	SearchWeightN []int32 `json:"search_weight__n,omitempty" mapstructure:"search_weight__n,omitempty"`
	// Type: The type of data this custom field holds
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: The type of data this custom field holds
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UiVisibility: Specifies the visibility of custom field in the UI
	UiVisibility string `json:"ui_visibility,omitempty" mapstructure:"ui_visibility,omitempty"`
	// UiVisibilityN: Specifies the visibility of custom field in the UI
	UiVisibilityN string `json:"ui_visibility__n,omitempty" mapstructure:"ui_visibility__n,omitempty"`
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
func (m ExtrasCustomFieldsListQueryParameters) Validate() error {
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
		"groupName": validation.Validate(
			m.GroupName,
		),
		"groupNameEmpty": validation.Validate(
			m.GroupNameEmpty,
		),
		"groupNameIc": validation.Validate(
			m.GroupNameIc,
		),
		"groupNameIe": validation.Validate(
			m.GroupNameIe,
		),
		"groupNameIew": validation.Validate(
			m.GroupNameIew,
		),
		"groupNameIsw": validation.Validate(
			m.GroupNameIsw,
		),
		"groupNameN": validation.Validate(
			m.GroupNameN,
		),
		"groupNameNic": validation.Validate(
			m.GroupNameNic,
		),
		"groupNameNie": validation.Validate(
			m.GroupNameNie,
		),
		"groupNameNiew": validation.Validate(
			m.GroupNameNiew,
		),
		"groupNameNisw": validation.Validate(
			m.GroupNameNisw,
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
		"searchWeight": validation.Validate(
			m.SearchWeight,
		),
		"searchWeightGt": validation.Validate(
			m.SearchWeightGt,
		),
		"searchWeightGte": validation.Validate(
			m.SearchWeightGte,
		),
		"searchWeightLt": validation.Validate(
			m.SearchWeightLt,
		),
		"searchWeightLte": validation.Validate(
			m.SearchWeightLte,
		),
		"searchWeightN": validation.Validate(
			m.SearchWeightN,
		),
		"type": validation.Validate(
			m.Type,
		),
		"typeN": validation.Validate(
			m.TypeN,
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
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypeId() []int32 {
	return m.ContentTypeId
}

// SetContentTypeId sets the ContentTypeId property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypeId(val []int32) {
	m.ContentTypeId = val
}

// GetContentTypeIdGt returns the ContentTypeIdGt property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypeIdGt() []int32 {
	return m.ContentTypeIdGt
}

// SetContentTypeIdGt sets the ContentTypeIdGt property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypeIdGt(val []int32) {
	m.ContentTypeIdGt = val
}

// GetContentTypeIdGte returns the ContentTypeIdGte property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypeIdGte() []int32 {
	return m.ContentTypeIdGte
}

// SetContentTypeIdGte sets the ContentTypeIdGte property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypeIdGte(val []int32) {
	m.ContentTypeIdGte = val
}

// GetContentTypeIdLt returns the ContentTypeIdLt property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypeIdLt() []int32 {
	return m.ContentTypeIdLt
}

// SetContentTypeIdLt sets the ContentTypeIdLt property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypeIdLt(val []int32) {
	m.ContentTypeIdLt = val
}

// GetContentTypeIdLte returns the ContentTypeIdLte property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypeIdLte() []int32 {
	return m.ContentTypeIdLte
}

// SetContentTypeIdLte sets the ContentTypeIdLte property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypeIdLte(val []int32) {
	m.ContentTypeIdLte = val
}

// GetContentTypeIdN returns the ContentTypeIdN property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypeIdN() []int32 {
	return m.ContentTypeIdN
}

// SetContentTypeIdN sets the ContentTypeIdN property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypeIdN(val []int32) {
	m.ContentTypeIdN = val
}

// GetContentTypes returns the ContentTypes property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypes() string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypes(val string) {
	m.ContentTypes = val
}

// GetContentTypesIc returns the ContentTypesIc property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesIc() string {
	return m.ContentTypesIc
}

// SetContentTypesIc sets the ContentTypesIc property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesIc(val string) {
	m.ContentTypesIc = val
}

// GetContentTypesIe returns the ContentTypesIe property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesIe() string {
	return m.ContentTypesIe
}

// SetContentTypesIe sets the ContentTypesIe property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesIe(val string) {
	m.ContentTypesIe = val
}

// GetContentTypesIew returns the ContentTypesIew property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesIew() string {
	return m.ContentTypesIew
}

// SetContentTypesIew sets the ContentTypesIew property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesIew(val string) {
	m.ContentTypesIew = val
}

// GetContentTypesIsw returns the ContentTypesIsw property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesIsw() string {
	return m.ContentTypesIsw
}

// SetContentTypesIsw sets the ContentTypesIsw property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesIsw(val string) {
	m.ContentTypesIsw = val
}

// GetContentTypesN returns the ContentTypesN property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesN() string {
	return m.ContentTypesN
}

// SetContentTypesN sets the ContentTypesN property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesN(val string) {
	m.ContentTypesN = val
}

// GetContentTypesNic returns the ContentTypesNic property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesNic() string {
	return m.ContentTypesNic
}

// SetContentTypesNic sets the ContentTypesNic property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesNic(val string) {
	m.ContentTypesNic = val
}

// GetContentTypesNie returns the ContentTypesNie property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesNie() string {
	return m.ContentTypesNie
}

// SetContentTypesNie sets the ContentTypesNie property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesNie(val string) {
	m.ContentTypesNie = val
}

// GetContentTypesNiew returns the ContentTypesNiew property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesNiew() string {
	return m.ContentTypesNiew
}

// SetContentTypesNiew sets the ContentTypesNiew property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesNiew(val string) {
	m.ContentTypesNiew = val
}

// GetContentTypesNisw returns the ContentTypesNisw property
func (m ExtrasCustomFieldsListQueryParameters) GetContentTypesNisw() string {
	return m.ContentTypesNisw
}

// SetContentTypesNisw sets the ContentTypesNisw property
func (m *ExtrasCustomFieldsListQueryParameters) SetContentTypesNisw(val string) {
	m.ContentTypesNisw = val
}

// GetDescription returns the Description property
func (m ExtrasCustomFieldsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m ExtrasCustomFieldsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *ExtrasCustomFieldsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetFilterLogic returns the FilterLogic property
func (m ExtrasCustomFieldsListQueryParameters) GetFilterLogic() string {
	return m.FilterLogic
}

// SetFilterLogic sets the FilterLogic property
func (m *ExtrasCustomFieldsListQueryParameters) SetFilterLogic(val string) {
	m.FilterLogic = val
}

// GetFilterLogicN returns the FilterLogicN property
func (m ExtrasCustomFieldsListQueryParameters) GetFilterLogicN() string {
	return m.FilterLogicN
}

// SetFilterLogicN sets the FilterLogicN property
func (m *ExtrasCustomFieldsListQueryParameters) SetFilterLogicN(val string) {
	m.FilterLogicN = val
}

// GetGroupName returns the GroupName property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupName() []string {
	return m.GroupName
}

// SetGroupName sets the GroupName property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupName(val []string) {
	m.GroupName = val
}

// GetGroupNameEmpty returns the GroupNameEmpty property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameEmpty() []string {
	return m.GroupNameEmpty
}

// SetGroupNameEmpty sets the GroupNameEmpty property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameEmpty(val []string) {
	m.GroupNameEmpty = val
}

// GetGroupNameIc returns the GroupNameIc property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameIc() []string {
	return m.GroupNameIc
}

// SetGroupNameIc sets the GroupNameIc property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameIc(val []string) {
	m.GroupNameIc = val
}

// GetGroupNameIe returns the GroupNameIe property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameIe() []string {
	return m.GroupNameIe
}

// SetGroupNameIe sets the GroupNameIe property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameIe(val []string) {
	m.GroupNameIe = val
}

// GetGroupNameIew returns the GroupNameIew property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameIew() []string {
	return m.GroupNameIew
}

// SetGroupNameIew sets the GroupNameIew property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameIew(val []string) {
	m.GroupNameIew = val
}

// GetGroupNameIsw returns the GroupNameIsw property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameIsw() []string {
	return m.GroupNameIsw
}

// SetGroupNameIsw sets the GroupNameIsw property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameIsw(val []string) {
	m.GroupNameIsw = val
}

// GetGroupNameN returns the GroupNameN property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameN() []string {
	return m.GroupNameN
}

// SetGroupNameN sets the GroupNameN property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameN(val []string) {
	m.GroupNameN = val
}

// GetGroupNameNic returns the GroupNameNic property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameNic() []string {
	return m.GroupNameNic
}

// SetGroupNameNic sets the GroupNameNic property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameNic(val []string) {
	m.GroupNameNic = val
}

// GetGroupNameNie returns the GroupNameNie property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameNie() []string {
	return m.GroupNameNie
}

// SetGroupNameNie sets the GroupNameNie property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameNie(val []string) {
	m.GroupNameNie = val
}

// GetGroupNameNiew returns the GroupNameNiew property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameNiew() []string {
	return m.GroupNameNiew
}

// SetGroupNameNiew sets the GroupNameNiew property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameNiew(val []string) {
	m.GroupNameNiew = val
}

// GetGroupNameNisw returns the GroupNameNisw property
func (m ExtrasCustomFieldsListQueryParameters) GetGroupNameNisw() []string {
	return m.GroupNameNisw
}

// SetGroupNameNisw sets the GroupNameNisw property
func (m *ExtrasCustomFieldsListQueryParameters) SetGroupNameNisw(val []string) {
	m.GroupNameNisw = val
}

// GetId returns the Id property
func (m ExtrasCustomFieldsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasCustomFieldsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasCustomFieldsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasCustomFieldsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasCustomFieldsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasCustomFieldsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasCustomFieldsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasCustomFieldsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasCustomFieldsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasCustomFieldsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasCustomFieldsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasCustomFieldsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetIsCloneable returns the IsCloneable property
func (m ExtrasCustomFieldsListQueryParameters) GetIsCloneable() bool {
	return m.IsCloneable
}

// SetIsCloneable sets the IsCloneable property
func (m *ExtrasCustomFieldsListQueryParameters) SetIsCloneable(val bool) {
	m.IsCloneable = val
}

// GetLimit returns the Limit property
func (m ExtrasCustomFieldsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasCustomFieldsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m ExtrasCustomFieldsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasCustomFieldsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasCustomFieldsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasCustomFieldsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasCustomFieldsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasCustomFieldsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasCustomFieldsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasCustomFieldsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasCustomFieldsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasCustomFieldsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasCustomFieldsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasCustomFieldsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasCustomFieldsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m ExtrasCustomFieldsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasCustomFieldsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasCustomFieldsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasCustomFieldsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasCustomFieldsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasCustomFieldsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRequired returns the Required property
func (m ExtrasCustomFieldsListQueryParameters) GetRequired() bool {
	return m.Required
}

// SetRequired sets the Required property
func (m *ExtrasCustomFieldsListQueryParameters) SetRequired(val bool) {
	m.Required = val
}

// GetSearchWeight returns the SearchWeight property
func (m ExtrasCustomFieldsListQueryParameters) GetSearchWeight() []int32 {
	return m.SearchWeight
}

// SetSearchWeight sets the SearchWeight property
func (m *ExtrasCustomFieldsListQueryParameters) SetSearchWeight(val []int32) {
	m.SearchWeight = val
}

// GetSearchWeightGt returns the SearchWeightGt property
func (m ExtrasCustomFieldsListQueryParameters) GetSearchWeightGt() []int32 {
	return m.SearchWeightGt
}

// SetSearchWeightGt sets the SearchWeightGt property
func (m *ExtrasCustomFieldsListQueryParameters) SetSearchWeightGt(val []int32) {
	m.SearchWeightGt = val
}

// GetSearchWeightGte returns the SearchWeightGte property
func (m ExtrasCustomFieldsListQueryParameters) GetSearchWeightGte() []int32 {
	return m.SearchWeightGte
}

// SetSearchWeightGte sets the SearchWeightGte property
func (m *ExtrasCustomFieldsListQueryParameters) SetSearchWeightGte(val []int32) {
	m.SearchWeightGte = val
}

// GetSearchWeightLt returns the SearchWeightLt property
func (m ExtrasCustomFieldsListQueryParameters) GetSearchWeightLt() []int32 {
	return m.SearchWeightLt
}

// SetSearchWeightLt sets the SearchWeightLt property
func (m *ExtrasCustomFieldsListQueryParameters) SetSearchWeightLt(val []int32) {
	m.SearchWeightLt = val
}

// GetSearchWeightLte returns the SearchWeightLte property
func (m ExtrasCustomFieldsListQueryParameters) GetSearchWeightLte() []int32 {
	return m.SearchWeightLte
}

// SetSearchWeightLte sets the SearchWeightLte property
func (m *ExtrasCustomFieldsListQueryParameters) SetSearchWeightLte(val []int32) {
	m.SearchWeightLte = val
}

// GetSearchWeightN returns the SearchWeightN property
func (m ExtrasCustomFieldsListQueryParameters) GetSearchWeightN() []int32 {
	return m.SearchWeightN
}

// SetSearchWeightN sets the SearchWeightN property
func (m *ExtrasCustomFieldsListQueryParameters) SetSearchWeightN(val []int32) {
	m.SearchWeightN = val
}

// GetType returns the Type property
func (m ExtrasCustomFieldsListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *ExtrasCustomFieldsListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m ExtrasCustomFieldsListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *ExtrasCustomFieldsListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUiVisibility returns the UiVisibility property
func (m ExtrasCustomFieldsListQueryParameters) GetUiVisibility() string {
	return m.UiVisibility
}

// SetUiVisibility sets the UiVisibility property
func (m *ExtrasCustomFieldsListQueryParameters) SetUiVisibility(val string) {
	m.UiVisibility = val
}

// GetUiVisibilityN returns the UiVisibilityN property
func (m ExtrasCustomFieldsListQueryParameters) GetUiVisibilityN() string {
	return m.UiVisibilityN
}

// SetUiVisibilityN sets the UiVisibilityN property
func (m *ExtrasCustomFieldsListQueryParameters) SetUiVisibilityN(val string) {
	m.UiVisibilityN = val
}

// GetWeight returns the Weight property
func (m ExtrasCustomFieldsListQueryParameters) GetWeight() []int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *ExtrasCustomFieldsListQueryParameters) SetWeight(val []int32) {
	m.Weight = val
}

// GetWeightGt returns the WeightGt property
func (m ExtrasCustomFieldsListQueryParameters) GetWeightGt() []int32 {
	return m.WeightGt
}

// SetWeightGt sets the WeightGt property
func (m *ExtrasCustomFieldsListQueryParameters) SetWeightGt(val []int32) {
	m.WeightGt = val
}

// GetWeightGte returns the WeightGte property
func (m ExtrasCustomFieldsListQueryParameters) GetWeightGte() []int32 {
	return m.WeightGte
}

// SetWeightGte sets the WeightGte property
func (m *ExtrasCustomFieldsListQueryParameters) SetWeightGte(val []int32) {
	m.WeightGte = val
}

// GetWeightLt returns the WeightLt property
func (m ExtrasCustomFieldsListQueryParameters) GetWeightLt() []int32 {
	return m.WeightLt
}

// SetWeightLt sets the WeightLt property
func (m *ExtrasCustomFieldsListQueryParameters) SetWeightLt(val []int32) {
	m.WeightLt = val
}

// GetWeightLte returns the WeightLte property
func (m ExtrasCustomFieldsListQueryParameters) GetWeightLte() []int32 {
	return m.WeightLte
}

// SetWeightLte sets the WeightLte property
func (m *ExtrasCustomFieldsListQueryParameters) SetWeightLte(val []int32) {
	m.WeightLte = val
}

// GetWeightN returns the WeightN property
func (m ExtrasCustomFieldsListQueryParameters) GetWeightN() []int32 {
	return m.WeightN
}

// SetWeightN sets the WeightN property
func (m *ExtrasCustomFieldsListQueryParameters) SetWeightN(val []int32) {
	m.WeightN = val
}
