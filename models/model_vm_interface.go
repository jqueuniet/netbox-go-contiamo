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

// VMInterface is an object. Adds support for custom fields and tags.
type VMInterface struct {
	// Bridge: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Bridge *NestedVMInterface `json:"bridge,omitempty" mapstructure:"bridge,omitempty"`
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
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// L2vpnTermination: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	L2vpnTermination *NestedL2VPNTermination `json:"l2vpn_termination" mapstructure:"l2vpn_termination"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// MacAddress:
	MacAddress *string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// Mode:
	Mode *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"mode,omitempty" mapstructure:"mode,omitempty"`
	// Mtu:
	Mtu *int32 `json:"mtu,omitempty" mapstructure:"mtu,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedVMInterface `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// TaggedVlans:
	TaggedVlans []int32 `json:"tagged_vlans,omitempty" mapstructure:"tagged_vlans,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// UntaggedVlan: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	UntaggedVlan *NestedVLAN `json:"untagged_vlan,omitempty" mapstructure:"untagged_vlan,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VirtualMachine: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	VirtualMachine NestedVirtualMachine `json:"virtual_machine" mapstructure:"virtual_machine"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRF `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m VMInterface) Validate() error {
	return validation.Errors{
		"bridge": validation.Validate(
			m.Bridge,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"l2vpnTermination": validation.Validate(
			m.L2vpnTermination,
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
		"taggedVlans": validation.Validate(
			m.TaggedVlans,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"untaggedVlan": validation.Validate(
			m.UntaggedVlan,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"virtualMachine": validation.Validate(
			m.VirtualMachine, validation.NotNil,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
	}.Filter()
}

// GetBridge returns the Bridge property
func (m VMInterface) GetBridge() *NestedVMInterface {
	return m.Bridge
}

// SetBridge sets the Bridge property
func (m *VMInterface) SetBridge(val *NestedVMInterface) {
	m.Bridge = val
}

// GetCountFhrpGroups returns the CountFhrpGroups property
func (m VMInterface) GetCountFhrpGroups() int32 {
	return m.CountFhrpGroups
}

// SetCountFhrpGroups sets the CountFhrpGroups property
func (m *VMInterface) SetCountFhrpGroups(val int32) {
	m.CountFhrpGroups = val
}

// GetCountIpaddresses returns the CountIpaddresses property
func (m VMInterface) GetCountIpaddresses() int32 {
	return m.CountIpaddresses
}

// SetCountIpaddresses sets the CountIpaddresses property
func (m *VMInterface) SetCountIpaddresses(val int32) {
	m.CountIpaddresses = val
}

// GetCreated returns the Created property
func (m VMInterface) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *VMInterface) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m VMInterface) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VMInterface) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VMInterface) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VMInterface) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m VMInterface) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *VMInterface) SetDisplay(val string) {
	m.Display = val
}

// GetEnabled returns the Enabled property
func (m VMInterface) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *VMInterface) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m VMInterface) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *VMInterface) SetId(val int32) {
	m.Id = val
}

// GetL2vpnTermination returns the L2vpnTermination property
func (m VMInterface) GetL2vpnTermination() *NestedL2VPNTermination {
	return m.L2vpnTermination
}

// SetL2vpnTermination sets the L2vpnTermination property
func (m *VMInterface) SetL2vpnTermination(val *NestedL2VPNTermination) {
	m.L2vpnTermination = val
}

// GetLastUpdated returns the LastUpdated property
func (m VMInterface) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *VMInterface) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMacAddress returns the MacAddress property
func (m VMInterface) GetMacAddress() *string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *VMInterface) SetMacAddress(val *string) {
	m.MacAddress = val
}

// GetMode returns the Mode property
func (m VMInterface) GetMode() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Mode
}

// SetMode sets the Mode property
func (m *VMInterface) SetMode(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Mode = val
}

// GetMtu returns the Mtu property
func (m VMInterface) GetMtu() *int32 {
	return m.Mtu
}

// SetMtu sets the Mtu property
func (m *VMInterface) SetMtu(val *int32) {
	m.Mtu = val
}

// GetName returns the Name property
func (m VMInterface) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VMInterface) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m VMInterface) GetParent() *NestedVMInterface {
	return m.Parent
}

// SetParent sets the Parent property
func (m *VMInterface) SetParent(val *NestedVMInterface) {
	m.Parent = val
}

// GetTaggedVlans returns the TaggedVlans property
func (m VMInterface) GetTaggedVlans() []int32 {
	return m.TaggedVlans
}

// SetTaggedVlans sets the TaggedVlans property
func (m *VMInterface) SetTaggedVlans(val []int32) {
	m.TaggedVlans = val
}

// GetTags returns the Tags property
func (m VMInterface) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VMInterface) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUntaggedVlan returns the UntaggedVlan property
func (m VMInterface) GetUntaggedVlan() *NestedVLAN {
	return m.UntaggedVlan
}

// SetUntaggedVlan sets the UntaggedVlan property
func (m *VMInterface) SetUntaggedVlan(val *NestedVLAN) {
	m.UntaggedVlan = val
}

// GetUrl returns the Url property
func (m VMInterface) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *VMInterface) SetUrl(val string) {
	m.Url = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m VMInterface) GetVirtualMachine() NestedVirtualMachine {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *VMInterface) SetVirtualMachine(val NestedVirtualMachine) {
	m.VirtualMachine = val
}

// GetVrf returns the Vrf property
func (m VMInterface) GetVrf() *NestedVRF {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *VMInterface) SetVrf(val *NestedVRF) {
	m.Vrf = val
}
