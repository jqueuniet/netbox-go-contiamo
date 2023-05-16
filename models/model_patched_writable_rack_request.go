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

// PatchedWritableRackRequest is an object. Adds support for custom fields and tags.
type PatchedWritableRackRequest struct {
	// AssetTag: A unique tag used to identify this rack
	AssetTag *string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// DescUnits: Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units,omitempty" mapstructure:"desc_units,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// FacilityId:
	FacilityId *string `json:"facility_id,omitempty" mapstructure:"facility_id,omitempty"`
	// Location:
	Location *int32 `json:"location,omitempty" mapstructure:"location,omitempty"`
	// MaxWeight: Maximum load capacity for the rack
	MaxWeight *int32 `json:"max_weight,omitempty" mapstructure:"max_weight,omitempty"`
	// MountingDepth: Maximum depth of a mounted device, in millimeters. For four-post racks, this is the distance between the front and rear rails.
	MountingDepth *int32 `json:"mounting_depth,omitempty" mapstructure:"mounting_depth,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// OuterDepth: Outer dimension of rack (depth)
	OuterDepth *int32 `json:"outer_depth,omitempty" mapstructure:"outer_depth,omitempty"`
	// OuterUnit: * `mm` - Millimeters
	// * `in` - Inches
	OuterUnit string `json:"outer_unit,omitempty" mapstructure:"outer_unit,omitempty"`
	// OuterWidth: Outer dimension of rack (width)
	OuterWidth *int32 `json:"outer_width,omitempty" mapstructure:"outer_width,omitempty"`
	// Role: Functional role
	Role *int32 `json:"role,omitempty" mapstructure:"role,omitempty"`
	// Serial:
	Serial string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// Site:
	Site int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Status: * `reserved` - Reserved
	// * `available` - Available
	// * `planned` - Planned
	// * `active` - Active
	// * `deprecated` - Deprecated
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Tenant:
	Tenant *int32 `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// Type: * `2-post-frame` - 2-post frame
	// * `4-post-frame` - 4-post frame
	// * `4-post-cabinet` - 4-post cabinet
	// * `wall-frame` - Wall-mounted frame
	// * `wall-frame-vertical` - Wall-mounted frame (vertical)
	// * `wall-cabinet` - Wall-mounted cabinet
	// * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// UHeight: Height in rack units
	UHeight int32 `json:"u_height,omitempty" mapstructure:"u_height,omitempty"`
	// Weight:
	Weight *float64 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightUnit: * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	WeightUnit string `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
	// Width: Rail-to-rail width
	//
	// * `10` - 10 inches
	// * `19` - 19 inches
	// * `21` - 21 inches
	// * `23` - 23 inches
	Width int32 `json:"width,omitempty" mapstructure:"width,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableRackRequest) Validate() error {
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
		"maxWeight": validation.Validate(
			m.MaxWeight, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"mountingDepth": validation.Validate(
			m.MountingDepth, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"outerDepth": validation.Validate(
			m.OuterDepth, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"outerWidth": validation.Validate(
			m.OuterWidth, validation.Min(*int32(0)), validation.Max(*int32(32767)),
		),
		"serial": validation.Validate(
			m.Serial, validation.Length(0, 50),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"uHeight": validation.Validate(
			m.UHeight, validation.Min(int32(1)), validation.Max(int32(100)),
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(*float64(-1e+06)), validation.Max(*float64(1e+06)),
		),
		"width": validation.Validate(
			m.Width, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m PatchedWritableRackRequest) GetAssetTag() *string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *PatchedWritableRackRequest) SetAssetTag(val *string) {
	m.AssetTag = val
}

// GetComments returns the Comments property
func (m PatchedWritableRackRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableRackRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableRackRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableRackRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescUnits returns the DescUnits property
func (m PatchedWritableRackRequest) GetDescUnits() bool {
	return m.DescUnits
}

// SetDescUnits sets the DescUnits property
func (m *PatchedWritableRackRequest) SetDescUnits(val bool) {
	m.DescUnits = val
}

// GetDescription returns the Description property
func (m PatchedWritableRackRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableRackRequest) SetDescription(val string) {
	m.Description = val
}

// GetFacilityId returns the FacilityId property
func (m PatchedWritableRackRequest) GetFacilityId() *string {
	return m.FacilityId
}

// SetFacilityId sets the FacilityId property
func (m *PatchedWritableRackRequest) SetFacilityId(val *string) {
	m.FacilityId = val
}

// GetLocation returns the Location property
func (m PatchedWritableRackRequest) GetLocation() *int32 {
	return m.Location
}

// SetLocation sets the Location property
func (m *PatchedWritableRackRequest) SetLocation(val *int32) {
	m.Location = val
}

// GetMaxWeight returns the MaxWeight property
func (m PatchedWritableRackRequest) GetMaxWeight() *int32 {
	return m.MaxWeight
}

// SetMaxWeight sets the MaxWeight property
func (m *PatchedWritableRackRequest) SetMaxWeight(val *int32) {
	m.MaxWeight = val
}

// GetMountingDepth returns the MountingDepth property
func (m PatchedWritableRackRequest) GetMountingDepth() *int32 {
	return m.MountingDepth
}

// SetMountingDepth sets the MountingDepth property
func (m *PatchedWritableRackRequest) SetMountingDepth(val *int32) {
	m.MountingDepth = val
}

// GetName returns the Name property
func (m PatchedWritableRackRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableRackRequest) SetName(val string) {
	m.Name = val
}

// GetOuterDepth returns the OuterDepth property
func (m PatchedWritableRackRequest) GetOuterDepth() *int32 {
	return m.OuterDepth
}

// SetOuterDepth sets the OuterDepth property
func (m *PatchedWritableRackRequest) SetOuterDepth(val *int32) {
	m.OuterDepth = val
}

// GetOuterUnit returns the OuterUnit property
func (m PatchedWritableRackRequest) GetOuterUnit() string {
	return m.OuterUnit
}

// SetOuterUnit sets the OuterUnit property
func (m *PatchedWritableRackRequest) SetOuterUnit(val string) {
	m.OuterUnit = val
}

// GetOuterWidth returns the OuterWidth property
func (m PatchedWritableRackRequest) GetOuterWidth() *int32 {
	return m.OuterWidth
}

// SetOuterWidth sets the OuterWidth property
func (m *PatchedWritableRackRequest) SetOuterWidth(val *int32) {
	m.OuterWidth = val
}

// GetRole returns the Role property
func (m PatchedWritableRackRequest) GetRole() *int32 {
	return m.Role
}

// SetRole sets the Role property
func (m *PatchedWritableRackRequest) SetRole(val *int32) {
	m.Role = val
}

// GetSerial returns the Serial property
func (m PatchedWritableRackRequest) GetSerial() string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *PatchedWritableRackRequest) SetSerial(val string) {
	m.Serial = val
}

// GetSite returns the Site property
func (m PatchedWritableRackRequest) GetSite() int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *PatchedWritableRackRequest) SetSite(val int32) {
	m.Site = val
}

// GetStatus returns the Status property
func (m PatchedWritableRackRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritableRackRequest) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m PatchedWritableRackRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableRackRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTenant returns the Tenant property
func (m PatchedWritableRackRequest) GetTenant() *int32 {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *PatchedWritableRackRequest) SetTenant(val *int32) {
	m.Tenant = val
}

// GetType returns the Type property
func (m PatchedWritableRackRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritableRackRequest) SetType(val string) {
	m.Type = val
}

// GetUHeight returns the UHeight property
func (m PatchedWritableRackRequest) GetUHeight() int32 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *PatchedWritableRackRequest) SetUHeight(val int32) {
	m.UHeight = val
}

// GetWeight returns the Weight property
func (m PatchedWritableRackRequest) GetWeight() *float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *PatchedWritableRackRequest) SetWeight(val *float64) {
	m.Weight = val
}

// GetWeightUnit returns the WeightUnit property
func (m PatchedWritableRackRequest) GetWeightUnit() string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *PatchedWritableRackRequest) SetWeightUnit(val string) {
	m.WeightUnit = val
}

// GetWidth returns the Width property
func (m PatchedWritableRackRequest) GetWidth() int32 {
	return m.Width
}

// SetWidth sets the Width property
func (m *PatchedWritableRackRequest) SetWidth(val int32) {
	m.Width = val
}
