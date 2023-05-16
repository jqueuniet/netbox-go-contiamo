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

// DcimInventoryItemTemplatesListQueryParameters is an object.
type DcimInventoryItemTemplatesListQueryParameters struct {
	// ComponentId:
	ComponentId []int32 `json:"component_id,omitempty" mapstructure:"component_id,omitempty"`
	// ComponentIdGt:
	ComponentIdGt []int32 `json:"component_id__gt,omitempty" mapstructure:"component_id__gt,omitempty"`
	// ComponentIdGte:
	ComponentIdGte []int32 `json:"component_id__gte,omitempty" mapstructure:"component_id__gte,omitempty"`
	// ComponentIdLt:
	ComponentIdLt []int32 `json:"component_id__lt,omitempty" mapstructure:"component_id__lt,omitempty"`
	// ComponentIdLte:
	ComponentIdLte []int32 `json:"component_id__lte,omitempty" mapstructure:"component_id__lte,omitempty"`
	// ComponentIdN:
	ComponentIdN []int32 `json:"component_id__n,omitempty" mapstructure:"component_id__n,omitempty"`
	// ComponentType:
	ComponentType string `json:"component_type,omitempty" mapstructure:"component_type,omitempty"`
	// ComponentTypeN:
	ComponentTypeN string `json:"component_type__n,omitempty" mapstructure:"component_type__n,omitempty"`
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
	// DevicetypeId: Device type (ID)
	DevicetypeId []int32 `json:"devicetype_id,omitempty" mapstructure:"devicetype_id,omitempty"`
	// DevicetypeIdN: Device type (ID)
	DevicetypeIdN []int32 `json:"devicetype_id__n,omitempty" mapstructure:"devicetype_id__n,omitempty"`
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
	// Label:
	Label []string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LabelEmpty:
	LabelEmpty []string `json:"label__empty,omitempty" mapstructure:"label__empty,omitempty"`
	// LabelIc:
	LabelIc []string `json:"label__ic,omitempty" mapstructure:"label__ic,omitempty"`
	// LabelIe:
	LabelIe []string `json:"label__ie,omitempty" mapstructure:"label__ie,omitempty"`
	// LabelIew:
	LabelIew []string `json:"label__iew,omitempty" mapstructure:"label__iew,omitempty"`
	// LabelIsw:
	LabelIsw []string `json:"label__isw,omitempty" mapstructure:"label__isw,omitempty"`
	// LabelN:
	LabelN []string `json:"label__n,omitempty" mapstructure:"label__n,omitempty"`
	// LabelNic:
	LabelNic []string `json:"label__nic,omitempty" mapstructure:"label__nic,omitempty"`
	// LabelNie:
	LabelNie []string `json:"label__nie,omitempty" mapstructure:"label__nie,omitempty"`
	// LabelNiew:
	LabelNiew []string `json:"label__niew,omitempty" mapstructure:"label__niew,omitempty"`
	// LabelNisw:
	LabelNisw []string `json:"label__nisw,omitempty" mapstructure:"label__nisw,omitempty"`
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
	ManufacturerId []*int32 `json:"manufacturer_id,omitempty" mapstructure:"manufacturer_id,omitempty"`
	// ManufacturerIdN: Manufacturer (ID)
	ManufacturerIdN []*int32 `json:"manufacturer_id__n,omitempty" mapstructure:"manufacturer_id__n,omitempty"`
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
	// ParentId: Parent inventory item (ID)
	ParentId []*int32 `json:"parent_id,omitempty" mapstructure:"parent_id,omitempty"`
	// ParentIdN: Parent inventory item (ID)
	ParentIdN []*int32 `json:"parent_id__n,omitempty" mapstructure:"parent_id__n,omitempty"`
	// PartId:
	PartId []string `json:"part_id,omitempty" mapstructure:"part_id,omitempty"`
	// PartIdEmpty:
	PartIdEmpty []string `json:"part_id__empty,omitempty" mapstructure:"part_id__empty,omitempty"`
	// PartIdIc:
	PartIdIc []string `json:"part_id__ic,omitempty" mapstructure:"part_id__ic,omitempty"`
	// PartIdIe:
	PartIdIe []string `json:"part_id__ie,omitempty" mapstructure:"part_id__ie,omitempty"`
	// PartIdIew:
	PartIdIew []string `json:"part_id__iew,omitempty" mapstructure:"part_id__iew,omitempty"`
	// PartIdIsw:
	PartIdIsw []string `json:"part_id__isw,omitempty" mapstructure:"part_id__isw,omitempty"`
	// PartIdN:
	PartIdN []string `json:"part_id__n,omitempty" mapstructure:"part_id__n,omitempty"`
	// PartIdNic:
	PartIdNic []string `json:"part_id__nic,omitempty" mapstructure:"part_id__nic,omitempty"`
	// PartIdNie:
	PartIdNie []string `json:"part_id__nie,omitempty" mapstructure:"part_id__nie,omitempty"`
	// PartIdNiew:
	PartIdNiew []string `json:"part_id__niew,omitempty" mapstructure:"part_id__niew,omitempty"`
	// PartIdNisw:
	PartIdNisw []string `json:"part_id__nisw,omitempty" mapstructure:"part_id__nisw,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Role: Role (slug)
	Role []string `json:"role,omitempty" mapstructure:"role,omitempty"`
	// RoleN: Role (slug)
	RoleN []string `json:"role__n,omitempty" mapstructure:"role__n,omitempty"`
	// RoleId: Role (ID)
	RoleId []*int32 `json:"role_id,omitempty" mapstructure:"role_id,omitempty"`
	// RoleIdN: Role (ID)
	RoleIdN []*int32 `json:"role_id__n,omitempty" mapstructure:"role_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimInventoryItemTemplatesListQueryParameters) Validate() error {
	return validation.Errors{
		"componentId": validation.Validate(
			m.ComponentId,
		),
		"componentIdGt": validation.Validate(
			m.ComponentIdGt,
		),
		"componentIdGte": validation.Validate(
			m.ComponentIdGte,
		),
		"componentIdLt": validation.Validate(
			m.ComponentIdLt,
		),
		"componentIdLte": validation.Validate(
			m.ComponentIdLte,
		),
		"componentIdN": validation.Validate(
			m.ComponentIdN,
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
		"devicetypeId": validation.Validate(
			m.DevicetypeId,
		),
		"devicetypeIdN": validation.Validate(
			m.DevicetypeIdN,
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
		"label": validation.Validate(
			m.Label,
		),
		"labelEmpty": validation.Validate(
			m.LabelEmpty,
		),
		"labelIc": validation.Validate(
			m.LabelIc,
		),
		"labelIe": validation.Validate(
			m.LabelIe,
		),
		"labelIew": validation.Validate(
			m.LabelIew,
		),
		"labelIsw": validation.Validate(
			m.LabelIsw,
		),
		"labelN": validation.Validate(
			m.LabelN,
		),
		"labelNic": validation.Validate(
			m.LabelNic,
		),
		"labelNie": validation.Validate(
			m.LabelNie,
		),
		"labelNiew": validation.Validate(
			m.LabelNiew,
		),
		"labelNisw": validation.Validate(
			m.LabelNisw,
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
		"parentId": validation.Validate(
			m.ParentId,
		),
		"parentIdN": validation.Validate(
			m.ParentIdN,
		),
		"partId": validation.Validate(
			m.PartId,
		),
		"partIdEmpty": validation.Validate(
			m.PartIdEmpty,
		),
		"partIdIc": validation.Validate(
			m.PartIdIc,
		),
		"partIdIe": validation.Validate(
			m.PartIdIe,
		),
		"partIdIew": validation.Validate(
			m.PartIdIew,
		),
		"partIdIsw": validation.Validate(
			m.PartIdIsw,
		),
		"partIdN": validation.Validate(
			m.PartIdN,
		),
		"partIdNic": validation.Validate(
			m.PartIdNic,
		),
		"partIdNie": validation.Validate(
			m.PartIdNie,
		),
		"partIdNiew": validation.Validate(
			m.PartIdNiew,
		),
		"partIdNisw": validation.Validate(
			m.PartIdNisw,
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
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetComponentId returns the ComponentId property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentId() []int32 {
	return m.ComponentId
}

// SetComponentId sets the ComponentId property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentId(val []int32) {
	m.ComponentId = val
}

// GetComponentIdGt returns the ComponentIdGt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentIdGt() []int32 {
	return m.ComponentIdGt
}

// SetComponentIdGt sets the ComponentIdGt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentIdGt(val []int32) {
	m.ComponentIdGt = val
}

// GetComponentIdGte returns the ComponentIdGte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentIdGte() []int32 {
	return m.ComponentIdGte
}

// SetComponentIdGte sets the ComponentIdGte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentIdGte(val []int32) {
	m.ComponentIdGte = val
}

// GetComponentIdLt returns the ComponentIdLt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentIdLt() []int32 {
	return m.ComponentIdLt
}

