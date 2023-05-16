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

// CoreJobsListQueryParameters is an object.
type CoreJobsListQueryParameters struct {
	// Completed:
	Completed time.Time `json:"completed" mapstructure:"completed"`
	// CompletedAfter:
	CompletedAfter time.Time `json:"completed__after" mapstructure:"completed__after"`
	// CompletedBefore:
	CompletedBefore time.Time `json:"completed__before" mapstructure:"completed__before"`
	// Created:
	Created time.Time `json:"created" mapstructure:"created"`
	// CreatedAfter:
	CreatedAfter time.Time `json:"created__after" mapstructure:"created__after"`
	// CreatedBefore:
	CreatedBefore time.Time `json:"created__before" mapstructure:"created__before"`
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
	// Interval:
	Interval []int32 `json:"interval,omitempty" mapstructure:"interval,omitempty"`
	// IntervalGt:
	IntervalGt []int32 `json:"interval__gt,omitempty" mapstructure:"interval__gt,omitempty"`
	// IntervalGte:
	IntervalGte []int32 `json:"interval__gte,omitempty" mapstructure:"interval__gte,omitempty"`
	// IntervalLt:
	IntervalLt []int32 `json:"interval__lt,omitempty" mapstructure:"interval__lt,omitempty"`
	// IntervalLte:
	IntervalLte []int32 `json:"interval__lte,omitempty" mapstructure:"interval__lte,omitempty"`
	// IntervalN:
	IntervalN []int32 `json:"interval__n,omitempty" mapstructure:"interval__n,omitempty"`
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
	// ObjectType:
	ObjectType int32 `json:"object_type,omitempty" mapstructure:"object_type,omitempty"`
	// ObjectTypeN:
	ObjectTypeN int32 `json:"object_type__n,omitempty" mapstructure:"object_type__n,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Scheduled:
	Scheduled time.Time `json:"scheduled" mapstructure:"scheduled"`
	// ScheduledAfter:
	ScheduledAfter time.Time `json:"scheduled__after" mapstructure:"scheduled__after"`
	// ScheduledBefore:
	ScheduledBefore time.Time `json:"scheduled__before" mapstructure:"scheduled__before"`
	// Started:
	Started time.Time `json:"started" mapstructure:"started"`
	// StartedAfter:
	StartedAfter time.Time `json:"started__after" mapstructure:"started__after"`
	// StartedBefore:
	StartedBefore time.Time `json:"started__before" mapstructure:"started__before"`
	// Status: * `pending` - Pending
	// * `scheduled` - Scheduled
	// * `running` - Running
	// * `completed` - Completed
	// * `errored` - Errored
	// * `failed` - Failed
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `pending` - Pending
	// * `scheduled` - Scheduled
	// * `running` - Running
	// * `completed` - Completed
	// * `errored` - Errored
	// * `failed` - Failed
	StatusN []string `json:"status__n,omitempty" mapstructure:"status__n,omitempty"`
	// User:
	User int32 `json:"user,omitempty" mapstructure:"user,omitempty"`
	// UserN:
	UserN int32 `json:"user__n,omitempty" mapstructure:"user__n,omitempty"`
}

// Validate implements basic validation for this model
func (m CoreJobsListQueryParameters) Validate() error {
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
		"interval": validation.Validate(
			m.Interval,
		),
		"intervalGt": validation.Validate(
			m.IntervalGt,
		),
		"intervalGte": validation.Validate(
			m.IntervalGte,
		),
		"intervalLt": validation.Validate(
			m.IntervalLt,
		),
		"intervalLte": validation.Validate(
			m.IntervalLte,
		),
		"intervalN": validation.Validate(
			m.IntervalN,
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
		"status": validation.Validate(
			m.Status,
		),
		"statusN": validation.Validate(
			m.StatusN,
		),
	}.Filter()
}

// GetCompleted returns the Completed property
func (m CoreJobsListQueryParameters) GetCompleted() time.Time {
	return m.Completed
}

// SetCompleted sets the Completed property
func (m *CoreJobsListQueryParameters) SetCompleted(val time.Time) {
	m.Completed = val
}

// GetCompletedAfter returns the CompletedAfter property
func (m CoreJobsListQueryParameters) GetCompletedAfter() time.Time {
	return m.CompletedAfter
}

// SetCompletedAfter sets the CompletedAfter property
func (m *CoreJobsListQueryParameters) SetCompletedAfter(val time.Time) {
	m.CompletedAfter = val
}

// GetCompletedBefore returns the CompletedBefore property
func (m CoreJobsListQueryParameters) GetCompletedBefore() time.Time {
	return m.CompletedBefore
}

// SetCompletedBefore sets the CompletedBefore property
func (m *CoreJobsListQueryParameters) SetCompletedBefore(val time.Time) {
	m.CompletedBefore = val
}

// GetCreated returns the Created property
func (m CoreJobsListQueryParameters) GetCreated() time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CoreJobsListQueryParameters) SetCreated(val time.Time) {
	m.Created = val
}

// GetCreatedAfter returns the CreatedAfter property
func (m CoreJobsListQueryParameters) GetCreatedAfter() time.Time {
	return m.CreatedAfter
}

// SetCreatedAfter sets the CreatedAfter property
func (m *CoreJobsListQueryParameters) SetCreatedAfter(val time.Time) {
	m.CreatedAfter = val
}

// GetCreatedBefore returns the CreatedBefore property
func (m CoreJobsListQueryParameters) GetCreatedBefore() time.Time {
	return m.CreatedBefore
}

// SetCreatedBefore sets the CreatedBefore property
func (m *CoreJobsListQueryParameters) SetCreatedBefore(val time.Time) {
	m.CreatedBefore = val
}

// GetId returns the Id property
func (m CoreJobsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CoreJobsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m CoreJobsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *CoreJobsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m CoreJobsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *CoreJobsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m CoreJobsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *CoreJobsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m CoreJobsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *CoreJobsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m CoreJobsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *CoreJobsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterval returns the Interval property
func (m CoreJobsListQueryParameters) GetInterval() []int32 {
	return m.Interval
}

// SetInterval sets the Interval property
func (m *CoreJobsListQueryParameters) SetInterval(val []int32) {
	m.Interval = val
}

// GetIntervalGt returns the IntervalGt property
func (m CoreJobsListQueryParameters) GetIntervalGt() []int32 {
	return m.IntervalGt
}

// SetIntervalGt sets the IntervalGt property
func (m *CoreJobsListQueryParameters) SetIntervalGt(val []int32) {
	m.IntervalGt = val
}

// GetIntervalGte returns the IntervalGte property
func (m CoreJobsListQueryParameters) GetIntervalGte() []int32 {
	return m.IntervalGte
}

// SetIntervalGte sets the IntervalGte property
func (m *CoreJobsListQueryParameters) SetIntervalGte(val []int32) {
	m.IntervalGte = val
}

// GetIntervalLt returns the IntervalLt property
func (m CoreJobsListQueryParameters) GetIntervalLt() []int32 {
	return m.IntervalLt
}

// SetIntervalLt sets the IntervalLt property
func (m *CoreJobsListQueryParameters) SetIntervalLt(val []int32) {
	m.IntervalLt = val
}

// GetIntervalLte returns the IntervalLte property
func (m CoreJobsListQueryParameters) GetIntervalLte() []int32 {
	return m.IntervalLte
}

// SetIntervalLte sets the IntervalLte property
func (m *CoreJobsListQueryParameters) SetIntervalLte(val []int32) {
	m.IntervalLte = val
}

// GetIntervalN returns the IntervalN property
func (m CoreJobsListQueryParameters) GetIntervalN() []int32 {
	return m.IntervalN
}

// SetIntervalN sets the IntervalN property
func (m *CoreJobsListQueryParameters) SetIntervalN(val []int32) {
	m.IntervalN = val
}

// GetLimit returns the Limit property
func (m CoreJobsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *CoreJobsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m CoreJobsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *CoreJobsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m CoreJobsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *CoreJobsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m CoreJobsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *CoreJobsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m CoreJobsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *CoreJobsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m CoreJobsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *CoreJobsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m CoreJobsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *CoreJobsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m CoreJobsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *CoreJobsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m CoreJobsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *CoreJobsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m CoreJobsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *CoreJobsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m CoreJobsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *CoreJobsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m CoreJobsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *CoreJobsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetObjectId returns the ObjectId property
func (m CoreJobsListQueryParameters) GetObjectId() []int32 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *CoreJobsListQueryParameters) SetObjectId(val []int32) {
	m.ObjectId = val
}

// GetObjectIdGt returns the ObjectIdGt property
func (m CoreJobsListQueryParameters) GetObjectIdGt() []int32 {
	return m.ObjectIdGt
}

// SetObjectIdGt sets the ObjectIdGt property
func (m *CoreJobsListQueryParameters) SetObjectIdGt(val []int32) {
	m.ObjectIdGt = val
}

// GetObjectIdGte returns the ObjectIdGte property
func (m CoreJobsListQueryParameters) GetObjectIdGte() []int32 {
	return m.ObjectIdGte
}

// SetObjectIdGte sets the ObjectIdGte property
func (m *CoreJobsListQueryParameters) SetObjectIdGte(val []int32) {
	m.ObjectIdGte = val
}

// GetObjectIdLt returns the ObjectIdLt property
func (m CoreJobsListQueryParameters) GetObjectIdLt() []int32 {
	return m.ObjectIdLt
}

// SetObjectIdLt sets the ObjectIdLt property
func (m *CoreJobsListQueryParameters) SetObjectIdLt(val []int32) {
	m.ObjectIdLt = val
}

// GetObjectIdLte returns the ObjectIdLte property
func (m CoreJobsListQueryParameters) GetObjectIdLte() []int32 {
	return m.ObjectIdLte
}

// SetObjectIdLte sets the ObjectIdLte property
func (m *CoreJobsListQueryParameters) SetObjectIdLte(val []int32) {
	m.ObjectIdLte = val
}

// GetObjectIdN returns the ObjectIdN property
func (m CoreJobsListQueryParameters) GetObjectIdN() []int32 {
	return m.ObjectIdN
}

// SetObjectIdN sets the ObjectIdN property
func (m *CoreJobsListQueryParameters) SetObjectIdN(val []int32) {
	m.ObjectIdN = val
}

// GetObjectType returns the ObjectType property
func (m CoreJobsListQueryParameters) GetObjectType() int32 {
	return m.ObjectType
}

// SetObjectType sets the ObjectType property
func (m *CoreJobsListQueryParameters) SetObjectType(val int32) {
	m.ObjectType = val
}

// GetObjectTypeN returns the ObjectTypeN property
func (m CoreJobsListQueryParameters) GetObjectTypeN() int32 {
	return m.ObjectTypeN
}

// SetObjectTypeN sets the ObjectTypeN property
func (m *CoreJobsListQueryParameters) SetObjectTypeN(val int32) {
	m.ObjectTypeN = val
}

// GetOffset returns the Offset property
func (m CoreJobsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *CoreJobsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m CoreJobsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *CoreJobsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m CoreJobsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *CoreJobsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetScheduled returns the Scheduled property
func (m CoreJobsListQueryParameters) GetScheduled() time.Time {
	return m.Scheduled
}

// SetScheduled sets the Scheduled property
func (m *CoreJobsListQueryParameters) SetScheduled(val time.Time) {
	m.Scheduled = val
}

// GetScheduledAfter returns the ScheduledAfter property
func (m CoreJobsListQueryParameters) GetScheduledAfter() time.Time {
	return m.ScheduledAfter
}

// SetScheduledAfter sets the ScheduledAfter property
func (m *CoreJobsListQueryParameters) SetScheduledAfter(val time.Time) {
	m.ScheduledAfter = val
}

// GetScheduledBefore returns the ScheduledBefore property
func (m CoreJobsListQueryParameters) GetScheduledBefore() time.Time {
	return m.ScheduledBefore
}

// SetScheduledBefore sets the ScheduledBefore property
func (m *CoreJobsListQueryParameters) SetScheduledBefore(val time.Time) {
	m.ScheduledBefore = val
}

// GetStarted returns the Started property
func (m CoreJobsListQueryParameters) GetStarted() time.Time {
	return m.Started
}

// SetStarted sets the Started property
func (m *CoreJobsListQueryParameters) SetStarted(val time.Time) {
	m.Started = val
}

// GetStartedAfter returns the StartedAfter property
func (m CoreJobsListQueryParameters) GetStartedAfter() time.Time {
	return m.StartedAfter
}

// SetStartedAfter sets the StartedAfter property
func (m *CoreJobsListQueryParameters) SetStartedAfter(val time.Time) {
	m.StartedAfter = val
}

// GetStartedBefore returns the StartedBefore property
func (m CoreJobsListQueryParameters) GetStartedBefore() time.Time {
	return m.StartedBefore
}

// SetStartedBefore sets the StartedBefore property
func (m *CoreJobsListQueryParameters) SetStartedBefore(val time.Time) {
	m.StartedBefore = val
}

// GetStatus returns the Status property
func (m CoreJobsListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *CoreJobsListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m CoreJobsListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *CoreJobsListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetUser returns the User property
func (m CoreJobsListQueryParameters) GetUser() int32 {
	return m.User
}

// SetUser sets the User property
func (m *CoreJobsListQueryParameters) SetUser(val int32) {
	m.User = val
}

// GetUserN returns the UserN property
func (m CoreJobsListQueryParameters) GetUserN() int32 {
	return m.UserN
}

// SetUserN sets the UserN property
func (m *CoreJobsListQueryParameters) SetUserN(val int32) {
	m.UserN = val
}
