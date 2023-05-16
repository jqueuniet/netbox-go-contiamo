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

// IpamL2vpnTerminationsListQueryParameters is an object.
type IpamL2vpnTerminationsListQueryParameters struct {
	// AssignedObjectType:
	AssignedObjectType string `json:"assigned_object_type,omitempty" mapstructure:"assigned_object_type,omitempty"`
	// AssignedObjectTypeN:
	AssignedObjectTypeN string `json:"assigned_object_type__n,omitempty" mapstructure:"assigned_object_type__n,omitempty"`
	// AssignedObjectTypeId:
	AssignedObjectTypeId int32 `json:"assigned_object_type_id,omitempty" mapstructure:"assigned_object_type_id,omitempty"`
	// AssignedObjectTypeIdN:
	AssignedObjectTypeIdN int32 `json:"assigned_object_type_id__n,omitempty" mapstructure:"assigned_object_type_id__n,omitempty"`
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
	// Device: Device (name)
	Device []*string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceN: Device (name)
	DeviceN []*string `json:"device__n,omitempty" mapstructure:"device__n,omitempty"`
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
	// Interface: Interface (name)
	Interface []string `json:"interface,omitempty" mapstructure:"interface,omitempty"`
	// InterfaceN: Interface (name)
	InterfaceN []string `json:"interface__n,omitempty" mapstructure:"interface__n,omitempty"`
	// InterfaceId: Interface (ID)
	InterfaceId []int32 `json:"interface_id,omitempty" mapstructure:"interface_id,omitempty"`
	// InterfaceIdN: Interface (ID)
	InterfaceIdN []int32 `json:"interface_id__n,omitempty" mapstructure:"interface_id__n,omitempty"`
	// L2vpn: L2VPN (slug)
	L2vpn []string `json:"l2vpn,omitempty" mapstructure:"l2vpn,omitempty"`
	// L2vpnN: L2VPN (slug)
	L2vpnN []string `json:"l2vpn__n,omitempty" mapstructure:"l2vpn__n,omitempty"`
	// L2vpnId: L2VPN (ID)
	L2vpnId []int32 `json:"l2vpn_id,omitempty" mapstructure:"l2vpn_id,omitempty"`
	// L2vpnIdN: L2VPN (ID)
	L2vpnIdN []int32 `json:"l2vpn_id__n,omitempty" mapstructure:"l2vpn_id__n,omitempty"`
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
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Region:
	Region []string `json:"region,omitempty" mapstructure:"region,omitempty"`
	// RegionId:
	RegionId []int32 `json:"region_id,omitempty" mapstructure:"region_id,omitempty"`
	// Site:
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteId:
	SiteId []int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
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
	VirtualMachineId []int32 `json:"virtual_machine_id,omitempty" mapstructure:"virtual_machine_id,omitempty"`
	// VirtualMachineIdN: Virtual machine (ID)
	VirtualMachineIdN []int32 `json:"virtual_machine_id__n,omitempty" mapstructure:"virtual_machine_id__n,omitempty"`
	// Vlan: VLAN (name)
	Vlan []string `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
	// VlanN: VLAN (name)
	VlanN []string `json:"vlan__n,omitempty" mapstructure:"vlan__n,omitempty"`
	// VlanId: VLAN (ID)
	VlanId []int32 `json:"vlan_id,omitempty" mapstructure:"vlan_id,omitempty"`
	// VlanIdN: VLAN (ID)
	VlanIdN []int32 `json:"vlan_id__n,omitempty" mapstructure:"vlan_id__n,omitempty"`
	// VlanVid: VLAN number (1-4094)
	VlanVid int32 `json:"vlan_vid,omitempty" mapstructure:"vlan_vid,omitempty"`
	// VlanVidGt: VLAN number (1-4094)
	VlanVidGt int32 `json:"vlan_vid__gt,omitempty" mapstructure:"vlan_vid__gt,omitempty"`
	// VlanVidGte: VLAN number (1-4094)
	VlanVidGte int32 `json:"vlan_vid__gte,omitempty" mapstructure:"vlan_vid__gte,omitempty"`
	// VlanVidLt: VLAN number (1-4094)
	VlanVidLt int32 `json:"vlan_vid__lt,omitempty" mapstructure:"vlan_vid__lt,omitempty"`
	// VlanVidLte: VLAN number (1-4094)
	VlanVidLte int32 `json:"vlan_vid__lte,omitempty" mapstructure:"vlan_vid__lte,omitempty"`
	// VlanVidN: VLAN number (1-4094)
	VlanVidN int32 `json:"vlan_vid__n,omitempty" mapstructure:"vlan_vid__n,omitempty"`
	// Vminterface: VM interface (name)
	Vminterface []string `json:"vminterface,omitempty" mapstructure:"vminterface,omitempty"`
	// VminterfaceN: VM interface (name)
	VminterfaceN []string `json:"vminterface__n,omitempty" mapstructure:"vminterface__n,omitempty"`
	// VminterfaceId: VM Interface (ID)
	VminterfaceId []int32 `json:"vminterface_id,omitempty" mapstructure:"vminterface_id,omitempty"`
	// VminterfaceIdN: VM Interface (ID)
	VminterfaceIdN []int32 `json:"vminterface_id__n,omitempty" mapstructure:"vminterface_id__n,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamL2vpnTerminationsListQueryParameters) Validate() error {
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
		"l2vpn": validation.Validate(
			m.L2vpn,
		),
		"l2vpnN": validation.Validate(
			m.L2vpnN,
		),
		"l2vpnId": validation.Validate(
			m.L2vpnId,
		),
		"l2vpnIdN": validation.Validate(
			m.L2vpnIdN,
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
		"region": validation.Validate(
			m.Region,
		),
		"regionId": validation.Validate(
			m.RegionId,
		),
		"site": validation.Validate(
			m.Site,
		),
		"siteId": validation.Validate(
			m.SiteId,
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
		"vlan": validation.Validate(
			m.Vlan,
		),
		"vlanN": validation.Validate(
			m.VlanN,
		),
		"vlanId": validation.Validate(
			m.VlanId,
		),
		"vlanIdN": validation.Validate(
			m.VlanIdN,
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
	}.Filter()
}

// GetAssignedObjectType returns the AssignedObjectType property
func (m IpamL2vpnTerminationsListQueryParameters) GetAssignedObjectType() string {
	return m.AssignedObjectType
}

// SetAssignedObjectType sets the AssignedObjectType property
func (m *IpamL2vpnTerminationsListQueryParameters) SetAssignedObjectType(val string) {
	m.AssignedObjectType = val
}

// GetAssignedObjectTypeN returns the AssignedObjectTypeN property
func (m IpamL2vpnTerminationsListQueryParameters) GetAssignedObjectTypeN() string {
	return m.AssignedObjectTypeN
}

// SetAssignedObjectTypeN sets the AssignedObjectTypeN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetAssignedObjectTypeN(val string) {
	m.AssignedObjectTypeN = val
}

// GetAssignedObjectTypeId returns the AssignedObjectTypeId property
func (m IpamL2vpnTerminationsListQueryParameters) GetAssignedObjectTypeId() int32 {
	return m.AssignedObjectTypeId
}

// SetAssignedObjectTypeId sets the AssignedObjectTypeId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetAssignedObjectTypeId(val int32) {
	m.AssignedObjectTypeId = val
}

