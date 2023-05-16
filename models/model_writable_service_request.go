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

// WritableServiceRequest is an object. Adds support for custom fields and tags.
type WritableServiceRequest struct {
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
	Name string `json:"name" mapstructure:"name"`
	// Ports:
	Ports []int32 `json:"ports" mapstructure:"ports"`
	// Protocol: * `tcp` - TCP
	// * `udp` - UDP
	// * `sctp` - SCTP
	Protocol string `json:"protocol" mapstructure:"protocol"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// VirtualMachine:
	VirtualMachine *int32 `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableServiceRequest) Validate() error {
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
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"ports": validation.Validate(
			m.Ports, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m WritableServiceRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableServiceRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableServiceRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableServiceRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableServiceRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableServiceRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m WritableServiceRequest) GetDevice() *int32 {
	return m.Device
}

// SetDevice sets the Device property
func (m *WritableServiceRequest) SetDevice(val *int32) {
	m.Device = val
}

// GetIpaddresses returns the Ipaddresses property
func (m WritableServiceRequest) GetIpaddresses() []int32 {
	return m.Ipaddresses
}

// SetIpaddresses sets the Ipaddresses property
func (m *WritableServiceRequest) SetIpaddresses(val []int32) {
	m.Ipaddresses = val
}

// GetName returns the Name property
func (m WritableServiceRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableServiceRequest) SetName(val string) {
	m.Name = val
}

// GetPorts returns the Ports property
func (m WritableServiceRequest) GetPorts() []int32 {
	return m.Ports
}

// SetPorts sets the Ports property
func (m *WritableServiceRequest) SetPorts(val []int32) {
	m.Ports = val
}

// GetProtocol returns the Protocol property
func (m WritableServiceRequest) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *WritableServiceRequest) SetProtocol(val string) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m WritableServiceRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableServiceRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m WritableServiceRequest) GetVirtualMachine() *int32 {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *WritableServiceRequest) SetVirtualMachine(val *int32) {
	m.VirtualMachine = val
}
