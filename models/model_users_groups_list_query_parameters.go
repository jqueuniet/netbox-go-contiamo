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

// UsersGroupsListQueryParameters is an object.
type UsersGroupsListQueryParameters struct {
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
}

// Validate implements basic validation for this model
func (m UsersGroupsListQueryParameters) Validate() error {
	return validation.Errors{
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
	}.Filter()
}

// GetId returns the Id property
func (m UsersGroupsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *UsersGroupsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m UsersGroupsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *UsersGroupsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m UsersGroupsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *UsersGroupsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m UsersGroupsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *UsersGroupsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m UsersGroupsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *UsersGroupsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m UsersGroupsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *UsersGroupsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m UsersGroupsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *UsersGroupsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m UsersGroupsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *UsersGroupsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m UsersGroupsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *UsersGroupsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m UsersGroupsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *UsersGroupsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m UsersGroupsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *UsersGroupsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m UsersGroupsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *UsersGroupsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m UsersGroupsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *UsersGroupsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m UsersGroupsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *UsersGroupsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m UsersGroupsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *UsersGroupsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m UsersGroupsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *UsersGroupsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m UsersGroupsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *UsersGroupsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m UsersGroupsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *UsersGroupsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m UsersGroupsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *UsersGroupsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m UsersGroupsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *UsersGroupsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m UsersGroupsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *UsersGroupsListQueryParameters) SetQ(val string) {
	m.Q = val
}
