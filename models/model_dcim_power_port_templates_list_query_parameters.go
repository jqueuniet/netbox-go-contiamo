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

// DcimPowerPortTemplatesListQueryParameters is an object.
type DcimPowerPortTemplatesListQueryParameters struct {
	// AllocatedDraw:
	AllocatedDraw []int32 `json:"allocated_draw,omitempty" mapstructure:"allocated_draw,omitempty"`
	// AllocatedDrawGt:
	AllocatedDrawGt []int32 `json:"allocated_draw__gt,omitempty" mapstructure:"allocated_draw__gt,omitempty"`
	// AllocatedDrawGte:
	AllocatedDrawGte []int32 `json:"allocated_draw__gte,omitempty" mapstructure:"allocated_draw__gte,omitempty"`
	// AllocatedDrawLt:
	AllocatedDrawLt []int32 `json:"allocated_draw__lt,omitempty" mapstructure:"allocated_draw__lt,omitempty"`
	// AllocatedDrawLte:
	AllocatedDrawLte []int32 `json:"allocated_draw__lte,omitempty" mapstructure:"allocated_draw__lte,omitempty"`
	// AllocatedDrawN:
	AllocatedDrawN []int32 `json:"allocated_draw__n,omitempty" mapstructure:"allocated_draw__n,omitempty"`
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
	// MaximumDraw:
	MaximumDraw []int32 `json:"maximum_draw,omitempty" mapstructure:"maximum_draw,omitempty"`
	// MaximumDrawGt:
	MaximumDrawGt []int32 `json:"maximum_draw__gt,omitempty" mapstructure:"maximum_draw__gt,omitempty"`
	// MaximumDrawGte:
	MaximumDrawGte []int32 `json:"maximum_draw__gte,omitempty" mapstructure:"maximum_draw__gte,omitempty"`
	// MaximumDrawLt:
	MaximumDrawLt []int32 `json:"maximum_draw__lt,omitempty" mapstructure:"maximum_draw__lt,omitempty"`
	// MaximumDrawLte:
	MaximumDrawLte []int32 `json:"maximum_draw__lte,omitempty" mapstructure:"maximum_draw__lte,omitempty"`
	// MaximumDrawN:
	MaximumDrawN []int32 `json:"maximum_draw__n,omitempty" mapstructure:"maximum_draw__n,omitempty"`
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
	// Type: * `iec-60320-c6` - C6
	// * `iec-60320-c8` - C8
	// * `iec-60320-c14` - C14
	// * `iec-60320-c16` - C16
	// * `iec-60320-c20` - C20
	// * `iec-60320-c22` - C22
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
	// * `nema-1-15p` - NEMA 1-15P
	// * `nema-5-15p` - NEMA 5-15P
	// * `nema-5-20p` - NEMA 5-20P
	// * `nema-5-30p` - NEMA 5-30P
	// * `nema-5-50p` - NEMA 5-50P
	// * `nema-6-15p` - NEMA 6-15P
	// * `nema-6-20p` - NEMA 6-20P
	// * `nema-6-30p` - NEMA 6-30P
	// * `nema-6-50p` - NEMA 6-50P
	// * `nema-10-30p` - NEMA 10-30P
	// * `nema-10-50p` - NEMA 10-50P
	// * `nema-14-20p` - NEMA 14-20P
	// * `nema-14-30p` - NEMA 14-30P
	// * `nema-14-50p` - NEMA 14-50P
	// * `nema-14-60p` - NEMA 14-60P
	// * `nema-15-15p` - NEMA 15-15P
	// * `nema-15-20p` - NEMA 15-20P
	// * `nema-15-30p` - NEMA 15-30P
	// * `nema-15-50p` - NEMA 15-50P
	// * `nema-15-60p` - NEMA 15-60P
	// * `nema-l1-15p` - NEMA L1-15P
	// * `nema-l5-15p` - NEMA L5-15P
	// * `nema-l5-20p` - NEMA L5-20P
	// * `nema-l5-30p` - NEMA L5-30P
	// * `nema-l5-50p` - NEMA L5-50P
	// * `nema-l6-15p` - NEMA L6-15P
	// * `nema-l6-20p` - NEMA L6-20P
	// * `nema-l6-30p` - NEMA L6-30P
	// * `nema-l6-50p` - NEMA L6-50P
	// * `nema-l10-30p` - NEMA L10-30P
	// * `nema-l14-20p` - NEMA L14-20P
	// * `nema-l14-30p` - NEMA L14-30P
	// * `nema-l14-50p` - NEMA L14-50P
	// * `nema-l14-60p` - NEMA L14-60P
	// * `nema-l15-20p` - NEMA L15-20P
	// * `nema-l15-30p` - NEMA L15-30P
	// * `nema-l15-50p` - NEMA L15-50P
	// * `nema-l15-60p` - NEMA L15-60P
	// * `nema-l21-20p` - NEMA L21-20P
	// * `nema-l21-30p` - NEMA L21-30P
	// * `nema-l22-30p` - NEMA L22-30P
	// * `cs6361c` - CS6361C
	// * `cs6365c` - CS6365C
	// * `cs8165c` - CS8165C
	// * `cs8265c` - CS8265C
	// * `cs8365c` - CS8365C
	// * `cs8465c` - CS8465C
	// * `ita-c` - ITA Type C (CEE 7/16)
	// * `ita-e` - ITA Type E (CEE 7/6)
	// * `ita-f` - ITA Type F (CEE 7/4)
	// * `ita-ef` - ITA Type E/F (CEE 7/7)
	// * `ita-g` - ITA Type G (BS 1363)
	// * `ita-h` - ITA Type H
	// * `ita-i` - ITA Type I
	// * `ita-j` - ITA Type J
	// * `ita-k` - ITA Type K
	// * `ita-l` - ITA Type L (CEI 23-50)
	// * `ita-m` - ITA Type M (BS 546)
	// * `ita-n` - ITA Type N
	// * `ita-o` - ITA Type O
	// * `usb-a` - USB Type A
	// * `usb-b` - USB Type B
	// * `usb-c` - USB Type C
	// * `usb-mini-a` - USB Mini A
	// * `usb-mini-b` - USB Mini B
	// * `usb-micro-a` - USB Micro A
	// * `usb-micro-b` - USB Micro B
	// * `usb-micro-ab` - USB Micro AB
	// * `usb-3-b` - USB 3.0 Type B
	// * `usb-3-micro-b` - USB 3.0 Micro B
	// * `dc-terminal` - DC Terminal
	// * `saf-d-grid` - Saf-D-Grid
	// * `neutrik-powercon-20` - Neutrik powerCON (20A)
	// * `neutrik-powercon-32` - Neutrik powerCON (32A)
	// * `neutrik-powercon-true1` - Neutrik powerCON TRUE1
	// * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP
	// * `ubiquiti-smartpower` - Ubiquiti SmartPower
	// * `hardwired` - Hardwired
	// * `other` - Other
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `iec-60320-c6` - C6
	// * `iec-60320-c8` - C8
	// * `iec-60320-c14` - C14
	// * `iec-60320-c16` - C16
	// * `iec-60320-c20` - C20
	// * `iec-60320-c22` - C22
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
	// * `nema-1-15p` - NEMA 1-15P
	// * `nema-5-15p` - NEMA 5-15P
	// * `nema-5-20p` - NEMA 5-20P
	// * `nema-5-30p` - NEMA 5-30P
	// * `nema-5-50p` - NEMA 5-50P
	// * `nema-6-15p` - NEMA 6-15P
	// * `nema-6-20p` - NEMA 6-20P
	// * `nema-6-30p` - NEMA 6-30P
	// * `nema-6-50p` - NEMA 6-50P
	// * `nema-10-30p` - NEMA 10-30P
	// * `nema-10-50p` - NEMA 10-50P
	// * `nema-14-20p` - NEMA 14-20P
	// * `nema-14-30p` - NEMA 14-30P
	// * `nema-14-50p` - NEMA 14-50P
	// * `nema-14-60p` - NEMA 14-60P
	// * `nema-15-15p` - NEMA 15-15P
	// * `nema-15-20p` - NEMA 15-20P
	// * `nema-15-30p` - NEMA 15-30P
	// * `nema-15-50p` - NEMA 15-50P
	// * `nema-15-60p` - NEMA 15-60P
	// * `nema-l1-15p` - NEMA L1-15P
	// * `nema-l5-15p` - NEMA L5-15P
	// * `nema-l5-20p` - NEMA L5-20P
	// * `nema-l5-30p` - NEMA L5-30P
	// * `nema-l5-50p` - NEMA L5-50P
	// * `nema-l6-15p` - NEMA L6-15P
	// * `nema-l6-20p` - NEMA L6-20P
	// * `nema-l6-30p` - NEMA L6-30P
	// * `nema-l6-50p` - NEMA L6-50P
	// * `nema-l10-30p` - NEMA L10-30P
	// * `nema-l14-20p` - NEMA L14-20P
	// * `nema-l14-30p` - NEMA L14-30P
	// * `nema-l14-50p` - NEMA L14-50P
	// * `nema-l14-60p` - NEMA L14-60P
	// * `nema-l15-20p` - NEMA L15-20P
	// * `nema-l15-30p` - NEMA L15-30P
	// * `nema-l15-50p` - NEMA L15-50P
	// * `nema-l15-60p` - NEMA L15-60P
	// * `nema-l21-20p` - NEMA L21-20P
	// * `nema-l21-30p` - NEMA L21-30P
	// * `nema-l22-30p` - NEMA L22-30P
	// * `cs6361c` - CS6361C
	// * `cs6365c` - CS6365C
	// * `cs8165c` - CS8165C
	// * `cs8265c` - CS8265C
	// * `cs8365c` - CS8365C
	// * `cs8465c` - CS8465C
	// * `ita-c` - ITA Type C (CEE 7/16)
	// * `ita-e` - ITA Type E (CEE 7/6)
	// * `ita-f` - ITA Type F (CEE 7/4)
	// * `ita-ef` - ITA Type E/F (CEE 7/7)
	// * `ita-g` - ITA Type G (BS 1363)
	// * `ita-h` - ITA Type H
	// * `ita-i` - ITA Type I
	// * `ita-j` - ITA Type J
	// * `ita-k` - ITA Type K
	// * `ita-l` - ITA Type L (CEI 23-50)
	// * `ita-m` - ITA Type M (BS 546)
	// * `ita-n` - ITA Type N
	// * `ita-o` - ITA Type O
	// * `usb-a` - USB Type A
	// * `usb-b` - USB Type B
	// * `usb-c` - USB Type C
	// * `usb-mini-a` - USB Mini A
	// * `usb-mini-b` - USB Mini B
	// * `usb-micro-a` - USB Micro A
	// * `usb-micro-b` - USB Micro B
	// * `usb-micro-ab` - USB Micro AB
	// * `usb-3-b` - USB 3.0 Type B
	// * `usb-3-micro-b` - USB 3.0 Micro B
	// * `dc-terminal` - DC Terminal
	// * `saf-d-grid` - Saf-D-Grid
	// * `neutrik-powercon-20` - Neutrik powerCON (20A)
	// * `neutrik-powercon-32` - Neutrik powerCON (32A)
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
func (m DcimPowerPortTemplatesListQueryParameters) Validate() error {
	return validation.Errors{
		"allocatedDraw": validation.Validate(
			m.AllocatedDraw,
		),
		"allocatedDrawGt": validation.Validate(
			m.AllocatedDrawGt,
		),
		"allocatedDrawGte": validation.Validate(
			m.AllocatedDrawGte,
		),
		"allocatedDrawLt": validation.Validate(
			m.AllocatedDrawLt,
		),
		"allocatedDrawLte": validation.Validate(
			m.AllocatedDrawLte,
		),
		"allocatedDrawN": validation.Validate(
			m.AllocatedDrawN,
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
		"maximumDraw": validation.Validate(
			m.MaximumDraw,
		),
		"maximumDrawGt": validation.Validate(
			m.MaximumDrawGt,
		),
		"maximumDrawGte": validation.Validate(
			m.MaximumDrawGte,
		),
		"maximumDrawLt": validation.Validate(
			m.MaximumDrawLt,
		),
		"maximumDrawLte": validation.Validate(
			m.MaximumDrawLte,
		),
		"maximumDrawN": validation.Validate(
			m.MaximumDrawN,
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

// GetAllocatedDraw returns the AllocatedDraw property
func (m DcimPowerPortTemplatesListQueryParameters) GetAllocatedDraw() []int32 {
	return m.AllocatedDraw
}

// SetAllocatedDraw sets the AllocatedDraw property
func (m *DcimPowerPortTemplatesListQueryParameters) SetAllocatedDraw(val []int32) {
	m.AllocatedDraw = val
}

// GetAllocatedDrawGt returns the AllocatedDrawGt property
func (m DcimPowerPortTemplatesListQueryParameters) GetAllocatedDrawGt() []int32 {
	return m.AllocatedDrawGt
}

// SetAllocatedDrawGt sets the AllocatedDrawGt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetAllocatedDrawGt(val []int32) {
	m.AllocatedDrawGt = val
}

// GetAllocatedDrawGte returns the AllocatedDrawGte property
func (m DcimPowerPortTemplatesListQueryParameters) GetAllocatedDrawGte() []int32 {
	return m.AllocatedDrawGte
}

// SetAllocatedDrawGte sets the AllocatedDrawGte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetAllocatedDrawGte(val []int32) {
	m.AllocatedDrawGte = val
}

// GetAllocatedDrawLt returns the AllocatedDrawLt property
func (m DcimPowerPortTemplatesListQueryParameters) GetAllocatedDrawLt() []int32 {
	return m.AllocatedDrawLt
}

// SetAllocatedDrawLt sets the AllocatedDrawLt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetAllocatedDrawLt(val []int32) {
	m.AllocatedDrawLt = val
}

// GetAllocatedDrawLte returns the AllocatedDrawLte property
func (m DcimPowerPortTemplatesListQueryParameters) GetAllocatedDrawLte() []int32 {
	return m.AllocatedDrawLte
}

// SetAllocatedDrawLte sets the AllocatedDrawLte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetAllocatedDrawLte(val []int32) {
	m.AllocatedDrawLte = val
}

// GetAllocatedDrawN returns the AllocatedDrawN property
func (m DcimPowerPortTemplatesListQueryParameters) GetAllocatedDrawN() []int32 {
	return m.AllocatedDrawN
}

// SetAllocatedDrawN sets the AllocatedDrawN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetAllocatedDrawN(val []int32) {
	m.AllocatedDrawN = val
}