// SetComponentIdLt sets the ComponentIdLt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentIdLt(val []int32) {
	m.ComponentIdLt = val
}

// GetComponentIdLte returns the ComponentIdLte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentIdLte() []int32 {
	return m.ComponentIdLte
}

// SetComponentIdLte sets the ComponentIdLte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentIdLte(val []int32) {
	m.ComponentIdLte = val
}

// GetComponentIdN returns the ComponentIdN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentIdN() []int32 {
	return m.ComponentIdN
}

// SetComponentIdN sets the ComponentIdN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentIdN(val []int32) {
	m.ComponentIdN = val
}

// GetComponentType returns the ComponentType property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentType() string {
	return m.ComponentType
}

// SetComponentType sets the ComponentType property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentType(val string) {
	m.ComponentType = val
}

// GetComponentTypeN returns the ComponentTypeN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetComponentTypeN() string {
	return m.ComponentTypeN
}

// SetComponentTypeN sets the ComponentTypeN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetComponentTypeN(val string) {
	m.ComponentTypeN = val
}

// GetCreated returns the Created property
func (m DcimInventoryItemTemplatesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimInventoryItemTemplatesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevicetypeId returns the DevicetypeId property
func (m DcimInventoryItemTemplatesListQueryParameters) GetDevicetypeId() []int32 {
	return m.DevicetypeId
}

// SetDevicetypeId sets the DevicetypeId property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetDevicetypeId(val []int32) {
	m.DevicetypeId = val
}