// GetAssignedObjectTypeIdN returns the AssignedObjectTypeIdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetAssignedObjectTypeIdN() int32 {
	return m.AssignedObjectTypeIdN
}

// SetAssignedObjectTypeIdN sets the AssignedObjectTypeIdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetAssignedObjectTypeIdN(val int32) {
	m.AssignedObjectTypeIdN = val
}

// GetCreated returns the Created property
func (m IpamL2vpnTerminationsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamL2vpnTerminationsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamL2vpnTerminationsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamL2vpnTerminationsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamL2vpnTerminationsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamL2vpnTerminationsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamL2vpnTerminationsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamL2vpnTerminationsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamL2vpnTerminationsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevice returns the Device property
func (m IpamL2vpnTerminationsListQueryParameters) GetDevice() []*string {
	return m.Device
}

// SetDevice sets the Device property
func (m *IpamL2vpnTerminationsListQueryParameters) SetDevice(val []*string) {
	m.Device = val
}

// GetDeviceN returns the DeviceN property
func (m IpamL2vpnTerminationsListQueryParameters) GetDeviceN() []*string {
	return m.DeviceN
}

// SetDeviceN sets the DeviceN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetDeviceN(val []*string) {
	m.DeviceN = val
}

// GetDeviceId returns the DeviceId property
func (m IpamL2vpnTerminationsListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetDeviceIdN returns the DeviceIdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetDeviceIdN() []int32 {
	return m.DeviceIdN
}

// SetDeviceIdN sets the DeviceIdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetDeviceIdN(val []int32) {
	m.DeviceIdN = val
}

