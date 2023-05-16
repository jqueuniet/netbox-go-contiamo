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

	"time"
)

// ExtrasObjectChangesListQueryParameters is an object.
type ExtrasObjectChangesListQueryParameters struct {
	// Action: * `create` - Created
	// * `update` - Updated
	// * `delete` - Deleted
	Action string `json:"action,omitempty" mapstructure:"action,omitempty"`
	// ActionN: * `create` - Created
	// * `update` - Updated
	// * `delete` - Deleted
	ActionN string `json:"action__n,omitempty" mapstructure:"action__n,omitempty"`
	// ChangedObjectId:
	ChangedObjectId []int32 `json:"changed_object_id,omitempty" mapstructure:"changed_object_id,omitempty"`
	// ChangedObjectIdGt:
	ChangedObjectIdGt []int32 `json:"changed_object_id__gt,omitempty" mapstructure:"changed_object_id__gt,omitempty"`
	// ChangedObjectIdGte:
	ChangedObjectIdGte []int32 `json:"changed_object_id__gte,omitempty" mapstructure:"changed_object_id__gte,omitempty"`
	// ChangedObjectIdLt:
	ChangedObjectIdLt []int32 `json:"changed_object_id__lt,omitempty" mapstructure:"changed_object_id__lt,omitempty"`
	// ChangedObjectIdLte:
	ChangedObjectIdLte []int32 `json:"changed_object_id__lte,omitempty" mapstructure:"changed_object_id__lte,omitempty"`
	// ChangedObjectIdN:
	ChangedObjectIdN []int32 `json:"changed_object_id__n,omitempty" mapstructure:"changed_object_id__n,omitempty"`
	// ChangedObjectType:
	ChangedObjectType string `json:"changed_object_type,omitempty" mapstructure:"changed_object_type,omitempty"`
	// ChangedObjectTypeN:
	ChangedObjectTypeN string `json:"changed_object_type__n,omitempty" mapstructure:"changed_object_type__n,omitempty"`
	// ChangedObjectTypeId:
	ChangedObjectTypeId []int32 `json:"changed_object_type_id,omitempty" mapstructure:"changed_object_type_id,omitempty"`
	// ChangedObjectTypeIdN:
	ChangedObjectTypeIdN []int32 `json:"changed_object_type_id__n,omitempty" mapstructure:"changed_object_type_id__n,omitempty"`
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
	// ObjectRepr:
	ObjectRepr []string `json:"object_repr,omitempty" mapstructure:"object_repr,omitempty"`
	// ObjectReprEmpty:
	ObjectReprEmpty []string `json:"object_repr__empty,omitempty" mapstructure:"object_repr__empty,omitempty"`
	// ObjectReprIc:
	ObjectReprIc []string `json:"object_repr__ic,omitempty" mapstructure:"object_repr__ic,omitempty"`
	// ObjectReprIe:
	ObjectReprIe []string `json:"object_repr__ie,omitempty" mapstructure:"object_repr__ie,omitempty"`
	// ObjectReprIew:
	ObjectReprIew []string `json:"object_repr__iew,omitempty" mapstructure:"object_repr__iew,omitempty"`
	// ObjectReprIsw:
	ObjectReprIsw []string `json:"object_repr__isw,omitempty" mapstructure:"object_repr__isw,omitempty"`
	// ObjectReprN:
	ObjectReprN []string `json:"object_repr__n,omitempty" mapstructure:"object_repr__n,omitempty"`
	// ObjectReprNic:
	ObjectReprNic []string `json:"object_repr__nic,omitempty" mapstructure:"object_repr__nic,omitempty"`
	// ObjectReprNie:
	ObjectReprNie []string `json:"object_repr__nie,omitempty" mapstructure:"object_repr__nie,omitempty"`
	// ObjectReprNiew:
	ObjectReprNiew []string `json:"object_repr__niew,omitempty" mapstructure:"object_repr__niew,omitempty"`
	// ObjectReprNisw:
	ObjectReprNisw []string `json:"object_repr__nisw,omitempty" mapstructure:"object_repr__nisw,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// RequestId:
	RequestId string `json:"request_id,omitempty" mapstructure:"request_id,omitempty"`
	// TimeAfter:
	TimeAfter time.Time `json:"time_after" mapstructure:"time_after"`
	// TimeBefore:
	TimeBefore time.Time `json:"time_before" mapstructure:"time_before"`
	// User: User name
	User []string `json:"user,omitempty" mapstructure:"user,omitempty"`
	// UserN: User name
	UserN []string `json:"user__n,omitempty" mapstructure:"user__n,omitempty"`
	// UserId: User (ID)
	UserId []*int32 `json:"user_id,omitempty" mapstructure:"user_id,omitempty"`
	// UserIdN: User (ID)
	UserIdN []*int32 `json:"user_id__n,omitempty" mapstructure:"user_id__n,omitempty"`
	// UserName:
	UserName []string `json:"user_name,omitempty" mapstructure:"user_name,omitempty"`
	// UserNameEmpty:
	UserNameEmpty []string `json:"user_name__empty,omitempty" mapstructure:"user_name__empty,omitempty"`
	// UserNameIc:
	UserNameIc []string `json:"user_name__ic,omitempty" mapstructure:"user_name__ic,omitempty"`
	// UserNameIe:
	UserNameIe []string `json:"user_name__ie,omitempty" mapstructure:"user_name__ie,omitempty"`
	// UserNameIew:
	UserNameIew []string `json:"user_name__iew,omitempty" mapstructure:"user_name__iew,omitempty"`
	// UserNameIsw:
	UserNameIsw []string `json:"user_name__isw,omitempty" mapstructure:"user_name__isw,omitempty"`
	// UserNameN:
	UserNameN []string `json:"user_name__n,omitempty" mapstructure:"user_name__n,omitempty"`
	// UserNameNic:
	UserNameNic []string `json:"user_name__nic,omitempty" mapstructure:"user_name__nic,omitempty"`
	// UserNameNie:
	UserNameNie []string `json:"user_name__nie,omitempty" mapstructure:"user_name__nie,omitempty"`
	// UserNameNiew:
	UserNameNiew []string `json:"user_name__niew,omitempty" mapstructure:"user_name__niew,omitempty"`
	// UserNameNisw:
	UserNameNisw []string `json:"user_name__nisw,omitempty" mapstructure:"user_name__nisw,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasObjectChangesListQueryParameters) Validate() error {
	return validation.Errors{
		"changedObjectId": validation.Validate(
			m.ChangedObjectId,
		),
		"changedObjectIdGt": validation.Validate(
			m.ChangedObjectIdGt,
		),
		"changedObjectIdGte": validation.Validate(
			m.ChangedObjectIdGte,
		),
		"changedObjectIdLt": validation.Validate(
			m.ChangedObjectIdLt,
		),
		"changedObjectIdLte": validation.Validate(
			m.ChangedObjectIdLte,
		),
		"changedObjectIdN": validation.Validate(
			m.ChangedObjectIdN,
		),
		"changedObjectTypeId": validation.Validate(
			m.ChangedObjectTypeId,
		),
		"changedObjectTypeIdN": validation.Validate(
			m.ChangedObjectTypeIdN,
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
		"objectRepr": validation.Validate(
			m.ObjectRepr,
		),
		"objectReprEmpty": validation.Validate(
			m.ObjectReprEmpty,
		),
		"objectReprIc": validation.Validate(
			m.ObjectReprIc,
		),
		"objectReprIe": validation.Validate(
			m.ObjectReprIe,
		),
		"objectReprIew": validation.Validate(
			m.ObjectReprIew,
		),
		"objectReprIsw": validation.Validate(
			m.ObjectReprIsw,
		),
		"objectReprN": validation.Validate(
			m.ObjectReprN,
		),
		"objectReprNic": validation.Validate(
			m.ObjectReprNic,
		),
		"objectReprNie": validation.Validate(
			m.ObjectReprNie,
		),
		"objectReprNiew": validation.Validate(
			m.ObjectReprNiew,
		),
		"objectReprNisw": validation.Validate(
			m.ObjectReprNisw,
		),
		"requestId": validation.Validate(
			m.RequestId, is.UUID,
		),
		"user": validation.Validate(
			m.User,
		),
		"userN": validation.Validate(
			m.UserN,
		),
		"userId": validation.Validate(
			m.UserId,
		),
		"userIdN": validation.Validate(
			m.UserIdN,
		),
		"userName": validation.Validate(
			m.UserName,
		),
		"userNameEmpty": validation.Validate(
			m.UserNameEmpty,
		),
		"userNameIc": validation.Validate(
			m.UserNameIc,
		),
		"userNameIe": validation.Validate(
			m.UserNameIe,
		),
		"userNameIew": validation.Validate(
			m.UserNameIew,
		),
		"userNameIsw": validation.Validate(
			m.UserNameIsw,
		),
		"userNameN": validation.Validate(
			m.UserNameN,
		),
		"userNameNic": validation.Validate(
			m.UserNameNic,
		),
		"userNameNie": validation.Validate(
			m.UserNameNie,
		),
		"userNameNiew": validation.Validate(
			m.UserNameNiew,
		),
		"userNameNisw": validation.Validate(
			m.UserNameNisw,
		),
	}.Filter()
}

// GetAction returns the Action property
func (m ExtrasObjectChangesListQueryParameters) GetAction() string {
	return m.Action
}

// SetAction sets the Action property
func (m *ExtrasObjectChangesListQueryParameters) SetAction(val string) {
	m.Action = val
}

// GetActionN returns the ActionN property
func (m ExtrasObjectChangesListQueryParameters) GetActionN() string {
	return m.ActionN
}

// SetActionN sets the ActionN property
func (m *ExtrasObjectChangesListQueryParameters) SetActionN(val string) {
	m.ActionN = val
}

// GetChangedObjectId returns the ChangedObjectId property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectId() []int32 {
	return m.ChangedObjectId
}

