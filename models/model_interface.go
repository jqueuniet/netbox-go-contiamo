// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Interface is an object. Adds support for custom fields and tags.
type Interface struct {
	// Occupied:
	Occupied bool `json:"_occupied" mapstructure:"_occupied"`
	// Bridge: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Bridge *NestedInterface `json:"bridge,omitempty" mapstructure:"bridge,omitempty"`
	// Cable: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Cable *NestedCable `json:"cable" mapstructure:"cable"`
	// CableEnd:
	CableEnd string `json:"cable_end" mapstructure:"cable_end"`
	// ConnectedEndpoints:
	ConnectedEndpoints []interface{} `json:"connected_endpoints" mapstructure:"connected_endpoints"`
	// ConnectedEndpointsReachable:
	ConnectedEndpointsReachable bool `json:"connected_endpoints_reachable" mapstructure:"connected_endpoints_reachable"`
	// ConnectedEndpointsType:
	ConnectedEndpointsType string `json:"connected_endpoints_type" mapstructure:"connected_endpoints_type"`
	// CountFhrpGroups:
	CountFhrpGroups int32 `json:"count_fhrp_groups" mapstructure:"count_fhrp_groups"`
	// CountIpaddresses:
	CountIpaddresses int32 `json:"count_ipaddresses" mapstructure:"count_ipaddresses"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device NestedDevice `json:"device" mapstructure:"device"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Duplex:
	Duplex *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"duplex,omitempty" mapstructure:"duplex,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// L2vpnTermination: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	L2vpnTermination *NestedL2VPNTermination `json:"l2vpn_termination" mapstructure:"l2vpn_termination"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// Lag: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Lag *NestedInterface `json:"lag,omitempty" mapstructure:"lag,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// LinkPeers:
	LinkPeers []interface{} `json:"link_peers" mapstructure:"link_peers"`
	// LinkPeersType: Return the type of the peer link terminations, or None.
	LinkPeersType string `json:"link_peers_type" mapstructure:"link_peers_type"`
	// MacAddress:
	MacAddress *string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// MgmtOnly: This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty" mapstructure:"mgmt_only,omitempty"`
	// Mode:
	Mode *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"mode,omitempty" mapstructure:"mode,omitempty"`
	// Module: Used by device component serializers.
	Module *ComponentNestedModule `json:"module,omitempty" mapstructure:"module,omitempty"`
	// Mtu:
	Mtu *int32 `json:"mtu,omitempty" mapstructure:"mtu,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedInterface `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// PoeMode:
	PoeMode *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"poe_mode,omitempty" mapstructure:"poe_mode,omitempty"`
	// PoeType:
	PoeType *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"poe_type,omitempty" mapstructure:"poe_type,omitempty"`
	// RfChannel:
	RfChannel *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"rf_channel,omitempty" mapstructure:"rf_channel,omitempty"`
	// RfChannelFrequency: Populated by selected channel (if set)
	RfChannelFrequency *float64 `json:"rf_channel_frequency,omitempty" mapstructure:"rf_channel_frequency,omitempty"`
	// RfChannelWidth: Populated by selected channel (if set)
	RfChannelWidth *float64 `json:"rf_channel_width,omitempty" mapstructure:"rf_channel_width,omitempty"`
	// RfRole:
	RfRole *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"rf_role,omitempty" mapstructure:"rf_role,omitempty"`
	// Speed:
	Speed *int32 `json:"speed,omitempty" mapstructure:"speed,omitempty"`
	// TaggedVlans:
	TaggedVlans []int32 `json:"tagged_vlans,omitempty" mapstructure:"tagged_vlans,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TxPower:
	TxPower *int32 `json:"tx_power,omitempty" mapstructure:"tx_power,omitempty"`
	// Type:
	Type struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type" mapstructure:"type"`
	// UntaggedVlan: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	UntaggedVlan *NestedVLAN `json:"untagged_vlan,omitempty" mapstructure:"untagged_vlan,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Vdcs:
	Vdcs []int32 `json:"vdcs,omitempty" mapstructure:"vdcs,omitempty"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRF `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
	// WirelessLans:
	WirelessLans []int32 `json:"wireless_lans,omitempty" mapstructure:"wireless_lans,omitempty"`
	// WirelessLink: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	WirelessLink *NestedWirelessLink `json:"wireless_link" mapstructure:"wireless_link"`
	// Wwn:
	Wwn string `json:"wwn,omitempty" mapstructure:"wwn,omitempty"`
}

