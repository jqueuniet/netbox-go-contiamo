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

// IpamFhrpGroupsListQueryParameters is an object.
type IpamFhrpGroupsListQueryParameters struct {
	// AuthKey:
	AuthKey []string `json:"auth_key,omitempty" mapstructure:"auth_key,omitempty"`
	// AuthKeyEmpty:
	AuthKeyEmpty []string `json:"auth_key__empty,omitempty" mapstructure:"auth_key__empty,omitempty"`
	// AuthKeyIc:
	AuthKeyIc []string `json:"auth_key__ic,omitempty" mapstructure:"auth_key__ic,omitempty"`
	// AuthKeyIe:
	AuthKeyIe []string `json:"auth_key__ie,omitempty" mapstructure:"auth_key__ie,omitempty"`
	// AuthKeyIew:
	AuthKeyIew []string `json:"auth_key__iew,omitempty" mapstructure:"auth_key__iew,omitempty"`
	// AuthKeyIsw:
	AuthKeyIsw []string `json:"auth_key__isw,omitempty" mapstructure:"auth_key__isw,omitempty"`
	// AuthKeyN:
	AuthKeyN []string `json:"auth_key__n,omitempty" mapstructure:"auth_key__n,omitempty"`
	// AuthKeyNic:
	AuthKeyNic []string `json:"auth_key__nic,omitempty" mapstructure:"auth_key__nic,omitempty"`
	// AuthKeyNie:
	AuthKeyNie []string `json:"auth_key__nie,omitempty" mapstructure:"auth_key__nie,omitempty"`
	// AuthKeyNiew:
	AuthKeyNiew []string `json:"auth_key__niew,omitempty" mapstructure:"auth_key__niew,omitempty"`
	// AuthKeyNisw:
	AuthKeyNisw []string `json:"auth_key__nisw,omitempty" mapstructure:"auth_key__nisw,omitempty"`
	// AuthType: * `plaintext` - Plaintext
	// * `md5` - MD5
	AuthType []string `json:"auth_type,omitempty" mapstructure:"auth_type,omitempty"`
	// AuthTypeN: * `plaintext` - Plaintext
	// * `md5` - MD5
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
	// GroupId:
	GroupId []int32 `json:"group_id,omitempty" mapstructure:"group_id,omitempty"`
	// GroupIdGt:
	GroupIdGt []int32 `json:"group_id__gt,omitempty" mapstructure:"group_id__gt,omitempty"`
	// GroupIdGte:
	GroupIdGte []int32 `json:"group_id__gte,omitempty" mapstructure:"group_id__gte,omitempty"`
	// GroupIdLt:
	GroupIdLt []int32 `json:"group_id__lt,omitempty" mapstructure:"group_id__lt,omitempty"`
	// GroupIdLte:
	GroupIdLte []int32 `json:"group_id__lte,omitempty" mapstructure:"group_id__lte,omitempty"`
	// GroupIdN:
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
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Protocol: * `vrrp2` - VRRPv2
	// * `vrrp3` - VRRPv3
	// * `carp` - CARP
	// * `clusterxl` - ClusterXL
	// * `hsrp` - HSRP
	// * `glbp` - GLBP
	// * `other` - Other
	Protocol []string `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	// ProtocolN: * `vrrp2` - VRRPv2
	// * `vrrp3` - VRRPv3
	// * `carp` - CARP
	// * `clusterxl` - ClusterXL
	// * `hsrp` - HSRP
	// * `glbp` - GLBP
	// * `other` - Other
	ProtocolN []string `json:"protocol__n,omitempty" mapstructure:"protocol__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// RelatedIp:
	RelatedIp []string `json:"related_ip,omitempty" mapstructure:"related_ip,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamFhrpGroupsListQueryParameters) Validate() error {
	return validation.Errors{
		"authKey": validation.Validate(
			m.AuthKey,
		),
		"authKeyEmpty": validation.Validate(
			m.AuthKeyEmpty,
		),
		"authKeyIc": validation.Validate(
			m.AuthKeyIc,
		),
		"authKeyIe": validation.Validate(
			m.AuthKeyIe,
		),
		"authKeyIew": validation.Validate(
			m.AuthKeyIew,
		),
		"authKeyIsw": validation.Validate(
			m.AuthKeyIsw,
		),
		"authKeyN": validation.Validate(
			m.AuthKeyN,
		),
		"authKeyNic": validation.Validate(
			m.AuthKeyNic,
		),
		"authKeyNie": validation.Validate(
			m.AuthKeyNie,
		),
		"authKeyNiew": validation.Validate(
			m.AuthKeyNiew,
		),
		"authKeyNisw": validation.Validate(
			m.AuthKeyNisw,
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
		"groupId": validation.Validate(
			m.GroupId,
		),
		"groupIdGt": validation.Validate(
			m.GroupIdGt,
		),
		"groupIdGte": validation.Validate(
			m.GroupIdGte,
		),
		"groupIdLt": validation.Validate(
			m.GroupIdLt,
		),
		"groupIdLte": validation.Validate(
			m.GroupIdLte,
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
		"protocol": validation.Validate(
			m.Protocol,
		),
		"protocolN": validation.Validate(
			m.ProtocolN,
		),
		"relatedIp": validation.Validate(
			m.RelatedIp,
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

// GetAuthKey returns the AuthKey property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKey() []string {
	return m.AuthKey
}

// SetAuthKey sets the AuthKey property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKey(val []string) {
	m.AuthKey = val
}

// GetAuthKeyEmpty returns the AuthKeyEmpty property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyEmpty() []string {
	return m.AuthKeyEmpty
}

// SetAuthKeyEmpty sets the AuthKeyEmpty property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyEmpty(val []string) {
	m.AuthKeyEmpty = val
}

// GetAuthKeyIc returns the AuthKeyIc property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyIc() []string {
	return m.AuthKeyIc
}

// SetAuthKeyIc sets the AuthKeyIc property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyIc(val []string) {
	m.AuthKeyIc = val
}

// GetAuthKeyIe returns the AuthKeyIe property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyIe() []string {
	return m.AuthKeyIe
}

// SetAuthKeyIe sets the AuthKeyIe property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyIe(val []string) {
	m.AuthKeyIe = val
}

// GetAuthKeyIew returns the AuthKeyIew property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyIew() []string {
	return m.AuthKeyIew
}

// SetAuthKeyIew sets the AuthKeyIew property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyIew(val []string) {
	m.AuthKeyIew = val
}

// GetAuthKeyIsw returns the AuthKeyIsw property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyIsw() []string {
	return m.AuthKeyIsw
}

// SetAuthKeyIsw sets the AuthKeyIsw property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyIsw(val []string) {
	m.AuthKeyIsw = val
}

// GetAuthKeyN returns the AuthKeyN property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyN() []string {
	return m.AuthKeyN
}

// SetAuthKeyN sets the AuthKeyN property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyN(val []string) {
	m.AuthKeyN = val
}

// GetAuthKeyNic returns the AuthKeyNic property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyNic() []string {
	return m.AuthKeyNic
}

// SetAuthKeyNic sets the AuthKeyNic property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyNic(val []string) {
	m.AuthKeyNic = val
}

// GetAuthKeyNie returns the AuthKeyNie property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyNie() []string {
	return m.AuthKeyNie
}

// SetAuthKeyNie sets the AuthKeyNie property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyNie(val []string) {
	m.AuthKeyNie = val
}

// GetAuthKeyNiew returns the AuthKeyNiew property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyNiew() []string {
	return m.AuthKeyNiew
}

// SetAuthKeyNiew sets the AuthKeyNiew property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyNiew(val []string) {
	m.AuthKeyNiew = val
}

// GetAuthKeyNisw returns the AuthKeyNisw property
func (m IpamFhrpGroupsListQueryParameters) GetAuthKeyNisw() []string {
	return m.AuthKeyNisw
}

// SetAuthKeyNisw sets the AuthKeyNisw property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthKeyNisw(val []string) {
	m.AuthKeyNisw = val
}

// GetAuthType returns the AuthType property
func (m IpamFhrpGroupsListQueryParameters) GetAuthType() []string {
	return m.AuthType
}

// SetAuthType sets the AuthType property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthType(val []string) {
	m.AuthType = val
}

// GetAuthTypeN returns the AuthTypeN property
func (m IpamFhrpGroupsListQueryParameters) GetAuthTypeN() []string {
	return m.AuthTypeN
}

// SetAuthTypeN sets the AuthTypeN property
func (m *IpamFhrpGroupsListQueryParameters) SetAuthTypeN(val []string) {
	m.AuthTypeN = val
}

// GetCreated returns the Created property
func (m IpamFhrpGroupsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamFhrpGroupsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamFhrpGroupsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamFhrpGroupsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamFhrpGroupsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamFhrpGroupsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamFhrpGroupsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamFhrpGroupsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamFhrpGroupsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamFhrpGroupsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamFhrpGroupsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamFhrpGroupsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamFhrpGroupsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamFhrpGroupsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetGroupId returns the GroupId property
func (m IpamFhrpGroupsListQueryParameters) GetGroupId() []int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *IpamFhrpGroupsListQueryParameters) SetGroupId(val []int32) {
	m.GroupId = val
}

// GetGroupIdGt returns the GroupIdGt property
func (m IpamFhrpGroupsListQueryParameters) GetGroupIdGt() []int32 {
	return m.GroupIdGt
}

// SetGroupIdGt sets the GroupIdGt property
func (m *IpamFhrpGroupsListQueryParameters) SetGroupIdGt(val []int32) {
	m.GroupIdGt = val
}

// GetGroupIdGte returns the GroupIdGte property
func (m IpamFhrpGroupsListQueryParameters) GetGroupIdGte() []int32 {
	return m.GroupIdGte
}

// SetGroupIdGte sets the GroupIdGte property
func (m *IpamFhrpGroupsListQueryParameters) SetGroupIdGte(val []int32) {
	m.GroupIdGte = val
}

// GetGroupIdLt returns the GroupIdLt property
func (m IpamFhrpGroupsListQueryParameters) GetGroupIdLt() []int32 {
	return m.GroupIdLt
}

// SetGroupIdLt sets the GroupIdLt property
func (m *IpamFhrpGroupsListQueryParameters) SetGroupIdLt(val []int32) {
	m.GroupIdLt = val
}

// GetGroupIdLte returns the GroupIdLte property
func (m IpamFhrpGroupsListQueryParameters) GetGroupIdLte() []int32 {
	return m.GroupIdLte
}

// SetGroupIdLte sets the GroupIdLte property
func (m *IpamFhrpGroupsListQueryParameters) SetGroupIdLte(val []int32) {
	m.GroupIdLte = val
}

// GetGroupIdN returns the GroupIdN property
func (m IpamFhrpGroupsListQueryParameters) GetGroupIdN() []int32 {
	return m.GroupIdN
}

// SetGroupIdN sets the GroupIdN property
func (m *IpamFhrpGroupsListQueryParameters) SetGroupIdN(val []int32) {
	m.GroupIdN = val
}

// GetId returns the Id property
func (m IpamFhrpGroupsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamFhrpGroupsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamFhrpGroupsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamFhrpGroupsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamFhrpGroupsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamFhrpGroupsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamFhrpGroupsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamFhrpGroupsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamFhrpGroupsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamFhrpGroupsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamFhrpGroupsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamFhrpGroupsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamFhrpGroupsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamFhrpGroupsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamFhrpGroupsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamFhrpGroupsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamFhrpGroupsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamFhrpGroupsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamFhrpGroupsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamFhrpGroupsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamFhrpGroupsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamFhrpGroupsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamFhrpGroupsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamFhrpGroupsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamFhrpGroupsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamFhrpGroupsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m IpamFhrpGroupsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *IpamFhrpGroupsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m IpamFhrpGroupsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *IpamFhrpGroupsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m IpamFhrpGroupsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *IpamFhrpGroupsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m IpamFhrpGroupsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *IpamFhrpGroupsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m IpamFhrpGroupsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *IpamFhrpGroupsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m IpamFhrpGroupsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *IpamFhrpGroupsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m IpamFhrpGroupsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *IpamFhrpGroupsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m IpamFhrpGroupsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *IpamFhrpGroupsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m IpamFhrpGroupsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *IpamFhrpGroupsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m IpamFhrpGroupsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *IpamFhrpGroupsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m IpamFhrpGroupsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *IpamFhrpGroupsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m IpamFhrpGroupsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamFhrpGroupsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamFhrpGroupsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamFhrpGroupsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetProtocol returns the Protocol property
func (m IpamFhrpGroupsListQueryParameters) GetProtocol() []string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *IpamFhrpGroupsListQueryParameters) SetProtocol(val []string) {
	m.Protocol = val
}

// GetProtocolN returns the ProtocolN property
func (m IpamFhrpGroupsListQueryParameters) GetProtocolN() []string {
	return m.ProtocolN
}

// SetProtocolN sets the ProtocolN property
func (m *IpamFhrpGroupsListQueryParameters) SetProtocolN(val []string) {
	m.ProtocolN = val
}

// GetQ returns the Q property
func (m IpamFhrpGroupsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamFhrpGroupsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRelatedIp returns the RelatedIp property
func (m IpamFhrpGroupsListQueryParameters) GetRelatedIp() []string {
	return m.RelatedIp
}

// SetRelatedIp sets the RelatedIp property
func (m *IpamFhrpGroupsListQueryParameters) SetRelatedIp(val []string) {
	m.RelatedIp = val
}

// GetTag returns the Tag property
func (m IpamFhrpGroupsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamFhrpGroupsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamFhrpGroupsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamFhrpGroupsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamFhrpGroupsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamFhrpGroupsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
