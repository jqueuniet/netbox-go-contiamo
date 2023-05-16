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

// DcimCableTerminationsListQueryParameters is an object.
type DcimCableTerminationsListQueryParameters struct {
	// Cable:
	Cable int32 `json:"cable,omitempty" mapstructure:"cable,omitempty"`
	// CableN:
	CableN int32 `json:"cable__n,omitempty" mapstructure:"cable__n,omitempty"`
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end,omitempty" mapstructure:"cable_end,omitempty"`
	// CableEndN: * `A` - A
	// * `B` - B
	CableEndN string `json:"cable_end__n,omitempty" mapstructure:"cable_end__n,omitempty"`
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
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// TerminationId:
	TerminationId []int32 `json:"termination_id,omitempty" mapstructure:"termination_id,omitempty"`
	// TerminationIdGt:
	TerminationIdGt []int32 `json:"termination_id__gt,omitempty" mapstructure:"termination_id__gt,omitempty"`
	// TerminationIdGte:
	TerminationIdGte []int32 `json:"termination_id__gte,omitempty" mapstructure:"termination_id__gte,omitempty"`
	// TerminationIdLt:
	TerminationIdLt []int32 `json:"termination_id__lt,omitempty" mapstructure:"termination_id__lt,omitempty"`
	// TerminationIdLte:
	TerminationIdLte []int32 `json:"termination_id__lte,omitempty" mapstructure:"termination_id__lte,omitempty"`
	// TerminationIdN:
	TerminationIdN []int32 `json:"termination_id__n,omitempty" mapstructure:"termination_id__n,omitempty"`
	// TerminationType:
	TerminationType string `json:"termination_type,omitempty" mapstructure:"termination_type,omitempty"`
	// TerminationTypeN:
	TerminationTypeN string `json:"termination_type__n,omitempty" mapstructure:"termination_type__n,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimCableTerminationsListQueryParameters) Validate() error {
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
		"terminationId": validation.Validate(
			m.TerminationId,
		),
		"terminationIdGt": validation.Validate(
			m.TerminationIdGt,
		),
		"terminationIdGte": validation.Validate(
			m.TerminationIdGte,
		),
		"terminationIdLt": validation.Validate(
			m.TerminationIdLt,
		),
		"terminationIdLte": validation.Validate(
			m.TerminationIdLte,
		),
		"terminationIdN": validation.Validate(
			m.TerminationIdN,
		),
	}.Filter()
}

// GetCable returns the Cable property
func (m DcimCableTerminationsListQueryParameters) GetCable() int32 {
	return m.Cable
}

// SetCable sets the Cable property
func (m *DcimCableTerminationsListQueryParameters) SetCable(val int32) {
	m.Cable = val
}

// GetCableN returns the CableN property
func (m DcimCableTerminationsListQueryParameters) GetCableN() int32 {
	return m.CableN
}

// SetCableN sets the CableN property
func (m *DcimCableTerminationsListQueryParameters) SetCableN(val int32) {
	m.CableN = val
}

// GetCableEnd returns the CableEnd property
func (m DcimCableTerminationsListQueryParameters) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *DcimCableTerminationsListQueryParameters) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCableEndN returns the CableEndN property
func (m DcimCableTerminationsListQueryParameters) GetCableEndN() string {
	return m.CableEndN
}

// SetCableEndN sets the CableEndN property
func (m *DcimCableTerminationsListQueryParameters) SetCableEndN(val string) {
	m.CableEndN = val
}

// GetId returns the Id property
func (m DcimCableTerminationsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimCableTerminationsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimCableTerminationsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimCableTerminationsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimCableTerminationsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimCableTerminationsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimCableTerminationsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimCableTerminationsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimCableTerminationsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimCableTerminationsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimCableTerminationsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimCableTerminationsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m DcimCableTerminationsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimCableTerminationsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m DcimCableTerminationsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimCableTerminationsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimCableTerminationsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimCableTerminationsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetTerminationId returns the TerminationId property
func (m DcimCableTerminationsListQueryParameters) GetTerminationId() []int32 {
	return m.TerminationId
}

// SetTerminationId sets the TerminationId property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationId(val []int32) {
	m.TerminationId = val
}

// GetTerminationIdGt returns the TerminationIdGt property
func (m DcimCableTerminationsListQueryParameters) GetTerminationIdGt() []int32 {
	return m.TerminationIdGt
}

// SetTerminationIdGt sets the TerminationIdGt property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationIdGt(val []int32) {
	m.TerminationIdGt = val
}

// GetTerminationIdGte returns the TerminationIdGte property
func (m DcimCableTerminationsListQueryParameters) GetTerminationIdGte() []int32 {
	return m.TerminationIdGte
}

// SetTerminationIdGte sets the TerminationIdGte property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationIdGte(val []int32) {
	m.TerminationIdGte = val
}

// GetTerminationIdLt returns the TerminationIdLt property
func (m DcimCableTerminationsListQueryParameters) GetTerminationIdLt() []int32 {
	return m.TerminationIdLt
}

// SetTerminationIdLt sets the TerminationIdLt property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationIdLt(val []int32) {
	m.TerminationIdLt = val
}

// GetTerminationIdLte returns the TerminationIdLte property
func (m DcimCableTerminationsListQueryParameters) GetTerminationIdLte() []int32 {
	return m.TerminationIdLte
}

// SetTerminationIdLte sets the TerminationIdLte property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationIdLte(val []int32) {
	m.TerminationIdLte = val
}

// GetTerminationIdN returns the TerminationIdN property
func (m DcimCableTerminationsListQueryParameters) GetTerminationIdN() []int32 {
	return m.TerminationIdN
}

// SetTerminationIdN sets the TerminationIdN property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationIdN(val []int32) {
	m.TerminationIdN = val
}

// GetTerminationType returns the TerminationType property
func (m DcimCableTerminationsListQueryParameters) GetTerminationType() string {
	return m.TerminationType
}

// SetTerminationType sets the TerminationType property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationType(val string) {
	m.TerminationType = val
}

// GetTerminationTypeN returns the TerminationTypeN property
func (m DcimCableTerminationsListQueryParameters) GetTerminationTypeN() string {
	return m.TerminationTypeN
}

// SetTerminationTypeN sets the TerminationTypeN property
func (m *DcimCableTerminationsListQueryParameters) SetTerminationTypeN(val string) {
	m.TerminationTypeN = val
}