// GetDevicetypeIdN returns the DevicetypeIdN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetDevicetypeIdN() []int32 {
	return m.DevicetypeIdN
}

// SetDevicetypeIdN sets the DevicetypeIdN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetDevicetypeIdN(val []int32) {
	m.DevicetypeIdN = val
}

// GetId returns the Id property
func (m DcimInventoryItemTemplatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLabel returns the Label property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabel() []string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabel(val []string) {
	m.Label = val
}

// GetLabelEmpty returns the LabelEmpty property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelEmpty() []string {
	return m.LabelEmpty
}

// SetLabelEmpty sets the LabelEmpty property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelEmpty(val []string) {
	m.LabelEmpty = val
}

// GetLabelIc returns the LabelIc property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelIc() []string {
	return m.LabelIc
}

// SetLabelIc sets the LabelIc property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelIc(val []string) {
	m.LabelIc = val
}

// GetLabelIe returns the LabelIe property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelIe() []string {
	return m.LabelIe
}

// SetLabelIe sets the LabelIe property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelIe(val []string) {
	m.LabelIe = val
}

// GetLabelIew returns the LabelIew property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelIew() []string {
	return m.LabelIew
}

// SetLabelIew sets the LabelIew property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelIew(val []string) {
	m.LabelIew = val
}

// GetLabelIsw returns the LabelIsw property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelIsw() []string {
	return m.LabelIsw
}

// SetLabelIsw sets the LabelIsw property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelIsw(val []string) {
	m.LabelIsw = val
}

// GetLabelN returns the LabelN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelN() []string {
	return m.LabelN
}

// SetLabelN sets the LabelN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelN(val []string) {
	m.LabelN = val
}

// GetLabelNic returns the LabelNic property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelNic() []string {
	return m.LabelNic
}

// SetLabelNic sets the LabelNic property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelNic(val []string) {
	m.LabelNic = val
}

// GetLabelNie returns the LabelNie property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelNie() []string {
	return m.LabelNie
}

// SetLabelNie sets the LabelNie property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelNie(val []string) {
	m.LabelNie = val
}

// GetLabelNiew returns the LabelNiew property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelNiew() []string {
	return m.LabelNiew
}

// SetLabelNiew sets the LabelNiew property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelNiew(val []string) {
	m.LabelNiew = val
}

// GetLabelNisw returns the LabelNisw property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLabelNisw() []string {
	return m.LabelNisw
}

// SetLabelNisw sets the LabelNisw property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLabelNisw(val []string) {
	m.LabelNisw = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimInventoryItemTemplatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetManufacturer returns the Manufacturer property
func (m DcimInventoryItemTemplatesListQueryParameters) GetManufacturer() []string {
	return m.Manufacturer
}

// SetManufacturer sets the Manufacturer property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetManufacturer(val []string) {
	m.Manufacturer = val
}

// GetManufacturerN returns the ManufacturerN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetManufacturerN() []string {
	return m.ManufacturerN
}

// SetManufacturerN sets the ManufacturerN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetManufacturerN(val []string) {
	m.ManufacturerN = val
}

