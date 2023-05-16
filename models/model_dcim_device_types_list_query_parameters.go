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

// DcimDeviceTypesListQueryParameters is an object.
type DcimDeviceTypesListQueryParameters struct {
	// Airflow: * `front-to-rear` - Front to rear
	// * `rear-to-front` - Rear to front
	// * `left-to-right` - Left to right
	// * `right-to-left` - Right to left
	// * `side-to-rear` - Side to rear
	// * `passive` - Passive
	// * `mixed` - Mixed
	Airflow string `json:"airflow,omitempty" mapstructure:"airflow,omitempty"`
	// AirflowN: * `front-to-rear` - Front to rear
	// * `rear-to-front` - Rear to front
	// * `left-to-right` - Left to right
	// * `right-to-left` - Right to left
	// * `side-to-rear` - Side to rear
	// * `passive` - Passive
	// * `mixed` - Mixed
	AirflowN string `json:"airflow__n,omitempty" mapstructure:"airflow__n,omitempty"`
	// ConsolePorts: Has console ports
	ConsolePorts bool `json:"console_ports,omitempty" mapstructure:"console_ports,omitempty"`
	// ConsoleServerPorts: Has console server ports
	ConsoleServerPorts bool `json:"console_server_ports,omitempty" mapstructure:"console_server_ports,omitempty"`
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
	// DefaultPlatform: Default platform (slug)
	DefaultPlatform []string `json:"default_platform,omitempty" mapstructure:"default_platform,omitempty"`
	// DefaultPlatformN: Default platform (slug)
	DefaultPlatformN []string `json:"default_platform__n,omitempty" mapstructure:"default_platform__n,omitempty"`
	// DefaultPlatformId: Default platform (ID)
	DefaultPlatformId []*int32 `json:"default_platform_id,omitempty" mapstructure:"default_platform_id,omitempty"`
	// DefaultPlatformIdN: Default platform (ID)
	DefaultPlatformIdN []*int32 `json:"default_platform_id__n,omitempty" mapstructure:"default_platform_id__n,omitempty"`
	// DeviceBays: Has device bays
	DeviceBays bool `json:"device_bays,omitempty" mapstructure:"device_bays,omitempty"`
	// HasFrontImage: Has a front image
	HasFrontImage bool `json:"has_front_image,omitempty" mapstructure:"has_front_image,omitempty"`
	// HasRearImage: Has a rear image
	HasRearImage bool `json:"has_rear_image,omitempty" mapstructure:"has_rear_image,omitempty"`
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
	// Interfaces: Has interfaces
	Interfaces bool `json:"interfaces,omitempty" mapstructure:"interfaces,omitempty"`
	// InventoryItems: Has inventory items
	InventoryItems bool `json:"inventory_items,omitempty" mapstructure:"inventory_items,omitempty"`
	// IsFullDepth:
	IsFullDepth bool `json:"is_full_depth,omitempty" mapstructure:"is_full_depth,omitempty"`
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
	// Model:
	Model []string `json:"model,omitempty" mapstructure:"model,omitempty"`
	// ModelEmpty:
	ModelEmpty []string `json:"model__empty,omitempty" mapstructure:"model__empty,omitempty"`
	// ModelIc:
	ModelIc []string `json:"model__ic,omitempty" mapstructure:"model__ic,omitempty"`
	// ModelIe:
	ModelIe []string `json:"model__ie,omitempty" mapstructure:"model__ie,omitempty"`
	// ModelIew:
	ModelIew []string `json:"model__iew,omitempty" mapstructure:"model__iew,omitempty"`
	// ModelIsw:
	ModelIsw []string `json:"model__isw,omitempty" mapstructure:"model__isw,omitempty"`
	// ModelN:
	ModelN []string `json:"model__n,omitempty" mapstructure:"model__n,omitempty"`
	// ModelNic:
	ModelNic []string `json:"model__nic,omitempty" mapstructure:"model__nic,omitempty"`
	// ModelNie:
	ModelNie []string `json:"model__nie,omitempty" mapstructure:"model__nie,omitempty"`
	// ModelNiew:
	ModelNiew []string `json:"model__niew,omitempty" mapstructure:"model__niew,omitempty"`
	// ModelNisw:
	ModelNisw []string `json:"model__nisw,omitempty" mapstructure:"model__nisw,omitempty"`
	// ModuleBays: Has module bays
	ModuleBays bool `json:"module_bays,omitempty" mapstructure:"module_bays,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// PartNumber:
	PartNumber []string `json:"part_number,omitempty" mapstructure:"part_number,omitempty"`
	// PartNumberEmpty:
	PartNumberEmpty []string `json:"part_number__empty,omitempty" mapstructure:"part_number__empty,omitempty"`
	// PartNumberIc:
	PartNumberIc []string `json:"part_number__ic,omitempty" mapstructure:"part_number__ic,omitempty"`
	// PartNumberIe:
	PartNumberIe []string `json:"part_number__ie,omitempty" mapstructure:"part_number__ie,omitempty"`
	// PartNumberIew:
	PartNumberIew []string `json:"part_number__iew,omitempty" mapstructure:"part_number__iew,omitempty"`
	// PartNumberIsw:
	PartNumberIsw []string `json:"part_number__isw,omitempty" mapstructure:"part_number__isw,omitempty"`
	// PartNumberN:
	PartNumberN []string `json:"part_number__n,omitempty" mapstructure:"part_number__n,omitempty"`
	// PartNumberNic:
	PartNumberNic []string `json:"part_number__nic,omitempty" mapstructure:"part_number__nic,omitempty"`
	// PartNumberNie:
	PartNumberNie []string `json:"part_number__nie,omitempty" mapstructure:"part_number__nie,omitempty"`
	// PartNumberNiew:
	PartNumberNiew []string `json:"part_number__niew,omitempty" mapstructure:"part_number__niew,omitempty"`
	// PartNumberNisw:
	PartNumberNisw []string `json:"part_number__nisw,omitempty" mapstructure:"part_number__nisw,omitempty"`
	// PassThroughPorts: Has pass-through ports
	PassThroughPorts bool `json:"pass_through_ports,omitempty" mapstructure:"pass_through_ports,omitempty"`
	// PowerOutlets: Has power outlets
	PowerOutlets bool `json:"power_outlets,omitempty" mapstructure:"power_outlets,omitempty"`
	// PowerPorts: Has power ports
	PowerPorts bool `json:"power_ports,omitempty" mapstructure:"power_ports,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
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
	// SubdeviceRole: Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.
	SubdeviceRole string `json:"subdevice_role,omitempty" mapstructure:"subdevice_role,omitempty"`
	// SubdeviceRoleN: Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.
	SubdeviceRoleN string `json:"subdevice_role__n,omitempty" mapstructure:"subdevice_role__n,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UHeight:
	UHeight []float64 `json:"u_height,omitempty" mapstructure:"u_height,omitempty"`
	// UHeightGt:
	UHeightGt []float64 `json:"u_height__gt,omitempty" mapstructure:"u_height__gt,omitempty"`
	// UHeightGte:
	UHeightGte []float64 `json:"u_height__gte,omitempty" mapstructure:"u_height__gte,omitempty"`
	// UHeightLt:
	UHeightLt []float64 `json:"u_height__lt,omitempty" mapstructure:"u_height__lt,omitempty"`
	// UHeightLte:
	UHeightLte []float64 `json:"u_height__lte,omitempty" mapstructure:"u_height__lte,omitempty"`
	// UHeightN:
	UHeightN []float64 `json:"u_height__n,omitempty" mapstructure:"u_height__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// Weight:
	Weight []float64 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
	// WeightGt:
	WeightGt []float64 `json:"weight__gt,omitempty" mapstructure:"weight__gt,omitempty"`
	// WeightGte:
	WeightGte []float64 `json:"weight__gte,omitempty" mapstructure:"weight__gte,omitempty"`
	// WeightLt:
	WeightLt []float64 `json:"weight__lt,omitempty" mapstructure:"weight__lt,omitempty"`
	// WeightLte:
	WeightLte []float64 `json:"weight__lte,omitempty" mapstructure:"weight__lte,omitempty"`
	// WeightN:
	WeightN []float64 `json:"weight__n,omitempty" mapstructure:"weight__n,omitempty"`
	// WeightUnit: * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	WeightUnit string `json:"weight_unit,omitempty" mapstructure:"weight_unit,omitempty"`
	// WeightUnitN: * `kg` - Kilograms
	// * `g` - Grams
	// * `lb` - Pounds
	// * `oz` - Ounces
	WeightUnitN string `json:"weight_unit__n,omitempty" mapstructure:"weight_unit__n,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimDeviceTypesListQueryParameters) Validate() error {
	return validation.Errors{
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
		"defaultPlatform": validation.Validate(
			m.DefaultPlatform,
		),
		"defaultPlatformN": validation.Validate(
			m.DefaultPlatformN,
		),
		"defaultPlatformId": validation.Validate(
			m.DefaultPlatformId,
		),
		"defaultPlatformIdN": validation.Validate(
			m.DefaultPlatformIdN,
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
		"model": validation.Validate(
			m.Model,
		),
		"modelEmpty": validation.Validate(
			m.ModelEmpty,
		),
		"modelIc": validation.Validate(
			m.ModelIc,
		),
		"modelIe": validation.Validate(
			m.ModelIe,
		),
		"modelIew": validation.Validate(
			m.ModelIew,
		),
		"modelIsw": validation.Validate(
			m.ModelIsw,
		),
		"modelN": validation.Validate(
			m.ModelN,
		),
		"modelNic": validation.Validate(
			m.ModelNic,
		),
		"modelNie": validation.Validate(
			m.ModelNie,
		),
		"modelNiew": validation.Validate(
			m.ModelNiew,
		),
		"modelNisw": validation.Validate(
			m.ModelNisw,
		),
		"partNumber": validation.Validate(
			m.PartNumber,
		),
		"partNumberEmpty": validation.Validate(
			m.PartNumberEmpty,
		),
		"partNumberIc": validation.Validate(
			m.PartNumberIc,
		),
		"partNumberIe": validation.Validate(
			m.PartNumberIe,
		),
		"partNumberIew": validation.Validate(
			m.PartNumberIew,
		),
		"partNumberIsw": validation.Validate(
			m.PartNumberIsw,
		),
		"partNumberN": validation.Validate(
			m.PartNumberN,
		),
		"partNumberNic": validation.Validate(
			m.PartNumberNic,
		),
		"partNumberNie": validation.Validate(
			m.PartNumberNie,
		),
		"partNumberNiew": validation.Validate(
			m.PartNumberNiew,
		),
		"partNumberNisw": validation.Validate(
			m.PartNumberNisw,
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
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"uHeight": validation.Validate(
			m.UHeight,
		),
		"uHeightGt": validation.Validate(
			m.UHeightGt,
		),
		"uHeightGte": validation.Validate(
			m.UHeightGte,
		),
		"uHeightLt": validation.Validate(
			m.UHeightLt,
		),
		"uHeightLte": validation.Validate(
			m.UHeightLte,
		),
		"uHeightN": validation.Validate(
			m.UHeightN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"weight": validation.Validate(
			m.Weight,
		),
		"weightGt": validation.Validate(
			m.WeightGt,
		),
		"weightGte": validation.Validate(
			m.WeightGte,
		),
		"weightLt": validation.Validate(
			m.WeightLt,
		),
		"weightLte": validation.Validate(
			m.WeightLte,
		),
		"weightN": validation.Validate(
			m.WeightN,
		),
	}.Filter()
}

// GetAirflow returns the Airflow property
func (m DcimDeviceTypesListQueryParameters) GetAirflow() string {
	return m.Airflow
}

// SetAirflow sets the Airflow property
func (m *DcimDeviceTypesListQueryParameters) SetAirflow(val string) {
	m.Airflow = val
}

// GetAirflowN returns the AirflowN property
func (m DcimDeviceTypesListQueryParameters) GetAirflowN() string {
	return m.AirflowN
}

// SetAirflowN sets the AirflowN property
func (m *DcimDeviceTypesListQueryParameters) SetAirflowN(val string) {
	m.AirflowN = val
}

// GetConsolePorts returns the ConsolePorts property
func (m DcimDeviceTypesListQueryParameters) GetConsolePorts() bool {
	return m.ConsolePorts
}

// SetConsolePorts sets the ConsolePorts property
func (m *DcimDeviceTypesListQueryParameters) SetConsolePorts(val bool) {
	m.ConsolePorts = val
}

// GetConsoleServerPorts returns the ConsoleServerPorts property
func (m DcimDeviceTypesListQueryParameters) GetConsoleServerPorts() bool {
	return m.ConsoleServerPorts
}

// SetConsoleServerPorts sets the ConsoleServerPorts property
func (m *DcimDeviceTypesListQueryParameters) SetConsoleServerPorts(val bool) {
	m.ConsoleServerPorts = val
}

// GetCreated returns the Created property
func (m DcimDeviceTypesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimDeviceTypesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimDeviceTypesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimDeviceTypesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimDeviceTypesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimDeviceTypesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimDeviceTypesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimDeviceTypesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimDeviceTypesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimDeviceTypesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimDeviceTypesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimDeviceTypesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimDeviceTypesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimDeviceTypesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDefaultPlatform returns the DefaultPlatform property
func (m DcimDeviceTypesListQueryParameters) GetDefaultPlatform() []string {
	return m.DefaultPlatform
}

// SetDefaultPlatform sets the DefaultPlatform property
func (m *DcimDeviceTypesListQueryParameters) SetDefaultPlatform(val []string) {
	m.DefaultPlatform = val
}

// GetDefaultPlatformN returns the DefaultPlatformN property
func (m DcimDeviceTypesListQueryParameters) GetDefaultPlatformN() []string {
	return m.DefaultPlatformN
}

// SetDefaultPlatformN sets the DefaultPlatformN property
func (m *DcimDeviceTypesListQueryParameters) SetDefaultPlatformN(val []string) {
	m.DefaultPlatformN = val
}

// GetDefaultPlatformId returns the DefaultPlatformId property
func (m DcimDeviceTypesListQueryParameters) GetDefaultPlatformId() []*int32 {
	return m.DefaultPlatformId
}

// SetDefaultPlatformId sets the DefaultPlatformId property
func (m *DcimDeviceTypesListQueryParameters) SetDefaultPlatformId(val []*int32) {
	m.DefaultPlatformId = val
}

// GetDefaultPlatformIdN returns the DefaultPlatformIdN property
func (m DcimDeviceTypesListQueryParameters) GetDefaultPlatformIdN() []*int32 {
	return m.DefaultPlatformIdN
}

// SetDefaultPlatformIdN sets the DefaultPlatformIdN property
func (m *DcimDeviceTypesListQueryParameters) SetDefaultPlatformIdN(val []*int32) {
	m.DefaultPlatformIdN = val
}

// GetDeviceBays returns the DeviceBays property
func (m DcimDeviceTypesListQueryParameters) GetDeviceBays() bool {
	return m.DeviceBays
}

// SetDeviceBays sets the DeviceBays property
func (m *DcimDeviceTypesListQueryParameters) SetDeviceBays(val bool) {
	m.DeviceBays = val
}

// GetHasFrontImage returns the HasFrontImage property
func (m DcimDeviceTypesListQueryParameters) GetHasFrontImage() bool {
	return m.HasFrontImage
}

// SetHasFrontImage sets the HasFrontImage property
func (m *DcimDeviceTypesListQueryParameters) SetHasFrontImage(val bool) {
	m.HasFrontImage = val
}

// GetHasRearImage returns the HasRearImage property
func (m DcimDeviceTypesListQueryParameters) GetHasRearImage() bool {
	return m.HasRearImage
}

// SetHasRearImage sets the HasRearImage property
func (m *DcimDeviceTypesListQueryParameters) SetHasRearImage(val bool) {
	m.HasRearImage = val
}

// GetId returns the Id property
func (m DcimDeviceTypesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimDeviceTypesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimDeviceTypesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimDeviceTypesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimDeviceTypesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimDeviceTypesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimDeviceTypesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimDeviceTypesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimDeviceTypesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimDeviceTypesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimDeviceTypesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimDeviceTypesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterfaces returns the Interfaces property
func (m DcimDeviceTypesListQueryParameters) GetInterfaces() bool {
	return m.Interfaces
}

// SetInterfaces sets the Interfaces property
func (m *DcimDeviceTypesListQueryParameters) SetInterfaces(val bool) {
	m.Interfaces = val
}

// GetInventoryItems returns the InventoryItems property
func (m DcimDeviceTypesListQueryParameters) GetInventoryItems() bool {
	return m.InventoryItems
}

// SetInventoryItems sets the InventoryItems property
func (m *DcimDeviceTypesListQueryParameters) SetInventoryItems(val bool) {
	m.InventoryItems = val
}

// GetIsFullDepth returns the IsFullDepth property
func (m DcimDeviceTypesListQueryParameters) GetIsFullDepth() bool {
	return m.IsFullDepth
}

// SetIsFullDepth sets the IsFullDepth property
func (m *DcimDeviceTypesListQueryParameters) SetIsFullDepth(val bool) {
	m.IsFullDepth = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimDeviceTypesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimDeviceTypesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimDeviceTypesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimDeviceTypesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimDeviceTypesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimDeviceTypesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimDeviceTypesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimDeviceTypesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimDeviceTypesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimDeviceTypesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimDeviceTypesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimDeviceTypesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimDeviceTypesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimDeviceTypesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetManufacturer returns the Manufacturer property
func (m DcimDeviceTypesListQueryParameters) GetManufacturer() []string {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DcimDeviceTypesListQueryParameters) SetManufacturer(val []string) {
	m.Manufacturer = val
}

// GetManufacturerN returns the ManufacturerN property
func (m DcimDeviceTypesListQueryParameters) GetManufacturerN() []string {
	return m.ManufacturerN
}

// SetManufacturerN sets the ManufacturerN property
func (m *DcimDeviceTypesListQueryParameters) SetManufacturerN(val []string) {
	m.ManufacturerN = val
}

// GetManufacturerId returns the ManufacturerId property
func (m DcimDeviceTypesListQueryParameters) GetManufacturerId() []int32 {
	return m.ManufacturerId
}

// SetManufacturerId sets the ManufacturerId property
func (m *DcimDeviceTypesListQueryParameters) SetManufacturerId(val []int32) {
	m.ManufacturerId = val
}

// GetManufacturerIdN returns the ManufacturerIdN property
func (m DcimDeviceTypesListQueryParameters) GetManufacturerIdN() []int32 {
	return m.ManufacturerIdN
}

// SetManufacturerIdN sets the ManufacturerIdN property
func (m *DcimDeviceTypesListQueryParameters) SetManufacturerIdN(val []int32) {
	m.ManufacturerIdN = val
}

// GetModel returns the Model property
func (m DcimDeviceTypesListQueryParameters) GetModel() []string {
	return m.Model
}

// SetModel sets the Model property
func (m *DcimDeviceTypesListQueryParameters) SetModel(val []string) {
	m.Model = val
}

// GetModelEmpty returns the ModelEmpty property
func (m DcimDeviceTypesListQueryParameters) GetModelEmpty() []string {
	return m.ModelEmpty
}

// SetModelEmpty sets the ModelEmpty property
func (m *DcimDeviceTypesListQueryParameters) SetModelEmpty(val []string) {
	m.ModelEmpty = val
}

// GetModelIc returns the ModelIc property
func (m DcimDeviceTypesListQueryParameters) GetModelIc() []string {
	return m.ModelIc
}

// SetModelIc sets the ModelIc property
func (m *DcimDeviceTypesListQueryParameters) SetModelIc(val []string) {
	m.ModelIc = val
}

// GetModelIe returns the ModelIe property
func (m DcimDeviceTypesListQueryParameters) GetModelIe() []string {
	return m.ModelIe
}

// SetModelIe sets the ModelIe property
func (m *DcimDeviceTypesListQueryParameters) SetModelIe(val []string) {
	m.ModelIe = val
}

// GetModelIew returns the ModelIew property
func (m DcimDeviceTypesListQueryParameters) GetModelIew() []string {
	return m.ModelIew
}

// SetModelIew sets the ModelIew property
func (m *DcimDeviceTypesListQueryParameters) SetModelIew(val []string) {
	m.ModelIew = val
}

// GetModelIsw returns the ModelIsw property
func (m DcimDeviceTypesListQueryParameters) GetModelIsw() []string {
	return m.ModelIsw
}

// SetModelIsw sets the ModelIsw property
func (m *DcimDeviceTypesListQueryParameters) SetModelIsw(val []string) {
	m.ModelIsw = val
}

// GetModelN returns the ModelN property
func (m DcimDeviceTypesListQueryParameters) GetModelN() []string {
	return m.ModelN
}

// SetModelN sets the ModelN property
func (m *DcimDeviceTypesListQueryParameters) SetModelN(val []string) {
	m.ModelN = val
}

// GetModelNic returns the ModelNic property
func (m DcimDeviceTypesListQueryParameters) GetModelNic() []string {
	return m.ModelNic
}

// SetModelNic sets the ModelNic property
func (m *DcimDeviceTypesListQueryParameters) SetModelNic(val []string) {
	m.ModelNic = val
}

// GetModelNie returns the ModelNie property
func (m DcimDeviceTypesListQueryParameters) GetModelNie() []string {
	return m.ModelNie
}

// SetModelNie sets the ModelNie property
func (m *DcimDeviceTypesListQueryParameters) SetModelNie(val []string) {
	m.ModelNie = val
}

// GetModelNiew returns the ModelNiew property
func (m DcimDeviceTypesListQueryParameters) GetModelNiew() []string {
	return m.ModelNiew
}

// SetModelNiew sets the ModelNiew property
func (m *DcimDeviceTypesListQueryParameters) SetModelNiew(val []string) {
	m.ModelNiew = val
}

// GetModelNisw returns the ModelNisw property
func (m DcimDeviceTypesListQueryParameters) GetModelNisw() []string {
	return m.ModelNisw
}

// SetModelNisw sets the ModelNisw property
func (m *DcimDeviceTypesListQueryParameters) SetModelNisw(val []string) {
	m.ModelNisw = val
}

// GetModuleBays returns the ModuleBays property
func (m DcimDeviceTypesListQueryParameters) GetModuleBays() bool {
	return m.ModuleBays
}

// SetModuleBays sets the ModuleBays property
func (m *DcimDeviceTypesListQueryParameters) SetModuleBays(val bool) {
	m.ModuleBays = val
}

// GetOffset returns the Offset property
func (m DcimDeviceTypesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimDeviceTypesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimDeviceTypesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimDeviceTypesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPartNumber returns the PartNumber property
func (m DcimDeviceTypesListQueryParameters) GetPartNumber() []string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumber(val []string) {
	m.PartNumber = val
}

// GetPartNumberEmpty returns the PartNumberEmpty property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberEmpty() []string {
	return m.PartNumberEmpty
}

// SetPartNumberEmpty sets the PartNumberEmpty property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberEmpty(val []string) {
	m.PartNumberEmpty = val
}

// GetPartNumberIc returns the PartNumberIc property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberIc() []string {
	return m.PartNumberIc
}

// SetPartNumberIc sets the PartNumberIc property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberIc(val []string) {
	m.PartNumberIc = val
}

// GetPartNumberIe returns the PartNumberIe property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberIe() []string {
	return m.PartNumberIe
}

// SetPartNumberIe sets the PartNumberIe property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberIe(val []string) {
	m.PartNumberIe = val
}

// GetPartNumberIew returns the PartNumberIew property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberIew() []string {
	return m.PartNumberIew
}

// SetPartNumberIew sets the PartNumberIew property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberIew(val []string) {
	m.PartNumberIew = val
}

// GetPartNumberIsw returns the PartNumberIsw property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberIsw() []string {
	return m.PartNumberIsw
}

// SetPartNumberIsw sets the PartNumberIsw property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberIsw(val []string) {
	m.PartNumberIsw = val
}

// GetPartNumberN returns the PartNumberN property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberN() []string {
	return m.PartNumberN
}

// SetPartNumberN sets the PartNumberN property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberN(val []string) {
	m.PartNumberN = val
}

// GetPartNumberNic returns the PartNumberNic property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberNic() []string {
	return m.PartNumberNic
}

// SetPartNumberNic sets the PartNumberNic property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberNic(val []string) {
	m.PartNumberNic = val
}

// GetPartNumberNie returns the PartNumberNie property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberNie() []string {
	return m.PartNumberNie
}

// SetPartNumberNie sets the PartNumberNie property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberNie(val []string) {
	m.PartNumberNie = val
}

// GetPartNumberNiew returns the PartNumberNiew property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberNiew() []string {
	return m.PartNumberNiew
}

// SetPartNumberNiew sets the PartNumberNiew property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberNiew(val []string) {
	m.PartNumberNiew = val
}

// GetPartNumberNisw returns the PartNumberNisw property
func (m DcimDeviceTypesListQueryParameters) GetPartNumberNisw() []string {
	return m.PartNumberNisw
}

// SetPartNumberNisw sets the PartNumberNisw property
func (m *DcimDeviceTypesListQueryParameters) SetPartNumberNisw(val []string) {
	m.PartNumberNisw = val
}

// GetPassThroughPorts returns the PassThroughPorts property
func (m DcimDeviceTypesListQueryParameters) GetPassThroughPorts() bool {
	return m.PassThroughPorts
}

// SetPassThroughPorts sets the PassThroughPorts property
func (m *DcimDeviceTypesListQueryParameters) SetPassThroughPorts(val bool) {
	m.PassThroughPorts = val
}

// GetPowerOutlets returns the PowerOutlets property
func (m DcimDeviceTypesListQueryParameters) GetPowerOutlets() bool {
	return m.PowerOutlets
}

// SetPowerOutlets sets the PowerOutlets property
func (m *DcimDeviceTypesListQueryParameters) SetPowerOutlets(val bool) {
	m.PowerOutlets = val
}

// GetPowerPorts returns the PowerPorts property
func (m DcimDeviceTypesListQueryParameters) GetPowerPorts() bool {
	return m.PowerPorts
}

// SetPowerPorts sets the PowerPorts property
func (m *DcimDeviceTypesListQueryParameters) SetPowerPorts(val bool) {
	m.PowerPorts = val
}

// GetQ returns the Q property
func (m DcimDeviceTypesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimDeviceTypesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSlug returns the Slug property
func (m DcimDeviceTypesListQueryParameters) GetSlug() []string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *DcimDeviceTypesListQueryParameters) SetSlug(val []string) {
	m.Slug = val
}

// GetSlugEmpty returns the SlugEmpty property
func (m DcimDeviceTypesListQueryParameters) GetSlugEmpty() []string {
	return m.SlugEmpty
}

// SetSlugEmpty sets the SlugEmpty property
func (m *DcimDeviceTypesListQueryParameters) SetSlugEmpty(val []string) {
	m.SlugEmpty = val
}

// GetSlugIc returns the SlugIc property
func (m DcimDeviceTypesListQueryParameters) GetSlugIc() []string {
	return m.SlugIc
}

// SetSlugIc sets the SlugIc property
func (m *DcimDeviceTypesListQueryParameters) SetSlugIc(val []string) {
	m.SlugIc = val
}

// GetSlugIe returns the SlugIe property
func (m DcimDeviceTypesListQueryParameters) GetSlugIe() []string {
	return m.SlugIe
}

// SetSlugIe sets the SlugIe property
func (m *DcimDeviceTypesListQueryParameters) SetSlugIe(val []string) {
	m.SlugIe = val
}

// GetSlugIew returns the SlugIew property
func (m DcimDeviceTypesListQueryParameters) GetSlugIew() []string {
	return m.SlugIew
}

// SetSlugIew sets the SlugIew property
func (m *DcimDeviceTypesListQueryParameters) SetSlugIew(val []string) {
	m.SlugIew = val
}

// GetSlugIsw returns the SlugIsw property
func (m DcimDeviceTypesListQueryParameters) GetSlugIsw() []string {
	return m.SlugIsw
}

// SetSlugIsw sets the SlugIsw property
func (m *DcimDeviceTypesListQueryParameters) SetSlugIsw(val []string) {
	m.SlugIsw = val
}

// GetSlugN returns the SlugN property
func (m DcimDeviceTypesListQueryParameters) GetSlugN() []string {
	return m.SlugN
}

// SetSlugN sets the SlugN property
func (m *DcimDeviceTypesListQueryParameters) SetSlugN(val []string) {
	m.SlugN = val
}

// GetSlugNic returns the SlugNic property
func (m DcimDeviceTypesListQueryParameters) GetSlugNic() []string {
	return m.SlugNic
}

// SetSlugNic sets the SlugNic property
func (m *DcimDeviceTypesListQueryParameters) SetSlugNic(val []string) {
	m.SlugNic = val
}

// GetSlugNie returns the SlugNie property
func (m DcimDeviceTypesListQueryParameters) GetSlugNie() []string {
	return m.SlugNie
}

// SetSlugNie sets the SlugNie property
func (m *DcimDeviceTypesListQueryParameters) SetSlugNie(val []string) {
	m.SlugNie = val
}

// GetSlugNiew returns the SlugNiew property
func (m DcimDeviceTypesListQueryParameters) GetSlugNiew() []string {
	return m.SlugNiew
}

// SetSlugNiew sets the SlugNiew property
func (m *DcimDeviceTypesListQueryParameters) SetSlugNiew(val []string) {
	m.SlugNiew = val
}

// GetSlugNisw returns the SlugNisw property
func (m DcimDeviceTypesListQueryParameters) GetSlugNisw() []string {
	return m.SlugNisw
}

// SetSlugNisw sets the SlugNisw property
func (m *DcimDeviceTypesListQueryParameters) SetSlugNisw(val []string) {
	m.SlugNisw = val
}

// GetSubdeviceRole returns the SubdeviceRole property
func (m DcimDeviceTypesListQueryParameters) GetSubdeviceRole() string {
	return m.SubdeviceRole
}

// SetSubdeviceRole sets the SubdeviceRole property
func (m *DcimDeviceTypesListQueryParameters) SetSubdeviceRole(val string) {
	m.SubdeviceRole = val
}

// GetSubdeviceRoleN returns the SubdeviceRoleN property
func (m DcimDeviceTypesListQueryParameters) GetSubdeviceRoleN() string {
	return m.SubdeviceRoleN
}

// SetSubdeviceRoleN sets the SubdeviceRoleN property
func (m *DcimDeviceTypesListQueryParameters) SetSubdeviceRoleN(val string) {
	m.SubdeviceRoleN = val
}

// GetTag returns the Tag property
func (m DcimDeviceTypesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimDeviceTypesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimDeviceTypesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimDeviceTypesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUHeight returns the UHeight property
func (m DcimDeviceTypesListQueryParameters) GetUHeight() []float64 {
	return m.UHeight
}

// SetUHeight sets the UHeight property
func (m *DcimDeviceTypesListQueryParameters) SetUHeight(val []float64) {
	m.UHeight = val
}

// GetUHeightGt returns the UHeightGt property
func (m DcimDeviceTypesListQueryParameters) GetUHeightGt() []float64 {
	return m.UHeightGt
}

// SetUHeightGt sets the UHeightGt property
func (m *DcimDeviceTypesListQueryParameters) SetUHeightGt(val []float64) {
	m.UHeightGt = val
}

// GetUHeightGte returns the UHeightGte property
func (m DcimDeviceTypesListQueryParameters) GetUHeightGte() []float64 {
	return m.UHeightGte
}

// SetUHeightGte sets the UHeightGte property
func (m *DcimDeviceTypesListQueryParameters) SetUHeightGte(val []float64) {
	m.UHeightGte = val
}

// GetUHeightLt returns the UHeightLt property
func (m DcimDeviceTypesListQueryParameters) GetUHeightLt() []float64 {
	return m.UHeightLt
}

// SetUHeightLt sets the UHeightLt property
func (m *DcimDeviceTypesListQueryParameters) SetUHeightLt(val []float64) {
	m.UHeightLt = val
}

// GetUHeightLte returns the UHeightLte property
func (m DcimDeviceTypesListQueryParameters) GetUHeightLte() []float64 {
	return m.UHeightLte
}

// SetUHeightLte sets the UHeightLte property
func (m *DcimDeviceTypesListQueryParameters) SetUHeightLte(val []float64) {
	m.UHeightLte = val
}

// GetUHeightN returns the UHeightN property
func (m DcimDeviceTypesListQueryParameters) GetUHeightN() []float64 {
	return m.UHeightN
}

// SetUHeightN sets the UHeightN property
func (m *DcimDeviceTypesListQueryParameters) SetUHeightN(val []float64) {
	m.UHeightN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimDeviceTypesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimDeviceTypesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetWeight returns the Weight property
func (m DcimDeviceTypesListQueryParameters) GetWeight() []float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *DcimDeviceTypesListQueryParameters) SetWeight(val []float64) {
	m.Weight = val
}

// GetWeightGt returns the WeightGt property
func (m DcimDeviceTypesListQueryParameters) GetWeightGt() []float64 {
	return m.WeightGt
}

// SetWeightGt sets the WeightGt property
func (m *DcimDeviceTypesListQueryParameters) SetWeightGt(val []float64) {
	m.WeightGt = val
}

// GetWeightGte returns the WeightGte property
func (m DcimDeviceTypesListQueryParameters) GetWeightGte() []float64 {
	return m.WeightGte
}

// SetWeightGte sets the WeightGte property
func (m *DcimDeviceTypesListQueryParameters) SetWeightGte(val []float64) {
	m.WeightGte = val
}

// GetWeightLt returns the WeightLt property
func (m DcimDeviceTypesListQueryParameters) GetWeightLt() []float64 {
	return m.WeightLt
}

// SetWeightLt sets the WeightLt property
func (m *DcimDeviceTypesListQueryParameters) SetWeightLt(val []float64) {
	m.WeightLt = val
}

// GetWeightLte returns the WeightLte property
func (m DcimDeviceTypesListQueryParameters) GetWeightLte() []float64 {
	return m.WeightLte
}

// SetWeightLte sets the WeightLte property
func (m *DcimDeviceTypesListQueryParameters) SetWeightLte(val []float64) {
	m.WeightLte = val
}

// GetWeightN returns the WeightN property
func (m DcimDeviceTypesListQueryParameters) GetWeightN() []float64 {
	return m.WeightN
}

// SetWeightN sets the WeightN property
func (m *DcimDeviceTypesListQueryParameters) SetWeightN(val []float64) {
	m.WeightN = val
}

// GetWeightUnit returns the WeightUnit property
func (m DcimDeviceTypesListQueryParameters) GetWeightUnit() string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *DcimDeviceTypesListQueryParameters) SetWeightUnit(val string) {
	m.WeightUnit = val
}

// GetWeightUnitN returns the WeightUnitN property
func (m DcimDeviceTypesListQueryParameters) GetWeightUnitN() string {
	return m.WeightUnitN
}

// SetWeightUnitN sets the WeightUnitN property
func (m *DcimDeviceTypesListQueryParameters) SetWeightUnitN(val string) {
	m.WeightUnitN = val
}
