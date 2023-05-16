// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// PatchedWritableInterfaceTemplateRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableInterfaceTemplateRequest struct {
	// Bridge:
	Bridge *int32 `json:"bridge,omitempty" mapstructure:"bridge,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType:
	DeviceType *int32 `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// MgmtOnly:
	MgmtOnly bool `json:"mgmt_only,omitempty" mapstructure:"mgmt_only,omitempty"`
	// ModuleType:
	ModuleType *int32 `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// PoeMode: * `pd` - PD
	// * `pse` - PSE
	PoeMode string `json:"poe_mode,omitempty" mapstructure:"poe_mode,omitempty"`
	// PoeType: * `type1-ieee802.3af` - 802.3af (Type 1)
	// * `type2-ieee802.3at` - 802.3at (Type 2)
	// * `type3-ieee802.3bt` - 802.3bt (Type 3)
	// * `type4-ieee802.3bt` - 802.3bt (Type 4)
	// * `passive-24v-2pair` - Passive 24V (2-pair)
	// * `passive-24v-4pair` - Passive 24V (4-pair)
	// * `passive-48v-2pair` - Passive 48V (2-pair)
	// * `passive-48v-4pair` - Passive 48V (4-pair)
	PoeType string `json:"poe_type,omitempty" mapstructure:"poe_type,omitempty"`
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
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableInterfaceTemplateRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 64),
		),
	}.Filter()
}

// GetBridge returns the Bridge property
func (m PatchedWritableInterfaceTemplateRequest) GetBridge() *int32 {
	return m.Bridge
}

// SetBridge sets the Bridge property
func (m *PatchedWritableInterfaceTemplateRequest) SetBridge(val *int32) {
	m.Bridge = val
}

// GetDescription returns the Description property
func (m PatchedWritableInterfaceTemplateRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableInterfaceTemplateRequest) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m PatchedWritableInterfaceTemplateRequest) GetDeviceType() *int32 {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *PatchedWritableInterfaceTemplateRequest) SetDeviceType(val *int32) {
	m.DeviceType = val
}

// GetEnabled returns the Enabled property
func (m PatchedWritableInterfaceTemplateRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *PatchedWritableInterfaceTemplateRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetLabel returns the Label property
func (m PatchedWritableInterfaceTemplateRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PatchedWritableInterfaceTemplateRequest) SetLabel(val string) {
	m.Label = val
}

// GetMgmtOnly returns the MgmtOnly property
func (m PatchedWritableInterfaceTemplateRequest) GetMgmtOnly() bool {
	return m.MgmtOnly
}

// SetMgmtOnly sets the MgmtOnly property
func (m *PatchedWritableInterfaceTemplateRequest) SetMgmtOnly(val bool) {
	m.MgmtOnly = val
}

// GetModuleType returns the ModuleType property
func (m PatchedWritableInterfaceTemplateRequest) GetModuleType() *int32 {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *PatchedWritableInterfaceTemplateRequest) SetModuleType(val *int32) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m PatchedWritableInterfaceTemplateRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableInterfaceTemplateRequest) SetName(val string) {
	m.Name = val
}

// GetPoeMode returns the PoeMode property
func (m PatchedWritableInterfaceTemplateRequest) GetPoeMode() string {
	return m.PoeMode
}

// SetPoeMode sets the PoeMode property
func (m *PatchedWritableInterfaceTemplateRequest) SetPoeMode(val string) {
	m.PoeMode = val
}

// GetPoeType returns the PoeType property
func (m PatchedWritableInterfaceTemplateRequest) GetPoeType() string {
	return m.PoeType
}

// SetPoeType sets the PoeType property
func (m *PatchedWritableInterfaceTemplateRequest) SetPoeType(val string) {
	m.PoeType = val
}

// GetType returns the Type property
func (m PatchedWritableInterfaceTemplateRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritableInterfaceTemplateRequest) SetType(val string) {
	m.Type = val
}
