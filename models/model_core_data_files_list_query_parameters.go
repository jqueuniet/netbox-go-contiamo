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

// CoreDataFilesListQueryParameters is an object.
type CoreDataFilesListQueryParameters struct {
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
	// Hash:
	Hash []string `json:"hash,omitempty" mapstructure:"hash,omitempty"`
	// HashEmpty:
	HashEmpty []string `json:"hash__empty,omitempty" mapstructure:"hash__empty,omitempty"`
	// HashIc:
	HashIc []string `json:"hash__ic,omitempty" mapstructure:"hash__ic,omitempty"`
	// HashIe:
	HashIe []string `json:"hash__ie,omitempty" mapstructure:"hash__ie,omitempty"`
	// HashIew:
	HashIew []string `json:"hash__iew,omitempty" mapstructure:"hash__iew,omitempty"`
	// HashIsw:
	HashIsw []string `json:"hash__isw,omitempty" mapstructure:"hash__isw,omitempty"`
	// HashN:
	HashN []string `json:"hash__n,omitempty" mapstructure:"hash__n,omitempty"`
	// HashNic:
	HashNic []string `json:"hash__nic,omitempty" mapstructure:"hash__nic,omitempty"`
	// HashNie:
	HashNie []string `json:"hash__nie,omitempty" mapstructure:"hash__nie,omitempty"`
	// HashNiew:
	HashNiew []string `json:"hash__niew,omitempty" mapstructure:"hash__niew,omitempty"`
	// HashNisw:
	HashNisw []string `json:"hash__nisw,omitempty" mapstructure:"hash__nisw,omitempty"`
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
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Path:
	Path []string `json:"path,omitempty" mapstructure:"path,omitempty"`
	// PathEmpty:
	PathEmpty []string `json:"path__empty,omitempty" mapstructure:"path__empty,omitempty"`
	// PathIc:
	PathIc []string `json:"path__ic,omitempty" mapstructure:"path__ic,omitempty"`
	// PathIe:
	PathIe []string `json:"path__ie,omitempty" mapstructure:"path__ie,omitempty"`
	// PathIew:
	PathIew []string `json:"path__iew,omitempty" mapstructure:"path__iew,omitempty"`
	// PathIsw:
	PathIsw []string `json:"path__isw,omitempty" mapstructure:"path__isw,omitempty"`
	// PathN:
	PathN []string `json:"path__n,omitempty" mapstructure:"path__n,omitempty"`
	// PathNic:
	PathNic []string `json:"path__nic,omitempty" mapstructure:"path__nic,omitempty"`
	// PathNie:
	PathNie []string `json:"path__nie,omitempty" mapstructure:"path__nie,omitempty"`
	// PathNiew:
	PathNiew []string `json:"path__niew,omitempty" mapstructure:"path__niew,omitempty"`
	// PathNisw:
	PathNisw []string `json:"path__nisw,omitempty" mapstructure:"path__nisw,omitempty"`
	// Q:
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Size:
	Size []int32 `json:"size,omitempty" mapstructure:"size,omitempty"`
	// SizeGt:
	SizeGt []int32 `json:"size__gt,omitempty" mapstructure:"size__gt,omitempty"`
	// SizeGte:
	SizeGte []int32 `json:"size__gte,omitempty" mapstructure:"size__gte,omitempty"`
	// SizeLt:
	SizeLt []int32 `json:"size__lt,omitempty" mapstructure:"size__lt,omitempty"`
	// SizeLte:
	SizeLte []int32 `json:"size__lte,omitempty" mapstructure:"size__lte,omitempty"`
	// SizeN:
	SizeN []int32 `json:"size__n,omitempty" mapstructure:"size__n,omitempty"`
	// Source: Data source (name)
	Source []string `json:"source,omitempty" mapstructure:"source,omitempty"`
	// SourceN: Data source (name)
	SourceN []string `json:"source__n,omitempty" mapstructure:"source__n,omitempty"`
	// SourceId: Data source (ID)
	SourceId []int32 `json:"source_id,omitempty" mapstructure:"source_id,omitempty"`
	// SourceIdN: Data source (ID)
	SourceIdN []int32 `json:"source_id__n,omitempty" mapstructure:"source_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m CoreDataFilesListQueryParameters) Validate() error {
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
		"hash": validation.Validate(
			m.Hash,
		),
		"hashEmpty": validation.Validate(
			m.HashEmpty,
		),
		"hashIc": validation.Validate(
			m.HashIc,
		),
		"hashIe": validation.Validate(
			m.HashIe,
		),
		"hashIew": validation.Validate(
			m.HashIew,
		),
		"hashIsw": validation.Validate(
			m.HashIsw,
		),
		"hashN": validation.Validate(
			m.HashN,
		),
		"hashNic": validation.Validate(
			m.HashNic,
		),
		"hashNie": validation.Validate(
			m.HashNie,
		),
		"hashNiew": validation.Validate(
			m.HashNiew,
		),
		"hashNisw": validation.Validate(
			m.HashNisw,
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
		"path": validation.Validate(
			m.Path,
		),
		"pathEmpty": validation.Validate(
			m.PathEmpty,
		),
		"pathIc": validation.Validate(
			m.PathIc,
		),
		"pathIe": validation.Validate(
			m.PathIe,
		),
		"pathIew": validation.Validate(
			m.PathIew,
		),
		"pathIsw": validation.Validate(
			m.PathIsw,
		),
		"pathN": validation.Validate(
			m.PathN,
		),
		"pathNic": validation.Validate(
			m.PathNic,
		),
		"pathNie": validation.Validate(
			m.PathNie,
		),
		"pathNiew": validation.Validate(
			m.PathNiew,
		),
		"pathNisw": validation.Validate(
			m.PathNisw,
		),
		"size": validation.Validate(
			m.Size,
		),
		"sizeGt": validation.Validate(
			m.SizeGt,
		),
		"sizeGte": validation.Validate(
			m.SizeGte,
		),
		"sizeLt": validation.Validate(
			m.SizeLt,
		),
		"sizeLte": validation.Validate(
			m.SizeLte,
		),
		"sizeN": validation.Validate(
			m.SizeN,
		),
		"source": validation.Validate(
			m.Source,
		),
		"sourceN": validation.Validate(
			m.SourceN,
		),
		"sourceId": validation.Validate(
			m.SourceId,
		),
		"sourceIdN": validation.Validate(
			m.SourceIdN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m CoreDataFilesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CoreDataFilesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m CoreDataFilesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *CoreDataFilesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m CoreDataFilesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *CoreDataFilesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m CoreDataFilesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *CoreDataFilesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m CoreDataFilesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *CoreDataFilesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m CoreDataFilesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *CoreDataFilesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m CoreDataFilesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *CoreDataFilesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetHash returns the Hash property
func (m CoreDataFilesListQueryParameters) GetHash() []string {
	return m.Hash
}

// SetHash sets the Hash property
func (m *CoreDataFilesListQueryParameters) SetHash(val []string) {
	m.Hash = val
}

// GetHashEmpty returns the HashEmpty property
func (m CoreDataFilesListQueryParameters) GetHashEmpty() []string {
	return m.HashEmpty
}

// SetHashEmpty sets the HashEmpty property
func (m *CoreDataFilesListQueryParameters) SetHashEmpty(val []string) {
	m.HashEmpty = val
}

// GetHashIc returns the HashIc property
func (m CoreDataFilesListQueryParameters) GetHashIc() []string {
	return m.HashIc
}

// SetHashIc sets the HashIc property
func (m *CoreDataFilesListQueryParameters) SetHashIc(val []string) {
	m.HashIc = val
}

// GetHashIe returns the HashIe property
func (m CoreDataFilesListQueryParameters) GetHashIe() []string {
	return m.HashIe
}

// SetHashIe sets the HashIe property
func (m *CoreDataFilesListQueryParameters) SetHashIe(val []string) {
	m.HashIe = val
}

// GetHashIew returns the HashIew property
func (m CoreDataFilesListQueryParameters) GetHashIew() []string {
	return m.HashIew
}

// SetHashIew sets the HashIew property
func (m *CoreDataFilesListQueryParameters) SetHashIew(val []string) {
	m.HashIew = val
}

// GetHashIsw returns the HashIsw property
func (m CoreDataFilesListQueryParameters) GetHashIsw() []string {
	return m.HashIsw
}

// SetHashIsw sets the HashIsw property
func (m *CoreDataFilesListQueryParameters) SetHashIsw(val []string) {
	m.HashIsw = val
}

// GetHashN returns the HashN property
func (m CoreDataFilesListQueryParameters) GetHashN() []string {
	return m.HashN
}

// SetHashN sets the HashN property
func (m *CoreDataFilesListQueryParameters) SetHashN(val []string) {
	m.HashN = val
}

// GetHashNic returns the HashNic property
func (m CoreDataFilesListQueryParameters) GetHashNic() []string {
	return m.HashNic
}

// SetHashNic sets the HashNic property
func (m *CoreDataFilesListQueryParameters) SetHashNic(val []string) {
	m.HashNic = val
}

// GetHashNie returns the HashNie property
func (m CoreDataFilesListQueryParameters) GetHashNie() []string {
	return m.HashNie
}

// SetHashNie sets the HashNie property
func (m *CoreDataFilesListQueryParameters) SetHashNie(val []string) {
	m.HashNie = val
}

// GetHashNiew returns the HashNiew property
func (m CoreDataFilesListQueryParameters) GetHashNiew() []string {
	return m.HashNiew
}

// SetHashNiew sets the HashNiew property
func (m *CoreDataFilesListQueryParameters) SetHashNiew(val []string) {
	m.HashNiew = val
}

// GetHashNisw returns the HashNisw property
func (m CoreDataFilesListQueryParameters) GetHashNisw() []string {
	return m.HashNisw
}

// SetHashNisw sets the HashNisw property
func (m *CoreDataFilesListQueryParameters) SetHashNisw(val []string) {
	m.HashNisw = val
}

// GetId returns the Id property
func (m CoreDataFilesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CoreDataFilesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m CoreDataFilesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *CoreDataFilesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m CoreDataFilesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *CoreDataFilesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m CoreDataFilesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *CoreDataFilesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m CoreDataFilesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *CoreDataFilesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m CoreDataFilesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *CoreDataFilesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m CoreDataFilesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CoreDataFilesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m CoreDataFilesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *CoreDataFilesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m CoreDataFilesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *CoreDataFilesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m CoreDataFilesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *CoreDataFilesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m CoreDataFilesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *CoreDataFilesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m CoreDataFilesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *CoreDataFilesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m CoreDataFilesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *CoreDataFilesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m CoreDataFilesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *CoreDataFilesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m CoreDataFilesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *CoreDataFilesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPath returns the Path property
func (m CoreDataFilesListQueryParameters) GetPath() []string {
	return m.Path
}

// SetPath sets the Path property
func (m *CoreDataFilesListQueryParameters) SetPath(val []string) {
	m.Path = val
}

// GetPathEmpty returns the PathEmpty property
func (m CoreDataFilesListQueryParameters) GetPathEmpty() []string {
	return m.PathEmpty
}

// SetPathEmpty sets the PathEmpty property
func (m *CoreDataFilesListQueryParameters) SetPathEmpty(val []string) {
	m.PathEmpty = val
}

// GetPathIc returns the PathIc property
func (m CoreDataFilesListQueryParameters) GetPathIc() []string {
	return m.PathIc
}

// SetPathIc sets the PathIc property
func (m *CoreDataFilesListQueryParameters) SetPathIc(val []string) {
	m.PathIc = val
}

// GetPathIe returns the PathIe property
func (m CoreDataFilesListQueryParameters) GetPathIe() []string {
	return m.PathIe
}

// SetPathIe sets the PathIe property
func (m *CoreDataFilesListQueryParameters) SetPathIe(val []string) {
	m.PathIe = val
}

// GetPathIew returns the PathIew property
func (m CoreDataFilesListQueryParameters) GetPathIew() []string {
	return m.PathIew
}

// SetPathIew sets the PathIew property
func (m *CoreDataFilesListQueryParameters) SetPathIew(val []string) {
	m.PathIew = val
}

// GetPathIsw returns the PathIsw property
func (m CoreDataFilesListQueryParameters) GetPathIsw() []string {
	return m.PathIsw
}

// SetPathIsw sets the PathIsw property
func (m *CoreDataFilesListQueryParameters) SetPathIsw(val []string) {
	m.PathIsw = val
}

// GetPathN returns the PathN property
func (m CoreDataFilesListQueryParameters) GetPathN() []string {
	return m.PathN
}

// SetPathN sets the PathN property
func (m *CoreDataFilesListQueryParameters) SetPathN(val []string) {
	m.PathN = val
}

// GetPathNic returns the PathNic property
func (m CoreDataFilesListQueryParameters) GetPathNic() []string {
	return m.PathNic
}

// SetPathNic sets the PathNic property
func (m *CoreDataFilesListQueryParameters) SetPathNic(val []string) {
	m.PathNic = val
}

// GetPathNie returns the PathNie property
func (m CoreDataFilesListQueryParameters) GetPathNie() []string {
	return m.PathNie
}

// SetPathNie sets the PathNie property
func (m *CoreDataFilesListQueryParameters) SetPathNie(val []string) {
	m.PathNie = val
}

// GetPathNiew returns the PathNiew property
func (m CoreDataFilesListQueryParameters) GetPathNiew() []string {
	return m.PathNiew
}

// SetPathNiew sets the PathNiew property
func (m *CoreDataFilesListQueryParameters) SetPathNiew(val []string) {
	m.PathNiew = val
}

// GetPathNisw returns the PathNisw property
func (m CoreDataFilesListQueryParameters) GetPathNisw() []string {
	return m.PathNisw
}

// SetPathNisw sets the PathNisw property
func (m *CoreDataFilesListQueryParameters) SetPathNisw(val []string) {
	m.PathNisw = val
}

// GetQ returns the Q property
func (m CoreDataFilesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *CoreDataFilesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSize returns the Size property
func (m CoreDataFilesListQueryParameters) GetSize() []int32 {
	return m.Size
}

// SetSize sets the Size property
func (m *CoreDataFilesListQueryParameters) SetSize(val []int32) {
	m.Size = val
}

// GetSizeGt returns the SizeGt property
func (m CoreDataFilesListQueryParameters) GetSizeGt() []int32 {
	return m.SizeGt
}

// SetSizeGt sets the SizeGt property
func (m *CoreDataFilesListQueryParameters) SetSizeGt(val []int32) {
	m.SizeGt = val
}

// GetSizeGte returns the SizeGte property
func (m CoreDataFilesListQueryParameters) GetSizeGte() []int32 {
	return m.SizeGte
}

// SetSizeGte sets the SizeGte property
func (m *CoreDataFilesListQueryParameters) SetSizeGte(val []int32) {
	m.SizeGte = val
}

// GetSizeLt returns the SizeLt property
func (m CoreDataFilesListQueryParameters) GetSizeLt() []int32 {
	return m.SizeLt
}

// SetSizeLt sets the SizeLt property
func (m *CoreDataFilesListQueryParameters) SetSizeLt(val []int32) {
	m.SizeLt = val
}

// GetSizeLte returns the SizeLte property
func (m CoreDataFilesListQueryParameters) GetSizeLte() []int32 {
	return m.SizeLte
}

// SetSizeLte sets the SizeLte property
func (m *CoreDataFilesListQueryParameters) SetSizeLte(val []int32) {
	m.SizeLte = val
}

// GetSizeN returns the SizeN property
func (m CoreDataFilesListQueryParameters) GetSizeN() []int32 {
	return m.SizeN
}

// SetSizeN sets the SizeN property
func (m *CoreDataFilesListQueryParameters) SetSizeN(val []int32) {
	m.SizeN = val
}

// GetSource returns the Source property
func (m CoreDataFilesListQueryParameters) GetSource() []string {
	return m.Source
}

// SetSource sets the Source property
func (m *CoreDataFilesListQueryParameters) SetSource(val []string) {
	m.Source = val
}

// GetSourceN returns the SourceN property
func (m CoreDataFilesListQueryParameters) GetSourceN() []string {
	return m.SourceN
}

// SetSourceN sets the SourceN property
func (m *CoreDataFilesListQueryParameters) SetSourceN(val []string) {
	m.SourceN = val
}

// GetSourceId returns the SourceId property
func (m CoreDataFilesListQueryParameters) GetSourceId() []int32 {
	return m.SourceId
}

// SetSourceId sets the SourceId property
func (m *CoreDataFilesListQueryParameters) SetSourceId(val []int32) {
	m.SourceId = val
}

// GetSourceIdN returns the SourceIdN property
func (m CoreDataFilesListQueryParameters) GetSourceIdN() []int32 {
	return m.SourceIdN
}

// SetSourceIdN sets the SourceIdN property
func (m *CoreDataFilesListQueryParameters) SetSourceIdN(val []int32) {
	m.SourceIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m CoreDataFilesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *CoreDataFilesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
