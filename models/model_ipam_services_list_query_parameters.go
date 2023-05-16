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

// IpamServicesListQueryParameters is an object.
type IpamServicesListQueryParameters struct {
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
	// Description:
	Description []string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DescriptionEmpty:
	DescriptionEmpty []string `json:"description__empty,omitempty" mapstructure:"description__empty,omitempty"`
	// DescriptionIc:
	DescriptionIc []string `json:"description__ic,omitempty" mapstructure:"description__ic,omitempty"`
	// DescriptionIe:
	DescriptionIe []string `json:"description__ie,omitempty" mapstructure:"description__ie,omitempty"`
	// DescriptionIew:
	DescriptionIew []string `json:"description__iew,omitempty" mapstructure:"description__iew,omitempty"`
	// DescriptionIsw:
	DescriptionIsw []string `json:"description__isw,omitempty" mapstructure:"description__isw,omitempty"`
	// DescriptionN:
	DescriptionN []string `json:"description__n,omitempty" mapstructure:"description__n,omitempty"`
	// DescriptionNic:
	DescriptionNic []string `json:"description__nic,omitempty" mapstructure:"description__nic,omitempty"`
	// DescriptionNie:
	DescriptionNie []string `json:"description__nie,omitempty" mapstructure:"description__nie,omitempty"`
	// DescriptionNiew:
	DescriptionNiew []string `json:"description__niew,omitempty" mapstructure:"description__niew,omitempty"`
	// DescriptionNisw:
	DescriptionNisw []string `json:"description__nisw,omitempty" mapstructure:"description__nisw,omitempty"`
	// Device: Device (name)
	Device []*string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceN: Device (name)
	DeviceN []*string `json:"device__n,omitempty" mapstructure:"device__n,omitempty"`
	// DeviceId: Device (ID)
	DeviceId []*int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
	// DeviceIdN: Device (ID)
	DeviceIdN []*int32 `json:"device_id__n,omitempty" mapstructure:"device_id__n,omitempty"`
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
	// Ipaddress: IP address
	Ipaddress []string `json:"ipaddress,omitempty" mapstructure:"ipaddress,omitempty"`
	// IpaddressN: IP address
	IpaddressN []string `json:"ipaddress__n,omitempty" mapstructure:"ipaddress__n,omitempty"`
	// IpaddressId: IP address (ID)
	IpaddressId []int32 `json:"ipaddress_id,omitempty" mapstructure:"ipaddress_id,omitempty"`
	// IpaddressIdN: IP address (ID)
	IpaddressIdN []int32 `json:"ipaddress_id__n,omitempty" mapstructure:"ipaddress_id__n,omitempty"`
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
	// Port:
	Port float32 `json:"port,omitempty" mapstructure:"port,omitempty"`
	// Protocol: * `tcp` - TCP
	// * `udp` - UDP
	// * `sctp` - SCTP
	Protocol string `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	// ProtocolN: * `tcp` - TCP
	// * `udp` - UDP
	// * `sctp` - SCTP
	ProtocolN string `json:"protocol__n,omitempty" mapstructure:"protocol__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// VirtualMachine: Virtual machine (name)
	VirtualMachine []string `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
	// VirtualMachineN: Virtual machine (name)
	VirtualMachineN []string `json:"virtual_machine__n,omitempty" mapstructure:"virtual_machine__n,omitempty"`
	// VirtualMachineId: Virtual machine (ID)
	VirtualMachineId []*int32 `json:"virtual_machine_id,omitempty" mapstructure:"virtual_machine_id,omitempty"`
	// VirtualMachineIdN: Virtual machine (ID)
	VirtualMachineIdN []*int32 `json:"virtual_machine_id__n,omitempty" mapstructure:"virtual_machine_id__n,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamServicesListQueryParameters) Validate() error {
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
		"description": validation.Validate(
			m.Description,
		),
		"descriptionEmpty": validation.Validate(
			m.DescriptionEmpty,
		),
		"descriptionIc": validation.Validate(
			m.DescriptionIc,
		),
		"descriptionIe": validation.Validate(
			m.DescriptionIe,
		),
		"descriptionIew": validation.Validate(
			m.DescriptionIew,
		),
		"descriptionIsw": validation.Validate(
			m.DescriptionIsw,
		),
		"descriptionN": validation.Validate(
			m.DescriptionN,
		),
		"descriptionNic": validation.Validate(
			m.DescriptionNic,
		),
		"descriptionNie": validation.Validate(
			m.DescriptionNie,
		),
		"descriptionNiew": validation.Validate(
			m.DescriptionNiew,
		),
		"descriptionNisw": validation.Validate(
			m.DescriptionNisw,
		),
		"device": validation.Validate(
			m.Device,
		),
		"deviceN": validation.Validate(
			m.DeviceN,
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
		"ipaddress": validation.Validate(
			m.Ipaddress,
		),
		"ipaddressN": validation.Validate(
			m.IpaddressN,
		),
		"ipaddressId": validation.Validate(
			m.IpaddressId,
		),
		"ipaddressIdN": validation.Validate(
			m.IpaddressIdN,
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
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"virtualMachine": validation.Validate(
			m.VirtualMachine,
		),
		"virtualMachineN": validation.Validate(
			m.VirtualMachineN,
		),
		"virtualMachineId": validation.Validate(
			m.VirtualMachineId,
		),
		"virtualMachineIdN": validation.Validate(
			m.VirtualMachineIdN,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m IpamServicesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamServicesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamServicesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamServicesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamServicesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamServicesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamServicesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamServicesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamServicesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamServicesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamServicesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamServicesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamServicesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamServicesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m IpamServicesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IpamServicesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m IpamServicesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *IpamServicesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m IpamServicesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *IpamServicesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m IpamServicesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *IpamServicesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m IpamServicesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *IpamServicesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m IpamServicesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *IpamServicesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m IpamServicesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *IpamServicesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m IpamServicesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *IpamServicesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m IpamServicesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *IpamServicesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m IpamServicesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *IpamServicesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m IpamServicesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *IpamServicesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetDevice returns the Device property
func (m IpamServicesListQueryParameters) GetDevice() []*string {
	return m.Device
}

// SetDevice sets the Device property
func (m *IpamServicesListQueryParameters) SetDevice(val []*string) {
	m.Device = val
}

// GetDeviceN returns the DeviceN property
func (m IpamServicesListQueryParameters) GetDeviceN() []*string {
	return m.DeviceN
}

// SetDeviceN sets the DeviceN property
func (m *IpamServicesListQueryParameters) SetDeviceN(val []*string) {
	m.DeviceN = val
}

// GetDeviceId returns the DeviceId property
func (m IpamServicesListQueryParameters) GetDeviceId() []*int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *IpamServicesListQueryParameters) SetDeviceId(val []*int32) {
	m.DeviceId = val
}

// GetDeviceIdN returns the DeviceIdN property
func (m IpamServicesListQueryParameters) GetDeviceIdN() []*int32 {
	return m.DeviceIdN
}

// SetDeviceIdN sets the DeviceIdN property
func (m *IpamServicesListQueryParameters) SetDeviceIdN(val []*int32) {
	m.DeviceIdN = val
}

// GetId returns the Id property
func (m IpamServicesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamServicesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamServicesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamServicesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamServicesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamServicesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamServicesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamServicesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamServicesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamServicesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamServicesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamServicesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetIpaddress returns the Ipaddress property
func (m IpamServicesListQueryParameters) GetIpaddress() []string {
	return m.Ipaddress
}

// SetIpaddress sets the Ipaddress property
func (m *IpamServicesListQueryParameters) SetIpaddress(val []string) {
	m.Ipaddress = val
}

// GetIpaddressN returns the IpaddressN property
func (m IpamServicesListQueryParameters) GetIpaddressN() []string {
	return m.IpaddressN
}

// SetIpaddressN sets the IpaddressN property
func (m *IpamServicesListQueryParameters) SetIpaddressN(val []string) {
	m.IpaddressN = val
}

// GetIpaddressId returns the IpaddressId property
func (m IpamServicesListQueryParameters) GetIpaddressId() []int32 {
	return m.IpaddressId
}

// SetIpaddressId sets the IpaddressId property
func (m *IpamServicesListQueryParameters) SetIpaddressId(val []int32) {
	m.IpaddressId = val
}

// GetIpaddressIdN returns the IpaddressIdN property
func (m IpamServicesListQueryParameters) GetIpaddressIdN() []int32 {
	return m.IpaddressIdN
}

// SetIpaddressIdN sets the IpaddressIdN property
func (m *IpamServicesListQueryParameters) SetIpaddressIdN(val []int32) {
	m.IpaddressIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamServicesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamServicesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamServicesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamServicesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamServicesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamServicesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamServicesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamServicesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamServicesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamServicesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamServicesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamServicesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamServicesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamServicesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m IpamServicesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *IpamServicesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m IpamServicesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *IpamServicesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m IpamServicesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *IpamServicesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m IpamServicesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *IpamServicesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m IpamServicesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *IpamServicesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m IpamServicesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *IpamServicesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m IpamServicesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *IpamServicesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m IpamServicesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *IpamServicesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m IpamServicesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *IpamServicesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m IpamServicesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *IpamServicesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m IpamServicesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *IpamServicesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m IpamServicesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamServicesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamServicesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamServicesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPort returns the Port property
func (m IpamServicesListQueryParameters) GetPort() float32 {
	return m.Port
}

// SetPort sets the Port property
func (m *IpamServicesListQueryParameters) SetPort(val float32) {
	m.Port = val
}

// GetProtocol returns the Protocol property
func (m IpamServicesListQueryParameters) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *IpamServicesListQueryParameters) SetProtocol(val string) {
	m.Protocol = val
}

// GetProtocolN returns the ProtocolN property
func (m IpamServicesListQueryParameters) GetProtocolN() string {
	return m.ProtocolN
}

// SetProtocolN sets the ProtocolN property
func (m *IpamServicesListQueryParameters) SetProtocolN(val string) {
	m.ProtocolN = val
}

// GetQ returns the Q property
func (m IpamServicesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamServicesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m IpamServicesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamServicesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamServicesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamServicesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamServicesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamServicesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m IpamServicesListQueryParameters) GetVirtualMachine() []string {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *IpamServicesListQueryParameters) SetVirtualMachine(val []string) {
	m.VirtualMachine = val
}

// GetVirtualMachineN returns the VirtualMachineN property
func (m IpamServicesListQueryParameters) GetVirtualMachineN() []string {
	return m.VirtualMachineN
}

// SetVirtualMachineN sets the VirtualMachineN property
func (m *IpamServicesListQueryParameters) SetVirtualMachineN(val []string) {
	m.VirtualMachineN = val
}

// GetVirtualMachineId returns the VirtualMachineId property
func (m IpamServicesListQueryParameters) GetVirtualMachineId() []*int32 {
	return m.VirtualMachineId
}

// SetVirtualMachineId sets the VirtualMachineId property
func (m *IpamServicesListQueryParameters) SetVirtualMachineId(val []*int32) {
	m.VirtualMachineId = val
}

// GetVirtualMachineIdN returns the VirtualMachineIdN property
func (m IpamServicesListQueryParameters) GetVirtualMachineIdN() []*int32 {
	return m.VirtualMachineIdN
}

// SetVirtualMachineIdN sets the VirtualMachineIdN property
func (m *IpamServicesListQueryParameters) SetVirtualMachineIdN(val []*int32) {
	m.VirtualMachineIdN = val
}
