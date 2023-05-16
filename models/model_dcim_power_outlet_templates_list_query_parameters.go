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

// DcimPowerOutletTemplatesListQueryParameters is an object.
type DcimPowerOutletTemplatesListQueryParameters struct {
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
	// FeedLeg: Phase (for three-phase feeds)
	FeedLeg []string `json:"feed_leg,omitempty" mapstructure:"feed_leg,omitempty"`
	// FeedLegN: Phase (for three-phase feeds)
	FeedLegN []string `json:"feed_leg__n,omitempty" mapstructure:"feed_leg__n,omitempty"`
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
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Type: * `iec-60320-c5` - C5
	// * `iec-60320-c7` - C7
	// * `iec-60320-c13` - C13
	// * `iec-60320-c15` - C15
	// * `iec-60320-c19` - C19
	// * `iec-60320-c21` - C21
	// * `iec-60309-p-n-e-4h` - P+N+E 4H
	// * `iec-60309-p-n-e-6h` - P+N+E 6H
	// * `iec-60309-p-n-e-9h` - P+N+E 9H
	// * `iec-60309-2p-e-4h` - 2P+E 4H
	// * `iec-60309-2p-e-6h` - 2P+E 6H
	// * `iec-60309-2p-e-9h` - 2P+E 9H
	// * `iec-60309-3p-e-4h` - 3P+E 4H
	// * `iec-60309-3p-e-6h` - 3P+E 6H
	// * `iec-60309-3p-e-9h` - 3P+E 9H
	// * `iec-60309-3p-n-e-4h` - 3P+N+E 4H
	// * `iec-60309-3p-n-e-6h` - 3P+N+E 6H
	// * `iec-60309-3p-n-e-9h` - 3P+N+E 9H
	// * `nema-1-15r` - NEMA 1-15R
	// * `nema-5-15r` - NEMA 5-15R
	// * `nema-5-20r` - NEMA 5-20R
	// * `nema-5-30r` - NEMA 5-30R
	// * `nema-5-50r` - NEMA 5-50R
	// * `nema-6-15r` - NEMA 6-15R
	// * `nema-6-20r` - NEMA 6-20R
	// * `nema-6-30r` - NEMA 6-30R
	// * `nema-6-50r` - NEMA 6-50R
	// * `nema-10-30r` - NEMA 10-30R
	// * `nema-10-50r` - NEMA 10-50R
	// * `nema-14-20r` - NEMA 14-20R
	// * `nema-14-30r` - NEMA 14-30R
	// * `nema-14-50r` - NEMA 14-50R
	// * `nema-14-60r` - NEMA 14-60R
	// * `nema-15-15r` - NEMA 15-15R
	// * `nema-15-20r` - NEMA 15-20R
	// * `nema-15-30r` - NEMA 15-30R
	// * `nema-15-50r` - NEMA 15-50R
	// * `nema-15-60r` - NEMA 15-60R
	// * `nema-l1-15r` - NEMA L1-15R
	// * `nema-l5-15r` - NEMA L5-15R
	// * `nema-l5-20r` - NEMA L5-20R
	// * `nema-l5-30r` - NEMA L5-30R
	// * `nema-l5-50r` - NEMA L5-50R
	// * `nema-l6-15r` - NEMA L6-15R
	// * `nema-l6-20r` - NEMA L6-20R
	// * `nema-l6-30r` - NEMA L6-30R
	// * `nema-l6-50r` - NEMA L6-50R
	// * `nema-l10-30r` - NEMA L10-30R
	// * `nema-l14-20r` - NEMA L14-20R
	// * `nema-l14-30r` - NEMA L14-30R
	// * `nema-l14-50r` - NEMA L14-50R
	// * `nema-l14-60r` - NEMA L14-60R
	// * `nema-l15-20r` - NEMA L15-20R
	// * `nema-l15-30r` - NEMA L15-30R
	// * `nema-l15-50r` - NEMA L15-50R
	// * `nema-l15-60r` - NEMA L15-60R
	// * `nema-l21-20r` - NEMA L21-20R
	// * `nema-l21-30r` - NEMA L21-30R
	// * `nema-l22-30r` - NEMA L22-30R
	// * `CS6360C` - CS6360C
	// * `CS6364C` - CS6364C
	// * `CS8164C` - CS8164C
	// * `CS8264C` - CS8264C
	// * `CS8364C` - CS8364C
	// * `CS8464C` - CS8464C
	// * `ita-e` - ITA Type E (CEE 7/5)
	// * `ita-f` - ITA Type F (CEE 7/3)
	// * `ita-g` - ITA Type G (BS 1363)
	// * `ita-h` - ITA Type H
	// * `ita-i` - ITA Type I
	// * `ita-j` - ITA Type J
	// * `ita-k` - ITA Type K
	// * `ita-l` - ITA Type L (CEI 23-50)
	// * `ita-m` - ITA Type M (BS 546)
	// * `ita-n` - ITA Type N
	// * `ita-o` - ITA Type O
	// * `ita-multistandard` - ITA Multistandard
	// * `usb-a` - USB Type A
	// * `usb-micro-b` - USB Micro B
	// * `usb-c` - USB Type C
	// * `dc-terminal` - DC Terminal
	// * `hdot-cx` - HDOT Cx
	// * `saf-d-grid` - Saf-D-Grid
	// * `neutrik-powercon-20a` - Neutrik powerCON (20A)
	// * `neutrik-powercon-32a` - Neutrik powerCON (32A)
	// * `neutrik-powercon-true1` - Neutrik powerCON TRUE1
	// * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP
	// * `ubiquiti-smartpower` - Ubiquiti SmartPower
	// * `hardwired` - Hardwired
	// * `other` - Other
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `iec-60320-c5` - C5
	// * `iec-60320-c7` - C7
	// * `iec-60320-c13` - C13
	// * `iec-60320-c15` - C15
	// * `iec-60320-c19` - C19
	// * `iec-60320-c21` - C21
	// * `iec-60309-p-n-e-4h` - P+N+E 4H
	// * `iec-60309-p-n-e-6h` - P+N+E 6H
	// * `iec-60309-p-n-e-9h` - P+N+E 9H
	// * `iec-60309-2p-e-4h` - 2P+E 4H
	// * `iec-60309-2p-e-6h` - 2P+E 6H
	// * `iec-60309-2p-e-9h` - 2P+E 9H
	// * `iec-60309-3p-e-4h` - 3P+E 4H
	// * `iec-60309-3p-e-6h` - 3P+E 6H
	// * `iec-60309-3p-e-9h` - 3P+E 9H
	// * `iec-60309-3p-n-e-4h` - 3P+N+E 4H
	// * `iec-60309-3p-n-e-6h` - 3P+N+E 6H
	// * `iec-60309-3p-n-e-9h` - 3P+N+E 9H
	// * `nema-1-15r` - NEMA 1-15R
	// * `nema-5-15r` - NEMA 5-15R
	// * `nema-5-20r` - NEMA 5-20R
	// * `nema-5-30r` - NEMA 5-30R
	// * `nema-5-50r` - NEMA 5-50R
	// * `nema-6-15r` - NEMA 6-15R
	// * `nema-6-20r` - NEMA 6-20R
	// * `nema-6-30r` - NEMA 6-30R
	// * `nema-6-50r` - NEMA 6-50R
	// * `nema-10-30r` - NEMA 10-30R
	// * `nema-10-50r` - NEMA 10-50R
	// * `nema-14-20r` - NEMA 14-20R
	// * `nema-14-30r` - NEMA 14-30R
	// * `nema-14-50r` - NEMA 14-50R
	// * `nema-14-60r` - NEMA 14-60R
	// * `nema-15-15r` - NEMA 15-15R
	// * `nema-15-20r` - NEMA 15-20R
	// * `nema-15-30r` - NEMA 15-30R
	// * `nema-15-50r` - NEMA 15-50R
	// * `nema-15-60r` - NEMA 15-60R
	// * `nema-l1-15r` - NEMA L1-15R
	// * `nema-l5-15r` - NEMA L5-15R
	// * `nema-l5-20r` - NEMA L5-20R
	// * `nema-l5-30r` - NEMA L5-30R
	// * `nema-l5-50r` - NEMA L5-50R
	// * `nema-l6-15r` - NEMA L6-15R
	// * `nema-l6-20r` - NEMA L6-20R
	// * `nema-l6-30r` - NEMA L6-30R
	// * `nema-l6-50r` - NEMA L6-50R
	// * `nema-l10-30r` - NEMA L10-30R
	// * `nema-l14-20r` - NEMA L14-20R
	// * `nema-l14-30r` - NEMA L14-30R
	// * `nema-l14-50r` - NEMA L14-50R
	// * `nema-l14-60r` - NEMA L14-60R
	// * `nema-l15-20r` - NEMA L15-20R
	// * `nema-l15-30r` - NEMA L15-30R
	// * `nema-l15-50r` - NEMA L15-50R
	// * `nema-l15-60r` - NEMA L15-60R
	// * `nema-l21-20r` - NEMA L21-20R
	// * `nema-l21-30r` - NEMA L21-30R
	// * `nema-l22-30r` - NEMA L22-30R
	// * `CS6360C` - CS6360C
	// * `CS6364C` - CS6364C
	// * `CS8164C` - CS8164C
	// * `CS8264C` - CS8264C
	// * `CS8364C` - CS8364C
	// * `CS8464C` - CS8464C
	// * `ita-e` - ITA Type E (CEE 7/5)
	// * `ita-f` - ITA Type F (CEE 7/3)
	// * `ita-g` - ITA Type G (BS 1363)
	// * `ita-h` - ITA Type H
	// * `ita-i` - ITA Type I
	// * `ita-j` - ITA Type J
	// * `ita-k` - ITA Type K
	// * `ita-l` - ITA Type L (CEI 23-50)
	// * `ita-m` - ITA Type M (BS 546)
	// * `ita-n` - ITA Type N
	// * `ita-o` - ITA Type O
	// * `ita-multistandard` - ITA Multistandard
	// * `usb-a` - USB Type A
	// * `usb-micro-b` - USB Micro B
	// * `usb-c` - USB Type C
	// * `dc-terminal` - DC Terminal
	// * `hdot-cx` - HDOT Cx
	// * `saf-d-grid` - Saf-D-Grid
	// * `neutrik-powercon-20a` - Neutrik powerCON (20A)
	// * `neutrik-powercon-32a` - Neutrik powerCON (32A)
	// * `neutrik-powercon-true1` - Neutrik powerCON TRUE1
	// * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP
	// * `ubiquiti-smartpower` - Ubiquiti SmartPower
	// * `hardwired` - Hardwired
	// * `other` - Other
	TypeN string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimPowerOutletTemplatesListQueryParameters) Validate() error {
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
		"devicetypeId": validation.Validate(
			m.DevicetypeId,
		),
		"devicetypeIdN": validation.Validate(
			m.DevicetypeIdN,
		),
		"feedLeg": validation.Validate(
			m.FeedLeg,
		),
		"feedLegN": validation.Validate(
			m.FeedLegN,
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
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m DcimPowerOutletTemplatesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimPowerOutletTemplatesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimPowerOutletTemplatesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimPowerOutletTemplatesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimPowerOutletTemplatesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimPowerOutletTemplatesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevicetypeId returns the DevicetypeId property
func (m DcimPowerOutletTemplatesListQueryParameters) GetDevicetypeId() []*int32 {
	return m.DevicetypeId
}

// SetDevicetypeId sets the DevicetypeId property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetDevicetypeId(val []*int32) {
	m.DevicetypeId = val
}

// GetDevicetypeIdN returns the DevicetypeIdN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetDevicetypeIdN() []*int32 {
	return m.DevicetypeIdN
}

