// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"

	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// frontPortColorPattern is the validation pattern for FrontPort.Color
var frontPortColorPattern = regexp.MustCompile(`^[0-9a-f]{6}$`)

// FrontPort is an object. Adds support for custom fields and tags.
type FrontPort struct {
	// Occupied:
	Occupied bool `json:"_occupied" mapstructure:"_occupied"`
	// Cable: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Cable *NestedCable `json:"cable" mapstructure:"cable"`
	// CableEnd:
	CableEnd string `json:"cable_end" mapstructure:"cable_end"`
	// Color:
	Color string `json:"color,omitempty" mapstructure:"color,omitempty"`
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
	// RearPort: NestedRearPortSerializer but with parent device omitted (since front and rear ports must belong to same device)
	RearPort FrontPortRearPort `json:"rear_port" mapstructure:"rear_port"`
	// RearPortPosition: Mapped position on corresponding rear port
	RearPortPosition int32 `json:"rear_port_position,omitempty" mapstructure:"rear_port_position,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type:
	Type struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m FrontPort) Validate() error {
	return validation.Errors{
		"cable": validation.Validate(
			m.Cable,
		),
		"color": validation.Validate(
			m.Color, validation.Length(0, 6), validation.Match(frontPortColorPattern),
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
		"rearPort": validation.Validate(
			m.RearPort, validation.NotNil,
		),
		"rearPortPosition": validation.Validate(
			m.RearPortPosition, validation.Min(int32(1)), validation.Max(int32(1024)),
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
func (m FrontPort) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *FrontPort) SetOccupied(val bool) {
	m.Occupied = val
}

// GetCable returns the Cable property
func (m FrontPort) GetCable() *NestedCable {
	return m.Cable
}

// SetCable sets the Cable property
func (m *FrontPort) SetCable(val *NestedCable) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m FrontPort) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *FrontPort) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetColor returns the Color property
func (m FrontPort) GetColor() string {
	return m.Color
}

// SetColor sets the Color property
func (m *FrontPort) SetColor(val string) {
	m.Color = val
}

// GetCreated returns the Created property
func (m FrontPort) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *FrontPort) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m FrontPort) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *FrontPort) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m FrontPort) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *FrontPort) SetDescription(val string) {
	m.Description = val
}

// GetDevice returns the Device property
func (m FrontPort) GetDevice() NestedDevice {
	return m.Device
}

// SetDevice sets the Device property
func (m *FrontPort) SetDevice(val NestedDevice) {
	m.Device = val
}

// GetDisplay returns the Display property
func (m FrontPort) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *FrontPort) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m FrontPort) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *FrontPort) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m FrontPort) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *FrontPort) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m FrontPort) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *FrontPort) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLinkPeers returns the LinkPeers property
func (m FrontPort) GetLinkPeers() []interface{} {
	return m.LinkPeers
}

// SetLinkPeers sets the LinkPeers property
func (m *FrontPort) SetLinkPeers(val []interface{}) {
	m.LinkPeers = val
}

// GetLinkPeersType returns the LinkPeersType property
func (m FrontPort) GetLinkPeersType() string {
	return m.LinkPeersType
}

// SetLinkPeersType sets the LinkPeersType property
func (m *FrontPort) SetLinkPeersType(val string) {
	m.LinkPeersType = val
}

// GetMarkConnected returns the MarkConnected property
func (m FrontPort) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *FrontPort) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetModule returns the Module property
func (m FrontPort) GetModule() *ComponentNestedModule {
	return m.Module
}

// SetModule sets the Module property
func (m *FrontPort) SetModule(val *ComponentNestedModule) {
	m.Module = val
}

// GetName returns the Name property
func (m FrontPort) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *FrontPort) SetName(val string) {
	m.Name = val
}

// GetRearPort returns the RearPort property
func (m FrontPort) GetRearPort() FrontPortRearPort {
	return m.RearPort
}

// SetRearPort sets the RearPort property
func (m *FrontPort) SetRearPort(val FrontPortRearPort) {
	m.RearPort = val
}

// GetRearPortPosition returns the RearPortPosition property
func (m FrontPort) GetRearPortPosition() int32 {
	return m.RearPortPosition
}

// SetRearPortPosition sets the RearPortPosition property
func (m *FrontPort) SetRearPortPosition(val int32) {
	m.RearPortPosition = val
}

// GetTags returns the Tags property
func (m FrontPort) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *FrontPort) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetType returns the Type property
func (m FrontPort) GetType() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *FrontPort) SetType(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m FrontPort) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *FrontPort) SetUrl(val string) {
	m.Url = val
}
