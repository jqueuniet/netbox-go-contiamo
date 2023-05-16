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

// PatchedWritableServiceRequest is an object. Adds support for custom fields and tags.
type PatchedWritableServiceRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device:
	Device *int32 `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Ipaddresses: The specific IP addresses (if any) to which this service is bound
	Ipaddresses []int32 `json:"ipaddresses,omitempty" mapstructure:"ipaddresses,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Ports:
	Ports []int32 `json:"ports,omitempty" mapstructure:"ports,omitempty"`
	// Protocol: * `tcp` - TCP
	// * `udp` - UDP
	// * `sctp` - SCTP
	Protocol string `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// VirtualMachine:
	VirtualMachine *int32 `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableServiceRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"ipaddresses": validation.Validate(
			m.Ipaddresses,
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"ports": validation.Validate(
			m.Ports,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m PatchedWritableServiceRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableServiceRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableServiceRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableServiceRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableServiceRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableServiceRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PatchedWritableServiceRequest) GetDevice() *int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *PatchedWritableServiceRequest) SetDevice(val *int32) {
	m.Device = val
}

// GetIpaddresses returns the Ipaddresses property
func (m PatchedWritableServiceRequest) GetIpaddresses() []int32 {
	return m.Ipaddresses
}

// SetIpaddresses sets the Ipaddresses property
func (m *PatchedWritableServiceRequest) SetIpaddresses(val []int32) {
	m.Ipaddresses = val
}

// GetName returns the Name property
func (m PatchedWritableServiceRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableServiceRequest) SetName(val string) {
	m.Name = val
}

// GetPorts returns the Ports property
func (m PatchedWritableServiceRequest) GetPorts() []int32 {
	return m.Ports
}

// SetPorts sets the Ports property
func (m *PatchedWritableServiceRequest) SetPorts(val []int32) {
	m.Ports = val
}

// GetProtocol returns the Protocol property
func (m PatchedWritableServiceRequest) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *PatchedWritableServiceRequest) SetProtocol(val string) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m PatchedWritableServiceRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableServiceRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m PatchedWritableServiceRequest) GetVirtualMachine() *int32 {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *PatchedWritableServiceRequest) SetVirtualMachine(val *int32) {
	m.VirtualMachine = val
}