// Validate implements basic validation for this model
func (m Interface) Validate() error {
	return validation.Errors{
		"bridge": validation.Validate(
			m.Bridge,
		),
		"cable": validation.Validate(
			m.Cable,
		),
		"connectedEndpoints": validation.Validate(
			m.ConnectedEndpoints, validation.NotNil,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"device": validation.Validate(
			m.Device, validation.NotNil,
		),
		"l2vpnTermination": validation.Validate(
			m.L2vpnTermination,
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"lag": validation.Validate(
			m.Lag,
		),
		"linkPeers": validation.Validate(
			m.LinkPeers, validation.NotNil,
		),
		"module": validation.Validate(
			m.Module,
		),
		"mtu": validation.Validate(
			m.Mtu, validation.Min(*int32(1)), validation.Max(*int32(65536)),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"parent": validation.Validate(
			m.Parent,
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
		"untaggedVlan": validation.Validate(
			m.UntaggedVlan,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"vdcs": validation.Validate(
			m.Vdcs,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
		"wirelessLans": validation.Validate(
			m.WirelessLans,
		),
		"wirelessLink": validation.Validate(
			m.WirelessLink,
		),
	}.Filter()
}

// GetOccupied returns the Occupied property
func (m Interface) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *Interface) SetOccupied(val bool) {
	m.Occupied = val
}

// GetBridge returns the Bridge property
func (m Interface) GetBridge() *NestedInterface {
	return m.Bridge
}

// SetBridge sets the Bridge property
func (m *Interface) SetBridge(val *NestedInterface) {
	m.Bridge = val
}

// GetCable returns the Cable property
func (m Interface) GetCable() *NestedCable {
	return m.Cable
}

// SetCable sets the Cable property
func (m *Interface) SetCable(val *NestedCable) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m Interface) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *Interface) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetConnectedEndpoints returns the ConnectedEndpoints property
func (m Interface) GetConnectedEndpoints() []interface{} {
	return m.ConnectedEndpoints
}

// SetConnectedEndpoints sets the ConnectedEndpoints property
func (m *Interface) SetConnectedEndpoints(val []interface{}) {
	m.ConnectedEndpoints = val
}

// GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable property
func (m Interface) GetConnectedEndpointsReachable() bool {
	return m.ConnectedEndpointsReachable
}

// SetConnectedEndpointsReachable sets the ConnectedEndpointsReachable property
func (m *Interface) SetConnectedEndpointsReachable(val bool) {
	m.ConnectedEndpointsReachable = val
}

// GetConnectedEndpointsType returns the ConnectedEndpointsType property
func (m Interface) GetConnectedEndpointsType() string {
	return m.ConnectedEndpointsType
}

// SetConnectedEndpointsType sets the ConnectedEndpointsType property
func (m *Interface) SetConnectedEndpointsType(val string) {
	m.ConnectedEndpointsType = val
}

// GetCountFhrpGroups returns the CountFhrpGroups property
func (m Interface) GetCountFhrpGroups() int32 {
	return m.CountFhrpGroups
}

// SetCountFhrpGroups sets the CountFhrpGroups property
func (m *Interface) SetCountFhrpGroups(val int32) {
	m.CountFhrpGroups = val
}

// GetCountIpaddresses returns the CountIpaddresses property
func (m Interface) GetCountIpaddresses() int32 {
	return m.CountIpaddresses
}

// SetCountIpaddresses sets the CountIpaddresses property
func (m *Interface) SetCountIpaddresses(val int32) {
	m.CountIpaddresses = val
}

// GetCreated returns the Created property
func (m Interface) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Interface) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Interface) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Interface) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Interface) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Interface) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m Interface) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *Interface) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m Interface) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Interface) SetDisplay(val string) {
	m.Display = val
}

