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

// IpamIpAddressesListQueryParameters is an object.
type IpamIpAddressesListQueryParameters struct {
	// Address:
	Address []string `json:"address,omitempty" mapstructure:"address,omitempty"`
	// AssignedToInterface: Is assigned to an interface
	AssignedToInterface bool `json:"assigned_to_interface,omitempty" mapstructure:"assigned_to_interface,omitempty"`
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
	// Device:
	Device []string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceId:
	DeviceId []int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
	// DnsName:
	DnsName []string `json:"dns_name,omitempty" mapstructure:"dns_name,omitempty"`
	// DnsNameEmpty:
	DnsNameEmpty []string `json:"dns_name__empty,omitempty" mapstructure:"dns_name__empty,omitempty"`
	// DnsNameIc:
	DnsNameIc []string `json:"dns_name__ic,omitempty" mapstructure:"dns_name__ic,omitempty"`
	// DnsNameIe:
	DnsNameIe []string `json:"dns_name__ie,omitempty" mapstructure:"dns_name__ie,omitempty"`
	// DnsNameIew:
	DnsNameIew []string `json:"dns_name__iew,omitempty" mapstructure:"dns_name__iew,omitempty"`
	// DnsNameIsw:
	DnsNameIsw []string `json:"dns_name__isw,omitempty" mapstructure:"dns_name__isw,omitempty"`
	// DnsNameN:
	DnsNameN []string `json:"dns_name__n,omitempty" mapstructure:"dns_name__n,omitempty"`
	// DnsNameNic:
	DnsNameNic []string `json:"dns_name__nic,omitempty" mapstructure:"dns_name__nic,omitempty"`
	// DnsNameNie:
	DnsNameNie []string `json:"dns_name__nie,omitempty" mapstructure:"dns_name__nie,omitempty"`
	// DnsNameNiew:
	DnsNameNiew []string `json:"dns_name__niew,omitempty" mapstructure:"dns_name__niew,omitempty"`
	// DnsNameNisw:
	DnsNameNisw []string `json:"dns_name__nisw,omitempty" mapstructure:"dns_name__nisw,omitempty"`
	// Family:
	Family float32 `json:"family,omitempty" mapstructure:"family,omitempty"`
	// FhrpgroupId: FHRP group (ID)
	FhrpgroupId []int32 `json:"fhrpgroup_id,omitempty" mapstructure:"fhrpgroup_id,omitempty"`
	// FhrpgroupIdN: FHRP group (ID)
	FhrpgroupIdN []int32 `json:"fhrpgroup_id__n,omitempty" mapstructure:"fhrpgroup_id__n,omitempty"`
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
	// Interface: Interface (name)
	Interface []string `json:"interface,omitempty" mapstructure:"interface,omitempty"`
	// InterfaceN: Interface (name)
	InterfaceN []string `json:"interface__n,omitempty" mapstructure:"interface__n,omitempty"`
	// InterfaceId: Interface (ID)
	InterfaceId []int32 `json:"interface_id,omitempty" mapstructure:"interface_id,omitempty"`
	// InterfaceIdN: Interface (ID)
	InterfaceIdN []int32 `json:"interface_id__n,omitempty" mapstructure:"interface_id__n,omitempty"`
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
	// MaskLength: Mask length
	MaskLength float32 `json:"mask_length,omitempty" mapstructure:"mask_length,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Parent:
	Parent []string `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// PresentInVrf: VRF (RD)
	PresentInVrf string `json:"present_in_vrf,omitempty" mapstructure:"present_in_vrf,omitempty"`
	// PresentInVrfId: VRF
	PresentInVrfId string `json:"present_in_vrf_id,omitempty" mapstructure:"present_in_vrf_id,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Role: The functional role of this IP
	Role []string `json:"role,omitempty" mapstructure:"role,omitempty"`
	// RoleN: The functional role of this IP
	RoleN []string `json:"role__n,omitempty" mapstructure:"role__n,omitempty"`
	// Status: The operational status of this IP
	Status []string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// StatusN: The operational status of this IP
	StatusN []string `json:"status__n,omitempty" mapstructure:"status__n,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// Tenant: Tenant (slug)
	Tenant []string `json:"tenant,omitempty" mapstructure:"tenant,omitempty"`
	// TenantN: Tenant (slug)
	TenantN []string `json:"tenant__n,omitempty" mapstructure:"tenant__n,omitempty"`
	// TenantGroup: Tenant Group (slug)
	TenantGroup []int32 `json:"tenant_group,omitempty" mapstructure:"tenant_group,omitempty"`
	// TenantGroupN: Tenant Group (slug)
	TenantGroupN []int32 `json:"tenant_group__n,omitempty" mapstructure:"tenant_group__n,omitempty"`
	// TenantGroupId: Tenant Group (ID)
	TenantGroupId []int32 `json:"tenant_group_id,omitempty" mapstructure:"tenant_group_id,omitempty"`
	// TenantGroupIdN: Tenant Group (ID)
	TenantGroupIdN []int32 `json:"tenant_group_id__n,omitempty" mapstructure:"tenant_group_id__n,omitempty"`
	// TenantId: Tenant (ID)
	TenantId []*int32 `json:"tenant_id,omitempty" mapstructure:"tenant_id,omitempty"`
	// TenantIdN: Tenant (ID)
	TenantIdN []*int32 `json:"tenant_id__n,omitempty" mapstructure:"tenant_id__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// VirtualMachine:
	VirtualMachine []string `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
	// VirtualMachineId:
	VirtualMachineId []int32 `json:"virtual_machine_id,omitempty" mapstructure:"virtual_machine_id,omitempty"`
	// Vminterface: VM interface (name)
	Vminterface []string `json:"vminterface,omitempty" mapstructure:"vminterface,omitempty"`
	// VminterfaceN: VM interface (name)
	VminterfaceN []string `json:"vminterface__n,omitempty" mapstructure:"vminterface__n,omitempty"`
	// VminterfaceId: VM interface (ID)
	VminterfaceId []int32 `json:"vminterface_id,omitempty" mapstructure:"vminterface_id,omitempty"`
	// VminterfaceIdN: VM interface (ID)
	VminterfaceIdN []int32 `json:"vminterface_id__n,omitempty" mapstructure:"vminterface_id__n,omitempty"`
	// Vrf: VRF (RD)
	Vrf []*string `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
	// VrfN: VRF (RD)
	VrfN []*string `json:"vrf__n,omitempty" mapstructure:"vrf__n,omitempty"`
	// VrfId: VRF
	VrfId []*int32 `json:"vrf_id,omitempty" mapstructure:"vrf_id,omitempty"`
	// VrfIdN: VRF
	VrfIdN []*int32 `json:"vrf_id__n,omitempty" mapstructure:"vrf_id__n,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamIpAddressesListQueryParameters) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address,
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
		"deviceId": validation.Validate(
			m.DeviceId,
		),
		"dnsName": validation.Validate(
			m.DnsName,
		),
		"dnsNameEmpty": validation.Validate(
			m.DnsNameEmpty,
		),
		"dnsNameIc": validation.Validate(
			m.DnsNameIc,
		),
		"dnsNameIe": validation.Validate(
			m.DnsNameIe,
		),
		"dnsNameIew": validation.Validate(
			m.DnsNameIew,
		),
		"dnsNameIsw": validation.Validate(
			m.DnsNameIsw,
		),
		"dnsNameN": validation.Validate(
			m.DnsNameN,
		),
		"dnsNameNic": validation.Validate(
			m.DnsNameNic,
		),
		"dnsNameNie": validation.Validate(
			m.DnsNameNie,
		),
		"dnsNameNiew": validation.Validate(
			m.DnsNameNiew,
		),
		"dnsNameNisw": validation.Validate(
			m.DnsNameNisw,
		),
		"fhrpgroupId": validation.Validate(
			m.FhrpgroupId,
		),
		"fhrpgroupIdN": validation.Validate(
			m.FhrpgroupIdN,
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
		"interface": validation.Validate(
			m.Interface,
		),
		"interfaceN": validation.Validate(
			m.InterfaceN,
		),
		"interfaceId": validation.Validate(
			m.InterfaceId,
		),
		"interfaceIdN": validation.Validate(
			m.InterfaceIdN,
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
		"parent": validation.Validate(
			m.Parent,
		),
		"role": validation.Validate(
			m.Role,
		),
		"roleN": validation.Validate(
			m.RoleN,
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
		"tenant": validation.Validate(
			m.Tenant,
		),
		"tenantN": validation.Validate(
			m.TenantN,
		),
		"tenantGroup": validation.Validate(
			m.TenantGroup,
		),
		"tenantGroupN": validation.Validate(
			m.TenantGroupN,
		),
		"tenantGroupId": validation.Validate(
			m.TenantGroupId,
		),
		"tenantGroupIdN": validation.Validate(
			m.TenantGroupIdN,
		),
		"tenantId": validation.Validate(
			m.TenantId,
		),
		"tenantIdN": validation.Validate(
			m.TenantIdN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"virtualMachine": validation.Validate(
			m.VirtualMachine,
		),
		"virtualMachineId": validation.Validate(
			m.VirtualMachineId,
		),
		"vminterface": validation.Validate(
			m.Vminterface,
		),
		"vminterfaceN": validation.Validate(
			m.VminterfaceN,
		),
		"vminterfaceId": validation.Validate(
			m.VminterfaceId,
		),
		"vminterfaceIdN": validation.Validate(
			m.VminterfaceIdN,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
		"vrfN": validation.Validate(
			m.VrfN,
		),
		"vrfId": validation.Validate(
			m.VrfId,
		),
		"vrfIdN": validation.Validate(
			m.VrfIdN,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m IpamIpAddressesListQueryParameters) GetAddress() []string {
	return m.Address
}

// SetAddress sets the Address property
func (m *IpamIpAddressesListQueryParameters) SetAddress(val []string) {
	m.Address = val
}

// GetAssignedToInterface returns the AssignedToInterface property
func (m IpamIpAddressesListQueryParameters) GetAssignedToInterface() bool {
	return m.AssignedToInterface
}

// SetAssignedToInterface sets the AssignedToInterface property
func (m *IpamIpAddressesListQueryParameters) SetAssignedToInterface(val bool) {
	m.AssignedToInterface = val
}

// GetCreated returns the Created property
func (m IpamIpAddressesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamIpAddressesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamIpAddressesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamIpAddressesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamIpAddressesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamIpAddressesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamIpAddressesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamIpAddressesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamIpAddressesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamIpAddressesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamIpAddressesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamIpAddressesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamIpAddressesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamIpAddressesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m IpamIpAddressesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *IpamIpAddressesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m IpamIpAddressesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m IpamIpAddressesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m IpamIpAddressesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m IpamIpAddressesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m IpamIpAddressesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m IpamIpAddressesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m IpamIpAddressesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m IpamIpAddressesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m IpamIpAddressesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m IpamIpAddressesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *IpamIpAddressesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetDevice returns the Device property
func (m IpamIpAddressesListQueryParameters) GetDevice() []string {
	return m.Device
}

// SetDevice sets the Device property
func (m *IpamIpAddressesListQueryParameters) SetDevice(val []string) {
	m.Device = val
}

// GetDeviceId returns the DeviceId property
func (m IpamIpAddressesListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *IpamIpAddressesListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetDnsName returns the DnsName property
func (m IpamIpAddressesListQueryParameters) GetDnsName() []string {
	return m.DnsName
}

// SetDnsName sets the DnsName property
func (m *IpamIpAddressesListQueryParameters) SetDnsName(val []string) {
	m.DnsName = val
}

// GetDnsNameEmpty returns the DnsNameEmpty property
func (m IpamIpAddressesListQueryParameters) GetDnsNameEmpty() []string {
	return m.DnsNameEmpty
}

// SetDnsNameEmpty sets the DnsNameEmpty property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameEmpty(val []string) {
	m.DnsNameEmpty = val
}

// GetDnsNameIc returns the DnsNameIc property
func (m IpamIpAddressesListQueryParameters) GetDnsNameIc() []string {
	return m.DnsNameIc
}

// SetDnsNameIc sets the DnsNameIc property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameIc(val []string) {
	m.DnsNameIc = val
}

// GetDnsNameIe returns the DnsNameIe property
func (m IpamIpAddressesListQueryParameters) GetDnsNameIe() []string {
	return m.DnsNameIe
}

// SetDnsNameIe sets the DnsNameIe property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameIe(val []string) {
	m.DnsNameIe = val
}

// GetDnsNameIew returns the DnsNameIew property
func (m IpamIpAddressesListQueryParameters) GetDnsNameIew() []string {
	return m.DnsNameIew
}

// SetDnsNameIew sets the DnsNameIew property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameIew(val []string) {
	m.DnsNameIew = val
}

// GetDnsNameIsw returns the DnsNameIsw property
func (m IpamIpAddressesListQueryParameters) GetDnsNameIsw() []string {
	return m.DnsNameIsw
}

// SetDnsNameIsw sets the DnsNameIsw property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameIsw(val []string) {
	m.DnsNameIsw = val
}

// GetDnsNameN returns the DnsNameN property
func (m IpamIpAddressesListQueryParameters) GetDnsNameN() []string {
	return m.DnsNameN
}

// SetDnsNameN sets the DnsNameN property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameN(val []string) {
	m.DnsNameN = val
}

// GetDnsNameNic returns the DnsNameNic property
func (m IpamIpAddressesListQueryParameters) GetDnsNameNic() []string {
	return m.DnsNameNic
}

// SetDnsNameNic sets the DnsNameNic property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameNic(val []string) {
	m.DnsNameNic = val
}

// GetDnsNameNie returns the DnsNameNie property
func (m IpamIpAddressesListQueryParameters) GetDnsNameNie() []string {
	return m.DnsNameNie
}

// SetDnsNameNie sets the DnsNameNie property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameNie(val []string) {
	m.DnsNameNie = val
}

