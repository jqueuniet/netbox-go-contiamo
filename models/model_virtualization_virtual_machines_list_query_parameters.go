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

// VirtualizationVirtualMachinesListQueryParameters is an object.
type VirtualizationVirtualMachinesListQueryParameters struct {
	// Cluster: Cluster
	Cluster []string `json:"cluster,omitempty" mapstructure:"cluster,omitempty"`
	// ClusterN: Cluster
	ClusterN []string `json:"cluster__n,omitempty" mapstructure:"cluster__n,omitempty"`
	// ClusterGroup: Cluster group (slug)
	ClusterGroup []string `json:"cluster_group,omitempty" mapstructure:"cluster_group,omitempty"`
	// ClusterGroupN: Cluster group (slug)
	ClusterGroupN []string `json:"cluster_group__n,omitempty" mapstructure:"cluster_group__n,omitempty"`
	// ClusterGroupId: Cluster group (ID)
	ClusterGroupId []int32 `json:"cluster_group_id,omitempty" mapstructure:"cluster_group_id,omitempty"`
	// ClusterGroupIdN: Cluster group (ID)
	ClusterGroupIdN []int32 `json:"cluster_group_id__n,omitempty" mapstructure:"cluster_group_id__n,omitempty"`
	// ClusterId: Cluster (ID)
	ClusterId []*int32 `json:"cluster_id,omitempty" mapstructure:"cluster_id,omitempty"`
	// ClusterIdN: Cluster (ID)
	ClusterIdN []*int32 `json:"cluster_id__n,omitempty" mapstructure:"cluster_id__n,omitempty"`
	// ClusterType: Cluster type (slug)
	ClusterType []string `json:"cluster_type,omitempty" mapstructure:"cluster_type,omitempty"`
	// ClusterTypeN: Cluster type (slug)
	ClusterTypeN []string `json:"cluster_type__n,omitempty" mapstructure:"cluster_type__n,omitempty"`
	// ClusterTypeId: Cluster type (ID)
	ClusterTypeId []int32 `json:"cluster_type_id,omitempty" mapstructure:"cluster_type_id,omitempty"`
	// ClusterTypeIdN: Cluster type (ID)
	ClusterTypeIdN []int32 `json:"cluster_type_id__n,omitempty" mapstructure:"cluster_type_id__n,omitempty"`
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
	// Device: Device
	Device []*string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceN: Device
	DeviceN []*string `json:"device__n,omitempty" mapstructure:"device__n,omitempty"`
	// DeviceId: Device (ID)
	DeviceId []*int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
	// DeviceIdN: Device (ID)
	DeviceIdN []*int32 `json:"device_id__n,omitempty" mapstructure:"device_id__n,omitempty"`
	// Disk:
	Disk []int32 `json:"disk,omitempty" mapstructure:"disk,omitempty"`
	// DiskGt:
	DiskGt []int32 `json:"disk__gt,omitempty" mapstructure:"disk__gt,omitempty"`
	// DiskGte:
	DiskGte []int32 `json:"disk__gte,omitempty" mapstructure:"disk__gte,omitempty"`
	// DiskLt:
	DiskLt []int32 `json:"disk__lt,omitempty" mapstructure:"disk__lt,omitempty"`
	// DiskLte:
	DiskLte []int32 `json:"disk__lte,omitempty" mapstructure:"disk__lte,omitempty"`
	// DiskN:
	DiskN []int32 `json:"disk__n,omitempty" mapstructure:"disk__n,omitempty"`
	// HasPrimaryIp: Has a primary IP
	HasPrimaryIp bool `json:"has_primary_ip,omitempty" mapstructure:"has_primary_ip,omitempty"`
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
	// LocalContextData: Has local config context data
	LocalContextData bool `json:"local_context_data,omitempty" mapstructure:"local_context_data,omitempty"`
	// MacAddress:
	MacAddress []string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// MacAddressIc:
	MacAddressIc []string `json:"mac_address__ic,omitempty" mapstructure:"mac_address__ic,omitempty"`
	// MacAddressIe:
	MacAddressIe []string `json:"mac_address__ie,omitempty" mapstructure:"mac_address__ie,omitempty"`
	// MacAddressIew:
	MacAddressIew []string `json:"mac_address__iew,omitempty" mapstructure:"mac_address__iew,omitempty"`
	// MacAddressIsw:
	MacAddressIsw []string `json:"mac_address__isw,omitempty" mapstructure:"mac_address__isw,omitempty"`
	// MacAddressN:
	MacAddressN []string `json:"mac_address__n,omitempty" mapstructure:"mac_address__n,omitempty"`
	// MacAddressNic:
	MacAddressNic []string `json:"mac_address__nic,omitempty" mapstructure:"mac_address__nic,omitempty"`
	// MacAddressNie:
	MacAddressNie []string `json:"mac_address__nie,omitempty" mapstructure:"mac_address__nie,omitempty"`
	// MacAddressNiew:
	MacAddressNiew []string `json:"mac_address__niew,omitempty" mapstructure:"mac_address__niew,omitempty"`
	// MacAddressNisw:
	MacAddressNisw []string `json:"mac_address__nisw,omitempty" mapstructure:"mac_address__nisw,omitempty"`
	// Memory:
	Memory []int32 `json:"memory,omitempty" mapstructure:"memory,omitempty"`
	// MemoryGt:
	MemoryGt []int32 `json:"memory__gt,omitempty" mapstructure:"memory__gt,omitempty"`
	// MemoryGte:
	MemoryGte []int32 `json:"memory__gte,omitempty" mapstructure:"memory__gte,omitempty"`
	// MemoryLt:
	MemoryLt []int32 `json:"memory__lt,omitempty" mapstructure:"memory__lt,omitempty"`
	// MemoryLte:
	MemoryLte []int32 `json:"memory__lte,omitempty" mapstructure:"memory__lte,omitempty"`
	// MemoryN:
	MemoryN []int32 `json:"memory__n,omitempty" mapstructure:"memory__n,omitempty"`
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
	// PlatformId: Platform (ID)
	PlatformId []*int32 `json:"platform_id,omitempty" mapstructure:"platform_id,omitempty"`
	// PlatformIdN: Platform (ID)
	PlatformIdN []*int32 `json:"platform_id__n,omitempty" mapstructure:"platform_id__n,omitempty"`
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
	// Status: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
	// * `decommissioning` - Decommissioning
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `staged` - Staged
	// * `failed` - Failed
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
	// Vcpus:
	Vcpus []float64 `json:"vcpus,omitempty" mapstructure:"vcpus,omitempty"`
	// VcpusGt:
	VcpusGt []float64 `json:"vcpus__gt,omitempty" mapstructure:"vcpus__gt,omitempty"`
	// VcpusGte:
	VcpusGte []float64 `json:"vcpus__gte,omitempty" mapstructure:"vcpus__gte,omitempty"`
	// VcpusLt:
	VcpusLt []float64 `json:"vcpus__lt,omitempty" mapstructure:"vcpus__lt,omitempty"`
	// VcpusLte:
	VcpusLte []float64 `json:"vcpus__lte,omitempty" mapstructure:"vcpus__lte,omitempty"`
	// VcpusN:
	VcpusN []float64 `json:"vcpus__n,omitempty" mapstructure:"vcpus__n,omitempty"`
}

