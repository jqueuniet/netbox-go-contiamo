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

// ServiceRequest is an object. Adds support for custom fields and tags.
type ServiceRequest struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device *NestedDeviceRequest `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Ipaddresses:
	Ipaddresses []int32 `json:"ipaddresses,omitempty" mapstructure:"ipaddresses,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Ports:
	Ports []int32 `json:"ports" mapstructure:"ports"`
	// Protocol: * `tcp` - TCP
	// * `udp` - UDP
	// * `sctp` - SCTP
	Protocol string `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// VirtualMachine: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	VirtualMachine *NestedVirtualMachineRequest `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
}

// Validate implements basic validation for this model
func (m ServiceRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"device": validation.Validate(
			m.Device,
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
		"virtualMachine": validation.Validate(
			m.VirtualMachine,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m ServiceRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ServiceRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m ServiceRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ServiceRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ServiceRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ServiceRequest) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m ServiceRequest) GetDevice() *NestedDeviceRequest {
	return m.Device
}

// SetDevice sets the Device property
func (m *ServiceRequest) SetDevice(val *NestedDeviceRequest) {
	m.Device = val
}

// GetIpaddresses returns the Ipaddresses property
func (m ServiceRequest) GetIpaddresses() []int32 {
	return m.Ipaddresses
}

// SetIpaddresses sets the Ipaddresses property
func (m *ServiceRequest) SetIpaddresses(val []int32) {
	m.Ipaddresses = val
}

// GetName returns the Name property
func (m ServiceRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ServiceRequest) SetName(val string) {
	m.Name = val
}

// GetPorts returns the Ports property
func (m ServiceRequest) GetPorts() []int32 {
	return m.Ports
}

// SetPorts sets the Ports property
func (m *ServiceRequest) SetPorts(val []int32) {
	m.Ports = val
}

// GetProtocol returns the Protocol property
func (m ServiceRequest) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *ServiceRequest) SetProtocol(val string) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m ServiceRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ServiceRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m ServiceRequest) GetVirtualMachine() *NestedVirtualMachineRequest {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *ServiceRequest) SetVirtualMachine(val *NestedVirtualMachineRequest) {
	m.VirtualMachine = val
}
