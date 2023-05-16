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

// PatchedWritableInterfaceRequest is an object. Adds support for custom fields and tags.
type PatchedWritableInterfaceRequest struct {
	// Bridge:
	Bridge *int32 `json:"bridge,omitempty" mapstructure:"bridge,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Duplex: * `half` - Half
	// * `full` - Full
	// * `auto` - Auto
	Duplex *string `json:"duplex,omitempty" mapstructure:"duplex,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Lag:
	Lag *int32 `json:"lag,omitempty" mapstructure:"lag,omitempty"`
	// MacAddress:
	MacAddress *string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// MgmtOnly: This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty" mapstructure:"mgmt_only,omitempty"`
	// Mode: IEEE 802.1Q tagging strategy
	//
	// * `access` - Access
	// * `tagged` - Tagged
	// * `tagged-all` - Tagged (All)
	Mode string `json:"mode,omitempty" mapstructure:"mode,omitempty"`
	// Module:
	Module *int32 `json:"module,omitempty" mapstructure:"module,omitempty"`
	// Mtu:
	Mtu *int32 `json:"mtu,omitempty" mapstructure:"mtu,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
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
	RfChannel string `json:"rf_channel,omitempty" mapstructure:"rf_channel,omitempty"`
	// RfChannelFrequency: Populated by selected channel (if set)
	RfChannelFrequency *float64 `json:"rf_channel_frequency,omitempty" mapstructure:"rf_channel_frequency,omitempty"`
	// RfChannelWidth: Populated by selected channel (if set)
	RfChannelWidth *float64 `json:"rf_channel_width,omitempty" mapstructure:"rf_channel_width,omitempty"`
	// RfRole: * `ap` - Access point
	// * `station` - Station
	RfRole string `json:"rf_role,omitempty" mapstructure:"rf_role,omitempty"`
	// Speed:
	Speed *int32 `json:"speed,omitempty" mapstructure:"speed,omitempty"`
	// TaggedVlans:
	TaggedVlans []int32 `json:"tagged_vlans,omitempty" mapstructure:"tagged_vlans,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TxPower:
	TxPower *int32 `json:"tx_power,omitempty" mapstructure:"tx_power,omitempty"`
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
	// UntaggedVlan:
	UntaggedVlan *int32 `json:"untagged_vlan,omitempty" mapstructure:"untagged_vlan,omitempty"`
	// Vdcs:
	Vdcs []int32 `json:"vdcs,omitempty" mapstructure:"vdcs,omitempty"`
	// Vrf:
	Vrf *int32 `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
	// WirelessLans:
	WirelessLans []int32 `json:"wireless_lans,omitempty" mapstructure:"wireless_lans,omitempty"`
	// Wwn:
	Wwn string `json:"wwn,omitempty" mapstructure:"wwn,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableInterfaceRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"macAddress": validation.Validate(
			m.MacAddress, validation.Length(1, 0),
		),
		"mtu": validation.Validate(
			m.Mtu, validation.Min(*int32(1)), validation.Max(*int32(65536)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 64),
		),
		"rfChannelFrequency": validation.Validate(
			m.RfChannelFrequency, validation.Min(*float64(-100000)), validation.Max(*float64(100000)),
		),
		"rfChannelWidth": validation.Validate(
			m.RfChannelWidth, validation.Min(*float64(-10000)), validation.Max(*float64(10000)),
		),
		"speed": validation.Validate(
			m.Speed, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"taggedVlans": validation.Validate(
			m.TaggedVlans,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"txPower": validation.Validate(
			m.TxPower, validation.Min(*int32(0)), validation.Max(*int32(127)),
		),
		"vdcs": validation.Validate(
			m.Vdcs,
		),
		"wirelessLans": validation.Validate(
			m.WirelessLans,
		),
		"wwn": validation.Validate(
			m.Wwn, validation.Length(1, 0),
		),
	}.Filter()
}