// SetChangedObjectId sets the ChangedObjectId property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectId(val []int32) {
	m.ChangedObjectId = val
}

// GetChangedObjectIdGt returns the ChangedObjectIdGt property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectIdGt() []int32 {
	return m.ChangedObjectIdGt
}

// SetChangedObjectIdGt sets the ChangedObjectIdGt property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectIdGt(val []int32) {
	m.ChangedObjectIdGt = val
}

// GetChangedObjectIdGte returns the ChangedObjectIdGte property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectIdGte() []int32 {
	return m.ChangedObjectIdGte
}

// SetChangedObjectIdGte sets the ChangedObjectIdGte property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectIdGte(val []int32) {
	m.ChangedObjectIdGte = val
}

// GetChangedObjectIdLt returns the ChangedObjectIdLt property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectIdLt() []int32 {
	return m.ChangedObjectIdLt
}

// SetChangedObjectIdLt sets the ChangedObjectIdLt property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectIdLt(val []int32) {
	m.ChangedObjectIdLt = val
}

// GetChangedObjectIdLte returns the ChangedObjectIdLte property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectIdLte() []int32 {
	return m.ChangedObjectIdLte
}

// SetChangedObjectIdLte sets the ChangedObjectIdLte property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectIdLte(val []int32) {
	m.ChangedObjectIdLte = val
}

