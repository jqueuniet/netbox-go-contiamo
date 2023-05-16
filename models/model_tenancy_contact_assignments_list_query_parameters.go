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

// TenancyContactAssignmentsListQueryParameters is an object.
type TenancyContactAssignmentsListQueryParameters struct {
	// ContactId: Contact (ID)
	ContactId []int32 `json:"contact_id,omitempty" mapstructure:"contact_id,omitempty"`
	// ContactIdN: Contact (ID)
	ContactIdN []int32 `json:"contact_id__n,omitempty" mapstructure:"contact_id__n,omitempty"`
	// ContentType:
	ContentType string `json:"content_type,omitempty" mapstructure:"content_type,omitempty"`
	// ContentTypeN:
	ContentTypeN string `json:"content_type__n,omitempty" mapstructure:"content_type__n,omitempty"`
	// ContentTypeId:
	ContentTypeId int32 `json:"content_type_id,omitempty" mapstructure:"content_type_id,omitempty"`
	// ContentTypeIdN:
	ContentTypeIdN int32 `json:"content_type_id__n,omitempty" mapstructure:"content_type_id__n,omitempty"`
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
	// ObjectId:
	ObjectId []int32 `json:"object_id,omitempty" mapstructure:"object_id,omitempty"`
	// ObjectIdGt:
	ObjectIdGt []int32 `json:"object_id__gt,omitempty" mapstructure:"object_id__gt,omitempty"`
	// ObjectIdGte:
	ObjectIdGte []int32 `json:"object_id__gte,omitempty" mapstructure:"object_id__gte,omitempty"`
	// ObjectIdLt:
	ObjectIdLt []int32 `json:"object_id__lt,omitempty" mapstructure:"object_id__lt,omitempty"`
	// ObjectIdLte:
	ObjectIdLte []int32 `json:"object_id__lte,omitempty" mapstructure:"object_id__lte,omitempty"`
	// ObjectIdN:
	ObjectIdN []int32 `json:"object_id__n,omitempty" mapstructure:"object_id__n,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Priority: * `primary` - Primary
	// * `secondary` - Secondary
	// * `tertiary` - Tertiary
	// * `inactive` - Inactive
	Priority string `json:"priority,omitempty" mapstructure:"priority,omitempty"`
	// PriorityN: * `primary` - Primary
	// * `secondary` - Secondary
	// * `tertiary` - Tertiary
	// * `inactive` - Inactive
	PriorityN string `json:"priority__n,omitempty" mapstructure:"priority__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Role: Contact role (slug)
	Role []string `json:"role,omitempty" mapstructure:"role,omitempty"`
	// RoleN: Contact role (slug)
	RoleN []string `json:"role__n,omitempty" mapstructure:"role__n,omitempty"`
	// RoleId: Contact role (ID)
	RoleId []int32 `json:"role_id,omitempty" mapstructure:"role_id,omitempty"`
	// RoleIdN: Contact role (ID)
	RoleIdN []int32 `json:"role_id__n,omitempty" mapstructure:"role_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m TenancyContactAssignmentsListQueryParameters) Validate() error {
	return validation.Errors{
		"contactId": validation.Validate(
			m.ContactId,
		),
		"contactIdN": validation.Validate(
			m.ContactIdN,
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
		"objectId": validation.Validate(
			m.ObjectId,
		),
		"objectIdGt": validation.Validate(
			m.ObjectIdGt,
		),
		"objectIdGte": validation.Validate(
			m.ObjectIdGte,
		),
		"objectIdLt": validation.Validate(
			m.ObjectIdLt,
		),
		"objectIdLte": validation.Validate(
			m.ObjectIdLte,
		),
		"objectIdN": validation.Validate(
			m.ObjectIdN,
		),
		"role": validation.Validate(
			m.Role,
		),
		"roleN": validation.Validate(
			m.RoleN,
		),
		"roleId": validation.Validate(
			m.RoleId,
		),
		"roleIdN": validation.Validate(
			m.RoleIdN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetContactId returns the ContactId property
func (m TenancyContactAssignmentsListQueryParameters) GetContactId() []int32 {
	return m.ContactId
}

// SetContactId sets the ContactId property
func (m *TenancyContactAssignmentsListQueryParameters) SetContactId(val []int32) {
	m.ContactId = val
}

// GetContactIdN returns the ContactIdN property
func (m TenancyContactAssignmentsListQueryParameters) GetContactIdN() []int32 {
	return m.ContactIdN
}

// SetContactIdN sets the ContactIdN property
func (m *TenancyContactAssignmentsListQueryParameters) SetContactIdN(val []int32) {
	m.ContactIdN = val
}

// GetContentType returns the ContentType property
func (m TenancyContactAssignmentsListQueryParameters) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *TenancyContactAssignmentsListQueryParameters) SetContentType(val string) {
	m.ContentType = val
}

// GetContentTypeN returns the ContentTypeN property
func (m TenancyContactAssignmentsListQueryParameters) GetContentTypeN() string {
	return m.ContentTypeN
}

// SetContentTypeN sets the ContentTypeN property
func (m *TenancyContactAssignmentsListQueryParameters) SetContentTypeN(val string) {
	m.ContentTypeN = val
}

// GetContentTypeId returns the ContentTypeId property
func (m TenancyContactAssignmentsListQueryParameters) GetContentTypeId() int32 {
	return m.ContentTypeId
}

// SetContentTypeId sets the ContentTypeId property
func (m *TenancyContactAssignmentsListQueryParameters) SetContentTypeId(val int32) {
	m.ContentTypeId = val
}

// GetContentTypeIdN returns the ContentTypeIdN property
func (m TenancyContactAssignmentsListQueryParameters) GetContentTypeIdN() int32 {
	return m.ContentTypeIdN
}

// SetContentTypeIdN sets the ContentTypeIdN property
func (m *TenancyContactAssignmentsListQueryParameters) SetContentTypeIdN(val int32) {
	m.ContentTypeIdN = val
}

// GetCreated returns the Created property
func (m TenancyContactAssignmentsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *TenancyContactAssignmentsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m TenancyContactAssignmentsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *TenancyContactAssignmentsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m TenancyContactAssignmentsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *TenancyContactAssignmentsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m TenancyContactAssignmentsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *TenancyContactAssignmentsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m TenancyContactAssignmentsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *TenancyContactAssignmentsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m TenancyContactAssignmentsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *TenancyContactAssignmentsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m TenancyContactAssignmentsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *TenancyContactAssignmentsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetId returns the Id property
func (m TenancyContactAssignmentsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *TenancyContactAssignmentsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m TenancyContactAssignmentsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *TenancyContactAssignmentsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m TenancyContactAssignmentsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *TenancyContactAssignmentsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m TenancyContactAssignmentsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *TenancyContactAssignmentsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m TenancyContactAssignmentsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *TenancyContactAssignmentsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m TenancyContactAssignmentsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *TenancyContactAssignmentsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m TenancyContactAssignmentsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *TenancyContactAssignmentsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m TenancyContactAssignmentsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *TenancyContactAssignmentsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m TenancyContactAssignmentsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *TenancyContactAssignmentsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m TenancyContactAssignmentsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *TenancyContactAssignmentsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m TenancyContactAssignmentsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *TenancyContactAssignmentsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m TenancyContactAssignmentsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *TenancyContactAssignmentsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m TenancyContactAssignmentsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *TenancyContactAssignmentsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetObjectId returns the ObjectId property
func (m TenancyContactAssignmentsListQueryParameters) GetObjectId() []int32 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *TenancyContactAssignmentsListQueryParameters) SetObjectId(val []int32) {
	m.ObjectId = val
}

// GetObjectIdGt returns the ObjectIdGt property
func (m TenancyContactAssignmentsListQueryParameters) GetObjectIdGt() []int32 {
	return m.ObjectIdGt
}

// SetObjectIdGt sets the ObjectIdGt property
func (m *TenancyContactAssignmentsListQueryParameters) SetObjectIdGt(val []int32) {
	m.ObjectIdGt = val
}

// GetObjectIdGte returns the ObjectIdGte property
func (m TenancyContactAssignmentsListQueryParameters) GetObjectIdGte() []int32 {
	return m.ObjectIdGte
}

// SetObjectIdGte sets the ObjectIdGte property
func (m *TenancyContactAssignmentsListQueryParameters) SetObjectIdGte(val []int32) {
	m.ObjectIdGte = val
}

// GetObjectIdLt returns the ObjectIdLt property
func (m TenancyContactAssignmentsListQueryParameters) GetObjectIdLt() []int32 {
	return m.ObjectIdLt
}

// SetObjectIdLt sets the ObjectIdLt property
func (m *TenancyContactAssignmentsListQueryParameters) SetObjectIdLt(val []int32) {
	m.ObjectIdLt = val
}

// GetObjectIdLte returns the ObjectIdLte property
func (m TenancyContactAssignmentsListQueryParameters) GetObjectIdLte() []int32 {
	return m.ObjectIdLte
}

// SetObjectIdLte sets the ObjectIdLte property
func (m *TenancyContactAssignmentsListQueryParameters) SetObjectIdLte(val []int32) {
	m.ObjectIdLte = val
}

// GetObjectIdN returns the ObjectIdN property
func (m TenancyContactAssignmentsListQueryParameters) GetObjectIdN() []int32 {
	return m.ObjectIdN
}

// SetObjectIdN sets the ObjectIdN property
func (m *TenancyContactAssignmentsListQueryParameters) SetObjectIdN(val []int32) {
	m.ObjectIdN = val
}

// GetOffset returns the Offset property
func (m TenancyContactAssignmentsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *TenancyContactAssignmentsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m TenancyContactAssignmentsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *TenancyContactAssignmentsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPriority returns the Priority property
func (m TenancyContactAssignmentsListQueryParameters) GetPriority() string {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *TenancyContactAssignmentsListQueryParameters) SetPriority(val string) {
	m.Priority = val
}

// GetPriorityN returns the PriorityN property
func (m TenancyContactAssignmentsListQueryParameters) GetPriorityN() string {
	return m.PriorityN
}

// SetPriorityN sets the PriorityN property
func (m *TenancyContactAssignmentsListQueryParameters) SetPriorityN(val string) {
	m.PriorityN = val
}

// GetQ returns the Q property
func (m TenancyContactAssignmentsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *TenancyContactAssignmentsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRole returns the Role property
func (m TenancyContactAssignmentsListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *TenancyContactAssignmentsListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m TenancyContactAssignmentsListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *TenancyContactAssignmentsListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m TenancyContactAssignmentsListQueryParameters) GetRoleId() []int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *TenancyContactAssignmentsListQueryParameters) SetRoleId(val []int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m TenancyContactAssignmentsListQueryParameters) GetRoleIdN() []int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *TenancyContactAssignmentsListQueryParameters) SetRoleIdN(val []int32) {
	m.RoleIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m TenancyContactAssignmentsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *TenancyContactAssignmentsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
