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

// DcimModuleTypesListQueryParameters is an object.
type DcimModuleTypesListQueryParameters struct {
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
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
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
func (m DcimModuleTypesListQueryParameters) Validate() error {
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
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
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

// GetConsolePorts returns the ConsolePorts property
func (m DcimModuleTypesListQueryParameters) GetConsolePorts() bool {
	return m.ConsolePorts
}

// SetConsolePorts sets the ConsolePorts property
func (m *DcimModuleTypesListQueryParameters) SetConsolePorts(val bool) {
	m.ConsolePorts = val
}

// GetConsoleServerPorts returns the ConsoleServerPorts property
func (m DcimModuleTypesListQueryParameters) GetConsoleServerPorts() bool {
	return m.ConsoleServerPorts
}

// SetConsoleServerPorts sets the ConsoleServerPorts property
func (m *DcimModuleTypesListQueryParameters) SetConsoleServerPorts(val bool) {
	m.ConsoleServerPorts = val
}

// GetCreated returns the Created property
func (m DcimModuleTypesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimModuleTypesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimModuleTypesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimModuleTypesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimModuleTypesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimModuleTypesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimModuleTypesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimModuleTypesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimModuleTypesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimModuleTypesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimModuleTypesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimModuleTypesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimModuleTypesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimModuleTypesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetId returns the Id property
func (m DcimModuleTypesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimModuleTypesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimModuleTypesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimModuleTypesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimModuleTypesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimModuleTypesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimModuleTypesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimModuleTypesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimModuleTypesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimModuleTypesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimModuleTypesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimModuleTypesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterfaces returns the Interfaces property
func (m DcimModuleTypesListQueryParameters) GetInterfaces() bool {
	return m.Interfaces
}

// SetInterfaces sets the Interfaces property
func (m *DcimModuleTypesListQueryParameters) SetInterfaces(val bool) {
	m.Interfaces = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimModuleTypesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimModuleTypesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimModuleTypesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimModuleTypesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimModuleTypesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimModuleTypesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimModuleTypesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimModuleTypesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimModuleTypesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimModuleTypesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimModuleTypesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimModuleTypesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimModuleTypesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimModuleTypesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetManufacturer returns the Manufacturer property
func (m DcimModuleTypesListQueryParameters) GetManufacturer() []string {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DcimModuleTypesListQueryParameters) SetManufacturer(val []string) {
	m.Manufacturer = val
}

// GetManufacturerN returns the ManufacturerN property
func (m DcimModuleTypesListQueryParameters) GetManufacturerN() []string {
	return m.ManufacturerN
}

// SetManufacturerN sets the ManufacturerN property
func (m *DcimModuleTypesListQueryParameters) SetManufacturerN(val []string) {
	m.ManufacturerN = val
}

// GetManufacturerId returns the ManufacturerId property
func (m DcimModuleTypesListQueryParameters) GetManufacturerId() []int32 {
	return m.ManufacturerId
}

// SetManufacturerId sets the ManufacturerId property
func (m *DcimModuleTypesListQueryParameters) SetManufacturerId(val []int32) {
	m.ManufacturerId = val
}

// GetManufacturerIdN returns the ManufacturerIdN property
func (m DcimModuleTypesListQueryParameters) GetManufacturerIdN() []int32 {
	return m.ManufacturerIdN
}

// SetManufacturerIdN sets the ManufacturerIdN property
func (m *DcimModuleTypesListQueryParameters) SetManufacturerIdN(val []int32) {
	m.ManufacturerIdN = val
}

// GetModel returns the Model property
func (m DcimModuleTypesListQueryParameters) GetModel() []string {
	return m.Model
}

// SetModel sets the Model property
func (m *DcimModuleTypesListQueryParameters) SetModel(val []string) {
	m.Model = val
}

// GetModelEmpty returns the ModelEmpty property
func (m DcimModuleTypesListQueryParameters) GetModelEmpty() []string {
	return m.ModelEmpty
}

// SetModelEmpty sets the ModelEmpty property
func (m *DcimModuleTypesListQueryParameters) SetModelEmpty(val []string) {
	m.ModelEmpty = val
}

// GetModelIc returns the ModelIc property
func (m DcimModuleTypesListQueryParameters) GetModelIc() []string {
	return m.ModelIc
}

// SetModelIc sets the ModelIc property
func (m *DcimModuleTypesListQueryParameters) SetModelIc(val []string) {
	m.ModelIc = val
}

// GetModelIe returns the ModelIe property
func (m DcimModuleTypesListQueryParameters) GetModelIe() []string {
	return m.ModelIe
}

// SetModelIe sets the ModelIe property
func (m *DcimModuleTypesListQueryParameters) SetModelIe(val []string) {
	m.ModelIe = val
}

// GetModelIew returns the ModelIew property
func (m DcimModuleTypesListQueryParameters) GetModelIew() []string {
	return m.ModelIew
}

// SetModelIew sets the ModelIew property
func (m *DcimModuleTypesListQueryParameters) SetModelIew(val []string) {
	m.ModelIew = val
}

// GetModelIsw returns the ModelIsw property
func (m DcimModuleTypesListQueryParameters) GetModelIsw() []string {
	return m.ModelIsw
}

// SetModelIsw sets the ModelIsw property
func (m *DcimModuleTypesListQueryParameters) SetModelIsw(val []string) {
	m.ModelIsw = val
}

// GetModelN returns the ModelN property
func (m DcimModuleTypesListQueryParameters) GetModelN() []string {
	return m.ModelN
}

// SetModelN sets the ModelN property
func (m *DcimModuleTypesListQueryParameters) SetModelN(val []string) {
	m.ModelN = val
}

// GetModelNic returns the ModelNic property
func (m DcimModuleTypesListQueryParameters) GetModelNic() []string {
	return m.ModelNic
}

// SetModelNic sets the ModelNic property
func (m *DcimModuleTypesListQueryParameters) SetModelNic(val []string) {
	m.ModelNic = val
}

// GetModelNie returns the ModelNie property
func (m DcimModuleTypesListQueryParameters) GetModelNie() []string {
	return m.ModelNie
}

// SetModelNie sets the ModelNie property
func (m *DcimModuleTypesListQueryParameters) SetModelNie(val []string) {
	m.ModelNie = val
}

// GetModelNiew returns the ModelNiew property
func (m DcimModuleTypesListQueryParameters) GetModelNiew() []string {
	return m.ModelNiew
}

// SetModelNiew sets the ModelNiew property
func (m *DcimModuleTypesListQueryParameters) SetModelNiew(val []string) {
	m.ModelNiew = val
}

// GetModelNisw returns the ModelNisw property
func (m DcimModuleTypesListQueryParameters) GetModelNisw() []string {
	return m.ModelNisw
}

// SetModelNisw sets the ModelNisw property
func (m *DcimModuleTypesListQueryParameters) SetModelNisw(val []string) {
	m.ModelNisw = val
}

// GetOffset returns the Offset property
func (m DcimModuleTypesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimModuleTypesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimModuleTypesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimModuleTypesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPartNumber returns the PartNumber property
func (m DcimModuleTypesListQueryParameters) GetPartNumber() []string {
	return m.PartNumber
}

// SetPartNumber sets the PartNumber property
func (m *DcimModuleTypesListQueryParameters) SetPartNumber(val []string) {
	m.PartNumber = val
}

// GetPartNumberEmpty returns the PartNumberEmpty property
func (m DcimModuleTypesListQueryParameters) GetPartNumberEmpty() []string {
	return m.PartNumberEmpty
}

// SetPartNumberEmpty sets the PartNumberEmpty property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberEmpty(val []string) {
	m.PartNumberEmpty = val
}

// GetPartNumberIc returns the PartNumberIc property
func (m DcimModuleTypesListQueryParameters) GetPartNumberIc() []string {
	return m.PartNumberIc
}

// SetPartNumberIc sets the PartNumberIc property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberIc(val []string) {
	m.PartNumberIc = val
}

// GetPartNumberIe returns the PartNumberIe property
func (m DcimModuleTypesListQueryParameters) GetPartNumberIe() []string {
	return m.PartNumberIe
}

// SetPartNumberIe sets the PartNumberIe property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberIe(val []string) {
	m.PartNumberIe = val
}

// GetPartNumberIew returns the PartNumberIew property
func (m DcimModuleTypesListQueryParameters) GetPartNumberIew() []string {
	return m.PartNumberIew
}

// SetPartNumberIew sets the PartNumberIew property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberIew(val []string) {
	m.PartNumberIew = val
}

// GetPartNumberIsw returns the PartNumberIsw property
func (m DcimModuleTypesListQueryParameters) GetPartNumberIsw() []string {
	return m.PartNumberIsw
}

// SetPartNumberIsw sets the PartNumberIsw property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberIsw(val []string) {
	m.PartNumberIsw = val
}

// GetPartNumberN returns the PartNumberN property
func (m DcimModuleTypesListQueryParameters) GetPartNumberN() []string {
	return m.PartNumberN
}

// SetPartNumberN sets the PartNumberN property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberN(val []string) {
	m.PartNumberN = val
}

