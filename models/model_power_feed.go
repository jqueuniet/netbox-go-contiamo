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

// PowerFeed is an object. Adds support for custom fields and tags.
type PowerFeed struct {
	// Occupied:
	Occupied bool `json:"_occupied" mapstructure:"_occupied"`
	// Amperage:
	Amperage int32 `json:"amperage,omitempty" mapstructure:"amperage,omitempty"`
	// Cable: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Cable *NestedCable `json:"cable" mapstructure:"cable"`
	// CableEnd:
	CableEnd string `json:"cable_end" mapstructure:"cable_end"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
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
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// LinkPeers:
	LinkPeers []interface{} `json:"link_peers" mapstructure:"link_peers"`
	// LinkPeersType: Return the type of the peer link terminations, or None.
	LinkPeersType string `json:"link_peers_type" mapstructure:"link_peers_type"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// MaxUtilization: Maximum permissible draw (percentage)
	MaxUtilization int32 `json:"max_utilization,omitempty" mapstructure:"max_utilization,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Phase:
	Phase *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"phase,omitempty" mapstructure:"phase,omitempty"`
	// PowerPanel: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	PowerPanel NestedPowerPanel `json:"power_panel" mapstructure:"power_panel"`
	// Rack: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Rack *NestedRack `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// Status:
	Status *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Supply:
	Supply *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"supply,omitempty" mapstructure:"supply,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type:
	Type *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type,omitempty" mapstructure:"type,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Voltage:
	Voltage int32 `json:"voltage,omitempty" mapstructure:"voltage,omitempty"`
}

// Validate implements basic validation for this model
func (m PowerFeed) Validate() error {
	return validation.Errors{
		"amperage": validation.Validate(
			m.Amperage, validation.Min(int32(1)), validation.Max(int32(32767)),
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
		"linkPeers": validation.Validate(
			m.LinkPeers, validation.NotNil,
		),
		"maxUtilization": validation.Validate(
			m.MaxUtilization, validation.Min(int32(1)), validation.Max(int32(100)),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"powerPanel": validation.Validate(
			m.PowerPanel, validation.NotNil,
		),
		"rack": validation.Validate(
			m.Rack,
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"voltage": validation.Validate(
			m.Voltage, validation.Min(int32(-32768)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetOccupied returns the Occupied property
func (m PowerFeed) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *PowerFeed) SetOccupied(val bool) {
	m.Occupied = val
}

// GetAmperage returns the Amperage property
func (m PowerFeed) GetAmperage() int32 {
	return m.Amperage
}

// SetAmperage sets the Amperage property
func (m *PowerFeed) SetAmperage(val int32) {
	m.Amperage = val
}

// GetCable returns the Cable property
func (m PowerFeed) GetCable() *NestedCable {
	return m.Cable
}

// SetCable sets the Cable property
func (m *PowerFeed) SetCable(val *NestedCable) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m PowerFeed) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *PowerFeed) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetComments returns the Comments property
func (m PowerFeed) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PowerFeed) SetComments(val string) {
	m.Comments = val
}

// GetConnectedEndpoints returns the ConnectedEndpoints property
func (m PowerFeed) GetConnectedEndpoints() []interface{} {
	return m.ConnectedEndpoints
}

// SetConnectedEndpoints sets the ConnectedEndpoints property
func (m *PowerFeed) SetConnectedEndpoints(val []interface{}) {
	m.ConnectedEndpoints = val
}

// GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable property
func (m PowerFeed) GetConnectedEndpointsReachable() bool {
	return m.ConnectedEndpointsReachable
}

// SetConnectedEndpointsReachable sets the ConnectedEndpointsReachable property
func (m *PowerFeed) SetConnectedEndpointsReachable(val bool) {
	m.ConnectedEndpointsReachable = val
}

// GetConnectedEndpointsType returns the ConnectedEndpointsType property
func (m PowerFeed) GetConnectedEndpointsType() string {
	return m.ConnectedEndpointsType
}

// SetConnectedEndpointsType sets the ConnectedEndpointsType property
func (m *PowerFeed) SetConnectedEndpointsType(val string) {
	m.ConnectedEndpointsType = val
}

// GetCreated returns the Created property
func (m PowerFeed) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *PowerFeed) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m PowerFeed) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PowerFeed) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PowerFeed) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PowerFeed) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m PowerFeed) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *PowerFeed) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m PowerFeed) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *PowerFeed) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m PowerFeed) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *PowerFeed) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLinkPeers returns the LinkPeers property
func (m PowerFeed) GetLinkPeers() []interface{} {
	return m.LinkPeers
}

// SetLinkPeers sets the LinkPeers property
func (m *PowerFeed) SetLinkPeers(val []interface{}) {
	m.LinkPeers = val
}

// GetLinkPeersType returns the LinkPeersType property
func (m PowerFeed) GetLinkPeersType() string {
	return m.LinkPeersType
}

// SetLinkPeersType sets the LinkPeersType property
func (m *PowerFeed) SetLinkPeersType(val string) {
	m.LinkPeersType = val
}

// GetMarkConnected returns the MarkConnected property
func (m PowerFeed) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *PowerFeed) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetMaxUtilization returns the MaxUtilization property
func (m PowerFeed) GetMaxUtilization() int32 {
	return m.MaxUtilization
}

// SetMaxUtilization sets the MaxUtilization property
func (m *PowerFeed) SetMaxUtilization(val int32) {
	m.MaxUtilization = val
}

// GetName returns the Name property
func (m PowerFeed) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PowerFeed) SetName(val string) {
	m.Name = val
}

// GetPhase returns the Phase property
func (m PowerFeed) GetPhase() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Phase
}

// SetPhase sets the Phase property
func (m *PowerFeed) SetPhase(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Phase = val
}

// GetPowerPanel returns the PowerPanel property
func (m PowerFeed) GetPowerPanel() NestedPowerPanel {
	return m.PowerPanel
}

// SetPowerPanel sets the PowerPanel property
func (m *PowerFeed) SetPowerPanel(val NestedPowerPanel) {
	m.PowerPanel = val
}

// GetRack returns the Rack property
func (m PowerFeed) GetRack() *NestedRack {
	return m.Rack
}

// SetRack sets the Rack property
func (m *PowerFeed) SetRack(val *NestedRack) {
	m.Rack = val
}

// GetStatus returns the Status property
func (m PowerFeed) GetStatus() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *PowerFeed) SetStatus(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetSupply returns the Supply property
func (m PowerFeed) GetSupply() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Supply
}

// SetSupply sets the Supply property
func (m *PowerFeed) SetSupply(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Supply = val
}

// GetTags returns the Tags property
func (m PowerFeed) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PowerFeed) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetType returns the Type property
func (m PowerFeed) GetType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *PowerFeed) SetType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m PowerFeed) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *PowerFeed) SetUrl(val string) {
	m.Url = val
}

// GetVoltage returns the Voltage property
func (m PowerFeed) GetVoltage() int32 {
	return m.Voltage
}

// SetVoltage sets the Voltage property
func (m *PowerFeed) SetVoltage(val int32) {
	m.Voltage = val
}