// Validate implements basic validation for this model
func (m VirtualizationVirtualMachinesListQueryParameters) Validate() error {
	return validation.Errors{
		"cluster": validation.Validate(
			m.Cluster,
		),
		"clusterN": validation.Validate(
			m.ClusterN,
		),
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
		"device": validation.Validate(
			m.Device,
		),
		"deviceN": validation.Validate(
			m.DeviceN,
		),
		"deviceId": validation.Validate(
			m.DeviceId,
		),
		"deviceIdN": validation.Validate(
			m.DeviceIdN,
		),
		"disk": validation.Validate(
			m.Disk,
		),
		"diskGt": validation.Validate(
			m.DiskGt,
		),
		"diskGte": validation.Validate(
			m.DiskGte,
		),
		"diskLt": validation.Validate(
			m.DiskLt,
		),
		"diskLte": validation.Validate(
			m.DiskLte,
		),
		"diskN": validation.Validate(
			m.DiskN,
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
		"macAddress": validation.Validate(
			m.MacAddress,
		),
		"macAddressIc": validation.Validate(
			m.MacAddressIc,
		),
		"macAddressIe": validation.Validate(
			m.MacAddressIe,
		),
		"macAddressIew": validation.Validate(
			m.MacAddressIew,
		),
		"macAddressIsw": validation.Validate(
			m.MacAddressIsw,
		),
		"macAddressN": validation.Validate(
			m.MacAddressN,
		),
		"macAddressNic": validation.Validate(
			m.MacAddressNic,
		),
		"macAddressNie": validation.Validate(
			m.MacAddressNie,
		),
		"macAddressNiew": validation.Validate(
			m.MacAddressNiew,
		),
		"macAddressNisw": validation.Validate(
			m.MacAddressNisw,
		),
		"memory": validation.Validate(
			m.Memory,
		),
		"memoryGt": validation.Validate(
			m.MemoryGt,
		),
		"memoryGte": validation.Validate(
			m.MemoryGte,
		),
		"memoryLt": validation.Validate(
			m.MemoryLt,
		),
		"memoryLte": validation.Validate(
			m.MemoryLte,
		),
		"memoryN": validation.Validate(
			m.MemoryN,
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
		"vcpus": validation.Validate(
			m.Vcpus,
		),
		"vcpusGt": validation.Validate(
			m.VcpusGt,
		),
		"vcpusGte": validation.Validate(
			m.VcpusGte,
		),
		"vcpusLt": validation.Validate(
			m.VcpusLt,
		),
		"vcpusLte": validation.Validate(
			m.VcpusLte,
		),
		"vcpusN": validation.Validate(
			m.VcpusN,
		),
	}.Filter()
}

// GetCluster returns the Cluster property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCluster() []string {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCluster(val []string) {
	m.Cluster = val
}

// GetClusterN returns the ClusterN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterN() []string {
	return m.ClusterN
}

