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

// PatchedWritableConfigContextRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableConfigContextRequest struct {
	// ClusterGroups:
	ClusterGroups []int32 `json:"cluster_groups,omitempty" mapstructure:"cluster_groups,omitempty"`
	// ClusterTypes:
	ClusterTypes []int32 `json:"cluster_types,omitempty" mapstructure:"cluster_types,omitempty"`
	// Clusters:
	Clusters []int32 `json:"clusters,omitempty" mapstructure:"clusters,omitempty"`
	// Data:
	Data map[string]interface{} `json:"data,omitempty" mapstructure:"data,omitempty"`
	// DataSource: Remote data source
	DataSource *int32 `json:"data_source,omitempty" mapstructure:"data_source,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceTypes:
	DeviceTypes []int32 `json:"device_types,omitempty" mapstructure:"device_types,omitempty"`
	// IsActive:
	IsActive bool `json:"is_active,omitempty" mapstructure:"is_active,omitempty"`
	// Locations:
	Locations []int32 `json:"locations,omitempty" mapstructure:"locations,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Platforms:
	Platforms []int32 `json:"platforms,omitempty" mapstructure:"platforms,omitempty"`
	// Regions:
	Regions []int32 `json:"regions,omitempty" mapstructure:"regions,omitempty"`
	// Roles:
	Roles []int32 `json:"roles,omitempty" mapstructure:"roles,omitempty"`
	// SiteGroups:
	SiteGroups []int32 `json:"site_groups,omitempty" mapstructure:"site_groups,omitempty"`
	// Sites:
	Sites []int32 `json:"sites,omitempty" mapstructure:"sites,omitempty"`
	// Tags:
	Tags []string `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TenantGroups:
	TenantGroups []int32 `json:"tenant_groups,omitempty" mapstructure:"tenant_groups,omitempty"`
	// Tenants:
	Tenants []int32 `json:"tenants,omitempty" mapstructure:"tenants,omitempty"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableConfigContextRequest) Validate() error {
	return validation.Errors{
		"clusterGroups": validation.Validate(
			m.ClusterGroups,
		),
		"clusterTypes": validation.Validate(
			m.ClusterTypes,
		),
		"clusters": validation.Validate(
			m.Clusters,
		),
		"data": validation.Validate(
			m.Data,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"deviceTypes": validation.Validate(
			m.DeviceTypes,
		),
		"locations": validation.Validate(
			m.Locations,
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"platforms": validation.Validate(
			m.Platforms,
		),
		"regions": validation.Validate(
			m.Regions,
		),
		"roles": validation.Validate(
			m.Roles,
		),
		"siteGroups": validation.Validate(
			m.SiteGroups,
		),
		"sites": validation.Validate(
			m.Sites,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenantGroups": validation.Validate(
			m.TenantGroups,
		),
		"tenants": validation.Validate(
			m.Tenants,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetClusterGroups returns the ClusterGroups property
func (m PatchedWritableConfigContextRequest) GetClusterGroups() []int32 {
	return m.ClusterGroups
}

// SetClusterGroups sets the ClusterGroups property
func (m *PatchedWritableConfigContextRequest) SetClusterGroups(val []int32) {
	m.ClusterGroups = val
}

// GetClusterTypes returns the ClusterTypes property
func (m PatchedWritableConfigContextRequest) GetClusterTypes() []int32 {
	return m.ClusterTypes
}

// SetClusterTypes sets the ClusterTypes property
func (m *PatchedWritableConfigContextRequest) SetClusterTypes(val []int32) {
	m.ClusterTypes = val
}

// GetClusters returns the Clusters property
func (m PatchedWritableConfigContextRequest) GetClusters() []int32 {
	return m.Clusters
}

// SetClusters sets the Clusters property
func (m *PatchedWritableConfigContextRequest) SetClusters(val []int32) {
	m.Clusters = val
}

// GetData returns the Data property
func (m PatchedWritableConfigContextRequest) GetData() map[string]interface{} {
	return m.Data
}

// SetData sets the Data property
func (m *PatchedWritableConfigContextRequest) SetData(val map[string]interface{}) {
	m.Data = val
}

// GetDataSource returns the DataSource property
func (m PatchedWritableConfigContextRequest) GetDataSource() *int32 {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *PatchedWritableConfigContextRequest) SetDataSource(val *int32) {
	m.DataSource = val
}

// GetDescription returns the Description property
func (m PatchedWritableConfigContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableConfigContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceTypes returns the DeviceTypes property
func (m PatchedWritableConfigContextRequest) GetDeviceTypes() []int32 {
	return m.DeviceTypes
}

// SetDeviceTypes sets the DeviceTypes property
func (m *PatchedWritableConfigContextRequest) SetDeviceTypes(val []int32) {
	m.DeviceTypes = val
}

// GetIsActive returns the IsActive property
func (m PatchedWritableConfigContextRequest) GetIsActive() bool {
	return m.IsActive
}

// SetIsActive sets the IsActive property
func (m *PatchedWritableConfigContextRequest) SetIsActive(val bool) {
	m.IsActive = val
}

// GetLocations returns the Locations property
func (m PatchedWritableConfigContextRequest) GetLocations() []int32 {
	return m.Locations
}

// SetLocations sets the Locations property
func (m *PatchedWritableConfigContextRequest) SetLocations(val []int32) {
	m.Locations = val
}

// GetName returns the Name property
func (m PatchedWritableConfigContextRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableConfigContextRequest) SetName(val string) {
	m.Name = val
}

// GetPlatforms returns the Platforms property
func (m PatchedWritableConfigContextRequest) GetPlatforms() []int32 {
	return m.Platforms
}

// SetPlatforms sets the Platforms property
func (m *PatchedWritableConfigContextRequest) SetPlatforms(val []int32) {
	m.Platforms = val
}

// GetRegions returns the Regions property
func (m PatchedWritableConfigContextRequest) GetRegions() []int32 {
	return m.Regions
}

// SetRegions sets the Regions property
func (m *PatchedWritableConfigContextRequest) SetRegions(val []int32) {
	m.Regions = val
}

// GetRoles returns the Roles property
func (m PatchedWritableConfigContextRequest) GetRoles() []int32 {
	return m.Roles
}

// SetRoles sets the Roles property
func (m *PatchedWritableConfigContextRequest) SetRoles(val []int32) {
	m.Roles = val
}

// GetSiteGroups returns the SiteGroups property
func (m PatchedWritableConfigContextRequest) GetSiteGroups() []int32 {
	return m.SiteGroups
}

// SetSiteGroups sets the SiteGroups property
func (m *PatchedWritableConfigContextRequest) SetSiteGroups(val []int32) {
	m.SiteGroups = val
}

// GetSites returns the Sites property
func (m PatchedWritableConfigContextRequest) GetSites() []int32 {
	return m.Sites
}

// SetSites sets the Sites property
func (m *PatchedWritableConfigContextRequest) SetSites(val []int32) {
	m.Sites = val
}

// GetTags returns the Tags property
func (m PatchedWritableConfigContextRequest) GetTags() []string {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableConfigContextRequest) SetTags(val []string) {
	m.Tags = val
}

// GetTenantGroups returns the TenantGroups property
func (m PatchedWritableConfigContextRequest) GetTenantGroups() []int32 {
	return m.TenantGroups
}

// SetTenantGroups sets the TenantGroups property
func (m *PatchedWritableConfigContextRequest) SetTenantGroups(val []int32) {
	m.TenantGroups = val
}

// GetTenants returns the Tenants property
func (m PatchedWritableConfigContextRequest) GetTenants() []int32 {
	return m.Tenants
}

// SetTenants sets the Tenants property
func (m *PatchedWritableConfigContextRequest) SetTenants(val []int32) {
	m.Tenants = val
}

// GetWeight returns the Weight property
func (m PatchedWritableConfigContextRequest) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *PatchedWritableConfigContextRequest) SetWeight(val int32) {
	m.Weight = val
}
