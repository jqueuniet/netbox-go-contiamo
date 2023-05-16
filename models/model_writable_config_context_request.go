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

// WritableConfigContextRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableConfigContextRequest struct {
	// ClusterGroups:
	ClusterGroups []int32 `json:"cluster_groups,omitempty" mapstructure:"cluster_groups,omitempty"`
	// ClusterTypes:
	ClusterTypes []int32 `json:"cluster_types,omitempty" mapstructure:"cluster_types,omitempty"`
	// Clusters:
	Clusters []int32 `json:"clusters,omitempty" mapstructure:"clusters,omitempty"`
	// Data:
	Data map[string]interface{} `json:"data" mapstructure:"data"`
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
	Name string `json:"name" mapstructure:"name"`
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
func (m WritableConfigContextRequest) Validate() error {
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
			m.Data, validation.NotNil,
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
			m.Name, validation.Required, validation.Length(1, 100),
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
func (m WritableConfigContextRequest) GetClusterGroups() []int32 {
	return m.ClusterGroups
}

// SetClusterGroups sets the ClusterGroups property
func (m *WritableConfigContextRequest) SetClusterGroups(val []int32) {
	m.ClusterGroups = val
}

// GetClusterTypes returns the ClusterTypes property
func (m WritableConfigContextRequest) GetClusterTypes() []int32 {
	return m.ClusterTypes
}

// SetClusterTypes sets the ClusterTypes property
func (m *WritableConfigContextRequest) SetClusterTypes(val []int32) {
	m.ClusterTypes = val
}

// GetClusters returns the Clusters property
func (m WritableConfigContextRequest) GetClusters() []int32 {
	return m.Clusters
}

// SetClusters sets the Clusters property
func (m *WritableConfigContextRequest) SetClusters(val []int32) {
	m.Clusters = val
}

// GetData returns the Data property
func (m WritableConfigContextRequest) GetData() map[string]interface{} {
	return m.Data
}

// SetData sets the Data property
func (m *WritableConfigContextRequest) SetData(val map[string]interface{}) {
	m.Data = val
}

// GetDataSource returns the DataSource property
func (m WritableConfigContextRequest) GetDataSource() *int32 {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *WritableConfigContextRequest) SetDataSource(val *int32) {
	m.DataSource = val
}

// GetDescription returns the Description property
func (m WritableConfigContextRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableConfigContextRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceTypes returns the DeviceTypes property
func (m WritableConfigContextRequest) GetDeviceTypes() []int32 {
	return m.DeviceTypes
}

// SetDeviceTypes sets the DeviceTypes property
func (m *WritableConfigContextRequest) SetDeviceTypes(val []int32) {
	m.DeviceTypes = val
}

// GetIsActive returns the IsActive property
func (m WritableConfigContextRequest) GetIsActive() bool {
	return m.IsActive
}

// SetIsActive sets the IsActive property
func (m *WritableConfigContextRequest) SetIsActive(val bool) {
	m.IsActive = val
}

// GetLocations returns the Locations property
func (m WritableConfigContextRequest) GetLocations() []int32 {
	return m.Locations
}

// SetLocations sets the Locations property
func (m *WritableConfigContextRequest) SetLocations(val []int32) {
	m.Locations = val
}

// GetName returns the Name property
func (m WritableConfigContextRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableConfigContextRequest) SetName(val string) {
	m.Name = val
}

// GetPlatforms returns the Platforms property
func (m WritableConfigContextRequest) GetPlatforms() []int32 {
	return m.Platforms
}

// SetPlatforms sets the Platforms property
func (m *WritableConfigContextRequest) SetPlatforms(val []int32) {
	m.Platforms = val
}

// GetRegions returns the Regions property
func (m WritableConfigContextRequest) GetRegions() []int32 {
	return m.Regions
}

// SetRegions sets the Regions property
func (m *WritableConfigContextRequest) SetRegions(val []int32) {
	m.Regions = val
}

// GetRoles returns the Roles property
func (m WritableConfigContextRequest) GetRoles() []int32 {
	return m.Roles
}

// SetRoles sets the Roles property
func (m *WritableConfigContextRequest) SetRoles(val []int32) {
	m.Roles = val
}

// GetSiteGroups returns the SiteGroups property
func (m WritableConfigContextRequest) GetSiteGroups() []int32 {
	return m.SiteGroups
}

// SetSiteGroups sets the SiteGroups property
func (m *WritableConfigContextRequest) SetSiteGroups(val []int32) {
	m.SiteGroups = val
}

// GetSites returns the Sites property
func (m WritableConfigContextRequest) GetSites() []int32 {
	return m.Sites
}

// SetSites sets the Sites property
func (m *WritableConfigContextRequest) SetSites(val []int32) {
	m.Sites = val
}

// GetTags returns the Tags property
func (m WritableConfigContextRequest) GetTags() []string {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableConfigContextRequest) SetTags(val []string) {
	m.Tags = val
}

// GetTenantGroups returns the TenantGroups property
func (m WritableConfigContextRequest) GetTenantGroups() []int32 {
	return m.TenantGroups
}

// SetTenantGroups sets the TenantGroups property
func (m *WritableConfigContextRequest) SetTenantGroups(val []int32) {
	m.TenantGroups = val
}

// GetTenants returns the Tenants property
func (m WritableConfigContextRequest) GetTenants() []int32 {
	return m.Tenants
}

// SetTenants sets the Tenants property
func (m *WritableConfigContextRequest) SetTenants(val []int32) {
	m.Tenants = val
}

// GetWeight returns the Weight property
func (m WritableConfigContextRequest) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *WritableConfigContextRequest) SetWeight(val int32) {
	m.Weight = val
}
