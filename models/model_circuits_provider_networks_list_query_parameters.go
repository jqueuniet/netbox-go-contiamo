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

// CircuitsProviderNetworksListQueryParameters is an object.
type CircuitsProviderNetworksListQueryParameters struct {
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
	// ServiceId:
	ServiceId []string `json:"service_id,omitempty" mapstructure:"service_id,omitempty"`
	// ServiceIdEmpty:
	ServiceIdEmpty []string `json:"service_id__empty,omitempty" mapstructure:"service_id__empty,omitempty"`
	// ServiceIdIc:
	ServiceIdIc []string `json:"service_id__ic,omitempty" mapstructure:"service_id__ic,omitempty"`
	// ServiceIdIe:
	ServiceIdIe []string `json:"service_id__ie,omitempty" mapstructure:"service_id__ie,omitempty"`
	// ServiceIdIew:
	ServiceIdIew []string `json:"service_id__iew,omitempty" mapstructure:"service_id__iew,omitempty"`
	// ServiceIdIsw:
	ServiceIdIsw []string `json:"service_id__isw,omitempty" mapstructure:"service_id__isw,omitempty"`
	// ServiceIdN:
	ServiceIdN []string `json:"service_id__n,omitempty" mapstructure:"service_id__n,omitempty"`
	// ServiceIdNic:
	ServiceIdNic []string `json:"service_id__nic,omitempty" mapstructure:"service_id__nic,omitempty"`
	// ServiceIdNie:
	ServiceIdNie []string `json:"service_id__nie,omitempty" mapstructure:"service_id__nie,omitempty"`
	// ServiceIdNiew:
	ServiceIdNiew []string `json:"service_id__niew,omitempty" mapstructure:"service_id__niew,omitempty"`
	// ServiceIdNisw:
	ServiceIdNisw []string `json:"service_id__nisw,omitempty" mapstructure:"service_id__nisw,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitsProviderNetworksListQueryParameters) Validate() error {
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
		"serviceId": validation.Validate(
			m.ServiceId,
		),
		"serviceIdEmpty": validation.Validate(
			m.ServiceIdEmpty,
		),
		"serviceIdIc": validation.Validate(
			m.ServiceIdIc,
		),
		"serviceIdIe": validation.Validate(
			m.ServiceIdIe,
		),
		"serviceIdIew": validation.Validate(
			m.ServiceIdIew,
		),
		"serviceIdIsw": validation.Validate(
			m.ServiceIdIsw,
		),
		"serviceIdN": validation.Validate(
			m.ServiceIdN,
		),
		"serviceIdNic": validation.Validate(
			m.ServiceIdNic,
		),
		"serviceIdNie": validation.Validate(
			m.ServiceIdNie,
		),
		"serviceIdNiew": validation.Validate(
			m.ServiceIdNiew,
		),
		"serviceIdNisw": validation.Validate(
			m.ServiceIdNisw,
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
func (m CircuitsProviderNetworksListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CircuitsProviderNetworksListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m CircuitsProviderNetworksListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *CircuitsProviderNetworksListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m CircuitsProviderNetworksListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *CircuitsProviderNetworksListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m CircuitsProviderNetworksListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *CircuitsProviderNetworksListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m CircuitsProviderNetworksListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *CircuitsProviderNetworksListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m CircuitsProviderNetworksListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *CircuitsProviderNetworksListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m CircuitsProviderNetworksListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *CircuitsProviderNetworksListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m CircuitsProviderNetworksListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m CircuitsProviderNetworksListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *CircuitsProviderNetworksListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m CircuitsProviderNetworksListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitsProviderNetworksListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m CircuitsProviderNetworksListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *CircuitsProviderNetworksListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m CircuitsProviderNetworksListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *CircuitsProviderNetworksListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m CircuitsProviderNetworksListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *CircuitsProviderNetworksListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m CircuitsProviderNetworksListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *CircuitsProviderNetworksListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m CircuitsProviderNetworksListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *CircuitsProviderNetworksListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m CircuitsProviderNetworksListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CircuitsProviderNetworksListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m CircuitsProviderNetworksListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *CircuitsProviderNetworksListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m CircuitsProviderNetworksListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *CircuitsProviderNetworksListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m CircuitsProviderNetworksListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *CircuitsProviderNetworksListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m CircuitsProviderNetworksListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *CircuitsProviderNetworksListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m CircuitsProviderNetworksListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *CircuitsProviderNetworksListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m CircuitsProviderNetworksListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *CircuitsProviderNetworksListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m CircuitsProviderNetworksListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *CircuitsProviderNetworksListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m CircuitsProviderNetworksListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m CircuitsProviderNetworksListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m CircuitsProviderNetworksListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m CircuitsProviderNetworksListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m CircuitsProviderNetworksListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m CircuitsProviderNetworksListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m CircuitsProviderNetworksListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m CircuitsProviderNetworksListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m CircuitsProviderNetworksListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m CircuitsProviderNetworksListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *CircuitsProviderNetworksListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m CircuitsProviderNetworksListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *CircuitsProviderNetworksListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m CircuitsProviderNetworksListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *CircuitsProviderNetworksListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetProvider returns the Provider property
func (m CircuitsProviderNetworksListQueryParameters) GetProvider() []string {
	return m.Provider
}

// SetProvider sets the Provider property
func (m *CircuitsProviderNetworksListQueryParameters) SetProvider(val []string) {
	m.Provider = val
}

// GetProviderN returns the ProviderN property
func (m CircuitsProviderNetworksListQueryParameters) GetProviderN() []string {
	return m.ProviderN
}

// SetProviderN sets the ProviderN property
func (m *CircuitsProviderNetworksListQueryParameters) SetProviderN(val []string) {
	m.ProviderN = val
}

// GetProviderId returns the ProviderId property
func (m CircuitsProviderNetworksListQueryParameters) GetProviderId() []int32 {
	return m.ProviderId
}

// SetProviderId sets the ProviderId property
func (m *CircuitsProviderNetworksListQueryParameters) SetProviderId(val []int32) {
	m.ProviderId = val
}

// GetProviderIdN returns the ProviderIdN property
func (m CircuitsProviderNetworksListQueryParameters) GetProviderIdN() []int32 {
	return m.ProviderIdN
}

// SetProviderIdN sets the ProviderIdN property
func (m *CircuitsProviderNetworksListQueryParameters) SetProviderIdN(val []int32) {
	m.ProviderIdN = val
}

// GetQ returns the Q property
func (m CircuitsProviderNetworksListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *CircuitsProviderNetworksListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetServiceId returns the ServiceId property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceId() []string {
	return m.ServiceId
}

// SetServiceId sets the ServiceId property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceId(val []string) {
	m.ServiceId = val
}

// GetServiceIdEmpty returns the ServiceIdEmpty property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdEmpty() []string {
	return m.ServiceIdEmpty
}

// SetServiceIdEmpty sets the ServiceIdEmpty property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdEmpty(val []string) {
	m.ServiceIdEmpty = val
}

// GetServiceIdIc returns the ServiceIdIc property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdIc() []string {
	return m.ServiceIdIc
}

// SetServiceIdIc sets the ServiceIdIc property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdIc(val []string) {
	m.ServiceIdIc = val
}

// GetServiceIdIe returns the ServiceIdIe property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdIe() []string {
	return m.ServiceIdIe
}

// SetServiceIdIe sets the ServiceIdIe property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdIe(val []string) {
	m.ServiceIdIe = val
}

// GetServiceIdIew returns the ServiceIdIew property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdIew() []string {
	return m.ServiceIdIew
}

// SetServiceIdIew sets the ServiceIdIew property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdIew(val []string) {
	m.ServiceIdIew = val
}

// GetServiceIdIsw returns the ServiceIdIsw property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdIsw() []string {
	return m.ServiceIdIsw
}

// SetServiceIdIsw sets the ServiceIdIsw property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdIsw(val []string) {
	m.ServiceIdIsw = val
}

// GetServiceIdN returns the ServiceIdN property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdN() []string {
	return m.ServiceIdN
}

// SetServiceIdN sets the ServiceIdN property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdN(val []string) {
	m.ServiceIdN = val
}

// GetServiceIdNic returns the ServiceIdNic property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdNic() []string {
	return m.ServiceIdNic
}

// SetServiceIdNic sets the ServiceIdNic property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdNic(val []string) {
	m.ServiceIdNic = val
}

// GetServiceIdNie returns the ServiceIdNie property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdNie() []string {
	return m.ServiceIdNie
}

// SetServiceIdNie sets the ServiceIdNie property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdNie(val []string) {
	m.ServiceIdNie = val
}

// GetServiceIdNiew returns the ServiceIdNiew property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdNiew() []string {
	return m.ServiceIdNiew
}

// SetServiceIdNiew sets the ServiceIdNiew property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdNiew(val []string) {
	m.ServiceIdNiew = val
}

// GetServiceIdNisw returns the ServiceIdNisw property
func (m CircuitsProviderNetworksListQueryParameters) GetServiceIdNisw() []string {
	return m.ServiceIdNisw
}

// SetServiceIdNisw sets the ServiceIdNisw property
func (m *CircuitsProviderNetworksListQueryParameters) SetServiceIdNisw(val []string) {
	m.ServiceIdNisw = val
}

// GetTag returns the Tag property
func (m CircuitsProviderNetworksListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *CircuitsProviderNetworksListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m CircuitsProviderNetworksListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *CircuitsProviderNetworksListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m CircuitsProviderNetworksListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *CircuitsProviderNetworksListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