// GetChangedObjectIdN returns the ChangedObjectIdN property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectIdN() []int32 {
	return m.ChangedObjectIdN
}

// SetChangedObjectIdN sets the ChangedObjectIdN property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectIdN(val []int32) {
	m.ChangedObjectIdN = val
}

// GetChangedObjectType returns the ChangedObjectType property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectType() string {
	return m.ChangedObjectType
}

// SetChangedObjectType sets the ChangedObjectType property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectType(val string) {
	m.ChangedObjectType = val
}

// GetChangedObjectTypeN returns the ChangedObjectTypeN property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectTypeN() string {
	return m.ChangedObjectTypeN
}

// SetChangedObjectTypeN sets the ChangedObjectTypeN property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectTypeN(val string) {
	m.ChangedObjectTypeN = val
}

// GetChangedObjectTypeId returns the ChangedObjectTypeId property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectTypeId() []int32 {
	return m.ChangedObjectTypeId
}

// SetChangedObjectTypeId sets the ChangedObjectTypeId property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectTypeId(val []int32) {
	m.ChangedObjectTypeId = val
}

// GetChangedObjectTypeIdN returns the ChangedObjectTypeIdN property
func (m ExtrasObjectChangesListQueryParameters) GetChangedObjectTypeIdN() []int32 {
	return m.ChangedObjectTypeIdN
}

