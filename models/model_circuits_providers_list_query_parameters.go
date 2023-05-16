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

// CircuitsProvidersListQueryParameters is an object.
type CircuitsProvidersListQueryParameters struct {
	// AsnId: ASN (ID)
	AsnId []int32 `json:"asn_id,omitempty" mapstructure:"asn_id,omitempty"`
	// AsnIdN: ASN (ID)
	AsnIdN []int32 `json:"asn_id__n,omitempty" mapstructure:"asn_id__n,omitempty"`
	// Contact: Contact
	Contact []int32 `json:"contact,omitempty" mapstructure:"contact,omitempty"`
	// ContactN: Contact
	ContactN []int32 `json:"contact__n,omitempty" mapstructure:"contact__n,omitempty"`
	// ContactGroup: Contact group
	ContactGroup []int32 `json:"contact_group,omitempty" mapstructure:"contact_group,omitempty"`
	// ContactGroupN: Contact group
	ContactGroupN []int32 `json:"contact_group__n,omitempty" mapstructure:"contact_group__n,omitempty"`
	// ContactRole: Contact Role
	ContactRole []int32 `json:"contact_role,omitempty" mapstructure:"contact_role,omitempty"`
	// ContactRoleN: Contact Role
	ContactRoleN []int32 `json:"contact_role__n,omitempty" mapstructure:"contact_role__n,omitempty"`
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
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Region: Region (slug)
	Region []int32 `json:"region,omitempty" mapstructure:"region,omitempty"`
	// RegionN: Region (slug)
	RegionN []int32 `json:"region__n,omitempty" mapstructure:"region__n,omitempty"`
	// RegionId: Region (ID)
	RegionId []int32 `json:"region_id,omitempty" mapstructure:"region_id,omitempty"`
	// RegionIdN: Region (ID)
	RegionIdN []int32 `json:"region_id__n,omitempty" mapstructure:"region_id__n,omitempty"`
	// Site: Site (slug)
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteN: Site (slug)
	SiteN []string `json:"site__n,omitempty" mapstructure:"site__n,omitempty"`
	// SiteGroup: Site group (slug)
	SiteGroup []int32 `json:"site_group,omitempty" mapstructure:"site_group,omitempty"`
	// SiteGroupN: Site group (slug)
	SiteGroupN []int32 `json:"site_group__n,omitempty" mapstructure:"site_group__n,omitempty"`
	// SiteGroupId: Site group (ID)
	SiteGroupId []int32 `json:"site_group_id,omitempty" mapstructure:"site_group_id,omitempty"`
	// SiteGroupIdN: Site group (ID)
	SiteGroupIdN []int32 `json:"site_group_id__n,omitempty" mapstructure:"site_group_id__n,omitempty"`
	// SiteId: Site
	SiteId []int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
	// SiteIdN: Site
	SiteIdN []int32 `json:"site_id__n,omitempty" mapstructure:"site_id__n,omitempty"`
	// Slug:
	Slug []string `json:"slug,omitempty" mapstructure:"slug,omitempty"`
	// SlugEmpty:
	SlugEmpty []string `json:"slug__empty,omitempty" mapstructure:"slug__empty,omitempty"`
	// SlugIc:
	SlugIc []string `json:"slug__ic,omitempty" mapstructure:"slug__ic,omitempty"`
	// SlugIe:
	SlugIe []string `json:"slug__ie,omitempty" mapstructure:"slug__ie,omitempty"`
	// SlugIew:
	SlugIew []string `json:"slug__iew,omitempty" mapstructure:"slug__iew,omitempty"`
	// SlugIsw:
	SlugIsw []string `json:"slug__isw,omitempty" mapstructure:"slug__isw,omitempty"`
	// SlugN:
	SlugN []string `json:"slug__n,omitempty" mapstructure:"slug__n,omitempty"`
	// SlugNic:
	SlugNic []string `json:"slug__nic,omitempty" mapstructure:"slug__nic,omitempty"`
	// SlugNie:
	SlugNie []string `json:"slug__nie,omitempty" mapstructure:"slug__nie,omitempty"`
	// SlugNiew:
	SlugNiew []string `json:"slug__niew,omitempty" mapstructure:"slug__niew,omitempty"`
	// SlugNisw:
	SlugNisw []string `json:"slug__nisw,omitempty" mapstructure:"slug__nisw,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitsProvidersListQueryParameters) Validate() error {
	return validation.Errors{
		"asnId": validation.Validate(
			m.AsnId,
		),
		"asnIdN": validation.Validate(
			m.AsnIdN,
		),
		"contact": validation.Validate(
			m.Contact,
		),
		"contactN": validation.Validate(
			m.ContactN,
		),
		"contactGroup": validation.Validate(
			m.ContactGroup,
		),
		"contactGroupN": validation.Validate(
			m.ContactGroupN,
		),
		"contactRole": validation.Validate(
			m.ContactRole,
		),
		"contactRoleN": validation.Validate(
			m.ContactRoleN,
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
		"region": validation.Validate(
			m.Region,
		),
		"regionN": validation.Validate(
			m.RegionN,
		),
		"regionId": validation.Validate(
			m.RegionId,
		),
		"regionIdN": validation.Validate(
			m.RegionIdN,
		),
		"site": validation.Validate(
			m.Site,
		),
		"siteN": validation.Validate(
			m.SiteN,
		),
		"siteGroup": validation.Validate(
			m.SiteGroup,
		),
		"siteGroupN": validation.Validate(
			m.SiteGroupN,
		),
		"siteGroupId": validation.Validate(
			m.SiteGroupId,
		),
		"siteGroupIdN": validation.Validate(
			m.SiteGroupIdN,
		),
		"siteId": validation.Validate(
			m.SiteId,
		),
		"siteIdN": validation.Validate(
			m.SiteIdN,
		),
		"slug": validation.Validate(
			m.Slug,
		),
		"slugEmpty": validation.Validate(
			m.SlugEmpty,
		),
		"slugIc": validation.Validate(
			m.SlugIc,
		),
		"slugIe": validation.Validate(
			m.SlugIe,
		),
		"slugIew": validation.Validate(
			m.SlugIew,
		),
		"slugIsw": validation.Validate(
			m.SlugIsw,
		),
		"slugN": validation.Validate(
			m.SlugN,
		),
		"slugNic": validation.Validate(
			m.SlugNic,
		),
		"slugNie": validation.Validate(
			m.SlugNie,
		),
		"slugNiew": validation.Validate(
			m.SlugNiew,
		),
		"slugNisw": validation.Validate(
			m.SlugNisw,
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

// GetAsnId returns the AsnId property
func (m CircuitsProvidersListQueryParameters) GetAsnId() []int32 {
	return m.AsnId
}

// SetAsnId sets the AsnId property
func (m *CircuitsProvidersListQueryParameters) SetAsnId(val []int32) {
	m.AsnId = val
}

// GetAsnIdN returns the AsnIdN property
func (m CircuitsProvidersListQueryParameters) GetAsnIdN() []int32 {
	return m.AsnIdN
}

// SetAsnIdN sets the AsnIdN property
func (m *CircuitsProvidersListQueryParameters) SetAsnIdN(val []int32) {
	m.AsnIdN = val
}

// GetContact returns the Contact property
func (m CircuitsProvidersListQueryParameters) GetContact() []int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *CircuitsProvidersListQueryParameters) SetContact(val []int32) {
	m.Contact = val
}

// GetContactN returns the ContactN property
func (m CircuitsProvidersListQueryParameters) GetContactN() []int32 {
	return m.ContactN
}

// SetContactN sets the ContactN property
func (m *CircuitsProvidersListQueryParameters) SetContactN(val []int32) {
	m.ContactN = val
}

// GetContactGroup returns the ContactGroup property
func (m CircuitsProvidersListQueryParameters) GetContactGroup() []int32 {
	return m.ContactGroup
}

// SetContactGroup sets the ContactGroup property
func (m *CircuitsProvidersListQueryParameters) SetContactGroup(val []int32) {
	m.ContactGroup = val
}

// GetContactGroupN returns the ContactGroupN property
func (m CircuitsProvidersListQueryParameters) GetContactGroupN() []int32 {
	return m.ContactGroupN
}

// SetContactGroupN sets the ContactGroupN property
func (m *CircuitsProvidersListQueryParameters) SetContactGroupN(val []int32) {
	m.ContactGroupN = val
}

// GetContactRole returns the ContactRole property
func (m CircuitsProvidersListQueryParameters) GetContactRole() []int32 {
	return m.ContactRole
}

// SetContactRole sets the ContactRole property
func (m *CircuitsProvidersListQueryParameters) SetContactRole(val []int32) {
	m.ContactRole = val
}

// GetContactRoleN returns the ContactRoleN property
func (m CircuitsProvidersListQueryParameters) GetContactRoleN() []int32 {
	return m.ContactRoleN
}

// SetContactRoleN sets the ContactRoleN property
func (m *CircuitsProvidersListQueryParameters) SetContactRoleN(val []int32) {
	m.ContactRoleN = val
}

// GetCreated returns the Created property
func (m CircuitsProvidersListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CircuitsProvidersListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m CircuitsProvidersListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *CircuitsProvidersListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m CircuitsProvidersListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *CircuitsProvidersListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m CircuitsProvidersListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *CircuitsProvidersListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m CircuitsProvidersListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *CircuitsProvidersListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m CircuitsProvidersListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *CircuitsProvidersListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m CircuitsProvidersListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *CircuitsProvidersListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetId returns the Id property
func (m CircuitsProvidersListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitsProvidersListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m CircuitsProvidersListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *CircuitsProvidersListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m CircuitsProvidersListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *CircuitsProvidersListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m CircuitsProvidersListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *CircuitsProvidersListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m CircuitsProvidersListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *CircuitsProvidersListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m CircuitsProvidersListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *CircuitsProvidersListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m CircuitsProvidersListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CircuitsProvidersListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m CircuitsProvidersListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *CircuitsProvidersListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m CircuitsProvidersListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *CircuitsProvidersListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m CircuitsProvidersListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *CircuitsProvidersListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m CircuitsProvidersListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *CircuitsProvidersListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m CircuitsProvidersListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *CircuitsProvidersListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m CircuitsProvidersListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *CircuitsProvidersListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m CircuitsProvidersListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *CircuitsProvidersListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m CircuitsProvidersListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *CircuitsProvidersListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m CircuitsProvidersListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *CircuitsProvidersListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m CircuitsProvidersListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *CircuitsProvidersListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m CircuitsProvidersListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *CircuitsProvidersListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m CircuitsProvidersListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *CircuitsProvidersListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m CircuitsProvidersListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *CircuitsProvidersListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m CircuitsProvidersListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *CircuitsProvidersListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m CircuitsProvidersListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *CircuitsProvidersListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m CircuitsProvidersListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *CircuitsProvidersListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m CircuitsProvidersListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *CircuitsProvidersListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m CircuitsProvidersListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *CircuitsProvidersListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m CircuitsProvidersListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *CircuitsProvidersListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m CircuitsProvidersListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *CircuitsProvidersListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m CircuitsProvidersListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *CircuitsProvidersListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m CircuitsProvidersListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *CircuitsProvidersListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m CircuitsProvidersListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *CircuitsProvidersListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m CircuitsProvidersListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *CircuitsProvidersListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSite returns the Site property
func (m CircuitsProvidersListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *CircuitsProvidersListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m CircuitsProvidersListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *CircuitsProvidersListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m CircuitsProvidersListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *CircuitsProvidersListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m CircuitsProvidersListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *CircuitsProvidersListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m CircuitsProvidersListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *CircuitsProvidersListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m CircuitsProvidersListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *CircuitsProvidersListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m CircuitsProvidersListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *CircuitsProvidersListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m CircuitsProvidersListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *CircuitsProvidersListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetSlug returns the Slug property
func (m CircuitsProvidersListQueryParameters) GetSlug() []string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *CircuitsProvidersListQueryParameters) SetSlug(val []string) {
	m.Slug = val
}

// GetSlugEmpty returns the SlugEmpty property
func (m CircuitsProvidersListQueryParameters) GetSlugEmpty() []string {
	return m.SlugEmpty
}

// SetSlugEmpty sets the SlugEmpty property
func (m *CircuitsProvidersListQueryParameters) SetSlugEmpty(val []string) {
	m.SlugEmpty = val
}

// GetSlugIc returns the SlugIc property
func (m CircuitsProvidersListQueryParameters) GetSlugIc() []string {
	return m.SlugIc
}

// SetSlugIc sets the SlugIc property
func (m *CircuitsProvidersListQueryParameters) SetSlugIc(val []string) {
	m.SlugIc = val
}

// GetSlugIe returns the SlugIe property
func (m CircuitsProvidersListQueryParameters) GetSlugIe() []string {
	return m.SlugIe
}

// SetSlugIe sets the SlugIe property
func (m *CircuitsProvidersListQueryParameters) SetSlugIe(val []string) {
	m.SlugIe = val
}

// GetSlugIew returns the SlugIew property
func (m CircuitsProvidersListQueryParameters) GetSlugIew() []string {
	return m.SlugIew
}

// SetSlugIew sets the SlugIew property
func (m *CircuitsProvidersListQueryParameters) SetSlugIew(val []string) {
	m.SlugIew = val
}

// GetSlugIsw returns the SlugIsw property
func (m CircuitsProvidersListQueryParameters) GetSlugIsw() []string {
	return m.SlugIsw
}

// SetSlugIsw sets the SlugIsw property
func (m *CircuitsProvidersListQueryParameters) SetSlugIsw(val []string) {
	m.SlugIsw = val
}

// GetSlugN returns the SlugN property
func (m CircuitsProvidersListQueryParameters) GetSlugN() []string {
	return m.SlugN
}

// SetSlugN sets the SlugN property
func (m *CircuitsProvidersListQueryParameters) SetSlugN(val []string) {
	m.SlugN = val
}

// GetSlugNic returns the SlugNic property
func (m CircuitsProvidersListQueryParameters) GetSlugNic() []string {
	return m.SlugNic
}

// SetSlugNic sets the SlugNic property
func (m *CircuitsProvidersListQueryParameters) SetSlugNic(val []string) {
	m.SlugNic = val
}

// GetSlugNie returns the SlugNie property
func (m CircuitsProvidersListQueryParameters) GetSlugNie() []string {
	return m.SlugNie
}

// SetSlugNie sets the SlugNie property
func (m *CircuitsProvidersListQueryParameters) SetSlugNie(val []string) {
	m.SlugNie = val
}

// GetSlugNiew returns the SlugNiew property
func (m CircuitsProvidersListQueryParameters) GetSlugNiew() []string {
	return m.SlugNiew
}

// SetSlugNiew sets the SlugNiew property
func (m *CircuitsProvidersListQueryParameters) SetSlugNiew(val []string) {
	m.SlugNiew = val
}

// GetSlugNisw returns the SlugNisw property
func (m CircuitsProvidersListQueryParameters) GetSlugNisw() []string {
	return m.SlugNisw
}

// SetSlugNisw sets the SlugNisw property
func (m *CircuitsProvidersListQueryParameters) SetSlugNisw(val []string) {
	m.SlugNisw = val
}

// GetTag returns the Tag property
func (m CircuitsProvidersListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *CircuitsProvidersListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m CircuitsProvidersListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *CircuitsProvidersListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m CircuitsProvidersListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *CircuitsProvidersListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
