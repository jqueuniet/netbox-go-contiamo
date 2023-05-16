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

// DcimModulesListQueryParameters is an object.
type DcimModulesListQueryParameters struct {
	// AssetTag:
	AssetTag []string `json:"asset_tag,omitempty" mapstructure:"asset_tag,omitempty"`
	// AssetTagEmpty:
	AssetTagEmpty []string `json:"asset_tag__empty,omitempty" mapstructure:"asset_tag__empty,omitempty"`
	// AssetTagIc:
	AssetTagIc []string `json:"asset_tag__ic,omitempty" mapstructure:"asset_tag__ic,omitempty"`
	// AssetTagIe:
	AssetTagIe []string `json:"asset_tag__ie,omitempty" mapstructure:"asset_tag__ie,omitempty"`
	// AssetTagIew:
	AssetTagIew []string `json:"asset_tag__iew,omitempty" mapstructure:"asset_tag__iew,omitempty"`
	// AssetTagIsw:
	AssetTagIsw []string `json:"asset_tag__isw,omitempty" mapstructure:"asset_tag__isw,omitempty"`
	// AssetTagN:
	AssetTagN []string `json:"asset_tag__n,omitempty" mapstructure:"asset_tag__n,omitempty"`
	// AssetTagNic:
	AssetTagNic []string `json:"asset_tag__nic,omitempty" mapstructure:"asset_tag__nic,omitempty"`
	// AssetTagNie:
	AssetTagNie []string `json:"asset_tag__nie,omitempty" mapstructure:"asset_tag__nie,omitempty"`
	// AssetTagNiew:
	AssetTagNiew []string `json:"asset_tag__niew,omitempty" mapstructure:"asset_tag__niew,omitempty"`
	// AssetTagNisw:
	AssetTagNisw []string `json:"asset_tag__nisw,omitempty" mapstructure:"asset_tag__nisw,omitempty"`
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
	// DeviceId: Device (ID)
	DeviceId []int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
	// DeviceIdN: Device (ID)
	DeviceIdN []int32 `json:"device_id__n,omitempty" mapstructure:"device_id__n,omitempty"`
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
	// Manufacturer: Manufacturer (slug)
	Manufacturer []string `json:"manufacturer,omitempty" mapstructure:"manufacturer,omitempty"`
	// ManufacturerN: Manufacturer (slug)
	ManufacturerN []string `json:"manufacturer__n,omitempty" mapstructure:"manufacturer__n,omitempty"`
	// ManufacturerId: Manufacturer (ID)
	ManufacturerId []int32 `json:"manufacturer_id,omitempty" mapstructure:"manufacturer_id,omitempty"`
	// ManufacturerIdN: Manufacturer (ID)
	ManufacturerIdN []int32 `json:"manufacturer_id__n,omitempty" mapstructure:"manufacturer_id__n,omitempty"`
	// ModuleBayId: Module Bay (ID)
	ModuleBayId []int32 `json:"module_bay_id,omitempty" mapstructure:"module_bay_id,omitempty"`
	// ModuleBayIdN: Module Bay (ID)
	ModuleBayIdN []int32 `json:"module_bay_id__n,omitempty" mapstructure:"module_bay_id__n,omitempty"`
	// ModuleType: Module type (model)
	ModuleType []string `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// ModuleTypeN: Module type (model)
	ModuleTypeN []string `json:"module_type__n,omitempty" mapstructure:"module_type__n,omitempty"`
	// ModuleTypeId: Module type (ID)
	ModuleTypeId []int32 `json:"module_type_id,omitempty" mapstructure:"module_type_id,omitempty"`
	// ModuleTypeIdN: Module type (ID)
	ModuleTypeIdN []int32 `json:"module_type_id__n,omitempty" mapstructure:"module_type_id__n,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Serial:
	Serial []string `json:"serial,omitempty" mapstructure:"serial,omitempty"`
	// SerialEmpty:
	SerialEmpty []string `json:"serial__empty,omitempty" mapstructure:"serial__empty,omitempty"`
	// SerialIc:
	SerialIc []string `json:"serial__ic,omitempty" mapstructure:"serial__ic,omitempty"`
	// SerialIe:
	SerialIe []string `json:"serial__ie,omitempty" mapstructure:"serial__ie,omitempty"`
	// SerialIew:
	SerialIew []string `json:"serial__iew,omitempty" mapstructure:"serial__iew,omitempty"`
	// SerialIsw:
	SerialIsw []string `json:"serial__isw,omitempty" mapstructure:"serial__isw,omitempty"`
	// SerialN:
	SerialN []string `json:"serial__n,omitempty" mapstructure:"serial__n,omitempty"`
	// SerialNic:
	SerialNic []string `json:"serial__nic,omitempty" mapstructure:"serial__nic,omitempty"`
	// SerialNie:
	SerialNie []string `json:"serial__nie,omitempty" mapstructure:"serial__nie,omitempty"`
	// SerialNiew:
	SerialNiew []string `json:"serial__niew,omitempty" mapstructure:"serial__niew,omitempty"`
	// SerialNisw:
	SerialNisw []string `json:"serial__nisw,omitempty" mapstructure:"serial__nisw,omitempty"`
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
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimModulesListQueryParameters) Validate() error {
	return validation.Errors{
		"assetTag": validation.Validate(
			m.AssetTag,
		),
		"assetTagEmpty": validation.Validate(
			m.AssetTagEmpty,
		),
		"assetTagIc": validation.Validate(
			m.AssetTagIc,
		),
		"assetTagIe": validation.Validate(
			m.AssetTagIe,
		),
		"assetTagIew": validation.Validate(
			m.AssetTagIew,
		),
		"assetTagIsw": validation.Validate(
			m.AssetTagIsw,
		),
		"assetTagN": validation.Validate(
			m.AssetTagN,
		),
		"assetTagNic": validation.Validate(
			m.AssetTagNic,
		),
		"assetTagNie": validation.Validate(
			m.AssetTagNie,
		),
		"assetTagNiew": validation.Validate(
			m.AssetTagNiew,
		),
		"assetTagNisw": validation.Validate(
			m.AssetTagNisw,
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
		"deviceId": validation.Validate(
			m.DeviceId,
		),
		"deviceIdN": validation.Validate(
			m.DeviceIdN,
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
		"manufacturer": validation.Validate(
			m.Manufacturer,
		),
		"manufacturerN": validation.Validate(
			m.ManufacturerN,
		),
		"manufacturerId": validation.Validate(
			m.ManufacturerId,
		),
		"manufacturerIdN": validation.Validate(
			m.ManufacturerIdN,
		),
		"moduleBayId": validation.Validate(
			m.ModuleBayId,
		),
		"moduleBayIdN": validation.Validate(
			m.ModuleBayIdN,
		),
		"moduleType": validation.Validate(
			m.ModuleType,
		),
		"moduleTypeN": validation.Validate(
			m.ModuleTypeN,
		),
		"moduleTypeId": validation.Validate(
			m.ModuleTypeId,
		),
		"moduleTypeIdN": validation.Validate(
			m.ModuleTypeIdN,
		),
		"serial": validation.Validate(
			m.Serial,
		),
		"serialEmpty": validation.Validate(
			m.SerialEmpty,
		),
		"serialIc": validation.Validate(
			m.SerialIc,
		),
		"serialIe": validation.Validate(
			m.SerialIe,
		),
		"serialIew": validation.Validate(
			m.SerialIew,
		),
		"serialIsw": validation.Validate(
			m.SerialIsw,
		),
		"serialN": validation.Validate(
			m.SerialN,
		),
		"serialNic": validation.Validate(
			m.SerialNic,
		),
		"serialNie": validation.Validate(
			m.SerialNie,
		),
		"serialNiew": validation.Validate(
			m.SerialNiew,
		),
		"serialNisw": validation.Validate(
			m.SerialNisw,
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
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetAssetTag returns the AssetTag property
func (m DcimModulesListQueryParameters) GetAssetTag() []string {
	return m.AssetTag
}

// SetAssetTag sets the AssetTag property
func (m *DcimModulesListQueryParameters) SetAssetTag(val []string) {
	m.AssetTag = val
}

// GetAssetTagEmpty returns the AssetTagEmpty property
func (m DcimModulesListQueryParameters) GetAssetTagEmpty() []string {
	return m.AssetTagEmpty
}

// SetAssetTagEmpty sets the AssetTagEmpty property
func (m *DcimModulesListQueryParameters) SetAssetTagEmpty(val []string) {
	m.AssetTagEmpty = val
}

// GetAssetTagIc returns the AssetTagIc property
func (m DcimModulesListQueryParameters) GetAssetTagIc() []string {
	return m.AssetTagIc
}

// SetAssetTagIc sets the AssetTagIc property
func (m *DcimModulesListQueryParameters) SetAssetTagIc(val []string) {
	m.AssetTagIc = val
}

// GetAssetTagIe returns the AssetTagIe property
func (m DcimModulesListQueryParameters) GetAssetTagIe() []string {
	return m.AssetTagIe
}

// SetAssetTagIe sets the AssetTagIe property
func (m *DcimModulesListQueryParameters) SetAssetTagIe(val []string) {
	m.AssetTagIe = val
}

// GetAssetTagIew returns the AssetTagIew property
func (m DcimModulesListQueryParameters) GetAssetTagIew() []string {
	return m.AssetTagIew
}

// SetAssetTagIew sets the AssetTagIew property
func (m *DcimModulesListQueryParameters) SetAssetTagIew(val []string) {
	m.AssetTagIew = val
}

// GetAssetTagIsw returns the AssetTagIsw property
func (m DcimModulesListQueryParameters) GetAssetTagIsw() []string {
	return m.AssetTagIsw
}

// SetAssetTagIsw sets the AssetTagIsw property
func (m *DcimModulesListQueryParameters) SetAssetTagIsw(val []string) {
	m.AssetTagIsw = val
}

// GetAssetTagN returns the AssetTagN property
func (m DcimModulesListQueryParameters) GetAssetTagN() []string {
	return m.AssetTagN
}

// SetAssetTagN sets the AssetTagN property
func (m *DcimModulesListQueryParameters) SetAssetTagN(val []string) {
	m.AssetTagN = val
}

// GetAssetTagNic returns the AssetTagNic property
func (m DcimModulesListQueryParameters) GetAssetTagNic() []string {
	return m.AssetTagNic
}

// SetAssetTagNic sets the AssetTagNic property
func (m *DcimModulesListQueryParameters) SetAssetTagNic(val []string) {
	m.AssetTagNic = val
}

// GetAssetTagNie returns the AssetTagNie property
func (m DcimModulesListQueryParameters) GetAssetTagNie() []string {
	return m.AssetTagNie
}

// SetAssetTagNie sets the AssetTagNie property
func (m *DcimModulesListQueryParameters) SetAssetTagNie(val []string) {
	m.AssetTagNie = val
}

// GetAssetTagNiew returns the AssetTagNiew property
func (m DcimModulesListQueryParameters) GetAssetTagNiew() []string {
	return m.AssetTagNiew
}

// SetAssetTagNiew sets the AssetTagNiew property
func (m *DcimModulesListQueryParameters) SetAssetTagNiew(val []string) {
	m.AssetTagNiew = val
}

// GetAssetTagNisw returns the AssetTagNisw property
func (m DcimModulesListQueryParameters) GetAssetTagNisw() []string {
	return m.AssetTagNisw
}

// SetAssetTagNisw sets the AssetTagNisw property
func (m *DcimModulesListQueryParameters) SetAssetTagNisw(val []string) {
	m.AssetTagNisw = val
}

// GetCreated returns the Created property
func (m DcimModulesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimModulesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimModulesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimModulesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimModulesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimModulesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimModulesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimModulesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimModulesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimModulesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimModulesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimModulesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimModulesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimModulesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDeviceId returns the DeviceId property
func (m DcimModulesListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *DcimModulesListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetDeviceIdN returns the DeviceIdN property
func (m DcimModulesListQueryParameters) GetDeviceIdN() []int32 {
	return m.DeviceIdN
}

// SetDeviceIdN sets the DeviceIdN property
func (m *DcimModulesListQueryParameters) SetDeviceIdN(val []int32) {
	m.DeviceIdN = val
}

// GetId returns the Id property
func (m DcimModulesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimModulesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimModulesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimModulesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimModulesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimModulesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimModulesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimModulesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimModulesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimModulesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimModulesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimModulesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimModulesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimModulesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimModulesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimModulesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimModulesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimModulesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimModulesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimModulesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimModulesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimModulesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimModulesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimModulesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimModulesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimModulesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetManufacturer returns the Manufacturer property
func (m DcimModulesListQueryParameters) GetManufacturer() []string {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DcimModulesListQueryParameters) SetManufacturer(val []string) {
	m.Manufacturer = val
}

// GetManufacturerN returns the ManufacturerN property
func (m DcimModulesListQueryParameters) GetManufacturerN() []string {
	return m.ManufacturerN
}

// SetManufacturerN sets the ManufacturerN property
func (m *DcimModulesListQueryParameters) SetManufacturerN(val []string) {
	m.ManufacturerN = val
}

// GetManufacturerId returns the ManufacturerId property
func (m DcimModulesListQueryParameters) GetManufacturerId() []int32 {
	return m.ManufacturerId
}

// SetManufacturerId sets the ManufacturerId property
func (m *DcimModulesListQueryParameters) SetManufacturerId(val []int32) {
	m.ManufacturerId = val
}

// GetManufacturerIdN returns the ManufacturerIdN property
func (m DcimModulesListQueryParameters) GetManufacturerIdN() []int32 {
	return m.ManufacturerIdN
}

// SetManufacturerIdN sets the ManufacturerIdN property
func (m *DcimModulesListQueryParameters) SetManufacturerIdN(val []int32) {
	m.ManufacturerIdN = val
}

// GetModuleBayId returns the ModuleBayId property
func (m DcimModulesListQueryParameters) GetModuleBayId() []int32 {
	return m.ModuleBayId
}

// SetModuleBayId sets the ModuleBayId property
func (m *DcimModulesListQueryParameters) SetModuleBayId(val []int32) {
	m.ModuleBayId = val
}

// GetModuleBayIdN returns the ModuleBayIdN property
func (m DcimModulesListQueryParameters) GetModuleBayIdN() []int32 {
	return m.ModuleBayIdN
}

// SetModuleBayIdN sets the ModuleBayIdN property
func (m *DcimModulesListQueryParameters) SetModuleBayIdN(val []int32) {
	m.ModuleBayIdN = val
}

// GetModuleType returns the ModuleType property
func (m DcimModulesListQueryParameters) GetModuleType() []string {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *DcimModulesListQueryParameters) SetModuleType(val []string) {
	m.ModuleType = val
}

// GetModuleTypeN returns the ModuleTypeN property
func (m DcimModulesListQueryParameters) GetModuleTypeN() []string {
	return m.ModuleTypeN
}

// SetModuleTypeN sets the ModuleTypeN property
func (m *DcimModulesListQueryParameters) SetModuleTypeN(val []string) {
	m.ModuleTypeN = val
}

// GetModuleTypeId returns the ModuleTypeId property
func (m DcimModulesListQueryParameters) GetModuleTypeId() []int32 {
	return m.ModuleTypeId
}

// SetModuleTypeId sets the ModuleTypeId property
func (m *DcimModulesListQueryParameters) SetModuleTypeId(val []int32) {
	m.ModuleTypeId = val
}

// GetModuleTypeIdN returns the ModuleTypeIdN property
func (m DcimModulesListQueryParameters) GetModuleTypeIdN() []int32 {
	return m.ModuleTypeIdN
}

// SetModuleTypeIdN sets the ModuleTypeIdN property
func (m *DcimModulesListQueryParameters) SetModuleTypeIdN(val []int32) {
	m.ModuleTypeIdN = val
}

// GetOffset returns the Offset property
func (m DcimModulesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimModulesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimModulesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimModulesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimModulesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimModulesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSerial returns the Serial property
func (m DcimModulesListQueryParameters) GetSerial() []string {
	return m.Serial
}

// SetSerial sets the Serial property
func (m *DcimModulesListQueryParameters) SetSerial(val []string) {
	m.Serial = val
}

// GetSerialEmpty returns the SerialEmpty property
func (m DcimModulesListQueryParameters) GetSerialEmpty() []string {
	return m.SerialEmpty
}

// SetSerialEmpty sets the SerialEmpty property
func (m *DcimModulesListQueryParameters) SetSerialEmpty(val []string) {
	m.SerialEmpty = val
}

// GetSerialIc returns the SerialIc property
func (m DcimModulesListQueryParameters) GetSerialIc() []string {
	return m.SerialIc
}

// SetSerialIc sets the SerialIc property
func (m *DcimModulesListQueryParameters) SetSerialIc(val []string) {
	m.SerialIc = val
}

// GetSerialIe returns the SerialIe property
func (m DcimModulesListQueryParameters) GetSerialIe() []string {
	return m.SerialIe
}

// SetSerialIe sets the SerialIe property
func (m *DcimModulesListQueryParameters) SetSerialIe(val []string) {
	m.SerialIe = val
}

// GetSerialIew returns the SerialIew property
func (m DcimModulesListQueryParameters) GetSerialIew() []string {
	return m.SerialIew
}

// SetSerialIew sets the SerialIew property
func (m *DcimModulesListQueryParameters) SetSerialIew(val []string) {
	m.SerialIew = val
}

// GetSerialIsw returns the SerialIsw property
func (m DcimModulesListQueryParameters) GetSerialIsw() []string {
	return m.SerialIsw
}

// SetSerialIsw sets the SerialIsw property
func (m *DcimModulesListQueryParameters) SetSerialIsw(val []string) {
	m.SerialIsw = val
}

// GetSerialN returns the SerialN property
func (m DcimModulesListQueryParameters) GetSerialN() []string {
	return m.SerialN
}

// SetSerialN sets the SerialN property
func (m *DcimModulesListQueryParameters) SetSerialN(val []string) {
	m.SerialN = val
}

// GetSerialNic returns the SerialNic property
func (m DcimModulesListQueryParameters) GetSerialNic() []string {
	return m.SerialNic
}

// SetSerialNic sets the SerialNic property
func (m *DcimModulesListQueryParameters) SetSerialNic(val []string) {
	m.SerialNic = val
}

// GetSerialNie returns the SerialNie property
func (m DcimModulesListQueryParameters) GetSerialNie() []string {
	return m.SerialNie
}

// SetSerialNie sets the SerialNie property
func (m *DcimModulesListQueryParameters) SetSerialNie(val []string) {
	m.SerialNie = val
}

// GetSerialNiew returns the SerialNiew property
func (m DcimModulesListQueryParameters) GetSerialNiew() []string {
	return m.SerialNiew
}

// SetSerialNiew sets the SerialNiew property
func (m *DcimModulesListQueryParameters) SetSerialNiew(val []string) {
	m.SerialNiew = val
}

// GetSerialNisw returns the SerialNisw property
func (m DcimModulesListQueryParameters) GetSerialNisw() []string {
	return m.SerialNisw
}

// SetSerialNisw sets the SerialNisw property
func (m *DcimModulesListQueryParameters) SetSerialNisw(val []string) {
	m.SerialNisw = val
}

// GetStatus returns the Status property
func (m DcimModulesListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *DcimModulesListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m DcimModulesListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *DcimModulesListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m DcimModulesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimModulesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimModulesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimModulesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimModulesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimModulesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