// GetDnsNameNiew returns the DnsNameNiew property
func (m IpamIpAddressesListQueryParameters) GetDnsNameNiew() []string {
	return m.DnsNameNiew
}

// SetDnsNameNiew sets the DnsNameNiew property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameNiew(val []string) {
	m.DnsNameNiew = val
}

// GetDnsNameNisw returns the DnsNameNisw property
func (m IpamIpAddressesListQueryParameters) GetDnsNameNisw() []string {
	return m.DnsNameNisw
}

// SetDnsNameNisw sets the DnsNameNisw property
func (m *IpamIpAddressesListQueryParameters) SetDnsNameNisw(val []string) {
	m.DnsNameNisw = val
}

// GetFamily returns the Family property
func (m IpamIpAddressesListQueryParameters) GetFamily() float32 {
	return m.Family
}

// SetFamily sets the Family property
func (m *IpamIpAddressesListQueryParameters) SetFamily(val float32) {
	m.Family = val
}

// GetFhrpgroupId returns the FhrpgroupId property
func (m IpamIpAddressesListQueryParameters) GetFhrpgroupId() []int32 {
	return m.FhrpgroupId
}

// SetFhrpgroupId sets the FhrpgroupId property
func (m *IpamIpAddressesListQueryParameters) SetFhrpgroupId(val []int32) {
	m.FhrpgroupId = val
}

