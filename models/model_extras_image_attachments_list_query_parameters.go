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
)

// ExtrasImageAttachmentsListQueryParameters is an object.
type ExtrasImageAttachmentsListQueryParameters struct {
	// ContentType:
	ContentType string `json:"content_type,omitempty" mapstructure:"content_type,omitempty"`
	// ContentTypeN:
	ContentTypeN string `json:"content_type__n,omitempty" mapstructure:"content_type__n,omitempty"`
	// ContentTypeId:
	ContentTypeId int32 `json:"content_type_id,omitempty" mapstructure:"content_type_id,omitempty"`
	// ContentTypeIdN:
	ContentTypeIdN int32 `json:"content_type_id__n,omitempty" mapstructure:"content_type_id__n,omitempty"`
	// Created:
	Created time.Time `json:"created" mapstructure:"created"`
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
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasImageAttachmentsListQueryParameters) Validate() error {
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
	}.Filter()
}

// GetContentType returns the ContentType property
func (m ExtrasImageAttachmentsListQueryParameters) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *ExtrasImageAttachmentsListQueryParameters) SetContentType(val string) {
	m.ContentType = val
}

// GetContentTypeN returns the ContentTypeN property
func (m ExtrasImageAttachmentsListQueryParameters) GetContentTypeN() string {
	return m.ContentTypeN
}

// SetContentTypeN sets the ContentTypeN property
func (m *ExtrasImageAttachmentsListQueryParameters) SetContentTypeN(val string) {
	m.ContentTypeN = val
}

// GetContentTypeId returns the ContentTypeId property
func (m ExtrasImageAttachmentsListQueryParameters) GetContentTypeId() int32 {
	return m.ContentTypeId
}

// SetContentTypeId sets the ContentTypeId property
func (m *ExtrasImageAttachmentsListQueryParameters) SetContentTypeId(val int32) {
	m.ContentTypeId = val
}

// GetContentTypeIdN returns the ContentTypeIdN property
func (m ExtrasImageAttachmentsListQueryParameters) GetContentTypeIdN() int32 {
	return m.ContentTypeIdN
}

// SetContentTypeIdN sets the ContentTypeIdN property
func (m *ExtrasImageAttachmentsListQueryParameters) SetContentTypeIdN(val int32) {
	m.ContentTypeIdN = val
}

// GetCreated returns the Created property
func (m ExtrasImageAttachmentsListQueryParameters) GetCreated() time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ExtrasImageAttachmentsListQueryParameters) SetCreated(val time.Time) {
	m.Created = val
}

// GetId returns the Id property
func (m ExtrasImageAttachmentsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasImageAttachmentsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasImageAttachmentsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasImageAttachmentsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasImageAttachmentsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasImageAttachmentsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasImageAttachmentsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasImageAttachmentsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasImageAttachmentsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasImageAttachmentsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasImageAttachmentsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasImageAttachmentsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m ExtrasImageAttachmentsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasImageAttachmentsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m ExtrasImageAttachmentsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasImageAttachmentsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasImageAttachmentsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasImageAttachmentsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetObjectId returns the ObjectId property
func (m ExtrasImageAttachmentsListQueryParameters) GetObjectId() []int32 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *ExtrasImageAttachmentsListQueryParameters) SetObjectId(val []int32) {
	m.ObjectId = val
}

// GetObjectIdGt returns the ObjectIdGt property
func (m ExtrasImageAttachmentsListQueryParameters) GetObjectIdGt() []int32 {
	return m.ObjectIdGt
}

// SetObjectIdGt sets the ObjectIdGt property
func (m *ExtrasImageAttachmentsListQueryParameters) SetObjectIdGt(val []int32) {
	m.ObjectIdGt = val
}

// GetObjectIdGte returns the ObjectIdGte property
func (m ExtrasImageAttachmentsListQueryParameters) GetObjectIdGte() []int32 {
	return m.ObjectIdGte
}

// SetObjectIdGte sets the ObjectIdGte property
func (m *ExtrasImageAttachmentsListQueryParameters) SetObjectIdGte(val []int32) {
	m.ObjectIdGte = val
}

// GetObjectIdLt returns the ObjectIdLt property
func (m ExtrasImageAttachmentsListQueryParameters) GetObjectIdLt() []int32 {
	return m.ObjectIdLt
}

// SetObjectIdLt sets the ObjectIdLt property
func (m *ExtrasImageAttachmentsListQueryParameters) SetObjectIdLt(val []int32) {
	m.ObjectIdLt = val
}

// GetObjectIdLte returns the ObjectIdLte property
func (m ExtrasImageAttachmentsListQueryParameters) GetObjectIdLte() []int32 {
	return m.ObjectIdLte
}

// SetObjectIdLte sets the ObjectIdLte property
func (m *ExtrasImageAttachmentsListQueryParameters) SetObjectIdLte(val []int32) {
	m.ObjectIdLte = val
}

// GetObjectIdN returns the ObjectIdN property
func (m ExtrasImageAttachmentsListQueryParameters) GetObjectIdN() []int32 {
	return m.ObjectIdN
}

// SetObjectIdN sets the ObjectIdN property
func (m *ExtrasImageAttachmentsListQueryParameters) SetObjectIdN(val []int32) {
	m.ObjectIdN = val
}

// GetOffset returns the Offset property
func (m ExtrasImageAttachmentsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasImageAttachmentsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasImageAttachmentsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasImageAttachmentsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasImageAttachmentsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasImageAttachmentsListQueryParameters) SetQ(val string) {
	m.Q = val
}