// GetPartNumberNic returns the PartNumberNic property
func (m DcimModuleTypesListQueryParameters) GetPartNumberNic() []string {
	return m.PartNumberNic
}

// SetPartNumberNic sets the PartNumberNic property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberNic(val []string) {
	m.PartNumberNic = val
}

// GetPartNumberNie returns the PartNumberNie property
func (m DcimModuleTypesListQueryParameters) GetPartNumberNie() []string {
	return m.PartNumberNie
}

// SetPartNumberNie sets the PartNumberNie property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberNie(val []string) {
	m.PartNumberNie = val
}

// GetPartNumberNiew returns the PartNumberNiew property
func (m DcimModuleTypesListQueryParameters) GetPartNumberNiew() []string {
	return m.PartNumberNiew
}

// SetPartNumberNiew sets the PartNumberNiew property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberNiew(val []string) {
	m.PartNumberNiew = val
}

// GetPartNumberNisw returns the PartNumberNisw property
func (m DcimModuleTypesListQueryParameters) GetPartNumberNisw() []string {
	return m.PartNumberNisw
}

// SetPartNumberNisw sets the PartNumberNisw property
func (m *DcimModuleTypesListQueryParameters) SetPartNumberNisw(val []string) {
	m.PartNumberNisw = val
}

// GetPassThroughPorts returns the PassThroughPorts property
func (m DcimModuleTypesListQueryParameters) GetPassThroughPorts() bool {
	return m.PassThroughPorts
}

