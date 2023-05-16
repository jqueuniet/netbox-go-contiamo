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

// DcimInterfacesListQueryParameters is an object.
type DcimInterfacesListQueryParameters struct {
	// BridgeId: Bridged interface (ID)
	BridgeId []int32 `json:"bridge_id,omitempty" mapstructure:"bridge_id,omitempty"`
	// BridgeIdN: Bridged interface (ID)
	BridgeIdN []int32 `json:"bridge_id__n,omitempty" mapstructure:"bridge_id__n,omitempty"`
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end,omitempty" mapstructure:"cable_end,omitempty"`
	// CableEndN: * `A` - A
	// * `B` - B
	CableEndN string `json:"cable_end__n,omitempty" mapstructure:"cable_end__n,omitempty"`
	// Cabled:
	Cabled bool `json:"cabled,omitempty" mapstructure:"cabled,omitempty"`
	// Connected:
	Connected bool `json:"connected,omitempty" mapstructure:"connected,omitempty"`
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
	// Duplex: * `half` - Half
	// * `full` - Full
	// * `auto` - Auto
	Duplex []*string `json:"duplex,omitempty" mapstructure:"duplex,omitempty"`
	// DuplexN: * `half` - Half
	// * `full` - Full
	// * `auto` - Auto
	DuplexN []*string `json:"duplex__n,omitempty" mapstructure:"duplex__n,omitempty"`
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
	// Kind: Kind of interface
	Kind string `json:"kind,omitempty" mapstructure:"kind,omitempty"`
	// L2vpn: L2VPN
	L2vpn []*int64 `json:"l2vpn,omitempty" mapstructure:"l2vpn,omitempty"`
	// L2vpnN: L2VPN
	L2vpnN []*int64 `json:"l2vpn__n,omitempty" mapstructure:"l2vpn__n,omitempty"`
	// L2vpnId: L2VPN (ID)
	L2vpnId []int32 `json:"l2vpn_id,omitempty" mapstructure:"l2vpn_id,omitempty"`
	// L2vpnIdN: L2VPN (ID)
	L2vpnIdN []int32 `json:"l2vpn_id__n,omitempty" mapstructure:"l2vpn_id__n,omitempty"`
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
	// LagId: LAG interface (ID)
	LagId []int32 `json:"lag_id,omitempty" mapstructure:"lag_id,omitempty"`
	// LagIdN: LAG interface (ID)
	LagIdN []int32 `json:"lag_id__n,omitempty" mapstructure:"lag_id__n,omitempty"`
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
	// Location: Location (slug)
	Location []string `json:"location,omitempty" mapstructure:"location,omitempty"`
	// LocationN: Location (slug)
	LocationN []string `json:"location__n,omitempty" mapstructure:"location__n,omitempty"`
	// LocationId: Location (ID)
	LocationId []int32 `json:"location_id,omitempty" mapstructure:"location_id,omitempty"`
	// LocationIdN: Location (ID)
	LocationIdN []int32 `json:"location_id__n,omitempty" mapstructure:"location_id__n,omitempty"`
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
	// MgmtOnly:
	MgmtOnly bool `json:"mgmt_only,omitempty" mapstructure:"mgmt_only,omitempty"`
	// Mode: IEEE 802.1Q tagging strategy
	Mode string `json:"mode,omitempty" mapstructure:"mode,omitempty"`
	// ModeN: IEEE 802.1Q tagging strategy
	ModeN string `json:"mode__n,omitempty" mapstructure:"mode__n,omitempty"`
	// ModuleId: Module (ID)
	ModuleId []*int32 `json:"module_id,omitempty" mapstructure:"module_id,omitempty"`
	// ModuleIdN: Module (ID)
	ModuleIdN []*int32 `json:"module_id__n,omitempty" mapstructure:"module_id__n,omitempty"`
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
	// Occupied:
	Occupied bool `json:"occupied,omitempty" mapstructure:"occupied,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// ParentId: Parent interface (ID)
	ParentId []int32 `json:"parent_id,omitempty" mapstructure:"parent_id,omitempty"`
	// ParentIdN: Parent interface (ID)
	ParentIdN []int32 `json:"parent_id__n,omitempty" mapstructure:"parent_id__n,omitempty"`
	// PoeMode: * `pd` - PD
	// * `pse` - PSE
	PoeMode []string `json:"poe_mode,omitempty" mapstructure:"poe_mode,omitempty"`
	// PoeModeN: * `pd` - PD
	// * `pse` - PSE
	PoeModeN []string `json:"poe_mode__n,omitempty" mapstructure:"poe_mode__n,omitempty"`
	// PoeType: * `type1-ieee802.3af` - 802.3af (Type 1)
	// * `type2-ieee802.3at` - 802.3at (Type 2)
	// * `type3-ieee802.3bt` - 802.3bt (Type 3)
	// * `type4-ieee802.3bt` - 802.3bt (Type 4)
	// * `passive-24v-2pair` - Passive 24V (2-pair)
	// * `passive-24v-4pair` - Passive 24V (4-pair)
	// * `passive-48v-2pair` - Passive 48V (2-pair)
	// * `passive-48v-4pair` - Passive 48V (4-pair)
	PoeType []string `json:"poe_type,omitempty" mapstructure:"poe_type,omitempty"`
	// PoeTypeN: * `type1-ieee802.3af` - 802.3af (Type 1)
	// * `type2-ieee802.3at` - 802.3at (Type 2)
	// * `type3-ieee802.3bt` - 802.3bt (Type 3)
	// * `type4-ieee802.3bt` - 802.3bt (Type 4)
	// * `passive-24v-2pair` - Passive 24V (2-pair)
	// * `passive-24v-4pair` - Passive 24V (4-pair)
	// * `passive-48v-2pair` - Passive 48V (2-pair)
	// * `passive-48v-4pair` - Passive 48V (4-pair)
	PoeTypeN []string `json:"poe_type__n,omitempty" mapstructure:"poe_type__n,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Rack: Rack (name)
	Rack []string `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// RackN: Rack (name)
	RackN []string `json:"rack__n,omitempty" mapstructure:"rack__n,omitempty"`
	// RackId: Rack (ID)
	RackId []int32 `json:"rack_id,omitempty" mapstructure:"rack_id,omitempty"`
	// RackIdN: Rack (ID)
	RackIdN []int32 `json:"rack_id__n,omitempty" mapstructure:"rack_id__n,omitempty"`
	// Region: Region (slug)
	Region []int32 `json:"region,omitempty" mapstructure:"region,omitempty"`
	// RegionN: Region (slug)
	RegionN []int32 `json:"region__n,omitempty" mapstructure:"region__n,omitempty"`
	// RegionId: Region (ID)
	RegionId []int32 `json:"region_id,omitempty" mapstructure:"region_id,omitempty"`
	// RegionIdN: Region (ID)
	RegionIdN []int32 `json:"region_id__n,omitempty" mapstructure:"region_id__n,omitempty"`
	// RfChannel: * `2.4g-1-2412-22` - 1 (2412 MHz)
	// * `2.4g-2-2417-22` - 2 (2417 MHz)
	// * `2.4g-3-2422-22` - 3 (2422 MHz)
	// * `2.4g-4-2427-22` - 4 (2427 MHz)
	// * `2.4g-5-2432-22` - 5 (2432 MHz)
	// * `2.4g-6-2437-22` - 6 (2437 MHz)
	// * `2.4g-7-2442-22` - 7 (2442 MHz)
	// * `2.4g-8-2447-22` - 8 (2447 MHz)
	// * `2.4g-9-2452-22` - 9 (2452 MHz)
	// * `2.4g-10-2457-22` - 10 (2457 MHz)
	// * `2.4g-11-2462-22` - 11 (2462 MHz)
	// * `2.4g-12-2467-22` - 12 (2467 MHz)
	// * `2.4g-13-2472-22` - 13 (2472 MHz)
	// * `5g-32-5160-20` - 32 (5160/20 MHz)
	// * `5g-34-5170-40` - 34 (5170/40 MHz)
	// * `5g-36-5180-20` - 36 (5180/20 MHz)
	// * `5g-38-5190-40` - 38 (5190/40 MHz)
	// * `5g-40-5200-20` - 40 (5200/20 MHz)
	// * `5g-42-5210-80` - 42 (5210/80 MHz)
	// * `5g-44-5220-20` - 44 (5220/20 MHz)
	// * `5g-46-5230-40` - 46 (5230/40 MHz)
	// * `5g-48-5240-20` - 48 (5240/20 MHz)
	// * `5g-50-5250-160` - 50 (5250/160 MHz)
	// * `5g-52-5260-20` - 52 (5260/20 MHz)
	// * `5g-54-5270-40` - 54 (5270/40 MHz)
	// * `5g-56-5280-20` - 56 (5280/20 MHz)
	// * `5g-58-5290-80` - 58 (5290/80 MHz)
	// * `5g-60-5300-20` - 60 (5300/20 MHz)
	// * `5g-62-5310-40` - 62 (5310/40 MHz)
	// * `5g-64-5320-20` - 64 (5320/20 MHz)
	// * `5g-100-5500-20` - 100 (5500/20 MHz)
	// * `5g-102-5510-40` - 102 (5510/40 MHz)
	// * `5g-104-5520-20` - 104 (5520/20 MHz)
	// * `5g-106-5530-80` - 106 (5530/80 MHz)
	// * `5g-108-5540-20` - 108 (5540/20 MHz)
	// * `5g-110-5550-40` - 110 (5550/40 MHz)
	// * `5g-112-5560-20` - 112 (5560/20 MHz)
	// * `5g-114-5570-160` - 114 (5570/160 MHz)
	// * `5g-116-5580-20` - 116 (5580/20 MHz)
	// * `5g-118-5590-40` - 118 (5590/40 MHz)
	// * `5g-120-5600-20` - 120 (5600/20 MHz)
	// * `5g-122-5610-80` - 122 (5610/80 MHz)
	// * `5g-124-5620-20` - 124 (5620/20 MHz)
	// * `5g-126-5630-40` - 126 (5630/40 MHz)
	// * `5g-128-5640-20` - 128 (5640/20 MHz)
	// * `5g-132-5660-20` - 132 (5660/20 MHz)
	// * `5g-134-5670-40` - 134 (5670/40 MHz)
	// * `5g-136-5680-20` - 136 (5680/20 MHz)
	// * `5g-138-5690-80` - 138 (5690/80 MHz)
	// * `5g-140-5700-20` - 140 (5700/20 MHz)
	// * `5g-142-5710-40` - 142 (5710/40 MHz)
	// * `5g-144-5720-20` - 144 (5720/20 MHz)
	// * `5g-149-5745-20` - 149 (5745/20 MHz)
	// * `5g-151-5755-40` - 151 (5755/40 MHz)
	// * `5g-153-5765-20` - 153 (5765/20 MHz)
	// * `5g-155-5775-80` - 155 (5775/80 MHz)
	// * `5g-157-5785-20` - 157 (5785/20 MHz)
	// * `5g-159-5795-40` - 159 (5795/40 MHz)
	// * `5g-161-5805-20` - 161 (5805/20 MHz)
	// * `5g-163-5815-160` - 163 (5815/160 MHz)
	// * `5g-165-5825-20` - 165 (5825/20 MHz)
	// * `5g-167-5835-40` - 167 (5835/40 MHz)
	// * `5g-169-5845-20` - 169 (5845/20 MHz)
	// * `5g-171-5855-80` - 171 (5855/80 MHz)
	// * `5g-173-5865-20` - 173 (5865/20 MHz)
	// * `5g-175-5875-40` - 175 (5875/40 MHz)
	// * `5g-177-5885-20` - 177 (5885/20 MHz)
	// * `6g-1-5955-20` - 1 (5955/20 MHz)
	// * `6g-3-5965-40` - 3 (5965/40 MHz)
	// * `6g-5-5975-20` - 5 (5975/20 MHz)
	// * `6g-7-5985-80` - 7 (5985/80 MHz)
	// * `6g-9-5995-20` - 9 (5995/20 MHz)
	// * `6g-11-6005-40` - 11 (6005/40 MHz)
	// * `6g-13-6015-20` - 13 (6015/20 MHz)
	// * `6g-15-6025-160` - 15 (6025/160 MHz)
	// * `6g-17-6035-20` - 17 (6035/20 MHz)
	// * `6g-19-6045-40` - 19 (6045/40 MHz)
	// * `6g-21-6055-20` - 21 (6055/20 MHz)
	// * `6g-23-6065-80` - 23 (6065/80 MHz)
	// * `6g-25-6075-20` - 25 (6075/20 MHz)
	// * `6g-27-6085-40` - 27 (6085/40 MHz)
	// * `6g-29-6095-20` - 29 (6095/20 MHz)
	// * `6g-31-6105-320` - 31 (6105/320 MHz)
	// * `6g-33-6115-20` - 33 (6115/20 MHz)
	// * `6g-35-6125-40` - 35 (6125/40 MHz)
	// * `6g-37-6135-20` - 37 (6135/20 MHz)
	// * `6g-39-6145-80` - 39 (6145/80 MHz)
	// * `6g-41-6155-20` - 41 (6155/20 MHz)
	// * `6g-43-6165-40` - 43 (6165/40 MHz)
	// * `6g-45-6175-20` - 45 (6175/20 MHz)
	// * `6g-47-6185-160` - 47 (6185/160 MHz)
	// * `6g-49-6195-20` - 49 (6195/20 MHz)
	// * `6g-51-6205-40` - 51 (6205/40 MHz)
	// * `6g-53-6215-20` - 53 (6215/20 MHz)
	// * `6g-55-6225-80` - 55 (6225/80 MHz)
	// * `6g-57-6235-20` - 57 (6235/20 MHz)
	// * `6g-59-6245-40` - 59 (6245/40 MHz)
	// * `6g-61-6255-20` - 61 (6255/20 MHz)
	// * `6g-65-6275-20` - 65 (6275/20 MHz)
	// * `6g-67-6285-40` - 67 (6285/40 MHz)
	// * `6g-69-6295-20` - 69 (6295/20 MHz)
	// * `6g-71-6305-80` - 71 (6305/80 MHz)
	// * `6g-73-6315-20` - 73 (6315/20 MHz)
	// * `6g-75-6325-40` - 75 (6325/40 MHz)
	// * `6g-77-6335-20` - 77 (6335/20 MHz)
	// * `6g-79-6345-160` - 79 (6345/160 MHz)
	// * `6g-81-6355-20` - 81 (6355/20 MHz)
	// * `6g-83-6365-40` - 83 (6365/40 MHz)
	// * `6g-85-6375-20` - 85 (6375/20 MHz)
	// * `6g-87-6385-80` - 87 (6385/80 MHz)
	// * `6g-89-6395-20` - 89 (6395/20 MHz)
	// * `6g-91-6405-40` - 91 (6405/40 MHz)
	// * `6g-93-6415-20` - 93 (6415/20 MHz)
	// * `6g-95-6425-320` - 95 (6425/320 MHz)
	// * `6g-97-6435-20` - 97 (6435/20 MHz)
	// * `6g-99-6445-40` - 99 (6445/40 MHz)
	// * `6g-101-6455-20` - 101 (6455/20 MHz)
	// * `6g-103-6465-80` - 103 (6465/80 MHz)
	// * `6g-105-6475-20` - 105 (6475/20 MHz)
	// * `6g-107-6485-40` - 107 (6485/40 MHz)
	// * `6g-109-6495-20` - 109 (6495/20 MHz)
	// * `6g-111-6505-160` - 111 (6505/160 MHz)
	// * `6g-113-6515-20` - 113 (6515/20 MHz)
	// * `6g-115-6525-40` - 115 (6525/40 MHz)
	// * `6g-117-6535-20` - 117 (6535/20 MHz)
	// * `6g-119-6545-80` - 119 (6545/80 MHz)
	// * `6g-121-6555-20` - 121 (6555/20 MHz)
	// * `6g-123-6565-40` - 123 (6565/40 MHz)
	// * `6g-125-6575-20` - 125 (6575/20 MHz)
	// * `6g-129-6595-20` - 129 (6595/20 MHz)
	// * `6g-131-6605-40` - 131 (6605/40 MHz)
	// * `6g-133-6615-20` - 133 (6615/20 MHz)
	// * `6g-135-6625-80` - 135 (6625/80 MHz)
	// * `6g-137-6635-20` - 137 (6635/20 MHz)
	// * `6g-139-6645-40` - 139 (6645/40 MHz)
	// * `6g-141-6655-20` - 141 (6655/20 MHz)
	// * `6g-143-6665-160` - 143 (6665/160 MHz)
	// * `6g-145-6675-20` - 145 (6675/20 MHz)
	// * `6g-147-6685-40` - 147 (6685/40 MHz)
	// * `6g-149-6695-20` - 149 (6695/20 MHz)
	// * `6g-151-6705-80` - 151 (6705/80 MHz)
	// * `6g-153-6715-20` - 153 (6715/20 MHz)
	// * `6g-155-6725-40` - 155 (6725/40 MHz)
	// * `6g-157-6735-20` - 157 (6735/20 MHz)
	// * `6g-159-6745-320` - 159 (6745/320 MHz)
	// * `6g-161-6755-20` - 161 (6755/20 MHz)
	// * `6g-163-6765-40` - 163 (6765/40 MHz)
	// * `6g-165-6775-20` - 165 (6775/20 MHz)
	// * `6g-167-6785-80` - 167 (6785/80 MHz)
	// * `6g-169-6795-20` - 169 (6795/20 MHz)
	// * `6g-171-6805-40` - 171 (6805/40 MHz)
	// * `6g-173-6815-20` - 173 (6815/20 MHz)
	// * `6g-175-6825-160` - 175 (6825/160 MHz)
	// * `6g-177-6835-20` - 177 (6835/20 MHz)
	// * `6g-179-6845-40` - 179 (6845/40 MHz)
	// * `6g-181-6855-20` - 181 (6855/20 MHz)
	// * `6g-183-6865-80` - 183 (6865/80 MHz)
	// * `6g-185-6875-20` - 185 (6875/20 MHz)
	// * `6g-187-6885-40` - 187 (6885/40 MHz)
	// * `6g-189-6895-20` - 189 (6895/20 MHz)
	// * `6g-193-6915-20` - 193 (6915/20 MHz)
	// * `6g-195-6925-40` - 195 (6925/40 MHz)
	// * `6g-197-6935-20` - 197 (6935/20 MHz)
	// * `6g-199-6945-80` - 199 (6945/80 MHz)
	// * `6g-201-6955-20` - 201 (6955/20 MHz)
	// * `6g-203-6965-40` - 203 (6965/40 MHz)
	// * `6g-205-6975-20` - 205 (6975/20 MHz)
	// * `6g-207-6985-160` - 207 (6985/160 MHz)
	// * `6g-209-6995-20` - 209 (6995/20 MHz)
	// * `6g-211-7005-40` - 211 (7005/40 MHz)
	// * `6g-213-7015-20` - 213 (7015/20 MHz)
	// * `6g-215-7025-80` - 215 (7025/80 MHz)
	// * `6g-217-7035-20` - 217 (7035/20 MHz)
	// * `6g-219-7045-40` - 219 (7045/40 MHz)
	// * `6g-221-7055-20` - 221 (7055/20 MHz)
	// * `6g-225-7075-20` - 225 (7075/20 MHz)
	// * `6g-227-7085-40` - 227 (7085/40 MHz)
	// * `6g-229-7095-20` - 229 (7095/20 MHz)
	// * `6g-233-7115-20` - 233 (7115/20 MHz)
	// * `60g-1-58320-2160` - 1 (58.32/2.16 GHz)
	// * `60g-2-60480-2160` - 2 (60.48/2.16 GHz)
	// * `60g-3-62640-2160` - 3 (62.64/2.16 GHz)
	// * `60g-4-64800-2160` - 4 (64.80/2.16 GHz)
	// * `60g-5-66960-2160` - 5 (66.96/2.16 GHz)
	// * `60g-6-69120-2160` - 6 (69.12/2.16 GHz)
	// * `60g-9-59400-4320` - 9 (59.40/4.32 GHz)
	// * `60g-10-61560-4320` - 10 (61.56/4.32 GHz)
	// * `60g-11-63720-4320` - 11 (63.72/4.32 GHz)
	// * `60g-12-65880-4320` - 12 (65.88/4.32 GHz)
	// * `60g-13-68040-4320` - 13 (68.04/4.32 GHz)
	// * `60g-17-60480-6480` - 17 (60.48/6.48 GHz)
	// * `60g-18-62640-6480` - 18 (62.64/6.48 GHz)
	// * `60g-19-64800-6480` - 19 (64.80/6.48 GHz)
	// * `60g-20-66960-6480` - 20 (66.96/6.48 GHz)
	// * `60g-25-61560-6480` - 25 (61.56/8.64 GHz)
	// * `60g-26-63720-6480` - 26 (63.72/8.64 GHz)
	// * `60g-27-65880-6480` - 27 (65.88/8.64 GHz)
	RfChannel []string `json:"rf_channel,omitempty" mapstructure:"rf_channel,omitempty"`
	// RfChannelN: * `2.4g-1-2412-22` - 1 (2412 MHz)
	// * `2.4g-2-2417-22` - 2 (2417 MHz)
	// * `2.4g-3-2422-22` - 3 (2422 MHz)
	// * `2.4g-4-2427-22` - 4 (2427 MHz)
	// * `2.4g-5-2432-22` - 5 (2432 MHz)
	// * `2.4g-6-2437-22` - 6 (2437 MHz)
	// * `2.4g-7-2442-22` - 7 (2442 MHz)
	// * `2.4g-8-2447-22` - 8 (2447 MHz)
	// * `2.4g-9-2452-22` - 9 (2452 MHz)
	// * `2.4g-10-2457-22` - 10 (2457 MHz)
	// * `2.4g-11-2462-22` - 11 (2462 MHz)
	// * `2.4g-12-2467-22` - 12 (2467 MHz)
	// * `2.4g-13-2472-22` - 13 (2472 MHz)
	// * `5g-32-5160-20` - 32 (5160/20 MHz)
	// * `5g-34-5170-40` - 34 (5170/40 MHz)
	// * `5g-36-5180-20` - 36 (5180/20 MHz)
	// * `5g-38-5190-40` - 38 (5190/40 MHz)
	// * `5g-40-5200-20` - 40 (5200/20 MHz)
	// * `5g-42-5210-80` - 42 (5210/80 MHz)
	// * `5g-44-5220-20` - 44 (5220/20 MHz)
	// * `5g-46-5230-40` - 46 (5230/40 MHz)
	// * `5g-48-5240-20` - 48 (5240/20 MHz)
	// * `5g-50-5250-160` - 50 (5250/160 MHz)
	// * `5g-52-5260-20` - 52 (5260/20 MHz)
	// * `5g-54-5270-40` - 54 (5270/40 MHz)
	// * `5g-56-5280-20` - 56 (5280/20 MHz)
	// * `5g-58-5290-80` - 58 (5290/80 MHz)
	// * `5g-60-5300-20` - 60 (5300/20 MHz)
	// * `5g-62-5310-40` - 62 (5310/40 MHz)
	// * `5g-64-5320-20` - 64 (5320/20 MHz)
	// * `5g-100-5500-20` - 100 (5500/20 MHz)
	// * `5g-102-5510-40` - 102 (5510/40 MHz)
	// * `5g-104-5520-20` - 104 (5520/20 MHz)
	// * `5g-106-5530-80` - 106 (5530/80 MHz)
	// * `5g-108-5540-20` - 108 (5540/20 MHz)
	// * `5g-110-5550-40` - 110 (5550/40 MHz)
	// * `5g-112-5560-20` - 112 (5560/20 MHz)
	// * `5g-114-5570-160` - 114 (5570/160 MHz)
	// * `5g-116-5580-20` - 116 (5580/20 MHz)
	// * `5g-118-5590-40` - 118 (5590/40 MHz)
	// * `5g-120-5600-20` - 120 (5600/20 MHz)
	// * `5g-122-5610-80` - 122 (5610/80 MHz)
	// * `5g-124-5620-20` - 124 (5620/20 MHz)
	// * `5g-126-5630-40` - 126 (5630/40 MHz)
	// * `5g-128-5640-20` - 128 (5640/20 MHz)
	// * `5g-132-5660-20` - 132 (5660/20 MHz)
	// * `5g-134-5670-40` - 134 (5670/40 MHz)
	// * `5g-136-5680-20` - 136 (5680/20 MHz)
	// * `5g-138-5690-80` - 138 (5690/80 MHz)
	// * `5g-140-5700-20` - 140 (5700/20 MHz)
	// * `5g-142-5710-40` - 142 (5710/40 MHz)
	// * `5g-144-5720-20` - 144 (5720/20 MHz)
	// * `5g-149-5745-20` - 149 (5745/20 MHz)
	// * `5g-151-5755-40` - 151 (5755/40 MHz)
	// * `5g-153-5765-20` - 153 (5765/20 MHz)
	// * `5g-155-5775-80` - 155 (5775/80 MHz)
	// * `5g-157-5785-20` - 157 (5785/20 MHz)
	// * `5g-159-5795-40` - 159 (5795/40 MHz)
	// * `5g-161-5805-20` - 161 (5805/20 MHz)
	// * `5g-163-5815-160` - 163 (5815/160 MHz)
	// * `5g-165-5825-20` - 165 (5825/20 MHz)
	// * `5g-167-5835-40` - 167 (5835/40 MHz)
	// * `5g-169-5845-20` - 169 (5845/20 MHz)
	// * `5g-171-5855-80` - 171 (5855/80 MHz)
	// * `5g-173-5865-20` - 173 (5865/20 MHz)
	// * `5g-175-5875-40` - 175 (5875/40 MHz)
	// * `5g-177-5885-20` - 177 (5885/20 MHz)
	// * `6g-1-5955-20` - 1 (5955/20 MHz)
	// * `6g-3-5965-40` - 3 (5965/40 MHz)
	// * `6g-5-5975-20` - 5 (5975/20 MHz)
	// * `6g-7-5985-80` - 7 (5985/80 MHz)
	// * `6g-9-5995-20` - 9 (5995/20 MHz)
	// * `6g-11-6005-40` - 11 (6005/40 MHz)
	// * `6g-13-6015-20` - 13 (6015/20 MHz)
	// * `6g-15-6025-160` - 15 (6025/160 MHz)
	// * `6g-17-6035-20` - 17 (6035/20 MHz)
	// * `6g-19-6045-40` - 19 (6045/40 MHz)
	// * `6g-21-6055-20` - 21 (6055/20 MHz)
	// * `6g-23-6065-80` - 23 (6065/80 MHz)
	// * `6g-25-6075-20` - 25 (6075/20 MHz)
	// * `6g-27-6085-40` - 27 (6085/40 MHz)
	// * `6g-29-6095-20` - 29 (6095/20 MHz)
	// * `6g-31-6105-320` - 31 (6105/320 MHz)
	// * `6g-33-6115-20` - 33 (6115/20 MHz)
	// * `6g-35-6125-40` - 35 (6125/40 MHz)
	// * `6g-37-6135-20` - 37 (6135/20 MHz)
	// * `6g-39-6145-80` - 39 (6145/80 MHz)
	// * `6g-41-6155-20` - 41 (6155/20 MHz)
	// * `6g-43-6165-40` - 43 (6165/40 MHz)
	// * `6g-45-6175-20` - 45 (6175/20 MHz)
	// * `6g-47-6185-160` - 47 (6185/160 MHz)
	// * `6g-49-6195-20` - 49 (6195/20 MHz)
	// * `6g-51-6205-40` - 51 (6205/40 MHz)
	// * `6g-53-6215-20` - 53 (6215/20 MHz)
	// * `6g-55-6225-80` - 55 (6225/80 MHz)
	// * `6g-57-6235-20` - 57 (6235/20 MHz)
	// * `6g-59-6245-40` - 59 (6245/40 MHz)
	// * `6g-61-6255-20` - 61 (6255/20 MHz)
	// * `6g-65-6275-20` - 65 (6275/20 MHz)
	// * `6g-67-6285-40` - 67 (6285/40 MHz)
	// * `6g-69-6295-20` - 69 (6295/20 MHz)
	// * `6g-71-6305-80` - 71 (6305/80 MHz)
	// * `6g-73-6315-20` - 73 (6315/20 MHz)
	// * `6g-75-6325-40` - 75 (6325/40 MHz)
	// * `6g-77-6335-20` - 77 (6335/20 MHz)
	// * `6g-79-6345-160` - 79 (6345/160 MHz)
	// * `6g-81-6355-20` - 81 (6355/20 MHz)
	// * `6g-83-6365-40` - 83 (6365/40 MHz)
	// * `6g-85-6375-20` - 85 (6375/20 MHz)
	// * `6g-87-6385-80` - 87 (6385/80 MHz)
	// * `6g-89-6395-20` - 89 (6395/20 MHz)
	// * `6g-91-6405-40` - 91 (6405/40 MHz)
	// * `6g-93-6415-20` - 93 (6415/20 MHz)
	// * `6g-95-6425-320` - 95 (6425/320 MHz)
	// * `6g-97-6435-20` - 97 (6435/20 MHz)
	// * `6g-99-6445-40` - 99 (6445/40 MHz)
	// * `6g-101-6455-20` - 101 (6455/20 MHz)
	// * `6g-103-6465-80` - 103 (6465/80 MHz)
	// * `6g-105-6475-20` - 105 (6475/20 MHz)
	// * `6g-107-6485-40` - 107 (6485/40 MHz)
	// * `6g-109-6495-20` - 109 (6495/20 MHz)
	// * `6g-111-6505-160` - 111 (6505/160 MHz)
	// * `6g-113-6515-20` - 113 (6515/20 MHz)
	// * `6g-115-6525-40` - 115 (6525/40 MHz)
	// * `6g-117-6535-20` - 117 (6535/20 MHz)
	// * `6g-119-6545-80` - 119 (6545/80 MHz)
	// * `6g-121-6555-20` - 121 (6555/20 MHz)
	// * `6g-123-6565-40` - 123 (6565/40 MHz)
	// * `6g-125-6575-20` - 125 (6575/20 MHz)
	// * `6g-129-6595-20` - 129 (6595/20 MHz)
	// * `6g-131-6605-40` - 131 (6605/40 MHz)
	// * `6g-133-6615-20` - 133 (6615/20 MHz)
	// * `6g-135-6625-80` - 135 (6625/80 MHz)
	// * `6g-137-6635-20` - 137 (6635/20 MHz)
	// * `6g-139-6645-40` - 139 (6645/40 MHz)
	// * `6g-141-6655-20` - 141 (6655/20 MHz)
	// * `6g-143-6665-160` - 143 (6665/160 MHz)
	// * `6g-145-6675-20` - 145 (6675/20 MHz)
	// * `6g-147-6685-40` - 147 (6685/40 MHz)
	// * `6g-149-6695-20` - 149 (6695/20 MHz)
	// * `6g-151-6705-80` - 151 (6705/80 MHz)
	// * `6g-153-6715-20` - 153 (6715/20 MHz)
	// * `6g-155-6725-40` - 155 (6725/40 MHz)
	// * `6g-157-6735-20` - 157 (6735/20 MHz)
	// * `6g-159-6745-320` - 159 (6745/320 MHz)
	// * `6g-161-6755-20` - 161 (6755/20 MHz)
	// * `6g-163-6765-40` - 163 (6765/40 MHz)
	// * `6g-165-6775-20` - 165 (6775/20 MHz)
	// * `6g-167-6785-80` - 167 (6785/80 MHz)
	// * `6g-169-6795-20` - 169 (6795/20 MHz)
	// * `6g-171-6805-40` - 171 (6805/40 MHz)
	// * `6g-173-6815-20` - 173 (6815/20 MHz)
	// * `6g-175-6825-160` - 175 (6825/160 MHz)
	// * `6g-177-6835-20` - 177 (6835/20 MHz)
	// * `6g-179-6845-40` - 179 (6845/40 MHz)
	// * `6g-181-6855-20` - 181 (6855/20 MHz)
	// * `6g-183-6865-80` - 183 (6865/80 MHz)
	// * `6g-185-6875-20` - 185 (6875/20 MHz)
	// * `6g-187-6885-40` - 187 (6885/40 MHz)
	// * `6g-189-6895-20` - 189 (6895/20 MHz)
	// * `6g-193-6915-20` - 193 (6915/20 MHz)
	// * `6g-195-6925-40` - 195 (6925/40 MHz)
	// * `6g-197-6935-20` - 197 (6935/20 MHz)
	// * `6g-199-6945-80` - 199 (6945/80 MHz)
	// * `6g-201-6955-20` - 201 (6955/20 MHz)
	// * `6g-203-6965-40` - 203 (6965/40 MHz)
	// * `6g-205-6975-20` - 205 (6975/20 MHz)
	// * `6g-207-6985-160` - 207 (6985/160 MHz)
	// * `6g-209-6995-20` - 209 (6995/20 MHz)
	// * `6g-211-7005-40` - 211 (7005/40 MHz)
	// * `6g-213-7015-20` - 213 (7015/20 MHz)
	// * `6g-215-7025-80` - 215 (7025/80 MHz)
	// * `6g-217-7035-20` - 217 (7035/20 MHz)
	// * `6g-219-7045-40` - 219 (7045/40 MHz)
	// * `6g-221-7055-20` - 221 (7055/20 MHz)
	// * `6g-225-7075-20` - 225 (7075/20 MHz)
	// * `6g-227-7085-40` - 227 (7085/40 MHz)
	// * `6g-229-7095-20` - 229 (7095/20 MHz)
	// * `6g-233-7115-20` - 233 (7115/20 MHz)
	// * `60g-1-58320-2160` - 1 (58.32/2.16 GHz)
	// * `60g-2-60480-2160` - 2 (60.48/2.16 GHz)
	// * `60g-3-62640-2160` - 3 (62.64/2.16 GHz)
	// * `60g-4-64800-2160` - 4 (64.80/2.16 GHz)
	// * `60g-5-66960-2160` - 5 (66.96/2.16 GHz)
	// * `60g-6-69120-2160` - 6 (69.12/2.16 GHz)
	// * `60g-9-59400-4320` - 9 (59.40/4.32 GHz)
	// * `60g-10-61560-4320` - 10 (61.56/4.32 GHz)
	// * `60g-11-63720-4320` - 11 (63.72/4.32 GHz)
	// * `60g-12-65880-4320` - 12 (65.88/4.32 GHz)
	// * `60g-13-68040-4320` - 13 (68.04/4.32 GHz)
	// * `60g-17-60480-6480` - 17 (60.48/6.48 GHz)
	// * `60g-18-62640-6480` - 18 (62.64/6.48 GHz)
	// * `60g-19-64800-6480` - 19 (64.80/6.48 GHz)
	// * `60g-20-66960-6480` - 20 (66.96/6.48 GHz)
	// * `60g-25-61560-6480` - 25 (61.56/8.64 GHz)
	// * `60g-26-63720-6480` - 26 (63.72/8.64 GHz)
	// * `60g-27-65880-6480` - 27 (65.88/8.64 GHz)
	RfChannelN []string `json:"rf_channel__n,omitempty" mapstructure:"rf_channel__n,omitempty"`
	// RfChannelFrequency:
	RfChannelFrequency []float64 `json:"rf_channel_frequency,omitempty" mapstructure:"rf_channel_frequency,omitempty"`
	// RfChannelFrequencyGt:
	RfChannelFrequencyGt []float64 `json:"rf_channel_frequency__gt,omitempty" mapstructure:"rf_channel_frequency__gt,omitempty"`
	// RfChannelFrequencyGte:
	RfChannelFrequencyGte []float64 `json:"rf_channel_frequency__gte,omitempty" mapstructure:"rf_channel_frequency__gte,omitempty"`
	// RfChannelFrequencyLt:
	RfChannelFrequencyLt []float64 `json:"rf_channel_frequency__lt,omitempty" mapstructure:"rf_channel_frequency__lt,omitempty"`
	// RfChannelFrequencyLte:
	RfChannelFrequencyLte []float64 `json:"rf_channel_frequency__lte,omitempty" mapstructure:"rf_channel_frequency__lte,omitempty"`
	// RfChannelFrequencyN:
	RfChannelFrequencyN []float64 `json:"rf_channel_frequency__n,omitempty" mapstructure:"rf_channel_frequency__n,omitempty"`
	// RfChannelWidth:
	RfChannelWidth []float64 `json:"rf_channel_width,omitempty" mapstructure:"rf_channel_width,omitempty"`
	// RfChannelWidthGt:
	RfChannelWidthGt []float64 `json:"rf_channel_width__gt,omitempty" mapstructure:"rf_channel_width__gt,omitempty"`
	// RfChannelWidthGte:
	RfChannelWidthGte []float64 `json:"rf_channel_width__gte,omitempty" mapstructure:"rf_channel_width__gte,omitempty"`
	// RfChannelWidthLt:
	RfChannelWidthLt []float64 `json:"rf_channel_width__lt,omitempty" mapstructure:"rf_channel_width__lt,omitempty"`
	// RfChannelWidthLte:
	RfChannelWidthLte []float64 `json:"rf_channel_width__lte,omitempty" mapstructure:"rf_channel_width__lte,omitempty"`
	// RfChannelWidthN:
	RfChannelWidthN []float64 `json:"rf_channel_width__n,omitempty" mapstructure:"rf_channel_width__n,omitempty"`
	// RfRole: * `ap` - Access point
	// * `station` - Station
	RfRole []string `json:"rf_role,omitempty" mapstructure:"rf_role,omitempty"`
	// RfRoleN: * `ap` - Access point
	// * `station` - Station
	RfRoleN []string `json:"rf_role__n,omitempty" mapstructure:"rf_role__n,omitempty"`
	// Site: Site name (slug)
	Site []string `json:"site,omitempty" mapstructure:"site,omitempty"`
	// SiteN: Site name (slug)
	SiteN []string `json:"site__n,omitempty" mapstructure:"site__n,omitempty"`
	// SiteGroup: Site group (slug)
	SiteGroup []int32 `json:"site_group,omitempty" mapstructure:"site_group,omitempty"`
	// SiteGroupN: Site group (slug)
	SiteGroupN []int32 `json:"site_group__n,omitempty" mapstructure:"site_group__n,omitempty"`
	// SiteGroupId: Site group (ID)
	SiteGroupId []int32 `json:"site_group_id,omitempty" mapstructure:"site_group_id,omitempty"`
	// SiteGroupIdN: Site group (ID)
	SiteGroupIdN []int32 `json:"site_group_id__n,omitempty" mapstructure:"site_group_id__n,omitempty"`
	// SiteId: Site (ID)
	SiteId []int32 `json:"site_id,omitempty" mapstructure:"site_id,omitempty"`
	// SiteIdN: Site (ID)
	SiteIdN []int32 `json:"site_id__n,omitempty" mapstructure:"site_id__n,omitempty"`
	// Speed:
	Speed []int32 `json:"speed,omitempty" mapstructure:"speed,omitempty"`
	// SpeedGt:
	SpeedGt []int32 `json:"speed__gt,omitempty" mapstructure:"speed__gt,omitempty"`
	// SpeedGte:
	SpeedGte []int32 `json:"speed__gte,omitempty" mapstructure:"speed__gte,omitempty"`
	// SpeedLt:
	SpeedLt []int32 `json:"speed__lt,omitempty" mapstructure:"speed__lt,omitempty"`
	// SpeedLte:
	SpeedLte []int32 `json:"speed__lte,omitempty" mapstructure:"speed__lte,omitempty"`
	// SpeedN:
	SpeedN []int32 `json:"speed__n,omitempty" mapstructure:"speed__n,omitempty"`
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
	// TxPower:
	TxPower []int32 `json:"tx_power,omitempty" mapstructure:"tx_power,omitempty"`
	// TxPowerGt:
	TxPowerGt []int32 `json:"tx_power__gt,omitempty" mapstructure:"tx_power__gt,omitempty"`
	// TxPowerGte:
	TxPowerGte []int32 `json:"tx_power__gte,omitempty" mapstructure:"tx_power__gte,omitempty"`
	// TxPowerLt:
	TxPowerLt []int32 `json:"tx_power__lt,omitempty" mapstructure:"tx_power__lt,omitempty"`
	// TxPowerLte:
	TxPowerLte []int32 `json:"tx_power__lte,omitempty" mapstructure:"tx_power__lte,omitempty"`
	// TxPowerN:
	TxPowerN []int32 `json:"tx_power__n,omitempty" mapstructure:"tx_power__n,omitempty"`
	// Type: * `virtual` - Virtual
	// * `bridge` - Bridge
	// * `lag` - Link Aggregation Group (LAG)
	// * `100base-fx` - 100BASE-FX (10/100ME FIBER)
	// * `100base-lfx` - 100BASE-LFX (10/100ME FIBER)
	// * `100base-tx` - 100BASE-TX (10/100ME)
	// * `100base-t1` - 100BASE-T1 (10/100ME Single Pair)
	// * `1000base-t` - 1000BASE-T (1GE)
	// * `2.5gbase-t` - 2.5GBASE-T (2.5GE)
	// * `5gbase-t` - 5GBASE-T (5GE)
	// * `10gbase-t` - 10GBASE-T (10GE)
	// * `10gbase-cx4` - 10GBASE-CX4 (10GE)
	// * `1000base-x-gbic` - GBIC (1GE)
	// * `1000base-x-sfp` - SFP (1GE)
	// * `10gbase-x-sfpp` - SFP+ (10GE)
	// * `10gbase-x-xfp` - XFP (10GE)
	// * `10gbase-x-xenpak` - XENPAK (10GE)
	// * `10gbase-x-x2` - X2 (10GE)
	// * `25gbase-x-sfp28` - SFP28 (25GE)
	// * `50gbase-x-sfp56` - SFP56 (50GE)
	// * `40gbase-x-qsfpp` - QSFP+ (40GE)
	// * `50gbase-x-sfp28` - QSFP28 (50GE)
	// * `100gbase-x-cfp` - CFP (100GE)
	// * `100gbase-x-cfp2` - CFP2 (100GE)
	// * `200gbase-x-cfp2` - CFP2 (200GE)
	// * `100gbase-x-cfp4` - CFP4 (100GE)
	// * `100gbase-x-cpak` - Cisco CPAK (100GE)
	// * `100gbase-x-qsfp28` - QSFP28 (100GE)
	// * `200gbase-x-qsfp56` - QSFP56 (200GE)
	// * `400gbase-x-qsfpdd` - QSFP-DD (400GE)
	// * `400gbase-x-osfp` - OSFP (400GE)
	// * `800gbase-x-qsfpdd` - QSFP-DD (800GE)
	// * `800gbase-x-osfp` - OSFP (800GE)
	// * `1000base-kx` - 1000BASE-KX (1GE)
	// * `10gbase-kr` - 10GBASE-KR (10GE)
	// * `10gbase-kx4` - 10GBASE-KX4 (10GE)
	// * `25gbase-kr` - 25GBASE-KR (25GE)
	// * `40gbase-kr4` - 40GBASE-KR4 (40GE)
	// * `50gbase-kr` - 50GBASE-KR (50GE)
	// * `100gbase-kp4` - 100GBASE-KP4 (100GE)
	// * `100gbase-kr2` - 100GBASE-KR2 (100GE)
	// * `100gbase-kr4` - 100GBASE-KR4 (100GE)
	// * `ieee802.11a` - IEEE 802.11a
	// * `ieee802.11g` - IEEE 802.11b/g
	// * `ieee802.11n` - IEEE 802.11n
	// * `ieee802.11ac` - IEEE 802.11ac
	// * `ieee802.11ad` - IEEE 802.11ad
	// * `ieee802.11ax` - IEEE 802.11ax
	// * `ieee802.11ay` - IEEE 802.11ay
	// * `ieee802.15.1` - IEEE 802.15.1 (Bluetooth)
	// * `other-wireless` - Other (Wireless)
	// * `gsm` - GSM
	// * `cdma` - CDMA
	// * `lte` - LTE
	// * `sonet-oc3` - OC-3/STM-1
	// * `sonet-oc12` - OC-12/STM-4
	// * `sonet-oc48` - OC-48/STM-16
	// * `sonet-oc192` - OC-192/STM-64
	// * `sonet-oc768` - OC-768/STM-256
	// * `sonet-oc1920` - OC-1920/STM-640
	// * `sonet-oc3840` - OC-3840/STM-1234
	// * `1gfc-sfp` - SFP (1GFC)
	// * `2gfc-sfp` - SFP (2GFC)
	// * `4gfc-sfp` - SFP (4GFC)
	// * `8gfc-sfpp` - SFP+ (8GFC)
	// * `16gfc-sfpp` - SFP+ (16GFC)
	// * `32gfc-sfp28` - SFP28 (32GFC)
	// * `64gfc-qsfpp` - QSFP+ (64GFC)
	// * `128gfc-qsfp28` - QSFP28 (128GFC)
	// * `infiniband-sdr` - SDR (2 Gbps)
	// * `infiniband-ddr` - DDR (4 Gbps)
	// * `infiniband-qdr` - QDR (8 Gbps)
	// * `infiniband-fdr10` - FDR10 (10 Gbps)
	// * `infiniband-fdr` - FDR (13.5 Gbps)
	// * `infiniband-edr` - EDR (25 Gbps)
	// * `infiniband-hdr` - HDR (50 Gbps)
	// * `infiniband-ndr` - NDR (100 Gbps)
	// * `infiniband-xdr` - XDR (250 Gbps)
	// * `t1` - T1 (1.544 Mbps)
	// * `e1` - E1 (2.048 Mbps)
	// * `t3` - T3 (45 Mbps)
	// * `e3` - E3 (34 Mbps)
	// * `xdsl` - xDSL
	// * `docsis` - DOCSIS
	// * `gpon` - GPON (2.5 Gbps / 1.25 Gps)
	// * `xg-pon` - XG-PON (10 Gbps / 2.5 Gbps)
	// * `xgs-pon` - XGS-PON (10 Gbps)
	// * `ng-pon2` - NG-PON2 (TWDM-PON) (4x10 Gbps)
	// * `epon` - EPON (1 Gbps)
	// * `10g-epon` - 10G-EPON (10 Gbps)
	// * `cisco-stackwise` - Cisco StackWise
	// * `cisco-stackwise-plus` - Cisco StackWise Plus
	// * `cisco-flexstack` - Cisco FlexStack
	// * `cisco-flexstack-plus` - Cisco FlexStack Plus
	// * `cisco-stackwise-80` - Cisco StackWise-80
	// * `cisco-stackwise-160` - Cisco StackWise-160
	// * `cisco-stackwise-320` - Cisco StackWise-320
	// * `cisco-stackwise-480` - Cisco StackWise-480
	// * `cisco-stackwise-1t` - Cisco StackWise-1T
	// * `juniper-vcp` - Juniper VCP
	// * `extreme-summitstack` - Extreme SummitStack
	// * `extreme-summitstack-128` - Extreme SummitStack-128
	// * `extreme-summitstack-256` - Extreme SummitStack-256
	// * `extreme-summitstack-512` - Extreme SummitStack-512
	// * `other` - Other
	Type []string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// TypeN: * `virtual` - Virtual
	// * `bridge` - Bridge
	// * `lag` - Link Aggregation Group (LAG)
	// * `100base-fx` - 100BASE-FX (10/100ME FIBER)
	// * `100base-lfx` - 100BASE-LFX (10/100ME FIBER)
	// * `100base-tx` - 100BASE-TX (10/100ME)
	// * `100base-t1` - 100BASE-T1 (10/100ME Single Pair)
	// * `1000base-t` - 1000BASE-T (1GE)
	// * `2.5gbase-t` - 2.5GBASE-T (2.5GE)
	// * `5gbase-t` - 5GBASE-T (5GE)
	// * `10gbase-t` - 10GBASE-T (10GE)
	// * `10gbase-cx4` - 10GBASE-CX4 (10GE)
	// * `1000base-x-gbic` - GBIC (1GE)
	// * `1000base-x-sfp` - SFP (1GE)
	// * `10gbase-x-sfpp` - SFP+ (10GE)
	// * `10gbase-x-xfp` - XFP (10GE)
	// * `10gbase-x-xenpak` - XENPAK (10GE)
	// * `10gbase-x-x2` - X2 (10GE)
	// * `25gbase-x-sfp28` - SFP28 (25GE)
	// * `50gbase-x-sfp56` - SFP56 (50GE)
	// * `40gbase-x-qsfpp` - QSFP+ (40GE)
	// * `50gbase-x-sfp28` - QSFP28 (50GE)
	// * `100gbase-x-cfp` - CFP (100GE)
	// * `100gbase-x-cfp2` - CFP2 (100GE)
	// * `200gbase-x-cfp2` - CFP2 (200GE)
	// * `100gbase-x-cfp4` - CFP4 (100GE)
	// * `100gbase-x-cpak` - Cisco CPAK (100GE)
	// * `100gbase-x-qsfp28` - QSFP28 (100GE)
	// * `200gbase-x-qsfp56` - QSFP56 (200GE)
	// * `400gbase-x-qsfpdd` - QSFP-DD (400GE)
	// * `400gbase-x-osfp` - OSFP (400GE)
	// * `800gbase-x-qsfpdd` - QSFP-DD (800GE)
	// * `800gbase-x-osfp` - OSFP (800GE)
	// * `1000base-kx` - 1000BASE-KX (1GE)
	// * `10gbase-kr` - 10GBASE-KR (10GE)
	// * `10gbase-kx4` - 10GBASE-KX4 (10GE)
	// * `25gbase-kr` - 25GBASE-KR (25GE)
	// * `40gbase-kr4` - 40GBASE-KR4 (40GE)
	// * `50gbase-kr` - 50GBASE-KR (50GE)
	// * `100gbase-kp4` - 100GBASE-KP4 (100GE)
	// * `100gbase-kr2` - 100GBASE-KR2 (100GE)
	// * `100gbase-kr4` - 100GBASE-KR4 (100GE)
	// * `ieee802.11a` - IEEE 802.11a
	// * `ieee802.11g` - IEEE 802.11b/g
	// * `ieee802.11n` - IEEE 802.11n
	// * `ieee802.11ac` - IEEE 802.11ac
	// * `ieee802.11ad` - IEEE 802.11ad
	// * `ieee802.11ax` - IEEE 802.11ax
	// * `ieee802.11ay` - IEEE 802.11ay
	// * `ieee802.15.1` - IEEE 802.15.1 (Bluetooth)
	// * `other-wireless` - Other (Wireless)
	// * `gsm` - GSM
	// * `cdma` - CDMA
	// * `lte` - LTE
	// * `sonet-oc3` - OC-3/STM-1
	// * `sonet-oc12` - OC-12/STM-4
	// * `sonet-oc48` - OC-48/STM-16
	// * `sonet-oc192` - OC-192/STM-64
	// * `sonet-oc768` - OC-768/STM-256
	// * `sonet-oc1920` - OC-1920/STM-640
	// * `sonet-oc3840` - OC-3840/STM-1234
	// * `1gfc-sfp` - SFP (1GFC)
	// * `2gfc-sfp` - SFP (2GFC)
	// * `4gfc-sfp` - SFP (4GFC)
	// * `8gfc-sfpp` - SFP+ (8GFC)
	// * `16gfc-sfpp` - SFP+ (16GFC)
	// * `32gfc-sfp28` - SFP28 (32GFC)
	// * `64gfc-qsfpp` - QSFP+ (64GFC)
	// * `128gfc-qsfp28` - QSFP28 (128GFC)
	// * `infiniband-sdr` - SDR (2 Gbps)
	// * `infiniband-ddr` - DDR (4 Gbps)
	// * `infiniband-qdr` - QDR (8 Gbps)
	// * `infiniband-fdr10` - FDR10 (10 Gbps)
	// * `infiniband-fdr` - FDR (13.5 Gbps)
	// * `infiniband-edr` - EDR (25 Gbps)
	// * `infiniband-hdr` - HDR (50 Gbps)
	// * `infiniband-ndr` - NDR (100 Gbps)
	// * `infiniband-xdr` - XDR (250 Gbps)
	// * `t1` - T1 (1.544 Mbps)
	// * `e1` - E1 (2.048 Mbps)
	// * `t3` - T3 (45 Mbps)
	// * `e3` - E3 (34 Mbps)
	// * `xdsl` - xDSL
	// * `docsis` - DOCSIS
	// * `gpon` - GPON (2.5 Gbps / 1.25 Gps)
	// * `xg-pon` - XG-PON (10 Gbps / 2.5 Gbps)
	// * `xgs-pon` - XGS-PON (10 Gbps)
	// * `ng-pon2` - NG-PON2 (TWDM-PON) (4x10 Gbps)
	// * `epon` - EPON (1 Gbps)
	// * `10g-epon` - 10G-EPON (10 Gbps)
	// * `cisco-stackwise` - Cisco StackWise
	// * `cisco-stackwise-plus` - Cisco StackWise Plus
	// * `cisco-flexstack` - Cisco FlexStack
	// * `cisco-flexstack-plus` - Cisco FlexStack Plus
	// * `cisco-stackwise-80` - Cisco StackWise-80
	// * `cisco-stackwise-160` - Cisco StackWise-160
	// * `cisco-stackwise-320` - Cisco StackWise-320
	// * `cisco-stackwise-480` - Cisco StackWise-480
	// * `cisco-stackwise-1t` - Cisco StackWise-1T
	// * `juniper-vcp` - Juniper VCP
	// * `extreme-summitstack` - Extreme SummitStack
	// * `extreme-summitstack-128` - Extreme SummitStack-128
	// * `extreme-summitstack-256` - Extreme SummitStack-256
	// * `extreme-summitstack-512` - Extreme SummitStack-512
	// * `other` - Other
	TypeN []string `json:"type__n,omitempty" mapstructure:"type__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// Vdc: Virtual Device Context
	Vdc []string `json:"vdc,omitempty" mapstructure:"vdc,omitempty"`
	// VdcN: Virtual Device Context
	VdcN []string `json:"vdc__n,omitempty" mapstructure:"vdc__n,omitempty"`
	// VdcId: Virtual Device Context
	VdcId []int32 `json:"vdc_id,omitempty" mapstructure:"vdc_id,omitempty"`
	// VdcIdN: Virtual Device Context
	VdcIdN []int32 `json:"vdc_id__n,omitempty" mapstructure:"vdc_id__n,omitempty"`
	// VdcIdentifier: Virtual Device Context (Identifier)
	VdcIdentifier []*int32 `json:"vdc_identifier,omitempty" mapstructure:"vdc_identifier,omitempty"`
	// VdcIdentifierN: Virtual Device Context (Identifier)
	VdcIdentifierN []*int32 `json:"vdc_identifier__n,omitempty" mapstructure:"vdc_identifier__n,omitempty"`
	// VirtualChassis: Virtual Chassis
	VirtualChassis []string `json:"virtual_chassis,omitempty" mapstructure:"virtual_chassis,omitempty"`
	// VirtualChassisN: Virtual Chassis
	VirtualChassisN []string `json:"virtual_chassis__n,omitempty" mapstructure:"virtual_chassis__n,omitempty"`
	// VirtualChassisId: Virtual Chassis (ID)
	VirtualChassisId []int32 `json:"virtual_chassis_id,omitempty" mapstructure:"virtual_chassis_id,omitempty"`
	// VirtualChassisIdN: Virtual Chassis (ID)
	VirtualChassisIdN []int32 `json:"virtual_chassis_id__n,omitempty" mapstructure:"virtual_chassis_id__n,omitempty"`
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
	// Wwn:
	Wwn []string `json:"wwn,omitempty" mapstructure:"wwn,omitempty"`
	// WwnIc:
	WwnIc []string `json:"wwn__ic,omitempty" mapstructure:"wwn__ic,omitempty"`
	// WwnIe:
	WwnIe []string `json:"wwn__ie,omitempty" mapstructure:"wwn__ie,omitempty"`
	// WwnIew:
	WwnIew []string `json:"wwn__iew,omitempty" mapstructure:"wwn__iew,omitempty"`
	// WwnIsw:
	WwnIsw []string `json:"wwn__isw,omitempty" mapstructure:"wwn__isw,omitempty"`
	// WwnN:
	WwnN []string `json:"wwn__n,omitempty" mapstructure:"wwn__n,omitempty"`
	// WwnNic:
	WwnNic []string `json:"wwn__nic,omitempty" mapstructure:"wwn__nic,omitempty"`
	// WwnNie:
	WwnNie []string `json:"wwn__nie,omitempty" mapstructure:"wwn__nie,omitempty"`
	// WwnNiew:
	WwnNiew []string `json:"wwn__niew,omitempty" mapstructure:"wwn__niew,omitempty"`
	// WwnNisw:
	WwnNisw []string `json:"wwn__nisw,omitempty" mapstructure:"wwn__nisw,omitempty"`
}

