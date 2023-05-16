// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ExtrasJournalEntriesListQueryParameters is an object.
type ExtrasJournalEntriesListQueryParameters struct {
	// AssignedObjectId:
	AssignedObjectId []int32 `json:"assigned_object_id,omitempty" mapstructure:"assigned_object_id,omitempty"`
	// AssignedObjectIdGt:
	AssignedObjectIdGt []int32 `json:"assigned_object_id__gt,omitempty" mapstructure:"assigned_object_id__gt,omitempty"`
	// AssignedObjectIdGte:
	AssignedObjectIdGte []int32 `json:"assigned_object_id__gte,omitempty" mapstructure:"assigned_object_id__gte,omitempty"`
	// AssignedObjectIdLt:
	AssignedObjectIdLt []int32 `json:"assigned_object_id__lt,omitempty" mapstructure:"assigned_object_id__lt,omitempty"`
	// AssignedObjectIdLte:
	AssignedObjectIdLte []int32 `json:"assigned_object_id__lte,omitempty" mapstructure:"assigned_object_id__lte,omitempty"`
	// AssignedObjectIdN:
	AssignedObjectIdN []int32 `json:"assigned_object_id__n,omitempty" mapstructure:"assigned_object_id__n,omitempty"`
	// AssignedObjectType:
	AssignedObjectType string `json:"assigned_object_type,omitempty" mapstructure:"assigned_object_type,omitempty"`
	// AssignedObjectTypeN:
	AssignedObjectTypeN string `json:"assigned_object_type__n,omitempty" mapstructure:"assigned_object_type__n,omitempty"`
	// AssignedObjectTypeId:
	AssignedObjectTypeId []int32 `json:"assigned_object_type_id,omitempty" mapstructure:"assigned_object_type_id,omitempty"`
	// AssignedObjectTypeIdN:
	AssignedObjectTypeIdN []int32 `json:"assigned_object_type_id__n,omitempty" mapstructure:"assigned_object_type_id__n,omitempty"`
	// CreatedAfter:
	CreatedAfter time.Time `json:"created_after" mapstructure:"created_after"`
	// CreatedBefore:
	CreatedBefore time.Time `json:"created_before" mapstructure:"created_before"`
	// CreatedBy: User (name)
	CreatedBy []string `json:"created_by,omitempty" mapstructure:"created_by,omitempty"`
	// CreatedByN: User (name)
	CreatedByN []string `json:"created_by__n,omitempty" mapstructure:"created_by__n,omitempty"`
	// CreatedById: User (ID)
	CreatedById []*int32 `json:"created_by_id,omitempty" mapstructure:"created_by_id,omitempty"`
	// CreatedByIdN: User (ID)
	CreatedByIdN []*int32 `json:"created_by_id__n,omitempty" mapstructure:"created_by_id__n,omitempty"`
	// CreatedByRequest:
	CreatedByRequest string `json:"created_by_request,omitempty" mapstructure:"created_by_request,omitempty"`
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
	// Kind: * `info` - Info
	// * `success` - Success
	// * `warning` - Warning
	// * `danger` - Danger
	Kind []string `json:"kind,omitempty" mapstructure:"kind,omitempty"`
	// KindN: * `info` - Info
	// * `success` - Success
	// * `warning` - Warning
	// * `danger` - Danger
	KindN []string `json:"kind__n,omitempty" mapstructure:"kind__n,omitempty"`
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
func (m ExtrasJournalEntriesListQueryParameters) Validate() error {
	return validation.Errors{
		"assignedObjectId": validation.Validate(
			m.AssignedObjectId,
		),
		"assignedObjectIdGt": validation.Validate(
			m.AssignedObjectIdGt,
		),
		"assignedObjectIdGte": validation.Validate(
			m.AssignedObjectIdGte,
		),
		"assignedObjectIdLt": validation.Validate(
			m.AssignedObjectIdLt,
		),
		"assignedObjectIdLte": validation.Validate(
			m.AssignedObjectIdLte,
		),
		"assignedObjectIdN": validation.Validate(
			m.AssignedObjectIdN,
		),
		"assignedObjectTypeId": validation.Validate(
			m.AssignedObjectTypeId,
		),
		"assignedObjectTypeIdN": validation.Validate(
			m.AssignedObjectTypeIdN,
		),
		"createdBy": validation.Validate(
			m.CreatedBy,
		),
		"createdByN": validation.Validate(
			m.CreatedByN,
		),
		"createdById": validation.Validate(
			m.CreatedById,
		),
		"createdByIdN": validation.Validate(
			m.CreatedByIdN,
		),
		"createdByRequest": validation.Validate(
			m.CreatedByRequest, is.UUID,
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
		"kind": validation.Validate(
			m.Kind,
		),
		"kindN": validation.Validate(
			m.KindN,
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

// GetAssignedObjectId returns the AssignedObjectId property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectId() []int32 {
	return m.AssignedObjectId
}

// SetAssignedObjectId sets the AssignedObjectId property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectId(val []int32) {
	m.AssignedObjectId = val
}

// GetAssignedObjectIdGt returns the AssignedObjectIdGt property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectIdGt() []int32 {
	return m.AssignedObjectIdGt
}

// SetAssignedObjectIdGt sets the AssignedObjectIdGt property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectIdGt(val []int32) {
	m.AssignedObjectIdGt = val
}

// GetAssignedObjectIdGte returns the AssignedObjectIdGte property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectIdGte() []int32 {
	return m.AssignedObjectIdGte
}