// GetDuplex returns the Duplex property
func (m Interface) GetDuplex() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Duplex
}

// SetDuplex sets the Duplex property
func (m *Interface) SetDuplex(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Duplex = val
}

// GetEnabled returns the Enabled property
func (m Interface) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *Interface) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m Interface) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Interface) SetId(val int32) {
	m.Id = val
}

// GetL2vpnTermination returns the L2vpnTermination property
func (m Interface) GetL2vpnTermination() *NestedL2VPNTermination {
	return m.L2vpnTermination
}

// SetL2vpnTermination sets the L2vpnTermination property
func (m *Interface) SetL2vpnTermination(val *NestedL2VPNTermination) {
	m.L2vpnTermination = val
}

// GetLabel returns the Label property
func (m Interface) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *Interface) SetLabel(val string) {
	m.Label = val
}

// GetLag returns the Lag property
func (m Interface) GetLag() *NestedInterface {
	return m.Lag
}

// SetLag sets the Lag property
func (m *Interface) SetLag(val *NestedInterface) {
	m.Lag = val
}

// GetLastUpdated returns the LastUpdated property
func (m Interface) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Interface) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLinkPeers returns the LinkPeers property
func (m Interface) GetLinkPeers() []interface{} {
	return m.LinkPeers
}

// SetLinkPeers sets the LinkPeers property
func (m *Interface) SetLinkPeers(val []interface{}) {
	m.LinkPeers = val
}

// GetLinkPeersType returns the LinkPeersType property
func (m Interface) GetLinkPeersType() string {
	return m.LinkPeersType
}

// SetLinkPeersType sets the LinkPeersType property
func (m *Interface) SetLinkPeersType(val string) {
	m.LinkPeersType = val
}

// GetMacAddress returns the MacAddress property
func (m Interface) GetMacAddress() *string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *Interface) SetMacAddress(val *string) {
	m.MacAddress = val
}

// GetMarkConnected returns the MarkConnected property
func (m Interface) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *Interface) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetMgmtOnly returns the MgmtOnly property
func (m Interface) GetMgmtOnly() bool {
	return m.MgmtOnly
}

// SetMgmtOnly sets the MgmtOnly property
func (m *Interface) SetMgmtOnly(val bool) {
	m.MgmtOnly = val
}

// GetMode returns the Mode property
func (m Interface) GetMode() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Mode
}

// SetMode sets the Mode property
func (m *Interface) SetMode(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Mode = val
}

// GetModule returns the Module property
func (m Interface) GetModule() *ComponentNestedModule {
	return m.Module
}

// SetModule sets the Module property
func (m *Interface) SetModule(val *ComponentNestedModule) {
	m.Module = val
}

// GetMtu returns the Mtu property
func (m Interface) GetMtu() *int32 {
	return m.Mtu
}

// SetMtu sets the Mtu property
func (m *Interface) SetMtu(val *int32) {
	m.Mtu = val
}

// GetName returns the Name property
func (m Interface) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Interface) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m Interface) GetParent() *NestedInterface {
	return m.Parent
}

// SetParent sets the Parent property
func (m *Interface) SetParent(val *NestedInterface) {
	m.Parent = val
}

// GetPoeMode returns the PoeMode property
func (m Interface) GetPoeMode() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.PoeMode
}

// SetPoeMode sets the PoeMode property
func (m *Interface) SetPoeMode(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.PoeMode = val
}

