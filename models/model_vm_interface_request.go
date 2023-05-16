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

// VMInterfaceRequest is an object. Adds support for custom fields and tags.
type VMInterfaceRequest struct {
	// Bridge: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Bridge *NestedVMInterfaceRequest `json:"bridge,omitempty" mapstructure:"bridge,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// MacAddress:
	MacAddress *string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// Mode: * `access` - Access
	// * `tagged` - Tagged
	// * `tagged-all` - Tagged (All)
	Mode string `json:"mode,omitempty" mapstructure:"mode,omitempty"`
	// Mtu:
	Mtu *int32 `json:"mtu,omitempty" mapstructure:"mtu,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Parent *NestedVMInterfaceRequest `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// TaggedVlans:
	TaggedVlans []int32 `json:"tagged_vlans,omitempty" mapstructure:"tagged_vlans,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// UntaggedVlan: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	UntaggedVlan *NestedVLANRequest `json:"untagged_vlan,omitempty" mapstructure:"untagged_vlan,omitempty"`
	// VirtualMachine: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	VirtualMachine NestedVirtualMachineRequest `json:"virtual_machine" mapstructure:"virtual_machine"`
	// Vrf: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Vrf *NestedVRFRequest `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m VMInterfaceRequest) Validate() error {
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
		"macAddress": validation.Validate(
			m.MacAddress, validation.Length(1, 0),
		),
		"mtu": validation.Validate(
			m.Mtu, validation.Min(*int32(1)), validation.Max(*int32(65536)),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 64),
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
		"virtualMachine": validation.Validate(
			m.VirtualMachine, validation.NotNil,
		),
		"vrf": validation.Validate(
			m.Vrf,
		),
	}.Filter()
}

// GetBridge returns the Bridge property
func (m VMInterfaceRequest) GetBridge() *NestedVMInterfaceRequest {
	return m.Bridge
}

// SetBridge sets the Bridge property
func (m *VMInterfaceRequest) SetBridge(val *NestedVMInterfaceRequest) {
	m.Bridge = val
}

// GetCustomFields returns the CustomFields property
func (m VMInterfaceRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *VMInterfaceRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m VMInterfaceRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *VMInterfaceRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnabled returns the Enabled property
func (m VMInterfaceRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *VMInterfaceRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetMacAddress returns the MacAddress property
func (m VMInterfaceRequest) GetMacAddress() *string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *VMInterfaceRequest) SetMacAddress(val *string) {
	m.MacAddress = val
}

// GetMode returns the Mode property
func (m VMInterfaceRequest) GetMode() string {
	return m.Mode
}

// SetMode sets the Mode property
func (m *VMInterfaceRequest) SetMode(val string) {
	m.Mode = val
}

// GetMtu returns the Mtu property
func (m VMInterfaceRequest) GetMtu() *int32 {
	return m.Mtu
}

// SetMtu sets the Mtu property
func (m *VMInterfaceRequest) SetMtu(val *int32) {
	m.Mtu = val
}

// GetName returns the Name property
func (m VMInterfaceRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *VMInterfaceRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m VMInterfaceRequest) GetParent() *NestedVMInterfaceRequest {
	return m.Parent
}

// SetParent sets the Parent property
func (m *VMInterfaceRequest) SetParent(val *NestedVMInterfaceRequest) {
	m.Parent = val
}

// GetTaggedVlans returns the TaggedVlans property
func (m VMInterfaceRequest) GetTaggedVlans() []int32 {
	return m.TaggedVlans
}

// SetTaggedVlans sets the TaggedVlans property
func (m *VMInterfaceRequest) SetTaggedVlans(val []int32) {
	m.TaggedVlans = val
}

// GetTags returns the Tags property
func (m VMInterfaceRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *VMInterfaceRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetUntaggedVlan returns the UntaggedVlan property
func (m VMInterfaceRequest) GetUntaggedVlan() *NestedVLANRequest {
	return m.UntaggedVlan
}

// SetUntaggedVlan sets the UntaggedVlan property
func (m *VMInterfaceRequest) SetUntaggedVlan(val *NestedVLANRequest) {
	m.UntaggedVlan = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m VMInterfaceRequest) GetVirtualMachine() NestedVirtualMachineRequest {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *VMInterfaceRequest) SetVirtualMachine(val NestedVirtualMachineRequest) {
	m.VirtualMachine = val
}

// GetVrf returns the Vrf property
func (m VMInterfaceRequest) GetVrf() *NestedVRFRequest {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *VMInterfaceRequest) SetVrf(val *NestedVRFRequest) {
	m.Vrf = val
}