// Validate implements basic validation for this model
func (m DcimInterfacesListQueryParameters) Validate() error {
	return validation.Errors{
		"bridgeId": validation.Validate(
			m.BridgeId,
		),
		"bridgeIdN": validation.Validate(
			m.BridgeIdN,
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
		"duplex": validation.Validate(
			m.Duplex,
		),
		"duplexN": validation.Validate(
			m.DuplexN,
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
		"lagId": validation.Validate(
			m.LagId,
		),
		"lagIdN": validation.Validate(
			m.LagIdN,
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
		"location": validation.Validate(
			m.Location,
		),
		"locationN": validation.Validate(
			m.LocationN,
		),
		"locationId": validation.Validate(
			m.LocationId,
		),
		"locationIdN": validation.Validate(
			m.LocationIdN,
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
		"moduleId": validation.Validate(
			m.ModuleId,
		),
		"moduleIdN": validation.Validate(
			m.ModuleIdN,
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
		"poeMode": validation.Validate(
			m.PoeMode,
		),
		"poeModeN": validation.Validate(
			m.PoeModeN,
		),
		"poeType": validation.Validate(
			m.PoeType,
		),
		"poeTypeN": validation.Validate(
			m.PoeTypeN,
		),
		"rack": validation.Validate(
			m.Rack,
		),
		"rackN": validation.Validate(
			m.RackN,
		),
		"rackId": validation.Validate(
			m.RackId,
		),
		"rackIdN": validation.Validate(
			m.RackIdN,
		),
		"region": validation.Validate(
			m.Region,
		),
		"regionN": validation.Validate(
			m.RegionN,
		),
		"regionId": validation.Validate(
			m.RegionId,
		),
		"regionIdN": validation.Validate(
			m.RegionIdN,
		),
		"rfChannel": validation.Validate(
			m.RfChannel,
		),
		"rfChannelN": validation.Validate(
			m.RfChannelN,
		),
		"rfChannelFrequency": validation.Validate(
			m.RfChannelFrequency,
		),
		"rfChannelFrequencyGt": validation.Validate(
			m.RfChannelFrequencyGt,
		),
		"rfChannelFrequencyGte": validation.Validate(
			m.RfChannelFrequencyGte,
		),
		"rfChannelFrequencyLt": validation.Validate(
			m.RfChannelFrequencyLt,
		),
		"rfChannelFrequencyLte": validation.Validate(
			m.RfChannelFrequencyLte,
		),
		"rfChannelFrequencyN": validation.Validate(
			m.RfChannelFrequencyN,
		),
		"rfChannelWidth": validation.Validate(
			m.RfChannelWidth,
		),
		"rfChannelWidthGt": validation.Validate(
			m.RfChannelWidthGt,
		),
		"rfChannelWidthGte": validation.Validate(
			m.RfChannelWidthGte,
		),
		"rfChannelWidthLt": validation.Validate(
			m.RfChannelWidthLt,
		),
		"rfChannelWidthLte": validation.Validate(
			m.RfChannelWidthLte,
		),
		"rfChannelWidthN": validation.Validate(
			m.RfChannelWidthN,
		),
		"rfRole": validation.Validate(
			m.RfRole,
		),
		"rfRoleN": validation.Validate(
			m.RfRoleN,
		),
		"site": validation.Validate(
			m.Site,
		),
		"siteN": validation.Validate(
			m.SiteN,
		),
		"siteGroup": validation.Validate(
			m.SiteGroup,
		),
		"siteGroupN": validation.Validate(
			m.SiteGroupN,
		),
		"siteGroupId": validation.Validate(
			m.SiteGroupId,
		),
		"siteGroupIdN": validation.Validate(
			m.SiteGroupIdN,
		),
		"siteId": validation.Validate(
			m.SiteId,
		),
		"siteIdN": validation.Validate(
			m.SiteIdN,
		),
		"speed": validation.Validate(
			m.Speed,
		),
		"speedGt": validation.Validate(
			m.SpeedGt,
		),
		"speedGte": validation.Validate(
			m.SpeedGte,
		),
		"speedLt": validation.Validate(
			m.SpeedLt,
		),
		"speedLte": validation.Validate(
			m.SpeedLte,
		),
		"speedN": validation.Validate(
			m.SpeedN,
		),
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
		"txPower": validation.Validate(
			m.TxPower,
		),
		"txPowerGt": validation.Validate(
			m.TxPowerGt,
		),
		"txPowerGte": validation.Validate(
			m.TxPowerGte,
		),
		"txPowerLt": validation.Validate(
			m.TxPowerLt,
		),
		"txPowerLte": validation.Validate(
			m.TxPowerLte,
		),
		"txPowerN": validation.Validate(
			m.TxPowerN,
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
		"vdc": validation.Validate(
			m.Vdc,
		),
		"vdcN": validation.Validate(
			m.VdcN,
		),
		"vdcId": validation.Validate(
			m.VdcId,
		),
		"vdcIdN": validation.Validate(
			m.VdcIdN,
		),
		"vdcIdentifier": validation.Validate(
			m.VdcIdentifier,
		),
		"vdcIdentifierN": validation.Validate(
			m.VdcIdentifierN,
		),
		"virtualChassis": validation.Validate(
			m.VirtualChassis,
		),
		"virtualChassisN": validation.Validate(
			m.VirtualChassisN,
		),
		"virtualChassisId": validation.Validate(
			m.VirtualChassisId,
		),
		"virtualChassisIdN": validation.Validate(
			m.VirtualChassisIdN,
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
		"wwn": validation.Validate(
			m.Wwn,
		),
		"wwnIc": validation.Validate(
			m.WwnIc,
		),
		"wwnIe": validation.Validate(
			m.WwnIe,
		),
		"wwnIew": validation.Validate(
			m.WwnIew,
		),
		"wwnIsw": validation.Validate(
			m.WwnIsw,
		),
		"wwnN": validation.Validate(
			m.WwnN,
		),
		"wwnNic": validation.Validate(
			m.WwnNic,
		),
		"wwnNie": validation.Validate(
			m.WwnNie,
		),
		"wwnNiew": validation.Validate(
			m.WwnNiew,
		),
		"wwnNisw": validation.Validate(
			m.WwnNisw,
		),
	}.Filter()
}

// GetBridgeId returns the BridgeId property
func (m DcimInterfacesListQueryParameters) GetBridgeId() []int32 {
	return m.BridgeId
}

// SetBridgeId sets the BridgeId property
func (m *DcimInterfacesListQueryParameters) SetBridgeId(val []int32) {
	m.BridgeId = val
}

// GetBridgeIdN returns the BridgeIdN property
func (m DcimInterfacesListQueryParameters) GetBridgeIdN() []int32 {
	return m.BridgeIdN
}

// SetBridgeIdN sets the BridgeIdN property
func (m *DcimInterfacesListQueryParameters) SetBridgeIdN(val []int32) {
	m.BridgeIdN = val
}

// GetCableEnd returns the CableEnd property
func (m DcimInterfacesListQueryParameters) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *DcimInterfacesListQueryParameters) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCableEndN returns the CableEndN property
func (m DcimInterfacesListQueryParameters) GetCableEndN() string {
	return m.CableEndN
}

// SetCableEndN sets the CableEndN property
func (m *DcimInterfacesListQueryParameters) SetCableEndN(val string) {
	m.CableEndN = val
}

// GetCabled returns the Cabled property
func (m DcimInterfacesListQueryParameters) GetCabled() bool {
	return m.Cabled
}

// SetCabled sets the Cabled property
func (m *DcimInterfacesListQueryParameters) SetCabled(val bool) {
	m.Cabled = val
}

// GetConnected returns the Connected property
func (m DcimInterfacesListQueryParameters) GetConnected() bool {
	return m.Connected
}

// SetConnected sets the Connected property
func (m *DcimInterfacesListQueryParameters) SetConnected(val bool) {
	m.Connected = val
}

// GetCreated returns the Created property
func (m DcimInterfacesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimInterfacesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimInterfacesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimInterfacesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimInterfacesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimInterfacesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimInterfacesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimInterfacesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimInterfacesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimInterfacesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimInterfacesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimInterfacesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimInterfacesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimInterfacesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDescription returns the Description property
func (m DcimInterfacesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *DcimInterfacesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m DcimInterfacesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *DcimInterfacesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m DcimInterfacesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *DcimInterfacesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m DcimInterfacesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *DcimInterfacesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m DcimInterfacesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *DcimInterfacesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m DcimInterfacesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *DcimInterfacesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m DcimInterfacesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *DcimInterfacesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m DcimInterfacesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *DcimInterfacesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m DcimInterfacesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *DcimInterfacesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m DcimInterfacesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *DcimInterfacesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m DcimInterfacesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *DcimInterfacesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetDevice returns the Device property
func (m DcimInterfacesListQueryParameters) GetDevice() []string {
	return m.Device
}

// SetDevice sets the Device property
func (m *DcimInterfacesListQueryParameters) SetDevice(val []string) {
	m.Device = val
}

// GetDeviceId returns the DeviceId property
func (m DcimInterfacesListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *DcimInterfacesListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetDuplex returns the Duplex property
func (m DcimInterfacesListQueryParameters) GetDuplex() []*string {
	return m.Duplex
}

// SetDuplex sets the Duplex property
func (m *DcimInterfacesListQueryParameters) SetDuplex(val []*string) {
	m.Duplex = val
}

// GetDuplexN returns the DuplexN property
func (m DcimInterfacesListQueryParameters) GetDuplexN() []*string {
	return m.DuplexN
}

// SetDuplexN sets the DuplexN property
func (m *DcimInterfacesListQueryParameters) SetDuplexN(val []*string) {
	m.DuplexN = val
}

// GetEnabled returns the Enabled property
func (m DcimInterfacesListQueryParameters) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *DcimInterfacesListQueryParameters) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m DcimInterfacesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimInterfacesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimInterfacesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimInterfacesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimInterfacesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimInterfacesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimInterfacesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimInterfacesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimInterfacesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimInterfacesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimInterfacesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimInterfacesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetKind returns the Kind property
func (m DcimInterfacesListQueryParameters) GetKind() string {
	return m.Kind
}

// SetKind sets the Kind property
func (m *DcimInterfacesListQueryParameters) SetKind(val string) {
	m.Kind = val
}

// GetL2vpn returns the L2vpn property
func (m DcimInterfacesListQueryParameters) GetL2vpn() []*int64 {
	return m.L2vpn
}

// SetL2vpn sets the L2vpn property
func (m *DcimInterfacesListQueryParameters) SetL2vpn(val []*int64) {
	m.L2vpn = val
}

// GetL2vpnN returns the L2vpnN property
func (m DcimInterfacesListQueryParameters) GetL2vpnN() []*int64 {
	return m.L2vpnN
}

// SetL2vpnN sets the L2vpnN property
func (m *DcimInterfacesListQueryParameters) SetL2vpnN(val []*int64) {
	m.L2vpnN = val
}

// GetL2vpnId returns the L2vpnId property
func (m DcimInterfacesListQueryParameters) GetL2vpnId() []int32 {
	return m.L2vpnId
}

// SetL2vpnId sets the L2vpnId property
func (m *DcimInterfacesListQueryParameters) SetL2vpnId(val []int32) {
	m.L2vpnId = val
}

// GetL2vpnIdN returns the L2vpnIdN property
func (m DcimInterfacesListQueryParameters) GetL2vpnIdN() []int32 {
	return m.L2vpnIdN
}

// SetL2vpnIdN sets the L2vpnIdN property
func (m *DcimInterfacesListQueryParameters) SetL2vpnIdN(val []int32) {
	m.L2vpnIdN = val
}

// GetLabel returns the Label property
func (m DcimInterfacesListQueryParameters) GetLabel() []string {
	return m.Label
}

// SetLabel sets the Label property
func (m *DcimInterfacesListQueryParameters) SetLabel(val []string) {
	m.Label = val
}

// GetLabelEmpty returns the LabelEmpty property
func (m DcimInterfacesListQueryParameters) GetLabelEmpty() []string {
	return m.LabelEmpty
}

// SetLabelEmpty sets the LabelEmpty property
func (m *DcimInterfacesListQueryParameters) SetLabelEmpty(val []string) {
	m.LabelEmpty = val
}

// GetLabelIc returns the LabelIc property
func (m DcimInterfacesListQueryParameters) GetLabelIc() []string {
	return m.LabelIc
}

// SetLabelIc sets the LabelIc property
func (m *DcimInterfacesListQueryParameters) SetLabelIc(val []string) {
	m.LabelIc = val
}

// GetLabelIe returns the LabelIe property
func (m DcimInterfacesListQueryParameters) GetLabelIe() []string {
	return m.LabelIe
}

// SetLabelIe sets the LabelIe property
func (m *DcimInterfacesListQueryParameters) SetLabelIe(val []string) {
	m.LabelIe = val
}

// GetLabelIew returns the LabelIew property
func (m DcimInterfacesListQueryParameters) GetLabelIew() []string {
	return m.LabelIew
}

// SetLabelIew sets the LabelIew property
func (m *DcimInterfacesListQueryParameters) SetLabelIew(val []string) {
	m.LabelIew = val
}

// GetLabelIsw returns the LabelIsw property
func (m DcimInterfacesListQueryParameters) GetLabelIsw() []string {
	return m.LabelIsw
}

// SetLabelIsw sets the LabelIsw property
func (m *DcimInterfacesListQueryParameters) SetLabelIsw(val []string) {
	m.LabelIsw = val
}

// GetLabelN returns the LabelN property
func (m DcimInterfacesListQueryParameters) GetLabelN() []string {
	return m.LabelN
}

// SetLabelN sets the LabelN property
func (m *DcimInterfacesListQueryParameters) SetLabelN(val []string) {
	m.LabelN = val
}

// GetLabelNic returns the LabelNic property
func (m DcimInterfacesListQueryParameters) GetLabelNic() []string {
	return m.LabelNic
}

// SetLabelNic sets the LabelNic property
func (m *DcimInterfacesListQueryParameters) SetLabelNic(val []string) {
	m.LabelNic = val
}

// GetLabelNie returns the LabelNie property
func (m DcimInterfacesListQueryParameters) GetLabelNie() []string {
	return m.LabelNie
}

// SetLabelNie sets the LabelNie property
func (m *DcimInterfacesListQueryParameters) SetLabelNie(val []string) {
	m.LabelNie = val
}

// GetLabelNiew returns the LabelNiew property
func (m DcimInterfacesListQueryParameters) GetLabelNiew() []string {
	return m.LabelNiew
}

// SetLabelNiew sets the LabelNiew property
func (m *DcimInterfacesListQueryParameters) SetLabelNiew(val []string) {
	m.LabelNiew = val
}

// GetLabelNisw returns the LabelNisw property
func (m DcimInterfacesListQueryParameters) GetLabelNisw() []string {
	return m.LabelNisw
}

// SetLabelNisw sets the LabelNisw property
func (m *DcimInterfacesListQueryParameters) SetLabelNisw(val []string) {
	m.LabelNisw = val
}

// GetLagId returns the LagId property
func (m DcimInterfacesListQueryParameters) GetLagId() []int32 {
	return m.LagId
}

// SetLagId sets the LagId property
func (m *DcimInterfacesListQueryParameters) SetLagId(val []int32) {
	m.LagId = val
}

// GetLagIdN returns the LagIdN property
func (m DcimInterfacesListQueryParameters) GetLagIdN() []int32 {
	return m.LagIdN
}

// SetLagIdN sets the LagIdN property
func (m *DcimInterfacesListQueryParameters) SetLagIdN(val []int32) {
	m.LagIdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimInterfacesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimInterfacesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimInterfacesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimInterfacesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimInterfacesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimInterfacesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimInterfacesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimInterfacesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimInterfacesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimInterfacesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimInterfacesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimInterfacesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimInterfacesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimInterfacesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetLocation returns the Location property
func (m DcimInterfacesListQueryParameters) GetLocation() []string {
	return m.Location
}

// SetLocation sets the Location property
func (m *DcimInterfacesListQueryParameters) SetLocation(val []string) {
	m.Location = val
}

// GetLocationN returns the LocationN property
func (m DcimInterfacesListQueryParameters) GetLocationN() []string {
	return m.LocationN
}

// SetLocationN sets the LocationN property
func (m *DcimInterfacesListQueryParameters) SetLocationN(val []string) {
	m.LocationN = val
}

// GetLocationId returns the LocationId property
func (m DcimInterfacesListQueryParameters) GetLocationId() []int32 {
	return m.LocationId
}

// SetLocationId sets the LocationId property
func (m *DcimInterfacesListQueryParameters) SetLocationId(val []int32) {
	m.LocationId = val
}

// GetLocationIdN returns the LocationIdN property
func (m DcimInterfacesListQueryParameters) GetLocationIdN() []int32 {
	return m.LocationIdN
}

// SetLocationIdN sets the LocationIdN property
func (m *DcimInterfacesListQueryParameters) SetLocationIdN(val []int32) {
	m.LocationIdN = val
}

// GetMacAddress returns the MacAddress property
func (m DcimInterfacesListQueryParameters) GetMacAddress() []string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *DcimInterfacesListQueryParameters) SetMacAddress(val []string) {
	m.MacAddress = val
}

// GetMacAddressIc returns the MacAddressIc property
func (m DcimInterfacesListQueryParameters) GetMacAddressIc() []string {
	return m.MacAddressIc
}

// SetMacAddressIc sets the MacAddressIc property
func (m *DcimInterfacesListQueryParameters) SetMacAddressIc(val []string) {
	m.MacAddressIc = val
}

// GetMacAddressIe returns the MacAddressIe property
func (m DcimInterfacesListQueryParameters) GetMacAddressIe() []string {
	return m.MacAddressIe
}

// SetMacAddressIe sets the MacAddressIe property
func (m *DcimInterfacesListQueryParameters) SetMacAddressIe(val []string) {
	m.MacAddressIe = val
}

// GetMacAddressIew returns the MacAddressIew property
func (m DcimInterfacesListQueryParameters) GetMacAddressIew() []string {
	return m.MacAddressIew
}

// SetMacAddressIew sets the MacAddressIew property
func (m *DcimInterfacesListQueryParameters) SetMacAddressIew(val []string) {
	m.MacAddressIew = val
}

// GetMacAddressIsw returns the MacAddressIsw property
func (m DcimInterfacesListQueryParameters) GetMacAddressIsw() []string {
	return m.MacAddressIsw
}

// SetMacAddressIsw sets the MacAddressIsw property
func (m *DcimInterfacesListQueryParameters) SetMacAddressIsw(val []string) {
	m.MacAddressIsw = val
}

// GetMacAddressN returns the MacAddressN property
func (m DcimInterfacesListQueryParameters) GetMacAddressN() []string {
	return m.MacAddressN
}

// SetMacAddressN sets the MacAddressN property
func (m *DcimInterfacesListQueryParameters) SetMacAddressN(val []string) {
	m.MacAddressN = val
}

// GetMacAddressNic returns the MacAddressNic property
func (m DcimInterfacesListQueryParameters) GetMacAddressNic() []string {
	return m.MacAddressNic
}

// SetMacAddressNic sets the MacAddressNic property
func (m *DcimInterfacesListQueryParameters) SetMacAddressNic(val []string) {
	m.MacAddressNic = val
}

// GetMacAddressNie returns the MacAddressNie property
func (m DcimInterfacesListQueryParameters) GetMacAddressNie() []string {
	return m.MacAddressNie
}

// SetMacAddressNie sets the MacAddressNie property
func (m *DcimInterfacesListQueryParameters) SetMacAddressNie(val []string) {
	m.MacAddressNie = val
}

// GetMacAddressNiew returns the MacAddressNiew property
func (m DcimInterfacesListQueryParameters) GetMacAddressNiew() []string {
	return m.MacAddressNiew
}

// SetMacAddressNiew sets the MacAddressNiew property
func (m *DcimInterfacesListQueryParameters) SetMacAddressNiew(val []string) {
	m.MacAddressNiew = val
}

// GetMacAddressNisw returns the MacAddressNisw property
func (m DcimInterfacesListQueryParameters) GetMacAddressNisw() []string {
	return m.MacAddressNisw
}

// SetMacAddressNisw sets the MacAddressNisw property
func (m *DcimInterfacesListQueryParameters) SetMacAddressNisw(val []string) {
	m.MacAddressNisw = val
}

// GetMgmtOnly returns the MgmtOnly property
func (m DcimInterfacesListQueryParameters) GetMgmtOnly() bool {
	return m.MgmtOnly
}

// SetMgmtOnly sets the MgmtOnly property
func (m *DcimInterfacesListQueryParameters) SetMgmtOnly(val bool) {
	m.MgmtOnly = val
}

// GetMode returns the Mode property
func (m DcimInterfacesListQueryParameters) GetMode() string {
	return m.Mode
}

// SetMode sets the Mode property
func (m *DcimInterfacesListQueryParameters) SetMode(val string) {
	m.Mode = val
}

// GetModeN returns the ModeN property
func (m DcimInterfacesListQueryParameters) GetModeN() string {
	return m.ModeN
}

// SetModeN sets the ModeN property
func (m *DcimInterfacesListQueryParameters) SetModeN(val string) {
	m.ModeN = val
}

// GetModuleId returns the ModuleId property
func (m DcimInterfacesListQueryParameters) GetModuleId() []*int32 {
	return m.ModuleId
}

// SetModuleId sets the ModuleId property
func (m *DcimInterfacesListQueryParameters) SetModuleId(val []*int32) {
	m.ModuleId = val
}

// GetModuleIdN returns the ModuleIdN property
func (m DcimInterfacesListQueryParameters) GetModuleIdN() []*int32 {
	return m.ModuleIdN
}

// SetModuleIdN sets the ModuleIdN property
func (m *DcimInterfacesListQueryParameters) SetModuleIdN(val []*int32) {
	m.ModuleIdN = val
}

// GetMtu returns the Mtu property
func (m DcimInterfacesListQueryParameters) GetMtu() []int32 {
	return m.Mtu
}

// SetMtu sets the Mtu property
func (m *DcimInterfacesListQueryParameters) SetMtu(val []int32) {
	m.Mtu = val
}

// GetMtuGt returns the MtuGt property
func (m DcimInterfacesListQueryParameters) GetMtuGt() []int32 {
	return m.MtuGt
}

// SetMtuGt sets the MtuGt property
func (m *DcimInterfacesListQueryParameters) SetMtuGt(val []int32) {
	m.MtuGt = val
}

// GetMtuGte returns the MtuGte property
func (m DcimInterfacesListQueryParameters) GetMtuGte() []int32 {
	return m.MtuGte
}

// SetMtuGte sets the MtuGte property
func (m *DcimInterfacesListQueryParameters) SetMtuGte(val []int32) {
	m.MtuGte = val
}

// GetMtuLt returns the MtuLt property
func (m DcimInterfacesListQueryParameters) GetMtuLt() []int32 {
	return m.MtuLt
}

// SetMtuLt sets the MtuLt property
func (m *DcimInterfacesListQueryParameters) SetMtuLt(val []int32) {
	m.MtuLt = val
}

// GetMtuLte returns the MtuLte property
func (m DcimInterfacesListQueryParameters) GetMtuLte() []int32 {
	return m.MtuLte
}

// SetMtuLte sets the MtuLte property
func (m *DcimInterfacesListQueryParameters) SetMtuLte(val []int32) {
	m.MtuLte = val
}

// GetMtuN returns the MtuN property
func (m DcimInterfacesListQueryParameters) GetMtuN() []int32 {
	return m.MtuN
}

// SetMtuN sets the MtuN property
func (m *DcimInterfacesListQueryParameters) SetMtuN(val []int32) {
	m.MtuN = val
}

// GetName returns the Name property
func (m DcimInterfacesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimInterfacesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimInterfacesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimInterfacesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimInterfacesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimInterfacesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimInterfacesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimInterfacesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimInterfacesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimInterfacesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimInterfacesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimInterfacesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimInterfacesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimInterfacesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimInterfacesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimInterfacesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimInterfacesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimInterfacesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimInterfacesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimInterfacesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimInterfacesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimInterfacesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOccupied returns the Occupied property
func (m DcimInterfacesListQueryParameters) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *DcimInterfacesListQueryParameters) SetOccupied(val bool) {
	m.Occupied = val
}

// GetOffset returns the Offset property
func (m DcimInterfacesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimInterfacesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimInterfacesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimInterfacesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetParentId returns the ParentId property
func (m DcimInterfacesListQueryParameters) GetParentId() []int32 {
	return m.ParentId
}

// SetParentId sets the ParentId property
func (m *DcimInterfacesListQueryParameters) SetParentId(val []int32) {
	m.ParentId = val
}

// GetParentIdN returns the ParentIdN property
func (m DcimInterfacesListQueryParameters) GetParentIdN() []int32 {
	return m.ParentIdN
}

// SetParentIdN sets the ParentIdN property
func (m *DcimInterfacesListQueryParameters) SetParentIdN(val []int32) {
	m.ParentIdN = val
}

// GetPoeMode returns the PoeMode property
func (m DcimInterfacesListQueryParameters) GetPoeMode() []string {
	return m.PoeMode
}

// SetPoeMode sets the PoeMode property
func (m *DcimInterfacesListQueryParameters) SetPoeMode(val []string) {
	m.PoeMode = val
}

// GetPoeModeN returns the PoeModeN property
func (m DcimInterfacesListQueryParameters) GetPoeModeN() []string {
	return m.PoeModeN
}

// SetPoeModeN sets the PoeModeN property
func (m *DcimInterfacesListQueryParameters) SetPoeModeN(val []string) {
	m.PoeModeN = val
}

// GetPoeType returns the PoeType property
func (m DcimInterfacesListQueryParameters) GetPoeType() []string {
	return m.PoeType
}

// SetPoeType sets the PoeType property
func (m *DcimInterfacesListQueryParameters) SetPoeType(val []string) {
	m.PoeType = val
}

// GetPoeTypeN returns the PoeTypeN property
func (m DcimInterfacesListQueryParameters) GetPoeTypeN() []string {
	return m.PoeTypeN
}

// SetPoeTypeN sets the PoeTypeN property
func (m *DcimInterfacesListQueryParameters) SetPoeTypeN(val []string) {
	m.PoeTypeN = val
}

// GetQ returns the Q property
func (m DcimInterfacesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimInterfacesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetRack returns the Rack property
func (m DcimInterfacesListQueryParameters) GetRack() []string {
	return m.Rack
}

// SetRack sets the Rack property
func (m *DcimInterfacesListQueryParameters) SetRack(val []string) {
	m.Rack = val
}

// GetRackN returns the RackN property
func (m DcimInterfacesListQueryParameters) GetRackN() []string {
	return m.RackN
}

// SetRackN sets the RackN property
func (m *DcimInterfacesListQueryParameters) SetRackN(val []string) {
	m.RackN = val
}

// GetRackId returns the RackId property
func (m DcimInterfacesListQueryParameters) GetRackId() []int32 {
	return m.RackId
}

// SetRackId sets the RackId property
func (m *DcimInterfacesListQueryParameters) SetRackId(val []int32) {
	m.RackId = val
}

// GetRackIdN returns the RackIdN property
func (m DcimInterfacesListQueryParameters) GetRackIdN() []int32 {
	return m.RackIdN
}

// SetRackIdN sets the RackIdN property
func (m *DcimInterfacesListQueryParameters) SetRackIdN(val []int32) {
	m.RackIdN = val
}

// GetRegion returns the Region property
func (m DcimInterfacesListQueryParameters) GetRegion() []int32 {
	return m.Region
}

// SetRegion sets the Region property
func (m *DcimInterfacesListQueryParameters) SetRegion(val []int32) {
	m.Region = val
}

// GetRegionN returns the RegionN property
func (m DcimInterfacesListQueryParameters) GetRegionN() []int32 {
	return m.RegionN
}

// SetRegionN sets the RegionN property
func (m *DcimInterfacesListQueryParameters) SetRegionN(val []int32) {
	m.RegionN = val
}

// GetRegionId returns the RegionId property
func (m DcimInterfacesListQueryParameters) GetRegionId() []int32 {
	return m.RegionId
}

// SetRegionId sets the RegionId property
func (m *DcimInterfacesListQueryParameters) SetRegionId(val []int32) {
	m.RegionId = val
}

// GetRegionIdN returns the RegionIdN property
func (m DcimInterfacesListQueryParameters) GetRegionIdN() []int32 {
	return m.RegionIdN
}

// SetRegionIdN sets the RegionIdN property
func (m *DcimInterfacesListQueryParameters) SetRegionIdN(val []int32) {
	m.RegionIdN = val
}

// GetRfChannel returns the RfChannel property
func (m DcimInterfacesListQueryParameters) GetRfChannel() []string {
	return m.RfChannel
}

// SetRfChannel sets the RfChannel property
func (m *DcimInterfacesListQueryParameters) SetRfChannel(val []string) {
	m.RfChannel = val
}

// GetRfChannelN returns the RfChannelN property
func (m DcimInterfacesListQueryParameters) GetRfChannelN() []string {
	return m.RfChannelN
}

// SetRfChannelN sets the RfChannelN property
func (m *DcimInterfacesListQueryParameters) SetRfChannelN(val []string) {
	m.RfChannelN = val
}

// GetRfChannelFrequency returns the RfChannelFrequency property
func (m DcimInterfacesListQueryParameters) GetRfChannelFrequency() []float64 {
	return m.RfChannelFrequency
}

// SetRfChannelFrequency sets the RfChannelFrequency property
func (m *DcimInterfacesListQueryParameters) SetRfChannelFrequency(val []float64) {
	m.RfChannelFrequency = val
}

// GetRfChannelFrequencyGt returns the RfChannelFrequencyGt property
func (m DcimInterfacesListQueryParameters) GetRfChannelFrequencyGt() []float64 {
	return m.RfChannelFrequencyGt
}

// SetRfChannelFrequencyGt sets the RfChannelFrequencyGt property
func (m *DcimInterfacesListQueryParameters) SetRfChannelFrequencyGt(val []float64) {
	m.RfChannelFrequencyGt = val
}

// GetRfChannelFrequencyGte returns the RfChannelFrequencyGte property
func (m DcimInterfacesListQueryParameters) GetRfChannelFrequencyGte() []float64 {
	return m.RfChannelFrequencyGte
}

// SetRfChannelFrequencyGte sets the RfChannelFrequencyGte property
func (m *DcimInterfacesListQueryParameters) SetRfChannelFrequencyGte(val []float64) {
	m.RfChannelFrequencyGte = val
}

// GetRfChannelFrequencyLt returns the RfChannelFrequencyLt property
func (m DcimInterfacesListQueryParameters) GetRfChannelFrequencyLt() []float64 {
	return m.RfChannelFrequencyLt
}

// SetRfChannelFrequencyLt sets the RfChannelFrequencyLt property
func (m *DcimInterfacesListQueryParameters) SetRfChannelFrequencyLt(val []float64) {
	m.RfChannelFrequencyLt = val
}

// GetRfChannelFrequencyLte returns the RfChannelFrequencyLte property
func (m DcimInterfacesListQueryParameters) GetRfChannelFrequencyLte() []float64 {
	return m.RfChannelFrequencyLte
}

// SetRfChannelFrequencyLte sets the RfChannelFrequencyLte property
func (m *DcimInterfacesListQueryParameters) SetRfChannelFrequencyLte(val []float64) {
	m.RfChannelFrequencyLte = val
}

// GetRfChannelFrequencyN returns the RfChannelFrequencyN property
func (m DcimInterfacesListQueryParameters) GetRfChannelFrequencyN() []float64 {
	return m.RfChannelFrequencyN
}

// SetRfChannelFrequencyN sets the RfChannelFrequencyN property
func (m *DcimInterfacesListQueryParameters) SetRfChannelFrequencyN(val []float64) {
	m.RfChannelFrequencyN = val
}

// GetRfChannelWidth returns the RfChannelWidth property
func (m DcimInterfacesListQueryParameters) GetRfChannelWidth() []float64 {
	return m.RfChannelWidth
}

// SetRfChannelWidth sets the RfChannelWidth property
func (m *DcimInterfacesListQueryParameters) SetRfChannelWidth(val []float64) {
	m.RfChannelWidth = val
}

// GetRfChannelWidthGt returns the RfChannelWidthGt property
func (m DcimInterfacesListQueryParameters) GetRfChannelWidthGt() []float64 {
	return m.RfChannelWidthGt
}

// SetRfChannelWidthGt sets the RfChannelWidthGt property
func (m *DcimInterfacesListQueryParameters) SetRfChannelWidthGt(val []float64) {
	m.RfChannelWidthGt = val
}

// GetRfChannelWidthGte returns the RfChannelWidthGte property
func (m DcimInterfacesListQueryParameters) GetRfChannelWidthGte() []float64 {
	return m.RfChannelWidthGte
}

// SetRfChannelWidthGte sets the RfChannelWidthGte property
func (m *DcimInterfacesListQueryParameters) SetRfChannelWidthGte(val []float64) {
	m.RfChannelWidthGte = val
}

// GetRfChannelWidthLt returns the RfChannelWidthLt property
func (m DcimInterfacesListQueryParameters) GetRfChannelWidthLt() []float64 {
	return m.RfChannelWidthLt
}

// SetRfChannelWidthLt sets the RfChannelWidthLt property
func (m *DcimInterfacesListQueryParameters) SetRfChannelWidthLt(val []float64) {
	m.RfChannelWidthLt = val
}

// GetRfChannelWidthLte returns the RfChannelWidthLte property
func (m DcimInterfacesListQueryParameters) GetRfChannelWidthLte() []float64 {
	return m.RfChannelWidthLte
}

// SetRfChannelWidthLte sets the RfChannelWidthLte property
func (m *DcimInterfacesListQueryParameters) SetRfChannelWidthLte(val []float64) {
	m.RfChannelWidthLte = val
}

// GetRfChannelWidthN returns the RfChannelWidthN property
func (m DcimInterfacesListQueryParameters) GetRfChannelWidthN() []float64 {
	return m.RfChannelWidthN
}

// SetRfChannelWidthN sets the RfChannelWidthN property
func (m *DcimInterfacesListQueryParameters) SetRfChannelWidthN(val []float64) {
	m.RfChannelWidthN = val
}

// GetRfRole returns the RfRole property
func (m DcimInterfacesListQueryParameters) GetRfRole() []string {
	return m.RfRole
}

// SetRfRole sets the RfRole property
func (m *DcimInterfacesListQueryParameters) SetRfRole(val []string) {
	m.RfRole = val
}

// GetRfRoleN returns the RfRoleN property
func (m DcimInterfacesListQueryParameters) GetRfRoleN() []string {
	return m.RfRoleN
}

// SetRfRoleN sets the RfRoleN property
func (m *DcimInterfacesListQueryParameters) SetRfRoleN(val []string) {
	m.RfRoleN = val
}

// GetSite returns the Site property
func (m DcimInterfacesListQueryParameters) GetSite() []string {
	return m.Site
}

// SetSite sets the Site property
func (m *DcimInterfacesListQueryParameters) SetSite(val []string) {
	m.Site = val
}

// GetSiteN returns the SiteN property
func (m DcimInterfacesListQueryParameters) GetSiteN() []string {
	return m.SiteN
}

// SetSiteN sets the SiteN property
func (m *DcimInterfacesListQueryParameters) SetSiteN(val []string) {
	m.SiteN = val
}

// GetSiteGroup returns the SiteGroup property
func (m DcimInterfacesListQueryParameters) GetSiteGroup() []int32 {
	return m.SiteGroup
}

// SetSiteGroup sets the SiteGroup property
func (m *DcimInterfacesListQueryParameters) SetSiteGroup(val []int32) {
	m.SiteGroup = val
}

// GetSiteGroupN returns the SiteGroupN property
func (m DcimInterfacesListQueryParameters) GetSiteGroupN() []int32 {
	return m.SiteGroupN
}

// SetSiteGroupN sets the SiteGroupN property
func (m *DcimInterfacesListQueryParameters) SetSiteGroupN(val []int32) {
	m.SiteGroupN = val
}

// GetSiteGroupId returns the SiteGroupId property
func (m DcimInterfacesListQueryParameters) GetSiteGroupId() []int32 {
	return m.SiteGroupId
}

// SetSiteGroupId sets the SiteGroupId property
func (m *DcimInterfacesListQueryParameters) SetSiteGroupId(val []int32) {
	m.SiteGroupId = val
}

// GetSiteGroupIdN returns the SiteGroupIdN property
func (m DcimInterfacesListQueryParameters) GetSiteGroupIdN() []int32 {
	return m.SiteGroupIdN
}

// SetSiteGroupIdN sets the SiteGroupIdN property
func (m *DcimInterfacesListQueryParameters) SetSiteGroupIdN(val []int32) {
	m.SiteGroupIdN = val
}

// GetSiteId returns the SiteId property
func (m DcimInterfacesListQueryParameters) GetSiteId() []int32 {
	return m.SiteId
}

// SetSiteId sets the SiteId property
func (m *DcimInterfacesListQueryParameters) SetSiteId(val []int32) {
	m.SiteId = val
}

// GetSiteIdN returns the SiteIdN property
func (m DcimInterfacesListQueryParameters) GetSiteIdN() []int32 {
	return m.SiteIdN
}

// SetSiteIdN sets the SiteIdN property
func (m *DcimInterfacesListQueryParameters) SetSiteIdN(val []int32) {
	m.SiteIdN = val
}

// GetSpeed returns the Speed property
func (m DcimInterfacesListQueryParameters) GetSpeed() []int32 {
	return m.Speed
}

// SetSpeed sets the Speed property
func (m *DcimInterfacesListQueryParameters) SetSpeed(val []int32) {
	m.Speed = val
}

// GetSpeedGt returns the SpeedGt property
func (m DcimInterfacesListQueryParameters) GetSpeedGt() []int32 {
	return m.SpeedGt
}

// SetSpeedGt sets the SpeedGt property
func (m *DcimInterfacesListQueryParameters) SetSpeedGt(val []int32) {
	m.SpeedGt = val
}

// GetSpeedGte returns the SpeedGte property
func (m DcimInterfacesListQueryParameters) GetSpeedGte() []int32 {
	return m.SpeedGte
}

// SetSpeedGte sets the SpeedGte property
func (m *DcimInterfacesListQueryParameters) SetSpeedGte(val []int32) {
	m.SpeedGte = val
}

// GetSpeedLt returns the SpeedLt property
func (m DcimInterfacesListQueryParameters) GetSpeedLt() []int32 {
	return m.SpeedLt
}

// SetSpeedLt sets the SpeedLt property
func (m *DcimInterfacesListQueryParameters) SetSpeedLt(val []int32) {
	m.SpeedLt = val
}

// GetSpeedLte returns the SpeedLte property
func (m DcimInterfacesListQueryParameters) GetSpeedLte() []int32 {
	return m.SpeedLte
}

// SetSpeedLte sets the SpeedLte property
func (m *DcimInterfacesListQueryParameters) SetSpeedLte(val []int32) {
	m.SpeedLte = val
}

// GetSpeedN returns the SpeedN property
func (m DcimInterfacesListQueryParameters) GetSpeedN() []int32 {
	return m.SpeedN
}

// SetSpeedN sets the SpeedN property
func (m *DcimInterfacesListQueryParameters) SetSpeedN(val []int32) {
	m.SpeedN = val
}

// GetTag returns the Tag property
func (m DcimInterfacesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *DcimInterfacesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m DcimInterfacesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *DcimInterfacesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}

// GetTxPower returns the TxPower property
func (m DcimInterfacesListQueryParameters) GetTxPower() []int32 {
	return m.TxPower
}

// SetTxPower sets the TxPower property
func (m *DcimInterfacesListQueryParameters) SetTxPower(val []int32) {
	m.TxPower = val
}

// GetTxPowerGt returns the TxPowerGt property
func (m DcimInterfacesListQueryParameters) GetTxPowerGt() []int32 {
	return m.TxPowerGt
}

// SetTxPowerGt sets the TxPowerGt property
func (m *DcimInterfacesListQueryParameters) SetTxPowerGt(val []int32) {
	m.TxPowerGt = val
}

// GetTxPowerGte returns the TxPowerGte property
func (m DcimInterfacesListQueryParameters) GetTxPowerGte() []int32 {
	return m.TxPowerGte
}

// SetTxPowerGte sets the TxPowerGte property
func (m *DcimInterfacesListQueryParameters) SetTxPowerGte(val []int32) {
	m.TxPowerGte = val
}

// GetTxPowerLt returns the TxPowerLt property
func (m DcimInterfacesListQueryParameters) GetTxPowerLt() []int32 {
	return m.TxPowerLt
}

// SetTxPowerLt sets the TxPowerLt property
func (m *DcimInterfacesListQueryParameters) SetTxPowerLt(val []int32) {
	m.TxPowerLt = val
}

// GetTxPowerLte returns the TxPowerLte property
func (m DcimInterfacesListQueryParameters) GetTxPowerLte() []int32 {
	return m.TxPowerLte
}

// SetTxPowerLte sets the TxPowerLte property
func (m *DcimInterfacesListQueryParameters) SetTxPowerLte(val []int32) {
	m.TxPowerLte = val
}

// GetTxPowerN returns the TxPowerN property
func (m DcimInterfacesListQueryParameters) GetTxPowerN() []int32 {
	return m.TxPowerN
}

// SetTxPowerN sets the TxPowerN property
func (m *DcimInterfacesListQueryParameters) SetTxPowerN(val []int32) {
	m.TxPowerN = val
}

// GetType returns the Type property
func (m DcimInterfacesListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimInterfacesListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimInterfacesListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimInterfacesListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimInterfacesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimInterfacesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVdc returns the Vdc property
func (m DcimInterfacesListQueryParameters) GetVdc() []string {
	return m.Vdc
}

// SetVdc sets the Vdc property
func (m *DcimInterfacesListQueryParameters) SetVdc(val []string) {
	m.Vdc = val
}

// GetVdcN returns the VdcN property
func (m DcimInterfacesListQueryParameters) GetVdcN() []string {
	return m.VdcN
}

// SetVdcN sets the VdcN property
func (m *DcimInterfacesListQueryParameters) SetVdcN(val []string) {
	m.VdcN = val
}

// GetVdcId returns the VdcId property
func (m DcimInterfacesListQueryParameters) GetVdcId() []int32 {
	return m.VdcId
}

// SetVdcId sets the VdcId property
func (m *DcimInterfacesListQueryParameters) SetVdcId(val []int32) {
	m.VdcId = val
}

// GetVdcIdN returns the VdcIdN property
func (m DcimInterfacesListQueryParameters) GetVdcIdN() []int32 {
	return m.VdcIdN
}

// SetVdcIdN sets the VdcIdN property
func (m *DcimInterfacesListQueryParameters) SetVdcIdN(val []int32) {
	m.VdcIdN = val
}

// GetVdcIdentifier returns the VdcIdentifier property
func (m DcimInterfacesListQueryParameters) GetVdcIdentifier() []*int32 {
	return m.VdcIdentifier
}

// SetVdcIdentifier sets the VdcIdentifier property
func (m *DcimInterfacesListQueryParameters) SetVdcIdentifier(val []*int32) {
	m.VdcIdentifier = val
}

// GetVdcIdentifierN returns the VdcIdentifierN property
func (m DcimInterfacesListQueryParameters) GetVdcIdentifierN() []*int32 {
	return m.VdcIdentifierN
}

// SetVdcIdentifierN sets the VdcIdentifierN property
func (m *DcimInterfacesListQueryParameters) SetVdcIdentifierN(val []*int32) {
	m.VdcIdentifierN = val
}

// GetVirtualChassis returns the VirtualChassis property
func (m DcimInterfacesListQueryParameters) GetVirtualChassis() []string {
	return m.VirtualChassis
}

// SetVirtualChassis sets the VirtualChassis property
func (m *DcimInterfacesListQueryParameters) SetVirtualChassis(val []string) {
	m.VirtualChassis = val
}

// GetVirtualChassisN returns the VirtualChassisN property
func (m DcimInterfacesListQueryParameters) GetVirtualChassisN() []string {
	return m.VirtualChassisN
}

// SetVirtualChassisN sets the VirtualChassisN property
func (m *DcimInterfacesListQueryParameters) SetVirtualChassisN(val []string) {
	m.VirtualChassisN = val
}

// GetVirtualChassisId returns the VirtualChassisId property
func (m DcimInterfacesListQueryParameters) GetVirtualChassisId() []int32 {
	return m.VirtualChassisId
}

// SetVirtualChassisId sets the VirtualChassisId property
func (m *DcimInterfacesListQueryParameters) SetVirtualChassisId(val []int32) {
	m.VirtualChassisId = val
}

// GetVirtualChassisIdN returns the VirtualChassisIdN property
func (m DcimInterfacesListQueryParameters) GetVirtualChassisIdN() []int32 {
	return m.VirtualChassisIdN
}

// SetVirtualChassisIdN sets the VirtualChassisIdN property
func (m *DcimInterfacesListQueryParameters) SetVirtualChassisIdN(val []int32) {
	m.VirtualChassisIdN = val
}

// GetVlan returns the Vlan property
func (m DcimInterfacesListQueryParameters) GetVlan() string {
	return m.Vlan
}

// SetVlan sets the Vlan property
func (m *DcimInterfacesListQueryParameters) SetVlan(val string) {
	m.Vlan = val
}

// GetVlanId returns the VlanId property
func (m DcimInterfacesListQueryParameters) GetVlanId() string {
	return m.VlanId
}

// SetVlanId sets the VlanId property
func (m *DcimInterfacesListQueryParameters) SetVlanId(val string) {
	m.VlanId = val
}

// GetVrf returns the Vrf property
func (m DcimInterfacesListQueryParameters) GetVrf() []*string {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *DcimInterfacesListQueryParameters) SetVrf(val []*string) {
	m.Vrf = val
}

// GetVrfN returns the VrfN property
func (m DcimInterfacesListQueryParameters) GetVrfN() []*string {
	return m.VrfN
}

// SetVrfN sets the VrfN property
func (m *DcimInterfacesListQueryParameters) SetVrfN(val []*string) {
	m.VrfN = val
}

// GetVrfId returns the VrfId property
func (m DcimInterfacesListQueryParameters) GetVrfId() []int32 {
	return m.VrfId
}

// SetVrfId sets the VrfId property
func (m *DcimInterfacesListQueryParameters) SetVrfId(val []int32) {
	m.VrfId = val
}

// GetVrfIdN returns the VrfIdN property
func (m DcimInterfacesListQueryParameters) GetVrfIdN() []int32 {
	return m.VrfIdN
}

// SetVrfIdN sets the VrfIdN property
func (m *DcimInterfacesListQueryParameters) SetVrfIdN(val []int32) {
	m.VrfIdN = val
}

// GetWwn returns the Wwn property
func (m DcimInterfacesListQueryParameters) GetWwn() []string {
	return m.Wwn
}

// SetWwn sets the Wwn property
func (m *DcimInterfacesListQueryParameters) SetWwn(val []string) {
	m.Wwn = val
}

// GetWwnIc returns the WwnIc property
func (m DcimInterfacesListQueryParameters) GetWwnIc() []string {
	return m.WwnIc
}

// SetWwnIc sets the WwnIc property
func (m *DcimInterfacesListQueryParameters) SetWwnIc(val []string) {
	m.WwnIc = val
}

// GetWwnIe returns the WwnIe property
func (m DcimInterfacesListQueryParameters) GetWwnIe() []string {
	return m.WwnIe
}

// SetWwnIe sets the WwnIe property
func (m *DcimInterfacesListQueryParameters) SetWwnIe(val []string) {
	m.WwnIe = val
}

// GetWwnIew returns the WwnIew property
func (m DcimInterfacesListQueryParameters) GetWwnIew() []string {
	return m.WwnIew
}

// SetWwnIew sets the WwnIew property
func (m *DcimInterfacesListQueryParameters) SetWwnIew(val []string) {
	m.WwnIew = val
}

// GetWwnIsw returns the WwnIsw property
func (m DcimInterfacesListQueryParameters) GetWwnIsw() []string {
	return m.WwnIsw
}

// SetWwnIsw sets the WwnIsw property
func (m *DcimInterfacesListQueryParameters) SetWwnIsw(val []string) {
	m.WwnIsw = val
}

// GetWwnN returns the WwnN property
func (m DcimInterfacesListQueryParameters) GetWwnN() []string {
	return m.WwnN
}

// SetWwnN sets the WwnN property
func (m *DcimInterfacesListQueryParameters) SetWwnN(val []string) {
	m.WwnN = val
}

// GetWwnNic returns the WwnNic property
func (m DcimInterfacesListQueryParameters) GetWwnNic() []string {
	return m.WwnNic
}

// SetWwnNic sets the WwnNic property
func (m *DcimInterfacesListQueryParameters) SetWwnNic(val []string) {
	m.WwnNic = val
}

// GetWwnNie returns the WwnNie property
func (m DcimInterfacesListQueryParameters) GetWwnNie() []string {
	return m.WwnNie
}

// SetWwnNie sets the WwnNie property
func (m *DcimInterfacesListQueryParameters) SetWwnNie(val []string) {
	m.WwnNie = val
}

// GetWwnNiew returns the WwnNiew property
func (m DcimInterfacesListQueryParameters) GetWwnNiew() []string {
	return m.WwnNiew
}

// SetWwnNiew sets the WwnNiew property
func (m *DcimInterfacesListQueryParameters) SetWwnNiew(val []string) {
	m.WwnNiew = val
}

// GetWwnNisw returns the WwnNisw property
func (m DcimInterfacesListQueryParameters) GetWwnNisw() []string {
	return m.WwnNisw
}

// SetWwnNisw sets the WwnNisw property
func (m *DcimInterfacesListQueryParameters) SetWwnNisw(val []string) {
	m.WwnNisw = val
}
