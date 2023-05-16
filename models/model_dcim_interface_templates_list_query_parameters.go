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

// DcimInterfaceTemplatesListQueryParameters is an object.
type DcimInterfaceTemplatesListQueryParameters struct {
	// BridgeId:
	BridgeId []int32 `json:"bridge_id,omitempty" mapstructure:"bridge_id,omitempty"`
	// BridgeIdN:
	BridgeIdN []int32 `json:"bridge_id__n,omitempty" mapstructure:"bridge_id__n,omitempty"`
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
	// MgmtOnly:
	MgmtOnly bool `json:"mgmt_only,omitempty" mapstructure:"mgmt_only,omitempty"`
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
}

// Validate implements basic validation for this model
func (m DcimInterfaceTemplatesListQueryParameters) Validate() error {
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

// GetBridgeId returns the BridgeId property
func (m DcimInterfaceTemplatesListQueryParameters) GetBridgeId() []int32 {
	return m.BridgeId
}

// SetBridgeId sets the BridgeId property
func (m *DcimInterfaceTemplatesListQueryParameters) SetBridgeId(val []int32) {
	m.BridgeId = val
}

// GetBridgeIdN returns the BridgeIdN property
func (m DcimInterfaceTemplatesListQueryParameters) GetBridgeIdN() []int32 {
	return m.BridgeIdN
}

// SetBridgeIdN sets the BridgeIdN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetBridgeIdN(val []int32) {
	m.BridgeIdN = val
}

// GetCreated returns the Created property
func (m DcimInterfaceTemplatesListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *DcimInterfaceTemplatesListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m DcimInterfaceTemplatesListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *DcimInterfaceTemplatesListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m DcimInterfaceTemplatesListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *DcimInterfaceTemplatesListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m DcimInterfaceTemplatesListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *DcimInterfaceTemplatesListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m DcimInterfaceTemplatesListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *DcimInterfaceTemplatesListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m DcimInterfaceTemplatesListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m DcimInterfaceTemplatesListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *DcimInterfaceTemplatesListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevicetypeId returns the DevicetypeId property
func (m DcimInterfaceTemplatesListQueryParameters) GetDevicetypeId() []*int32 {
	return m.DevicetypeId
}

// SetDevicetypeId sets the DevicetypeId property
func (m *DcimInterfaceTemplatesListQueryParameters) SetDevicetypeId(val []*int32) {
	m.DevicetypeId = val
}

// GetDevicetypeIdN returns the DevicetypeIdN property
func (m DcimInterfaceTemplatesListQueryParameters) GetDevicetypeIdN() []*int32 {
	return m.DevicetypeIdN
}

// SetDevicetypeIdN sets the DevicetypeIdN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetDevicetypeIdN(val []*int32) {
	m.DevicetypeIdN = val
}

// GetEnabled returns the Enabled property
func (m DcimInterfaceTemplatesListQueryParameters) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *DcimInterfaceTemplatesListQueryParameters) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m DcimInterfaceTemplatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *DcimInterfaceTemplatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m DcimInterfaceTemplatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *DcimInterfaceTemplatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m DcimInterfaceTemplatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *DcimInterfaceTemplatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m DcimInterfaceTemplatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *DcimInterfaceTemplatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m DcimInterfaceTemplatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *DcimInterfaceTemplatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m DcimInterfaceTemplatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLastUpdated returns the LastUpdated property
func (m DcimInterfaceTemplatesListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *DcimInterfaceTemplatesListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m DcimInterfaceTemplatesListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *DcimInterfaceTemplatesListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m DcimInterfaceTemplatesListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *DcimInterfaceTemplatesListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m DcimInterfaceTemplatesListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *DcimInterfaceTemplatesListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m DcimInterfaceTemplatesListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *DcimInterfaceTemplatesListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m DcimInterfaceTemplatesListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m DcimInterfaceTemplatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *DcimInterfaceTemplatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetMgmtOnly returns the MgmtOnly property
func (m DcimInterfaceTemplatesListQueryParameters) GetMgmtOnly() bool {
	return m.MgmtOnly
}

// SetMgmtOnly sets the MgmtOnly property
func (m *DcimInterfaceTemplatesListQueryParameters) SetMgmtOnly(val bool) {
	m.MgmtOnly = val
}

// GetModuletypeId returns the ModuletypeId property
func (m DcimInterfaceTemplatesListQueryParameters) GetModuletypeId() []*int32 {
	return m.ModuletypeId
}

// SetModuletypeId sets the ModuletypeId property
func (m *DcimInterfaceTemplatesListQueryParameters) SetModuletypeId(val []*int32) {
	m.ModuletypeId = val
}

// GetModuletypeIdN returns the ModuletypeIdN property
func (m DcimInterfaceTemplatesListQueryParameters) GetModuletypeIdN() []*int32 {
	return m.ModuletypeIdN
}

// SetModuletypeIdN sets the ModuletypeIdN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetModuletypeIdN(val []*int32) {
	m.ModuletypeIdN = val
}

// GetName returns the Name property
func (m DcimInterfaceTemplatesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *DcimInterfaceTemplatesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m DcimInterfaceTemplatesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *DcimInterfaceTemplatesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m DcimInterfaceTemplatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *DcimInterfaceTemplatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m DcimInterfaceTemplatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *DcimInterfaceTemplatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPoeMode returns the PoeMode property
func (m DcimInterfaceTemplatesListQueryParameters) GetPoeMode() []string {
	return m.PoeMode
}

// SetPoeMode sets the PoeMode property
func (m *DcimInterfaceTemplatesListQueryParameters) SetPoeMode(val []string) {
	m.PoeMode = val
}

// GetPoeModeN returns the PoeModeN property
func (m DcimInterfaceTemplatesListQueryParameters) GetPoeModeN() []string {
	return m.PoeModeN
}

// SetPoeModeN sets the PoeModeN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetPoeModeN(val []string) {
	m.PoeModeN = val
}

// GetPoeType returns the PoeType property
func (m DcimInterfaceTemplatesListQueryParameters) GetPoeType() []string {
	return m.PoeType
}

// SetPoeType sets the PoeType property
func (m *DcimInterfaceTemplatesListQueryParameters) SetPoeType(val []string) {
	m.PoeType = val
}

// GetPoeTypeN returns the PoeTypeN property
func (m DcimInterfaceTemplatesListQueryParameters) GetPoeTypeN() []string {
	return m.PoeTypeN
}

// SetPoeTypeN sets the PoeTypeN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetPoeTypeN(val []string) {
	m.PoeTypeN = val
}

// GetQ returns the Q property
func (m DcimInterfaceTemplatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *DcimInterfaceTemplatesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetType returns the Type property
func (m DcimInterfaceTemplatesListQueryParameters) GetType() []string {
	return m.Type
}

// SetType sets the Type property
func (m *DcimInterfaceTemplatesListQueryParameters) SetType(val []string) {
	m.Type = val
}

// GetTypeN returns the TypeN property
func (m DcimInterfaceTemplatesListQueryParameters) GetTypeN() []string {
	return m.TypeN
}

// SetTypeN sets the TypeN property
func (m *DcimInterfaceTemplatesListQueryParameters) SetTypeN(val []string) {
	m.TypeN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m DcimInterfaceTemplatesListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *DcimInterfaceTemplatesListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}