// GetCreated returns the Created property
func (m DcimPowerPortTemplatesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimPowerPortTemplatesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimPowerPortTemplatesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimPowerPortTemplatesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimPowerPortTemplatesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimPowerPortTemplatesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimPowerPortTemplatesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimPowerPortTemplatesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimPowerPortTemplatesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevicetypeId returns the DevicetypeId property
func (m DcimPowerPortTemplatesListQueryParameters) GetDevicetypeId() []*int32 {
	return m.DevicetypeId
}

// SetDevicetypeId sets the DevicetypeId property
func (m *DcimPowerPortTemplatesListQueryParameters) SetDevicetypeId(val []*int32) {
	m.DevicetypeId = val
}

// GetDevicetypeIdN returns the DevicetypeIdN property
func (m DcimPowerPortTemplatesListQueryParameters) GetDevicetypeIdN() []*int32 {
	return m.DevicetypeIdN
}

// SetDevicetypeIdN sets the DevicetypeIdN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetDevicetypeIdN(val []*int32) {
	m.DevicetypeIdN = val
}

// GetId returns the Id property
func (m DcimPowerPortTemplatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimPowerPortTemplatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimPowerPortTemplatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimPowerPortTemplatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimPowerPortTemplatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimPowerPortTemplatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimPowerPortTemplatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimPowerPortTemplatesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimPowerPortTemplatesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimPowerPortTemplatesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimPowerPortTemplatesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimPowerPortTemplatesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimPowerPortTemplatesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimPowerPortTemplatesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimPowerPortTemplatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimPowerPortTemplatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetMaximumDraw returns the MaximumDraw property
func (m DcimPowerPortTemplatesListQueryParameters) GetMaximumDraw() []int32 {
	return m.MaximumDraw
}

// SetMaximumDraw sets the MaximumDraw property
func (m *DcimPowerPortTemplatesListQueryParameters) SetMaximumDraw(val []int32) {
	m.MaximumDraw = val
}

// GetMaximumDrawGt returns the MaximumDrawGt property
func (m DcimPowerPortTemplatesListQueryParameters) GetMaximumDrawGt() []int32 {
	return m.MaximumDrawGt
}

// SetMaximumDrawGt sets the MaximumDrawGt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetMaximumDrawGt(val []int32) {
	m.MaximumDrawGt = val
}

