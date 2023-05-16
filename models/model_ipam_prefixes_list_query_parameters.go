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

// IpamPrefixesListQueryParameters is an object.
type IpamPrefixesListQueryParameters struct {
	// Children:
	Children []int32 `json:"children,omitempty" mapstructure:"children,omitempty"`
	// ChildrenGt:
	ChildrenGt []int32 `json:"children__gt,omitempty" mapstructure:"children__gt,omitempty"`
	// ChildrenGte:
	ChildrenGte []int32 `json:"children__gte,omitempty" mapstructure:"children__gte,omitempty"`
	// ChildrenLt:
	ChildrenLt []int32 `json:"children__lt,omitempty" mapstructure:"children__lt,omitempty"`
	// ChildrenLte:
	ChildrenLte []int32 `json:"children__lte,omitempty" mapstructure:"children__lte,omitempty"`
	// ChildrenN:
	ChildrenN []int32 `json:"children__n,omitempty" mapstructure:"children__n,omitempty"`
	// Contains: Prefixes which contain this prefix or IP
	Contains string `json:"contains,omitempty" mapstructure:"contains,omitempty"`
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
	// Depth:
	Depth []int32 `json:"depth,omitempty" mapstructure:"depth,omitempty"`
	// DepthGt:
	DepthGt []int32 `json:"depth__gt,omitempty" mapstructure:"depth__gt,omitempty"`
	// DepthGte:
	DepthGte []int32 `json:"depth__gte,omitempty" mapstructure:"depth__gte,omitempty"`
	// DepthLt:
	DepthLt []int32 `json:"depth__lt,omitempty" mapstructure:"depth__lt,omitempty"`
	// DepthLte:
	DepthLte []int32 `json:"depth__lte,omitempty" mapstructure:"depth__lte,omitempty"`
	// DepthN:
	DepthN []int32 `json:"depth__n,omitempty" mapstructure:"depth__n,omitempty"`
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
	// Family:
	Family float32 `json:"family,omitempty" mapstructure:"family,omitempty"`
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
	// IsPool:
	IsPool bool `json:"is_pool,omitempty" mapstructure:"is_pool,omitempty"`
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
	// MarkUtilized:
	MarkUtilized bool `json:"mark_utilized,omitempty" mapstructure:"mark_utilized,omitempty"`
	// MaskLength:
	MaskLength []int32 `json:"mask_length,omitempty" mapstructure:"mask_length,omitempty"`
	// MaskLengthGte:
	MaskLengthGte float32 `json:"mask_length__gte,omitempty" mapstructure:"mask_length__gte,omitempty"`
	// MaskLengthLte:
	MaskLengthLte float32 `json:"mask_length__lte,omitempty" mapstructure:"mask_length__lte,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Prefix:
	Prefix []string `json:"prefix,omitempty" mapstructure:"prefix,omitempty"`
	// PresentInVrf: VRF (RD)
	PresentInVrf string `json:"present_in_vrf,omitempty" mapstructure:"present_in_vrf,omitempty"`
	// PresentInVrfId: VRF
	PresentInVrfId string `json:"present_in_vrf_id,omitempty" mapstructure:"present_in_vrf_id,omitempty"`
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
	// Role: Role (slug)
	Role []string `json:"role,omitempty" mapstructure:"role,omitempty"`
	// RoleN: Role (slug)
	RoleN []string `json:"role__n,omitempty" mapstructure:"role__n,omitempty"`
	// RoleId: Role (ID)
	RoleId []*int32 `json:"role_id,omitempty" mapstructure:"role_id,omitempty"`
	// RoleIdN: Role (ID)
	RoleIdN []*int32 `json:"role_id__n,omitempty" mapstructure:"role_id__n,omitempty"`
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
	// SiteId: Site (ID)
	SiteId []*int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
	// SiteIdN: Site (ID)
	SiteIdN []*int32 `json:"site_id__n,omitempty" mapstructure:"site_id__n,omitempty"`
	// Status: Operational status of this prefix
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: Operational status of this prefix
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
	// VlanId: VLAN (ID)
	VlanId []*int32 `json:"vlan_id,omitempty" mapstructure:"vlan_id,omitempty"`
	// VlanIdN: VLAN (ID)
	VlanIdN []*int32 `json:"vlan_id__n,omitempty" mapstructure:"vlan_id__n,omitempty"`
	// VlanVid: VLAN number (1-4094)
	VlanVid int32 `json:"vlan_vid,omitempty" mapstructure:"vlan_vid,omitempty"`
	// VlanVidGt: VLAN number (1-4094)
	VlanVidGt int32 `json:"vlan_vid__gt,omitempty" mapstructure:"vlan_vid__gt,omitempty"`
	// VlanVidGte: VLAN number (1-4094)
	VlanVidGte int32 `json:"vlan_vid__gte,omitempty" mapstructure:"vlan_vid__gte,omitempty"`
	// VlanVidLt: VLAN number (1-4094)
	VlanVidLt int32 `json:"vlan_vid__lt,omitempty" mapstructure:"vlan_vid__lt,omitempty"`
	// VlanVidLte: VLAN number (1-4094)
	VlanVidLte int32 `json:"vlan_vid__lte,omitempty" mapstructure:"vlan_vid__lte,omitempty"`
	// VlanVidN: VLAN number (1-4094)
	VlanVidN int32 `json:"vlan_vid__n,omitempty" mapstructure:"vlan_vid__n,omitempty"`
	// Vrf: VRF (RD)
	Vrf []*string `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
	// VrfN: VRF (RD)
	VrfN []*string `json:"vrf__n,omitempty" mapstructure:"vrf__n,omitempty"`
	// VrfId: VRF
	VrfId []*int32 `json:"vrf_id,omitempty" mapstructure:"vrf_id,omitempty"`
	// VrfIdN: VRF
	VrfIdN []*int32 `json:"vrf_id__n,omitempty" mapstructure:"vrf_id__n,omitempty"`
	// Within: Within prefix
	Within string `json:"within,omitempty" mapstructure:"within,omitempty"`
	// WithinInclude: Within and including prefix
	WithinInclude string `json:"within_include,omitempty" mapstructure:"within_include,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamPrefixesListQueryParameters) Validate() error {
	return validation.Errors{
		"children": validation.Validate(
			m.Children,
		),
		"childrenGt": validation.Validate(
			m.ChildrenGt,
		),
		"childrenGte": validation.Validate(
			m.ChildrenGte,
		),
		"childrenLt": validation.Validate(
			m.ChildrenLt,
		),
		"childrenLte": validation.Validate(
			m.ChildrenLte,
		),
		"childrenN": validation.Validate(
			m.ChildrenN,
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
		"depth": validation.Validate(
			m.Depth,
		),
		"depthGt": validation.Validate(
			m.DepthGt,
		),
		"depthGte": validation.Validate(
			m.DepthGte,
		),
		"depthLt": validation.Validate(
			m.DepthLt,
		),
		"depthLte": validation.Validate(
			m.DepthLte,
		),
		"depthN": validation.Validate(
			m.DepthN,
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
		"maskLength": validation.Validate(
			m.MaskLength,
		),
		"prefix": validation.Validate(
			m.Prefix,
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
		"role": validation.Validate(
			m.Role,
		),
		"roleN": validation.Validate(
			m.RoleN,
		),
		"roleId": validation.Validate(
			m.RoleId,
		),
		"roleIdN": validation.Validate(
			m.RoleIdN,
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
		"vlanId": validation.Validate(
			m.VlanId,
		),
		"vlanIdN": validation.Validate(
			m.VlanIdN,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
		"vrfN": validation.Validate(
			m.VrfN,
		),
		"vrfId": validation.Validate(
			m.VrfId,
		),
		"vrfIdN": validation.Validate(
			m.VrfIdN,
		),
	}.Filter()
}

// GetChildren returns the Children property
func (m IpamPrefixesListQueryParameters) GetChildren() []int32 {
	return m.Children
}

// SetChildren sets the Children property
func (m *IpamPrefixesListQueryParameters) SetChildren(val []int32) {
	m.Children = val
}

// GetChildrenGt returns the ChildrenGt property
func (m IpamPrefixesListQueryParameters) GetChildrenGt() []int32 {
	return m.ChildrenGt
}

// SetChildrenGt sets the ChildrenGt property
func (m *IpamPrefixesListQueryParameters) SetChildrenGt(val []int32) {
	m.ChildrenGt = val
}

// GetChildrenGte returns the ChildrenGte property
func (m IpamPrefixesListQueryParameters) GetChildrenGte() []int32 {
	return m.ChildrenGte
}

// SetChildrenGte sets the ChildrenGte property
func (m *IpamPrefixesListQueryParameters) SetChildrenGte(val []int32) {
	m.ChildrenGte = val
}

// GetChildrenLt returns the ChildrenLt property
func (m IpamPrefixesListQueryParameters) GetChildrenLt() []int32 {
	return m.ChildrenLt
}

// SetChildrenLt sets the ChildrenLt property
func (m *IpamPrefixesListQueryParameters) SetChildrenLt(val []int32) {
	m.ChildrenLt = val
}

// GetChildrenLte returns the ChildrenLte property
func (m IpamPrefixesListQueryParameters) GetChildrenLte() []int32 {
	return m.ChildrenLte
}

// SetChildrenLte sets the ChildrenLte property
func (m *IpamPrefixesListQueryParameters) SetChildrenLte(val []int32) {
	m.ChildrenLte = val
}

// GetChildrenN returns the ChildrenN property
func (m IpamPrefixesListQueryParameters) GetChildrenN() []int32 {
	return m.ChildrenN
}

// SetChildrenN sets the ChildrenN property
func (m *IpamPrefixesListQueryParameters) SetChildrenN(val []int32) {
	m.ChildrenN = val
}

// GetContains returns the Contains property
func (m IpamPrefixesListQueryParameters) GetContains() string {
	return m.Contains
}

// SetContains sets the Contains property
func (m *IpamPrefixesListQueryParameters) SetContains(val string) {
	m.Contains = val
}

// GetCreated returns the Created property
func (m IpamPrefixesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamPrefixesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamPrefixesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamPrefixesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamPrefixesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamPrefixesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamPrefixesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamPrefixesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamPrefixesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamPrefixesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamPrefixesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamPrefixesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamPrefixesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamPrefixesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDepth returns the Depth property
func (m IpamPrefixesListQueryParameters) GetDepth() []int32 {
	return m.Depth
}

// SetDepth sets the Depth property
func (m *IpamPrefixesListQueryParameters) SetDepth(val []int32) {
	m.Depth = val
}

// GetDepthGt returns the DepthGt property
func (m IpamPrefixesListQueryParameters) GetDepthGt() []int32 {
	return m.DepthGt
}

// SetDepthGt sets the DepthGt property
func (m *IpamPrefixesListQueryParameters) SetDepthGt(val []int32) {
	m.DepthGt = val
}

// GetDepthGte returns the DepthGte property
func (m IpamPrefixesListQueryParameters) GetDepthGte() []int32 {
	return m.DepthGte
}

// SetDepthGte sets the DepthGte property
func (m *IpamPrefixesListQueryParameters) SetDepthGte(val []int32) {
	m.DepthGte = val
}

// GetDepthLt returns the DepthLt property
func (m IpamPrefixesListQueryParameters) GetDepthLt() []int32 {
	return m.DepthLt
}

// SetDepthLt sets the DepthLt property
func (m *IpamPrefixesListQueryParameters) SetDepthLt(val []int32) {
	m.DepthLt = val
}

// GetDepthLte returns the DepthLte property
func (m IpamPrefixesListQueryParameters) GetDepthLte() []int32 {
	return m.DepthLte
}

// SetDepthLte sets the DepthLte property
func (m *IpamPrefixesListQueryParameters) SetDepthLte(val []int32) {
	m.DepthLte = val
}

// GetDepthN returns the DepthN property
func (m IpamPrefixesListQueryParameters) GetDepthN() []int32 {
	return m.DepthN
}

// SetDepthN sets the DepthN property
func (m *IpamPrefixesListQueryParameters) SetDepthN(val []int32) {
	m.DepthN = val
}

// GetDescription returns the Description property
func (m IpamPrefixesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IpamPrefixesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m IpamPrefixesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *IpamPrefixesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m IpamPrefixesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *IpamPrefixesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m IpamPrefixesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *IpamPrefixesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m IpamPrefixesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *IpamPrefixesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m IpamPrefixesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *IpamPrefixesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m IpamPrefixesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *IpamPrefixesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m IpamPrefixesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *IpamPrefixesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m IpamPrefixesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *IpamPrefixesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m IpamPrefixesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *IpamPrefixesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m IpamPrefixesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *IpamPrefixesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetFamily returns the Family property
func (m IpamPrefixesListQueryParameters) GetFamily() float32 {
	return m.Family
}

// SetFamily sets the Family property
func (m *IpamPrefixesListQueryParameters) SetFamily(val float32) {
	m.Family = val
}

// GetId returns the Id property
func (m IpamPrefixesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamPrefixesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamPrefixesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamPrefixesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamPrefixesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamPrefixesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamPrefixesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamPrefixesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamPrefixesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamPrefixesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamPrefixesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamPrefixesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetIsPool returns the IsPool property
func (m IpamPrefixesListQueryParameters) GetIsPool() bool {
	return m.IsPool
}

// SetIsPool sets the IsPool property
func (m *IpamPrefixesListQueryParameters) SetIsPool(val bool) {
	m.IsPool = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamPrefixesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamPrefixesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamPrefixesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamPrefixesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamPrefixesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamPrefixesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamPrefixesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamPrefixesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamPrefixesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamPrefixesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamPrefixesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamPrefixesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamPrefixesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamPrefixesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetMarkUtilized returns the MarkUtilized property
func (m IpamPrefixesListQueryParameters) GetMarkUtilized() bool {
	return m.MarkUtilized
}

// SetMarkUtilized sets the MarkUtilized property
func (m *IpamPrefixesListQueryParameters) SetMarkUtilized(val bool) {
	m.MarkUtilized = val
}

// GetMaskLength returns the MaskLength property
func (m IpamPrefixesListQueryParameters) GetMaskLength() []int32 {
	return m.MaskLength
}

// SetMaskLength sets the MaskLength property
func (m *IpamPrefixesListQueryParameters) SetMaskLength(val []int32) {
	m.MaskLength = val
}

// GetMaskLengthGte returns the MaskLengthGte property
func (m IpamPrefixesListQueryParameters) GetMaskLengthGte() float32 {
	return m.MaskLengthGte
}

// SetMaskLengthGte sets the MaskLengthGte property
func (m *IpamPrefixesListQueryParameters) SetMaskLengthGte(val float32) {
	m.MaskLengthGte = val
}

// GetMaskLengthLte returns the MaskLengthLte property
func (m IpamPrefixesListQueryParameters) GetMaskLengthLte() float32 {
	return m.MaskLengthLte
}

// SetMaskLengthLte sets the MaskLengthLte property
func (m *IpamPrefixesListQueryParameters) SetMaskLengthLte(val float32) {
	m.MaskLengthLte = val
}

// GetOffset returns the Offset property
func (m IpamPrefixesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamPrefixesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamPrefixesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamPrefixesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPrefix returns the Prefix property
func (m IpamPrefixesListQueryParameters) GetPrefix() []string {
	return m.Prefix
}

// SetPrefix sets the Prefix property
func (m *IpamPrefixesListQueryParameters) SetPrefix(val []string) {
	m.Prefix = val
}

// GetPresentInVrf returns the PresentInVrf property
func (m IpamPrefixesListQueryParameters) GetPresentInVrf() string {
	return m.PresentInVrf
}

// SetPresentInVrf sets the PresentInVrf property
func (m *IpamPrefixesListQueryParameters) SetPresentInVrf(val string) {
	m.PresentInVrf = val
}

// GetPresentInVrfId returns the PresentInVrfId property
func (m IpamPrefixesListQueryParameters) GetPresentInVrfId() string {
	return m.PresentInVrfId
}

// SetPresentInVrfId sets the PresentInVrfId property
func (m *IpamPrefixesListQueryParameters) SetPresentInVrfId(val string) {
	m.PresentInVrfId = val
}

// GetQ returns the Q property
func (m IpamPrefixesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamPrefixesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m IpamPrefixesListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *IpamPrefixesListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m IpamPrefixesListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *IpamPrefixesListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m IpamPrefixesListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *IpamPrefixesListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m IpamPrefixesListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *IpamPrefixesListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetRole returns the Role property
func (m IpamPrefixesListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *IpamPrefixesListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m IpamPrefixesListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *IpamPrefixesListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m IpamPrefixesListQueryParameters) GetRoleId() []*int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *IpamPrefixesListQueryParameters) SetRoleId(val []*int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m IpamPrefixesListQueryParameters) GetRoleIdN() []*int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *IpamPrefixesListQueryParameters) SetRoleIdN(val []*int32) {
	m.RoleIdN = val
}

// GetSite returns the Site property
func (m IpamPrefixesListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *IpamPrefixesListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m IpamPrefixesListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *IpamPrefixesListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m IpamPrefixesListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *IpamPrefixesListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m IpamPrefixesListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *IpamPrefixesListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m IpamPrefixesListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *IpamPrefixesListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m IpamPrefixesListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *IpamPrefixesListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m IpamPrefixesListQueryParameters) GetSiteId() []*int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *IpamPrefixesListQueryParameters) SetSiteId(val []*int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m IpamPrefixesListQueryParameters) GetSiteIdN() []*int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *IpamPrefixesListQueryParameters) SetSiteIdN(val []*int32) {
	m.SiteIdN = val
}

// GetStatus returns the Status property
func (m IpamPrefixesListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *IpamPrefixesListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m IpamPrefixesListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *IpamPrefixesListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m IpamPrefixesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamPrefixesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamPrefixesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamPrefixesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m IpamPrefixesListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IpamPrefixesListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m IpamPrefixesListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *IpamPrefixesListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m IpamPrefixesListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *IpamPrefixesListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m IpamPrefixesListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *IpamPrefixesListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m IpamPrefixesListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *IpamPrefixesListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m IpamPrefixesListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *IpamPrefixesListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m IpamPrefixesListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *IpamPrefixesListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m IpamPrefixesListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *IpamPrefixesListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamPrefixesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamPrefixesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVlanId returns the VlanId property
func (m IpamPrefixesListQueryParameters) GetVlanId() []*int32 {
	return m.VlanId
}

// SetVlanId sets the VlanId property
func (m *IpamPrefixesListQueryParameters) SetVlanId(val []*int32) {
	m.VlanId = val
}

// GetVlanIdN returns the VlanIdN property
func (m IpamPrefixesListQueryParameters) GetVlanIdN() []*int32 {
	return m.VlanIdN
}

// SetVlanIdN sets the VlanIdN property
func (m *IpamPrefixesListQueryParameters) SetVlanIdN(val []*int32) {
	m.VlanIdN = val
}

// GetVlanVid returns the VlanVid property
func (m IpamPrefixesListQueryParameters) GetVlanVid() int32 {
	return m.VlanVid
}

// SetVlanVid sets the VlanVid property
func (m *IpamPrefixesListQueryParameters) SetVlanVid(val int32) {
	m.VlanVid = val
}

// GetVlanVidGt returns the VlanVidGt property
func (m IpamPrefixesListQueryParameters) GetVlanVidGt() int32 {
	return m.VlanVidGt
}

// SetVlanVidGt sets the VlanVidGt property
func (m *IpamPrefixesListQueryParameters) SetVlanVidGt(val int32) {
	m.VlanVidGt = val
}

// GetVlanVidGte returns the VlanVidGte property
func (m IpamPrefixesListQueryParameters) GetVlanVidGte() int32 {
	return m.VlanVidGte
}

// SetVlanVidGte sets the VlanVidGte property
func (m *IpamPrefixesListQueryParameters) SetVlanVidGte(val int32) {
	m.VlanVidGte = val
}

// GetVlanVidLt returns the VlanVidLt property
func (m IpamPrefixesListQueryParameters) GetVlanVidLt() int32 {
	return m.VlanVidLt
}

// SetVlanVidLt sets the VlanVidLt property
func (m *IpamPrefixesListQueryParameters) SetVlanVidLt(val int32) {
	m.VlanVidLt = val
}

// GetVlanVidLte returns the VlanVidLte property
func (m IpamPrefixesListQueryParameters) GetVlanVidLte() int32 {
	return m.VlanVidLte
}

// SetVlanVidLte sets the VlanVidLte property
func (m *IpamPrefixesListQueryParameters) SetVlanVidLte(val int32) {
	m.VlanVidLte = val
}

// GetVlanVidN returns the VlanVidN property
func (m IpamPrefixesListQueryParameters) GetVlanVidN() int32 {
	return m.VlanVidN
}

// SetVlanVidN sets the VlanVidN property
func (m *IpamPrefixesListQueryParameters) SetVlanVidN(val int32) {
	m.VlanVidN = val
}

// GetVrf returns the Vrf property
func (m IpamPrefixesListQueryParameters) GetVrf() []*string {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *IpamPrefixesListQueryParameters) SetVrf(val []*string) {
	m.Vrf = val
}

// GetVrfN returns the VrfN property
func (m IpamPrefixesListQueryParameters) GetVrfN() []*string {
	return m.VrfN
}

// SetVrfN sets the VrfN property
func (m *IpamPrefixesListQueryParameters) SetVrfN(val []*string) {
	m.VrfN = val
}

// GetVrfId returns the VrfId property
func (m IpamPrefixesListQueryParameters) GetVrfId() []*int32 {
	return m.VrfId
}

// SetVrfId sets the VrfId property
func (m *IpamPrefixesListQueryParameters) SetVrfId(val []*int32) {
	m.VrfId = val
}

// GetVrfIdN returns the VrfIdN property
func (m IpamPrefixesListQueryParameters) GetVrfIdN() []*int32 {
	return m.VrfIdN
}

// SetVrfIdN sets the VrfIdN property
func (m *IpamPrefixesListQueryParameters) SetVrfIdN(val []*int32) {
	m.VrfIdN = val
}

// GetWithin returns the Within property
func (m IpamPrefixesListQueryParameters) GetWithin() string {
	return m.Within
}

// SetWithin sets the Within property
func (m *IpamPrefixesListQueryParameters) SetWithin(val string) {
	m.Within = val
}

// GetWithinInclude returns the WithinInclude property
func (m IpamPrefixesListQueryParameters) GetWithinInclude() string {
	return m.WithinInclude
}

// SetWithinInclude sets the WithinInclude property
func (m *IpamPrefixesListQueryParameters) SetWithinInclude(val string) {
	m.WithinInclude = val
}