// GetManufacturerId returns the ManufacturerId property
func (m DcimInventoryItemTemplatesListQueryParameters) GetManufacturerId() []*int32 {
	return m.ManufacturerId
}

// SetManufacturerId sets the ManufacturerId property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetManufacturerId(val []*int32) {
	m.ManufacturerId = val
}

// GetManufacturerIdN returns the ManufacturerIdN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetManufacturerIdN() []*int32 {
	return m.ManufacturerIdN
}

// SetManufacturerIdN sets the ManufacturerIdN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetManufacturerIdN(val []*int32) {
	m.ManufacturerIdN = val
}

// GetName returns the Name property
func (m DcimInventoryItemTemplatesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimInventoryItemTemplatesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimInventoryItemTemplatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimInventoryItemTemplatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetParentId returns the ParentId property
func (m DcimInventoryItemTemplatesListQueryParameters) GetParentId() []*int32 {
	return m.ParentId
}

// SetParentId sets the ParentId property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetParentId(val []*int32) {
	m.ParentId = val
}

// GetParentIdN returns the ParentIdN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetParentIdN() []*int32 {
	return m.ParentIdN
}

// SetParentIdN sets the ParentIdN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetParentIdN(val []*int32) {
	m.ParentIdN = val
}

// GetPartId returns the PartId property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartId() []string {
	return m.PartId
}

// SetPartId sets the PartId property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartId(val []string) {
	m.PartId = val
}

// GetPartIdEmpty returns the PartIdEmpty property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdEmpty() []string {
	return m.PartIdEmpty
}

// SetPartIdEmpty sets the PartIdEmpty property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdEmpty(val []string) {
	m.PartIdEmpty = val
}

// GetPartIdIc returns the PartIdIc property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdIc() []string {
	return m.PartIdIc
}

// SetPartIdIc sets the PartIdIc property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdIc(val []string) {
	m.PartIdIc = val
}

// GetPartIdIe returns the PartIdIe property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdIe() []string {
	return m.PartIdIe
}

// SetPartIdIe sets the PartIdIe property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdIe(val []string) {
	m.PartIdIe = val
}

// GetPartIdIew returns the PartIdIew property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdIew() []string {
	return m.PartIdIew
}

// SetPartIdIew sets the PartIdIew property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdIew(val []string) {
	m.PartIdIew = val
}

// GetPartIdIsw returns the PartIdIsw property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdIsw() []string {
	return m.PartIdIsw
}

// SetPartIdIsw sets the PartIdIsw property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdIsw(val []string) {
	m.PartIdIsw = val
}

// GetPartIdN returns the PartIdN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdN() []string {
	return m.PartIdN
}

// SetPartIdN sets the PartIdN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdN(val []string) {
	m.PartIdN = val
}

// GetPartIdNic returns the PartIdNic property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdNic() []string {
	return m.PartIdNic
}

// SetPartIdNic sets the PartIdNic property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdNic(val []string) {
	m.PartIdNic = val
}

// GetPartIdNie returns the PartIdNie property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdNie() []string {
	return m.PartIdNie
}

// SetPartIdNie sets the PartIdNie property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdNie(val []string) {
	m.PartIdNie = val
}

// GetPartIdNiew returns the PartIdNiew property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdNiew() []string {
	return m.PartIdNiew
}

// SetPartIdNiew sets the PartIdNiew property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdNiew(val []string) {
	m.PartIdNiew = val
}

// GetPartIdNisw returns the PartIdNisw property
func (m DcimInventoryItemTemplatesListQueryParameters) GetPartIdNisw() []string {
	return m.PartIdNisw
}

// SetPartIdNisw sets the PartIdNisw property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetPartIdNisw(val []string) {
	m.PartIdNisw = val
}

// GetQ returns the Q property
func (m DcimInventoryItemTemplatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRole returns the Role property
func (m DcimInventoryItemTemplatesListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetRoleId returns the RoleId property
func (m DcimInventoryItemTemplatesListQueryParameters) GetRoleId() []*int32 {
	return m.RoleId
}

// SetRoleId sets the RoleId property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetRoleId(val []*int32) {
	m.RoleId = val
}

// GetRoleIdN returns the RoleIdN property
func (m DcimInventoryItemTemplatesListQueryParameters) GetRoleIdN() []*int32 {
	return m.RoleIdN
}

// SetRoleIdN sets the RoleIdN property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetRoleIdN(val []*int32) {
	m.RoleIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimInventoryItemTemplatesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimInventoryItemTemplatesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
