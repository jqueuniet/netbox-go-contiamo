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

// WritableVMInterfaceRequest is an object. Adds support for custom fields and tags.
type WritableVMInterfaceRequest struct {
	// Bridge:
	Bridge *int32 `json:"bridge,omitempty" mapstructure:"bridge,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// MacAddress:
	MacAddress *string `json:"mac_address,omitempty" mapstructure:"mac_address,omitempty"`
	// Mode: IEEE 802.1Q tagging strategy
	//
	// * `access` - Access
	// * `tagged` - Tagged
	// * `tagged-all` - Tagged (All)
	Mode string `json:"mode,omitempty" mapstructure:"mode,omitempty"`
	// Mtu:
	Mtu *int32 `json:"mtu,omitempty" mapstructure:"mtu,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Parent:
	Parent *int32 `json:"parent,omitempty" mapstructure:"parent,omitempty"`
	// TaggedVlans:
	TaggedVlans []int32 `json:"tagged_vlans,omitempty" mapstructure:"tagged_vlans,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// UntaggedVlan:
	UntaggedVlan *int32 `json:"untagged_vlan,omitempty" mapstructure:"untagged_vlan,omitempty"`
	// VirtualMachine:
	VirtualMachine int32 `json:"virtual_machine" mapstructure:"virtual_machine"`
	// Vrf:
	Vrf *int32 `json:"vrf,omitempty" mapstructure:"vrf,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableVMInterfaceRequest) Validate() error {
	return validation.Errors{
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
		"taggedVlans": validation.Validate(
			m.TaggedVlans,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetBridge returns the Bridge property
func (m WritableVMInterfaceRequest) GetBridge() *int32 {
	return m.Bridge
}

// SetBridge sets the Bridge property
func (m *WritableVMInterfaceRequest) SetBridge(val *int32) {
	m.Bridge = val
}

// GetCustomFields returns the CustomFields property
func (m WritableVMInterfaceRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableVMInterfaceRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableVMInterfaceRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableVMInterfaceRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnabled returns the Enabled property
func (m WritableVMInterfaceRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *WritableVMInterfaceRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetMacAddress returns the MacAddress property
func (m WritableVMInterfaceRequest) GetMacAddress() *string {
	return m.MacAddress
}

// SetMacAddress sets the MacAddress property
func (m *WritableVMInterfaceRequest) SetMacAddress(val *string) {
	m.MacAddress = val
}

// GetMode returns the Mode property
func (m WritableVMInterfaceRequest) GetMode() string {
	return m.Mode
}

// SetMode sets the Mode property
func (m *WritableVMInterfaceRequest) SetMode(val string) {
	m.Mode = val
}

// GetMtu returns the Mtu property
func (m WritableVMInterfaceRequest) GetMtu() *int32 {
	return m.Mtu
}

// SetMtu sets the Mtu property
func (m *WritableVMInterfaceRequest) SetMtu(val *int32) {
	m.Mtu = val
}

// GetName returns the Name property
func (m WritableVMInterfaceRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableVMInterfaceRequest) SetName(val string) {
	m.Name = val
}

// GetParent returns the Parent property
func (m WritableVMInterfaceRequest) GetParent() *int32 {
	return m.Parent
}

// SetParent sets the Parent property
func (m *WritableVMInterfaceRequest) SetParent(val *int32) {
	m.Parent = val
}

// GetTaggedVlans returns the TaggedVlans property
func (m WritableVMInterfaceRequest) GetTaggedVlans() []int32 {
	return m.TaggedVlans
}

// SetTaggedVlans sets the TaggedVlans property
func (m *WritableVMInterfaceRequest) SetTaggedVlans(val []int32) {
	m.TaggedVlans = val
}

// GetTags returns the Tags property
func (m WritableVMInterfaceRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableVMInterfaceRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetUntaggedVlan returns the UntaggedVlan property
func (m WritableVMInterfaceRequest) GetUntaggedVlan() *int32 {
	return m.UntaggedVlan
}

// SetUntaggedVlan sets the UntaggedVlan property
func (m *WritableVMInterfaceRequest) SetUntaggedVlan(val *int32) {
	m.UntaggedVlan = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m WritableVMInterfaceRequest) GetVirtualMachine() int32 {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *WritableVMInterfaceRequest) SetVirtualMachine(val int32) {
	m.VirtualMachine = val
}

// GetVrf returns the Vrf property
func (m WritableVMInterfaceRequest) GetVrf() *int32 {
	return m.Vrf
}

// SetVrf sets the Vrf property
func (m *WritableVMInterfaceRequest) SetVrf(val *int32) {
	m.Vrf = val
}