// SetChangedObjectTypeIdN sets the ChangedObjectTypeIdN property
func (m *ExtrasObjectChangesListQueryParameters) SetChangedObjectTypeIdN(val []int32) {
	m.ChangedObjectTypeIdN = val
}

// GetId returns the Id property
func (m ExtrasObjectChangesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasObjectChangesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasObjectChangesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasObjectChangesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasObjectChangesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasObjectChangesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasObjectChangesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasObjectChangesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasObjectChangesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasObjectChangesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasObjectChangesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasObjectChangesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m ExtrasObjectChangesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasObjectChangesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetObjectRepr returns the ObjectRepr property
func (m ExtrasObjectChangesListQueryParameters) GetObjectRepr() []string {
	return m.ObjectRepr
}

// SetObjectRepr sets the ObjectRepr property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectRepr(val []string) {
	m.ObjectRepr = val
}

// GetObjectReprEmpty returns the ObjectReprEmpty property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprEmpty() []string {
	return m.ObjectReprEmpty
}

// SetObjectReprEmpty sets the ObjectReprEmpty property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprEmpty(val []string) {
	m.ObjectReprEmpty = val
}

// GetObjectReprIc returns the ObjectReprIc property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprIc() []string {
	return m.ObjectReprIc
}

// SetObjectReprIc sets the ObjectReprIc property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprIc(val []string) {
	m.ObjectReprIc = val
}

// GetObjectReprIe returns the ObjectReprIe property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprIe() []string {
	return m.ObjectReprIe
}

// SetObjectReprIe sets the ObjectReprIe property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprIe(val []string) {
	m.ObjectReprIe = val
}

// GetObjectReprIew returns the ObjectReprIew property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprIew() []string {
	return m.ObjectReprIew
}

// SetObjectReprIew sets the ObjectReprIew property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprIew(val []string) {
	m.ObjectReprIew = val
}

// GetObjectReprIsw returns the ObjectReprIsw property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprIsw() []string {
	return m.ObjectReprIsw
}

// SetObjectReprIsw sets the ObjectReprIsw property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprIsw(val []string) {
	m.ObjectReprIsw = val
}

// GetObjectReprN returns the ObjectReprN property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprN() []string {
	return m.ObjectReprN
}

// SetObjectReprN sets the ObjectReprN property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprN(val []string) {
	m.ObjectReprN = val
}

// GetObjectReprNic returns the ObjectReprNic property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprNic() []string {
	return m.ObjectReprNic
}

// SetObjectReprNic sets the ObjectReprNic property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprNic(val []string) {
	m.ObjectReprNic = val
}

// GetObjectReprNie returns the ObjectReprNie property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprNie() []string {
	return m.ObjectReprNie
}

// SetObjectReprNie sets the ObjectReprNie property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprNie(val []string) {
	m.ObjectReprNie = val
}

// GetObjectReprNiew returns the ObjectReprNiew property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprNiew() []string {
	return m.ObjectReprNiew
}

// SetObjectReprNiew sets the ObjectReprNiew property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprNiew(val []string) {
	m.ObjectReprNiew = val
}

// GetObjectReprNisw returns the ObjectReprNisw property
func (m ExtrasObjectChangesListQueryParameters) GetObjectReprNisw() []string {
	return m.ObjectReprNisw
}

// SetObjectReprNisw sets the ObjectReprNisw property
func (m *ExtrasObjectChangesListQueryParameters) SetObjectReprNisw(val []string) {
	m.ObjectReprNisw = val
}

