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

// UsersUsersListQueryParameters is an object.
type UsersUsersListQueryParameters struct {
	// Email:
	Email []string `json:"email,omitempty" mapstructure:"email,omitempty"`
	// EmailEmpty:
	EmailEmpty []string `json:"email__empty,omitempty" mapstructure:"email__empty,omitempty"`
	// EmailIc:
	EmailIc []string `json:"email__ic,omitempty" mapstructure:"email__ic,omitempty"`
	// EmailIe:
	EmailIe []string `json:"email__ie,omitempty" mapstructure:"email__ie,omitempty"`
	// EmailIew:
	EmailIew []string `json:"email__iew,omitempty" mapstructure:"email__iew,omitempty"`
	// EmailIsw:
	EmailIsw []string `json:"email__isw,omitempty" mapstructure:"email__isw,omitempty"`
	// EmailN:
	EmailN []string `json:"email__n,omitempty" mapstructure:"email__n,omitempty"`
	// EmailNic:
	EmailNic []string `json:"email__nic,omitempty" mapstructure:"email__nic,omitempty"`
	// EmailNie:
	EmailNie []string `json:"email__nie,omitempty" mapstructure:"email__nie,omitempty"`
	// EmailNiew:
	EmailNiew []string `json:"email__niew,omitempty" mapstructure:"email__niew,omitempty"`
	// EmailNisw:
	EmailNisw []string `json:"email__nisw,omitempty" mapstructure:"email__nisw,omitempty"`
	// FirstName:
	FirstName []string `json:"first_name,omitempty" mapstructure:"first_name,omitempty"`
	// FirstNameEmpty:
	FirstNameEmpty []string `json:"first_name__empty,omitempty" mapstructure:"first_name__empty,omitempty"`
	// FirstNameIc:
	FirstNameIc []string `json:"first_name__ic,omitempty" mapstructure:"first_name__ic,omitempty"`
	// FirstNameIe:
	FirstNameIe []string `json:"first_name__ie,omitempty" mapstructure:"first_name__ie,omitempty"`
	// FirstNameIew:
	FirstNameIew []string `json:"first_name__iew,omitempty" mapstructure:"first_name__iew,omitempty"`
	// FirstNameIsw:
	FirstNameIsw []string `json:"first_name__isw,omitempty" mapstructure:"first_name__isw,omitempty"`
	// FirstNameN:
	FirstNameN []string `json:"first_name__n,omitempty" mapstructure:"first_name__n,omitempty"`
	// FirstNameNic:
	FirstNameNic []string `json:"first_name__nic,omitempty" mapstructure:"first_name__nic,omitempty"`
	// FirstNameNie:
	FirstNameNie []string `json:"first_name__nie,omitempty" mapstructure:"first_name__nie,omitempty"`
	// FirstNameNiew:
	FirstNameNiew []string `json:"first_name__niew,omitempty" mapstructure:"first_name__niew,omitempty"`
	// FirstNameNisw:
	FirstNameNisw []string `json:"first_name__nisw,omitempty" mapstructure:"first_name__nisw,omitempty"`
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
	// IsActive:
	IsActive bool `json:"is_active,omitempty" mapstructure:"is_active,omitempty"`
	// IsStaff:
	IsStaff bool `json:"is_staff,omitempty" mapstructure:"is_staff,omitempty"`
	// LastName:
	LastName []string `json:"last_name,omitempty" mapstructure:"last_name,omitempty"`
	// LastNameEmpty:
	LastNameEmpty []string `json:"last_name__empty,omitempty" mapstructure:"last_name__empty,omitempty"`
	// LastNameIc:
	LastNameIc []string `json:"last_name__ic,omitempty" mapstructure:"last_name__ic,omitempty"`
	// LastNameIe:
	LastNameIe []string `json:"last_name__ie,omitempty" mapstructure:"last_name__ie,omitempty"`
	// LastNameIew:
	LastNameIew []string `json:"last_name__iew,omitempty" mapstructure:"last_name__iew,omitempty"`
	// LastNameIsw:
	LastNameIsw []string `json:"last_name__isw,omitempty" mapstructure:"last_name__isw,omitempty"`
	// LastNameN:
	LastNameN []string `json:"last_name__n,omitempty" mapstructure:"last_name__n,omitempty"`
	// LastNameNic:
	LastNameNic []string `json:"last_name__nic,omitempty" mapstructure:"last_name__nic,omitempty"`
	// LastNameNie:
	LastNameNie []string `json:"last_name__nie,omitempty" mapstructure:"last_name__nie,omitempty"`
	// LastNameNiew:
	LastNameNiew []string `json:"last_name__niew,omitempty" mapstructure:"last_name__niew,omitempty"`
	// LastNameNisw:
	LastNameNisw []string `json:"last_name__nisw,omitempty" mapstructure:"last_name__nisw,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Username:
	Username []string `json:"username,omitempty" mapstructure:"username,omitempty"`
	// UsernameEmpty:
	UsernameEmpty []string `json:"username__empty,omitempty" mapstructure:"username__empty,omitempty"`
	// UsernameIc:
	UsernameIc []string `json:"username__ic,omitempty" mapstructure:"username__ic,omitempty"`
	// UsernameIe:
	UsernameIe []string `json:"username__ie,omitempty" mapstructure:"username__ie,omitempty"`
	// UsernameIew:
	UsernameIew []string `json:"username__iew,omitempty" mapstructure:"username__iew,omitempty"`
	// UsernameIsw:
	UsernameIsw []string `json:"username__isw,omitempty" mapstructure:"username__isw,omitempty"`
	// UsernameN:
	UsernameN []string `json:"username__n,omitempty" mapstructure:"username__n,omitempty"`
	// UsernameNic:
	UsernameNic []string `json:"username__nic,omitempty" mapstructure:"username__nic,omitempty"`
	// UsernameNie:
	UsernameNie []string `json:"username__nie,omitempty" mapstructure:"username__nie,omitempty"`
	// UsernameNiew:
	UsernameNiew []string `json:"username__niew,omitempty" mapstructure:"username__niew,omitempty"`
	// UsernameNisw:
	UsernameNisw []string `json:"username__nisw,omitempty" mapstructure:"username__nisw,omitempty"`
}

// Validate implements basic validation for this model
func (m UsersUsersListQueryParameters) Validate() error {
	return validation.Errors{
		"email": validation.Validate(
			m.Email,
		),
		"emailEmpty": validation.Validate(
			m.EmailEmpty,
		),
		"emailIc": validation.Validate(
			m.EmailIc,
		),
		"emailIe": validation.Validate(
			m.EmailIe,
		),
		"emailIew": validation.Validate(
			m.EmailIew,
		),
		"emailIsw": validation.Validate(
			m.EmailIsw,
		),
		"emailN": validation.Validate(
			m.EmailN,
		),
		"emailNic": validation.Validate(
			m.EmailNic,
		),
		"emailNie": validation.Validate(
			m.EmailNie,
		),
		"emailNiew": validation.Validate(
			m.EmailNiew,
		),
		"emailNisw": validation.Validate(
			m.EmailNisw,
		),
		"firstName": validation.Validate(
			m.FirstName,
		),
		"firstNameEmpty": validation.Validate(
			m.FirstNameEmpty,
		),
		"firstNameIc": validation.Validate(
			m.FirstNameIc,
		),
		"firstNameIe": validation.Validate(
			m.FirstNameIe,
		),
		"firstNameIew": validation.Validate(
			m.FirstNameIew,
		),
		"firstNameIsw": validation.Validate(
			m.FirstNameIsw,
		),
		"firstNameN": validation.Validate(
			m.FirstNameN,
		),
		"firstNameNic": validation.Validate(
			m.FirstNameNic,
		),
		"firstNameNie": validation.Validate(
			m.FirstNameNie,
		),
		"firstNameNiew": validation.Validate(
			m.FirstNameNiew,
		),
		"firstNameNisw": validation.Validate(
			m.FirstNameNisw,
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
		"lastName": validation.Validate(
			m.LastName,
		),
		"lastNameEmpty": validation.Validate(
			m.LastNameEmpty,
		),
		"lastNameIc": validation.Validate(
			m.LastNameIc,
		),
		"lastNameIe": validation.Validate(
			m.LastNameIe,
		),
		"lastNameIew": validation.Validate(
			m.LastNameIew,
		),
		"lastNameIsw": validation.Validate(
			m.LastNameIsw,
		),
		"lastNameN": validation.Validate(
			m.LastNameN,
		),
		"lastNameNic": validation.Validate(
			m.LastNameNic,
		),
		"lastNameNie": validation.Validate(
			m.LastNameNie,
		),
		"lastNameNiew": validation.Validate(
			m.LastNameNiew,
		),
		"lastNameNisw": validation.Validate(
			m.LastNameNisw,
		),
		"username": validation.Validate(
			m.Username,
		),
		"usernameEmpty": validation.Validate(
			m.UsernameEmpty,
		),
		"usernameIc": validation.Validate(
			m.UsernameIc,
		),
		"usernameIe": validation.Validate(
			m.UsernameIe,
		),
		"usernameIew": validation.Validate(
			m.UsernameIew,
		),
		"usernameIsw": validation.Validate(
			m.UsernameIsw,
		),
		"usernameN": validation.Validate(
			m.UsernameN,
		),
		"usernameNic": validation.Validate(
			m.UsernameNic,
		),
		"usernameNie": validation.Validate(
			m.UsernameNie,
		),
		"usernameNiew": validation.Validate(
			m.UsernameNiew,
		),
		"usernameNisw": validation.Validate(
			m.UsernameNisw,
		),
	}.Filter()
}

// GetEmail returns the Email property
func (m UsersUsersListQueryParameters) GetEmail() []string {
	return m.Email
}

// SetEmail sets the Email property
func (m *UsersUsersListQueryParameters) SetEmail(val []string) {
	m.Email = val
}

// GetEmailEmpty returns the EmailEmpty property
func (m UsersUsersListQueryParameters) GetEmailEmpty() []string {
	return m.EmailEmpty
}

// SetEmailEmpty sets the EmailEmpty property
func (m *UsersUsersListQueryParameters) SetEmailEmpty(val []string) {
	m.EmailEmpty = val
}

// GetEmailIc returns the EmailIc property
func (m UsersUsersListQueryParameters) GetEmailIc() []string {
	return m.EmailIc
}

// SetEmailIc sets the EmailIc property
func (m *UsersUsersListQueryParameters) SetEmailIc(val []string) {
	m.EmailIc = val
}

// GetEmailIe returns the EmailIe property
func (m UsersUsersListQueryParameters) GetEmailIe() []string {
	return m.EmailIe
}

// SetEmailIe sets the EmailIe property
func (m *UsersUsersListQueryParameters) SetEmailIe(val []string) {
	m.EmailIe = val
}

// GetEmailIew returns the EmailIew property
func (m UsersUsersListQueryParameters) GetEmailIew() []string {
	return m.EmailIew
}

// SetEmailIew sets the EmailIew property
func (m *UsersUsersListQueryParameters) SetEmailIew(val []string) {
	m.EmailIew = val
}

// GetEmailIsw returns the EmailIsw property
func (m UsersUsersListQueryParameters) GetEmailIsw() []string {
	return m.EmailIsw
}

// SetEmailIsw sets the EmailIsw property
func (m *UsersUsersListQueryParameters) SetEmailIsw(val []string) {
	m.EmailIsw = val
}

// GetEmailN returns the EmailN property
func (m UsersUsersListQueryParameters) GetEmailN() []string {
	return m.EmailN
}

// SetEmailN sets the EmailN property
func (m *UsersUsersListQueryParameters) SetEmailN(val []string) {
	m.EmailN = val
}

// GetEmailNic returns the EmailNic property
func (m UsersUsersListQueryParameters) GetEmailNic() []string {
	return m.EmailNic
}

// SetEmailNic sets the EmailNic property
func (m *UsersUsersListQueryParameters) SetEmailNic(val []string) {
	m.EmailNic = val
}

// GetEmailNie returns the EmailNie property
func (m UsersUsersListQueryParameters) GetEmailNie() []string {
	return m.EmailNie
}

// SetEmailNie sets the EmailNie property
func (m *UsersUsersListQueryParameters) SetEmailNie(val []string) {
	m.EmailNie = val
}

// GetEmailNiew returns the EmailNiew property
func (m UsersUsersListQueryParameters) GetEmailNiew() []string {
	return m.EmailNiew
}

// SetEmailNiew sets the EmailNiew property
func (m *UsersUsersListQueryParameters) SetEmailNiew(val []string) {
	m.EmailNiew = val
}

// GetEmailNisw returns the EmailNisw property
func (m UsersUsersListQueryParameters) GetEmailNisw() []string {
	return m.EmailNisw
}

// SetEmailNisw sets the EmailNisw property
func (m *UsersUsersListQueryParameters) SetEmailNisw(val []string) {
	m.EmailNisw = val
}

// GetFirstName returns the FirstName property
func (m UsersUsersListQueryParameters) GetFirstName() []string {
	return m.FirstName
}

// SetFirstName sets the FirstName property
func (m *UsersUsersListQueryParameters) SetFirstName(val []string) {
	m.FirstName = val
}

// GetFirstNameEmpty returns the FirstNameEmpty property
func (m UsersUsersListQueryParameters) GetFirstNameEmpty() []string {
	return m.FirstNameEmpty
}

// SetFirstNameEmpty sets the FirstNameEmpty property
func (m *UsersUsersListQueryParameters) SetFirstNameEmpty(val []string) {
	m.FirstNameEmpty = val
}

// GetFirstNameIc returns the FirstNameIc property
func (m UsersUsersListQueryParameters) GetFirstNameIc() []string {
	return m.FirstNameIc
}

// SetFirstNameIc sets the FirstNameIc property
func (m *UsersUsersListQueryParameters) SetFirstNameIc(val []string) {
	m.FirstNameIc = val
}

// GetFirstNameIe returns the FirstNameIe property
func (m UsersUsersListQueryParameters) GetFirstNameIe() []string {
	return m.FirstNameIe
}

// SetFirstNameIe sets the FirstNameIe property
func (m *UsersUsersListQueryParameters) SetFirstNameIe(val []string) {
	m.FirstNameIe = val
}

// GetFirstNameIew returns the FirstNameIew property
func (m UsersUsersListQueryParameters) GetFirstNameIew() []string {
	return m.FirstNameIew
}

// SetFirstNameIew sets the FirstNameIew property
func (m *UsersUsersListQueryParameters) SetFirstNameIew(val []string) {
	m.FirstNameIew = val
}

// GetFirstNameIsw returns the FirstNameIsw property
func (m UsersUsersListQueryParameters) GetFirstNameIsw() []string {
	return m.FirstNameIsw
}

// SetFirstNameIsw sets the FirstNameIsw property
func (m *UsersUsersListQueryParameters) SetFirstNameIsw(val []string) {
	m.FirstNameIsw = val
}

// GetFirstNameN returns the FirstNameN property
func (m UsersUsersListQueryParameters) GetFirstNameN() []string {
	return m.FirstNameN
}

// SetFirstNameN sets the FirstNameN property
func (m *UsersUsersListQueryParameters) SetFirstNameN(val []string) {
	m.FirstNameN = val
}

// GetFirstNameNic returns the FirstNameNic property
func (m UsersUsersListQueryParameters) GetFirstNameNic() []string {
	return m.FirstNameNic
}

// SetFirstNameNic sets the FirstNameNic property
func (m *UsersUsersListQueryParameters) SetFirstNameNic(val []string) {
	m.FirstNameNic = val
}

// GetFirstNameNie returns the FirstNameNie property
func (m UsersUsersListQueryParameters) GetFirstNameNie() []string {
	return m.FirstNameNie
}

// SetFirstNameNie sets the FirstNameNie property
func (m *UsersUsersListQueryParameters) SetFirstNameNie(val []string) {
	m.FirstNameNie = val
}

// GetFirstNameNiew returns the FirstNameNiew property
func (m UsersUsersListQueryParameters) GetFirstNameNiew() []string {
	return m.FirstNameNiew
}

// SetFirstNameNiew sets the FirstNameNiew property
func (m *UsersUsersListQueryParameters) SetFirstNameNiew(val []string) {
	m.FirstNameNiew = val
}

// GetFirstNameNisw returns the FirstNameNisw property
func (m UsersUsersListQueryParameters) GetFirstNameNisw() []string {
	return m.FirstNameNisw
}

// SetFirstNameNisw sets the FirstNameNisw property
func (m *UsersUsersListQueryParameters) SetFirstNameNisw(val []string) {
	m.FirstNameNisw = val
}

// GetGroup returns the Group property
func (m UsersUsersListQueryParameters) GetGroup() []string {
	return m.Group
}

// SetGroup sets the Group property
func (m *UsersUsersListQueryParameters) SetGroup(val []string) {
	m.Group = val
}

// GetGroupN returns the GroupN property
func (m UsersUsersListQueryParameters) GetGroupN() []string {
	return m.GroupN
}

// SetGroupN sets the GroupN property
func (m *UsersUsersListQueryParameters) SetGroupN(val []string) {
	m.GroupN = val
}

// GetGroupId returns the GroupId property
func (m UsersUsersListQueryParameters) GetGroupId() []int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *UsersUsersListQueryParameters) SetGroupId(val []int32) {
	m.GroupId = val
}

// GetGroupIdN returns the GroupIdN property
func (m UsersUsersListQueryParameters) GetGroupIdN() []int32 {
	return m.GroupIdN
}

// SetGroupIdN sets the GroupIdN property
func (m *UsersUsersListQueryParameters) SetGroupIdN(val []int32) {
	m.GroupIdN = val
}

// GetId returns the Id property
func (m UsersUsersListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *UsersUsersListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m UsersUsersListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *UsersUsersListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m UsersUsersListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *UsersUsersListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m UsersUsersListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *UsersUsersListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m UsersUsersListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *UsersUsersListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m UsersUsersListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *UsersUsersListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetIsActive returns the IsActive property
func (m UsersUsersListQueryParameters) GetIsActive() bool {
	return m.IsActive
}

// SetIsActive sets the IsActive property
func (m *UsersUsersListQueryParameters) SetIsActive(val bool) {
	m.IsActive = val
}

// GetIsStaff returns the IsStaff property
func (m UsersUsersListQueryParameters) GetIsStaff() bool {
	return m.IsStaff
}

// SetIsStaff sets the IsStaff property
func (m *UsersUsersListQueryParameters) SetIsStaff(val bool) {
	m.IsStaff = val
}

// GetLastName returns the LastName property
func (m UsersUsersListQueryParameters) GetLastName() []string {
	return m.LastName
}

// SetLastName sets the LastName property
func (m *UsersUsersListQueryParameters) SetLastName(val []string) {
	m.LastName = val
}

// GetLastNameEmpty returns the LastNameEmpty property
func (m UsersUsersListQueryParameters) GetLastNameEmpty() []string {
	return m.LastNameEmpty
}

// SetLastNameEmpty sets the LastNameEmpty property
func (m *UsersUsersListQueryParameters) SetLastNameEmpty(val []string) {
	m.LastNameEmpty = val
}

// GetLastNameIc returns the LastNameIc property
func (m UsersUsersListQueryParameters) GetLastNameIc() []string {
	return m.LastNameIc
}

// SetLastNameIc sets the LastNameIc property
func (m *UsersUsersListQueryParameters) SetLastNameIc(val []string) {
	m.LastNameIc = val
}

// GetLastNameIe returns the LastNameIe property
func (m UsersUsersListQueryParameters) GetLastNameIe() []string {
	return m.LastNameIe
}

// SetLastNameIe sets the LastNameIe property
func (m *UsersUsersListQueryParameters) SetLastNameIe(val []string) {
	m.LastNameIe = val
}

// GetLastNameIew returns the LastNameIew property
func (m UsersUsersListQueryParameters) GetLastNameIew() []string {
	return m.LastNameIew
}

// SetLastNameIew sets the LastNameIew property
func (m *UsersUsersListQueryParameters) SetLastNameIew(val []string) {
	m.LastNameIew = val
}

// GetLastNameIsw returns the LastNameIsw property
func (m UsersUsersListQueryParameters) GetLastNameIsw() []string {
	return m.LastNameIsw
}

// SetLastNameIsw sets the LastNameIsw property
func (m *UsersUsersListQueryParameters) SetLastNameIsw(val []string) {
	m.LastNameIsw = val
}

// GetLastNameN returns the LastNameN property
func (m UsersUsersListQueryParameters) GetLastNameN() []string {
	return m.LastNameN
}

// SetLastNameN sets the LastNameN property
func (m *UsersUsersListQueryParameters) SetLastNameN(val []string) {
	m.LastNameN = val
}

// GetLastNameNic returns the LastNameNic property
func (m UsersUsersListQueryParameters) GetLastNameNic() []string {
	return m.LastNameNic
}

// SetLastNameNic sets the LastNameNic property
func (m *UsersUsersListQueryParameters) SetLastNameNic(val []string) {
	m.LastNameNic = val
}

// GetLastNameNie returns the LastNameNie property
func (m UsersUsersListQueryParameters) GetLastNameNie() []string {
	return m.LastNameNie
}

// SetLastNameNie sets the LastNameNie property
func (m *UsersUsersListQueryParameters) SetLastNameNie(val []string) {
	m.LastNameNie = val
}

// GetLastNameNiew returns the LastNameNiew property
func (m UsersUsersListQueryParameters) GetLastNameNiew() []string {
	return m.LastNameNiew
}

// SetLastNameNiew sets the LastNameNiew property
func (m *UsersUsersListQueryParameters) SetLastNameNiew(val []string) {
	m.LastNameNiew = val
}

// GetLastNameNisw returns the LastNameNisw property
func (m UsersUsersListQueryParameters) GetLastNameNisw() []string {
	return m.LastNameNisw
}

// SetLastNameNisw sets the LastNameNisw property
func (m *UsersUsersListQueryParameters) SetLastNameNisw(val []string) {
	m.LastNameNisw = val
}

// GetLimit returns the Limit property
func (m UsersUsersListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *UsersUsersListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m UsersUsersListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *UsersUsersListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m UsersUsersListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *UsersUsersListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m UsersUsersListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *UsersUsersListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetUsername returns the Username property
func (m UsersUsersListQueryParameters) GetUsername() []string {
	return m.Username
}

// SetUsername sets the Username property
func (m *UsersUsersListQueryParameters) SetUsername(val []string) {
	m.Username = val
}

// GetUsernameEmpty returns the UsernameEmpty property
func (m UsersUsersListQueryParameters) GetUsernameEmpty() []string {
	return m.UsernameEmpty
}

// SetUsernameEmpty sets the UsernameEmpty property
func (m *UsersUsersListQueryParameters) SetUsernameEmpty(val []string) {
	m.UsernameEmpty = val
}

// GetUsernameIc returns the UsernameIc property
func (m UsersUsersListQueryParameters) GetUsernameIc() []string {
	return m.UsernameIc
}

// SetUsernameIc sets the UsernameIc property
func (m *UsersUsersListQueryParameters) SetUsernameIc(val []string) {
	m.UsernameIc = val
}

// GetUsernameIe returns the UsernameIe property
func (m UsersUsersListQueryParameters) GetUsernameIe() []string {
	return m.UsernameIe
}

// SetUsernameIe sets the UsernameIe property
func (m *UsersUsersListQueryParameters) SetUsernameIe(val []string) {
	m.UsernameIe = val
}

// GetUsernameIew returns the UsernameIew property
func (m UsersUsersListQueryParameters) GetUsernameIew() []string {
	return m.UsernameIew
}

// SetUsernameIew sets the UsernameIew property
func (m *UsersUsersListQueryParameters) SetUsernameIew(val []string) {
	m.UsernameIew = val
}

// GetUsernameIsw returns the UsernameIsw property
func (m UsersUsersListQueryParameters) GetUsernameIsw() []string {
	return m.UsernameIsw
}

// SetUsernameIsw sets the UsernameIsw property
func (m *UsersUsersListQueryParameters) SetUsernameIsw(val []string) {
	m.UsernameIsw = val
}

// GetUsernameN returns the UsernameN property
func (m UsersUsersListQueryParameters) GetUsernameN() []string {
	return m.UsernameN
}

// SetUsernameN sets the UsernameN property
func (m *UsersUsersListQueryParameters) SetUsernameN(val []string) {
	m.UsernameN = val
}

// GetUsernameNic returns the UsernameNic property
func (m UsersUsersListQueryParameters) GetUsernameNic() []string {
	return m.UsernameNic
}

// SetUsernameNic sets the UsernameNic property
func (m *UsersUsersListQueryParameters) SetUsernameNic(val []string) {
	m.UsernameNic = val
}

// GetUsernameNie returns the UsernameNie property
func (m UsersUsersListQueryParameters) GetUsernameNie() []string {
	return m.UsernameNie
}

// SetUsernameNie sets the UsernameNie property
func (m *UsersUsersListQueryParameters) SetUsernameNie(val []string) {
	m.UsernameNie = val
}

// GetUsernameNiew returns the UsernameNiew property
func (m UsersUsersListQueryParameters) GetUsernameNiew() []string {
	return m.UsernameNiew
}

// SetUsernameNiew sets the UsernameNiew property
func (m *UsersUsersListQueryParameters) SetUsernameNiew(val []string) {
	m.UsernameNiew = val
}

// GetUsernameNisw returns the UsernameNisw property
func (m UsersUsersListQueryParameters) GetUsernameNisw() []string {
	return m.UsernameNisw
}

// SetUsernameNisw sets the UsernameNisw property
func (m *UsersUsersListQueryParameters) SetUsernameNisw(val []string) {
	m.UsernameNisw = val
}
