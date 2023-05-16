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

// UsersPermissionsListQueryParameters is an object.
type UsersPermissionsListQueryParameters struct {
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
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Group: Group (name)
	Group []string `json:"group,omitempty" mapstructure:"group,omitempty"`
	// GroupN: Group (name)
	GroupN []string `json:"group__n,omitempty" mapstructure:"group__n,omitempty"`
	// GroupId: Group
	GroupId []int32 `json:"group_id,omitempty" mapstructure:"group_id,omitempty"`
	// GroupIdN: Group
	GroupIdN []int32 `json:"group_id__n,omitempty" mapstructure:"group_id__n,omitempty"`
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
	// ObjectTypes:
	ObjectTypes []int32 `json:"object_types,omitempty" mapstructure:"object_types,omitempty"`
	// ObjectTypesN:
	ObjectTypesN []int32 `json:"object_types__n,omitempty" mapstructure:"object_types__n,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// User: User (name)
	User []string `json:"user,omitempty" mapstructure:"user,omitempty"`
	// UserN: User (name)
	UserN []string `json:"user__n,omitempty" mapstructure:"user__n,omitempty"`
	// UserId: User
	UserId []int32 `json:"user_id,omitempty" mapstructure:"user_id,omitempty"`
	// UserIdN: User
	UserIdN []int32 `json:"user_id__n,omitempty" mapstructure:"user_id__n,omitempty"`
}

// Validate implements basic validation for this model
func (m UsersPermissionsListQueryParameters) Validate() error {
	return validation.Errors{
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
		"group": validation.Validate(
			m.Group,
		),
		"groupN": validation.Validate(
			m.GroupN,
		),
		"groupId": validation.Validate(
			m.GroupId,
		),
		"groupIdN": validation.Validate(
			m.GroupIdN,
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
		"objectTypes": validation.Validate(
			m.ObjectTypes,
		),
		"objectTypesN": validation.Validate(
			m.ObjectTypesN,
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
	}.Filter()
}

// GetDescription returns the Description property
func (m UsersPermissionsListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *UsersPermissionsListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m UsersPermissionsListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *UsersPermissionsListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m UsersPermissionsListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *UsersPermissionsListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m UsersPermissionsListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *UsersPermissionsListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m UsersPermissionsListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *UsersPermissionsListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m UsersPermissionsListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *UsersPermissionsListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m UsersPermissionsListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *UsersPermissionsListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m UsersPermissionsListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *UsersPermissionsListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m UsersPermissionsListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *UsersPermissionsListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m UsersPermissionsListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *UsersPermissionsListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m UsersPermissionsListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *UsersPermissionsListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetEnabled returns the Enabled property
func (m UsersPermissionsListQueryParameters) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *UsersPermissionsListQueryParameters) SetEnabled(val bool) {
	m.Enabled = val
}

// GetGroup returns the Group property
func (m UsersPermissionsListQueryParameters) GetGroup() []string {
	return m.Group
}

// SetGroup sets the Group property
func (m *UsersPermissionsListQueryParameters) SetGroup(val []string) {
	m.Group = val
}

// GetGroupN returns the GroupN property
func (m UsersPermissionsListQueryParameters) GetGroupN() []string {
	return m.GroupN
}

// SetGroupN sets the GroupN property
func (m *UsersPermissionsListQueryParameters) SetGroupN(val []string) {
	m.GroupN = val
}

// GetGroupId returns the GroupId property
func (m UsersPermissionsListQueryParameters) GetGroupId() []int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *UsersPermissionsListQueryParameters) SetGroupId(val []int32) {
	m.GroupId = val
}

// GetGroupIdN returns the GroupIdN property
func (m UsersPermissionsListQueryParameters) GetGroupIdN() []int32 {
	return m.GroupIdN
}

// SetGroupIdN sets the GroupIdN property
func (m *UsersPermissionsListQueryParameters) SetGroupIdN(val []int32) {
	m.GroupIdN = val
}

// GetId returns the Id property
func (m UsersPermissionsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *UsersPermissionsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m UsersPermissionsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *UsersPermissionsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m UsersPermissionsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *UsersPermissionsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m UsersPermissionsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *UsersPermissionsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m UsersPermissionsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *UsersPermissionsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m UsersPermissionsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *UsersPermissionsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m UsersPermissionsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *UsersPermissionsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m UsersPermissionsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *UsersPermissionsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m UsersPermissionsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *UsersPermissionsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m UsersPermissionsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *UsersPermissionsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m UsersPermissionsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *UsersPermissionsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m UsersPermissionsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *UsersPermissionsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m UsersPermissionsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *UsersPermissionsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m UsersPermissionsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *UsersPermissionsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m UsersPermissionsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *UsersPermissionsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m UsersPermissionsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *UsersPermissionsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m UsersPermissionsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *UsersPermissionsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m UsersPermissionsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *UsersPermissionsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetObjectTypes returns the ObjectTypes property
func (m UsersPermissionsListQueryParameters) GetObjectTypes() []int32 {
	return m.ObjectTypes
}

// SetObjectTypes sets the ObjectTypes property
func (m *UsersPermissionsListQueryParameters) SetObjectTypes(val []int32) {
	m.ObjectTypes = val
}

// GetObjectTypesN returns the ObjectTypesN property
func (m UsersPermissionsListQueryParameters) GetObjectTypesN() []int32 {
	return m.ObjectTypesN
}

// SetObjectTypesN sets the ObjectTypesN property
func (m *UsersPermissionsListQueryParameters) SetObjectTypesN(val []int32) {
	m.ObjectTypesN = val
}

// GetOffset returns the Offset property
func (m UsersPermissionsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *UsersPermissionsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m UsersPermissionsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *UsersPermissionsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m UsersPermissionsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *UsersPermissionsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetUser returns the User property
func (m UsersPermissionsListQueryParameters) GetUser() []string {
	return m.User
}

// SetUser sets the User property
func (m *UsersPermissionsListQueryParameters) SetUser(val []string) {
	m.User = val
}

// GetUserN returns the UserN property
func (m UsersPermissionsListQueryParameters) GetUserN() []string {
	return m.UserN
}

// SetUserN sets the UserN property
func (m *UsersPermissionsListQueryParameters) SetUserN(val []string) {
	m.UserN = val
}

// GetUserId returns the UserId property
func (m UsersPermissionsListQueryParameters) GetUserId() []int32 {
	return m.UserId
}

// SetUserId sets the UserId property
func (m *UsersPermissionsListQueryParameters) SetUserId(val []int32) {
	m.UserId = val
}

// GetUserIdN returns the UserIdN property
func (m UsersPermissionsListQueryParameters) GetUserIdN() []int32 {
	return m.UserIdN
}

// SetUserIdN sets the UserIdN property
func (m *UsersPermissionsListQueryParameters) SetUserIdN(val []int32) {
	m.UserIdN = val
}