// SetAssignedObjectIdGte sets the AssignedObjectIdGte property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectIdGte(val []int32) {
	m.AssignedObjectIdGte = val
}

// GetAssignedObjectIdLt returns the AssignedObjectIdLt property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectIdLt() []int32 {
	return m.AssignedObjectIdLt
}

// SetAssignedObjectIdLt sets the AssignedObjectIdLt property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectIdLt(val []int32) {
	m.AssignedObjectIdLt = val
}

// GetAssignedObjectIdLte returns the AssignedObjectIdLte property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectIdLte() []int32 {
	return m.AssignedObjectIdLte
}

// SetAssignedObjectIdLte sets the AssignedObjectIdLte property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectIdLte(val []int32) {
	m.AssignedObjectIdLte = val
}

// GetAssignedObjectIdN returns the AssignedObjectIdN property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectIdN() []int32 {
	return m.AssignedObjectIdN
}

// SetAssignedObjectIdN sets the AssignedObjectIdN property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectIdN(val []int32) {
	m.AssignedObjectIdN = val
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectType() string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectType(val string) {
	m.AssignedObjectType = val
}

// GetAssignedObjectTypeN returns the AssignedObjectTypeN property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectTypeN() string {
	return m.AssignedObjectTypeN
}

// SetAssignedObjectTypeN sets the AssignedObjectTypeN property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectTypeN(val string) {
	m.AssignedObjectTypeN = val
}

// GetAssignedObjectTypeId returns the AssignedObjectTypeId property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectTypeId() []int32 {
	return m.AssignedObjectTypeId
}

// SetAssignedObjectTypeId sets the AssignedObjectTypeId property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectTypeId(val []int32) {
	m.AssignedObjectTypeId = val
}

// GetAssignedObjectTypeIdN returns the AssignedObjectTypeIdN property
func (m ExtrasJournalEntriesListQueryParameters) GetAssignedObjectTypeIdN() []int32 {
	return m.AssignedObjectTypeIdN
}

// SetAssignedObjectTypeIdN sets the AssignedObjectTypeIdN property
func (m *ExtrasJournalEntriesListQueryParameters) SetAssignedObjectTypeIdN(val []int32) {
	m.AssignedObjectTypeIdN = val
}

