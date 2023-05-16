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

// DcimSitesListQueryParameters is an object.
type DcimSitesListQueryParameters struct {
	// Asn: AS (ID)
	Asn []int64 `json:"asn,omitempty" mapstructure:"asn,omitempty"`
	// AsnN: AS (ID)
	AsnN []int64 `json:"asn__n,omitempty" mapstructure:"asn__n,omitempty"`
	// AsnId: AS (ID)
	AsnId []int32 `json:"asn_id,omitempty" mapstructure:"asn_id,omitempty"`
	// AsnIdN: AS (ID)
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
	// Facility:
	Facility []string `json:"facility,omitempty" mapstructure:"facility,omitempty"`
	// FacilityEmpty:
	FacilityEmpty []string `json:"facility__empty,omitempty" mapstructure:"facility__empty,omitempty"`
	// FacilityIc:
	FacilityIc []string `json:"facility__ic,omitempty" mapstructure:"facility__ic,omitempty"`
	// FacilityIe:
	FacilityIe []string `json:"facility__ie,omitempty" mapstructure:"facility__ie,omitempty"`
	// FacilityIew:
	FacilityIew []string `json:"facility__iew,omitempty" mapstructure:"facility__iew,omitempty"`
	// FacilityIsw:
	FacilityIsw []string `json:"facility__isw,omitempty" mapstructure:"facility__isw,omitempty"`
	// FacilityN:
	FacilityN []string `json:"facility__n,omitempty" mapstructure:"facility__n,omitempty"`
	// FacilityNic:
	FacilityNic []string `json:"facility__nic,omitempty" mapstructure:"facility__nic,omitempty"`
	// FacilityNie:
	FacilityNie []string `json:"facility__nie,omitempty" mapstructure:"facility__nie,omitempty"`
	// FacilityNiew:
	FacilityNiew []string `json:"facility__niew,omitempty" mapstructure:"facility__niew,omitempty"`
	// FacilityNisw:
	FacilityNisw []string `json:"facility__nisw,omitempty" mapstructure:"facility__nisw,omitempty"`
	// Group: Group (slug)
	Group []int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// GroupN: Group (slug)
	GroupN []int32 `json:"group__n,omitempty" mapstructure:"group__n,omitempty"`
	// GroupId: Group (ID)
	GroupId []int32 `json:"group_id,omitempty" mapstructure:"group_id,omitempty"`
	// GroupIdN: Group (ID)
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
	// Latitude:
	Latitude []float64 `json:"latitude,omitempty" mapstructure:"latitude,omitempty"`
	// LatitudeGt:
	LatitudeGt []float64 `json:"latitude__gt,omitempty" mapstructure:"latitude__gt,omitempty"`
	// LatitudeGte:
	LatitudeGte []float64 `json:"latitude__gte,omitempty" mapstructure:"latitude__gte,omitempty"`
	// LatitudeLt:
	LatitudeLt []float64 `json:"latitude__lt,omitempty" mapstructure:"latitude__lt,omitempty"`
	// LatitudeLte:
	LatitudeLte []float64 `json:"latitude__lte,omitempty" mapstructure:"latitude__lte,omitempty"`
	// LatitudeN:
	LatitudeN []float64 `json:"latitude__n,omitempty" mapstructure:"latitude__n,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// Longitude:
	Longitude []float64 `json:"longitude,omitempty" mapstructure:"longitude,omitempty"`
	// LongitudeGt:
	LongitudeGt []float64 `json:"longitude__gt,omitempty" mapstructure:"longitude__gt,omitempty"`
	// LongitudeGte:
	LongitudeGte []float64 `json:"longitude__gte,omitempty" mapstructure:"longitude__gte,omitempty"`
	// LongitudeLt:
	LongitudeLt []float64 `json:"longitude__lt,omitempty" mapstructure:"longitude__lt,omitempty"`
	// LongitudeLte:
	LongitudeLte []float64 `json:"longitude__lte,omitempty" mapstructure:"longitude__lte,omitempty"`
	// LongitudeN:
	LongitudeN []float64 `json:"longitude__n,omitempty" mapstructure:"longitude__n,omitempty"`
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
	// Status: * `planned` - Planned
	// * `staging` - Staging
	// * `active` - Active
	// * `decommissioning` - Decommissioning
	// * `retired` - Retired
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `planned` - Planned
	// * `staging` - Staging
	// * `active` - Active
	// * `decommissioning` - Decommissioning
	// * `retired` - Retired
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
func (m DcimSitesListQueryParameters) Validate() error {
	return validation.Errors{
		"asn": validation.Validate(
			m.Asn,
		),
		"asnN": validation.Validate(
			m.AsnN,
		),
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
		"facility": validation.Validate(
			m.Facility,
		),
		"facilityEmpty": validation.Validate(
			m.FacilityEmpty,
		),
		"facilityIc": validation.Validate(
			m.FacilityIc,
		),
		"facilityIe": validation.Validate(
			m.FacilityIe,
		),
		"facilityIew": validation.Validate(
			m.FacilityIew,
		),
		"facilityIsw": validation.Validate(
			m.FacilityIsw,
		),
		"facilityN": validation.Validate(
			m.FacilityN,
		),
		"facilityNic": validation.Validate(
			m.FacilityNic,
		),
		"facilityNie": validation.Validate(
			m.FacilityNie,
		),
		"facilityNiew": validation.Validate(
			m.FacilityNiew,
		),
		"facilityNisw": validation.Validate(
			m.FacilityNisw,
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
		"latitude": validation.Validate(
			m.Latitude,
		),
		"latitudeGt": validation.Validate(
			m.LatitudeGt,
		),
		"latitudeGte": validation.Validate(
			m.LatitudeGte,
		),
		"latitudeLt": validation.Validate(
			m.LatitudeLt,
		),
		"latitudeLte": validation.Validate(
			m.LatitudeLte,
		),
		"latitudeN": validation.Validate(
			m.LatitudeN,
		),
		"longitude": validation.Validate(
			m.Longitude,
		),
		"longitudeGt": validation.Validate(
			m.LongitudeGt,
		),
		"longitudeGte": validation.Validate(
			m.LongitudeGte,
		),
		"longitudeLt": validation.Validate(
			m.LongitudeLt,
		),
		"longitudeLte": validation.Validate(
			m.LongitudeLte,
		),
		"longitudeN": validation.Validate(
			m.LongitudeN,
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

// GetAsn returns the Asn property
func (m DcimSitesListQueryParameters) GetAsn() []int64 {
	return m.Asn
}

// SetAsn sets the Asn property
func (m *DcimSitesListQueryParameters) SetAsn(val []int64) {
	m.Asn = val
}

// GetAsnN returns the AsnN property
func (m DcimSitesListQueryParameters) GetAsnN() []int64 {
	return m.AsnN
}

// SetAsnN sets the AsnN property
func (m *DcimSitesListQueryParameters) SetAsnN(val []int64) {
	m.AsnN = val
}

// GetAsnId returns the AsnId property
func (m DcimSitesListQueryParameters) GetAsnId() []int32 {
	return m.AsnId
}

// SetAsnId sets the AsnId property
func (m *DcimSitesListQueryParameters) SetAsnId(val []int32) {
	m.AsnId = val
}

// GetAsnIdN returns the AsnIdN property
func (m DcimSitesListQueryParameters) GetAsnIdN() []int32 {
	return m.AsnIdN
}

// SetAsnIdN sets the AsnIdN property
func (m *DcimSitesListQueryParameters) SetAsnIdN(val []int32) {
	m.AsnIdN = val
}

// GetContact returns the Contact property
func (m DcimSitesListQueryParameters) GetContact() []int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *DcimSitesListQueryParameters) SetContact(val []int32) {
	m.Contact = val
}

// GetContactN returns the ContactN property
func (m DcimSitesListQueryParameters) GetContactN() []int32 {
	return m.ContactN
}

// SetContactN sets the ContactN property
func (m *DcimSitesListQueryParameters) SetContactN(val []int32) {
	m.ContactN = val
}

// GetContactGroup returns the ContactGroup property
func (m DcimSitesListQueryParameters) GetContactGroup() []int32 {
	return m.ContactGroup
}

// SetContactGroup sets the ContactGroup property
func (m *DcimSitesListQueryParameters) SetContactGroup(val []int32) {
	m.ContactGroup = val
}

// GetContactGroupN returns the ContactGroupN property
func (m DcimSitesListQueryParameters) GetContactGroupN() []int32 {
	return m.ContactGroupN
}

// SetContactGroupN sets the ContactGroupN property
func (m *DcimSitesListQueryParameters) SetContactGroupN(val []int32) {
	m.ContactGroupN = val
}

// GetContactRole returns the ContactRole property
func (m DcimSitesListQueryParameters) GetContactRole() []int32 {
	return m.ContactRole
}

// SetContactRole sets the ContactRole property
func (m *DcimSitesListQueryParameters) SetContactRole(val []int32) {
	m.ContactRole = val
}

// GetContactRoleN returns the ContactRoleN property
func (m DcimSitesListQueryParameters) GetContactRoleN() []int32 {
	return m.ContactRoleN
}

// SetContactRoleN sets the ContactRoleN property
func (m *DcimSitesListQueryParameters) SetContactRoleN(val []int32) {
	m.ContactRoleN = val
}

// GetCreated returns the Created property
func (m DcimSitesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimSitesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimSitesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimSitesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimSitesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimSitesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimSitesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimSitesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimSitesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimSitesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimSitesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimSitesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimSitesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimSitesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m DcimSitesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DcimSitesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m DcimSitesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *DcimSitesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m DcimSitesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *DcimSitesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m DcimSitesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *DcimSitesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m DcimSitesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *DcimSitesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m DcimSitesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *DcimSitesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m DcimSitesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *DcimSitesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m DcimSitesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *DcimSitesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m DcimSitesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *DcimSitesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m DcimSitesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *DcimSitesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m DcimSitesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *DcimSitesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetFacility returns the Facility property
func (m DcimSitesListQueryParameters) GetFacility() []string {
	return m.Facility
}

// SetFacility sets the Facility property
func (m *DcimSitesListQueryParameters) SetFacility(val []string) {
	m.Facility = val
}

// GetFacilityEmpty returns the FacilityEmpty property
func (m DcimSitesListQueryParameters) GetFacilityEmpty() []string {
	return m.FacilityEmpty
}

// SetFacilityEmpty sets the FacilityEmpty property
func (m *DcimSitesListQueryParameters) SetFacilityEmpty(val []string) {
	m.FacilityEmpty = val
}

// GetFacilityIc returns the FacilityIc property
func (m DcimSitesListQueryParameters) GetFacilityIc() []string {
	return m.FacilityIc
}

// SetFacilityIc sets the FacilityIc property
func (m *DcimSitesListQueryParameters) SetFacilityIc(val []string) {
	m.FacilityIc = val
}

// GetFacilityIe returns the FacilityIe property
func (m DcimSitesListQueryParameters) GetFacilityIe() []string {
	return m.FacilityIe
}

// SetFacilityIe sets the FacilityIe property
func (m *DcimSitesListQueryParameters) SetFacilityIe(val []string) {
	m.FacilityIe = val
}

// GetFacilityIew returns the FacilityIew property
func (m DcimSitesListQueryParameters) GetFacilityIew() []string {
	return m.FacilityIew
}

// SetFacilityIew sets the FacilityIew property
func (m *DcimSitesListQueryParameters) SetFacilityIew(val []string) {
	m.FacilityIew = val
}

// GetFacilityIsw returns the FacilityIsw property
func (m DcimSitesListQueryParameters) GetFacilityIsw() []string {
	return m.FacilityIsw
}

// SetFacilityIsw sets the FacilityIsw property
func (m *DcimSitesListQueryParameters) SetFacilityIsw(val []string) {
	m.FacilityIsw = val
}

// GetFacilityN returns the FacilityN property
func (m DcimSitesListQueryParameters) GetFacilityN() []string {
	return m.FacilityN
}

// SetFacilityN sets the FacilityN property
func (m *DcimSitesListQueryParameters) SetFacilityN(val []string) {
	m.FacilityN = val
}

// GetFacilityNic returns the FacilityNic property
func (m DcimSitesListQueryParameters) GetFacilityNic() []string {
	return m.FacilityNic
}

// SetFacilityNic sets the FacilityNic property
func (m *DcimSitesListQueryParameters) SetFacilityNic(val []string) {
	m.FacilityNic = val
}

// GetFacilityNie returns the FacilityNie property
func (m DcimSitesListQueryParameters) GetFacilityNie() []string {
	return m.FacilityNie
}

// SetFacilityNie sets the FacilityNie property
func (m *DcimSitesListQueryParameters) SetFacilityNie(val []string) {
	m.FacilityNie = val
}

// GetFacilityNiew returns the FacilityNiew property
func (m DcimSitesListQueryParameters) GetFacilityNiew() []string {
	return m.FacilityNiew
}

// SetFacilityNiew sets the FacilityNiew property
func (m *DcimSitesListQueryParameters) SetFacilityNiew(val []string) {
	m.FacilityNiew = val
}

// GetFacilityNisw returns the FacilityNisw property
func (m DcimSitesListQueryParameters) GetFacilityNisw() []string {
	return m.FacilityNisw
}

// SetFacilityNisw sets the FacilityNisw property
func (m *DcimSitesListQueryParameters) SetFacilityNisw(val []string) {
	m.FacilityNisw = val
}

// GetGroup returns the Group property
func (m DcimSitesListQueryParameters) GetGroup() []int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *DcimSitesListQueryParameters) SetGroup(val []int32) {
	m.Group = val
}

// GetGroupN returns the GroupN property
func (m DcimSitesListQueryParameters) GetGroupN() []int32 {
	return m.GroupN
}

// SetGroupN sets the GroupN property
func (m *DcimSitesListQueryParameters) SetGroupN(val []int32) {
	m.GroupN = val
}

// GetGroupId returns the GroupId property
func (m DcimSitesListQueryParameters) GetGroupId() []int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *DcimSitesListQueryParameters) SetGroupId(val []int32) {
	m.GroupId = val
}

// GetGroupIdN returns the GroupIdN property
func (m DcimSitesListQueryParameters) GetGroupIdN() []int32 {
	return m.GroupIdN
}

// SetGroupIdN sets the GroupIdN property
func (m *DcimSitesListQueryParameters) SetGroupIdN(val []int32) {
	m.GroupIdN = val
}

// GetId returns the Id property
func (m DcimSitesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimSitesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimSitesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimSitesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimSitesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimSitesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimSitesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimSitesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimSitesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimSitesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimSitesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimSitesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimSitesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimSitesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimSitesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimSitesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimSitesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimSitesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimSitesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimSitesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimSitesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimSitesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimSitesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimSitesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLatitude returns the Latitude property
func (m DcimSitesListQueryParameters) GetLatitude() []float64 {
	return m.Latitude
}

// SetLatitude sets the Latitude property
func (m *DcimSitesListQueryParameters) SetLatitude(val []float64) {
	m.Latitude = val
}

// GetLatitudeGt returns the LatitudeGt property
func (m DcimSitesListQueryParameters) GetLatitudeGt() []float64 {
	return m.LatitudeGt
}

// SetLatitudeGt sets the LatitudeGt property
func (m *DcimSitesListQueryParameters) SetLatitudeGt(val []float64) {
	m.LatitudeGt = val
}

// GetLatitudeGte returns the LatitudeGte property
func (m DcimSitesListQueryParameters) GetLatitudeGte() []float64 {
	return m.LatitudeGte
}

// SetLatitudeGte sets the LatitudeGte property
func (m *DcimSitesListQueryParameters) SetLatitudeGte(val []float64) {
	m.LatitudeGte = val
}

// GetLatitudeLt returns the LatitudeLt property
func (m DcimSitesListQueryParameters) GetLatitudeLt() []float64 {
	return m.LatitudeLt
}

// SetLatitudeLt sets the LatitudeLt property
func (m *DcimSitesListQueryParameters) SetLatitudeLt(val []float64) {
	m.LatitudeLt = val
}

// GetLatitudeLte returns the LatitudeLte property
func (m DcimSitesListQueryParameters) GetLatitudeLte() []float64 {
	return m.LatitudeLte
}

// SetLatitudeLte sets the LatitudeLte property
func (m *DcimSitesListQueryParameters) SetLatitudeLte(val []float64) {
	m.LatitudeLte = val
}

// GetLatitudeN returns the LatitudeN property
func (m DcimSitesListQueryParameters) GetLatitudeN() []float64 {
	return m.LatitudeN
}

// SetLatitudeN sets the LatitudeN property
func (m *DcimSitesListQueryParameters) SetLatitudeN(val []float64) {
	m.LatitudeN = val
}

// GetLimit returns the Limit property
func (m DcimSitesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimSitesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLongitude returns the Longitude property
func (m DcimSitesListQueryParameters) GetLongitude() []float64 {
	return m.Longitude
}

// SetLongitude sets the Longitude property
func (m *DcimSitesListQueryParameters) SetLongitude(val []float64) {
	m.Longitude = val
}

// GetLongitudeGt returns the LongitudeGt property
func (m DcimSitesListQueryParameters) GetLongitudeGt() []float64 {
	return m.LongitudeGt
}

// SetLongitudeGt sets the LongitudeGt property
func (m *DcimSitesListQueryParameters) SetLongitudeGt(val []float64) {
	m.LongitudeGt = val
}

// GetLongitudeGte returns the LongitudeGte property
func (m DcimSitesListQueryParameters) GetLongitudeGte() []float64 {
	return m.LongitudeGte
}

// SetLongitudeGte sets the LongitudeGte property
func (m *DcimSitesListQueryParameters) SetLongitudeGte(val []float64) {
	m.LongitudeGte = val
}

// GetLongitudeLt returns the LongitudeLt property
func (m DcimSitesListQueryParameters) GetLongitudeLt() []float64 {
	return m.LongitudeLt
}

// SetLongitudeLt sets the LongitudeLt property
func (m *DcimSitesListQueryParameters) SetLongitudeLt(val []float64) {
	m.LongitudeLt = val
}

// GetLongitudeLte returns the LongitudeLte property
func (m DcimSitesListQueryParameters) GetLongitudeLte() []float64 {
	return m.LongitudeLte
}

// SetLongitudeLte sets the LongitudeLte property
func (m *DcimSitesListQueryParameters) SetLongitudeLte(val []float64) {
	m.LongitudeLte = val
}

// GetLongitudeN returns the LongitudeN property
func (m DcimSitesListQueryParameters) GetLongitudeN() []float64 {
	return m.LongitudeN
}

// SetLongitudeN sets the LongitudeN property
func (m *DcimSitesListQueryParameters) SetLongitudeN(val []float64) {
	m.LongitudeN = val
}

// GetName returns the Name property
func (m DcimSitesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimSitesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimSitesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimSitesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimSitesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimSitesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimSitesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimSitesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimSitesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimSitesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimSitesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimSitesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimSitesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimSitesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimSitesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimSitesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimSitesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimSitesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimSitesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimSitesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimSitesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimSitesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimSitesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimSitesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimSitesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimSitesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimSitesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimSitesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m DcimSitesListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimSitesListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimSitesListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimSitesListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimSitesListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimSitesListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimSitesListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimSitesListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetSlug returns the Slug property
func (m DcimSitesListQueryParameters) GetSlug() []string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *DcimSitesListQueryParameters) SetSlug(val []string) {
	m.Slug = val
}

// GetSlugEmpty returns the SlugEmpty property
func (m DcimSitesListQueryParameters) GetSlugEmpty() []string {
	return m.SlugEmpty
}

// SetSlugEmpty sets the SlugEmpty property
func (m *DcimSitesListQueryParameters) SetSlugEmpty(val []string) {
	m.SlugEmpty = val
}

// GetSlugIc returns the SlugIc property
func (m DcimSitesListQueryParameters) GetSlugIc() []string {
	return m.SlugIc
}

// SetSlugIc sets the SlugIc property
func (m *DcimSitesListQueryParameters) SetSlugIc(val []string) {
	m.SlugIc = val
}

// GetSlugIe returns the SlugIe property
func (m DcimSitesListQueryParameters) GetSlugIe() []string {
	return m.SlugIe
}

// SetSlugIe sets the SlugIe property
func (m *DcimSitesListQueryParameters) SetSlugIe(val []string) {
	m.SlugIe = val
}

// GetSlugIew returns the SlugIew property
func (m DcimSitesListQueryParameters) GetSlugIew() []string {
	return m.SlugIew
}

// SetSlugIew sets the SlugIew property
func (m *DcimSitesListQueryParameters) SetSlugIew(val []string) {
	m.SlugIew = val
}

// GetSlugIsw returns the SlugIsw property
func (m DcimSitesListQueryParameters) GetSlugIsw() []string {
	return m.SlugIsw
}

// SetSlugIsw sets the SlugIsw property
func (m *DcimSitesListQueryParameters) SetSlugIsw(val []string) {
	m.SlugIsw = val
}

// GetSlugN returns the SlugN property
func (m DcimSitesListQueryParameters) GetSlugN() []string {
	return m.SlugN
}

// SetSlugN sets the SlugN property
func (m *DcimSitesListQueryParameters) SetSlugN(val []string) {
	m.SlugN = val
}

// GetSlugNic returns the SlugNic property
func (m DcimSitesListQueryParameters) GetSlugNic() []string {
	return m.SlugNic
}

// SetSlugNic sets the SlugNic property
func (m *DcimSitesListQueryParameters) SetSlugNic(val []string) {
	m.SlugNic = val
}

// GetSlugNie returns the SlugNie property
func (m DcimSitesListQueryParameters) GetSlugNie() []string {
	return m.SlugNie
}

// SetSlugNie sets the SlugNie property
func (m *DcimSitesListQueryParameters) SetSlugNie(val []string) {
	m.SlugNie = val
}

// GetSlugNiew returns the SlugNiew property
func (m DcimSitesListQueryParameters) GetSlugNiew() []string {
	return m.SlugNiew
}

// SetSlugNiew sets the SlugNiew property
func (m *DcimSitesListQueryParameters) SetSlugNiew(val []string) {
	m.SlugNiew = val
}

// GetSlugNisw returns the SlugNisw property
func (m DcimSitesListQueryParameters) GetSlugNisw() []string {
	return m.SlugNisw
}

// SetSlugNisw sets the SlugNisw property
func (m *DcimSitesListQueryParameters) SetSlugNisw(val []string) {
	m.SlugNisw = val
}

// GetStatus returns the Status property
func (m DcimSitesListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *DcimSitesListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m DcimSitesListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *DcimSitesListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m DcimSitesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimSitesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimSitesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimSitesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m DcimSitesListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *DcimSitesListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m DcimSitesListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *DcimSitesListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m DcimSitesListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *DcimSitesListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m DcimSitesListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *DcimSitesListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m DcimSitesListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *DcimSitesListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m DcimSitesListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *DcimSitesListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m DcimSitesListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *DcimSitesListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m DcimSitesListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *DcimSitesListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimSitesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimSitesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