// SetPassThroughPorts sets the PassThroughPorts property
func (m *DcimModuleTypesListQueryParameters) SetPassThroughPorts(val bool) {
	m.PassThroughPorts = val
}

// GetPowerOutlets returns the PowerOutlets property
func (m DcimModuleTypesListQueryParameters) GetPowerOutlets() bool {
	return m.PowerOutlets
}

// SetPowerOutlets sets the PowerOutlets property
func (m *DcimModuleTypesListQueryParameters) SetPowerOutlets(val bool) {
	m.PowerOutlets = val
}

// GetPowerPorts returns the PowerPorts property
func (m DcimModuleTypesListQueryParameters) GetPowerPorts() bool {
	return m.PowerPorts
}

// SetPowerPorts sets the PowerPorts property
func (m *DcimModuleTypesListQueryParameters) SetPowerPorts(val bool) {
	m.PowerPorts = val
}

// GetQ returns the Q property
func (m DcimModuleTypesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimModuleTypesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m DcimModuleTypesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimModuleTypesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimModuleTypesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimModuleTypesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimModuleTypesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimModuleTypesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetWeight returns the Weight property
func (m DcimModuleTypesListQueryParameters) GetWeight() []float64 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *DcimModuleTypesListQueryParameters) SetWeight(val []float64) {
	m.Weight = val
}

// GetWeightGt returns the WeightGt property
func (m DcimModuleTypesListQueryParameters) GetWeightGt() []float64 {
	return m.WeightGt
}

// SetWeightGt sets the WeightGt property
func (m *DcimModuleTypesListQueryParameters) SetWeightGt(val []float64) {
	m.WeightGt = val
}

// GetWeightGte returns the WeightGte property
func (m DcimModuleTypesListQueryParameters) GetWeightGte() []float64 {
	return m.WeightGte
}

// SetWeightGte sets the WeightGte property
func (m *DcimModuleTypesListQueryParameters) SetWeightGte(val []float64) {
	m.WeightGte = val
}

// GetWeightLt returns the WeightLt property
func (m DcimModuleTypesListQueryParameters) GetWeightLt() []float64 {
	return m.WeightLt
}

// SetWeightLt sets the WeightLt property
func (m *DcimModuleTypesListQueryParameters) SetWeightLt(val []float64) {
	m.WeightLt = val
}

// GetWeightLte returns the WeightLte property
func (m DcimModuleTypesListQueryParameters) GetWeightLte() []float64 {
	return m.WeightLte
}

// SetWeightLte sets the WeightLte property
func (m *DcimModuleTypesListQueryParameters) SetWeightLte(val []float64) {
	m.WeightLte = val
}

// GetWeightN returns the WeightN property
func (m DcimModuleTypesListQueryParameters) GetWeightN() []float64 {
	return m.WeightN
}

// SetWeightN sets the WeightN property
func (m *DcimModuleTypesListQueryParameters) SetWeightN(val []float64) {
	m.WeightN = val
}

// GetWeightUnit returns the WeightUnit property
func (m DcimModuleTypesListQueryParameters) GetWeightUnit() string {
	return m.WeightUnit
}

// SetWeightUnit sets the WeightUnit property
func (m *DcimModuleTypesListQueryParameters) SetWeightUnit(val string) {
	m.WeightUnit = val
}

// GetWeightUnitN returns the WeightUnitN property
func (m DcimModuleTypesListQueryParameters) GetWeightUnitN() string {
	return m.WeightUnitN
}

// SetWeightUnitN sets the WeightUnitN property
func (m *DcimModuleTypesListQueryParameters) SetWeightUnitN(val string) {
	m.WeightUnitN = val
}