// SetClusterN sets the ClusterN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterN(val []string) {
	m.ClusterN = val
}

// GetClusterGroup returns the ClusterGroup property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterGroup() []string {
	return m.ClusterGroup
}

// SetClusterGroup sets the ClusterGroup property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterGroup(val []string) {
	m.ClusterGroup = val
}

// GetClusterGroupN returns the ClusterGroupN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterGroupN() []string {
	return m.ClusterGroupN
}

// SetClusterGroupN sets the ClusterGroupN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterGroupN(val []string) {
	m.ClusterGroupN = val
}

// GetClusterGroupId returns the ClusterGroupId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterGroupId() []int32 {
	return m.ClusterGroupId
}

// SetClusterGroupId sets the ClusterGroupId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterGroupId(val []int32) {
	m.ClusterGroupId = val
}

// GetClusterGroupIdN returns the ClusterGroupIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterGroupIdN() []int32 {
	return m.ClusterGroupIdN
}

// SetClusterGroupIdN sets the ClusterGroupIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterGroupIdN(val []int32) {
	m.ClusterGroupIdN = val
}

// GetClusterId returns the ClusterId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterId() []*int32 {
	return m.ClusterId
}

// SetClusterId sets the ClusterId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterId(val []*int32) {
	m.ClusterId = val
}

// GetClusterIdN returns the ClusterIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterIdN() []*int32 {
	return m.ClusterIdN
}

// SetClusterIdN sets the ClusterIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterIdN(val []*int32) {
	m.ClusterIdN = val
}

// GetClusterType returns the ClusterType property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterType() []string {
	return m.ClusterType
}

// SetClusterType sets the ClusterType property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterType(val []string) {
	m.ClusterType = val
}

