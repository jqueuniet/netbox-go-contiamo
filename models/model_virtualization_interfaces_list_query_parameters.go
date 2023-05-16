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

// VirtualizationInterfacesListQueryParameters is an object.
type VirtualizationInterfacesListQueryParameters struct {
	// BridgeId: Bridged interface (ID)
	BridgeId []int32 `json:"bridge_id,omitempty" mapstructure:"bridge_id,omitempty"`
	// BridgeIdN: Bridged interface (ID)
	BridgeIdN []int32 `json:"bridge_id__n,omitempty" mapstructure:"bridge_id__n,omitempty"`
	// Cluster: Cluster
	Cluster []string `json:"cluster,omitempty" mapstructure:"cluster,omitempty"`
	// ClusterN: Cluster
	ClusterN []string `json:"cluster__n,omitempty" mapstructure:"cluster__n,omitempty"`
	// ClusterId: Cluster (ID)
	ClusterId []int32 `json:"cluster_id,omitempty" mapstructure:"cluster_id,omitempty"`
	// ClusterIdN: Cluster (ID)
	ClusterIdN []int32 `json:"cluster_id__n,omitempty" mapstructure:"cluster_id__n,omitempty"`
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
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
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
	// L2vpn: L2VPN
	L2vpn []*int64 `json:"l2vpn,omitempty" mapstructure:"l2vpn,omitempty"`
	// L2vpnN: L2VPN
	L2vpnN []*int64 `json:"l2vpn__n,omitempty" mapstructure:"l2vpn__n,omitempty"`
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
	// MacAddress:
	MacAddress []string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// MacAddressIc:
	MacAddressIc []string `json:"mac_address__ic,omitempty" mapstructure:"mac_address__ic,omitempty"`
	// MacAddressIe:
	MacAddressIe []string `json:"mac_address__ie,omitempty" mapstructure:"mac_address__ie,omitempty"`
	// MacAddressIew:
	MacAddressIew []string `json:"mac_address__iew,omitempty" mapstructure:"mac_address__iew,omitempty"`
	// MacAddressIsw:
	MacAddressIsw []string `json:"mac_address__isw,omitempty" mapstructure:"mac_address__isw,omitempty"`
	// MacAddressN:
	MacAddressN []string `json:"mac_address__n,omitempty" mapstructure:"mac_address__n,omitempty"`
	// MacAddressNic:
	MacAddressNic []string `json:"mac_address__nic,omitempty" mapstructure:"mac_address__nic,omitempty"`
	// MacAddressNie:
	MacAddressNie []string `json:"mac_address__nie,omitempty" mapstructure:"mac_address__nie,omitempty"`
	// MacAddressNiew:
	MacAddressNiew []string `json:"mac_address__niew,omitempty" mapstructure:"mac_address__niew,omitempty"`
	// MacAddressNisw:
	MacAddressNisw []string `json:"mac_address__nisw,omitempty" mapstructure:"mac_address__nisw,omitempty"`
	// Mtu:
	Mtu []int32 `json:"mtu,omitempty" mapstructure:"mtu,omitempty"`
	// MtuGt:
	MtuGt []int32 `json:"mtu__gt,omitempty" mapstructure:"mtu__gt,omitempty"`
	// MtuGte:
	MtuGte []int32 `json:"mtu__gte,omitempty" mapstructure:"mtu__gte,omitempty"`
	// MtuLt:
	MtuLt []int32 `json:"mtu__lt,omitempty" mapstructure:"mtu__lt,omitempty"`
	// MtuLte:
	MtuLte []int32 `json:"mtu__lte,omitempty" mapstructure:"mtu__lte,omitempty"`
	// MtuN:
	MtuN []int32 `json:"mtu__n,omitempty" mapstructure:"mtu__n,omitempty"`
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
	// ParentId: Parent interface (ID)
	ParentId []int32 `json:"parent_id,omitempty" mapstructure:"parent_id,omitempty"`
	// ParentIdN: Parent interface (ID)
	ParentIdN []int32 `json:"parent_id__n,omitempty" mapstructure:"parent_id__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// VirtualMachine: Virtual machine
	VirtualMachine []string `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
	// VirtualMachineN: Virtual machine
	VirtualMachineN []string `json:"virtual_machine__n,omitempty" mapstructure:"virtual_machine__n,omitempty"`
	// VirtualMachineId: Virtual machine (ID)
	VirtualMachineId []int32 `json:"virtual_machine_id,omitempty" mapstructure:"virtual_machine_id,omitempty"`
	// VirtualMachineIdN: Virtual machine (ID)
	VirtualMachineIdN []int32 `json:"virtual_machine_id__n,omitempty" mapstructure:"virtual_machine_id__n,omitempty"`
	// Vlan: Assigned VID
	Vlan string `json:"vlan,omitempty" mapstructure:"vlan,omitempty"`
	// VlanId: Assigned VLAN
	VlanId string `json:"vlan_id,omitempty" mapstructure:"vlan_id,omitempty"`
	// Vrf: VRF (RD)
	Vrf []*string `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
	// VrfN: VRF (RD)
	VrfN []*string `json:"vrf__n,omitempty" mapstructure:"vrf__n,omitempty"`
	// VrfId: VRF
	VrfId []int32 `json:"vrf_id,omitempty" mapstructure:"vrf_id,omitempty"`
	// VrfIdN: VRF
	VrfIdN []int32 `json:"vrf_id__n,omitempty" mapstructure:"vrf_id__n,omitempty"`
}

// Validate implements basic validation for this model
func (m VirtualizationInterfacesListQueryParameters) Validate() error {
	return validation.Errors{
		"bridgeId": validation.Validate(
			m.BridgeId,
		),
		"bridgeIdN": validation.Validate(
			m.BridgeIdN,
		),
		"cluster": validation.Validate(
			m.Cluster,
		),
		"clusterN": validation.Validate(
			m.ClusterN,
		),
		"clusterId": validation.Validate(
			m.ClusterId,
		),
		"clusterIdN": validation.Validate(
			m.ClusterIdN,
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
		"macAddress": validation.Validate(
			m.MacAddress,
		),
		"macAddressIc": validation.Validate(
			m.MacAddressIc,
		),
		"macAddressIe": validation.Validate(
			m.MacAddressIe,
		),
		"macAddressIew": validation.Validate(
			m.MacAddressIew,
		),
		"macAddressIsw": validation.Validate(
			m.MacAddressIsw,
		),
		"macAddressN": validation.Validate(
			m.MacAddressN,
		),
		"macAddressNic": validation.Validate(
			m.MacAddressNic,
		),
		"macAddressNie": validation.Validate(
			m.MacAddressNie,
		),
		"macAddressNiew": validation.Validate(
			m.MacAddressNiew,
		),
		"macAddressNisw": validation.Validate(
			m.MacAddressNisw,
		),
		"mtu": validation.Validate(
			m.Mtu,
		),
		"mtuGt": validation.Validate(
			m.MtuGt,
		),
		"mtuGte": validation.Validate(
			m.MtuGte,
		),
		"mtuLt": validation.Validate(
			m.MtuLt,
		),
		"mtuLte": validation.Validate(
			m.MtuLte,
		),
		"mtuN": validation.Validate(
			m.MtuN,
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

// GetBridgeId returns the BridgeId property
func (m VirtualizationInterfacesListQueryParameters) GetBridgeId() []int32 {
	return m.BridgeId
}

// SetBridgeId sets the BridgeId property
func (m *VirtualizationInterfacesListQueryParameters) SetBridgeId(val []int32) {
	m.BridgeId = val
}

// GetBridgeIdN returns the BridgeIdN property
func (m VirtualizationInterfacesListQueryParameters) GetBridgeIdN() []int32 {
	return m.BridgeIdN
}

// SetBridgeIdN sets the BridgeIdN property
func (m *VirtualizationInterfacesListQueryParameters) SetBridgeIdN(val []int32) {
	m.BridgeIdN = val
}

// GetCluster returns the Cluster property
func (m VirtualizationInterfacesListQueryParameters) GetCluster() []string {
	return m.Cluster
}

// SetCluster sets the Cluster property
func (m *VirtualizationInterfacesListQueryParameters) SetCluster(val []string) {
	m.Cluster = val
}

// GetClusterN returns the ClusterN property
func (m VirtualizationInterfacesListQueryParameters) GetClusterN() []string {
	return m.ClusterN
}

// SetClusterN sets the ClusterN property
func (m *VirtualizationInterfacesListQueryParameters) SetClusterN(val []string) {
	m.ClusterN = val
}

// GetClusterId returns the ClusterId property
func (m VirtualizationInterfacesListQueryParameters) GetClusterId() []int32 {
	return m.ClusterId
}

// SetClusterId sets the ClusterId property
func (m *VirtualizationInterfacesListQueryParameters) SetClusterId(val []int32) {
	m.ClusterId = val
}

// GetClusterIdN returns the ClusterIdN property
func (m VirtualizationInterfacesListQueryParameters) GetClusterIdN() []int32 {
	return m.ClusterIdN
}

// SetClusterIdN sets the ClusterIdN property
func (m *VirtualizationInterfacesListQueryParameters) SetClusterIdN(val []int32) {
	m.ClusterIdN = val
}

// GetCreated returns the Created property
func (m VirtualizationInterfacesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VirtualizationInterfacesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m VirtualizationInterfacesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *VirtualizationInterfacesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m VirtualizationInterfacesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *VirtualizationInterfacesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m VirtualizationInterfacesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *VirtualizationInterfacesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m VirtualizationInterfacesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *VirtualizationInterfacesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m VirtualizationInterfacesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *VirtualizationInterfacesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m VirtualizationInterfacesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *VirtualizationInterfacesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m VirtualizationInterfacesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VirtualizationInterfacesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m VirtualizationInterfacesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *VirtualizationInterfacesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetEnabled returns the Enabled property
func (m VirtualizationInterfacesListQueryParameters) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *VirtualizationInterfacesListQueryParameters) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m VirtualizationInterfacesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VirtualizationInterfacesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m VirtualizationInterfacesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *VirtualizationInterfacesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m VirtualizationInterfacesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *VirtualizationInterfacesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m VirtualizationInterfacesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *VirtualizationInterfacesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m VirtualizationInterfacesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *VirtualizationInterfacesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m VirtualizationInterfacesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *VirtualizationInterfacesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetL2vpn returns the L2vpn property
func (m VirtualizationInterfacesListQueryParameters) GetL2vpn() []*int64 {
	return m.L2vpn
}

// SetL2vpn sets the L2vpn property
func (m *VirtualizationInterfacesListQueryParameters) SetL2vpn(val []*int64) {
	m.L2vpn = val
}

// GetL2vpnN returns the L2vpnN property
func (m VirtualizationInterfacesListQueryParameters) GetL2vpnN() []*int64 {
	return m.L2vpnN
}

// SetL2vpnN sets the L2vpnN property
func (m *VirtualizationInterfacesListQueryParameters) SetL2vpnN(val []*int64) {
	m.L2vpnN = val
}

// GetL2vpnId returns the L2vpnId property
func (m VirtualizationInterfacesListQueryParameters) GetL2vpnId() []int32 {
	return m.L2vpnId
}

// SetL2vpnId sets the L2vpnId property
func (m *VirtualizationInterfacesListQueryParameters) SetL2vpnId(val []int32) {
	m.L2vpnId = val
}

// GetL2vpnIdN returns the L2vpnIdN property
func (m VirtualizationInterfacesListQueryParameters) GetL2vpnIdN() []int32 {
	return m.L2vpnIdN
}

// SetL2vpnIdN sets the L2vpnIdN property
func (m *VirtualizationInterfacesListQueryParameters) SetL2vpnIdN(val []int32) {
	m.L2vpnIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m VirtualizationInterfacesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VirtualizationInterfacesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m VirtualizationInterfacesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *VirtualizationInterfacesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m VirtualizationInterfacesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *VirtualizationInterfacesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m VirtualizationInterfacesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *VirtualizationInterfacesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m VirtualizationInterfacesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *VirtualizationInterfacesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m VirtualizationInterfacesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *VirtualizationInterfacesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m VirtualizationInterfacesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *VirtualizationInterfacesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetMacAddress returns the MacAddress property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddress() []string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddress(val []string) {
	m.MacAddress = val
}

// GetMacAddressIc returns the MacAddressIc property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressIc() []string {
	return m.MacAddressIc
}

// SetMacAddressIc sets the MacAddressIc property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressIc(val []string) {
	m.MacAddressIc = val
}

// GetMacAddressIe returns the MacAddressIe property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressIe() []string {
	return m.MacAddressIe
}

// SetMacAddressIe sets the MacAddressIe property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressIe(val []string) {
	m.MacAddressIe = val
}

// GetMacAddressIew returns the MacAddressIew property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressIew() []string {
	return m.MacAddressIew
}

// SetMacAddressIew sets the MacAddressIew property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressIew(val []string) {
	m.MacAddressIew = val
}

// GetMacAddressIsw returns the MacAddressIsw property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressIsw() []string {
	return m.MacAddressIsw
}

// SetMacAddressIsw sets the MacAddressIsw property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressIsw(val []string) {
	m.MacAddressIsw = val
}

// GetMacAddressN returns the MacAddressN property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressN() []string {
	return m.MacAddressN
}

// SetMacAddressN sets the MacAddressN property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressN(val []string) {
	m.MacAddressN = val
}

// GetMacAddressNic returns the MacAddressNic property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressNic() []string {
	return m.MacAddressNic
}

// SetMacAddressNic sets the MacAddressNic property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressNic(val []string) {
	m.MacAddressNic = val
}

// GetMacAddressNie returns the MacAddressNie property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressNie() []string {
	return m.MacAddressNie
}

// SetMacAddressNie sets the MacAddressNie property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressNie(val []string) {
	m.MacAddressNie = val
}

// GetMacAddressNiew returns the MacAddressNiew property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressNiew() []string {
	return m.MacAddressNiew
}

// SetMacAddressNiew sets the MacAddressNiew property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressNiew(val []string) {
	m.MacAddressNiew = val
}

// GetMacAddressNisw returns the MacAddressNisw property
func (m VirtualizationInterfacesListQueryParameters) GetMacAddressNisw() []string {
	return m.MacAddressNisw
}

// SetMacAddressNisw sets the MacAddressNisw property
func (m *VirtualizationInterfacesListQueryParameters) SetMacAddressNisw(val []string) {
	m.MacAddressNisw = val
}

// GetMtu returns the Mtu property
func (m VirtualizationInterfacesListQueryParameters) GetMtu() []int32 {
	return m.Mtu
}

// SetMtu sets the Mtu property
func (m *VirtualizationInterfacesListQueryParameters) SetMtu(val []int32) {
	m.Mtu = val
}

// GetMtuGt returns the MtuGt property
func (m VirtualizationInterfacesListQueryParameters) GetMtuGt() []int32 {
	return m.MtuGt
}

// SetMtuGt sets the MtuGt property
func (m *VirtualizationInterfacesListQueryParameters) SetMtuGt(val []int32) {
	m.MtuGt = val
}

// GetMtuGte returns the MtuGte property
func (m VirtualizationInterfacesListQueryParameters) GetMtuGte() []int32 {
	return m.MtuGte
}

// SetMtuGte sets the MtuGte property
func (m *VirtualizationInterfacesListQueryParameters) SetMtuGte(val []int32) {
	m.MtuGte = val
}

// GetMtuLt returns the MtuLt property
func (m VirtualizationInterfacesListQueryParameters) GetMtuLt() []int32 {
	return m.MtuLt
}

// SetMtuLt sets the MtuLt property
func (m *VirtualizationInterfacesListQueryParameters) SetMtuLt(val []int32) {
	m.MtuLt = val
}

// GetMtuLte returns the MtuLte property
func (m VirtualizationInterfacesListQueryParameters) GetMtuLte() []int32 {
	return m.MtuLte
}

// SetMtuLte sets the MtuLte property
func (m *VirtualizationInterfacesListQueryParameters) SetMtuLte(val []int32) {
	m.MtuLte = val
}

// GetMtuN returns the MtuN property
func (m VirtualizationInterfacesListQueryParameters) GetMtuN() []int32 {
	return m.MtuN
}

// SetMtuN sets the MtuN property
func (m *VirtualizationInterfacesListQueryParameters) SetMtuN(val []int32) {
	m.MtuN = val
}

// GetName returns the Name property
func (m VirtualizationInterfacesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *VirtualizationInterfacesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m VirtualizationInterfacesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *VirtualizationInterfacesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m VirtualizationInterfacesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *VirtualizationInterfacesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m VirtualizationInterfacesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *VirtualizationInterfacesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m VirtualizationInterfacesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *VirtualizationInterfacesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m VirtualizationInterfacesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *VirtualizationInterfacesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m VirtualizationInterfacesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *VirtualizationInterfacesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m VirtualizationInterfacesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *VirtualizationInterfacesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m VirtualizationInterfacesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *VirtualizationInterfacesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m VirtualizationInterfacesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *VirtualizationInterfacesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m VirtualizationInterfacesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *VirtualizationInterfacesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m VirtualizationInterfacesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *VirtualizationInterfacesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m VirtualizationInterfacesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *VirtualizationInterfacesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetParentId returns the ParentId property
func (m VirtualizationInterfacesListQueryParameters) GetParentId() []int32 {
	return m.ParentId
}

// SetParentId sets the ParentId property
func (m *VirtualizationInterfacesListQueryParameters) SetParentId(val []int32) {
	m.ParentId = val
}

// GetParentIdN returns the ParentIdN property
func (m VirtualizationInterfacesListQueryParameters) GetParentIdN() []int32 {
	return m.ParentIdN
}

// SetParentIdN sets the ParentIdN property
func (m *VirtualizationInterfacesListQueryParameters) SetParentIdN(val []int32) {
	m.ParentIdN = val
}

// GetQ returns the Q property
func (m VirtualizationInterfacesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *VirtualizationInterfacesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m VirtualizationInterfacesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *VirtualizationInterfacesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m VirtualizationInterfacesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *VirtualizationInterfacesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m VirtualizationInterfacesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *VirtualizationInterfacesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m VirtualizationInterfacesListQueryParameters) GetVirtualMachine() []string {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *VirtualizationInterfacesListQueryParameters) SetVirtualMachine(val []string) {
	m.VirtualMachine = val
}

// GetVirtualMachineN returns the VirtualMachineN property
func (m VirtualizationInterfacesListQueryParameters) GetVirtualMachineN() []string {
	return m.VirtualMachineN
}

// SetVirtualMachineN sets the VirtualMachineN property
func (m *VirtualizationInterfacesListQueryParameters) SetVirtualMachineN(val []string) {
	m.VirtualMachineN = val
}

// GetVirtualMachineId returns the VirtualMachineId property
func (m VirtualizationInterfacesListQueryParameters) GetVirtualMachineId() []int32 {
	return m.VirtualMachineId
}

// SetVirtualMachineId sets the VirtualMachineId property
func (m *VirtualizationInterfacesListQueryParameters) SetVirtualMachineId(val []int32) {
	m.VirtualMachineId = val
}

// GetVirtualMachineIdN returns the VirtualMachineIdN property
func (m VirtualizationInterfacesListQueryParameters) GetVirtualMachineIdN() []int32 {
	return m.VirtualMachineIdN
}

// SetVirtualMachineIdN sets the VirtualMachineIdN property
func (m *VirtualizationInterfacesListQueryParameters) SetVirtualMachineIdN(val []int32) {
	m.VirtualMachineIdN = val
}

// GetVlan returns the Vlan property
func (m VirtualizationInterfacesListQueryParameters) GetVlan() string {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *VirtualizationInterfacesListQueryParameters) SetVlan(val string) {
	m.Vlan = val
}

// GetVlanId returns the VlanId property
func (m VirtualizationInterfacesListQueryParameters) GetVlanId() string {
	return m.VlanId
}

// SetVlanId sets the VlanId property
func (m *VirtualizationInterfacesListQueryParameters) SetVlanId(val string) {
	m.VlanId = val
}

// GetVrf returns the Vrf property
func (m VirtualizationInterfacesListQueryParameters) GetVrf() []*string {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *VirtualizationInterfacesListQueryParameters) SetVrf(val []*string) {
	m.Vrf = val
}

// GetVrfN returns the VrfN property
func (m VirtualizationInterfacesListQueryParameters) GetVrfN() []*string {
	return m.VrfN
}

// SetVrfN sets the VrfN property
func (m *VirtualizationInterfacesListQueryParameters) SetVrfN(val []*string) {
	m.VrfN = val
}

// GetVrfId returns the VrfId property
func (m VirtualizationInterfacesListQueryParameters) GetVrfId() []int32 {
	return m.VrfId
}

// SetVrfId sets the VrfId property
func (m *VirtualizationInterfacesListQueryParameters) SetVrfId(val []int32) {
	m.VrfId = val
}

// GetVrfIdN returns the VrfIdN property
func (m VirtualizationInterfacesListQueryParameters) GetVrfIdN() []int32 {
	return m.VrfIdN
}

// SetVrfIdN sets the VrfIdN property
func (m *VirtualizationInterfacesListQueryParameters) SetVrfIdN(val []int32) {
	m.VrfIdN = val
}
