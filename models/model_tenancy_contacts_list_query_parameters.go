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

// TenancyContactsListQueryParameters is an object.
type TenancyContactsListQueryParameters struct {
	// Address:
	Address []string `json:"address,omitempty" mapstructure:"address,omitempty"`
	// AddressEmpty:
	AddressEmpty []string `json:"address__empty,omitempty" mapstructure:"address__empty,omitempty"`
	// AddressIc:
	AddressIc []string `json:"address__ic,omitempty" mapstructure:"address__ic,omitempty"`
	// AddressIe:
	AddressIe []string `json:"address__ie,omitempty" mapstructure:"address__ie,omitempty"`
	// AddressIew:
	AddressIew []string `json:"address__iew,omitempty" mapstructure:"address__iew,omitempty"`
	// AddressIsw:
	AddressIsw []string `json:"address__isw,omitempty" mapstructure:"address__isw,omitempty"`
	// AddressN:
	AddressN []string `json:"address__n,omitempty" mapstructure:"address__n,omitempty"`
	// AddressNic:
	AddressNic []string `json:"address__nic,omitempty" mapstructure:"address__nic,omitempty"`
	// AddressNie:
	AddressNie []string `json:"address__nie,omitempty" mapstructure:"address__nie,omitempty"`
	// AddressNiew:
	AddressNiew []string `json:"address__niew,omitempty" mapstructure:"address__niew,omitempty"`
	// AddressNisw:
	AddressNisw []string `json:"address__nisw,omitempty" mapstructure:"address__nisw,omitempty"`
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
	// Email:
	Email []string `json:"email,omitempty" mapstructure:"email,omitempty"`
	// EmailEmpty:
	EmailEmpty []string `json:"email__empty,omitempty" mapstructure:"email__empty,omitempty"`
	// EmailIc:
	EmailIc []string `json:"email__ic,omitempty" mapstructure:"email__ic,omitempty"`
	// EmailIe:
	EmailIe []string `json:"email__ie,omitempty" mapstructure:"email__ie,omitempty"`
	// EmailIew:
	EmailIew []string `json:"email__iew,omitempty" mapstructure:"email__iew,omitempty"`
	// EmailIsw:
	EmailIsw []string `json:"email__isw,omitempty" mapstructure:"email__isw,omitempty"`
	// EmailN:
	EmailN []string `json:"email__n,omitempty" mapstructure:"email__n,omitempty"`
	// EmailNic:
	EmailNic []string `json:"email__nic,omitempty" mapstructure:"email__nic,omitempty"`
	// EmailNie:
	EmailNie []string `json:"email__nie,omitempty" mapstructure:"email__nie,omitempty"`
	// EmailNiew:
	EmailNiew []string `json:"email__niew,omitempty" mapstructure:"email__niew,omitempty"`
	// EmailNisw:
	EmailNisw []string `json:"email__nisw,omitempty" mapstructure:"email__nisw,omitempty"`
	// Group: Contact group (slug)
	Group []int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// GroupN: Contact group (slug)
	GroupN []int32 `json:"group__n,omitempty" mapstructure:"group__n,omitempty"`
	// GroupId: Contact group (ID)
	GroupId []int32 `json:"group_id,omitempty" mapstructure:"group_id,omitempty"`
	// GroupIdN: Contact group (ID)
	GroupIdN []int32 `json:"group_id__n,omitempty" mapstructure:"group_id__n,omitempty"`
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
	// Link:
	Link []string `json:"link,omitempty" mapstructure:"link,omitempty"`
	// LinkEmpty:
	LinkEmpty []string `json:"link__empty,omitempty" mapstructure:"link__empty,omitempty"`
	// LinkIc:
	LinkIc []string `json:"link__ic,omitempty" mapstructure:"link__ic,omitempty"`
	// LinkIe:
	LinkIe []string `json:"link__ie,omitempty" mapstructure:"link__ie,omitempty"`
	// LinkIew:
	LinkIew []string `json:"link__iew,omitempty" mapstructure:"link__iew,omitempty"`
	// LinkIsw:
	LinkIsw []string `json:"link__isw,omitempty" mapstructure:"link__isw,omitempty"`
	// LinkN:
	LinkN []string `json:"link__n,omitempty" mapstructure:"link__n,omitempty"`
	// LinkNic:
	LinkNic []string `json:"link__nic,omitempty" mapstructure:"link__nic,omitempty"`
	// LinkNie:
	LinkNie []string `json:"link__nie,omitempty" mapstructure:"link__nie,omitempty"`
	// LinkNiew:
	LinkNiew []string `json:"link__niew,omitempty" mapstructure:"link__niew,omitempty"`
	// LinkNisw:
	LinkNisw []string `json:"link__nisw,omitempty" mapstructure:"link__nisw,omitempty"`
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
	// Phone:
	Phone []string `json:"phone,omitempty" mapstructure:"phone,omitempty"`
	// PhoneEmpty:
	PhoneEmpty []string `json:"phone__empty,omitempty" mapstructure:"phone__empty,omitempty"`
	// PhoneIc:
	PhoneIc []string `json:"phone__ic,omitempty" mapstructure:"phone__ic,omitempty"`
	// PhoneIe:
	PhoneIe []string `json:"phone__ie,omitempty" mapstructure:"phone__ie,omitempty"`
	// PhoneIew:
	PhoneIew []string `json:"phone__iew,omitempty" mapstructure:"phone__iew,omitempty"`
	// PhoneIsw:
	PhoneIsw []string `json:"phone__isw,omitempty" mapstructure:"phone__isw,omitempty"`
	// PhoneN:
	PhoneN []string `json:"phone__n,omitempty" mapstructure:"phone__n,omitempty"`
	// PhoneNic:
	PhoneNic []string `json:"phone__nic,omitempty" mapstructure:"phone__nic,omitempty"`
	// PhoneNie:
	PhoneNie []string `json:"phone__nie,omitempty" mapstructure:"phone__nie,omitempty"`
	// PhoneNiew:
	PhoneNiew []string `json:"phone__niew,omitempty" mapstructure:"phone__niew,omitempty"`
	// PhoneNisw:
	PhoneNisw []string `json:"phone__nisw,omitempty" mapstructure:"phone__nisw,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// Title:
	Title []string `json:"title,omitempty" mapstructure:"title,omitempty"`
	// TitleEmpty:
	TitleEmpty []string `json:"title__empty,omitempty" mapstructure:"title__empty,omitempty"`
	// TitleIc:
	TitleIc []string `json:"title__ic,omitempty" mapstructure:"title__ic,omitempty"`
	// TitleIe:
	TitleIe []string `json:"title__ie,omitempty" mapstructure:"title__ie,omitempty"`
	// TitleIew:
	TitleIew []string `json:"title__iew,omitempty" mapstructure:"title__iew,omitempty"`
	// TitleIsw:
	TitleIsw []string `json:"title__isw,omitempty" mapstructure:"title__isw,omitempty"`
	// TitleN:
	TitleN []string `json:"title__n,omitempty" mapstructure:"title__n,omitempty"`
	// TitleNic:
	TitleNic []string `json:"title__nic,omitempty" mapstructure:"title__nic,omitempty"`
	// TitleNie:
	TitleNie []string `json:"title__nie,omitempty" mapstructure:"title__nie,omitempty"`
	// TitleNiew:
	TitleNiew []string `json:"title__niew,omitempty" mapstructure:"title__niew,omitempty"`
	// TitleNisw:
	TitleNisw []string `json:"title__nisw,omitempty" mapstructure:"title__nisw,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m TenancyContactsListQueryParameters) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address,
		),
		"addressEmpty": validation.Validate(
			m.AddressEmpty,
		),
		"addressIc": validation.Validate(
			m.AddressIc,
		),
		"addressIe": validation.Validate(
			m.AddressIe,
		),
		"addressIew": validation.Validate(
			m.AddressIew,
		),
		"addressIsw": validation.Validate(
			m.AddressIsw,
		),
		"addressN": validation.Validate(
			m.AddressN,
		),
		"addressNic": validation.Validate(
			m.AddressNic,
		),
		"addressNie": validation.Validate(
			m.AddressNie,
		),
		"addressNiew": validation.Validate(
			m.AddressNiew,
		),
		"addressNisw": validation.Validate(
			m.AddressNisw,
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
		"email": validation.Validate(
			m.Email,
		),
		"emailEmpty": validation.Validate(
			m.EmailEmpty,
		),
		"emailIc": validation.Validate(
			m.EmailIc,
		),
		"emailIe": validation.Validate(
			m.EmailIe,
		),
		"emailIew": validation.Validate(
			m.EmailIew,
		),
		"emailIsw": validation.Validate(
			m.EmailIsw,
		),
		"emailN": validation.Validate(
			m.EmailN,
		),
		"emailNic": validation.Validate(
			m.EmailNic,
		),
		"emailNie": validation.Validate(
			m.EmailNie,
		),
		"emailNiew": validation.Validate(
			m.EmailNiew,
		),
		"emailNisw": validation.Validate(
			m.EmailNisw,
		),
		"group": validation.Validate(
			m.Group,
		),
		"groupN": validation.Validate(
			m.GroupN,
		),
		"groupId": validation.Validate(
			m.GroupId,
		),
		"groupIdN": validation.Validate(
			m.GroupIdN,
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
		"link": validation.Validate(
			m.Link,
		),
		"linkEmpty": validation.Validate(
			m.LinkEmpty,
		),
		"linkIc": validation.Validate(
			m.LinkIc,
		),
		"linkIe": validation.Validate(
			m.LinkIe,
		),
		"linkIew": validation.Validate(
			m.LinkIew,
		),
		"linkIsw": validation.Validate(
			m.LinkIsw,
		),
		"linkN": validation.Validate(
			m.LinkN,
		),
		"linkNic": validation.Validate(
			m.LinkNic,
		),
		"linkNie": validation.Validate(
			m.LinkNie,
		),
		"linkNiew": validation.Validate(
			m.LinkNiew,
		),
		"linkNisw": validation.Validate(
			m.LinkNisw,
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
		"phone": validation.Validate(
			m.Phone,
		),
		"phoneEmpty": validation.Validate(
			m.PhoneEmpty,
		),
		"phoneIc": validation.Validate(
			m.PhoneIc,
		),
		"phoneIe": validation.Validate(
			m.PhoneIe,
		),
		"phoneIew": validation.Validate(
			m.PhoneIew,
		),
		"phoneIsw": validation.Validate(
			m.PhoneIsw,
		),
		"phoneN": validation.Validate(
			m.PhoneN,
		),
		"phoneNic": validation.Validate(
			m.PhoneNic,
		),
		"phoneNie": validation.Validate(
			m.PhoneNie,
		),
		"phoneNiew": validation.Validate(
			m.PhoneNiew,
		),
		"phoneNisw": validation.Validate(
			m.PhoneNisw,
		),
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"title": validation.Validate(
			m.Title,
		),
		"titleEmpty": validation.Validate(
			m.TitleEmpty,
		),
		"titleIc": validation.Validate(
			m.TitleIc,
		),
		"titleIe": validation.Validate(
			m.TitleIe,
		),
		"titleIew": validation.Validate(
			m.TitleIew,
		),
		"titleIsw": validation.Validate(
			m.TitleIsw,
		),
		"titleN": validation.Validate(
			m.TitleN,
		),
		"titleNic": validation.Validate(
			m.TitleNic,
		),
		"titleNie": validation.Validate(
			m.TitleNie,
		),
		"titleNiew": validation.Validate(
			m.TitleNiew,
		),
		"titleNisw": validation.Validate(
			m.TitleNisw,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m TenancyContactsListQueryParameters) GetAddress() []string {
	return m.Address
}

// SetAddress sets the Address property
func (m *TenancyContactsListQueryParameters) SetAddress(val []string) {
	m.Address = val
}

// GetAddressEmpty returns the AddressEmpty property
func (m TenancyContactsListQueryParameters) GetAddressEmpty() []string {
	return m.AddressEmpty
}

// SetAddressEmpty sets the AddressEmpty property
func (m *TenancyContactsListQueryParameters) SetAddressEmpty(val []string) {
	m.AddressEmpty = val
}

// GetAddressIc returns the AddressIc property
func (m TenancyContactsListQueryParameters) GetAddressIc() []string {
	return m.AddressIc
}

// SetAddressIc sets the AddressIc property
func (m *TenancyContactsListQueryParameters) SetAddressIc(val []string) {
	m.AddressIc = val
}

// GetAddressIe returns the AddressIe property
func (m TenancyContactsListQueryParameters) GetAddressIe() []string {
	return m.AddressIe
}

// SetAddressIe sets the AddressIe property
func (m *TenancyContactsListQueryParameters) SetAddressIe(val []string) {
	m.AddressIe = val
}

// GetAddressIew returns the AddressIew property
func (m TenancyContactsListQueryParameters) GetAddressIew() []string {
	return m.AddressIew
}

// SetAddressIew sets the AddressIew property
func (m *TenancyContactsListQueryParameters) SetAddressIew(val []string) {
	m.AddressIew = val
}

// GetAddressIsw returns the AddressIsw property
func (m TenancyContactsListQueryParameters) GetAddressIsw() []string {
	return m.AddressIsw
}

// SetAddressIsw sets the AddressIsw property
func (m *TenancyContactsListQueryParameters) SetAddressIsw(val []string) {
	m.AddressIsw = val
}

// GetAddressN returns the AddressN property
func (m TenancyContactsListQueryParameters) GetAddressN() []string {
	return m.AddressN
}

// SetAddressN sets the AddressN property
func (m *TenancyContactsListQueryParameters) SetAddressN(val []string) {
	m.AddressN = val
}

// GetAddressNic returns the AddressNic property
func (m TenancyContactsListQueryParameters) GetAddressNic() []string {
	return m.AddressNic
}

// SetAddressNic sets the AddressNic property
func (m *TenancyContactsListQueryParameters) SetAddressNic(val []string) {
	m.AddressNic = val
}

// GetAddressNie returns the AddressNie property
func (m TenancyContactsListQueryParameters) GetAddressNie() []string {
	return m.AddressNie
}

// SetAddressNie sets the AddressNie property
func (m *TenancyContactsListQueryParameters) SetAddressNie(val []string) {
	m.AddressNie = val
}

// GetAddressNiew returns the AddressNiew property
func (m TenancyContactsListQueryParameters) GetAddressNiew() []string {
	return m.AddressNiew
}

// SetAddressNiew sets the AddressNiew property
func (m *TenancyContactsListQueryParameters) SetAddressNiew(val []string) {
	m.AddressNiew = val
}

// GetAddressNisw returns the AddressNisw property
func (m TenancyContactsListQueryParameters) GetAddressNisw() []string {
	return m.AddressNisw
}

// SetAddressNisw sets the AddressNisw property
func (m *TenancyContactsListQueryParameters) SetAddressNisw(val []string) {
	m.AddressNisw = val
}

// GetCreated returns the Created property
func (m TenancyContactsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *TenancyContactsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m TenancyContactsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *TenancyContactsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m TenancyContactsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *TenancyContactsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m TenancyContactsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *TenancyContactsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m TenancyContactsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *TenancyContactsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m TenancyContactsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *TenancyContactsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m TenancyContactsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *TenancyContactsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetEmail returns the Email property
func (m TenancyContactsListQueryParameters) GetEmail() []string {
	return m.Email
}

// SetEmail sets the Email property
func (m *TenancyContactsListQueryParameters) SetEmail(val []string) {
	m.Email = val
}

// GetEmailEmpty returns the EmailEmpty property
func (m TenancyContactsListQueryParameters) GetEmailEmpty() []string {
	return m.EmailEmpty
}

// SetEmailEmpty sets the EmailEmpty property
func (m *TenancyContactsListQueryParameters) SetEmailEmpty(val []string) {
	m.EmailEmpty = val
}

// GetEmailIc returns the EmailIc property
func (m TenancyContactsListQueryParameters) GetEmailIc() []string {
	return m.EmailIc
}

// SetEmailIc sets the EmailIc property
func (m *TenancyContactsListQueryParameters) SetEmailIc(val []string) {
	m.EmailIc = val
}

// GetEmailIe returns the EmailIe property
func (m TenancyContactsListQueryParameters) GetEmailIe() []string {
	return m.EmailIe
}

// SetEmailIe sets the EmailIe property
func (m *TenancyContactsListQueryParameters) SetEmailIe(val []string) {
	m.EmailIe = val
}

// GetEmailIew returns the EmailIew property
func (m TenancyContactsListQueryParameters) GetEmailIew() []string {
	return m.EmailIew
}

// SetEmailIew sets the EmailIew property
func (m *TenancyContactsListQueryParameters) SetEmailIew(val []string) {
	m.EmailIew = val
}

// GetEmailIsw returns the EmailIsw property
func (m TenancyContactsListQueryParameters) GetEmailIsw() []string {
	return m.EmailIsw
}

// SetEmailIsw sets the EmailIsw property
func (m *TenancyContactsListQueryParameters) SetEmailIsw(val []string) {
	m.EmailIsw = val
}

// GetEmailN returns the EmailN property
func (m TenancyContactsListQueryParameters) GetEmailN() []string {
	return m.EmailN
}

// SetEmailN sets the EmailN property
func (m *TenancyContactsListQueryParameters) SetEmailN(val []string) {
	m.EmailN = val
}

// GetEmailNic returns the EmailNic property
func (m TenancyContactsListQueryParameters) GetEmailNic() []string {
	return m.EmailNic
}

// SetEmailNic sets the EmailNic property
func (m *TenancyContactsListQueryParameters) SetEmailNic(val []string) {
	m.EmailNic = val
}

// GetEmailNie returns the EmailNie property
func (m TenancyContactsListQueryParameters) GetEmailNie() []string {
	return m.EmailNie
}

// SetEmailNie sets the EmailNie property
func (m *TenancyContactsListQueryParameters) SetEmailNie(val []string) {
	m.EmailNie = val
}

// GetEmailNiew returns the EmailNiew property
func (m TenancyContactsListQueryParameters) GetEmailNiew() []string {
	return m.EmailNiew
}

// SetEmailNiew sets the EmailNiew property
func (m *TenancyContactsListQueryParameters) SetEmailNiew(val []string) {
	m.EmailNiew = val
}

// GetEmailNisw returns the EmailNisw property
func (m TenancyContactsListQueryParameters) GetEmailNisw() []string {
	return m.EmailNisw
}

// SetEmailNisw sets the EmailNisw property
func (m *TenancyContactsListQueryParameters) SetEmailNisw(val []string) {
	m.EmailNisw = val
}

// GetGroup returns the Group property
func (m TenancyContactsListQueryParameters) GetGroup() []int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *TenancyContactsListQueryParameters) SetGroup(val []int32) {
	m.Group = val
}

// GetGroupN returns the GroupN property
func (m TenancyContactsListQueryParameters) GetGroupN() []int32 {
	return m.GroupN
}

// SetGroupN sets the GroupN property
func (m *TenancyContactsListQueryParameters) SetGroupN(val []int32) {
	m.GroupN = val
}

// GetGroupId returns the GroupId property
func (m TenancyContactsListQueryParameters) GetGroupId() []int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *TenancyContactsListQueryParameters) SetGroupId(val []int32) {
	m.GroupId = val
}

// GetGroupIdN returns the GroupIdN property
func (m TenancyContactsListQueryParameters) GetGroupIdN() []int32 {
	return m.GroupIdN
}

// SetGroupIdN sets the GroupIdN property
func (m *TenancyContactsListQueryParameters) SetGroupIdN(val []int32) {
	m.GroupIdN = val
}

// GetId returns the Id property
func (m TenancyContactsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *TenancyContactsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m TenancyContactsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *TenancyContactsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m TenancyContactsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *TenancyContactsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m TenancyContactsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *TenancyContactsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m TenancyContactsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *TenancyContactsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m TenancyContactsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *TenancyContactsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m TenancyContactsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *TenancyContactsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m TenancyContactsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *TenancyContactsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m TenancyContactsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *TenancyContactsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m TenancyContactsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *TenancyContactsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m TenancyContactsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *TenancyContactsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m TenancyContactsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *TenancyContactsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m TenancyContactsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *TenancyContactsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLink returns the Link property
func (m TenancyContactsListQueryParameters) GetLink() []string {
	return m.Link
}

// SetLink sets the Link property
func (m *TenancyContactsListQueryParameters) SetLink(val []string) {
	m.Link = val
}

// GetLinkEmpty returns the LinkEmpty property
func (m TenancyContactsListQueryParameters) GetLinkEmpty() []string {
	return m.LinkEmpty
}

// SetLinkEmpty sets the LinkEmpty property
func (m *TenancyContactsListQueryParameters) SetLinkEmpty(val []string) {
	m.LinkEmpty = val
}

// GetLinkIc returns the LinkIc property
func (m TenancyContactsListQueryParameters) GetLinkIc() []string {
	return m.LinkIc
}

// SetLinkIc sets the LinkIc property
func (m *TenancyContactsListQueryParameters) SetLinkIc(val []string) {
	m.LinkIc = val
}

// GetLinkIe returns the LinkIe property
func (m TenancyContactsListQueryParameters) GetLinkIe() []string {
	return m.LinkIe
}

// SetLinkIe sets the LinkIe property
func (m *TenancyContactsListQueryParameters) SetLinkIe(val []string) {
	m.LinkIe = val
}

// GetLinkIew returns the LinkIew property
func (m TenancyContactsListQueryParameters) GetLinkIew() []string {
	return m.LinkIew
}

// SetLinkIew sets the LinkIew property
func (m *TenancyContactsListQueryParameters) SetLinkIew(val []string) {
	m.LinkIew = val
}

// GetLinkIsw returns the LinkIsw property
func (m TenancyContactsListQueryParameters) GetLinkIsw() []string {
	return m.LinkIsw
}

// SetLinkIsw sets the LinkIsw property
func (m *TenancyContactsListQueryParameters) SetLinkIsw(val []string) {
	m.LinkIsw = val
}

// GetLinkN returns the LinkN property
func (m TenancyContactsListQueryParameters) GetLinkN() []string {
	return m.LinkN
}

// SetLinkN sets the LinkN property
func (m *TenancyContactsListQueryParameters) SetLinkN(val []string) {
	m.LinkN = val
}

// GetLinkNic returns the LinkNic property
func (m TenancyContactsListQueryParameters) GetLinkNic() []string {
	return m.LinkNic
}

// SetLinkNic sets the LinkNic property
func (m *TenancyContactsListQueryParameters) SetLinkNic(val []string) {
	m.LinkNic = val
}

// GetLinkNie returns the LinkNie property
func (m TenancyContactsListQueryParameters) GetLinkNie() []string {
	return m.LinkNie
}

// SetLinkNie sets the LinkNie property
func (m *TenancyContactsListQueryParameters) SetLinkNie(val []string) {
	m.LinkNie = val
}

// GetLinkNiew returns the LinkNiew property
func (m TenancyContactsListQueryParameters) GetLinkNiew() []string {
	return m.LinkNiew
}

// SetLinkNiew sets the LinkNiew property
func (m *TenancyContactsListQueryParameters) SetLinkNiew(val []string) {
	m.LinkNiew = val
}

// GetLinkNisw returns the LinkNisw property
func (m TenancyContactsListQueryParameters) GetLinkNisw() []string {
	return m.LinkNisw
}

// SetLinkNisw sets the LinkNisw property
func (m *TenancyContactsListQueryParameters) SetLinkNisw(val []string) {
	m.LinkNisw = val
}

// GetName returns the Name property
func (m TenancyContactsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *TenancyContactsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m TenancyContactsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *TenancyContactsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m TenancyContactsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *TenancyContactsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m TenancyContactsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *TenancyContactsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m TenancyContactsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *TenancyContactsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m TenancyContactsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *TenancyContactsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m TenancyContactsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *TenancyContactsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m TenancyContactsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *TenancyContactsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m TenancyContactsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *TenancyContactsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m TenancyContactsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *TenancyContactsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m TenancyContactsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *TenancyContactsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m TenancyContactsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *TenancyContactsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m TenancyContactsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *TenancyContactsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPhone returns the Phone property
func (m TenancyContactsListQueryParameters) GetPhone() []string {
	return m.Phone
}

// SetPhone sets the Phone property
func (m *TenancyContactsListQueryParameters) SetPhone(val []string) {
	m.Phone = val
}

// GetPhoneEmpty returns the PhoneEmpty property
func (m TenancyContactsListQueryParameters) GetPhoneEmpty() []string {
	return m.PhoneEmpty
}

// SetPhoneEmpty sets the PhoneEmpty property
func (m *TenancyContactsListQueryParameters) SetPhoneEmpty(val []string) {
	m.PhoneEmpty = val
}

// GetPhoneIc returns the PhoneIc property
func (m TenancyContactsListQueryParameters) GetPhoneIc() []string {
	return m.PhoneIc
}

// SetPhoneIc sets the PhoneIc property
func (m *TenancyContactsListQueryParameters) SetPhoneIc(val []string) {
	m.PhoneIc = val
}

// GetPhoneIe returns the PhoneIe property
func (m TenancyContactsListQueryParameters) GetPhoneIe() []string {
	return m.PhoneIe
}

// SetPhoneIe sets the PhoneIe property
func (m *TenancyContactsListQueryParameters) SetPhoneIe(val []string) {
	m.PhoneIe = val
}

// GetPhoneIew returns the PhoneIew property
func (m TenancyContactsListQueryParameters) GetPhoneIew() []string {
	return m.PhoneIew
}

// SetPhoneIew sets the PhoneIew property
func (m *TenancyContactsListQueryParameters) SetPhoneIew(val []string) {
	m.PhoneIew = val
}

// GetPhoneIsw returns the PhoneIsw property
func (m TenancyContactsListQueryParameters) GetPhoneIsw() []string {
	return m.PhoneIsw
}

// SetPhoneIsw sets the PhoneIsw property
func (m *TenancyContactsListQueryParameters) SetPhoneIsw(val []string) {
	m.PhoneIsw = val
}

// GetPhoneN returns the PhoneN property
func (m TenancyContactsListQueryParameters) GetPhoneN() []string {
	return m.PhoneN
}

// SetPhoneN sets the PhoneN property
func (m *TenancyContactsListQueryParameters) SetPhoneN(val []string) {
	m.PhoneN = val
}

// GetPhoneNic returns the PhoneNic property
func (m TenancyContactsListQueryParameters) GetPhoneNic() []string {
	return m.PhoneNic
}

// SetPhoneNic sets the PhoneNic property
func (m *TenancyContactsListQueryParameters) SetPhoneNic(val []string) {
	m.PhoneNic = val
}

// GetPhoneNie returns the PhoneNie property
func (m TenancyContactsListQueryParameters) GetPhoneNie() []string {
	return m.PhoneNie
}

// SetPhoneNie sets the PhoneNie property
func (m *TenancyContactsListQueryParameters) SetPhoneNie(val []string) {
	m.PhoneNie = val
}

// GetPhoneNiew returns the PhoneNiew property
func (m TenancyContactsListQueryParameters) GetPhoneNiew() []string {
	return m.PhoneNiew
}

// SetPhoneNiew sets the PhoneNiew property
func (m *TenancyContactsListQueryParameters) SetPhoneNiew(val []string) {
	m.PhoneNiew = val
}

// GetPhoneNisw returns the PhoneNisw property
func (m TenancyContactsListQueryParameters) GetPhoneNisw() []string {
	return m.PhoneNisw
}

// SetPhoneNisw sets the PhoneNisw property
func (m *TenancyContactsListQueryParameters) SetPhoneNisw(val []string) {
	m.PhoneNisw = val
}

// GetQ returns the Q property
func (m TenancyContactsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *TenancyContactsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m TenancyContactsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *TenancyContactsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m TenancyContactsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *TenancyContactsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTitle returns the Title property
func (m TenancyContactsListQueryParameters) GetTitle() []string {
	return m.Title
}

// SetTitle sets the Title property
func (m *TenancyContactsListQueryParameters) SetTitle(val []string) {
	m.Title = val
}

// GetTitleEmpty returns the TitleEmpty property
func (m TenancyContactsListQueryParameters) GetTitleEmpty() []string {
	return m.TitleEmpty
}

// SetTitleEmpty sets the TitleEmpty property
func (m *TenancyContactsListQueryParameters) SetTitleEmpty(val []string) {
	m.TitleEmpty = val
}

// GetTitleIc returns the TitleIc property
func (m TenancyContactsListQueryParameters) GetTitleIc() []string {
	return m.TitleIc
}

// SetTitleIc sets the TitleIc property
func (m *TenancyContactsListQueryParameters) SetTitleIc(val []string) {
	m.TitleIc = val
}

// GetTitleIe returns the TitleIe property
func (m TenancyContactsListQueryParameters) GetTitleIe() []string {
	return m.TitleIe
}

// SetTitleIe sets the TitleIe property
func (m *TenancyContactsListQueryParameters) SetTitleIe(val []string) {
	m.TitleIe = val
}

// GetTitleIew returns the TitleIew property
func (m TenancyContactsListQueryParameters) GetTitleIew() []string {
	return m.TitleIew
}

// SetTitleIew sets the TitleIew property
func (m *TenancyContactsListQueryParameters) SetTitleIew(val []string) {
	m.TitleIew = val
}

// GetTitleIsw returns the TitleIsw property
func (m TenancyContactsListQueryParameters) GetTitleIsw() []string {
	return m.TitleIsw
}

// SetTitleIsw sets the TitleIsw property
func (m *TenancyContactsListQueryParameters) SetTitleIsw(val []string) {
	m.TitleIsw = val
}

// GetTitleN returns the TitleN property
func (m TenancyContactsListQueryParameters) GetTitleN() []string {
	return m.TitleN
}

// SetTitleN sets the TitleN property
func (m *TenancyContactsListQueryParameters) SetTitleN(val []string) {
	m.TitleN = val
}

// GetTitleNic returns the TitleNic property
func (m TenancyContactsListQueryParameters) GetTitleNic() []string {
	return m.TitleNic
}

// SetTitleNic sets the TitleNic property
func (m *TenancyContactsListQueryParameters) SetTitleNic(val []string) {
	m.TitleNic = val
}

// GetTitleNie returns the TitleNie property
func (m TenancyContactsListQueryParameters) GetTitleNie() []string {
	return m.TitleNie
}

// SetTitleNie sets the TitleNie property
func (m *TenancyContactsListQueryParameters) SetTitleNie(val []string) {
	m.TitleNie = val
}

// GetTitleNiew returns the TitleNiew property
func (m TenancyContactsListQueryParameters) GetTitleNiew() []string {
	return m.TitleNiew
}

// SetTitleNiew sets the TitleNiew property
func (m *TenancyContactsListQueryParameters) SetTitleNiew(val []string) {
	m.TitleNiew = val
}

// GetTitleNisw returns the TitleNisw property
func (m TenancyContactsListQueryParameters) GetTitleNisw() []string {
	return m.TitleNisw
}

// SetTitleNisw sets the TitleNisw property
func (m *TenancyContactsListQueryParameters) SetTitleNisw(val []string) {
	m.TitleNisw = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m TenancyContactsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *TenancyContactsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