// GetMaximumDrawGte returns the MaximumDrawGte property
func (m DcimPowerPortTemplatesListQueryParameters) GetMaximumDrawGte() []int32 {
	return m.MaximumDrawGte
}

// SetMaximumDrawGte sets the MaximumDrawGte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetMaximumDrawGte(val []int32) {
	m.MaximumDrawGte = val
}

// GetMaximumDrawLt returns the MaximumDrawLt property
func (m DcimPowerPortTemplatesListQueryParameters) GetMaximumDrawLt() []int32 {
	return m.MaximumDrawLt
}

// SetMaximumDrawLt sets the MaximumDrawLt property
func (m *DcimPowerPortTemplatesListQueryParameters) SetMaximumDrawLt(val []int32) {
	m.MaximumDrawLt = val
}

// GetMaximumDrawLte returns the MaximumDrawLte property
func (m DcimPowerPortTemplatesListQueryParameters) GetMaximumDrawLte() []int32 {
	return m.MaximumDrawLte
}

// SetMaximumDrawLte sets the MaximumDrawLte property
func (m *DcimPowerPortTemplatesListQueryParameters) SetMaximumDrawLte(val []int32) {
	m.MaximumDrawLte = val
}

// GetMaximumDrawN returns the MaximumDrawN property
func (m DcimPowerPortTemplatesListQueryParameters) GetMaximumDrawN() []int32 {
	return m.MaximumDrawN
}