// GetBridge returns the Bridge property
func (m PatchedWritableInterfaceRequest) GetBridge() *int32 {
	return m.Bridge
}

// SetBridge sets the Bridge property
func (m *PatchedWritableInterfaceRequest) SetBridge(val *int32) {
	m.Bridge = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableInterfaceRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableInterfaceRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableInterfaceRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableInterfaceRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PatchedWritableInterfaceRequest) GetDevice() int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *PatchedWritableInterfaceRequest) SetDevice(val int32) {
	m.Device = val
}

// GetDuplex returns the Duplex property
func (m PatchedWritableInterfaceRequest) GetDuplex() *string {
	return m.Duplex
}

// SetDuplex sets the Duplex property
func (m *PatchedWritableInterfaceRequest) SetDuplex(val *string) {
	m.Duplex = val
}

// GetEnabled returns the Enabled property
func (m PatchedWritableInterfaceRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *PatchedWritableInterfaceRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetLabel returns the Label property
func (m PatchedWritableInterfaceRequest) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PatchedWritableInterfaceRequest) SetLabel(val string) {
	m.Label = val
}

// GetLag returns the Lag property
func (m PatchedWritableInterfaceRequest) GetLag() *int32 {
	return m.Lag
}

// SetLag sets the Lag property
func (m *PatchedWritableInterfaceRequest) SetLag(val *int32) {
	m.Lag = val
}

// GetMacAddress returns the MacAddress property
func (m PatchedWritableInterfaceRequest) GetMacAddress() *string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *PatchedWritableInterfaceRequest) SetMacAddress(val *string) {
	m.MacAddress = val
}