// GetCreatedAfter returns the CreatedAfter property
func (m ExtrasJournalEntriesListQueryParameters) GetCreatedAfter() time.Time {
	return m.CreatedAfter
}

// SetCreatedAfter sets the CreatedAfter property
func (m *ExtrasJournalEntriesListQueryParameters) SetCreatedAfter(val time.Time) {
	m.CreatedAfter = val
}

// GetCreatedBefore returns the CreatedBefore property
func (m ExtrasJournalEntriesListQueryParameters) GetCreatedBefore() time.Time {
	return m.CreatedBefore
}

// SetCreatedBefore sets the CreatedBefore property
func (m *ExtrasJournalEntriesListQueryParameters) SetCreatedBefore(val time.Time) {
	m.CreatedBefore = val
}

// GetCreatedBy returns the CreatedBy property
func (m ExtrasJournalEntriesListQueryParameters) GetCreatedBy() []string {
	return m.CreatedBy
}

// SetCreatedBy sets the CreatedBy property
func (m *ExtrasJournalEntriesListQueryParameters) SetCreatedBy(val []string) {
	m.CreatedBy = val
}

// GetCreatedByN returns the CreatedByN property
func (m ExtrasJournalEntriesListQueryParameters) GetCreatedByN() []string {
	return m.CreatedByN
}

// SetCreatedByN sets the CreatedByN property
func (m *ExtrasJournalEntriesListQueryParameters) SetCreatedByN(val []string) {
	m.CreatedByN = val
}

// GetCreatedById returns the CreatedById property
func (m ExtrasJournalEntriesListQueryParameters) GetCreatedById() []*int32 {
	return m.CreatedById
}

// SetCreatedById sets the CreatedById property
func (m *ExtrasJournalEntriesListQueryParameters) SetCreatedById(val []*int32) {
	m.CreatedById = val
}

// GetCreatedByIdN returns the CreatedByIdN property
func (m ExtrasJournalEntriesListQueryParameters) GetCreatedByIdN() []*int32 {
	return m.CreatedByIdN
}

// SetCreatedByIdN sets the CreatedByIdN property
func (m *ExtrasJournalEntriesListQueryParameters) SetCreatedByIdN(val []*int32) {
	m.CreatedByIdN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m ExtrasJournalEntriesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *ExtrasJournalEntriesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetId returns the Id property
func (m ExtrasJournalEntriesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasJournalEntriesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasJournalEntriesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasJournalEntriesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasJournalEntriesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasJournalEntriesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasJournalEntriesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasJournalEntriesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasJournalEntriesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasJournalEntriesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasJournalEntriesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasJournalEntriesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetKind returns the Kind property
func (m ExtrasJournalEntriesListQueryParameters) GetKind() []string {
	return m.Kind
}

// SetKind sets the Kind property
func (m *ExtrasJournalEntriesListQueryParameters) SetKind(val []string) {
	m.Kind = val
}

// GetKindN returns the KindN property
func (m ExtrasJournalEntriesListQueryParameters) GetKindN() []string {
	return m.KindN
}

// SetKindN sets the KindN property
func (m *ExtrasJournalEntriesListQueryParameters) SetKindN(val []string) {
	m.KindN = val
}

// GetLastUpdated returns the LastUpdated property
func (m ExtrasJournalEntriesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ExtrasJournalEntriesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m ExtrasJournalEntriesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *ExtrasJournalEntriesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m ExtrasJournalEntriesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *ExtrasJournalEntriesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m ExtrasJournalEntriesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *ExtrasJournalEntriesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m ExtrasJournalEntriesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *ExtrasJournalEntriesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m ExtrasJournalEntriesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *ExtrasJournalEntriesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m ExtrasJournalEntriesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasJournalEntriesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m ExtrasJournalEntriesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasJournalEntriesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasJournalEntriesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasJournalEntriesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasJournalEntriesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasJournalEntriesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m ExtrasJournalEntriesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *ExtrasJournalEntriesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m ExtrasJournalEntriesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *ExtrasJournalEntriesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m ExtrasJournalEntriesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *ExtrasJournalEntriesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
