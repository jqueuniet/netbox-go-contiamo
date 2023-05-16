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

// DcimRearPortTemplatesListQueryParameters is an object.
type DcimRearPortTemplatesListQueryParameters struct {
	// Color:
	Color []string `json:"color,omitempty" mapstructure:"color,omitempty"`
	// ColorEmpty:
	ColorEmpty []string `json:"color__empty,omitempty" mapstructure:"color__empty,omitempty"`
	// ColorIc:
	ColorIc []string `json:"color__ic,omitempty" mapstructure:"color__ic,omitempty"`
	// ColorIe:
	ColorIe []string `json:"color__ie,omitempty" mapstructure:"color__ie,omitempty"`
	// ColorIew:
	ColorIew []string `json:"color__iew,omitempty" mapstructure:"color__iew,omitempty"`
	// ColorIsw:
	ColorIsw []string `json:"color__isw,omitempty" mapstructure:"color__isw,omitempty"`
	// ColorN:
	ColorN []string `json:"color__n,omitempty" mapstructure:"color__n,omitempty"`
	// ColorNic:
	ColorNic []string `json:"color__nic,omitempty" mapstructure:"color__nic,omitempty"`
	// ColorNie:
	ColorNie []string `json:"color__nie,omitempty" mapstructure:"color__nie,omitempty"`
	// ColorNiew:
	ColorNiew []string `json:"color__niew,omitempty" mapstructure:"color__niew,omitempty"`
	// ColorNisw:
	ColorNisw []string `json:"color__nisw,omitempty" mapstructure:"color__nisw,omitempty"`
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
	DevicetypeId []*int32 `json:"devicetype_id,omitempty" mapstructure:"devicetype_id,omitempty"`
	// DevicetypeIdN: Device type (ID)
	DevicetypeIdN []*int32 `json:"devicetype_id__n,omitempty" mapstructure:"devicetype_id__n,omitempty"`
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
	// ModuletypeId: Module type (ID)
	ModuletypeId []*int32 `json:"moduletype_id,omitempty" mapstructure:"moduletype_id,omitempty"`
	// ModuletypeIdN: Module type (ID)
	ModuletypeIdN []*int32 `json:"moduletype_id__n,omitempty" mapstructure:"moduletype_id__n,omitempty"`
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
	// Positions:
	Positions []int32 `json:"positions,omitempty" mapstructure:"positions,omitempty"`
	// PositionsGt:
	PositionsGt []int32 `json:"positions__gt,omitempty" mapstructure:"positions__gt,omitempty"`
	// PositionsGte:
	PositionsGte []int32 `json:"positions__gte,omitempty" mapstructure:"positions__gte,omitempty"`
	// PositionsLt:
	PositionsLt []int32 `json:"positions__lt,omitempty" mapstructure:"positions__lt,omitempty"`
	// PositionsLte:
	PositionsLte []int32 `json:"positions__lte,omitempty" mapstructure:"positions__lte,omitempty"`
	// PositionsN:
	PositionsN []int32 `json:"positions__n,omitempty" mapstructure:"positions__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Type: * `8p8c` - 8P8C
	// * `8p6c` - 8P6C
	// * `8p4c` - 8P4C
	// * `8p2c` - 8P2C
	// * `6p6c` - 6P6C
	// * `6p4c` - 6P4C
	// * `6p2c` - 6P2C
	// * `4p4c` - 4P4C
	// * `4p2c` - 4P2C
	// * `gg45` - GG45
	// * `tera-4p` - TERA 4P
	// * `tera-2p` - TERA 2P
	// * `tera-1p` - TERA 1P
	// * `110-punch` - 110 Punch
	// * `bnc` - BNC
	// * `f` - F Connector
	// * `n` - N Connector
	// * `mrj21` - MRJ21
	// * `fc` - FC
	// * `lc` - LC
	// * `lc-pc` - LC/PC
	// * `lc-upc` - LC/UPC
	// * `lc-apc` - LC/APC
	// * `lsh` - LSH
	// * `lsh-pc` - LSH/PC
	// * `lsh-upc` - LSH/UPC
	// * `lsh-apc` - LSH/APC
	// * `mpo` - MPO
	// * `mtrj` - MTRJ
	// * `sc` - SC
	// * `sc-pc` - SC/PC
	// * `sc-upc` - SC/UPC
	// * `sc-apc` - SC/APC
	// * `st` - ST
	// * `cs` - CS
	// * `sn` - SN
	// * `sma-905` - SMA 905
	// * `sma-906` - SMA 906
	// * `urm-p2` - URM-P2
	// * `urm-p4` - URM-P4
	// * `urm-p8` - URM-P8
	// * `splice` - Splice
	// * `other` - Other
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `8p8c` - 8P8C
	// * `8p6c` - 8P6C
	// * `8p4c` - 8P4C
	// * `8p2c` - 8P2C
	// * `6p6c` - 6P6C
	// * `6p4c` - 6P4C
	// * `6p2c` - 6P2C
	// * `4p4c` - 4P4C
	// * `4p2c` - 4P2C
	// * `gg45` - GG45
	// * `tera-4p` - TERA 4P
	// * `tera-2p` - TERA 2P
	// * `tera-1p` - TERA 1P
	// * `110-punch` - 110 Punch
	// * `bnc` - BNC
	// * `f` - F Connector
	// * `n` - N Connector
	// * `mrj21` - MRJ21
	// * `fc` - FC
	// * `lc` - LC
	// * `lc-pc` - LC/PC
	// * `lc-upc` - LC/UPC
	// * `lc-apc` - LC/APC
	// * `lsh` - LSH
	// * `lsh-pc` - LSH/PC
	// * `lsh-upc` - LSH/UPC
	// * `lsh-apc` - LSH/APC
	// * `mpo` - MPO
	// * `mtrj` - MTRJ
	// * `sc` - SC
	// * `sc-pc` - SC/PC
	// * `sc-upc` - SC/UPC
	// * `sc-apc` - SC/APC
	// * `st` - ST
	// * `cs` - CS
	// * `sn` - SN
	// * `sma-905` - SMA 905
	// * `sma-906` - SMA 906
	// * `urm-p2` - URM-P2
	// * `urm-p4` - URM-P4
	// * `urm-p8` - URM-P8
	// * `splice` - Splice
	// * `other` - Other
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimRearPortTemplatesListQueryParameters) Validate() error {
	return validation.Errors{
		"color": validation.Validate(
			m.Color,
		),
		"colorEmpty": validation.Validate(
			m.ColorEmpty,
		),
		"colorIc": validation.Validate(
			m.ColorIc,
		),
		"colorIe": validation.Validate(
			m.ColorIe,
		),
		"colorIew": validation.Validate(
			m.ColorIew,
		),
		"colorIsw": validation.Validate(
			m.ColorIsw,
		),
		"colorN": validation.Validate(
			m.ColorN,
		),
		"colorNic": validation.Validate(
			m.ColorNic,
		),
		"colorNie": validation.Validate(
			m.ColorNie,
		),
		"colorNiew": validation.Validate(
			m.ColorNiew,
		),
		"colorNisw": validation.Validate(
			m.ColorNisw,
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
		"moduletypeId": validation.Validate(
			m.ModuletypeId,
		),
		"moduletypeIdN": validation.Validate(
			m.ModuletypeIdN,
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
		"positions": validation.Validate(
			m.Positions,
		),
		"positionsGt": validation.Validate(
			m.PositionsGt,
		),
		"positionsGte": validation.Validate(
			m.PositionsGte,
		),
		"positionsLt": validation.Validate(
			m.PositionsLt,
		),
		"positionsLte": validation.Validate(
			m.PositionsLte,
		),
		"positionsN": validation.Validate(
			m.PositionsN,
		),
		"type": validation.Validate(
			m.Type,
		),
		"typeN": validation.Validate(
			m.TypeN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetColor returns the Color property
func (m DcimRearPortTemplatesListQueryParameters) GetColor() []string {
	return m.Color
}

// SetColor sets the Color property
func (m *DcimRearPortTemplatesListQueryParameters) SetColor(val []string) {
	m.Color = val
}

// GetColorEmpty returns the ColorEmpty property
func (m DcimRearPortTemplatesListQueryParameters) GetColorEmpty() []string {
	return m.ColorEmpty
}

// SetColorEmpty sets the ColorEmpty property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorEmpty(val []string) {
	m.ColorEmpty = val
}

// GetColorIc returns the ColorIc property
func (m DcimRearPortTemplatesListQueryParameters) GetColorIc() []string {
	return m.ColorIc
}

// SetColorIc sets the ColorIc property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorIc(val []string) {
	m.ColorIc = val
}

// GetColorIe returns the ColorIe property
func (m DcimRearPortTemplatesListQueryParameters) GetColorIe() []string {
	return m.ColorIe
}

// SetColorIe sets the ColorIe property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorIe(val []string) {
	m.ColorIe = val
}

// GetColorIew returns the ColorIew property
func (m DcimRearPortTemplatesListQueryParameters) GetColorIew() []string {
	return m.ColorIew
}

// SetColorIew sets the ColorIew property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorIew(val []string) {
	m.ColorIew = val
}

// GetColorIsw returns the ColorIsw property
func (m DcimRearPortTemplatesListQueryParameters) GetColorIsw() []string {
	return m.ColorIsw
}

// SetColorIsw sets the ColorIsw property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorIsw(val []string) {
	m.ColorIsw = val
}

// GetColorN returns the ColorN property
func (m DcimRearPortTemplatesListQueryParameters) GetColorN() []string {
	return m.ColorN
}

// SetColorN sets the ColorN property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorN(val []string) {
	m.ColorN = val
}

// GetColorNic returns the ColorNic property
func (m DcimRearPortTemplatesListQueryParameters) GetColorNic() []string {
	return m.ColorNic
}

// SetColorNic sets the ColorNic property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorNic(val []string) {
	m.ColorNic = val
}

// GetColorNie returns the ColorNie property
func (m DcimRearPortTemplatesListQueryParameters) GetColorNie() []string {
	return m.ColorNie
}

// SetColorNie sets the ColorNie property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorNie(val []string) {
	m.ColorNie = val
}

// GetColorNiew returns the ColorNiew property
func (m DcimRearPortTemplatesListQueryParameters) GetColorNiew() []string {
	return m.ColorNiew
}

// SetColorNiew sets the ColorNiew property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorNiew(val []string) {
	m.ColorNiew = val
}

// GetColorNisw returns the ColorNisw property
func (m DcimRearPortTemplatesListQueryParameters) GetColorNisw() []string {
	return m.ColorNisw
}

// SetColorNisw sets the ColorNisw property
func (m *DcimRearPortTemplatesListQueryParameters) SetColorNisw(val []string) {
	m.ColorNisw = val
}

// GetCreated returns the Created property
func (m DcimRearPortTemplatesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimRearPortTemplatesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimRearPortTemplatesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimRearPortTemplatesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimRearPortTemplatesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimRearPortTemplatesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimRearPortTemplatesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimRearPortTemplatesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimRearPortTemplatesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimRearPortTemplatesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimRearPortTemplatesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimRearPortTemplatesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimRearPortTemplatesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimRearPortTemplatesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevicetypeId returns the DevicetypeId property
func (m DcimRearPortTemplatesListQueryParameters) GetDevicetypeId() []*int32 {
	return m.DevicetypeId
}

// SetDevicetypeId sets the DevicetypeId property
func (m *DcimRearPortTemplatesListQueryParameters) SetDevicetypeId(val []*int32) {
	m.DevicetypeId = val
}

// GetDevicetypeIdN returns the DevicetypeIdN property
func (m DcimRearPortTemplatesListQueryParameters) GetDevicetypeIdN() []*int32 {
	return m.DevicetypeIdN
}

// SetDevicetypeIdN sets the DevicetypeIdN property
func (m *DcimRearPortTemplatesListQueryParameters) SetDevicetypeIdN(val []*int32) {
	m.DevicetypeIdN = val
}

// GetId returns the Id property
func (m DcimRearPortTemplatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimRearPortTemplatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimRearPortTemplatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimRearPortTemplatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimRearPortTemplatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimRearPortTemplatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimRearPortTemplatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimRearPortTemplatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimRearPortTemplatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimRearPortTemplatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimRearPortTemplatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimRearPortTemplatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimRearPortTemplatesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimRearPortTemplatesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimRearPortTemplatesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimRearPortTemplatesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimRearPortTemplatesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimRearPortTemplatesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimRearPortTemplatesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimRearPortTemplatesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimRearPortTemplatesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimRearPortTemplatesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimRearPortTemplatesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimRearPortTemplatesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimRearPortTemplatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimRearPortTemplatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetModuletypeId returns the ModuletypeId property
func (m DcimRearPortTemplatesListQueryParameters) GetModuletypeId() []*int32 {
	return m.ModuletypeId
}

// SetModuletypeId sets the ModuletypeId property
func (m *DcimRearPortTemplatesListQueryParameters) SetModuletypeId(val []*int32) {
	m.ModuletypeId = val
}

// GetModuletypeIdN returns the ModuletypeIdN property
func (m DcimRearPortTemplatesListQueryParameters) GetModuletypeIdN() []*int32 {
	return m.ModuletypeIdN
}

// SetModuletypeIdN sets the ModuletypeIdN property
func (m *DcimRearPortTemplatesListQueryParameters) SetModuletypeIdN(val []*int32) {
	m.ModuletypeIdN = val
}

// GetName returns the Name property
func (m DcimRearPortTemplatesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimRearPortTemplatesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimRearPortTemplatesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimRearPortTemplatesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimRearPortTemplatesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimRearPortTemplatesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimRearPortTemplatesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimRearPortTemplatesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimRearPortTemplatesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimRearPortTemplatesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimRearPortTemplatesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimRearPortTemplatesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimRearPortTemplatesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimRearPortTemplatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimRearPortTemplatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimRearPortTemplatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimRearPortTemplatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPositions returns the Positions property
func (m DcimRearPortTemplatesListQueryParameters) GetPositions() []int32 {
	return m.Positions
}

// SetPositions sets the Positions property
func (m *DcimRearPortTemplatesListQueryParameters) SetPositions(val []int32) {
	m.Positions = val
}

// GetPositionsGt returns the PositionsGt property
func (m DcimRearPortTemplatesListQueryParameters) GetPositionsGt() []int32 {
	return m.PositionsGt
}

// SetPositionsGt sets the PositionsGt property
func (m *DcimRearPortTemplatesListQueryParameters) SetPositionsGt(val []int32) {
	m.PositionsGt = val
}

// GetPositionsGte returns the PositionsGte property
func (m DcimRearPortTemplatesListQueryParameters) GetPositionsGte() []int32 {
	return m.PositionsGte
}

// SetPositionsGte sets the PositionsGte property
func (m *DcimRearPortTemplatesListQueryParameters) SetPositionsGte(val []int32) {
	m.PositionsGte = val
}

// GetPositionsLt returns the PositionsLt property
func (m DcimRearPortTemplatesListQueryParameters) GetPositionsLt() []int32 {
	return m.PositionsLt
}

// SetPositionsLt sets the PositionsLt property
func (m *DcimRearPortTemplatesListQueryParameters) SetPositionsLt(val []int32) {
	m.PositionsLt = val
}

// GetPositionsLte returns the PositionsLte property
func (m DcimRearPortTemplatesListQueryParameters) GetPositionsLte() []int32 {
	return m.PositionsLte
}

// SetPositionsLte sets the PositionsLte property
func (m *DcimRearPortTemplatesListQueryParameters) SetPositionsLte(val []int32) {
	m.PositionsLte = val
}

// GetPositionsN returns the PositionsN property
func (m DcimRearPortTemplatesListQueryParameters) GetPositionsN() []int32 {
	return m.PositionsN
}

// SetPositionsN sets the PositionsN property
func (m *DcimRearPortTemplatesListQueryParameters) SetPositionsN(val []int32) {
	m.PositionsN = val
}

// GetQ returns the Q property
func (m DcimRearPortTemplatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimRearPortTemplatesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetType returns the Type property
func (m DcimRearPortTemplatesListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimRearPortTemplatesListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimRearPortTemplatesListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimRearPortTemplatesListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimRearPortTemplatesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimRearPortTemplatesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