// GetOffset returns the Offset property
func (m ExtrasObjectChangesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasObjectChangesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasObjectChangesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasObjectChangesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasObjectChangesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasObjectChangesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRequestId returns the RequestId property
func (m ExtrasObjectChangesListQueryParameters) GetRequestId() string {
	return m.RequestId
}

// SetRequestId sets the RequestId property
func (m *ExtrasObjectChangesListQueryParameters) SetRequestId(val string) {
	m.RequestId = val
}

// GetTimeAfter returns the TimeAfter property
func (m ExtrasObjectChangesListQueryParameters) GetTimeAfter() time.Time {
	return m.TimeAfter
}

// SetTimeAfter sets the TimeAfter property
func (m *ExtrasObjectChangesListQueryParameters) SetTimeAfter(val time.Time) {
	m.TimeAfter = val
}

// GetTimeBefore returns the TimeBefore property
func (m ExtrasObjectChangesListQueryParameters) GetTimeBefore() time.Time {
	return m.TimeBefore
}

// SetTimeBefore sets the TimeBefore property
func (m *ExtrasObjectChangesListQueryParameters) SetTimeBefore(val time.Time) {
	m.TimeBefore = val
}

// GetUser returns the User property
func (m ExtrasObjectChangesListQueryParameters) GetUser() []string {
	return m.User
}

// SetUser sets the User property
func (m *ExtrasObjectChangesListQueryParameters) SetUser(val []string) {
	m.User = val
}

// GetUserN returns the UserN property
func (m ExtrasObjectChangesListQueryParameters) GetUserN() []string {
	return m.UserN
}

// SetUserN sets the UserN property
func (m *ExtrasObjectChangesListQueryParameters) SetUserN(val []string) {
	m.UserN = val
}

// GetUserId returns the UserId property
func (m ExtrasObjectChangesListQueryParameters) GetUserId() []*int32 {
	return m.UserId
}

// SetUserId sets the UserId property
func (m *ExtrasObjectChangesListQueryParameters) SetUserId(val []*int32) {
	m.UserId = val
}

// GetUserIdN returns the UserIdN property
func (m ExtrasObjectChangesListQueryParameters) GetUserIdN() []*int32 {
	return m.UserIdN
}

// SetUserIdN sets the UserIdN property
func (m *ExtrasObjectChangesListQueryParameters) SetUserIdN(val []*int32) {
	m.UserIdN = val
}

// GetUserName returns the UserName property
func (m ExtrasObjectChangesListQueryParameters) GetUserName() []string {
	return m.UserName
}

// SetUserName sets the UserName property
func (m *ExtrasObjectChangesListQueryParameters) SetUserName(val []string) {
	m.UserName = val
}

// GetUserNameEmpty returns the UserNameEmpty property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameEmpty() []string {
	return m.UserNameEmpty
}

// SetUserNameEmpty sets the UserNameEmpty property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameEmpty(val []string) {
	m.UserNameEmpty = val
}

// GetUserNameIc returns the UserNameIc property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameIc() []string {
	return m.UserNameIc
}

// SetUserNameIc sets the UserNameIc property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameIc(val []string) {
	m.UserNameIc = val
}

// GetUserNameIe returns the UserNameIe property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameIe() []string {
	return m.UserNameIe
}

// SetUserNameIe sets the UserNameIe property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameIe(val []string) {
	m.UserNameIe = val
}

// GetUserNameIew returns the UserNameIew property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameIew() []string {
	return m.UserNameIew
}

// SetUserNameIew sets the UserNameIew property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameIew(val []string) {
	m.UserNameIew = val
}

// GetUserNameIsw returns the UserNameIsw property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameIsw() []string {
	return m.UserNameIsw
}

// SetUserNameIsw sets the UserNameIsw property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameIsw(val []string) {
	m.UserNameIsw = val
}

// GetUserNameN returns the UserNameN property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameN() []string {
	return m.UserNameN
}

// SetUserNameN sets the UserNameN property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameN(val []string) {
	m.UserNameN = val
}

// GetUserNameNic returns the UserNameNic property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameNic() []string {
	return m.UserNameNic
}

// SetUserNameNic sets the UserNameNic property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameNic(val []string) {
	m.UserNameNic = val
}

// GetUserNameNie returns the UserNameNie property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameNie() []string {
	return m.UserNameNie
}

// SetUserNameNie sets the UserNameNie property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameNie(val []string) {
	m.UserNameNie = val
}

// GetUserNameNiew returns the UserNameNiew property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameNiew() []string {
	return m.UserNameNiew
}

// SetUserNameNiew sets the UserNameNiew property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameNiew(val []string) {
	m.UserNameNiew = val
}

// GetUserNameNisw returns the UserNameNisw property
func (m ExtrasObjectChangesListQueryParameters) GetUserNameNisw() []string {
	return m.UserNameNisw
}

// SetUserNameNisw sets the UserNameNisw property
func (m *ExtrasObjectChangesListQueryParameters) SetUserNameNisw(val []string) {
	m.UserNameNisw = val
}
