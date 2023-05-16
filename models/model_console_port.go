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

// ConsolePort is an object. Adds support for custom fields and tags.
type ConsolePort struct {
	// Occupied:
	Occupied bool `json:"_occupied" mapstructure:"_occupied"`
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
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// LinkPeers:
	LinkPeers []interface{} `json:"link_peers" mapstructure:"link_peers"`
	// LinkPeersType: Return the type of the peer link terminations, or None.
	LinkPeersType string `json:"link_peers_type" mapstructure:"link_peers_type"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// Module: Used by device component serializers.
	Module *ComponentNestedModule `json:"module,omitempty" mapstructure:"module,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Speed:
	Speed *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"speed,omitempty" mapstructure:"speed,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type:
	Type *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type,omitempty" mapstructure:"type,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ConsolePort) Validate() error {
	return validation.Errors{
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
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"linkPeers": validation.Validate(
			m.LinkPeers, validation.NotNil,
		),
		"module": validation.Validate(
			m.Module,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetOccupied returns the Occupied property
func (m ConsolePort) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *ConsolePort) SetOccupied(val bool) {
	m.Occupied = val
}

// GetCable returns the Cable property
func (m ConsolePort) GetCable() *NestedCable {
	return m.Cable
}

// SetCable sets the Cable property
func (m *ConsolePort) SetCable(val *NestedCable) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m ConsolePort) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *ConsolePort) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetConnectedEndpoints returns the ConnectedEndpoints property
func (m ConsolePort) GetConnectedEndpoints() []interface{} {
	return m.ConnectedEndpoints
}

// SetConnectedEndpoints sets the ConnectedEndpoints property
func (m *ConsolePort) SetConnectedEndpoints(val []interface{}) {
	m.ConnectedEndpoints = val
}

// GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable property
func (m ConsolePort) GetConnectedEndpointsReachable() bool {
	return m.ConnectedEndpointsReachable
}

// SetConnectedEndpointsReachable sets the ConnectedEndpointsReachable property
func (m *ConsolePort) SetConnectedEndpointsReachable(val bool) {
	m.ConnectedEndpointsReachable = val
}

// GetConnectedEndpointsType returns the ConnectedEndpointsType property
func (m ConsolePort) GetConnectedEndpointsType() string {
	return m.ConnectedEndpointsType
}

// SetConnectedEndpointsType sets the ConnectedEndpointsType property
func (m *ConsolePort) SetConnectedEndpointsType(val string) {
	m.ConnectedEndpointsType = val
}

// GetCreated returns the Created property
func (m ConsolePort) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ConsolePort) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m ConsolePort) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ConsolePort) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ConsolePort) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ConsolePort) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m ConsolePort) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *ConsolePort) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m ConsolePort) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ConsolePort) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ConsolePort) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ConsolePort) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m ConsolePort) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *ConsolePort) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m ConsolePort) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ConsolePort) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLinkPeers returns the LinkPeers property
func (m ConsolePort) GetLinkPeers() []interface{} {
	return m.LinkPeers
}

// SetLinkPeers sets the LinkPeers property
func (m *ConsolePort) SetLinkPeers(val []interface{}) {
	m.LinkPeers = val
}

// GetLinkPeersType returns the LinkPeersType property
func (m ConsolePort) GetLinkPeersType() string {
	return m.LinkPeersType
}

// SetLinkPeersType sets the LinkPeersType property
func (m *ConsolePort) SetLinkPeersType(val string) {
	m.LinkPeersType = val
}

// GetMarkConnected returns the MarkConnected property
func (m ConsolePort) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *ConsolePort) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetModule returns the Module property
func (m ConsolePort) GetModule() *ComponentNestedModule {
	return m.Module
}

// SetModule sets the Module property
func (m *ConsolePort) SetModule(val *ComponentNestedModule) {
	m.Module = val
}

// GetName returns the Name property
func (m ConsolePort) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ConsolePort) SetName(val string) {
	m.Name = val
}

// GetSpeed returns the Speed property
func (m ConsolePort) GetSpeed() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Speed
}

// SetSpeed sets the Speed property
func (m *ConsolePort) SetSpeed(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Speed = val
}

// GetTags returns the Tags property
func (m ConsolePort) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ConsolePort) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetType returns the Type property
func (m ConsolePort) GetType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *ConsolePort) SetType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m ConsolePort) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ConsolePort) SetUrl(val string) {
	m.Url = val
}