// GetMarkConnected returns the MarkConnected property
func (m PatchedWritableInterfaceRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *PatchedWritableInterfaceRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetMgmtOnly returns the MgmtOnly property
func (m PatchedWritableInterfaceRequest) GetMgmtOnly() bool {
	return m.MgmtOnly
}

// SetMgmtOnly sets the MgmtOnly property
func (m *PatchedWritableInterfaceRequest) SetMgmtOnly(val bool) {
	m.MgmtOnly = val
}

// GetMode returns the Mode property
func (m PatchedWritableInterfaceRequest) GetMode() string {
	return m.Mode
}

// SetMode sets the Mode property
func (m *PatchedWritableInterfaceRequest) SetMode(val string) {
	m.Mode = val
}

// GetModule returns the Module property
func (m PatchedWritableInterfaceRequest) GetModule() *int32 {
	return m.Module
}

// SetModule sets the Module property
func (m *PatchedWritableInterfaceRequest) SetModule(val *int32) {
	m.Module = val
}

// GetMtu returns the Mtu property
func (m PatchedWritableInterfaceRequest) GetMtu() *int32 {
	return m.Mtu
}

// SetMtu sets the Mtu property
func (m *PatchedWritableInterfaceRequest) SetMtu(val *int32) {
	m.Mtu = val
}

// GetName returns the Name property
func (m PatchedWritableInterfaceRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableInterfaceRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m PatchedWritableInterfaceRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *PatchedWritableInterfaceRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetPoeMode returns the PoeMode property
func (m PatchedWritableInterfaceRequest) GetPoeMode() string {
	return m.PoeMode
}

// SetPoeMode sets the PoeMode property
func (m *PatchedWritableInterfaceRequest) SetPoeMode(val string) {
	m.PoeMode = val
}

// GetPoeType returns the PoeType property
func (m PatchedWritableInterfaceRequest) GetPoeType() string {
	return m.PoeType
}

// SetPoeType sets the PoeType property
func (m *PatchedWritableInterfaceRequest) SetPoeType(val string) {
	m.PoeType = val
}

// GetRfChannel returns the RfChannel property
func (m PatchedWritableInterfaceRequest) GetRfChannel() string {
	return m.RfChannel
}

// SetRfChannel sets the RfChannel property
func (m *PatchedWritableInterfaceRequest) SetRfChannel(val string) {
	m.RfChannel = val
}

// GetRfChannelFrequency returns the RfChannelFrequency property
func (m PatchedWritableInterfaceRequest) GetRfChannelFrequency() *float64 {
	return m.RfChannelFrequency
}

// SetRfChannelFrequency sets the RfChannelFrequency property
func (m *PatchedWritableInterfaceRequest) SetRfChannelFrequency(val *float64) {
	m.RfChannelFrequency = val
}

// GetRfChannelWidth returns the RfChannelWidth property
func (m PatchedWritableInterfaceRequest) GetRfChannelWidth() *float64 {
	return m.RfChannelWidth
}

// SetRfChannelWidth sets the RfChannelWidth property
func (m *PatchedWritableInterfaceRequest) SetRfChannelWidth(val *float64) {
	m.RfChannelWidth = val
}

// GetRfRole returns the RfRole property
func (m PatchedWritableInterfaceRequest) GetRfRole() string {
	return m.RfRole
}

// SetRfRole sets the RfRole property
func (m *PatchedWritableInterfaceRequest) SetRfRole(val string) {
	m.RfRole = val
}

// GetSpeed returns the Speed property
func (m PatchedWritableInterfaceRequest) GetSpeed() *int32 {
	return m.Speed
}

// SetSpeed sets the Speed property
func (m *PatchedWritableInterfaceRequest) SetSpeed(val *int32) {
	m.Speed = val
}

// GetTaggedVlans returns the TaggedVlans property
func (m PatchedWritableInterfaceRequest) GetTaggedVlans() []int32 {
	return m.TaggedVlans
}

// SetTaggedVlans sets the TaggedVlans property
func (m *PatchedWritableInterfaceRequest) SetTaggedVlans(val []int32) {
	m.TaggedVlans = val
}

// GetTags returns the Tags property
func (m PatchedWritableInterfaceRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableInterfaceRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTxPower returns the TxPower property
func (m PatchedWritableInterfaceRequest) GetTxPower() *int32 {
	return m.TxPower
}

// SetTxPower sets the TxPower property
func (m *PatchedWritableInterfaceRequest) SetTxPower(val *int32) {
	m.TxPower = val
}

// GetType returns the Type property
func (m PatchedWritableInterfaceRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritableInterfaceRequest) SetType(val string) {
	m.Type = val
}

// GetUntaggedVlan returns the UntaggedVlan property
func (m PatchedWritableInterfaceRequest) GetUntaggedVlan() *int32 {
	return m.UntaggedVlan
}

// SetUntaggedVlan sets the UntaggedVlan property
func (m *PatchedWritableInterfaceRequest) SetUntaggedVlan(val *int32) {
	m.UntaggedVlan = val
}

// GetVdcs returns the Vdcs property
func (m PatchedWritableInterfaceRequest) GetVdcs() []int32 {
	return m.Vdcs
}

// SetVdcs sets the Vdcs property
func (m *PatchedWritableInterfaceRequest) SetVdcs(val []int32) {
	m.Vdcs = val
}

// GetVrf returns the Vrf property
func (m PatchedWritableInterfaceRequest) GetVrf() *int32 {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *PatchedWritableInterfaceRequest) SetVrf(val *int32) {
	m.Vrf = val
}

// GetWirelessLans returns the WirelessLans property
func (m PatchedWritableInterfaceRequest) GetWirelessLans() []int32 {
	return m.WirelessLans
}

// SetWirelessLans sets the WirelessLans property
func (m *PatchedWritableInterfaceRequest) SetWirelessLans(val []int32) {
	m.WirelessLans = val
}

// GetWwn returns the Wwn property
func (m PatchedWritableInterfaceRequest) GetWwn() string {
	return m.Wwn
}

// SetWwn sets the Wwn property
func (m *PatchedWritableInterfaceRequest) SetWwn(val string) {
	m.Wwn = val
}