// GetId returns the Id property
func (m IpamL2vpnTerminationsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamL2vpnTerminationsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamL2vpnTerminationsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamL2vpnTerminationsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamL2vpnTerminationsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamL2vpnTerminationsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterface returns the Interface property
func (m IpamL2vpnTerminationsListQueryParameters) GetInterface() []string {
	return m.Interface
}

// SetInterface sets the Interface property
func (m *IpamL2vpnTerminationsListQueryParameters) SetInterface(val []string) {
	m.Interface = val
}

// GetInterfaceN returns the InterfaceN property
func (m IpamL2vpnTerminationsListQueryParameters) GetInterfaceN() []string {
	return m.InterfaceN
}

// SetInterfaceN sets the InterfaceN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetInterfaceN(val []string) {
	m.InterfaceN = val
}

// GetInterfaceId returns the InterfaceId property
func (m IpamL2vpnTerminationsListQueryParameters) GetInterfaceId() []int32 {
	return m.InterfaceId
}

// SetInterfaceId sets the InterfaceId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetInterfaceId(val []int32) {
	m.InterfaceId = val
}

// GetInterfaceIdN returns the InterfaceIdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetInterfaceIdN() []int32 {
	return m.InterfaceIdN
}

// SetInterfaceIdN sets the InterfaceIdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetInterfaceIdN(val []int32) {
	m.InterfaceIdN = val
}

// GetL2vpn returns the L2vpn property
func (m IpamL2vpnTerminationsListQueryParameters) GetL2vpn() []string {
	return m.L2vpn
}

// SetL2vpn sets the L2vpn property
func (m *IpamL2vpnTerminationsListQueryParameters) SetL2vpn(val []string) {
	m.L2vpn = val
}

// GetL2vpnN returns the L2vpnN property
func (m IpamL2vpnTerminationsListQueryParameters) GetL2vpnN() []string {
	return m.L2vpnN
}

// SetL2vpnN sets the L2vpnN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetL2vpnN(val []string) {
	m.L2vpnN = val
}

// GetL2vpnId returns the L2vpnId property
func (m IpamL2vpnTerminationsListQueryParameters) GetL2vpnId() []int32 {
	return m.L2vpnId
}

// SetL2vpnId sets the L2vpnId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetL2vpnId(val []int32) {
	m.L2vpnId = val
}

// GetL2vpnIdN returns the L2vpnIdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetL2vpnIdN() []int32 {
	return m.L2vpnIdN
}

// SetL2vpnIdN sets the L2vpnIdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetL2vpnIdN(val []int32) {
	m.L2vpnIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamL2vpnTerminationsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamL2vpnTerminationsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamL2vpnTerminationsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamL2vpnTerminationsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamL2vpnTerminationsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamL2vpnTerminationsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamL2vpnTerminationsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamL2vpnTerminationsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamL2vpnTerminationsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m IpamL2vpnTerminationsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamL2vpnTerminationsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamL2vpnTerminationsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamL2vpnTerminationsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m IpamL2vpnTerminationsListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *IpamL2vpnTerminationsListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRegion returns the Region property
func (m IpamL2vpnTerminationsListQueryParameters) GetRegion() []string {
	return m.Region
}

// SetRegion sets the Region property
func (m *IpamL2vpnTerminationsListQueryParameters) SetRegion(val []string) {
	m.Region = val
}