// GetPoeType returns the PoeType property
func (m Interface) GetPoeType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.PoeType
}

// SetPoeType sets the PoeType property
func (m *Interface) SetPoeType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.PoeType = val
}

// GetRfChannel returns the RfChannel property
func (m Interface) GetRfChannel() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.RfChannel
}

// SetRfChannel sets the RfChannel property
func (m *Interface) SetRfChannel(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.RfChannel = val
}

// GetRfChannelFrequency returns the RfChannelFrequency property
func (m Interface) GetRfChannelFrequency() *float64 {
	return m.RfChannelFrequency
}

// SetRfChannelFrequency sets the RfChannelFrequency property
func (m *Interface) SetRfChannelFrequency(val *float64) {
	m.RfChannelFrequency = val
}

// GetRfChannelWidth returns the RfChannelWidth property
func (m Interface) GetRfChannelWidth() *float64 {
	return m.RfChannelWidth
}

// SetRfChannelWidth sets the RfChannelWidth property
func (m *Interface) SetRfChannelWidth(val *float64) {
	m.RfChannelWidth = val
}

// GetRfRole returns the RfRole property
func (m Interface) GetRfRole() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.RfRole
}

// SetRfRole sets the RfRole property
func (m *Interface) SetRfRole(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.RfRole = val
}

// GetSpeed returns the Speed property
func (m Interface) GetSpeed() *int32 {
	return m.Speed
}

// SetSpeed sets the Speed property
func (m *Interface) SetSpeed(val *int32) {
	m.Speed = val
}

// GetTaggedVlans returns the TaggedVlans property
func (m Interface) GetTaggedVlans() []int32 {
	return m.TaggedVlans
}

// SetTaggedVlans sets the TaggedVlans property
func (m *Interface) SetTaggedVlans(val []int32) {
	m.TaggedVlans = val
}

// GetTags returns the Tags property
func (m Interface) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Interface) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTxPower returns the TxPower property
func (m Interface) GetTxPower() *int32 {
	return m.TxPower
}

// SetTxPower sets the TxPower property
func (m *Interface) SetTxPower(val *int32) {
	m.TxPower = val
}

// GetType returns the Type property
func (m Interface) GetType() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *Interface) SetType(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUntaggedVlan returns the UntaggedVlan property
func (m Interface) GetUntaggedVlan() *NestedVLAN {
	return m.UntaggedVlan
}

// SetUntaggedVlan sets the UntaggedVlan property
func (m *Interface) SetUntaggedVlan(val *NestedVLAN) {
	m.UntaggedVlan = val
}

// GetUrl returns the Url property
func (m Interface) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Interface) SetUrl(val string) {
	m.Url = val
}

// GetVdcs returns the Vdcs property
func (m Interface) GetVdcs() []int32 {
	return m.Vdcs
}

// SetVdcs sets the Vdcs property
func (m *Interface) SetVdcs(val []int32) {
	m.Vdcs = val
}

// GetVrf returns the Vrf property
func (m Interface) GetVrf() *NestedVRF {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *Interface) SetVrf(val *NestedVRF) {
	m.Vrf = val
}

// GetWirelessLans returns the WirelessLans property
func (m Interface) GetWirelessLans() []int32 {
	return m.WirelessLans
}

// SetWirelessLans sets the WirelessLans property
func (m *Interface) SetWirelessLans(val []int32) {
	m.WirelessLans = val
}

// GetWirelessLink returns the WirelessLink property
func (m Interface) GetWirelessLink() *NestedWirelessLink {
	return m.WirelessLink
}

// SetWirelessLink sets the WirelessLink property
func (m *Interface) SetWirelessLink(val *NestedWirelessLink) {
	m.WirelessLink = val
}

// GetWwn returns the Wwn property
func (m Interface) GetWwn() string {
	return m.Wwn
}

// SetWwn sets the Wwn property
func (m *Interface) SetWwn(val string) {
	m.Wwn = val
}