// SetMaximumDrawN sets the MaximumDrawN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetMaximumDrawN(val []int32) {
	m.MaximumDrawN = val
}

// GetModuletypeId returns the ModuletypeId property
func (m DcimPowerPortTemplatesListQueryParameters) GetModuletypeId() []*int32 {
	return m.ModuletypeId
}

// SetModuletypeId sets the ModuletypeId property
func (m *DcimPowerPortTemplatesListQueryParameters) SetModuletypeId(val []*int32) {
	m.ModuletypeId = val
}

// GetModuletypeIdN returns the ModuletypeIdN property
func (m DcimPowerPortTemplatesListQueryParameters) GetModuletypeIdN() []*int32 {
	return m.ModuletypeIdN
}

// SetModuletypeIdN sets the ModuletypeIdN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetModuletypeIdN(val []*int32) {
	m.ModuletypeIdN = val
}

// GetName returns the Name property
func (m DcimPowerPortTemplatesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimPowerPortTemplatesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimPowerPortTemplatesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimPowerPortTemplatesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimPowerPortTemplatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimPowerPortTemplatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimPowerPortTemplatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimPowerPortTemplatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m DcimPowerPortTemplatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimPowerPortTemplatesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetType returns the Type property
func (m DcimPowerPortTemplatesListQueryParameters) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimPowerPortTemplatesListQueryParameters) SetType(val string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimPowerPortTemplatesListQueryParameters) GetTypeN() string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimPowerPortTemplatesListQueryParameters) SetTypeN(val string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimPowerPortTemplatesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimPowerPortTemplatesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