// GetRegionId returns the RegionId property
func (m IpamL2vpnTerminationsListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetSite returns the Site property
func (m IpamL2vpnTerminationsListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *IpamL2vpnTerminationsListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteId returns the SiteId property
func (m IpamL2vpnTerminationsListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetTag returns the Tag property
func (m IpamL2vpnTerminationsListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *IpamL2vpnTerminationsListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m IpamL2vpnTerminationsListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamL2vpnTerminationsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamL2vpnTerminationsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m IpamL2vpnTerminationsListQueryParameters) GetVirtualMachine() []string {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVirtualMachine(val []string) {
	m.VirtualMachine = val
}

// GetVirtualMachineN returns the VirtualMachineN property
func (m IpamL2vpnTerminationsListQueryParameters) GetVirtualMachineN() []string {
	return m.VirtualMachineN
}

// SetVirtualMachineN sets the VirtualMachineN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVirtualMachineN(val []string) {
	m.VirtualMachineN = val
}

// GetVirtualMachineId returns the VirtualMachineId property
func (m IpamL2vpnTerminationsListQueryParameters) GetVirtualMachineId() []int32 {
	return m.VirtualMachineId
}

// SetVirtualMachineId sets the VirtualMachineId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVirtualMachineId(val []int32) {
	m.VirtualMachineId = val
}

// GetVirtualMachineIdN returns the VirtualMachineIdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetVirtualMachineIdN() []int32 {
	return m.VirtualMachineIdN
}

// SetVirtualMachineIdN sets the VirtualMachineIdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVirtualMachineIdN(val []int32) {
	m.VirtualMachineIdN = val
}

// GetVlan returns the Vlan property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlan() []string {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlan(val []string) {
	m.Vlan = val
}

// GetVlanN returns the VlanN property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanN() []string {
	return m.VlanN
}

// SetVlanN sets the VlanN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanN(val []string) {
	m.VlanN = val
}

// GetVlanId returns the VlanId property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanId() []int32 {
	return m.VlanId
}

// SetVlanId sets the VlanId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanId(val []int32) {
	m.VlanId = val
}

// GetVlanIdN returns the VlanIdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanIdN() []int32 {
	return m.VlanIdN
}

// SetVlanIdN sets the VlanIdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanIdN(val []int32) {
	m.VlanIdN = val
}

// GetVlanVid returns the VlanVid property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanVid() int32 {
	return m.VlanVid
}

// SetVlanVid sets the VlanVid property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanVid(val int32) {
	m.VlanVid = val
}

// GetVlanVidGt returns the VlanVidGt property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanVidGt() int32 {
	return m.VlanVidGt
}

// SetVlanVidGt sets the VlanVidGt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanVidGt(val int32) {
	m.VlanVidGt = val
}

// GetVlanVidGte returns the VlanVidGte property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanVidGte() int32 {
	return m.VlanVidGte
}

// SetVlanVidGte sets the VlanVidGte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanVidGte(val int32) {
	m.VlanVidGte = val
}

// GetVlanVidLt returns the VlanVidLt property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanVidLt() int32 {
	return m.VlanVidLt
}

// SetVlanVidLt sets the VlanVidLt property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanVidLt(val int32) {
	m.VlanVidLt = val
}

// GetVlanVidLte returns the VlanVidLte property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanVidLte() int32 {
	return m.VlanVidLte
}

// SetVlanVidLte sets the VlanVidLte property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanVidLte(val int32) {
	m.VlanVidLte = val
}

// GetVlanVidN returns the VlanVidN property
func (m IpamL2vpnTerminationsListQueryParameters) GetVlanVidN() int32 {
	return m.VlanVidN
}

// SetVlanVidN sets the VlanVidN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVlanVidN(val int32) {
	m.VlanVidN = val
}

// GetVminterface returns the Vminterface property
func (m IpamL2vpnTerminationsListQueryParameters) GetVminterface() []string {
	return m.Vminterface
}

// SetVminterface sets the Vminterface property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVminterface(val []string) {
	m.Vminterface = val
}

// GetVminterfaceN returns the VminterfaceN property
func (m IpamL2vpnTerminationsListQueryParameters) GetVminterfaceN() []string {
	return m.VminterfaceN
}

// SetVminterfaceN sets the VminterfaceN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVminterfaceN(val []string) {
	m.VminterfaceN = val
}

// GetVminterfaceId returns the VminterfaceId property
func (m IpamL2vpnTerminationsListQueryParameters) GetVminterfaceId() []int32 {
	return m.VminterfaceId
}

// SetVminterfaceId sets the VminterfaceId property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVminterfaceId(val []int32) {
	m.VminterfaceId = val
}

// GetVminterfaceIdN returns the VminterfaceIdN property
func (m IpamL2vpnTerminationsListQueryParameters) GetVminterfaceIdN() []int32 {
	return m.VminterfaceIdN
}

// SetVminterfaceIdN sets the VminterfaceIdN property
func (m *IpamL2vpnTerminationsListQueryParameters) SetVminterfaceIdN(val []int32) {
	m.VminterfaceIdN = val
}
