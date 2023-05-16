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

// CircuitsProviderAccountsListQueryParameters is an object.
type CircuitsProviderAccountsListQueryParameters struct {
	// Account:
	Account []string `json:"account,omitempty" mapstructure:"account,omitempty"`
	// AccountEmpty:
	AccountEmpty []string `json:"account__empty,omitempty" mapstructure:"account__empty,omitempty"`
	// AccountIc:
	AccountIc []string `json:"account__ic,omitempty" mapstructure:"account__ic,omitempty"`
	// AccountIe:
	AccountIe []string `json:"account__ie,omitempty" mapstructure:"account__ie,omitempty"`
	// AccountIew:
	AccountIew []string `json:"account__iew,omitempty" mapstructure:"account__iew,omitempty"`
	// AccountIsw:
	AccountIsw []string `json:"account__isw,omitempty" mapstructure:"account__isw,omitempty"`
	// AccountN:
	AccountN []string `json:"account__n,omitempty" mapstructure:"account__n,omitempty"`
	// AccountNic:
	AccountNic []string `json:"account__nic,omitempty" mapstructure:"account__nic,omitempty"`
	// AccountNie:
	AccountNie []string `json:"account__nie,omitempty" mapstructure:"account__nie,omitempty"`
	// AccountNiew:
	AccountNiew []string `json:"account__niew,omitempty" mapstructure:"account__niew,omitempty"`
	// AccountNisw:
	AccountNisw []string `json:"account__nisw,omitempty" mapstructure:"account__nisw,omitempty"`
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
	// Provider: Provider (slug)
	Provider []string `json:"provider,omitempty" mapstructure:"provider,omitempty"`
	// ProviderN: Provider (slug)
	ProviderN []string `json:"provider__n,omitempty" mapstructure:"provider__n,omitempty"`
	// ProviderId: Provider (ID)
	ProviderId []int32 `json:"provider_id,omitempty" mapstructure:"provider_id,omitempty"`
	// ProviderIdN: Provider (ID)
	ProviderIdN []int32 `json:"provider_id__n,omitempty" mapstructure:"provider_id__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitsProviderAccountsListQueryParameters) Validate() error {
	return validation.Errors{
		"account": validation.Validate(
			m.Account,
		),
		"accountEmpty": validation.Validate(
			m.AccountEmpty,
		),
		"accountIc": validation.Validate(
			m.AccountIc,
		),
		"accountIe": validation.Validate(
			m.AccountIe,
		),
		"accountIew": validation.Validate(
			m.AccountIew,
		),
		"accountIsw": validation.Validate(
			m.AccountIsw,
		),
		"accountN": validation.Validate(
			m.AccountN,
		),
		"accountNic": validation.Validate(
			m.AccountNic,
		),
		"accountNie": validation.Validate(
			m.AccountNie,
		),
		"accountNiew": validation.Validate(
			m.AccountNiew,
		),
		"accountNisw": validation.Validate(
			m.AccountNisw,
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
		"provider": validation.Validate(
			m.Provider,
		),
		"providerN": validation.Validate(
			m.ProviderN,
		),
		"providerId": validation.Validate(
			m.ProviderId,
		),
		"providerIdN": validation.Validate(
			m.ProviderIdN,
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

// GetAccount returns the Account property
func (m CircuitsProviderAccountsListQueryParameters) GetAccount() []string {
	return m.Account
}

// SetAccount sets the Account property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccount(val []string) {
	m.Account = val
}

// GetAccountEmpty returns the AccountEmpty property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountEmpty() []string {
	return m.AccountEmpty
}

// SetAccountEmpty sets the AccountEmpty property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountEmpty(val []string) {
	m.AccountEmpty = val
}

// GetAccountIc returns the AccountIc property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountIc() []string {
	return m.AccountIc
}

// SetAccountIc sets the AccountIc property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountIc(val []string) {
	m.AccountIc = val
}

// GetAccountIe returns the AccountIe property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountIe() []string {
	return m.AccountIe
}

// SetAccountIe sets the AccountIe property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountIe(val []string) {
	m.AccountIe = val
}

// GetAccountIew returns the AccountIew property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountIew() []string {
	return m.AccountIew
}

// SetAccountIew sets the AccountIew property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountIew(val []string) {
	m.AccountIew = val
}

// GetAccountIsw returns the AccountIsw property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountIsw() []string {
	return m.AccountIsw
}

// SetAccountIsw sets the AccountIsw property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountIsw(val []string) {
	m.AccountIsw = val
}

// GetAccountN returns the AccountN property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountN() []string {
	return m.AccountN
}

// SetAccountN sets the AccountN property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountN(val []string) {
	m.AccountN = val
}

// GetAccountNic returns the AccountNic property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountNic() []string {
	return m.AccountNic
}

// SetAccountNic sets the AccountNic property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountNic(val []string) {
	m.AccountNic = val
}

// GetAccountNie returns the AccountNie property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountNie() []string {
	return m.AccountNie
}

// SetAccountNie sets the AccountNie property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountNie(val []string) {
	m.AccountNie = val
}

// GetAccountNiew returns the AccountNiew property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountNiew() []string {
	return m.AccountNiew
}

// SetAccountNiew sets the AccountNiew property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountNiew(val []string) {
	m.AccountNiew = val
}

// GetAccountNisw returns the AccountNisw property
func (m CircuitsProviderAccountsListQueryParameters) GetAccountNisw() []string {
	return m.AccountNisw
}

// SetAccountNisw sets the AccountNisw property
func (m *CircuitsProviderAccountsListQueryParameters) SetAccountNisw(val []string) {
	m.AccountNisw = val
}

// GetCreated returns the Created property
func (m CircuitsProviderAccountsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CircuitsProviderAccountsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m CircuitsProviderAccountsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *CircuitsProviderAccountsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m CircuitsProviderAccountsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *CircuitsProviderAccountsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m CircuitsProviderAccountsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *CircuitsProviderAccountsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m CircuitsProviderAccountsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *CircuitsProviderAccountsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m CircuitsProviderAccountsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *CircuitsProviderAccountsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m CircuitsProviderAccountsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *CircuitsProviderAccountsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m CircuitsProviderAccountsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m CircuitsProviderAccountsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *CircuitsProviderAccountsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m CircuitsProviderAccountsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitsProviderAccountsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m CircuitsProviderAccountsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *CircuitsProviderAccountsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m CircuitsProviderAccountsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *CircuitsProviderAccountsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m CircuitsProviderAccountsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *CircuitsProviderAccountsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m CircuitsProviderAccountsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *CircuitsProviderAccountsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m CircuitsProviderAccountsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *CircuitsProviderAccountsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m CircuitsProviderAccountsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CircuitsProviderAccountsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m CircuitsProviderAccountsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *CircuitsProviderAccountsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m CircuitsProviderAccountsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *CircuitsProviderAccountsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m CircuitsProviderAccountsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *CircuitsProviderAccountsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m CircuitsProviderAccountsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *CircuitsProviderAccountsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m CircuitsProviderAccountsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *CircuitsProviderAccountsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m CircuitsProviderAccountsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *CircuitsProviderAccountsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m CircuitsProviderAccountsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *CircuitsProviderAccountsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m CircuitsProviderAccountsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m CircuitsProviderAccountsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m CircuitsProviderAccountsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m CircuitsProviderAccountsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m CircuitsProviderAccountsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m CircuitsProviderAccountsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m CircuitsProviderAccountsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m CircuitsProviderAccountsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m CircuitsProviderAccountsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m CircuitsProviderAccountsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *CircuitsProviderAccountsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m CircuitsProviderAccountsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *CircuitsProviderAccountsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m CircuitsProviderAccountsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *CircuitsProviderAccountsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetProvider returns the Provider property
func (m CircuitsProviderAccountsListQueryParameters) GetProvider() []string {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *CircuitsProviderAccountsListQueryParameters) SetProvider(val []string) {
	m.Provider = val
}

// GetProviderN returns the ProviderN property
func (m CircuitsProviderAccountsListQueryParameters) GetProviderN() []string {
	return m.ProviderN
}

// SetProviderN sets the ProviderN property
func (m *CircuitsProviderAccountsListQueryParameters) SetProviderN(val []string) {
	m.ProviderN = val
}

// GetProviderId returns the ProviderId property
func (m CircuitsProviderAccountsListQueryParameters) GetProviderId() []int32 {
	return m.ProviderId
}

// SetProviderId sets the ProviderId property
func (m *CircuitsProviderAccountsListQueryParameters) SetProviderId(val []int32) {
	m.ProviderId = val
}

// GetProviderIdN returns the ProviderIdN property
func (m CircuitsProviderAccountsListQueryParameters) GetProviderIdN() []int32 {
	return m.ProviderIdN
}

// SetProviderIdN sets the ProviderIdN property
func (m *CircuitsProviderAccountsListQueryParameters) SetProviderIdN(val []int32) {
	m.ProviderIdN = val
}

// GetQ returns the Q property
func (m CircuitsProviderAccountsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *CircuitsProviderAccountsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m CircuitsProviderAccountsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *CircuitsProviderAccountsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m CircuitsProviderAccountsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *CircuitsProviderAccountsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m CircuitsProviderAccountsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *CircuitsProviderAccountsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