// GetFhrpgroupIdN returns the FhrpgroupIdN property
func (m IpamIpAddressesListQueryParameters) GetFhrpgroupIdN() []int32 {
	return m.FhrpgroupIdN
}

// SetFhrpgroupIdN sets the FhrpgroupIdN property
func (m *IpamIpAddressesListQueryParameters) SetFhrpgroupIdN(val []int32) {
	m.FhrpgroupIdN = val
}

// GetId returns the Id property
func (m IpamIpAddressesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamIpAddressesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamIpAddressesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamIpAddressesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamIpAddressesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamIpAddressesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamIpAddressesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamIpAddressesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamIpAddressesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamIpAddressesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamIpAddressesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamIpAddressesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterface returns the Interface property
func (m IpamIpAddressesListQueryParameters) GetInterface() []string {
	return m.Interface
}

// SetInterface sets the Interface property
func (m *IpamIpAddressesListQueryParameters) SetInterface(val []string) {
	m.Interface = val
}

// GetInterfaceN returns the InterfaceN property
func (m IpamIpAddressesListQueryParameters) GetInterfaceN() []string {
	return m.InterfaceN
}

// SetInterfaceN sets the InterfaceN property
func (m *IpamIpAddressesListQueryParameters) SetInterfaceN(val []string) {
	m.InterfaceN = val
}

// GetInterfaceId returns the InterfaceId property
func (m IpamIpAddressesListQueryParameters) GetInterfaceId() []int32 {
	return m.InterfaceId
}

// SetInterfaceId sets the InterfaceId property
func (m *IpamIpAddressesListQueryParameters) SetInterfaceId(val []int32) {
	m.InterfaceId = val
}

// GetInterfaceIdN returns the InterfaceIdN property
func (m IpamIpAddressesListQueryParameters) GetInterfaceIdN() []int32 {
	return m.InterfaceIdN
}

// SetInterfaceIdN sets the InterfaceIdN property
func (m *IpamIpAddressesListQueryParameters) SetInterfaceIdN(val []int32) {
	m.InterfaceIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamIpAddressesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamIpAddressesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamIpAddressesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamIpAddressesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamIpAddressesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamIpAddressesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamIpAddressesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamIpAddressesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamIpAddressesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamIpAddressesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamIpAddressesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamIpAddressesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamIpAddressesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamIpAddressesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetMaskLength returns the MaskLength property
func (m IpamIpAddressesListQueryParameters) GetMaskLength() float32 {
	return m.MaskLength
}

// SetMaskLength sets the MaskLength property
func (m *IpamIpAddressesListQueryParameters) SetMaskLength(val float32) {
	m.MaskLength = val
}

// GetOffset returns the Offset property
func (m IpamIpAddressesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamIpAddressesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamIpAddressesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamIpAddressesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetParent returns the Parent property
func (m IpamIpAddressesListQueryParameters) GetParent() []string {
	return m.Parent
}

// SetParent sets the Parent property
func (m *IpamIpAddressesListQueryParameters) SetParent(val []string) {
	m.Parent = val
}

// GetPresentInVrf returns the PresentInVrf property
func (m IpamIpAddressesListQueryParameters) GetPresentInVrf() string {
	return m.PresentInVrf
}

// SetPresentInVrf sets the PresentInVrf property
func (m *IpamIpAddressesListQueryParameters) SetPresentInVrf(val string) {
	m.PresentInVrf = val
}

// GetPresentInVrfId returns the PresentInVrfId property
func (m IpamIpAddressesListQueryParameters) GetPresentInVrfId() string {
	return m.PresentInVrfId
}

// SetPresentInVrfId sets the PresentInVrfId property
func (m *IpamIpAddressesListQueryParameters) SetPresentInVrfId(val string) {
	m.PresentInVrfId = val
}

// GetQ returns the Q property
func (m IpamIpAddressesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamIpAddressesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRole returns the Role property
func (m IpamIpAddressesListQueryParameters) GetRole() []string {
	return m.Role
}

// SetRole sets the Role property
func (m *IpamIpAddressesListQueryParameters) SetRole(val []string) {
	m.Role = val
}

// GetRoleN returns the RoleN property
func (m IpamIpAddressesListQueryParameters) GetRoleN() []string {
	return m.RoleN
}

// SetRoleN sets the RoleN property
func (m *IpamIpAddressesListQueryParameters) SetRoleN(val []string) {
	m.RoleN = val
}

// GetStatus returns the Status property
func (m IpamIpAddressesListQueryParameters) GetStatus() []string {
	return m.Status
}

// SetStatus sets the Status property
func (m *IpamIpAddressesListQueryParameters) SetStatus(val []string) {
	m.Status = val
}

// GetStatusN returns the StatusN property
func (m IpamIpAddressesListQueryParameters) GetStatusN() []string {
	return m.StatusN
}

// SetStatusN sets the StatusN property
func (m *IpamIpAddressesListQueryParameters) SetStatusN(val []string) {
	m.StatusN = val
}

// GetTag returns the Tag property
func (m IpamIpAddressesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamIpAddressesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamIpAddressesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamIpAddressesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTenant returns the Tenant property
func (m IpamIpAddressesListQueryParameters) GetTenant() []string {
	return m.Tenant
}

// SetTenant sets the Tenant property
func (m *IpamIpAddressesListQueryParameters) SetTenant(val []string) {
	m.Tenant = val
}

// GetTenantN returns the TenantN property
func (m IpamIpAddressesListQueryParameters) GetTenantN() []string {
	return m.TenantN
}

// SetTenantN sets the TenantN property
func (m *IpamIpAddressesListQueryParameters) SetTenantN(val []string) {
	m.TenantN = val
}

// GetTenantGroup returns the TenantGroup property
func (m IpamIpAddressesListQueryParameters) GetTenantGroup() []int32 {
	return m.TenantGroup
}

// SetTenantGroup sets the TenantGroup property
func (m *IpamIpAddressesListQueryParameters) SetTenantGroup(val []int32) {
	m.TenantGroup = val
}

// GetTenantGroupN returns the TenantGroupN property
func (m IpamIpAddressesListQueryParameters) GetTenantGroupN() []int32 {
	return m.TenantGroupN
}

// SetTenantGroupN sets the TenantGroupN property
func (m *IpamIpAddressesListQueryParameters) SetTenantGroupN(val []int32) {
	m.TenantGroupN = val
}

// GetTenantGroupId returns the TenantGroupId property
func (m IpamIpAddressesListQueryParameters) GetTenantGroupId() []int32 {
	return m.TenantGroupId
}

// SetTenantGroupId sets the TenantGroupId property
func (m *IpamIpAddressesListQueryParameters) SetTenantGroupId(val []int32) {
	m.TenantGroupId = val
}

// GetTenantGroupIdN returns the TenantGroupIdN property
func (m IpamIpAddressesListQueryParameters) GetTenantGroupIdN() []int32 {
	return m.TenantGroupIdN
}

// SetTenantGroupIdN sets the TenantGroupIdN property
func (m *IpamIpAddressesListQueryParameters) SetTenantGroupIdN(val []int32) {
	m.TenantGroupIdN = val
}

// GetTenantId returns the TenantId property
func (m IpamIpAddressesListQueryParameters) GetTenantId() []*int32 {
	return m.TenantId
}

// SetTenantId sets the TenantId property
func (m *IpamIpAddressesListQueryParameters) SetTenantId(val []*int32) {
	m.TenantId = val
}

// GetTenantIdN returns the TenantIdN property
func (m IpamIpAddressesListQueryParameters) GetTenantIdN() []*int32 {
	return m.TenantIdN
}

// SetTenantIdN sets the TenantIdN property
func (m *IpamIpAddressesListQueryParameters) SetTenantIdN(val []*int32) {
	m.TenantIdN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamIpAddressesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamIpAddressesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m IpamIpAddressesListQueryParameters) GetVirtualMachine() []string {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *IpamIpAddressesListQueryParameters) SetVirtualMachine(val []string) {
	m.VirtualMachine = val
}

// GetVirtualMachineId returns the VirtualMachineId property
func (m IpamIpAddressesListQueryParameters) GetVirtualMachineId() []int32 {
	return m.VirtualMachineId
}

// SetVirtualMachineId sets the VirtualMachineId property
func (m *IpamIpAddressesListQueryParameters) SetVirtualMachineId(val []int32) {
	m.VirtualMachineId = val
}

// GetVminterface returns the Vminterface property
func (m IpamIpAddressesListQueryParameters) GetVminterface() []string {
	return m.Vminterface
}

// SetVminterface sets the Vminterface property
func (m *IpamIpAddressesListQueryParameters) SetVminterface(val []string) {
	m.Vminterface = val
}

// GetVminterfaceN returns the VminterfaceN property
func (m IpamIpAddressesListQueryParameters) GetVminterfaceN() []string {
	return m.VminterfaceN
}

// SetVminterfaceN sets the VminterfaceN property
func (m *IpamIpAddressesListQueryParameters) SetVminterfaceN(val []string) {
	m.VminterfaceN = val
}

// GetVminterfaceId returns the VminterfaceId property
func (m IpamIpAddressesListQueryParameters) GetVminterfaceId() []int32 {
	return m.VminterfaceId
}

// SetVminterfaceId sets the VminterfaceId property
func (m *IpamIpAddressesListQueryParameters) SetVminterfaceId(val []int32) {
	m.VminterfaceId = val
}

// GetVminterfaceIdN returns the VminterfaceIdN property
func (m IpamIpAddressesListQueryParameters) GetVminterfaceIdN() []int32 {
	return m.VminterfaceIdN
}

// SetVminterfaceIdN sets the VminterfaceIdN property
func (m *IpamIpAddressesListQueryParameters) SetVminterfaceIdN(val []int32) {
	m.VminterfaceIdN = val
}

// GetVrf returns the Vrf property
func (m IpamIpAddressesListQueryParameters) GetVrf() []*string {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *IpamIpAddressesListQueryParameters) SetVrf(val []*string) {
	m.Vrf = val
}

// GetVrfN returns the VrfN property
func (m IpamIpAddressesListQueryParameters) GetVrfN() []*string {
	return m.VrfN
}

// SetVrfN sets the VrfN property
func (m *IpamIpAddressesListQueryParameters) SetVrfN(val []*string) {
	m.VrfN = val
}

// GetVrfId returns the VrfId property
func (m IpamIpAddressesListQueryParameters) GetVrfId() []*int32 {
	return m.VrfId
}

// SetVrfId sets the VrfId property
func (m *IpamIpAddressesListQueryParameters) SetVrfId(val []*int32) {
	m.VrfId = val
}

// GetVrfIdN returns the VrfIdN property
func (m IpamIpAddressesListQueryParameters) GetVrfIdN() []*int32 {
	return m.VrfIdN
}

// SetVrfIdN sets the VrfIdN property
func (m *IpamIpAddressesListQueryParameters) SetVrfIdN(val []*int32) {
	m.VrfIdN = val
}
