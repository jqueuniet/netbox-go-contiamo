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

// WirelessWirelessLinksListQueryParameters is an object.
type WirelessWirelessLinksListQueryParameters struct {
	// AuthCipher: * `auto` - Auto
	// * `tkip` - TKIP
	// * `aes` - AES
	AuthCipher []string `json:"auth_cipher,omitempty" mapstructure:"auth_cipher,omitempty"`
	// AuthCipherN: * `auto` - Auto
	// * `tkip` - TKIP
	// * `aes` - AES
	AuthCipherN []string `json:"auth_cipher__n,omitempty" mapstructure:"auth_cipher__n,omitempty"`
	// AuthPsk:
	AuthPsk []string `json:"auth_psk,omitempty" mapstructure:"auth_psk,omitempty"`
	// AuthPskEmpty:
	AuthPskEmpty []string `json:"auth_psk__empty,omitempty" mapstructure:"auth_psk__empty,omitempty"`
	// AuthPskIc:
	AuthPskIc []string `json:"auth_psk__ic,omitempty" mapstructure:"auth_psk__ic,omitempty"`
	// AuthPskIe:
	AuthPskIe []string `json:"auth_psk__ie,omitempty" mapstructure:"auth_psk__ie,omitempty"`
	// AuthPskIew:
	AuthPskIew []string `json:"auth_psk__iew,omitempty" mapstructure:"auth_psk__iew,omitempty"`
	// AuthPskIsw:
	AuthPskIsw []string `json:"auth_psk__isw,omitempty" mapstructure:"auth_psk__isw,omitempty"`
	// AuthPskN:
	AuthPskN []string `json:"auth_psk__n,omitempty" mapstructure:"auth_psk__n,omitempty"`
	// AuthPskNic:
	AuthPskNic []string `json:"auth_psk__nic,omitempty" mapstructure:"auth_psk__nic,omitempty"`
	// AuthPskNie:
	AuthPskNie []string `json:"auth_psk__nie,omitempty" mapstructure:"auth_psk__nie,omitempty"`
	// AuthPskNiew:
	AuthPskNiew []string `json:"auth_psk__niew,omitempty" mapstructure:"auth_psk__niew,omitempty"`
	// AuthPskNisw:
	AuthPskNisw []string `json:"auth_psk__nisw,omitempty" mapstructure:"auth_psk__nisw,omitempty"`
	// AuthType: * `open` - Open
	// * `wep` - WEP
	// * `wpa-personal` - WPA Personal (PSK)
	// * `wpa-enterprise` - WPA Enterprise
	AuthType []string `json:"auth_type,omitempty" mapstructure:"auth_type,omitempty"`
	// AuthTypeN: * `open` - Open
	// * `wep` - WEP
	// * `wpa-personal` - WPA Personal (PSK)
	// * `wpa-enterprise` - WPA Enterprise
	AuthTypeN []string `json:"auth_type__n,omitempty" mapstructure:"auth_type__n,omitempty"`
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
	// InterfaceAId:
	InterfaceAId []int32 `json:"interface_a_id,omitempty" mapstructure:"interface_a_id,omitempty"`
	// InterfaceAIdGt:
	InterfaceAIdGt []int32 `json:"interface_a_id__gt,omitempty" mapstructure:"interface_a_id__gt,omitempty"`
	// InterfaceAIdGte:
	InterfaceAIdGte []int32 `json:"interface_a_id__gte,omitempty" mapstructure:"interface_a_id__gte,omitempty"`
	// InterfaceAIdLt:
	InterfaceAIdLt []int32 `json:"interface_a_id__lt,omitempty" mapstructure:"interface_a_id__lt,omitempty"`
	// InterfaceAIdLte:
	InterfaceAIdLte []int32 `json:"interface_a_id__lte,omitempty" mapstructure:"interface_a_id__lte,omitempty"`
	// InterfaceAIdN:
	InterfaceAIdN []int32 `json:"interface_a_id__n,omitempty" mapstructure:"interface_a_id__n,omitempty"`
	// InterfaceBId:
	InterfaceBId []int32 `json:"interface_b_id,omitempty" mapstructure:"interface_b_id,omitempty"`
	// InterfaceBIdGt:
	InterfaceBIdGt []int32 `json:"interface_b_id__gt,omitempty" mapstructure:"interface_b_id__gt,omitempty"`
	// InterfaceBIdGte:
	InterfaceBIdGte []int32 `json:"interface_b_id__gte,omitempty" mapstructure:"interface_b_id__gte,omitempty"`
	// InterfaceBIdLt:
	InterfaceBIdLt []int32 `json:"interface_b_id__lt,omitempty" mapstructure:"interface_b_id__lt,omitempty"`
	// InterfaceBIdLte:
	InterfaceBIdLte []int32 `json:"interface_b_id__lte,omitempty" mapstructure:"interface_b_id__lte,omitempty"`
	// InterfaceBIdN:
	InterfaceBIdN []int32 `json:"interface_b_id__n,omitempty" mapstructure:"interface_b_id__n,omitempty"`
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
	// Ssid:
	Ssid []string `json:"ssid,omitempty" mapstructure:"ssid,omitempty"`
	// SsidEmpty:
	SsidEmpty []string `json:"ssid__empty,omitempty" mapstructure:"ssid__empty,omitempty"`
	// SsidIc:
	SsidIc []string `json:"ssid__ic,omitempty" mapstructure:"ssid__ic,omitempty"`
	// SsidIe:
	SsidIe []string `json:"ssid__ie,omitempty" mapstructure:"ssid__ie,omitempty"`
	// SsidIew:
	SsidIew []string `json:"ssid__iew,omitempty" mapstructure:"ssid__iew,omitempty"`
	// SsidIsw:
	SsidIsw []string `json:"ssid__isw,omitempty" mapstructure:"ssid__isw,omitempty"`
	// SsidN:
	SsidN []string `json:"ssid__n,omitempty" mapstructure:"ssid__n,omitempty"`
	// SsidNic:
	SsidNic []string `json:"ssid__nic,omitempty" mapstructure:"ssid__nic,omitempty"`
	// SsidNie:
	SsidNie []string `json:"ssid__nie,omitempty" mapstructure:"ssid__nie,omitempty"`
	// SsidNiew:
	SsidNiew []string `json:"ssid__niew,omitempty" mapstructure:"ssid__niew,omitempty"`
	// SsidNisw:
	SsidNisw []string `json:"ssid__nisw,omitempty" mapstructure:"ssid__nisw,omitempty"`
	// Status: * `connected` - Connected
	// * `planned` - Planned
	// * `decommissioning` - Decommissioning
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `connected` - Connected
	// * `planned` - Planned
	// * `decommissioning` - Decommissioning
	StatusN []string `json:"status__n,omitempty" mapstructure:"status__n,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// Tenant: Tenant (slug)
	Tenant []string `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TenantN: Tenant (slug)
	TenantN []string `json:"tenant__n,omitempty" mapstructure:"tenant__n,omitempty"`
	// TenantGroup: Tenant Group (slug)
	TenantGroup []int32 `json:"tenant_group,omitempty" mapstructure:"tenant_group,omitempty"`
	// TenantGroupN: Tenant Group (slug)
	TenantGroupN []int32 `json:"tenant_group__n,omitempty" mapstructure:"tenant_group__n,omitempty"`
	// TenantGroupId: Tenant Group (ID)
	TenantGroupId []int32 `json:"tenant_group_id,omitempty" mapstructure:"tenant_group_id,omitempty"`
	// TenantGroupIdN: Tenant Group (ID)
	TenantGroupIdN []int32 `json:"tenant_group_id__n,omitempty" mapstructure:"tenant_group_id__n,omitempty"`
	// TenantId: Tenant (ID)
	TenantId []*int32 `json:"tenant_id,omitempty" mapstructure:"tenant_id,omitempty"`
	// TenantIdN: Tenant (ID)
	TenantIdN []*int32 `json:"tenant_id__n,omitempty" mapstructure:"tenant_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m WirelessWirelessLinksListQueryParameters) Validate() error {
	return validation.Errors{
		"authCipher": validation.Validate(
			m.AuthCipher,
		),
		"authCipherN": validation.Validate(
			m.AuthCipherN,
		),
		"authPsk": validation.Validate(
			m.AuthPsk,
		),
		"authPskEmpty": validation.Validate(
			m.AuthPskEmpty,
		),
		"authPskIc": validation.Validate(
			m.AuthPskIc,
		),
		"authPskIe": validation.Validate(
			m.AuthPskIe,
		),
		"authPskIew": validation.Validate(
			m.AuthPskIew,
		),
		"authPskIsw": validation.Validate(
			m.AuthPskIsw,
		),
		"authPskN": validation.Validate(
			m.AuthPskN,
		),
		"authPskNic": validation.Validate(
			m.AuthPskNic,
		),
		"authPskNie": validation.Validate(
			m.AuthPskNie,
		),
		"authPskNiew": validation.Validate(
			m.AuthPskNiew,
		),
		"authPskNisw": validation.Validate(
			m.AuthPskNisw,
		),
		"authType": validation.Validate(
			m.AuthType,
		),
		"authTypeN": validation.Validate(
			m.AuthTypeN,
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
		"interfaceAId": validation.Validate(
			m.InterfaceAId,
		),
		"interfaceAIdGt": validation.Validate(
			m.InterfaceAIdGt,
		),
		"interfaceAIdGte": validation.Validate(
			m.InterfaceAIdGte,
		),
		"interfaceAIdLt": validation.Validate(
			m.InterfaceAIdLt,
		),
		"interfaceAIdLte": validation.Validate(
			m.InterfaceAIdLte,
		),
		"interfaceAIdN": validation.Validate(
			m.InterfaceAIdN,
		),
		"interfaceBId": validation.Validate(
			m.InterfaceBId,
		),
		"interfaceBIdGt": validation.Validate(
			m.InterfaceBIdGt,
		),
		"interfaceBIdGte": validation.Validate(
			m.InterfaceBIdGte,
		),
		"interfaceBIdLt": validation.Validate(
			m.InterfaceBIdLt,
		),
		"interfaceBIdLte": validation.Validate(
			m.InterfaceBIdLte,
		),
		"interfaceBIdN": validation.Validate(
			m.InterfaceBIdN,
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
		"ssid": validation.Validate(
			m.Ssid,
		),
		"ssidEmpty": validation.Validate(
			m.SsidEmpty,
		),
		"ssidIc": validation.Validate(
			m.SsidIc,
		),
		"ssidIe": validation.Validate(
			m.SsidIe,
		),
		"ssidIew": validation.Validate(
			m.SsidIew,
		),
		"ssidIsw": validation.Validate(
			m.SsidIsw,
		),
		"ssidN": validation.Validate(
			m.SsidN,
		),
		"ssidNic": validation.Validate(
			m.SsidNic,
		),
		"ssidNie": validation.Validate(
			m.SsidNie,
		),
		"ssidNiew": validation.Validate(
			m.SsidNiew,
		),
		"ssidNisw": validation.Validate(
			m.SsidNisw,
		),
		"status": validation.Validate(
			m.Status,
		),
		"statusN": validation.Validate(
			m.StatusN,
		),
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"tenantN": validation.Validate(
			m.TenantN,
		),
		"tenantGroup": validation.Validate(
			m.TenantGroup,
		),
		"tenantGroupN": validation.Validate(
			m.TenantGroupN,
		),
		"tenantGroupId": validation.Validate(
			m.TenantGroupId,
		),
		"tenantGroupIdN": validation.Validate(
			m.TenantGroupIdN,
		),
		"tenantId": validation.Validate(
			m.TenantId,
		),
		"tenantIdN": validation.Validate(
			m.TenantIdN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetAuthCipher returns the AuthCipher property
func (m WirelessWirelessLinksListQueryParameters) GetAuthCipher() []string {
	return m.AuthCipher
}

// SetAuthCipher sets the AuthCipher property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthCipher(val []string) {
	m.AuthCipher = val
}

// GetAuthCipherN returns the AuthCipherN property
func (m WirelessWirelessLinksListQueryParameters) GetAuthCipherN() []string {
	return m.AuthCipherN
}

// SetAuthCipherN sets the AuthCipherN property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthCipherN(val []string) {
	m.AuthCipherN = val
}

// GetAuthPsk returns the AuthPsk property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPsk() []string {
	return m.AuthPsk
}

// SetAuthPsk sets the AuthPsk property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPsk(val []string) {
	m.AuthPsk = val
}

// GetAuthPskEmpty returns the AuthPskEmpty property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskEmpty() []string {
	return m.AuthPskEmpty
}

// SetAuthPskEmpty sets the AuthPskEmpty property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskEmpty(val []string) {
	m.AuthPskEmpty = val
}

// GetAuthPskIc returns the AuthPskIc property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskIc() []string {
	return m.AuthPskIc
}

// SetAuthPskIc sets the AuthPskIc property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskIc(val []string) {
	m.AuthPskIc = val
}

// GetAuthPskIe returns the AuthPskIe property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskIe() []string {
	return m.AuthPskIe
}

// SetAuthPskIe sets the AuthPskIe property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskIe(val []string) {
	m.AuthPskIe = val
}

// GetAuthPskIew returns the AuthPskIew property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskIew() []string {
	return m.AuthPskIew
}

// SetAuthPskIew sets the AuthPskIew property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskIew(val []string) {
	m.AuthPskIew = val
}

// GetAuthPskIsw returns the AuthPskIsw property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskIsw() []string {
	return m.AuthPskIsw
}

// SetAuthPskIsw sets the AuthPskIsw property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskIsw(val []string) {
	m.AuthPskIsw = val
}

// GetAuthPskN returns the AuthPskN property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskN() []string {
	return m.AuthPskN
}

// SetAuthPskN sets the AuthPskN property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskN(val []string) {
	m.AuthPskN = val
}

// GetAuthPskNic returns the AuthPskNic property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskNic() []string {
	return m.AuthPskNic
}

// SetAuthPskNic sets the AuthPskNic property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskNic(val []string) {
	m.AuthPskNic = val
}

// GetAuthPskNie returns the AuthPskNie property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskNie() []string {
	return m.AuthPskNie
}

// SetAuthPskNie sets the AuthPskNie property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskNie(val []string) {
	m.AuthPskNie = val
}

// GetAuthPskNiew returns the AuthPskNiew property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskNiew() []string {
	return m.AuthPskNiew
}

// SetAuthPskNiew sets the AuthPskNiew property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskNiew(val []string) {
	m.AuthPskNiew = val
}

// GetAuthPskNisw returns the AuthPskNisw property
func (m WirelessWirelessLinksListQueryParameters) GetAuthPskNisw() []string {
	return m.AuthPskNisw
}

// SetAuthPskNisw sets the AuthPskNisw property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthPskNisw(val []string) {
	m.AuthPskNisw = val
}

// GetAuthType returns the AuthType property
func (m WirelessWirelessLinksListQueryParameters) GetAuthType() []string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthType(val []string) {
	m.AuthType = val
}

// GetAuthTypeN returns the AuthTypeN property
func (m WirelessWirelessLinksListQueryParameters) GetAuthTypeN() []string {
	return m.AuthTypeN
}

// SetAuthTypeN sets the AuthTypeN property
func (m *WirelessWirelessLinksListQueryParameters) SetAuthTypeN(val []string) {
	m.AuthTypeN = val
}

// GetCreated returns the Created property
func (m WirelessWirelessLinksListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *WirelessWirelessLinksListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m WirelessWirelessLinksListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *WirelessWirelessLinksListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m WirelessWirelessLinksListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *WirelessWirelessLinksListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m WirelessWirelessLinksListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *WirelessWirelessLinksListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m WirelessWirelessLinksListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *WirelessWirelessLinksListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m WirelessWirelessLinksListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *WirelessWirelessLinksListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m WirelessWirelessLinksListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *WirelessWirelessLinksListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m WirelessWirelessLinksListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WirelessWirelessLinksListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m WirelessWirelessLinksListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *WirelessWirelessLinksListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m WirelessWirelessLinksListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *WirelessWirelessLinksListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m WirelessWirelessLinksListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *WirelessWirelessLinksListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m WirelessWirelessLinksListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *WirelessWirelessLinksListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m WirelessWirelessLinksListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *WirelessWirelessLinksListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m WirelessWirelessLinksListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *WirelessWirelessLinksListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m WirelessWirelessLinksListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *WirelessWirelessLinksListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterfaceAId returns the InterfaceAId property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceAId() []int32 {
	return m.InterfaceAId
}

// SetInterfaceAId sets the InterfaceAId property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceAId(val []int32) {
	m.InterfaceAId = val
}

// GetInterfaceAIdGt returns the InterfaceAIdGt property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceAIdGt() []int32 {
	return m.InterfaceAIdGt
}

// SetInterfaceAIdGt sets the InterfaceAIdGt property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceAIdGt(val []int32) {
	m.InterfaceAIdGt = val
}

// GetInterfaceAIdGte returns the InterfaceAIdGte property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceAIdGte() []int32 {
	return m.InterfaceAIdGte
}

// SetInterfaceAIdGte sets the InterfaceAIdGte property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceAIdGte(val []int32) {
	m.InterfaceAIdGte = val
}

// GetInterfaceAIdLt returns the InterfaceAIdLt property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceAIdLt() []int32 {
	return m.InterfaceAIdLt
}

// SetInterfaceAIdLt sets the InterfaceAIdLt property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceAIdLt(val []int32) {
	m.InterfaceAIdLt = val
}

// GetInterfaceAIdLte returns the InterfaceAIdLte property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceAIdLte() []int32 {
	return m.InterfaceAIdLte
}

// SetInterfaceAIdLte sets the InterfaceAIdLte property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceAIdLte(val []int32) {
	m.InterfaceAIdLte = val
}

// GetInterfaceAIdN returns the InterfaceAIdN property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceAIdN() []int32 {
	return m.InterfaceAIdN
}

// SetInterfaceAIdN sets the InterfaceAIdN property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceAIdN(val []int32) {
	m.InterfaceAIdN = val
}

// GetInterfaceBId returns the InterfaceBId property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceBId() []int32 {
	return m.InterfaceBId
}

// SetInterfaceBId sets the InterfaceBId property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceBId(val []int32) {
	m.InterfaceBId = val
}

// GetInterfaceBIdGt returns the InterfaceBIdGt property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceBIdGt() []int32 {
	return m.InterfaceBIdGt
}

// SetInterfaceBIdGt sets the InterfaceBIdGt property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceBIdGt(val []int32) {
	m.InterfaceBIdGt = val
}

// GetInterfaceBIdGte returns the InterfaceBIdGte property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceBIdGte() []int32 {
	return m.InterfaceBIdGte
}

// SetInterfaceBIdGte sets the InterfaceBIdGte property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceBIdGte(val []int32) {
	m.InterfaceBIdGte = val
}

// GetInterfaceBIdLt returns the InterfaceBIdLt property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceBIdLt() []int32 {
	return m.InterfaceBIdLt
}

// SetInterfaceBIdLt sets the InterfaceBIdLt property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceBIdLt(val []int32) {
	m.InterfaceBIdLt = val
}

// GetInterfaceBIdLte returns the InterfaceBIdLte property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceBIdLte() []int32 {
	return m.InterfaceBIdLte
}

// SetInterfaceBIdLte sets the InterfaceBIdLte property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceBIdLte(val []int32) {
	m.InterfaceBIdLte = val
}

// GetInterfaceBIdN returns the InterfaceBIdN property
func (m WirelessWirelessLinksListQueryParameters) GetInterfaceBIdN() []int32 {
	return m.InterfaceBIdN
}

// SetInterfaceBIdN sets the InterfaceBIdN property
func (m *WirelessWirelessLinksListQueryParameters) SetInterfaceBIdN(val []int32) {
	m.InterfaceBIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m WirelessWirelessLinksListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *WirelessWirelessLinksListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m WirelessWirelessLinksListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *WirelessWirelessLinksListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m WirelessWirelessLinksListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *WirelessWirelessLinksListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m WirelessWirelessLinksListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *WirelessWirelessLinksListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m WirelessWirelessLinksListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *WirelessWirelessLinksListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m WirelessWirelessLinksListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *WirelessWirelessLinksListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m WirelessWirelessLinksListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *WirelessWirelessLinksListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m WirelessWirelessLinksListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *WirelessWirelessLinksListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m WirelessWirelessLinksListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *WirelessWirelessLinksListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m WirelessWirelessLinksListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *WirelessWirelessLinksListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSsid returns the Ssid property
func (m WirelessWirelessLinksListQueryParameters) GetSsid() []string {
	return m.Ssid
}

// SetSsid sets the Ssid property
func (m *WirelessWirelessLinksListQueryParameters) SetSsid(val []string) {
	m.Ssid = val
}

// GetSsidEmpty returns the SsidEmpty property
func (m WirelessWirelessLinksListQueryParameters) GetSsidEmpty() []string {
	return m.SsidEmpty
}

// SetSsidEmpty sets the SsidEmpty property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidEmpty(val []string) {
	m.SsidEmpty = val
}

// GetSsidIc returns the SsidIc property
func (m WirelessWirelessLinksListQueryParameters) GetSsidIc() []string {
	return m.SsidIc
}

// SetSsidIc sets the SsidIc property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidIc(val []string) {
	m.SsidIc = val
}

// GetSsidIe returns the SsidIe property
func (m WirelessWirelessLinksListQueryParameters) GetSsidIe() []string {
	return m.SsidIe
}

// SetSsidIe sets the SsidIe property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidIe(val []string) {
	m.SsidIe = val
}

// GetSsidIew returns the SsidIew property
func (m WirelessWirelessLinksListQueryParameters) GetSsidIew() []string {
	return m.SsidIew
}

// SetSsidIew sets the SsidIew property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidIew(val []string) {
	m.SsidIew = val
}

// GetSsidIsw returns the SsidIsw property
func (m WirelessWirelessLinksListQueryParameters) GetSsidIsw() []string {
	return m.SsidIsw
}

// SetSsidIsw sets the SsidIsw property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidIsw(val []string) {
	m.SsidIsw = val
}

// GetSsidN returns the SsidN property
func (m WirelessWirelessLinksListQueryParameters) GetSsidN() []string {
	return m.SsidN
}

// SetSsidN sets the SsidN property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidN(val []string) {
	m.SsidN = val
}

// GetSsidNic returns the SsidNic property
func (m WirelessWirelessLinksListQueryParameters) GetSsidNic() []string {
	return m.SsidNic
}

// SetSsidNic sets the SsidNic property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidNic(val []string) {
	m.SsidNic = val
}

// GetSsidNie returns the SsidNie property
func (m WirelessWirelessLinksListQueryParameters) GetSsidNie() []string {
	return m.SsidNie
}

// SetSsidNie sets the SsidNie property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidNie(val []string) {
	m.SsidNie = val
}

// GetSsidNiew returns the SsidNiew property
func (m WirelessWirelessLinksListQueryParameters) GetSsidNiew() []string {
	return m.SsidNiew
}

// SetSsidNiew sets the SsidNiew property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidNiew(val []string) {
	m.SsidNiew = val
}

// GetSsidNisw returns the SsidNisw property
func (m WirelessWirelessLinksListQueryParameters) GetSsidNisw() []string {
	return m.SsidNisw
}

// SetSsidNisw sets the SsidNisw property
func (m *WirelessWirelessLinksListQueryParameters) SetSsidNisw(val []string) {
	m.SsidNisw = val
}

// GetStatus returns the Status property
func (m WirelessWirelessLinksListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WirelessWirelessLinksListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m WirelessWirelessLinksListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *WirelessWirelessLinksListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m WirelessWirelessLinksListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *WirelessWirelessLinksListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m WirelessWirelessLinksListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *WirelessWirelessLinksListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m WirelessWirelessLinksListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *WirelessWirelessLinksListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m WirelessWirelessLinksListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *WirelessWirelessLinksListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m WirelessWirelessLinksListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *WirelessWirelessLinksListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m WirelessWirelessLinksListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *WirelessWirelessLinksListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m WirelessWirelessLinksListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *WirelessWirelessLinksListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m WirelessWirelessLinksListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *WirelessWirelessLinksListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m WirelessWirelessLinksListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *WirelessWirelessLinksListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m WirelessWirelessLinksListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *WirelessWirelessLinksListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m WirelessWirelessLinksListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *WirelessWirelessLinksListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
