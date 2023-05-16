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

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ConfigContext is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ConfigContext struct {
	// ClusterGroups:
	ClusterGroups []int32 `json:"cluster_groups,omitempty" mapstructure:"cluster_groups,omitempty"`
	// ClusterTypes:
	ClusterTypes []int32 `json:"cluster_types,omitempty" mapstructure:"cluster_types,omitempty"`
	// Clusters:
	Clusters []int32 `json:"clusters,omitempty" mapstructure:"clusters,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Data:
	Data map[string]interface{} `json:"data" mapstructure:"data"`
	// DataFile: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DataFile NestedDataFile `json:"data_file" mapstructure:"data_file"`
	// DataPath: Path to remote file (relative to data source root)
	DataPath string `json:"data_path" mapstructure:"data_path"`
	// DataSource: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DataSource *NestedDataSource `json:"data_source,omitempty" mapstructure:"data_source,omitempty"`
	// DataSynced:
	DataSynced *time.Time `json:"data_synced" mapstructure:"data_synced"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceTypes:
	DeviceTypes []int32 `json:"device_types,omitempty" mapstructure:"device_types,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// IsActive:
	IsActive bool `json:"is_active,omitempty" mapstructure:"is_active,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
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
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m ConfigContext) Validate() error {
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
		"dataFile": validation.Validate(
			m.DataFile, validation.NotNil,
		),
		"dataSource": validation.Validate(
			m.DataSource,
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
			m.Name, validation.NotNil, validation.Length(0, 100),
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetClusterGroups returns the ClusterGroups property
func (m ConfigContext) GetClusterGroups() []int32 {
	return m.ClusterGroups
}

// SetClusterGroups sets the ClusterGroups property
func (m *ConfigContext) SetClusterGroups(val []int32) {
	m.ClusterGroups = val
}

// GetClusterTypes returns the ClusterTypes property
func (m ConfigContext) GetClusterTypes() []int32 {
	return m.ClusterTypes
}

// SetClusterTypes sets the ClusterTypes property
func (m *ConfigContext) SetClusterTypes(val []int32) {
	m.ClusterTypes = val
}

// GetClusters returns the Clusters property
func (m ConfigContext) GetClusters() []int32 {
	return m.Clusters
}

// SetClusters sets the Clusters property
func (m *ConfigContext) SetClusters(val []int32) {
	m.Clusters = val
}

// GetCreated returns the Created property
func (m ConfigContext) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ConfigContext) SetCreated(val *time.Time) {
	m.Created = val
}

// GetData returns the Data property
func (m ConfigContext) GetData() map[string]interface{} {
	return m.Data
}

// SetData sets the Data property
func (m *ConfigContext) SetData(val map[string]interface{}) {
	m.Data = val
}

// GetDataFile returns the DataFile property
func (m ConfigContext) GetDataFile() NestedDataFile {
	return m.DataFile
}

// SetDataFile sets the DataFile property
func (m *ConfigContext) SetDataFile(val NestedDataFile) {
	m.DataFile = val
}

// GetDataPath returns the DataPath property
func (m ConfigContext) GetDataPath() string {
	return m.DataPath
}

// SetDataPath sets the DataPath property
func (m *ConfigContext) SetDataPath(val string) {
	m.DataPath = val
}

// GetDataSource returns the DataSource property
func (m ConfigContext) GetDataSource() *NestedDataSource {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m *ConfigContext) SetDataSource(val *NestedDataSource) {
	m.DataSource = val
}

// GetDataSynced returns the DataSynced property
func (m ConfigContext) GetDataSynced() *time.Time {
	return m.DataSynced
}

// SetDataSynced sets the DataSynced property
func (m *ConfigContext) SetDataSynced(val *time.Time) {
	m.DataSynced = val
}

// GetDescription returns the Description property
func (m ConfigContext) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ConfigContext) SetDescription(val string) {
	m.Description = val
}

// GetDeviceTypes returns the DeviceTypes property
func (m ConfigContext) GetDeviceTypes() []int32 {
	return m.DeviceTypes
}

// SetDeviceTypes sets the DeviceTypes property
func (m *ConfigContext) SetDeviceTypes(val []int32) {
	m.DeviceTypes = val
}

// GetDisplay returns the Display property
func (m ConfigContext) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ConfigContext) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ConfigContext) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ConfigContext) SetId(val int32) {
	m.Id = val
}

// GetIsActive returns the IsActive property
func (m ConfigContext) GetIsActive() bool {
	return m.IsActive
}

// SetIsActive sets the IsActive property
func (m *ConfigContext) SetIsActive(val bool) {
	m.IsActive = val
}

// GetLastUpdated returns the LastUpdated property
func (m ConfigContext) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ConfigContext) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLocations returns the Locations property
func (m ConfigContext) GetLocations() []int32 {
	return m.Locations
}

// SetLocations sets the Locations property
func (m *ConfigContext) SetLocations(val []int32) {
	m.Locations = val
}

// GetName returns the Name property
func (m ConfigContext) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ConfigContext) SetName(val string) {
	m.Name = val
}

// GetPlatforms returns the Platforms property
func (m ConfigContext) GetPlatforms() []int32 {
	return m.Platforms
}

// SetPlatforms sets the Platforms property
func (m *ConfigContext) SetPlatforms(val []int32) {
	m.Platforms = val
}

// GetRegions returns the Regions property
func (m ConfigContext) GetRegions() []int32 {
	return m.Regions
}

// SetRegions sets the Regions property
func (m *ConfigContext) SetRegions(val []int32) {
	m.Regions = val
}

// GetRoles returns the Roles property
func (m ConfigContext) GetRoles() []int32 {
	return m.Roles
}

// SetRoles sets the Roles property
func (m *ConfigContext) SetRoles(val []int32) {
	m.Roles = val
}

// GetSiteGroups returns the SiteGroups property
func (m ConfigContext) GetSiteGroups() []int32 {
	return m.SiteGroups
}

// SetSiteGroups sets the SiteGroups property
func (m *ConfigContext) SetSiteGroups(val []int32) {
	m.SiteGroups = val
}

// GetSites returns the Sites property
func (m ConfigContext) GetSites() []int32 {
	return m.Sites
}

// SetSites sets the Sites property
func (m *ConfigContext) SetSites(val []int32) {
	m.Sites = val
}

// GetTags returns the Tags property
func (m ConfigContext) GetTags() []string {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ConfigContext) SetTags(val []string) {
	m.Tags = val
}

// GetTenantGroups returns the TenantGroups property
func (m ConfigContext) GetTenantGroups() []int32 {
	return m.TenantGroups
}

// SetTenantGroups sets the TenantGroups property
func (m *ConfigContext) SetTenantGroups(val []int32) {
	m.TenantGroups = val
}

// GetTenants returns the Tenants property
func (m ConfigContext) GetTenants() []int32 {
	return m.Tenants
}

// SetTenants sets the Tenants property
func (m *ConfigContext) SetTenants(val []int32) {
	m.Tenants = val
}

// GetUrl returns the Url property
func (m ConfigContext) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ConfigContext) SetUrl(val string) {
	m.Url = val
}

// GetWeight returns the Weight property
func (m ConfigContext) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *ConfigContext) SetWeight(val int32) {
	m.Weight = val
}