// GetClusterTypeN returns the ClusterTypeN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterTypeN() []string {
	return m.ClusterTypeN
}

// SetClusterTypeN sets the ClusterTypeN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterTypeN(val []string) {
	m.ClusterTypeN = val
}

// GetClusterTypeId returns the ClusterTypeId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterTypeId() []int32 {
	return m.ClusterTypeId
}

// SetClusterTypeId sets the ClusterTypeId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterTypeId(val []int32) {
	m.ClusterTypeId = val
}

// GetClusterTypeIdN returns the ClusterTypeIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetClusterTypeIdN() []int32 {
	return m.ClusterTypeIdN
}

// SetClusterTypeIdN sets the ClusterTypeIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetClusterTypeIdN(val []int32) {
	m.ClusterTypeIdN = val
}

// GetContact returns the Contact property
func (m VirtualizationVirtualMachinesListQueryParameters) GetContact() []int32 {
	return m.Contact
}

// SetContact sets the Contact property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetContact(val []int32) {
	m.Contact = val
}

// GetContactN returns the ContactN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetContactN() []int32 {
	return m.ContactN
}

// SetContactN sets the ContactN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetContactN(val []int32) {
	m.ContactN = val
}

// GetContactGroup returns the ContactGroup property
func (m VirtualizationVirtualMachinesListQueryParameters) GetContactGroup() []int32 {
	return m.ContactGroup
}

// SetContactGroup sets the ContactGroup property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetContactGroup(val []int32) {
	m.ContactGroup = val
}

// GetContactGroupN returns the ContactGroupN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetContactGroupN() []int32 {
	return m.ContactGroupN
}

// SetContactGroupN sets the ContactGroupN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetContactGroupN(val []int32) {
	m.ContactGroupN = val
}

// GetContactRole returns the ContactRole property
func (m VirtualizationVirtualMachinesListQueryParameters) GetContactRole() []int32 {
	return m.ContactRole
}

// SetContactRole sets the ContactRole property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetContactRole(val []int32) {
	m.ContactRole = val
}

// GetContactRoleN returns the ContactRoleN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetContactRoleN() []int32 {
	return m.ContactRoleN
}

// SetContactRoleN sets the ContactRoleN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetContactRoleN(val []int32) {
	m.ContactRoleN = val
}

// GetCreated returns the Created property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m VirtualizationVirtualMachinesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevice returns the Device property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDevice() []*string {
	return m.Device
}

// SetDevice sets the Device property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDevice(val []*string) {
	m.Device = val
}

// GetDeviceN returns the DeviceN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDeviceN() []*string {
	return m.DeviceN
}

// SetDeviceN sets the DeviceN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDeviceN(val []*string) {
	m.DeviceN = val
}

// GetDeviceId returns the DeviceId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDeviceId() []*int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDeviceId(val []*int32) {
	m.DeviceId = val
}

// GetDeviceIdN returns the DeviceIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDeviceIdN() []*int32 {
	return m.DeviceIdN
}

// SetDeviceIdN sets the DeviceIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDeviceIdN(val []*int32) {
	m.DeviceIdN = val
}

// GetDisk returns the Disk property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDisk() []int32 {
	return m.Disk
}

// SetDisk sets the Disk property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDisk(val []int32) {
	m.Disk = val
}

// GetDiskGt returns the DiskGt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDiskGt() []int32 {
	return m.DiskGt
}

// SetDiskGt sets the DiskGt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDiskGt(val []int32) {
	m.DiskGt = val
}

// GetDiskGte returns the DiskGte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDiskGte() []int32 {
	return m.DiskGte
}

// SetDiskGte sets the DiskGte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDiskGte(val []int32) {
	m.DiskGte = val
}

// GetDiskLt returns the DiskLt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDiskLt() []int32 {
	return m.DiskLt
}

// SetDiskLt sets the DiskLt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDiskLt(val []int32) {
	m.DiskLt = val
}

// GetDiskLte returns the DiskLte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDiskLte() []int32 {
	return m.DiskLte
}

