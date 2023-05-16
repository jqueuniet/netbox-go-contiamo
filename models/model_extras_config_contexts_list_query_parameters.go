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

// ExtrasConfigContextsListQueryParameters is an object.
type ExtrasConfigContextsListQueryParameters struct {
	// ClusterGroup: Cluster group (slug)
	ClusterGroup []string `json:"cluster_group,omitempty" mapstructure:"cluster_group,omitempty"`
	// ClusterGroupN: Cluster group (slug)
	ClusterGroupN []string `json:"cluster_group__n,omitempty" mapstructure:"cluster_group__n,omitempty"`
	// ClusterGroupId: Cluster group
	ClusterGroupId []int32 `json:"cluster_group_id,omitempty" mapstructure:"cluster_group_id,omitempty"`
	// ClusterGroupIdN: Cluster group
	ClusterGroupIdN []int32 `json:"cluster_group_id__n,omitempty" mapstructure:"cluster_group_id__n,omitempty"`
	// ClusterId: Cluster
	ClusterId []int32 `json:"cluster_id,omitempty" mapstructure:"cluster_id,omitempty"`
	// ClusterIdN: Cluster
	ClusterIdN []int32 `json:"cluster_id__n,omitempty" mapstructure:"cluster_id__n,omitempty"`
	// ClusterType: Cluster type (slug)
	ClusterType []string `json:"cluster_type,omitempty" mapstructure:"cluster_type,omitempty"`
	// ClusterTypeN: Cluster type (slug)
	ClusterTypeN []string `json:"cluster_type__n,omitempty" mapstructure:"cluster_type__n,omitempty"`
	// ClusterTypeId: Cluster type
	ClusterTypeId []int32 `json:"cluster_type_id,omitempty" mapstructure:"cluster_type_id,omitempty"`
	// ClusterTypeIdN: Cluster type
	ClusterTypeIdN []int32 `json:"cluster_type_id__n,omitempty" mapstructure:"cluster_type_id__n,omitempty"`
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
	// DataFileId: Data file (ID)
	DataFileId []*int32 `json:"data_file_id,omitempty" mapstructure:"data_file_id,omitempty"`
	// DataFileIdN: Data file (ID)
	DataFileIdN []*int32 `json:"data_file_id__n,omitempty" mapstructure:"data_file_id__n,omitempty"`
	// DataSourceId: Data source (ID)
	DataSourceId []*int32 `json:"data_source_id,omitempty" mapstructure:"data_source_id,omitempty"`
	// DataSourceIdN: Data source (ID)
	DataSourceIdN []*int32 `json:"data_source_id__n,omitempty" mapstructure:"data_source_id__n,omitempty"`
	// DataSynced:
	DataSynced []time.Time `json:"data_synced,omitempty" mapstructure:"data_synced,omitempty"`
	// DataSyncedGt:
	DataSyncedGt []time.Time `json:"data_synced__gt,omitempty" mapstructure:"data_synced__gt,omitempty"`
	// DataSyncedGte:
	DataSyncedGte []time.Time `json:"data_synced__gte,omitempty" mapstructure:"data_synced__gte,omitempty"`
	// DataSyncedLt:
	DataSyncedLt []time.Time `json:"data_synced__lt,omitempty" mapstructure:"data_synced__lt,omitempty"`
	// DataSyncedLte:
	DataSyncedLte []time.Time `json:"data_synced__lte,omitempty" mapstructure:"data_synced__lte,omitempty"`
	// DataSyncedN:
	DataSyncedN []time.Time `json:"data_synced__n,omitempty" mapstructure:"data_synced__n,omitempty"`
	// DeviceTypeId: Device type
	DeviceTypeId []int32 `json:"device_type_id,omitempty" mapstructure:"device_type_id,omitempty"`
	// DeviceTypeIdN: Device type
	DeviceTypeIdN []int32 `json:"device_type_id__n,omitempty" mapstructure:"device_type_id__n,omitempty"`
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
	// Location: Location (slug)
	Location []string `json:"location,omitempty" mapstructure:"location,omitempty"`
	// LocationN: Location (slug)
	LocationN []string `json:"location__n,omitempty" mapstructure:"location__n,omitempty"`
	// LocationId: Location
	LocationId []int32 `json:"location_id,omitempty" mapstructure:"location_id,omitempty"`
	// LocationIdN: Location
	LocationIdN []int32 `json:"location_id__n,omitempty" mapstructure:"location_id__n,omitempty"`
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
	// Platform: Platform (slug)
	Platform []string `json:"platform,omitempty" mapstructure:"platform,omitempty"`
	// PlatformN: Platform (slug)
	PlatformN []string `json:"platform__n,omitempty" mapstructure:"platform__n,omitempty"`
	// PlatformId: Platform
	PlatformId []int32 `json:"platform_id,omitempty" mapstructure:"platform_id,omitempty"`
	// PlatformIdN: Platform
	PlatformIdN []int32 `json:"platform_id__n,omitempty" mapstructure:"platform_id__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Region: Region (slug)
	Region []string `json:"region,omitempty" mapstructure:"region,omitempty"`
	// RegionN: Region (slug)
	RegionN []string `json:"region__n,omitempty" mapstructure:"region__n,omitempty"`
	// RegionId: Region
	RegionId []int32 `json:"region_id,omitempty" mapstructure:"region_id,omitempty"`
	// RegionIdN: Region
	RegionIdN []int32 `json:"region_id__n,omitempty" mapstructure:"region_id__n,omitempty"`
	// Role: Role (slug)
	Role []string `json:"role,omitempty" mapstructure:"role,omitempty"`
	// RoleN: Role (slug)
	RoleN []string `json:"role__n,omitempty" mapstructure:"role__n,omitempty"`
	// RoleId: Role
	RoleId []int32 `json:"role_id,omitempty" mapstructure:"role_id,omitempty"`
	// RoleIdN: Role
	RoleIdN []int32 `json:"role_id__n,omitempty" mapstructure:"role_id__n,omitempty"`
	// Site: Site (slug)
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteN: Site (slug)
	SiteN []string `json:"site__n,omitempty" mapstructure:"site__n,omitempty"`
	// SiteGroup: Site group (slug)
	SiteGroup []string `json:"site_group,omitempty" mapstructure:"site_group,omitempty"`
	// SiteGroupN: Site group (slug)
	SiteGroupN []string `json:"site_group__n,omitempty" mapstructure:"site_group__n,omitempty"`
	// SiteGroupId: Site group
	SiteGroupId []int32 `json:"site_group_id,omitempty" mapstructure:"site_group_id,omitempty"`
	// SiteGroupIdN: Site group
	SiteGroupIdN []int32 `json:"site_group_id__n,omitempty" mapstructure:"site_group_id__n,omitempty"`
	// SiteId: Site
	SiteId []int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
	// SiteIdN: Site
	SiteIdN []int32 `json:"site_id__n,omitempty" mapstructure:"site_id__n,omitempty"`
	// Tag: Tag (slug)
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN: Tag (slug)
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// TagId: Tag
	TagId []int32 `json:"tag_id,omitempty" mapstructure:"tag_id,omitempty"`
	// TagIdN: Tag
	TagIdN []int32 `json:"tag_id__n,omitempty" mapstructure:"tag_id__n,omitempty"`
	// Tenant: Tenant (slug)
	Tenant []string `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TenantN: Tenant (slug)
	TenantN []string `json:"tenant__n,omitempty" mapstructure:"tenant__n,omitempty"`
	// TenantGroup: Tenant group (slug)
	TenantGroup []string `json:"tenant_group,omitempty" mapstructure:"tenant_group,omitempty"`
	// TenantGroupN: Tenant group (slug)
	TenantGroupN []string `json:"tenant_group__n,omitempty" mapstructure:"tenant_group__n,omitempty"`
	// TenantGroupId: Tenant group
	TenantGroupId []int32 `json:"tenant_group_id,omitempty" mapstructure:"tenant_group_id,omitempty"`
	// TenantGroupIdN: Tenant group
	TenantGroupIdN []int32 `json:"tenant_group_id__n,omitempty" mapstructure:"tenant_group_id__n,omitempty"`
	// TenantId: Tenant
	TenantId []int32 `json:"tenant_id,omitempty" mapstructure:"tenant_id,omitempty"`
	// TenantIdN: Tenant
	TenantIdN []int32 `json:"tenant_id__n,omitempty" mapstructure:"tenant_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasConfigContextsListQueryParameters) Validate() error {
	return validation.Errors{
		"clusterGroup": validation.Validate(
			m.ClusterGroup,
		),
		"clusterGroupN": validation.Validate(
			m.ClusterGroupN,
		),
		"clusterGroupId": validation.Validate(
			m.ClusterGroupId,
		),
		"clusterGroupIdN": validation.Validate(
			m.ClusterGroupIdN,
		),
		"clusterId": validation.Validate(
			m.ClusterId,
		),
		"clusterIdN": validation.Validate(
			m.ClusterIdN,
		),
		"clusterType": validation.Validate(
			m.ClusterType,
		),
		"clusterTypeN": validation.Validate(
			m.ClusterTypeN,
		),
		"clusterTypeId": validation.Validate(
			m.ClusterTypeId,
		),
		"clusterTypeIdN": validation.Validate(
			m.ClusterTypeIdN,
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
		"dataFileId": validation.Validate(
			m.DataFileId,
		),
		"dataFileIdN": validation.Validate(
			m.DataFileIdN,
		),
		"dataSourceId": validation.Validate(
			m.DataSourceId,
		),
		"dataSourceIdN": validation.Validate(
			m.DataSourceIdN,
		),
		"dataSynced": validation.Validate(
			m.DataSynced,
		),
		"dataSyncedGt": validation.Validate(
			m.DataSyncedGt,
		),
		"dataSyncedGte": validation.Validate(
			m.DataSyncedGte,
		),
		"dataSyncedLt": validation.Validate(
			m.DataSyncedLt,
		),
		"dataSyncedLte": validation.Validate(
			m.DataSyncedLte,
		),
		"dataSyncedN": validation.Validate(
			m.DataSyncedN,
		),
		"deviceTypeId": validation.Validate(
			m.DeviceTypeId,
		),
		"deviceTypeIdN": validation.Validate(
			m.DeviceTypeIdN,
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
		"location": validation.Validate(
			m.Location,
		),
		"locationN": validation.Validate(
			m.LocationN,
		),
		"locationId": validation.Validate(
			m.LocationId,
		),
		"locationIdN": validation.Validate(
			m.LocationIdN,
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
		"platform": validation.Validate(
			m.Platform,
		),
		"platformN": validation.Validate(
			m.PlatformN,
		),
		"platformId": validation.Validate(
			m.PlatformId,
		),
		"platformIdN": validation.Validate(
			m.PlatformIdN,
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
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"tagId": validation.Validate(
			m.TagId,
		),
		"tagIdN": validation.Validate(
			m.TagIdN,
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

// GetClusterGroup returns the ClusterGroup property
func (m ExtrasConfigContextsListQueryParameters) GetClusterGroup() []string {
	return m.ClusterGroup
}

// SetClusterGroup sets the ClusterGroup property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterGroup(val []string) {
	m.ClusterGroup = val
}

// GetClusterGroupN returns the ClusterGroupN property
func (m ExtrasConfigContextsListQueryParameters) GetClusterGroupN() []string {
	return m.ClusterGroupN
}

// SetClusterGroupN sets the ClusterGroupN property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterGroupN(val []string) {
	m.ClusterGroupN = val
}

// GetClusterGroupId returns the ClusterGroupId property
func (m ExtrasConfigContextsListQueryParameters) GetClusterGroupId() []int32 {
	return m.ClusterGroupId
}

// SetClusterGroupId sets the ClusterGroupId property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterGroupId(val []int32) {
	m.ClusterGroupId = val
}

// GetClusterGroupIdN returns the ClusterGroupIdN property
func (m ExtrasConfigContextsListQueryParameters) GetClusterGroupIdN() []int32 {
	return m.ClusterGroupIdN
}

// SetClusterGroupIdN sets the ClusterGroupIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterGroupIdN(val []int32) {
	m.ClusterGroupIdN = val
}

// GetClusterId returns the ClusterId property
func (m ExtrasConfigContextsListQueryParameters) GetClusterId() []int32 {
	return m.ClusterId
}

// SetClusterId sets the ClusterId property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterId(val []int32) {
	m.ClusterId = val
}

// GetClusterIdN returns the ClusterIdN property
func (m ExtrasConfigContextsListQueryParameters) GetClusterIdN() []int32 {
	return m.ClusterIdN
}

// SetClusterIdN sets the ClusterIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterIdN(val []int32) {
	m.ClusterIdN = val
}

// GetClusterType returns the ClusterType property
func (m ExtrasConfigContextsListQueryParameters) GetClusterType() []string {
	return m.ClusterType
}

// SetClusterType sets the ClusterType property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterType(val []string) {
	m.ClusterType = val
}

// GetClusterTypeN returns the ClusterTypeN property
func (m ExtrasConfigContextsListQueryParameters) GetClusterTypeN() []string {
	return m.ClusterTypeN
}

// SetClusterTypeN sets the ClusterTypeN property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterTypeN(val []string) {
	m.ClusterTypeN = val
}

// GetClusterTypeId returns the ClusterTypeId property
func (m ExtrasConfigContextsListQueryParameters) GetClusterTypeId() []int32 {
	return m.ClusterTypeId
}

// SetClusterTypeId sets the ClusterTypeId property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterTypeId(val []int32) {
	m.ClusterTypeId = val
}

// GetClusterTypeIdN returns the ClusterTypeIdN property
func (m ExtrasConfigContextsListQueryParameters) GetClusterTypeIdN() []int32 {
	return m.ClusterTypeIdN
}

// SetClusterTypeIdN sets the ClusterTypeIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetClusterTypeIdN(val []int32) {
	m.ClusterTypeIdN = val
}

// GetCreated returns the Created property
func (m ExtrasConfigContextsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ExtrasConfigContextsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m ExtrasConfigContextsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *ExtrasConfigContextsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m ExtrasConfigContextsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *ExtrasConfigContextsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m ExtrasConfigContextsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *ExtrasConfigContextsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m ExtrasConfigContextsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *ExtrasConfigContextsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m ExtrasConfigContextsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *ExtrasConfigContextsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m ExtrasConfigContextsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *ExtrasConfigContextsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDataFileId returns the DataFileId property
func (m ExtrasConfigContextsListQueryParameters) GetDataFileId() []*int32 {
	return m.DataFileId
}

// SetDataFileId sets the DataFileId property
func (m *ExtrasConfigContextsListQueryParameters) SetDataFileId(val []*int32) {
	m.DataFileId = val
}

// GetDataFileIdN returns the DataFileIdN property
func (m ExtrasConfigContextsListQueryParameters) GetDataFileIdN() []*int32 {
	return m.DataFileIdN
}

// SetDataFileIdN sets the DataFileIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetDataFileIdN(val []*int32) {
	m.DataFileIdN = val
}

// GetDataSourceId returns the DataSourceId property
func (m ExtrasConfigContextsListQueryParameters) GetDataSourceId() []*int32 {
	return m.DataSourceId
}

// SetDataSourceId sets the DataSourceId property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSourceId(val []*int32) {
	m.DataSourceId = val
}

// GetDataSourceIdN returns the DataSourceIdN property
func (m ExtrasConfigContextsListQueryParameters) GetDataSourceIdN() []*int32 {
	return m.DataSourceIdN
}

// SetDataSourceIdN sets the DataSourceIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSourceIdN(val []*int32) {
	m.DataSourceIdN = val
}

// GetDataSynced returns the DataSynced property
func (m ExtrasConfigContextsListQueryParameters) GetDataSynced() []time.Time {
	return m.DataSynced
}

// SetDataSynced sets the DataSynced property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSynced(val []time.Time) {
	m.DataSynced = val
}

// GetDataSyncedGt returns the DataSyncedGt property
func (m ExtrasConfigContextsListQueryParameters) GetDataSyncedGt() []time.Time {
	return m.DataSyncedGt
}

// SetDataSyncedGt sets the DataSyncedGt property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSyncedGt(val []time.Time) {
	m.DataSyncedGt = val
}

// GetDataSyncedGte returns the DataSyncedGte property
func (m ExtrasConfigContextsListQueryParameters) GetDataSyncedGte() []time.Time {
	return m.DataSyncedGte
}

// SetDataSyncedGte sets the DataSyncedGte property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSyncedGte(val []time.Time) {
	m.DataSyncedGte = val
}

// GetDataSyncedLt returns the DataSyncedLt property
func (m ExtrasConfigContextsListQueryParameters) GetDataSyncedLt() []time.Time {
	return m.DataSyncedLt
}

// SetDataSyncedLt sets the DataSyncedLt property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSyncedLt(val []time.Time) {
	m.DataSyncedLt = val
}

// GetDataSyncedLte returns the DataSyncedLte property
func (m ExtrasConfigContextsListQueryParameters) GetDataSyncedLte() []time.Time {
	return m.DataSyncedLte
}

// SetDataSyncedLte sets the DataSyncedLte property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSyncedLte(val []time.Time) {
	m.DataSyncedLte = val
}

// GetDataSyncedN returns the DataSyncedN property
func (m ExtrasConfigContextsListQueryParameters) GetDataSyncedN() []time.Time {
	return m.DataSyncedN
}

// SetDataSyncedN sets the DataSyncedN property
func (m *ExtrasConfigContextsListQueryParameters) SetDataSyncedN(val []time.Time) {
	m.DataSyncedN = val
}

// GetDeviceTypeId returns the DeviceTypeId property
func (m ExtrasConfigContextsListQueryParameters) GetDeviceTypeId() []int32 {
	return m.DeviceTypeId
}

// SetDeviceTypeId sets the DeviceTypeId property
func (m *ExtrasConfigContextsListQueryParameters) SetDeviceTypeId(val []int32) {
	m.DeviceTypeId = val
}

// GetDeviceTypeIdN returns the DeviceTypeIdN property
func (m ExtrasConfigContextsListQueryParameters) GetDeviceTypeIdN() []int32 {
	return m.DeviceTypeIdN
}

// SetDeviceTypeIdN sets the DeviceTypeIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetDeviceTypeIdN(val []int32) {
	m.DeviceTypeIdN = val
}

// GetId returns the Id property
func (m ExtrasConfigContextsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasConfigContextsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasConfigContextsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasConfigContextsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasConfigContextsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasConfigContextsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasConfigContextsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasConfigContextsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasConfigContextsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasConfigContextsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasConfigContextsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasConfigContextsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetIsActive returns the IsActive property
func (m ExtrasConfigContextsListQueryParameters) GetIsActive() bool {
	return m.IsActive
}

// SetIsActive sets the IsActive property
func (m *ExtrasConfigContextsListQueryParameters) SetIsActive(val bool) {
	m.IsActive = val
}

// GetLastUpdated returns the LastUpdated property
func (m ExtrasConfigContextsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ExtrasConfigContextsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m ExtrasConfigContextsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *ExtrasConfigContextsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m ExtrasConfigContextsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *ExtrasConfigContextsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m ExtrasConfigContextsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *ExtrasConfigContextsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m ExtrasConfigContextsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *ExtrasConfigContextsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m ExtrasConfigContextsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *ExtrasConfigContextsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m ExtrasConfigContextsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasConfigContextsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m ExtrasConfigContextsListQueryParameters) GetLocation() []string {
	return m.Location
}

// SetLocation sets the Location property
func (m *ExtrasConfigContextsListQueryParameters) SetLocation(val []string) {
	m.Location = val
}

// GetLocationN returns the LocationN property
func (m ExtrasConfigContextsListQueryParameters) GetLocationN() []string {
	return m.LocationN
}

// SetLocationN sets the LocationN property
func (m *ExtrasConfigContextsListQueryParameters) SetLocationN(val []string) {
	m.LocationN = val
}

// GetLocationId returns the LocationId property
func (m ExtrasConfigContextsListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *ExtrasConfigContextsListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m ExtrasConfigContextsListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetName returns the Name property
func (m ExtrasConfigContextsListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasConfigContextsListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasConfigContextsListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasConfigContextsListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasConfigContextsListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasConfigContextsListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasConfigContextsListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasConfigContextsListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasConfigContextsListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasConfigContextsListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasConfigContextsListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasConfigContextsListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasConfigContextsListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasConfigContextsListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasConfigContextsListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasConfigContextsListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasConfigContextsListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasConfigContextsListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasConfigContextsListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasConfigContextsListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasConfigContextsListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasConfigContextsListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m ExtrasConfigContextsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasConfigContextsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasConfigContextsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasConfigContextsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPlatform returns the Platform property
func (m ExtrasConfigContextsListQueryParameters) GetPlatform() []string {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *ExtrasConfigContextsListQueryParameters) SetPlatform(val []string) {
	m.Platform = val
}

// GetPlatformN returns the PlatformN property
func (m ExtrasConfigContextsListQueryParameters) GetPlatformN() []string {
	return m.PlatformN
}

// SetPlatformN sets the PlatformN property
func (m *ExtrasConfigContextsListQueryParameters) SetPlatformN(val []string) {
	m.PlatformN = val
}

// GetPlatformId returns the PlatformId property
func (m ExtrasConfigContextsListQueryParameters) GetPlatformId() []int32 {
	return m.PlatformId
}

// SetPlatformId sets the PlatformId property
func (m *ExtrasConfigContextsListQueryParameters) SetPlatformId(val []int32) {
	m.PlatformId = val
}

// GetPlatformIdN returns the PlatformIdN property
func (m ExtrasConfigContextsListQueryParameters) GetPlatformIdN() []int32 {
	return m.PlatformIdN
}

// SetPlatformIdN sets the PlatformIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetPlatformIdN(val []int32) {
	m.PlatformIdN = val
}

// GetQ returns the Q property
func (m ExtrasConfigContextsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasConfigContextsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m ExtrasConfigContextsListQueryParameters) GetRegion() []string {
	return m.Region
}

// SetRegion sets the Region property
func (m *ExtrasConfigContextsListQueryParameters) SetRegion(val []string) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m ExtrasConfigContextsListQueryParameters) GetRegionN() []string {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *ExtrasConfigContextsListQueryParameters) SetRegionN(val []string) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m ExtrasConfigContextsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *ExtrasConfigContextsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m ExtrasConfigContextsListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetRole returns the Role property
func (m ExtrasConfigContextsListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *ExtrasConfigContextsListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m ExtrasConfigContextsListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *ExtrasConfigContextsListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m ExtrasConfigContextsListQueryParameters) GetRoleId() []int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *ExtrasConfigContextsListQueryParameters) SetRoleId(val []int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m ExtrasConfigContextsListQueryParameters) GetRoleIdN() []int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetRoleIdN(val []int32) {
	m.RoleIdN = val
}

// GetSite returns the Site property
func (m ExtrasConfigContextsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *ExtrasConfigContextsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m ExtrasConfigContextsListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *ExtrasConfigContextsListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m ExtrasConfigContextsListQueryParameters) GetSiteGroup() []string {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *ExtrasConfigContextsListQueryParameters) SetSiteGroup(val []string) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m ExtrasConfigContextsListQueryParameters) GetSiteGroupN() []string {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *ExtrasConfigContextsListQueryParameters) SetSiteGroupN(val []string) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m ExtrasConfigContextsListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *ExtrasConfigContextsListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m ExtrasConfigContextsListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m ExtrasConfigContextsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *ExtrasConfigContextsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m ExtrasConfigContextsListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetTag returns the Tag property
func (m ExtrasConfigContextsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *ExtrasConfigContextsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m ExtrasConfigContextsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *ExtrasConfigContextsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTagId returns the TagId property
func (m ExtrasConfigContextsListQueryParameters) GetTagId() []int32 {
	return m.TagId
}

// SetTagId sets the TagId property
func (m *ExtrasConfigContextsListQueryParameters) SetTagId(val []int32) {
	m.TagId = val
}

// GetTagIdN returns the TagIdN property
func (m ExtrasConfigContextsListQueryParameters) GetTagIdN() []int32 {
	return m.TagIdN
}

// SetTagIdN sets the TagIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetTagIdN(val []int32) {
	m.TagIdN = val
}

// GetTenant returns the Tenant property
func (m ExtrasConfigContextsListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *ExtrasConfigContextsListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m ExtrasConfigContextsListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *ExtrasConfigContextsListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m ExtrasConfigContextsListQueryParameters) GetTenantGroup() []string {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *ExtrasConfigContextsListQueryParameters) SetTenantGroup(val []string) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m ExtrasConfigContextsListQueryParameters) GetTenantGroupN() []string {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *ExtrasConfigContextsListQueryParameters) SetTenantGroupN(val []string) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m ExtrasConfigContextsListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *ExtrasConfigContextsListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m ExtrasConfigContextsListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m ExtrasConfigContextsListQueryParameters) GetTenantId() []int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *ExtrasConfigContextsListQueryParameters) SetTenantId(val []int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m ExtrasConfigContextsListQueryParameters) GetTenantIdN() []int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *ExtrasConfigContextsListQueryParameters) SetTenantIdN(val []int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m ExtrasConfigContextsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *ExtrasConfigContextsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
