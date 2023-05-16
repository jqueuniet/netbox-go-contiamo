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

// Service is an object. Adds support for custom fields and tags.
type Service struct {
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Device: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Device *NestedDevice `json:"device,omitempty" mapstructure:"device,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Ipaddresses:
	Ipaddresses []int32 `json:"ipaddresses,omitempty" mapstructure:"ipaddresses,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Ports:
	Ports []int32 `json:"ports" mapstructure:"ports"`
	// Protocol:
	Protocol *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"protocol,omitempty" mapstructure:"protocol,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// VirtualMachine: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	VirtualMachine *NestedVirtualMachine `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
}

// Validate implements basic validation for this model
func (m Service) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"ports": validation.Validate(
			m.Ports, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"virtualMachine": validation.Validate(
			m.VirtualMachine,
		),
	}.Filter()
}

// GetComments returns the Comments property
func (m Service) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Service) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Service) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Service) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Service) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Service) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Service) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Service) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m Service) GetDevice() *NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *Service) SetDevice(val *NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m Service) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Service) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Service) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Service) SetId(val int32) {
	m.Id = val
}

// GetIpaddresses returns the Ipaddresses property
func (m Service) GetIpaddresses() []int32 {
	return m.Ipaddresses
}

// SetIpaddresses sets the Ipaddresses property
func (m *Service) SetIpaddresses(val []int32) {
	m.Ipaddresses = val
}

// GetLastUpdated returns the LastUpdated property
func (m Service) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Service) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Service) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Service) SetName(val string) {
	m.Name = val
}

// GetPorts returns the Ports property
func (m Service) GetPorts() []int32 {
	return m.Ports
}

// SetPorts sets the Ports property
func (m *Service) SetPorts(val []int32) {
	m.Ports = val
}

// GetProtocol returns the Protocol property
func (m Service) GetProtocol() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *Service) SetProtocol(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Protocol = val
}

// GetTags returns the Tags property
func (m Service) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Service) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetUrl returns the Url property
func (m Service) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Service) SetUrl(val string) {
	m.Url = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m Service) GetVirtualMachine() *NestedVirtualMachine {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *Service) SetVirtualMachine(val *NestedVirtualMachine) {
	m.VirtualMachine = val
}
