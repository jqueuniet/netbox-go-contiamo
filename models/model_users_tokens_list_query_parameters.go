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

// UsersTokensListQueryParameters is an object.
type UsersTokensListQueryParameters struct {
	// Created:
	Created time.Time `json:"created" mapstructure:"created"`
	// CreatedGte:
	CreatedGte time.Time `json:"created__gte" mapstructure:"created__gte"`
	// CreatedLte:
	CreatedLte time.Time `json:"created__lte" mapstructure:"created__lte"`
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
	// Expires:
	Expires time.Time `json:"expires" mapstructure:"expires"`
	// ExpiresGte:
	ExpiresGte time.Time `json:"expires__gte" mapstructure:"expires__gte"`
	// ExpiresLte:
	ExpiresLte time.Time `json:"expires__lte" mapstructure:"expires__lte"`
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
	// Key:
	Key []string `json:"key,omitempty" mapstructure:"key,omitempty"`
	// KeyEmpty:
	KeyEmpty []string `json:"key__empty,omitempty" mapstructure:"key__empty,omitempty"`
	// KeyIc:
	KeyIc []string `json:"key__ic,omitempty" mapstructure:"key__ic,omitempty"`
	// KeyIe:
	KeyIe []string `json:"key__ie,omitempty" mapstructure:"key__ie,omitempty"`
	// KeyIew:
	KeyIew []string `json:"key__iew,omitempty" mapstructure:"key__iew,omitempty"`
	// KeyIsw:
	KeyIsw []string `json:"key__isw,omitempty" mapstructure:"key__isw,omitempty"`
	// KeyN:
	KeyN []string `json:"key__n,omitempty" mapstructure:"key__n,omitempty"`
	// KeyNic:
	KeyNic []string `json:"key__nic,omitempty" mapstructure:"key__nic,omitempty"`
	// KeyNie:
	KeyNie []string `json:"key__nie,omitempty" mapstructure:"key__nie,omitempty"`
	// KeyNiew:
	KeyNiew []string `json:"key__niew,omitempty" mapstructure:"key__niew,omitempty"`
	// KeyNisw:
	KeyNisw []string `json:"key__nisw,omitempty" mapstructure:"key__nisw,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
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
	// WriteEnabled:
	WriteEnabled bool `json:"write_enabled,omitempty" mapstructure:"write_enabled,omitempty"`
}

// Validate implements basic validation for this model
func (m UsersTokensListQueryParameters) Validate() error {
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
		"key": validation.Validate(
			m.Key,
		),
		"keyEmpty": validation.Validate(
			m.KeyEmpty,
		),
		"keyIc": validation.Validate(
			m.KeyIc,
		),
		"keyIe": validation.Validate(
			m.KeyIe,
		),
		"keyIew": validation.Validate(
			m.KeyIew,
		),
		"keyIsw": validation.Validate(
			m.KeyIsw,
		),
		"keyN": validation.Validate(
			m.KeyN,
		),
		"keyNic": validation.Validate(
			m.KeyNic,
		),
		"keyNie": validation.Validate(
			m.KeyNie,
		),
		"keyNiew": validation.Validate(
			m.KeyNiew,
		),
		"keyNisw": validation.Validate(
			m.KeyNisw,
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

// GetCreated returns the Created property
func (m UsersTokensListQueryParameters) GetCreated() time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *UsersTokensListQueryParameters) SetCreated(val time.Time) {
	m.Created = val
}

// GetCreatedGte returns the CreatedGte property
func (m UsersTokensListQueryParameters) GetCreatedGte() time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *UsersTokensListQueryParameters) SetCreatedGte(val time.Time) {
	m.CreatedGte = val
}

// GetCreatedLte returns the CreatedLte property
func (m UsersTokensListQueryParameters) GetCreatedLte() time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *UsersTokensListQueryParameters) SetCreatedLte(val time.Time) {
	m.CreatedLte = val
}

// GetDescription returns the Description property
func (m UsersTokensListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *UsersTokensListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m UsersTokensListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *UsersTokensListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m UsersTokensListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *UsersTokensListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m UsersTokensListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *UsersTokensListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m UsersTokensListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *UsersTokensListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m UsersTokensListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *UsersTokensListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m UsersTokensListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *UsersTokensListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m UsersTokensListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *UsersTokensListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m UsersTokensListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *UsersTokensListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m UsersTokensListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *UsersTokensListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m UsersTokensListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *UsersTokensListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetExpires returns the Expires property
func (m UsersTokensListQueryParameters) GetExpires() time.Time {
	return m.Expires
}

// SetExpires sets the Expires property
func (m *UsersTokensListQueryParameters) SetExpires(val time.Time) {
	m.Expires = val
}

// GetExpiresGte returns the ExpiresGte property
func (m UsersTokensListQueryParameters) GetExpiresGte() time.Time {
	return m.ExpiresGte
}

// SetExpiresGte sets the ExpiresGte property
func (m *UsersTokensListQueryParameters) SetExpiresGte(val time.Time) {
	m.ExpiresGte = val
}

// GetExpiresLte returns the ExpiresLte property
func (m UsersTokensListQueryParameters) GetExpiresLte() time.Time {
	return m.ExpiresLte
}

// SetExpiresLte sets the ExpiresLte property
func (m *UsersTokensListQueryParameters) SetExpiresLte(val time.Time) {
	m.ExpiresLte = val
}

// GetId returns the Id property
func (m UsersTokensListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *UsersTokensListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m UsersTokensListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *UsersTokensListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m UsersTokensListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *UsersTokensListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m UsersTokensListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *UsersTokensListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m UsersTokensListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *UsersTokensListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m UsersTokensListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *UsersTokensListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetKey returns the Key property
func (m UsersTokensListQueryParameters) GetKey() []string {
	return m.Key
}

// SetKey sets the Key property
func (m *UsersTokensListQueryParameters) SetKey(val []string) {
	m.Key = val
}

// GetKeyEmpty returns the KeyEmpty property
func (m UsersTokensListQueryParameters) GetKeyEmpty() []string {
	return m.KeyEmpty
}

// SetKeyEmpty sets the KeyEmpty property
func (m *UsersTokensListQueryParameters) SetKeyEmpty(val []string) {
	m.KeyEmpty = val
}

// GetKeyIc returns the KeyIc property
func (m UsersTokensListQueryParameters) GetKeyIc() []string {
	return m.KeyIc
}

// SetKeyIc sets the KeyIc property
func (m *UsersTokensListQueryParameters) SetKeyIc(val []string) {
	m.KeyIc = val
}

// GetKeyIe returns the KeyIe property
func (m UsersTokensListQueryParameters) GetKeyIe() []string {
	return m.KeyIe
}

// SetKeyIe sets the KeyIe property
func (m *UsersTokensListQueryParameters) SetKeyIe(val []string) {
	m.KeyIe = val
}

// GetKeyIew returns the KeyIew property
func (m UsersTokensListQueryParameters) GetKeyIew() []string {
	return m.KeyIew
}

// SetKeyIew sets the KeyIew property
func (m *UsersTokensListQueryParameters) SetKeyIew(val []string) {
	m.KeyIew = val
}

// GetKeyIsw returns the KeyIsw property
func (m UsersTokensListQueryParameters) GetKeyIsw() []string {
	return m.KeyIsw
}

// SetKeyIsw sets the KeyIsw property
func (m *UsersTokensListQueryParameters) SetKeyIsw(val []string) {
	m.KeyIsw = val
}

// GetKeyN returns the KeyN property
func (m UsersTokensListQueryParameters) GetKeyN() []string {
	return m.KeyN
}

// SetKeyN sets the KeyN property
func (m *UsersTokensListQueryParameters) SetKeyN(val []string) {
	m.KeyN = val
}

// GetKeyNic returns the KeyNic property
func (m UsersTokensListQueryParameters) GetKeyNic() []string {
	return m.KeyNic
}

// SetKeyNic sets the KeyNic property
func (m *UsersTokensListQueryParameters) SetKeyNic(val []string) {
	m.KeyNic = val
}

// GetKeyNie returns the KeyNie property
func (m UsersTokensListQueryParameters) GetKeyNie() []string {
	return m.KeyNie
}

// SetKeyNie sets the KeyNie property
func (m *UsersTokensListQueryParameters) SetKeyNie(val []string) {
	m.KeyNie = val
}

// GetKeyNiew returns the KeyNiew property
func (m UsersTokensListQueryParameters) GetKeyNiew() []string {
	return m.KeyNiew
}

// SetKeyNiew sets the KeyNiew property
func (m *UsersTokensListQueryParameters) SetKeyNiew(val []string) {
	m.KeyNiew = val
}

// GetKeyNisw returns the KeyNisw property
func (m UsersTokensListQueryParameters) GetKeyNisw() []string {
	return m.KeyNisw
}

// SetKeyNisw sets the KeyNisw property
func (m *UsersTokensListQueryParameters) SetKeyNisw(val []string) {
	m.KeyNisw = val
}

// GetLimit returns the Limit property
func (m UsersTokensListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *UsersTokensListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m UsersTokensListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *UsersTokensListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m UsersTokensListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *UsersTokensListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m UsersTokensListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *UsersTokensListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetUser returns the User property
func (m UsersTokensListQueryParameters) GetUser() []string {
	return m.User
}

// SetUser sets the User property
func (m *UsersTokensListQueryParameters) SetUser(val []string) {
	m.User = val
}

// GetUserN returns the UserN property
func (m UsersTokensListQueryParameters) GetUserN() []string {
	return m.UserN
}

// SetUserN sets the UserN property
func (m *UsersTokensListQueryParameters) SetUserN(val []string) {
	m.UserN = val
}

// GetUserId returns the UserId property
func (m UsersTokensListQueryParameters) GetUserId() []int32 {
	return m.UserId
}

// SetUserId sets the UserId property
func (m *UsersTokensListQueryParameters) SetUserId(val []int32) {
	m.UserId = val
}

// GetUserIdN returns the UserIdN property
func (m UsersTokensListQueryParameters) GetUserIdN() []int32 {
	return m.UserIdN
}

// SetUserIdN sets the UserIdN property
func (m *UsersTokensListQueryParameters) SetUserIdN(val []int32) {
	m.UserIdN = val
}

// GetWriteEnabled returns the WriteEnabled property
func (m UsersTokensListQueryParameters) GetWriteEnabled() bool {
	return m.WriteEnabled
}

// SetWriteEnabled sets the WriteEnabled property
func (m *UsersTokensListQueryParameters) SetWriteEnabled(val bool) {
	m.WriteEnabled = val
}