// SetDiskLte sets the DiskLte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDiskLte(val []int32) {
	m.DiskLte = val
}

// GetDiskN returns the DiskN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetDiskN() []int32 {
	return m.DiskN
}

// SetDiskN sets the DiskN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetDiskN(val []int32) {
	m.DiskN = val
}

// GetHasPrimaryIp returns the HasPrimaryIp property
func (m VirtualizationVirtualMachinesListQueryParameters) GetHasPrimaryIp() bool {
	return m.HasPrimaryIp
}

// SetHasPrimaryIp sets the HasPrimaryIp property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetHasPrimaryIp(val bool) {
	m.HasPrimaryIp = val
}

// GetId returns the Id property
func (m VirtualizationVirtualMachinesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocalContextData returns the LocalContextData property
func (m VirtualizationVirtualMachinesListQueryParameters) GetLocalContextData() bool {
	return m.LocalContextData
}

// SetLocalContextData sets the LocalContextData property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetLocalContextData(val bool) {
	m.LocalContextData = val
}

// GetMacAddress returns the MacAddress property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddress() []string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddress(val []string) {
	m.MacAddress = val
}

// GetMacAddressIc returns the MacAddressIc property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressIc() []string {
	return m.MacAddressIc
}

// SetMacAddressIc sets the MacAddressIc property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressIc(val []string) {
	m.MacAddressIc = val
}

// GetMacAddressIe returns the MacAddressIe property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressIe() []string {
	return m.MacAddressIe
}

// SetMacAddressIe sets the MacAddressIe property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressIe(val []string) {
	m.MacAddressIe = val
}

// GetMacAddressIew returns the MacAddressIew property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressIew() []string {
	return m.MacAddressIew
}

// SetMacAddressIew sets the MacAddressIew property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressIew(val []string) {
	m.MacAddressIew = val
}

// GetMacAddressIsw returns the MacAddressIsw property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressIsw() []string {
	return m.MacAddressIsw
}

// SetMacAddressIsw sets the MacAddressIsw property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressIsw(val []string) {
	m.MacAddressIsw = val
}

// GetMacAddressN returns the MacAddressN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressN() []string {
	return m.MacAddressN
}

// SetMacAddressN sets the MacAddressN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressN(val []string) {
	m.MacAddressN = val
}

// GetMacAddressNic returns the MacAddressNic property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressNic() []string {
	return m.MacAddressNic
}

// SetMacAddressNic sets the MacAddressNic property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressNic(val []string) {
	m.MacAddressNic = val
}

// GetMacAddressNie returns the MacAddressNie property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressNie() []string {
	return m.MacAddressNie
}

// SetMacAddressNie sets the MacAddressNie property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressNie(val []string) {
	m.MacAddressNie = val
}

// GetMacAddressNiew returns the MacAddressNiew property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressNiew() []string {
	return m.MacAddressNiew
}

// SetMacAddressNiew sets the MacAddressNiew property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressNiew(val []string) {
	m.MacAddressNiew = val
}

// GetMacAddressNisw returns the MacAddressNisw property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMacAddressNisw() []string {
	return m.MacAddressNisw
}

// SetMacAddressNisw sets the MacAddressNisw property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMacAddressNisw(val []string) {
	m.MacAddressNisw = val
}

// GetMemory returns the Memory property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMemory() []int32 {
	return m.Memory
}

// SetMemory sets the Memory property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMemory(val []int32) {
	m.Memory = val
}

// GetMemoryGt returns the MemoryGt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMemoryGt() []int32 {
	return m.MemoryGt
}

// SetMemoryGt sets the MemoryGt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMemoryGt(val []int32) {
	m.MemoryGt = val
}

// GetMemoryGte returns the MemoryGte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMemoryGte() []int32 {
	return m.MemoryGte
}

// SetMemoryGte sets the MemoryGte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMemoryGte(val []int32) {
	m.MemoryGte = val
}

