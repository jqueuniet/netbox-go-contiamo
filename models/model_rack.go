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

// Rack is an object. Adds support for custom fields and tags.
type Rack struct {
	// AssetTag: A unique tag used to identify this rack
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DescUnits: Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units,omitempty" mapstructure:"desc_units,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceCount:
	DeviceCount int32 `json:"device_count" mapstructure:"device_count"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// FacilityId:
	FacilityId *string `json:"facility_id,omitempty" mapstructure:"facility_id,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Location: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Location *NestedLocation `json:"location,omitempty" mapstructure:"location,omitempty"`
	// MaxWeight: Maximum load capacity for the rack
	MaxWeight *int32 `json:"max_weight,omitempty" mapstructure:"max_weight,omitempty"`
	// MountingDepth: Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	MountingDepth *int32 `json:"mounting_depth,omitempty" mapstructure:"mounting_depth,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// OuterDepth: Outer dimension of rack (depth)
	OuterDepth *int32 `json:"outer_depth,omitempty" mapstructure:"outer_depth,omitempty"`
	// OuterUnit:
	OuterUnit *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"outer_unit,omitempty" mapstructure:"outer_unit,omitempty"`
	// OuterWidth: Outer dimension of rack (width)
	OuterWidth *int32 `json:"outer_width,omitempty" mapstructure:"outer_width,omitempty"`
	// PowerfeedCount:
	PowerfeedCount int32 `json:"powerfeed_count" mapstructure:"powerfeed_count"`
	// Role: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Role *NestedRackRole `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site NestedSite `json:"site" mapstructure:"site"`
	// Status:
	Status *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Tenant *NestedTenant `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type:
	Type *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type,omitempty" mapstructure:"type,omitempty"`
	// UHeight: Height in rack units
	UHeight int32 `json:"u_height,omitempty" mapstructure:"u_height,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Weight:
	Weight *float64 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightUnit:
	WeightUnit *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
	// Width:
	Width *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"width,omitempty" mapstructure:"width,omitempty"`
}

// Validate implements basic validation for this model
func (m Rack) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag, validation.Length(0, 50),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"facilityId": validation.Validate(
			m.FacilityId, validation.Length(0, 50),
		),
		"location": validation.Validate(
			m.Location,
		),
		"maxWeight": validation.Validate(
			m.MaxWeight, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"mountingDepth": validation.Validate(
			m.MountingDepth, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"outerDepth": validation.Validate(
			m.OuterDepth, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"outerWidth": validation.Validate(
			m.OuterWidth, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"role": validation.Validate(
			m.Role,
		),
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
		"site": validation.Validate(
			m.Site, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"tenant": validation.Validate(
			m.Tenant,
		),
		"uHeight": validation.Validate(
			m.UHeight, validation.Min(int32(1)), validation.Max(int32(100)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(*float64(-1e+06)), validation.Max(*float64(1e+06)),
		),
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m Rack) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *Rack) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComments returns the Comments property
func (m Rack) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Rack) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Rack) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Rack) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Rack) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Rack) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescUnits returns the DescUnits property
func (m Rack) GetDescUnits() bool {
	return m.DescUnits
}

// SetDescUnits sets the DescUnits property
func (m *Rack) SetDescUnits(val bool) {
	m.DescUnits = val
}

// GetDescription returns the Description property
func (m Rack) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Rack) SetDescription(val string) {
	m.Description = val
}

// GetDeviceCount returns the DeviceCount property
func (m Rack) GetDeviceCount() int32 {
	return m.DeviceCount
}

// SetDeviceCount sets the DeviceCount property
func (m *Rack) SetDeviceCount(val int32) {
	m.DeviceCount = val
}

// GetDisplay returns the Display property
func (m Rack) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Rack) SetDisplay(val string) {
	m.Display = val
}

// GetFacilityId returns the FacilityId property
func (m Rack) GetFacilityId() *string {
	return m.FacilityId
}

// SetFacilityId sets the FacilityId property
func (m *Rack) SetFacilityId(val *string) {
	m.FacilityId = val
}

// GetId returns the Id property
func (m Rack) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Rack) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Rack) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Rack) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLocation returns the Location property
func (m Rack) GetLocation() *NestedLocation {
	return m.Location
}

// SetLocation sets the Location property
func (m *Rack) SetLocation(val *NestedLocation) {
	m.Location = val
}

// GetMaxWeight returns the MaxWeight property
func (m Rack) GetMaxWeight() *int32 {
	return m.MaxWeight
}

// SetMaxWeight sets the MaxWeight property
func (m *Rack) SetMaxWeight(val *int32) {
	m.MaxWeight = val
}

// GetMountingDepth returns the MountingDepth property
func (m Rack) GetMountingDepth() *int32 {
	return m.MountingDepth
}

// SetMountingDepth sets the MountingDepth property
func (m *Rack) SetMountingDepth(val *int32) {
	m.MountingDepth = val
}

// GetName returns the Name property
func (m Rack) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Rack) SetName(val string) {
	m.Name = val
}

// GetOuterDepth returns the OuterDepth property
func (m Rack) GetOuterDepth() *int32 {
	return m.OuterDepth
}

// SetOuterDepth sets the OuterDepth property
func (m *Rack) SetOuterDepth(val *int32) {
	m.OuterDepth = val
}

// GetOuterUnit returns the OuterUnit property
func (m Rack) GetOuterUnit() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.OuterUnit
}

// SetOuterUnit sets the OuterUnit property
func (m *Rack) SetOuterUnit(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.OuterUnit = val
}

// GetOuterWidth returns the OuterWidth property
func (m Rack) GetOuterWidth() *int32 {
	return m.OuterWidth
}

// SetOuterWidth sets the OuterWidth property
func (m *Rack) SetOuterWidth(val *int32) {
	m.OuterWidth = val
}

// GetPowerfeedCount returns the PowerfeedCount property
func (m Rack) GetPowerfeedCount() int32 {
	return m.PowerfeedCount
}

// SetPowerfeedCount sets the PowerfeedCount property
func (m *Rack) SetPowerfeedCount(val int32) {
	m.PowerfeedCount = val
}

// GetRole returns the Role property
func (m Rack) GetRole() *NestedRackRole {
	return m.Role
}

// SetRole sets the Role property
func (m *Rack) SetRole(val *NestedRackRole) {
	m.Role = val
}

// GetSerial returns the Serial property
func (m Rack) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *Rack) SetSerial(val string) {
	m.Serial = val
}

// GetSite returns the Site property
func (m Rack) GetSite() NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *Rack) SetSite(val NestedSite) {
	m.Site = val
}

// GetStatus returns the Status property
func (m Rack) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Rack) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Rack) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Rack) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m Rack) GetTenant() *NestedTenant {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *Rack) SetTenant(val *NestedTenant) {
	m.Tenant = val
}

// GetType returns the Type property
func (m Rack) GetType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *Rack) SetType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUHeight returns the UHeight property
func (m Rack) GetUHeight() int32 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *Rack) SetUHeight(val int32) {
	m.UHeight = val
}

// GetUrl returns the Url property
func (m Rack) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Rack) SetUrl(val string) {
	m.Url = val
}

// GetWeight returns the Weight property
func (m Rack) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *Rack) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m Rack) GetWeightUnit() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *Rack) SetWeightUnit(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.WeightUnit = val
}

// GetWidth returns the Width property
func (m Rack) GetWidth() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Width
}

// SetWidth sets the Width property
func (m *Rack) SetWidth(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Width = val
}
