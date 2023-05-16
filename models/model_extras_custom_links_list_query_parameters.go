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

// ExtrasCustomLinksListQueryParameters is an object.
type ExtrasCustomLinksListQueryParameters struct {
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
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
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
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// LinkText:
	LinkText string `json:"link_text,omitempty" mapstructure:"link_text,omitempty"`
	// LinkTextIc:
	LinkTextIc string `json:"link_text__ic,omitempty" mapstructure:"link_text__ic,omitempty"`
	// LinkTextIe:
	LinkTextIe string `json:"link_text__ie,omitempty" mapstructure:"link_text__ie,omitempty"`
	// LinkTextIew:
	LinkTextIew string `json:"link_text__iew,omitempty" mapstructure:"link_text__iew,omitempty"`
	// LinkTextIsw:
	LinkTextIsw string `json:"link_text__isw,omitempty" mapstructure:"link_text__isw,omitempty"`
	// LinkTextN:
	LinkTextN string `json:"link_text__n,omitempty" mapstructure:"link_text__n,omitempty"`
	// LinkTextNic:
	LinkTextNic string `json:"link_text__nic,omitempty" mapstructure:"link_text__nic,omitempty"`
	// LinkTextNie:
	LinkTextNie string `json:"link_text__nie,omitempty" mapstructure:"link_text__nie,omitempty"`
	// LinkTextNiew:
	LinkTextNiew string `json:"link_text__niew,omitempty" mapstructure:"link_text__niew,omitempty"`
	// LinkTextNisw:
	LinkTextNisw string `json:"link_text__nisw,omitempty" mapstructure:"link_text__nisw,omitempty"`
	// LinkUrl:
	LinkUrl string `json:"link_url,omitempty" mapstructure:"link_url,omitempty"`
	// LinkUrlIc:
	LinkUrlIc string `json:"link_url__ic,omitempty" mapstructure:"link_url__ic,omitempty"`
	// LinkUrlIe:
	LinkUrlIe string `json:"link_url__ie,omitempty" mapstructure:"link_url__ie,omitempty"`
	// LinkUrlIew:
	LinkUrlIew string `json:"link_url__iew,omitempty" mapstructure:"link_url__iew,omitempty"`
	// LinkUrlIsw:
	LinkUrlIsw string `json:"link_url__isw,omitempty" mapstructure:"link_url__isw,omitempty"`
	// LinkUrlN:
	LinkUrlN string `json:"link_url__n,omitempty" mapstructure:"link_url__n,omitempty"`
	// LinkUrlNic:
	LinkUrlNic string `json:"link_url__nic,omitempty" mapstructure:"link_url__nic,omitempty"`
	// LinkUrlNie:
	LinkUrlNie string `json:"link_url__nie,omitempty" mapstructure:"link_url__nie,omitempty"`
	// LinkUrlNiew:
	LinkUrlNiew string `json:"link_url__niew,omitempty" mapstructure:"link_url__niew,omitempty"`
	// LinkUrlNisw:
	LinkUrlNisw string `json:"link_url__nisw,omitempty" mapstructure:"link_url__nisw,omitempty"`
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
	// NewWindow:
	NewWindow bool `json:"new_window,omitempty" mapstructure:"new_window,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
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
func (m ExtrasCustomLinksListQueryParameters) Validate() error {
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
func (m ExtrasCustomLinksListQueryParameters) GetContentTypeId() []int32 {
	return m.ContentTypeId
}

// SetContentTypeId sets the ContentTypeId property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypeId(val []int32) {
	m.ContentTypeId = val
}

// GetContentTypeIdGt returns the ContentTypeIdGt property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypeIdGt() []int32 {
	return m.ContentTypeIdGt
}

// SetContentTypeIdGt sets the ContentTypeIdGt property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypeIdGt(val []int32) {
	m.ContentTypeIdGt = val
}

// GetContentTypeIdGte returns the ContentTypeIdGte property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypeIdGte() []int32 {
	return m.ContentTypeIdGte
}

// SetContentTypeIdGte sets the ContentTypeIdGte property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypeIdGte(val []int32) {
	m.ContentTypeIdGte = val
}

// GetContentTypeIdLt returns the ContentTypeIdLt property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypeIdLt() []int32 {
	return m.ContentTypeIdLt
}

// SetContentTypeIdLt sets the ContentTypeIdLt property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypeIdLt(val []int32) {
	m.ContentTypeIdLt = val
}

// GetContentTypeIdLte returns the ContentTypeIdLte property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypeIdLte() []int32 {
	return m.ContentTypeIdLte
}

// SetContentTypeIdLte sets the ContentTypeIdLte property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypeIdLte(val []int32) {
	m.ContentTypeIdLte = val
}

// GetContentTypeIdN returns the ContentTypeIdN property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypeIdN() []int32 {
	return m.ContentTypeIdN
}

// SetContentTypeIdN sets the ContentTypeIdN property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypeIdN(val []int32) {
	m.ContentTypeIdN = val
}

// GetContentTypes returns the ContentTypes property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypes() string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypes(val string) {
	m.ContentTypes = val
}

// GetContentTypesIc returns the ContentTypesIc property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesIc() string {
	return m.ContentTypesIc
}

// SetContentTypesIc sets the ContentTypesIc property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesIc(val string) {
	m.ContentTypesIc = val
}

// GetContentTypesIe returns the ContentTypesIe property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesIe() string {
	return m.ContentTypesIe
}

// SetContentTypesIe sets the ContentTypesIe property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesIe(val string) {
	m.ContentTypesIe = val
}

// GetContentTypesIew returns the ContentTypesIew property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesIew() string {
	return m.ContentTypesIew
}

// SetContentTypesIew sets the ContentTypesIew property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesIew(val string) {
	m.ContentTypesIew = val
}

// GetContentTypesIsw returns the ContentTypesIsw property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesIsw() string {
	return m.ContentTypesIsw
}

// SetContentTypesIsw sets the ContentTypesIsw property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesIsw(val string) {
	m.ContentTypesIsw = val
}

// GetContentTypesN returns the ContentTypesN property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesN() string {
	return m.ContentTypesN
}

// SetContentTypesN sets the ContentTypesN property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesN(val string) {
	m.ContentTypesN = val
}

// GetContentTypesNic returns the ContentTypesNic property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesNic() string {
	return m.ContentTypesNic
}

// SetContentTypesNic sets the ContentTypesNic property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesNic(val string) {
	m.ContentTypesNic = val
}

// GetContentTypesNie returns the ContentTypesNie property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesNie() string {
	return m.ContentTypesNie
}

// SetContentTypesNie sets the ContentTypesNie property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesNie(val string) {
	m.ContentTypesNie = val
}

// GetContentTypesNiew returns the ContentTypesNiew property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesNiew() string {
	return m.ContentTypesNiew
}

// SetContentTypesNiew sets the ContentTypesNiew property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesNiew(val string) {
	m.ContentTypesNiew = val
}

// GetContentTypesNisw returns the ContentTypesNisw property
func (m ExtrasCustomLinksListQueryParameters) GetContentTypesNisw() string {
	return m.ContentTypesNisw
}

// SetContentTypesNisw sets the ContentTypesNisw property
func (m *ExtrasCustomLinksListQueryParameters) SetContentTypesNisw(val string) {
	m.ContentTypesNisw = val
}

// GetEnabled returns the Enabled property
func (m ExtrasCustomLinksListQueryParameters) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *ExtrasCustomLinksListQueryParameters) SetEnabled(val bool) {
	m.Enabled = val
}

// GetGroupName returns the GroupName property
func (m ExtrasCustomLinksListQueryParameters) GetGroupName() []string {
	return m.GroupName
}

// SetGroupName sets the GroupName property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupName(val []string) {
	m.GroupName = val
}

// GetGroupNameEmpty returns the GroupNameEmpty property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameEmpty() []string {
	return m.GroupNameEmpty
}

// SetGroupNameEmpty sets the GroupNameEmpty property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameEmpty(val []string) {
	m.GroupNameEmpty = val
}

// GetGroupNameIc returns the GroupNameIc property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameIc() []string {
	return m.GroupNameIc
}

// SetGroupNameIc sets the GroupNameIc property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameIc(val []string) {
	m.GroupNameIc = val
}

// GetGroupNameIe returns the GroupNameIe property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameIe() []string {
	return m.GroupNameIe
}

// SetGroupNameIe sets the GroupNameIe property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameIe(val []string) {
	m.GroupNameIe = val
}

// GetGroupNameIew returns the GroupNameIew property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameIew() []string {
	return m.GroupNameIew
}

// SetGroupNameIew sets the GroupNameIew property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameIew(val []string) {
	m.GroupNameIew = val
}

// GetGroupNameIsw returns the GroupNameIsw property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameIsw() []string {
	return m.GroupNameIsw
}

// SetGroupNameIsw sets the GroupNameIsw property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameIsw(val []string) {
	m.GroupNameIsw = val
}

// GetGroupNameN returns the GroupNameN property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameN() []string {
	return m.GroupNameN
}

// SetGroupNameN sets the GroupNameN property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameN(val []string) {
	m.GroupNameN = val
}

// GetGroupNameNic returns the GroupNameNic property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameNic() []string {
	return m.GroupNameNic
}

// SetGroupNameNic sets the GroupNameNic property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameNic(val []string) {
	m.GroupNameNic = val
}

// GetGroupNameNie returns the GroupNameNie property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameNie() []string {
	return m.GroupNameNie
}

// SetGroupNameNie sets the GroupNameNie property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameNie(val []string) {
	m.GroupNameNie = val
}

// GetGroupNameNiew returns the GroupNameNiew property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameNiew() []string {
	return m.GroupNameNiew
}

// SetGroupNameNiew sets the GroupNameNiew property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameNiew(val []string) {
	m.GroupNameNiew = val
}

// GetGroupNameNisw returns the GroupNameNisw property
func (m ExtrasCustomLinksListQueryParameters) GetGroupNameNisw() []string {
	return m.GroupNameNisw
}

// SetGroupNameNisw sets the GroupNameNisw property
func (m *ExtrasCustomLinksListQueryParameters) SetGroupNameNisw(val []string) {
	m.GroupNameNisw = val
}

// GetId returns the Id property
func (m ExtrasCustomLinksListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasCustomLinksListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasCustomLinksListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasCustomLinksListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasCustomLinksListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasCustomLinksListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasCustomLinksListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasCustomLinksListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasCustomLinksListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasCustomLinksListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasCustomLinksListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasCustomLinksListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m ExtrasCustomLinksListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasCustomLinksListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLinkText returns the LinkText property
func (m ExtrasCustomLinksListQueryParameters) GetLinkText() string {
	return m.LinkText
}

// SetLinkText sets the LinkText property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkText(val string) {
	m.LinkText = val
}

// GetLinkTextIc returns the LinkTextIc property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextIc() string {
	return m.LinkTextIc
}

// SetLinkTextIc sets the LinkTextIc property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextIc(val string) {
	m.LinkTextIc = val
}

// GetLinkTextIe returns the LinkTextIe property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextIe() string {
	return m.LinkTextIe
}

// SetLinkTextIe sets the LinkTextIe property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextIe(val string) {
	m.LinkTextIe = val
}

// GetLinkTextIew returns the LinkTextIew property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextIew() string {
	return m.LinkTextIew
}

// SetLinkTextIew sets the LinkTextIew property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextIew(val string) {
	m.LinkTextIew = val
}

// GetLinkTextIsw returns the LinkTextIsw property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextIsw() string {
	return m.LinkTextIsw
}

// SetLinkTextIsw sets the LinkTextIsw property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextIsw(val string) {
	m.LinkTextIsw = val
}

// GetLinkTextN returns the LinkTextN property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextN() string {
	return m.LinkTextN
}

// SetLinkTextN sets the LinkTextN property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextN(val string) {
	m.LinkTextN = val
}

// GetLinkTextNic returns the LinkTextNic property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextNic() string {
	return m.LinkTextNic
}

// SetLinkTextNic sets the LinkTextNic property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextNic(val string) {
	m.LinkTextNic = val
}

// GetLinkTextNie returns the LinkTextNie property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextNie() string {
	return m.LinkTextNie
}

// SetLinkTextNie sets the LinkTextNie property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextNie(val string) {
	m.LinkTextNie = val
}

// GetLinkTextNiew returns the LinkTextNiew property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextNiew() string {
	return m.LinkTextNiew
}

// SetLinkTextNiew sets the LinkTextNiew property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextNiew(val string) {
	m.LinkTextNiew = val
}

// GetLinkTextNisw returns the LinkTextNisw property
func (m ExtrasCustomLinksListQueryParameters) GetLinkTextNisw() string {
	return m.LinkTextNisw
}

// SetLinkTextNisw sets the LinkTextNisw property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkTextNisw(val string) {
	m.LinkTextNisw = val
}

// GetLinkUrl returns the LinkUrl property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrl() string {
	return m.LinkUrl
}

// SetLinkUrl sets the LinkUrl property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrl(val string) {
	m.LinkUrl = val
}

// GetLinkUrlIc returns the LinkUrlIc property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlIc() string {
	return m.LinkUrlIc
}

// SetLinkUrlIc sets the LinkUrlIc property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlIc(val string) {
	m.LinkUrlIc = val
}

// GetLinkUrlIe returns the LinkUrlIe property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlIe() string {
	return m.LinkUrlIe
}

// SetLinkUrlIe sets the LinkUrlIe property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlIe(val string) {
	m.LinkUrlIe = val
}

// GetLinkUrlIew returns the LinkUrlIew property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlIew() string {
	return m.LinkUrlIew
}

// SetLinkUrlIew sets the LinkUrlIew property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlIew(val string) {
	m.LinkUrlIew = val
}

// GetLinkUrlIsw returns the LinkUrlIsw property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlIsw() string {
	return m.LinkUrlIsw
}

// SetLinkUrlIsw sets the LinkUrlIsw property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlIsw(val string) {
	m.LinkUrlIsw = val
}

// GetLinkUrlN returns the LinkUrlN property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlN() string {
	return m.LinkUrlN
}

// SetLinkUrlN sets the LinkUrlN property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlN(val string) {
	m.LinkUrlN = val
}

// GetLinkUrlNic returns the LinkUrlNic property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlNic() string {
	return m.LinkUrlNic
}

// SetLinkUrlNic sets the LinkUrlNic property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlNic(val string) {
	m.LinkUrlNic = val
}

// GetLinkUrlNie returns the LinkUrlNie property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlNie() string {
	return m.LinkUrlNie
}

// SetLinkUrlNie sets the LinkUrlNie property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlNie(val string) {
	m.LinkUrlNie = val
}

// GetLinkUrlNiew returns the LinkUrlNiew property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlNiew() string {
	return m.LinkUrlNiew
}

// SetLinkUrlNiew sets the LinkUrlNiew property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlNiew(val string) {
	m.LinkUrlNiew = val
}

// GetLinkUrlNisw returns the LinkUrlNisw property
func (m ExtrasCustomLinksListQueryParameters) GetLinkUrlNisw() string {
	return m.LinkUrlNisw
}

// SetLinkUrlNisw sets the LinkUrlNisw property
func (m *ExtrasCustomLinksListQueryParameters) SetLinkUrlNisw(val string) {
	m.LinkUrlNisw = val
}

// GetName returns the Name property
func (m ExtrasCustomLinksListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasCustomLinksListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasCustomLinksListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasCustomLinksListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasCustomLinksListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasCustomLinksListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasCustomLinksListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasCustomLinksListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasCustomLinksListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasCustomLinksListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasCustomLinksListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasCustomLinksListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasCustomLinksListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasCustomLinksListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasCustomLinksListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasCustomLinksListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasCustomLinksListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasCustomLinksListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasCustomLinksListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasCustomLinksListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasCustomLinksListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasCustomLinksListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetNewWindow returns the NewWindow property
func (m ExtrasCustomLinksListQueryParameters) GetNewWindow() bool {
	return m.NewWindow
}

// SetNewWindow sets the NewWindow property
func (m *ExtrasCustomLinksListQueryParameters) SetNewWindow(val bool) {
	m.NewWindow = val
}

// GetOffset returns the Offset property
func (m ExtrasCustomLinksListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasCustomLinksListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasCustomLinksListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasCustomLinksListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasCustomLinksListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasCustomLinksListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetWeight returns the Weight property
func (m ExtrasCustomLinksListQueryParameters) GetWeight() []int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *ExtrasCustomLinksListQueryParameters) SetWeight(val []int32) {
	m.Weight = val
}

// GetWeightGt returns the WeightGt property
func (m ExtrasCustomLinksListQueryParameters) GetWeightGt() []int32 {
	return m.WeightGt
}

// SetWeightGt sets the WeightGt property
func (m *ExtrasCustomLinksListQueryParameters) SetWeightGt(val []int32) {
	m.WeightGt = val
}

// GetWeightGte returns the WeightGte property
func (m ExtrasCustomLinksListQueryParameters) GetWeightGte() []int32 {
	return m.WeightGte
}

// SetWeightGte sets the WeightGte property
func (m *ExtrasCustomLinksListQueryParameters) SetWeightGte(val []int32) {
	m.WeightGte = val
}

// GetWeightLt returns the WeightLt property
func (m ExtrasCustomLinksListQueryParameters) GetWeightLt() []int32 {
	return m.WeightLt
}

// SetWeightLt sets the WeightLt property
func (m *ExtrasCustomLinksListQueryParameters) SetWeightLt(val []int32) {
	m.WeightLt = val
}

// GetWeightLte returns the WeightLte property
func (m ExtrasCustomLinksListQueryParameters) GetWeightLte() []int32 {
	return m.WeightLte
}

// SetWeightLte sets the WeightLte property
func (m *ExtrasCustomLinksListQueryParameters) SetWeightLte(val []int32) {
	m.WeightLte = val
}

// GetWeightN returns the WeightN property
func (m ExtrasCustomLinksListQueryParameters) GetWeightN() []int32 {
	return m.WeightN
}

// SetWeightN sets the WeightN property
func (m *ExtrasCustomLinksListQueryParameters) SetWeightN(val []int32) {
	m.WeightN = val
}