// GetMemoryLt returns the MemoryLt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMemoryLt() []int32 {
	return m.MemoryLt
}

// SetMemoryLt sets the MemoryLt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMemoryLt(val []int32) {
	m.MemoryLt = val
}

// GetMemoryLte returns the MemoryLte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMemoryLte() []int32 {
	return m.MemoryLte
}

// SetMemoryLte sets the MemoryLte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMemoryLte(val []int32) {
	m.MemoryLte = val
}

// GetMemoryN returns the MemoryN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetMemoryN() []int32 {
	return m.MemoryN
}

// SetMemoryN sets the MemoryN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetMemoryN(val []int32) {
	m.MemoryN = val
}

// GetName returns the Name property
func (m VirtualizationVirtualMachinesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m VirtualizationVirtualMachinesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m VirtualizationVirtualMachinesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m VirtualizationVirtualMachinesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPlatform returns the Platform property
func (m VirtualizationVirtualMachinesListQueryParameters) GetPlatform() []string {
	return m.Platform
}

// SetPlatform sets the Platform property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetPlatform(val []string) {
	m.Platform = val
}

// GetPlatformN returns the PlatformN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetPlatformN() []string {
	return m.PlatformN
}

// SetPlatformN sets the PlatformN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetPlatformN(val []string) {
	m.PlatformN = val
}

// GetPlatformId returns the PlatformId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetPlatformId() []*int32 {
	return m.PlatformId
}

// SetPlatformId sets the PlatformId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetPlatformId(val []*int32) {
	m.PlatformId = val
}

// GetPlatformIdN returns the PlatformIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetPlatformIdN() []*int32 {
	return m.PlatformIdN
}

// SetPlatformIdN sets the PlatformIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetPlatformIdN(val []*int32) {
	m.PlatformIdN = val
}

// GetQ returns the Q property
func (m VirtualizationVirtualMachinesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetRole returns the Role property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRoleId() []*int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRoleId(val []*int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetRoleIdN() []*int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetRoleIdN(val []*int32) {
	m.RoleIdN = val
}

// GetSite returns the Site property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSiteId() []*int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSiteId(val []*int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetSiteIdN() []*int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetSiteIdN(val []*int32) {
	m.SiteIdN = val
}

// GetStatus returns the Status property
func (m VirtualizationVirtualMachinesListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m VirtualizationVirtualMachinesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVcpus returns the Vcpus property
func (m VirtualizationVirtualMachinesListQueryParameters) GetVcpus() []float64 {
	return m.Vcpus
}

// SetVcpus sets the Vcpus property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetVcpus(val []float64) {
	m.Vcpus = val
}

// GetVcpusGt returns the VcpusGt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetVcpusGt() []float64 {
	return m.VcpusGt
}

// SetVcpusGt sets the VcpusGt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetVcpusGt(val []float64) {
	m.VcpusGt = val
}

// GetVcpusGte returns the VcpusGte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetVcpusGte() []float64 {
	return m.VcpusGte
}

// SetVcpusGte sets the VcpusGte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetVcpusGte(val []float64) {
	m.VcpusGte = val
}

// GetVcpusLt returns the VcpusLt property
func (m VirtualizationVirtualMachinesListQueryParameters) GetVcpusLt() []float64 {
	return m.VcpusLt
}

// SetVcpusLt sets the VcpusLt property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetVcpusLt(val []float64) {
	m.VcpusLt = val
}

// GetVcpusLte returns the VcpusLte property
func (m VirtualizationVirtualMachinesListQueryParameters) GetVcpusLte() []float64 {
	return m.VcpusLte
}

// SetVcpusLte sets the VcpusLte property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetVcpusLte(val []float64) {
	m.VcpusLte = val
}

// GetVcpusN returns the VcpusN property
func (m VirtualizationVirtualMachinesListQueryParameters) GetVcpusN() []float64 {
	return m.VcpusN
}

// SetVcpusN sets the VcpusN property
func (m *VirtualizationVirtualMachinesListQueryParameters) SetVcpusN(val []float64) {
	m.VcpusN = val
}