// SetDevicetypeIdN sets the DevicetypeIdN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetDevicetypeIdN(val []*int32) {
	m.DevicetypeIdN = val
}

// GetFeedLeg returns the FeedLeg property
func (m DcimPowerOutletTemplatesListQueryParameters) GetFeedLeg() []string {
	return m.FeedLeg
}

// SetFeedLeg sets the FeedLeg property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetFeedLeg(val []string) {
	m.FeedLeg = val
}

// GetFeedLegN returns the FeedLegN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetFeedLegN() []string {
	return m.FeedLegN
}

// SetFeedLegN sets the FeedLegN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetFeedLegN(val []string) {
	m.FeedLegN = val
}

// GetId returns the Id property
func (m DcimPowerOutletTemplatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimPowerOutletTemplatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimPowerOutletTemplatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimPowerOutletTemplatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimPowerOutletTemplatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimPowerOutletTemplatesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimPowerOutletTemplatesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimPowerOutletTemplatesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimPowerOutletTemplatesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimPowerOutletTemplatesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimPowerOutletTemplatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetModuletypeId returns the ModuletypeId property
func (m DcimPowerOutletTemplatesListQueryParameters) GetModuletypeId() []*int32 {
	return m.ModuletypeId
}

// SetModuletypeId sets the ModuletypeId property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetModuletypeId(val []*int32) {
	m.ModuletypeId = val
}

// GetModuletypeIdN returns the ModuletypeIdN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetModuletypeIdN() []*int32 {
	return m.ModuletypeIdN
}

// SetModuletypeIdN sets the ModuletypeIdN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetModuletypeIdN(val []*int32) {
	m.ModuletypeIdN = val
}

// GetName returns the Name property
func (m DcimPowerOutletTemplatesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimPowerOutletTemplatesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimPowerOutletTemplatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimPowerOutletTemplatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimPowerOutletTemplatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetType returns the Type property
func (m DcimPowerOutletTemplatesListQueryParameters) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetType(val string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimPowerOutletTemplatesListQueryParameters) GetTypeN() string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetTypeN(val string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimPowerOutletTemplatesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimPowerOutletTemplatesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
