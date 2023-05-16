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

// CircuitsCircuitTerminationsListQueryParameters is an object.
type CircuitsCircuitTerminationsListQueryParameters struct {
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end,omitempty" mapstructure:"cable_end,omitempty"`
	// CableEndN: * `A` - A
	// * `B` - B
	CableEndN string `json:"cable_end__n,omitempty" mapstructure:"cable_end__n,omitempty"`
	// Cabled:
	Cabled bool `json:"cabled,omitempty" mapstructure:"cabled,omitempty"`
	// CircuitId: Circuit
	CircuitId []int32 `json:"circuit_id,omitempty" mapstructure:"circuit_id,omitempty"`
	// CircuitIdN: Circuit
	CircuitIdN []int32 `json:"circuit_id__n,omitempty" mapstructure:"circuit_id__n,omitempty"`
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
	// Occupied:
	Occupied bool `json:"occupied,omitempty" mapstructure:"occupied,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// PortSpeed:
	PortSpeed []int32 `json:"port_speed,omitempty" mapstructure:"port_speed,omitempty"`
	// PortSpeedGt:
	PortSpeedGt []int32 `json:"port_speed__gt,omitempty" mapstructure:"port_speed__gt,omitempty"`
	// PortSpeedGte:
	PortSpeedGte []int32 `json:"port_speed__gte,omitempty" mapstructure:"port_speed__gte,omitempty"`
	// PortSpeedLt:
	PortSpeedLt []int32 `json:"port_speed__lt,omitempty" mapstructure:"port_speed__lt,omitempty"`
	// PortSpeedLte:
	PortSpeedLte []int32 `json:"port_speed__lte,omitempty" mapstructure:"port_speed__lte,omitempty"`
	// PortSpeedN:
	PortSpeedN []int32 `json:"port_speed__n,omitempty" mapstructure:"port_speed__n,omitempty"`
	// ProviderNetworkId: ProviderNetwork (ID)
	ProviderNetworkId []*int32 `json:"provider_network_id,omitempty" mapstructure:"provider_network_id,omitempty"`
	// ProviderNetworkIdN: ProviderNetwork (ID)
	ProviderNetworkIdN []*int32 `json:"provider_network_id__n,omitempty" mapstructure:"provider_network_id__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Site: Site (slug)
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteN: Site (slug)
	SiteN []string `json:"site__n,omitempty" mapstructure:"site__n,omitempty"`
	// SiteId: Site (ID)
	SiteId []*int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
	// SiteIdN: Site (ID)
	SiteIdN []*int32 `json:"site_id__n,omitempty" mapstructure:"site_id__n,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// TermSide: * `A` - A
	// * `Z` - Z
	TermSide string `json:"term_side,omitempty" mapstructure:"term_side,omitempty"`
	// TermSideN: * `A` - A
	// * `Z` - Z
	TermSideN string `json:"term_side__n,omitempty" mapstructure:"term_side__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// UpstreamSpeed:
	UpstreamSpeed []int32 `json:"upstream_speed,omitempty" mapstructure:"upstream_speed,omitempty"`
	// UpstreamSpeedGt:
	UpstreamSpeedGt []int32 `json:"upstream_speed__gt,omitempty" mapstructure:"upstream_speed__gt,omitempty"`
	// UpstreamSpeedGte:
	UpstreamSpeedGte []int32 `json:"upstream_speed__gte,omitempty" mapstructure:"upstream_speed__gte,omitempty"`
	// UpstreamSpeedLt:
	UpstreamSpeedLt []int32 `json:"upstream_speed__lt,omitempty" mapstructure:"upstream_speed__lt,omitempty"`
	// UpstreamSpeedLte:
	UpstreamSpeedLte []int32 `json:"upstream_speed__lte,omitempty" mapstructure:"upstream_speed__lte,omitempty"`
	// UpstreamSpeedN:
	UpstreamSpeedN []int32 `json:"upstream_speed__n,omitempty" mapstructure:"upstream_speed__n,omitempty"`
	// XconnectId:
	XconnectId []string `json:"xconnect_id,omitempty" mapstructure:"xconnect_id,omitempty"`
	// XconnectIdEmpty:
	XconnectIdEmpty []string `json:"xconnect_id__empty,omitempty" mapstructure:"xconnect_id__empty,omitempty"`
	// XconnectIdIc:
	XconnectIdIc []string `json:"xconnect_id__ic,omitempty" mapstructure:"xconnect_id__ic,omitempty"`
	// XconnectIdIe:
	XconnectIdIe []string `json:"xconnect_id__ie,omitempty" mapstructure:"xconnect_id__ie,omitempty"`
	// XconnectIdIew:
	XconnectIdIew []string `json:"xconnect_id__iew,omitempty" mapstructure:"xconnect_id__iew,omitempty"`
	// XconnectIdIsw:
	XconnectIdIsw []string `json:"xconnect_id__isw,omitempty" mapstructure:"xconnect_id__isw,omitempty"`
	// XconnectIdN:
	XconnectIdN []string `json:"xconnect_id__n,omitempty" mapstructure:"xconnect_id__n,omitempty"`
	// XconnectIdNic:
	XconnectIdNic []string `json:"xconnect_id__nic,omitempty" mapstructure:"xconnect_id__nic,omitempty"`
	// XconnectIdNie:
	XconnectIdNie []string `json:"xconnect_id__nie,omitempty" mapstructure:"xconnect_id__nie,omitempty"`
	// XconnectIdNiew:
	XconnectIdNiew []string `json:"xconnect_id__niew,omitempty" mapstructure:"xconnect_id__niew,omitempty"`
	// XconnectIdNisw:
	XconnectIdNisw []string `json:"xconnect_id__nisw,omitempty" mapstructure:"xconnect_id__nisw,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitsCircuitTerminationsListQueryParameters) Validate() error {
	return validation.Errors{
		"circuitId": validation.Validate(
			m.CircuitId,
		),
		"circuitIdN": validation.Validate(
			m.CircuitIdN,
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
		"portSpeed": validation.Validate(
			m.PortSpeed,
		),
		"portSpeedGt": validation.Validate(
			m.PortSpeedGt,
		),
		"portSpeedGte": validation.Validate(
			m.PortSpeedGte,
		),
		"portSpeedLt": validation.Validate(
			m.PortSpeedLt,
		),
		"portSpeedLte": validation.Validate(
			m.PortSpeedLte,
		),
		"portSpeedN": validation.Validate(
			m.PortSpeedN,
		),
		"providerNetworkId": validation.Validate(
			m.ProviderNetworkId,
		),
		"providerNetworkIdN": validation.Validate(
			m.ProviderNetworkIdN,
		),
		"site": validation.Validate(
			m.Site,
		),
		"siteN": validation.Validate(
			m.SiteN,
		),
		"siteId": validation.Validate(
			m.SiteId,
		),
		"siteIdN": validation.Validate(
			m.SiteIdN,
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
		"upstreamSpeed": validation.Validate(
			m.UpstreamSpeed,
		),
		"upstreamSpeedGt": validation.Validate(
			m.UpstreamSpeedGt,
		),
		"upstreamSpeedGte": validation.Validate(
			m.UpstreamSpeedGte,
		),
		"upstreamSpeedLt": validation.Validate(
			m.UpstreamSpeedLt,
		),
		"upstreamSpeedLte": validation.Validate(
			m.UpstreamSpeedLte,
		),
		"upstreamSpeedN": validation.Validate(
			m.UpstreamSpeedN,
		),
		"xconnectId": validation.Validate(
			m.XconnectId,
		),
		"xconnectIdEmpty": validation.Validate(
			m.XconnectIdEmpty,
		),
		"xconnectIdIc": validation.Validate(
			m.XconnectIdIc,
		),
		"xconnectIdIe": validation.Validate(
			m.XconnectIdIe,
		),
		"xconnectIdIew": validation.Validate(
			m.XconnectIdIew,
		),
		"xconnectIdIsw": validation.Validate(
			m.XconnectIdIsw,
		),
		"xconnectIdN": validation.Validate(
			m.XconnectIdN,
		),
		"xconnectIdNic": validation.Validate(
			m.XconnectIdNic,
		),
		"xconnectIdNie": validation.Validate(
			m.XconnectIdNie,
		),
		"xconnectIdNiew": validation.Validate(
			m.XconnectIdNiew,
		),
		"xconnectIdNisw": validation.Validate(
			m.XconnectIdNisw,
		),
	}.Filter()
}

// GetCableEnd returns the CableEnd property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCableEndN returns the CableEndN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCableEndN() string {
	return m.CableEndN
}

// SetCableEndN sets the CableEndN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCableEndN(val string) {
	m.CableEndN = val
}

// GetCabled returns the Cabled property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCabled() bool {
	return m.Cabled
}

// SetCabled sets the Cabled property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCabled(val bool) {
	m.Cabled = val
}

// GetCircuitId returns the CircuitId property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCircuitId() []int32 {
	return m.CircuitId
}

// SetCircuitId sets the CircuitId property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCircuitId(val []int32) {
	m.CircuitId = val
}

// GetCircuitIdN returns the CircuitIdN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCircuitIdN() []int32 {
	return m.CircuitIdN
}

// SetCircuitIdN sets the CircuitIdN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCircuitIdN(val []int32) {
	m.CircuitIdN = val
}

// GetCreated returns the Created property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m CircuitsCircuitTerminationsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m CircuitsCircuitTerminationsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m CircuitsCircuitTerminationsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m CircuitsCircuitTerminationsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m CircuitsCircuitTerminationsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOccupied returns the Occupied property
func (m CircuitsCircuitTerminationsListQueryParameters) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetOccupied(val bool) {
	m.Occupied = val
}

// GetOffset returns the Offset property
func (m CircuitsCircuitTerminationsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m CircuitsCircuitTerminationsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPortSpeed returns the PortSpeed property
func (m CircuitsCircuitTerminationsListQueryParameters) GetPortSpeed() []int32 {
	return m.PortSpeed
}

// SetPortSpeed sets the PortSpeed property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetPortSpeed(val []int32) {
	m.PortSpeed = val
}

// GetPortSpeedGt returns the PortSpeedGt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetPortSpeedGt() []int32 {
	return m.PortSpeedGt
}

// SetPortSpeedGt sets the PortSpeedGt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetPortSpeedGt(val []int32) {
	m.PortSpeedGt = val
}

// GetPortSpeedGte returns the PortSpeedGte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetPortSpeedGte() []int32 {
	return m.PortSpeedGte
}

// SetPortSpeedGte sets the PortSpeedGte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetPortSpeedGte(val []int32) {
	m.PortSpeedGte = val
}

// GetPortSpeedLt returns the PortSpeedLt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetPortSpeedLt() []int32 {
	return m.PortSpeedLt
}

// SetPortSpeedLt sets the PortSpeedLt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetPortSpeedLt(val []int32) {
	m.PortSpeedLt = val
}

// GetPortSpeedLte returns the PortSpeedLte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetPortSpeedLte() []int32 {
	return m.PortSpeedLte
}

// SetPortSpeedLte sets the PortSpeedLte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetPortSpeedLte(val []int32) {
	m.PortSpeedLte = val
}

// GetPortSpeedN returns the PortSpeedN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetPortSpeedN() []int32 {
	return m.PortSpeedN
}

// SetPortSpeedN sets the PortSpeedN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetPortSpeedN(val []int32) {
	m.PortSpeedN = val
}

// GetProviderNetworkId returns the ProviderNetworkId property
func (m CircuitsCircuitTerminationsListQueryParameters) GetProviderNetworkId() []*int32 {
	return m.ProviderNetworkId
}

// SetProviderNetworkId sets the ProviderNetworkId property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetProviderNetworkId(val []*int32) {
	m.ProviderNetworkId = val
}

// GetProviderNetworkIdN returns the ProviderNetworkIdN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetProviderNetworkIdN() []*int32 {
	return m.ProviderNetworkIdN
}

// SetProviderNetworkIdN sets the ProviderNetworkIdN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetProviderNetworkIdN(val []*int32) {
	m.ProviderNetworkIdN = val
}

// GetQ returns the Q property
func (m CircuitsCircuitTerminationsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSite returns the Site property
func (m CircuitsCircuitTerminationsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteId returns the SiteId property
func (m CircuitsCircuitTerminationsListQueryParameters) GetSiteId() []*int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetSiteId(val []*int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetSiteIdN() []*int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetSiteIdN(val []*int32) {
	m.SiteIdN = val
}

// GetTag returns the Tag property
func (m CircuitsCircuitTerminationsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTermSide returns the TermSide property
func (m CircuitsCircuitTerminationsListQueryParameters) GetTermSide() string {
	return m.TermSide
}

// SetTermSide sets the TermSide property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetTermSide(val string) {
	m.TermSide = val
}

// GetTermSideN returns the TermSideN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetTermSideN() string {
	return m.TermSideN
}

// SetTermSideN sets the TermSideN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetTermSideN(val string) {
	m.TermSideN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m CircuitsCircuitTerminationsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetUpstreamSpeed returns the UpstreamSpeed property
func (m CircuitsCircuitTerminationsListQueryParameters) GetUpstreamSpeed() []int32 {
	return m.UpstreamSpeed
}

// SetUpstreamSpeed sets the UpstreamSpeed property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetUpstreamSpeed(val []int32) {
	m.UpstreamSpeed = val
}

// GetUpstreamSpeedGt returns the UpstreamSpeedGt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetUpstreamSpeedGt() []int32 {
	return m.UpstreamSpeedGt
}

// SetUpstreamSpeedGt sets the UpstreamSpeedGt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetUpstreamSpeedGt(val []int32) {
	m.UpstreamSpeedGt = val
}

// GetUpstreamSpeedGte returns the UpstreamSpeedGte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetUpstreamSpeedGte() []int32 {
	return m.UpstreamSpeedGte
}

// SetUpstreamSpeedGte sets the UpstreamSpeedGte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetUpstreamSpeedGte(val []int32) {
	m.UpstreamSpeedGte = val
}

// GetUpstreamSpeedLt returns the UpstreamSpeedLt property
func (m CircuitsCircuitTerminationsListQueryParameters) GetUpstreamSpeedLt() []int32 {
	return m.UpstreamSpeedLt
}

// SetUpstreamSpeedLt sets the UpstreamSpeedLt property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetUpstreamSpeedLt(val []int32) {
	m.UpstreamSpeedLt = val
}

// GetUpstreamSpeedLte returns the UpstreamSpeedLte property
func (m CircuitsCircuitTerminationsListQueryParameters) GetUpstreamSpeedLte() []int32 {
	return m.UpstreamSpeedLte
}

// SetUpstreamSpeedLte sets the UpstreamSpeedLte property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetUpstreamSpeedLte(val []int32) {
	m.UpstreamSpeedLte = val
}

// GetUpstreamSpeedN returns the UpstreamSpeedN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetUpstreamSpeedN() []int32 {
	return m.UpstreamSpeedN
}

// SetUpstreamSpeedN sets the UpstreamSpeedN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetUpstreamSpeedN(val []int32) {
	m.UpstreamSpeedN = val
}

// GetXconnectId returns the XconnectId property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectId() []string {
	return m.XconnectId
}

// SetXconnectId sets the XconnectId property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectId(val []string) {
	m.XconnectId = val
}

// GetXconnectIdEmpty returns the XconnectIdEmpty property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdEmpty() []string {
	return m.XconnectIdEmpty
}

// SetXconnectIdEmpty sets the XconnectIdEmpty property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdEmpty(val []string) {
	m.XconnectIdEmpty = val
}

// GetXconnectIdIc returns the XconnectIdIc property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdIc() []string {
	return m.XconnectIdIc
}

// SetXconnectIdIc sets the XconnectIdIc property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdIc(val []string) {
	m.XconnectIdIc = val
}

// GetXconnectIdIe returns the XconnectIdIe property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdIe() []string {
	return m.XconnectIdIe
}

// SetXconnectIdIe sets the XconnectIdIe property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdIe(val []string) {
	m.XconnectIdIe = val
}

// GetXconnectIdIew returns the XconnectIdIew property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdIew() []string {
	return m.XconnectIdIew
}

// SetXconnectIdIew sets the XconnectIdIew property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdIew(val []string) {
	m.XconnectIdIew = val
}

// GetXconnectIdIsw returns the XconnectIdIsw property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdIsw() []string {
	return m.XconnectIdIsw
}

// SetXconnectIdIsw sets the XconnectIdIsw property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdIsw(val []string) {
	m.XconnectIdIsw = val
}

// GetXconnectIdN returns the XconnectIdN property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdN() []string {
	return m.XconnectIdN
}

// SetXconnectIdN sets the XconnectIdN property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdN(val []string) {
	m.XconnectIdN = val
}

// GetXconnectIdNic returns the XconnectIdNic property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdNic() []string {
	return m.XconnectIdNic
}

// SetXconnectIdNic sets the XconnectIdNic property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdNic(val []string) {
	m.XconnectIdNic = val
}

// GetXconnectIdNie returns the XconnectIdNie property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdNie() []string {
	return m.XconnectIdNie
}

// SetXconnectIdNie sets the XconnectIdNie property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdNie(val []string) {
	m.XconnectIdNie = val
}

// GetXconnectIdNiew returns the XconnectIdNiew property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdNiew() []string {
	return m.XconnectIdNiew
}

// SetXconnectIdNiew sets the XconnectIdNiew property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdNiew(val []string) {
	m.XconnectIdNiew = val
}

// GetXconnectIdNisw returns the XconnectIdNisw property
func (m CircuitsCircuitTerminationsListQueryParameters) GetXconnectIdNisw() []string {
	return m.XconnectIdNisw
}

// SetXconnectIdNisw sets the XconnectIdNisw property
func (m *CircuitsCircuitTerminationsListQueryParameters) SetXconnectIdNisw(val []string) {
	m.XconnectIdNisw = val
}
