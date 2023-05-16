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

// PowerOutlet is an object. Adds support for custom fields and tags.
type PowerOutlet struct {
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
	// FeedLeg:
	FeedLeg *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"feed_leg,omitempty" mapstructure:"feed_leg,omitempty"`
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
	// PowerPort: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PowerPort *NestedPowerPort `json:"power_port,omitempty" mapstructure:"power_port,omitempty"`
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
func (m PowerOutlet) Validate() error {
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
		"powerPort": validation.Validate(
			m.PowerPort,
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
func (m PowerOutlet) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *PowerOutlet) SetOccupied(val bool) {
	m.Occupied = val
}

// GetCable returns the Cable property
func (m PowerOutlet) GetCable() *NestedCable {
	return m.Cable
}

// SetCable sets the Cable property
func (m *PowerOutlet) SetCable(val *NestedCable) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m PowerOutlet) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *PowerOutlet) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetConnectedEndpoints returns the ConnectedEndpoints property
func (m PowerOutlet) GetConnectedEndpoints() []interface{} {
	return m.ConnectedEndpoints
}

// SetConnectedEndpoints sets the ConnectedEndpoints property
func (m *PowerOutlet) SetConnectedEndpoints(val []interface{}) {
	m.ConnectedEndpoints = val
}

// GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable property
func (m PowerOutlet) GetConnectedEndpointsReachable() bool {
	return m.ConnectedEndpointsReachable
}

// SetConnectedEndpointsReachable sets the ConnectedEndpointsReachable property
func (m *PowerOutlet) SetConnectedEndpointsReachable(val bool) {
	m.ConnectedEndpointsReachable = val
}

// GetConnectedEndpointsType returns the ConnectedEndpointsType property
func (m PowerOutlet) GetConnectedEndpointsType() string {
	return m.ConnectedEndpointsType
}

// SetConnectedEndpointsType sets the ConnectedEndpointsType property
func (m *PowerOutlet) SetConnectedEndpointsType(val string) {
	m.ConnectedEndpointsType = val
}

// GetCreated returns the Created property
func (m PowerOutlet) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *PowerOutlet) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m PowerOutlet) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PowerOutlet) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PowerOutlet) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PowerOutlet) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m PowerOutlet) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *PowerOutlet) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m PowerOutlet) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *PowerOutlet) SetDisplay(val string) {
	m.Display = val
}

// GetFeedLeg returns the FeedLeg property
func (m PowerOutlet) GetFeedLeg() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.FeedLeg
}

// SetFeedLeg sets the FeedLeg property
func (m *PowerOutlet) SetFeedLeg(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.FeedLeg = val
}

// GetId returns the Id property
func (m PowerOutlet) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *PowerOutlet) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m PowerOutlet) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *PowerOutlet) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m PowerOutlet) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *PowerOutlet) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLinkPeers returns the LinkPeers property
func (m PowerOutlet) GetLinkPeers() []interface{} {
	return m.LinkPeers
}

// SetLinkPeers sets the LinkPeers property
func (m *PowerOutlet) SetLinkPeers(val []interface{}) {
	m.LinkPeers = val
}

// GetLinkPeersType returns the LinkPeersType property
func (m PowerOutlet) GetLinkPeersType() string {
	return m.LinkPeersType
}

// SetLinkPeersType sets the LinkPeersType property
func (m *PowerOutlet) SetLinkPeersType(val string) {
	m.LinkPeersType = val
}

// GetMarkConnected returns the MarkConnected property
func (m PowerOutlet) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *PowerOutlet) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetModule returns the Module property
func (m PowerOutlet) GetModule() *ComponentNestedModule {
	return m.Module
}

// SetModule sets the Module property
func (m *PowerOutlet) SetModule(val *ComponentNestedModule) {
	m.Module = val
}

// GetName returns the Name property
func (m PowerOutlet) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PowerOutlet) SetName(val string) {
	m.Name = val
}

// GetPowerPort returns the PowerPort property
func (m PowerOutlet) GetPowerPort() *NestedPowerPort {
	return m.PowerPort
}

// SetPowerPort sets the PowerPort property
func (m *PowerOutlet) SetPowerPort(val *NestedPowerPort) {
	m.PowerPort = val
}

// GetTags returns the Tags property
func (m PowerOutlet) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PowerOutlet) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetType returns the Type property
func (m PowerOutlet) GetType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *PowerOutlet) SetType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m PowerOutlet) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *PowerOutlet) SetUrl(val string) {
	m.Url = val
}
